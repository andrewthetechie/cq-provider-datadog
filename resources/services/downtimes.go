package services

import (
	"context"

	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"github.com/andrewthetechie/cq-provider-datadog/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource downtimes --config ./resources/services/v1/downtimes.hcl --output .
func Downtimes() *schema.Table {
	return &schema.Table{
		Name:        "datadog_downtimes",
		Description: "Downtime Downtiming gives you greater control over monitor notifications by allowing you to globally exclude scopes from alerting. Downtime settings, which can be scheduled with start and end times, prevent all alerting related to specified Datadog tags.",
		Multiplex:   client.AccountMultiplex,
		Resolver:    fetchDowntimes,
		Options:     schema.TableCreationOptions{PrimaryKeys: []string{"id"}},
		Columns: []schema.Column{
			{
				Name:        "account_name",
				Description: "The name of this datadog account from your config.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAccountName,
			},
			//Todo: Handle all the NullableInt64 columns
			{
				Name:        "active",
				Description: "If a scheduled downtime currently exists.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "active_child",
				Description: "",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "creator_id",
				Description: "User ID of the downtime creator.",
				Type:        schema.TypeInt,
			},
			{
				Name:        "disabled",
				Description: "If a downtime has been disabled.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "downtime_type",
				Description: "`0` for a downtime applied on `*` or all, `1` when the downtime is only scoped to hosts, or `2` when the downtime is scoped to anything but hosts.",
				Type:        schema.TypeInt,
			},
			{
				Name:        "id",
				Description: "The downtime ID.",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "message",
				Description: "A message to include with notifications for this downtime. Email notifications can be sent to specific users by using the same `@username` notation as events.",
				Type:        schema.TypeString,
			},
			{
				Name:        "monitor_tags",
				Description: "A comma-separated list of monitor tags",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "mute_first_recovery_notification",
				Description: "If the first recovery notification during a downtime should be muted.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "recurrence",
				Description: "",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "scope",
				Description: "The scope(s) to which the downtime applies",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "start",
				Description: "POSIX timestamp to start the downtime. If not provided, the downtime starts the moment it is created.",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "timezone",
				Description: "The timezone in which to display the downtime's start and end times in Datadog applications.",
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

func fetchDowntimes(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	logger := c.Logger()
	logger.Debug("in fetchDowntimes")
	// TODO: multiplexing
	apiClient := datadog.NewAPIClient(&c.MultiPlexedAccount.V1Config)
	resp, r, err := apiClient.DowntimesApi.ListDowntimes(c.MultiPlexedAccount.V1Context, *datadog.NewListDowntimesOptionalParameters().WithCurrentOnly(true))
	logger.Debug(r.Status)
	if err != nil {
		return diag.FromError(err, diag.ACCESS)
	}
	res <- resp
	return nil
}
