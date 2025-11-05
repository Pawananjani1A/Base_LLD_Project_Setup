package config

import (
	"github.com/newrelic/go-agent/v3/newrelic"
	"golang.org/x/exp/slices"
	"net/http"
)

const (
	maxSampleRecord = 10000
)

func LoadNewRelicApplication(appName string) (*newrelic.Application, error) {
	nrConfigKey := "new_relic_app_name"
	if appName == "uss-admin-backend" {
		nrConfigKey = "new_relic_uss_admin_backend_app_name"
	}
	return newrelic.NewApplication(
		newrelic.ConfigAppName(getStringOrPanic(nrConfigKey)),
		newrelic.ConfigLicense(getStringOrPanic("new_relic_licence_key")),
		newrelic.ConfigEnabled(getBoolOrPanic("enable_new_relic")),
		func(config *newrelic.Config) {
			config.ErrorCollector.IgnoreStatusCodes = []int{
				http.StatusBadRequest,
				http.StatusNotFound,
				http.StatusTooManyRequests,
			}
		},
	)
}

func LoadNewRelicWorkerApplication(newRelicWorkerAppName string) (*newrelic.Application, error) {
	if slices.Contains(dynamicConfig.GetStringArrayWithFallback("applicable_workers_for_sampling", []string{}), newRelicWorkerAppName) {
		return newrelic.NewApplication(
			newrelic.ConfigAppName(newRelicWorkerAppName),
			newrelic.ConfigLicense(getStringOrPanic("new_relic_licence_key")),
			newrelic.ConfigEnabled(getBoolOrPanic("enable_new_relic")),
			newrelic.ConfigAppLogForwardingEnabled(false),
			newrelic.ConfigAppLogMetricsEnabled(false),
			func(config *newrelic.Config) {
				config.ErrorCollector.IgnoreStatusCodes = []int{
					http.StatusBadRequest,
					http.StatusNotFound,
				}
				config.TransactionEvents.MaxSamplesStored = int(maxSampleRecord * getFloatOrPanic("sampling_rate"))
				config.CustomInsightsEvents.MaxSamplesStored = int(maxSampleRecord * getFloatOrPanic("sampling_rate"))
			},
		)
	}
	return newrelic.NewApplication(
		newrelic.ConfigAppName(newRelicWorkerAppName),
		newrelic.ConfigLicense(getStringOrPanic("new_relic_licence_key")),
		newrelic.ConfigEnabled(getBoolOrPanic("enable_new_relic")),
		func(config *newrelic.Config) {
			config.ErrorCollector.IgnoreStatusCodes = []int{
				http.StatusBadRequest,
				http.StatusNotFound,
			}
		},
	)
}
