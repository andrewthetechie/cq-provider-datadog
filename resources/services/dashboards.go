package services

import (
	"context"

	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"github.com/andrewthetechie/cq-provider-datadog/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource dashboards --config ./resources/services/v1/dashboards.hcl --output .
func Dashboards() *schema.Table {
	return &schema.Table{
		Name:        "datadog_dashboards",
		Description: "Dashboard A dashboard is Datadog’s tool for visually tracking, analyzing, and displaying key performance metrics, which enable you to monitor the health of your infrastructure.",
		Resolver:    fetchDashboards,
		Options:     schema.TableCreationOptions{PrimaryKeys: []string{"id"}},
		Columns: []schema.Column{
			{
				Name:        "author_handle",
				Description: "Identifier of the dashboard author.",
				Type:        schema.TypeString,
			},
			{
				Name:        "created_at",
				Description: "Creation date of the dashboard.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "id",
				Description: "ID of the dashboard.",
				Type:        schema.TypeString,
			},
			{
				Name:        "is_read_only",
				Description: "Whether this dashboard is read-only",
				Type:        schema.TypeBool,
			},
			{
				Name:        "layout_type",
				Description: "Layout type of the dashboard.",
				Type:        schema.TypeString,
			},
			{
				Name:        "modified_at",
				Description: "Modification date of the dashboard.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "notify_list",
				Description: "List of handles of users to notify when changes are made to this dashboard.",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "reflow_type",
				Description: "Reflow type for a **new dashboard layout** dashboard",
				Type:        schema.TypeString,
			},
			{
				Name:        "restricted_roles",
				Description: "A list of role identifiers",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "title",
				Description: "Title of the dashboard.",
				Type:        schema.TypeString,
			},
			{
				Name:        "url",
				Description: "The URL of the dashboard.",
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

func fetchDashboards(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	logger := c.Logger()
	logger.Debug("in fetchHosts")
	// TODO: multiplexing
	apiClient := datadog.NewAPIClient(&c.Accounts[0].V1Config)
	resp, r, err := apiClient.DashboardsApi.ListDashboards(c.Accounts[0].V1Context)
	logger.Debug(r.Status)
	if err != nil {
		return diag.FromError(err, diag.ACCESS)
	}
	res <- resp.Dashboards
	return nil
}
