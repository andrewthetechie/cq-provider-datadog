package client

// Provider Configuration

type Account struct {
	Name   string `hcl:",label"`
	APIKey string `hcl:"api_key"`
	AppKey string `hcl:"app_key"`
	APIUrl string `hcl:"api_url,optional"`
}

type Config struct {
	Accounts    []Account `hcl:"accounts,block"`
	DatdogDebug bool      `hcl:"datadog_debug,optional"`
	MaxRetries  int       `hcl:"max_retries,optional" default:"10"`
	MaxBackoff  int       `hcl:"max_backoff,optional" default:"30"`
}

func (c Config) Example() string {
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
}
