package httplib

import (
	"base_app/logger"
	"context"
	"github.com/go-resty/resty/v2"
	"time"
)

type HttpClient struct {
	C *resty.Client
}

func NewHttpClient(retryCount int, retryWaitTime time.Duration, requestTimeout time.Duration) *HttpClient {
	c := resty.New()
	c.SetRetryCount(retryCount)
	c.SetRetryWaitTime(retryWaitTime)
	c.SetTimeout(requestTimeout)

	c.OnBeforeRequest(func(c *resty.Client, req *resty.Request) error {
		logger.Log.Debugf("Request : %v", req)
		return nil
	})

	c.OnAfterResponse(func(client *resty.Client, response *resty.Response) error {
		if !response.IsSuccess() {
			logger.Log.Infof("Downstream non OK response : %v, for request %v", response, response.Request)
		}
		logger.Log.Debugf("Downstream response : %v", response)
		return nil
	})

	return &HttpClient{C: c}
}

func (h *HttpClient) CreateRequest(ctx context.Context, body interface{}, res interface{}) *resty.Request {
	req := h.C.R().SetBody(body).SetResult(res)
	req.SetContext(ctx)
	return req
}
