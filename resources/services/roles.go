package services

import (
	"context"

	datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
	"github.com/andrewthetechie/cq-provider-datadog/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource roles --config ./resources/services/v2/roles.hcl --output .
func Roles() *schema.Table {
	return &schema.Table{
		Name:        "datadog_roles",
		Description: "Role Role object returned by the API.",
		Resolver:    fetchRoles,
		Options:     schema.TableCreationOptions{PrimaryKeys: []string{"id"}},
		Columns: []schema.Column{
			{
				Name:        "attributes_created_at",
				Description: "Creation time of the role.",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("Attributes.CreatedAt"),
			},
			{
				Name:        "attributes_modified_at",
				Description: "Time of last role modification.",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("Attributes.ModifiedAt"),
			},
			{
				Name:        "attributes_name",
				Description: "The name of the role",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Attributes.Name"),
			},
			{
				Name:        "attributes_user_count",
				Description: "Number of users with that role.",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("Attributes.UserCount"),
			},
			{
				Name:     "attributes_additional_properties",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Attributes.AdditionalProperties"),
			},
			{
				Name:        "id",
				Description: "The unique identifier of the role.",
				Type:        schema.TypeString,
			},
			{
				Name:     "relationships_permissions_additional_properties",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Relationships.Permissions.AdditionalProperties"),
			},
			{
				Name:     "relationships_additional_properties",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Relationships.AdditionalProperties"),
			},
			{
				Name:        "type",
				Description: "Roles type.",
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

func fetchRoles(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	logger := c.Logger()
	logger.Debug("in fetchHosts")
	// TODO: multiplexing
	apiClient := datadog.NewAPIClient(&c.Accounts[0].V2Config)

	resp, r, err := apiClient.RolesApi.ListRoles(c.Accounts[0].V2Context)
	logger.Debug(r.Status)
	if err != nil {
		return diag.FromError(err, diag.ACCESS)
	}
	res <- resp.Data
	return nil
}
