package provider

import (
	"github.com/cloudquery/cq-provider-sdk/cqproto"
	"github.com/cloudquery/cq-provider-sdk/provider"
	sdkschema "github.com/cloudquery/cq-provider-sdk/provider/schema"

	"github.com/andrewthetechie/cq-provider-datadog/client"
	"github.com/andrewthetechie/cq-provider-datadog/resources/services"
)

var (
	Version = "Development"
)

func Provider() *provider.Provider {

	return &provider.Provider{
		Version:   Version,
		Name:      "datadog",
		Configure: client.Configure,
		ResourceMap: map[string]*sdkschema.Table{
			"monitors":        services.Monitors(),
			"notebooks":       services.Notebooks(),
			"hosts":           services.Hosts(),
			"dashboard_lists": services.DashboardLists(),
			"dashboards":      services.Dashboards(),
			"downtimes":       services.Downtimes(),
			"incidents":       services.Incidents(),
			"permissions":     services.Permissions(),
			"roles":           services.Roles(),
			"synthetics":      services.Synthetics(),
			"users":           services.Users(),
		},
		Config: func(f cqproto.ConfigFormat) provider.Config {
			return client.NewConfig(f)
		},
	}

}
