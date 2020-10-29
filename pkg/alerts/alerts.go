package alerts

import (
	"github.com/joeyparsons/newrelic-alerts-client-go/internal/http"
	"github.com/joeyparsons/newrelic-alerts-client-go/internal/logging"
	"github.com/joeyparsons/newrelic-alerts-client-go/pkg/config"
	"github.com/joeyparsons/newrelic-alerts-client-go/pkg/infrastructure"
)

// Alerts is used to communicate with New Relic Alerts.
type Alerts struct {
	client      http.Client
	config      config.Config
	infraClient http.Client
	logger      logging.Logger
	pager       http.Pager
}

// New is used to create a new Alerts client instance.
func New(config config.Config) Alerts {
	infraConfig := config

	infraClient := http.NewClient(infraConfig)
	infraClient.SetAuthStrategy(&http.PersonalAPIKeyCapableV2Authorizer{})
	infraClient.SetErrorValue(&infrastructure.ErrorResponse{})

	client := http.NewClient(config)
	client.SetAuthStrategy(&http.PersonalAPIKeyCapableV2Authorizer{})

	pkg := Alerts{
		client:      client,
		config:      config,
		infraClient: infraClient,
		logger:      config.GetLogger(),
		pager:       &http.LinkHeaderPager{},
	}

	return pkg
}