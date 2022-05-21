package services

import (
	"context"

	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"github.com/andrewthetechie/cq-provider-datadog/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource dashboard_lists --config resources/services/dashboard_lists.hcl --output .
func DashboardLists() *schema.Table {
	return &schema.Table{
		Name:        "datadog_dashboard_lists",
		Description: "DashboardList Your Datadog Dashboards.",
		Resolver:    fetchDashboardLists,
		Options:     schema.TableCreationOptions{PrimaryKeys: []string{"id"}},
		Columns: []schema.Column{
			{
				Name:        "author_email",
				Description: "Email of the creator.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Author.Email"),
			},
			{
				Name:        "author_handle",
				Description: "Handle of the creator.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Author.Handle"),
			},
			{
				Name:     "author_additional_properties",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Author.AdditionalProperties"),
			},
			{
				Name:        "created",
				Description: "Date of creation of the dashboard list.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "dashboard_count",
				Description: "The number of dashboards in the list.",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "id",
				Description: "The ID of the dashboard list.",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "is_favorite",
				Description: "Whether or not the list is in the favorites.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "modified",
				Description: "Date of last edition of the dashboard list.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "name",
				Description: "The name of the dashboard list.",
				Type:        schema.TypeString,
			},
			{
				Name:        "type",
				Description: "The type of dashboard list.",
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

func fetchDashboardLists(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	logger := c.Logger()
	logger.Debug("in fetchHosts")
	// TODO: multiplexing
	apiClient := datadog.NewAPIClient(&c.Accounts[0].V1Config)
	resp, _, err := apiClient.DashboardListsApi.ListDashboardLists(c.Accounts[0].V1Context)
	if err != nil {
		return diag.FromError(err, diag.ACCESS)
	}
	res <- resp.DashboardLists

	return nil
}