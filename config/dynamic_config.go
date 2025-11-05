package config

import (
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
	"strconv"
	"strings"
	"sync"
)

type DynamicConfig struct {
	keys     []string
	listener *viper.Viper
	reader   sync.Map
}

func (*DynamicConfig) GetString(key string) (string, error) {
	val, ok := dynamicConfig.reader.Load(key)
	if !ok {
		return "", errors.New("key not found in dynamic config: " + key)
	}
	return cast.ToString(val), nil
}

func (*DynamicConfig) GetBoolWithFallback(key string, fallback bool) bool {
	val, ok := dynamicConfig.reader.Load(key)
	if !ok || val == nil {
		return fallback
	}
	return cast.ToBool(val)
}

func (*DynamicConfig) GetStringWithFallback(key string, fallback string) string {
	val, ok := dynamicConfig.reader.Load(key)
	if !ok || val == nil {
		return fallback
	}
	return cast.ToString(val)
}

func (*DynamicConfig) GetInt64WithFallback(key string, fallback int64) int64 {
	val, ok := dynamicConfig.reader.Load(key)
	if !ok || val == nil {
		return fallback
	}
	return cast.ToInt64(val)
}

func (*DynamicConfig) GetFloat64WithFallback(key string, fallback float64) float64 {
	val, ok := dynamicConfig.reader.Load(key)
	if !ok || val == nil {
		return fallback
	}
	return cast.ToFloat64(val)
}

func (*DynamicConfig) GetStringArrayWithFallback(key string, fallback []string) []string {
	val, ok := dynamicConfig.reader.Load(key)

	if !ok || val == nil {
		return fallback
	}
	stringArray := strings.Split(cast.ToString(val), ",")

	for i, str := range stringArray {
		stringArray[i] = strings.TrimSpace(str)
	}
	return stringArray
}

func (*DynamicConfig) getIntSliceWithFallback(key string, fallback []int) []int {
	val, ok := dynamicConfig.reader.Load(key)
	if !ok || val == nil {
		return fallback
	}
	intSlice := cast.ToIntSlice(val)
	return intSlice
}

func (dc *DynamicConfig) getIntSliceStringWithFallback(key string, fallback []int) []int {
	val, ok := dc.reader.Load(key)
	if !ok || val == nil {
		return fallback
	}

	// Convert to string and split by comma
	strVal := cast.ToString(val)
	if strVal == "" {
		return fallback
	}

	parts := strings.Split(strVal, ",")
	result := make([]int, 0, len(parts))

	for _, part := range parts {
		numStr := strings.TrimSpace(part)
		if numStr == "" {
			continue
		}
		num, err := strconv.Atoi(numStr)
		if err != nil {
			// If any entry is invalid, fall back
			return fallback
		}
		result = append(result, num)
	}

	return result
}

func (*DynamicConfig) GetStringMapBoolWithFallback(key string, fallback map[string]bool) map[string]bool {
	val, ok := dynamicConfig.reader.Load(key)
	if !ok || val == nil {
		return fallback
	}
	return cast.ToStringMapBool(val)
}

func (*DynamicConfig) GetStringMapStringWithFallback(key string, fallback map[string]string) map[string]string {
	val, ok := dynamicConfig.reader.Load(key)
	if !ok || val == nil {
		return fallback
	}
	return cast.ToStringMapString(val)
}

func (*DynamicConfig) GetStringMapWithFallback(key string, fallback map[string]interface{}) map[string]interface{} {
	val, ok := dynamicConfig.reader.Load(key)
	if !ok || val == nil {
		return fallback
	}
	return cast.ToStringMap(val)
}

func getFromString[T any](key string, out *T) *T {
	val, ok := dynamicConfig.reader.Load(key)
	if !ok || val == nil {
		return new(T)
	}
	valString := cast.ToString(val)
	err := sonic.Unmarshal([]byte(valString), out)
	if err != nil {
		return new(T)
	}
	return out
}

func (*DynamicConfig) SyncDynamicConfigsToReader() {
	for _, key := range dynamicConfig.keys {
		dynamicConfig.reader.Store(key, dynamicConfig.listener.Get(key))
		fmt.Printf("Dynamic config key synced key: %+v, value : %+v\n", key, dynamicConfig.listener.Get(key))
	}
}

func getAllDynamicConfigKeys() []string {
	return []string{}
}
