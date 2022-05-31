package services

import (
	"context"

	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"github.com/andrewthetechie/cq-provider-datadog/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource monitors --config ./resources/services/v1/monitors.hcl --output .
func Monitors() *schema.Table {
	return &schema.Table{
		Name:        "datadog_monitors",
		Description: "Monitor Object describing a monitor.",
		Multiplex:   client.AccountMultiplex,
		Resolver:    fetchMonitors,
		Options:     schema.TableCreationOptions{PrimaryKeys: []string{"id"}},
		Columns: []schema.Column{
			{
				Name:        "account_name",
				Description: "The name of this datadog account from your config.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAccountName,
			},
			{
				Name:        "created",
				Description: "Timestamp of the monitor creation.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "creator_email",
				Description: "Email of the creator.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Creator.Email"),
			},
			{
				Name:        "creator_handle",
				Description: "Handle of the creator.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Creator.Handle"),
			},
			{
				Name:     "creator_additional_properties",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Creator.AdditionalProperties"),
			},
			{
				Name:        "id",
				Description: "ID of this monitor.",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "message",
				Description: "A message to include with notifications for this monitor.",
				Type:        schema.TypeString,
			},
			{
				Name:        "modified",
				Description: "Last timestamp when the monitor was edited.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "multi",
				Description: "Whether or not the monitor is broken down on different groups.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "name",
				Description: "The monitor name.",
				Type:        schema.TypeString,
			},
			{
				Name:        "options_aggregation_group_by",
				Description: "Group to break down the monitor on.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Options.Aggregation.GroupBy"),
			},
			{
				Name:        "options_aggregation_metric",
				Description: "Metric name used in the monitor.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Options.Aggregation.Metric"),
			},
			{
				Name:        "options_aggregation_type",
				Description: "Metric type used in the monitor.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Options.Aggregation.Type"),
			},
			{
				Name:     "options_aggregation_additional_properties",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Options.Aggregation.AdditionalProperties"),
			},
			{
				Name:        "options_device_ids",
				Description: "IDs of the device the Synthetics monitor is running on. Deprecated",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("Options.DeviceIds"),
			},
			{
				Name:        "options_enable_logs_sample",
				Description: "Whether or not to send a log sample when the log monitor triggers.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Options.EnableLogsSample"),
			},
			{
				Name:        "options_escalation_message",
				Description: "We recommend using the [is_renotify](https://docs.datadoghq.com/monitors/notify/?tab=is_alert#renotify), block in the original message instead. A message to include with a re-notification",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Options.EscalationMessage"),
			},
			{
				Name:        "options_groupby_simple_monitor",
				Description: "Whether the log alert monitor triggers a single alert or multiple alerts when any group breaches a threshold.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Options.GroupbySimpleMonitor"),
			},
			{
				Name:        "options_include_tags",
				Description: "A Boolean indicating whether notifications from this monitor automatically inserts its triggering tags into the title.  **Examples** - If `True`, `[Triggered on {host:h1}] Monitor Title` - If `False`, `[Triggered] Monitor Title`",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Options.IncludeTags"),
			},
			{
				Name:        "options_locked",
				Description: "Whether or not the monitor is locked (only editable by creator and admins)",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Options.Locked"),
			},
			{
				Name:        "options_notify_audit",
				Description: "A Boolean indicating whether tagged users is notified on changes to this monitor.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Options.NotifyAudit"),
			},
			{
				Name:        "options_notify_no_data",
				Description: "A Boolean indicating whether this monitor notifies when data stops reporting.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Options.NotifyNoData"),
			},
			{
				Name:        "options_renotify_statuses",
				Description: "The types of monitor statuses for which re-notification messages are sent.",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("Options.RenotifyStatuses"),
			},
			{
				Name:        "options_require_full_window",
				Description: "A Boolean indicating whether this monitor needs a full window of data before it’s evaluated. We highly recommend you set this to `false` for sparse metrics, otherwise some evaluations are skipped",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Options.RequireFullWindow"),
			},
			{
				Name:        "options_silenced",
				Description: "Information about the downtime applied to the monitor. Deprecated",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("Options.Silenced"),
			},
			{
				Name:     "options_threshold_windows_additional_properties",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Options.ThresholdWindows.AdditionalProperties"),
			},
			{
				Name:        "options_thresholds_critical",
				Description: "The monitor `CRITICAL` threshold.",
				Type:        schema.TypeFloat,
				Resolver:    schema.PathResolver("Options.Thresholds.Critical"),
			},
			{
				Name:     "options_thresholds_additional_properties",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Options.Thresholds.AdditionalProperties"),
			},
			{
				Name:     "options_additional_properties",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Options.AdditionalProperties"),
			},
			{
				Name:        "overall_state",
				Description: "The different states your monitor can be in.",
				Type:        schema.TypeString,
			},
			{
				Name:        "query",
				Description: "The monitor query.",
				Type:        schema.TypeString,
			},
			{
				Name:        "restricted_roles",
				Description: "A list of unique role identifiers to define which roles are allowed to edit the monitor",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "state_groups",
				Description: "Dictionary where the keys are groups (comma separated lists of tags) and the values are the list of groups your monitor is broken down on.",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("State.Groups"),
			},
			{
				Name:     "state_additional_properties",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("State.AdditionalProperties"),
			},
			{
				Name:        "tags",
				Description: "Tags associated to your monitor.",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "type",
				Description: "The type of the monitor",
				Type:        schema.TypeString,
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

func fetchMonitors(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	logger := c.Logger()
	logger.Debug("in fetchDatadogMonitors")
	// TODO: multiplexing
	apiClient := datadog.NewAPIClient(&c.MultiPlexedAccount.V1Config)
	resp, r, err := apiClient.MonitorsApi.ListMonitors(c.MultiPlexedAccount.V1Context)
	logger.Debug(r.Status)
	if err != nil {
		return diag.FromError(err, diag.ACCESS)
	}
	res <- resp
	return nil
}
