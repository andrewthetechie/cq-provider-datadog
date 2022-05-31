package services

import (
	"context"
	"fmt"

	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"github.com/andrewthetechie/cq-provider-datadog/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource hosts --config ./resources/services/v1/hosts.hcl --output .
func Hosts() *schema.Table {
	return &schema.Table{
		Name:        "datadog_hosts",
		Description: "Host Object representing a host.",
		Multiplex:   client.AccountMultiplex,
		Resolver:    fetchHosts,
		Options:     schema.TableCreationOptions{PrimaryKeys: []string{"id"}},
		Columns: []schema.Column{
			{
				Name:        "account_name",
				Description: "The name of this datadog account from your config.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAccountName,
			},
			{
				Name:        "aliases",
				Description: "Host aliases collected by Datadog.",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "apps",
				Description: "The Datadog integrations reporting metrics for the host.",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "aws_name",
				Description: "AWS name of your host.",
				Type:        schema.TypeString,
			},
			{
				Name:        "host_name",
				Description: "The host name.",
				Type:        schema.TypeString,
			},
			{
				Name:        "id",
				Description: "The host ID.",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "is_muted",
				Description: "If a host is muted or unmuted.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "last_reported_time",
				Description: "Last time the host reported a metric data point.",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "meta_agent_version",
				Description: "The Datadog Agent version.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Meta.AgentVersion"),
			},
			{
				Name:        "meta_cpu_cores",
				Description: "The number of cores.",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("Meta.CpuCores"),
			},
			{
				Name:        "meta_fbsd_v",
				Description: "An array of Mac versions.",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("Meta.FbsdV"),
			},
			{
				Name:        "meta_gohai",
				Description: "JSON string containing system information.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Meta.Gohai"),
			},
			{
				Name:        "meta_install_method_installer_version",
				Description: "The installer version.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Meta.InstallMethod.InstallerVersion"),
			},
			{
				Name:        "meta_install_method_tool",
				Description: "Tool used to install the agent.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Meta.InstallMethod.Tool"),
			},
			{
				Name:        "meta_install_method_tool_version",
				Description: "The tool version.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Meta.InstallMethod.ToolVersion"),
			},
			{
				Name:     "meta_install_method_additional_properties",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Meta.InstallMethod.AdditionalProperties"),
			},
			{
				Name:        "meta_mac_v",
				Description: "An array of Mac versions.",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("Meta.MacV"),
			},
			{
				Name:        "meta_machine",
				Description: "The machine architecture.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Meta.Machine"),
			},
			{
				Name:        "meta_nix_v",
				Description: "Array of Unix versions.",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("Meta.NixV"),
			},
			{
				Name:        "meta_platform",
				Description: "The OS platform.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Meta.Platform"),
			},
			{
				Name:        "meta_processor",
				Description: "The processor.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Meta.Processor"),
			},
			{
				Name:        "meta_python_v",
				Description: "The Python version.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Meta.PythonV"),
			},
			{
				Name:        "meta_win_v",
				Description: "An array of Windows versions.",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("Meta.WinV"),
			},
			{
				Name:     "meta_additional_properties",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Meta.AdditionalProperties"),
			},
			{
				Name:        "metrics_cpu",
				Description: "The percent of CPU used (everything but idle).",
				Type:        schema.TypeFloat,
				Resolver:    schema.PathResolver("Metrics.Cpu"),
			},
			{
				Name:        "metrics_iowait",
				Description: "The percent of CPU spent waiting on the IO (not reported for all platforms).",
				Type:        schema.TypeFloat,
				Resolver:    schema.PathResolver("Metrics.Iowait"),
			},
			{
				Name:        "metrics_load",
				Description: "The system load over the last 15 minutes.",
				Type:        schema.TypeFloat,
				Resolver:    schema.PathResolver("Metrics.Load"),
			},
			{
				Name:     "metrics_additional_properties",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Metrics.AdditionalProperties"),
			},
			{
				Name:        "mute_timeout",
				Description: "Timeout of the mute applied to your host.",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "name",
				Description: "The host name.",
				Type:        schema.TypeString,
			},
			{
				Name:        "sources",
				Description: "Source or cloud provider associated with your host.",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "tags_by_source",
				Description: "List of tags for each source (AWS, Datadog Agent, Chef..).",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "up",
				Description: "Displays UP when the expected metrics are received and displays `???` if no metrics are received.",
				Type:        schema.TypeBool,
			},
			{
				Name: "additional_properties",
				Type: schema.TypeJSON,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchHosts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	logger := c.Logger()
	logger.Debug("in fetchHosts")
	// TODO: multiplexing
	apiClient := datadog.NewAPIClient(&c.MultiPlexedAccount.V1Config)
	var step int64 = 1000
	params := datadog.NewListHostsOptionalParameters().WithIncludeMutedHostsData(true).WithIncludeHostsMetadata(true).WithCount(step)
	resp, r, err := apiClient.HostsApi.ListHosts(c.MultiPlexedAccount.V1Context, *params)
	logger.Debug(r.Status)
	if err != nil {
		return diag.FromError(err, diag.ACCESS)
	}
	var retrievedCount int64 = *resp.TotalReturned
	var totalCount int64 = *resp.TotalMatching
	res <- resp.HostList
	logger.Debug(fmt.Sprintf("TotalCount: %d, RetrievedCount: %d", totalCount, retrievedCount))
	for retrievedCount < totalCount {
		logger.Debug("Looping host requests")
		params := datadog.NewListHostsOptionalParameters().WithIncludeMutedHostsData(true).WithIncludeHostsMetadata(true).WithCount(step).WithStart(retrievedCount)
		resp, r, err := apiClient.HostsApi.ListHosts(c.MultiPlexedAccount.V1Context, *params)
		logger.Debug(r.Status)
		if err != nil {
			return diag.FromError(err, diag.ACCESS)
		}
		res <- resp.HostList
		retrievedCount += *resp.TotalReturned
	}
	logger.Debug(fmt.Sprintf("Total Retrieved: %d", retrievedCount))
	return nil
}
