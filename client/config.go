package client

import "github.com/cloudquery/cq-provider-sdk/cqproto"

// Provider Configuration

type Account struct {
	Name   string `yaml:"name" hcl:",label"`
	APIKey string `yaml:"api_key" hcl:"api_key"`
	AppKey string `yaml:"app_key" hcl:"app_key"`
	APIUrl string `yaml:"api_url,omitempty" hcl:"api_url,optional"`
}

type Config struct {
	Accounts    []Account `hcl:"accounts,block"`
	DatdogDebug bool      `hcl:"datadog_debug,optional"`
	MaxRetries  int       `hcl:"max_retries,optional" default:"10"`
	MaxBackoff  int       `hcl:"max_backoff,optional" default:"30"`

	requestedFormat cqproto.ConfigFormat
}

func (c Config) Example() string {
	switch c.requestedFormat {
	case cqproto.ConfigHCL:
		return `configuration {
			accounts "main" {
				api_key = "datadog api key"
				app_key = "datadog app key"
				// api_url = "datadoghq.eu" use to set a custom datadog api url
			}
		// Optional or required parameters
		// datadog_debug will turn on debug logging from the datadog client
		// datadog_debug = false
		
	}
	`
	default:
		return `
accounts:
- name: "main"
  api_key: "datadog api key"
  app_key: "datadog app key"
  # Optional, you can add an API url to use a custom datadog api url
  # api_url: "datadoghq.eu"
	`
	}

}

func (c Config) Format() cqproto.ConfigFormat {
	return c.requestedFormat
}
