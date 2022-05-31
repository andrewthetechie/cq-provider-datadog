package client

import (
	"context"
	"errors"

	datadogv1 "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	datadogv2 "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/hashicorp/go-hclog"
)

type Client struct {
	logger hclog.Logger

	Accounts       []DDAccount
	AccountConfigs []Account
	// this is set by the table client multiplexer
	MultiPlexedAccount DDAccount
}

func (c *Client) Logger() hclog.Logger {
	return c.logger
}

type DDAccount struct {
	V1Context context.Context
	V2Context context.Context
	V1Config  datadogv1.Configuration
	V2Config  datadogv2.Configuration
	Name      string
}

func (c Client) withAccount(account Account) *Client {
	return &Client{
		logger:   c.logger.With("id", account.Name),
		Accounts: c.Accounts,
		MultiPlexedAccount: DDAccount{
			Name: account.Name,
			V1Context: context.WithValue(
				context.Background(),
				datadogv1.ContextAPIKeys,
				map[string]datadogv1.APIKey{
					"apiKeyAuth": {
						Key: account.APIKey,
					},
					"appKeyAuth": {
						Key: account.AppKey,
					},
				},
			),
			V2Context: context.WithValue(
				context.Background(),
				datadogv2.ContextAPIKeys,
				map[string]datadogv2.APIKey{
					"apiKeyAuth": {
						Key: account.APIKey,
					},
					"appKeyAuth": {
						Key: account.AppKey,
					},
				},
			),
			V1Config: *datadogv1.NewConfiguration(),
			V2Config: *datadogv2.NewConfiguration(),
		}}
}

func Configure(logger hclog.Logger, config interface{}) (schema.ClientMeta, diag.Diagnostics) {
	logger.Info("in datadog configure")
	diag.WithDetails("In Datadog configure")
	ddConfig := config.(*Config)
	if len(ddConfig.Accounts) == 0 {
		logger.Error("error no datadog accounts configured")
		return nil, diag.FromError(errors.New("error no datadog accounts configured"), diag.USER)
	}
	ddAccounts := make([]DDAccount, 0)
	for _, account := range ddConfig.Accounts {
		logger.Debug("user definied account", "account", account.Name)
		ddAccounts = append(ddAccounts, DDAccount{
			Name: account.Name,
			V1Context: context.WithValue(
				context.Background(),
				datadogv1.ContextAPIKeys,
				map[string]datadogv1.APIKey{
					"apiKeyAuth": {
						Key: account.APIKey,
					},
					"appKeyAuth": {
						Key: account.AppKey,
					},
				},
			),
			V2Context: context.WithValue(
				context.Background(),
				datadogv2.ContextAPIKeys,
				map[string]datadogv2.APIKey{
					"apiKeyAuth": {
						Key: account.APIKey,
					},
					"appKeyAuth": {
						Key: account.AppKey,
					},
				},
			),
			V1Config: *datadogv1.NewConfiguration(),
			V2Config: *datadogv2.NewConfiguration(),
		})
	}

	// TODO: Figure out how to initialize all the datadog clients and pass them in here
	// For now, going to do that in the resources. Still not sure how multiplexers work
	client := Client{
		logger:         logger,
		Accounts:       ddAccounts,
		AccountConfigs: ddConfig.Accounts,
	}

	// Return the initialized client and it will be passed to your resources
	return &client, nil
}
