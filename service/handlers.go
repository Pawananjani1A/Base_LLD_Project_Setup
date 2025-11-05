package service

import (
	"base_app/config"
	"base_app/constants/api_constants"
	"base_app/logger"
	"base_app/schema"
	"base_app/util"
	"context"
	"encoding/json"
	"net/http"
	"strconv"
)

type PingResponse struct {
	Status string `json:"status"`
}

func PingHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, _ *http.Request) {
		payload, _ := json.Marshal(PingResponse{
			Status: "healthy",
		})
		w.Header().Set("Content-Type", "application/json")
		w.Write(payload)
	}
}

type ApiHandlers interface {
	GetResults() func(w http.ResponseWriter, r *http.Request)
}

type ApiHandlersImpl struct {
}

func ApiHandler() ApiHandlers {
	return &ApiHandlersImpl{}
}

func (handler *ApiHandlersImpl) GetResults() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		decoder := json.NewDecoder(r.Body)
		decoder.DisallowUnknownFields()

		ctx = logger.AddEnabledLoggingComponentsToContext(ctx, config.GetLoggingEnabledComponents())

		var searchRequest schema.RequestPayload
		err := decoder.Decode(&searchRequest)
		if err != nil {
			logger.Errorf(ctx, "error while decoding request body: %v, err: %s", r.Body, err.Error())
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Please provide valid payload to continue"))
			return
		}

		bearerToken := r.Header.Get("Authorization")
		platform := r.Header.Get("platform")
		source := r.Header.Get("source")
		appVersion := r.Header.Get("appVersion")
		bundleVersion := r.Header.Get("bundleVersion")
		deviceId := r.Header.Get("deviceId")

		// Added requestId and UserId in logs to trace the API flow
		userId, requestId := util.GetUserIdAndRequestIdFromContext(ctx)

		ctx = context.WithValue(ctx, api_constants.UserId, userId)
		ctx = context.WithValue(ctx, api_constants.RequestId, requestId)
		ctx = context.WithValue(ctx, api_constants.DeviceId, deviceId)
		ctx = context.WithValue(ctx, api_constants.AppVersion, appVersion)

		log := logger.GetLoggerFromContext(ctx)

		searchRequestJson, _ := json.Marshal(searchRequest)
		log.Infof("[GetResults] searchRequest = %v, userId=%v, platform=%v, source=%v, appversion=%v, bundleVersion=%v, deviceId=%v, ", string(searchRequestJson), userId, platform, source, appVersion, bundleVersion, deviceId)

		searchHeaderRequest := &schema.RequestHeaders{
			BearerToken: bearerToken,
			DeviceId:    deviceId,
			Platform:    platform,
			Source:      source,
		}

		log.Infof("[v1/search] searchHeaderRequest = %+v", searchHeaderRequest)

		response := schema.ResponsePayload{
			Success: true,
		}
		w.Header().Set("Content-Type", "application/json")
		payload, _ := json.Marshal(response)
		w.Header().Set("Content-Length", strconv.Itoa(len(payload)))
		w.Write(payload)
	}
}
