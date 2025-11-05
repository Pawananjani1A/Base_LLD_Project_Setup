package config

import (
	"base_app/constants"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"
)

var dynamicConfig = DynamicConfig{
	keys:     getAllDynamicConfigKeys(),
	listener: viper.New(),
	reader:   sync.Map{},
}

type Configuration struct {
	port         int64
	logLevel     string
	logLocation  string
	enableSentry bool
	sentryDsn    string
	authSecret   string
	kafkaConfig  KafkaConfig
	authSecretV2 []byte
	Grpc         GrpcConfig
	stage        string
	Bifrost      Bifrost
	RedisConfig  RedisConfig
	MongoConfig  MongoConfig
}

type Bifrost struct {
	Key    string
	NewKey string
}

type ExternalServicesHystrixConfig struct {
	TimeoutInMS       int
	MaxErrorPercent   int
	RecoveryCheckTime int
}

type UserAppUpdate struct {
	userAppUpdateWidget                  string
	userAppUpdateDeeplinkIos             string
	userAppUpdateDeeplinkAndroid         string
	userAppUpdateBannerId                string
	userAppUpdateAndroidVersionThreshold string
	userAppUpdateIosVersionThreshold     string
}

type KafkaTopicsMap struct {
	KafkaTopic string
}

type NewKafkaFlowEnabled struct {
}

type KafkaConfig struct {
	Brokers                   []string
	TopicsMap                 KafkaTopicsMap
	NewKafkaFlowEnabled       NewKafkaFlowEnabled
	GroupId                   string
	KafkaProducerRequiredAcks int
}

type GrpcConfig struct {
	Port int64
	Host string
}

type ClusterRedisConfig struct {
	NodeAddresses             []string
	PerNodePoolSize           int64
	PerNodeMinIdleConnections int64
	DialTimeout               time.Duration
	ReadTimeout               time.Duration
	WriteTimeout              time.Duration
	PerNodeMaxIdleConnections int64
}

type RedisConfig struct {
	Addr         string
	Password     string
	Db           int
	DialTimeout  time.Duration
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type MongoConfig struct {
	Uri            string
	Database       string
	ConnectTimeout time.Duration
	Username       string
	Password       string
	AuthSource     string
}

type ExternalServicesConfig struct {
}

type ExternalServiceHttpConfig struct {
	DNS         string
	TimeoutInMS time.Duration
}

type ExternalServiceGrpcConfig struct {
	AnalyticsDNS      string
	DNS               string
	RetryCount        int64
	RetryWaitTimeInMS time.Duration
	TimeoutInMS       time.Duration
	TLSEnabled        bool
}

var config *Configuration

type Error struct {
	Error error
}

func panicIfError(err error) {
	if err != nil {
		fmt.Println(err)
		panic(fmt.Errorf("unable to load config: %v", err))
	}
}

func checkKey(key string) {
	if !viper.IsSet(key) {
		panicIfError(fmt.Errorf("%s key is not set", key))
	}
}

func panicIfErrorForKey(err error, key string) {
	if err != nil {
		panicIfError(fmt.Errorf("could not parse key: %s. Error: %v", key, err))
	}
}

func getIntOrPanic(key string) int64 {
	checkKey(key)
	v, err := strconv.Atoi(viper.GetString(key))
	panicIfErrorForKey(err, key)
	return int64(v)
}

func getInt(key string, fallback int64) int64 {
	if !viper.IsSet(key) {
		return fallback
	}
	v, err := strconv.Atoi(viper.GetString(key))
	if err != nil {
		return fallback
	}
	return int64(v)
}

func getStringOrPanic(key string) string {
	checkKey(key)
	return viper.GetString(key)
}

func getString(key string, fallback string) string {
	if !viper.IsSet(key) {
		return fallback
	}
	return viper.GetString(key)
}

func getFloatOrPanic(key string) float64 {
	checkKey(key)
	v, err := strconv.ParseFloat(viper.GetString(key), 64)
	panicIfErrorForKey(err, key)
	return v
}

func getBoolOrPanic(key string) bool {
	if !viper.IsSet(key) {
		return false
	}

	v, err := strconv.ParseBool(viper.GetString(key))
	panicIfErrorForKey(err, key)
	return v
}

func getBool(key string, fallback bool) bool {
	if !viper.IsSet(key) {
		return fallback
	}

	v, err := strconv.ParseBool(viper.GetString(key))
	if err != nil {
		return fallback
	}
	return v
}

func getFloat(key string, fallback float64) float64 {
	if !viper.IsSet(key) {
		return fallback
	}

	v, err := strconv.ParseFloat(viper.GetString(key), 64)
	if err != nil {
		return fallback
	}
	return v
}

func getStringArray(key string) []string {
	stringArray := strings.Split(viper.GetString(key), ",")
	for i, str := range stringArray {
		stringArray[i] = strings.TrimSpace(str)
	}
	return stringArray
}

func getStringArrayWithFallback(key string, fallback []string) []string {
	if !viper.IsSet(key) {
		return fallback
	}
	stringArray := strings.Split(viper.GetString(key), ",")
	for i, str := range stringArray {
		stringArray[i] = strings.TrimSpace(str)
	}
	return stringArray
}

func getStringSliceOrPanic(key string) []string {
	checkKey(key)
	return viper.GetStringSlice(key)
}

func getIntSliceWithFallback(key string, fallback []int) []int {
	if !viper.IsSet(key) {
		return fallback
	}
	return getIntSliceOrPanic(key)
}

func getIntSliceOrPanic(key string) []int {
	checkKey(key)
	strArr := strings.Split(viper.GetString(key), ",")
	intArr := make([]int, len(strArr))

	for i, str := range strArr {
		v, err := strconv.Atoi(str)
		panicIfError(err)
		intArr[i] = v
	}

	return intArr
}

func GetMapOrDefault(key string, defaultValue map[string]string) map[string]string {
	if !viper.IsSet(key) {
		return defaultValue
	}
	var mapLoaded map[string]string
	rawString := dynamicConfig.GetStringWithFallback(key, "")
	rawString = strings.Replace(rawString, "\\", "", -1)
	err := json.Unmarshal([]byte(rawString), &mapLoaded)
	if err != nil {
		return nil
	}
	return mapLoaded
}

func getStringMapOrPanic(key string) map[string]string {
	checkKey(key)
	return viper.GetStringMapString(key)
}

func unescapeJsonStringForKey(key string, fallback string) string {
	jsonString := dynamicConfig.GetStringWithFallback(key, fallback)
	jsonString = strings.Replace(jsonString, "\\", "", -1)
	return jsonString
}

func SetConfigFileFromArgs(commandArgs []string) {
	if len(commandArgs) > 2 {
		viper.SetConfigName(commandArgs[2])
	} else {
		logrus.Fatal("Failed to startup server, please provide a config file name")
	}
}

func SetConfigFileFromFilePath(filepath string) {
	viper.SetConfigName(filepath)
}

func LoadWithoutStoreSdk(commandArgs []string) {
	var staticFilePath, dynamicFilePath string
	if len(commandArgs) > 3 {
		staticFilePath = commandArgs[2]

		dynamicFilePath = commandArgs[3]
	} else {
		logrus.Fatal("Failed to startup server, please provide a config file name")
	}

	viper.SetDefault("log_level", "debug")
	viper.AutomaticEnv()

	viper.AddConfigPath("./")
	viper.AddConfigPath("../")
	viper.AddConfigPath("../..")
	viper.AddConfigPath("./profiles")
	viper.AddConfigPath("/vault/secrets/")
	viper.SetConfigType("yaml")

	viper.SetConfigName(staticFilePath)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	//Reading Dynamic Configs
	dynamicConfig.listener.SetDefault("log_level", "debug")
	dynamicConfig.listener.AutomaticEnv()
	dynamicConfig.listener.AddConfigPath("./")
	dynamicConfig.listener.AddConfigPath("../")
	dynamicConfig.listener.AddConfigPath("./profiles")
	dynamicConfig.listener.AddConfigPath("/vault/secrets/")
	dynamicConfig.listener.SetConfigType("yaml")
	dynamicConfig.listener.SetConfigName(dynamicFilePath)
	err = dynamicConfig.listener.ReadInConfig()
	if err != nil {
		fmt.Println("error while reading dynamic config file")
		panic(err)
	}

	dynamicConfig.SyncDynamicConfigsToReader()
	dynamicConfig.listener.WatchConfig()
	dynamicConfig.listener.OnConfigChange(func(in fsnotify.Event) {
		// sync configs to sync Map for concurrent Reads
		fmt.Println("Dynamic Configs Changed")
		dynamicConfig.SyncDynamicConfigsToReader()
		if commandArgs[1] == "tagsEventConsumerWorker" {
			TagsConsumerEnableChannel <- true
		}
	})

	enableSentry := getBoolOrPanic("enable_sentry")

	var sentryDsn string
	if enableSentry {
		sentryDsn = getStringOrPanic("sentry_dsn")
	}
	authSecretV2, err := base64.StdEncoding.DecodeString(getStringOrPanic("auth_secret_v2"))

	if err != nil {
		panic("Error in auth secret")
	}
	config = &Configuration{
		port:         getIntOrPanic("port"),
		logLevel:     getStringOrPanic("log_level"),
		logLocation:  getStringOrPanic("log_location"),
		sentryDsn:    sentryDsn,
		enableSentry: enableSentry,
		authSecret:   getStringOrPanic("auth_sercet"),
		authSecretV2: authSecretV2,
		RedisConfig: RedisConfig{
			Addr:         getStringOrPanic("redis_addr"),
			Password:     getStringOrPanic("redis_password"),
			Db:           int(getIntOrPanic("redis_db")),
			ReadTimeout:  time.Duration(getIntOrPanic("redis_read_timeout")),
			WriteTimeout: time.Duration(getIntOrPanic("redis_write_timeout")),
			DialTimeout:  time.Duration(getIntOrPanic("redis_dial_timeout")),
		},
		MongoConfig: MongoConfig{
			Uri:            getStringOrPanic("mongo_uri"),
			Username:       getStringOrPanic("mongo_username"),
			Password:       getStringOrPanic("mongo_password"),
			Database:       getStringOrPanic("mongo_db"),
			ConnectTimeout: time.Duration(getIntOrPanic("mongo_connect_timeout")),
			AuthSource:     getStringOrPanic("mongo_auth_source"),
		},
		Grpc: GrpcConfig{
			Port: getIntOrPanic("grpc_port"),
			Host: getStringOrPanic("grpc_host"),
		},
		kafkaConfig: KafkaConfig{
			Brokers:             getStringSliceOrPanic("kafka_brokers"),
			TopicsMap:           KafkaTopicsMap{},
			NewKafkaFlowEnabled: NewKafkaFlowEnabled{},
		},
	}
}

func Load(commandArgs []string) {
	LoadWithoutStoreSdk(commandArgs)
}

func Port() int64 {
	return config.port
}

func LogLevel() string {
	return config.logLevel
}

func GetLoggingEnabledComponents() []string {
	return dynamicConfig.GetStringArrayWithFallback("logging_enabled_components", []string{})
}

func LogLocation() string {
	return config.logLocation
}

func SentryDsn() string {
	return config.sentryDsn
}

func EnableSentry() bool {
	return config.enableSentry
}

func AuthSecret() string {
	return config.authSecret
}

func AuthSecretV2() []byte {
	return config.authSecretV2
}

func LoadDefault() {
	pc, b, line, ok := runtime.Caller(0)
	if !ok {
		fmt.Println("runtime caller failed", pc, b, line, ok)
	}
	staticConfig := []string{"", "", "profiles/test.yml", "profiles/test.yml"}
	Load(staticConfig)
}
func LoadConfigWithOverride(t *testing.T, key string, value string) func() {
	t.Helper()

	origVal := os.Getenv(key)
	origCfg := *config
	_ = os.Setenv(key, value)
	LoadDefault()

	return func() {
		_ = os.Setenv(key, origVal)
		LoadDefault()
		*config = origCfg
	}
}

func LoadTestDynamicConfigWithOverride(t *testing.T, key string, value string) func() {
	t.Helper()
	origVal := viper.GetString(key)
	viper.Set(key, value)
	old, ok := dynamicConfig.reader.Load(key)
	dynamicConfig.reader.Store(key, value) // override reader directly

	return func() {
		viper.Set(key, origVal)
		if ok {
			dynamicConfig.reader.Store(key, old)
		} else {
			dynamicConfig.reader.Delete(key)
		}
	}
}

func GetServerEnvironment() string {
	if config.stage == constants.Production {
		return config.stage
	}
	return constants.Development
}

func KafkaConf() KafkaConfig {
	return config.kafkaConfig
}
func GetKafkaBrokers() []string {
	return config.kafkaConfig.Brokers
}

func GetBifrostAuth() Bifrost {
	return config.Bifrost
}

func GetRedisConf() RedisConfig {
	return config.RedisConfig
}

func GetMongoConf() MongoConfig {
	return config.MongoConfig
}
