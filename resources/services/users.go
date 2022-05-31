package services

import (
	"context"

	datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
	"github.com/andrewthetechie/cq-provider-datadog/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource users --config ./resources/services/v2/users.hcl --output .
func Users() *schema.Table {
	return &schema.Table{
		Name:        "datadog_users",
		Description: "User User object returned by the API.",
		Multiplex:   client.AccountMultiplex,
		Resolver:    fetchUsers,
		Options:     schema.TableCreationOptions{PrimaryKeys: []string{"id"}},
		Columns: []schema.Column{
			{
				Name:        "account_name",
				Description: "The name of this datadog account from your config.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAccountName,
			},
			{
				Name:        "attributes_created_at",
				Description: "Creation time of the user.",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("Attributes.CreatedAt"),
			},
			{
				Name:        "attributes_disabled",
				Description: "Whether the user is disabled.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Attributes.Disabled"),
			},
			{
				Name:        "attributes_email",
				Description: "Email of the user.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Attributes.Email"),
			},
			{
				Name:        "attributes_handle",
				Description: "Handle of the user.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Attributes.Handle"),
			},
			{
				Name:        "attributes_icon",
				Description: "URL of the user's icon.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Attributes.Icon"),
			},
			{
				Name:        "attributes_modified_at",
				Description: "Time that the user was last modified.",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("Attributes.ModifiedAt"),
			},
			{
				Name:        "attributes_service_account",
				Description: "Whether the user is a service account.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Attributes.ServiceAccount"),
			},
			{
				Name:        "attributes_status",
				Description: "Status of the user.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Attributes.Status"),
			},
			{
				Name:        "attributes_verified",
				Description: "Whether the user is verified.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Attributes.Verified"),
			},
			{
				Name:     "attributes_additional_properties",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Attributes.AdditionalProperties"),
			},
			{
				Name:        "id",
				Description: "ID of the user.",
				Type:        schema.TypeString,
			},
			{
				Name:        "relationships_org_data_id",
				Description: "ID of the organization.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Relationships.Org.Data.Id"),
			},
			{
				Name:        "relationships_org_data_type",
				Description: "Organizations resource type.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Relationships.Org.Data.Type"),
			},
			{
				Name:     "relationships_org_data_additional_properties",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Relationships.Org.Data.AdditionalProperties"),
			},
			{
				Name:     "relationships_org_additional_properties",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Relationships.Org.AdditionalProperties"),
			},
			{
				Name:     "relationships_other_orgs_additional_properties",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Relationships.OtherOrgs.AdditionalProperties"),
			},
			{
				Name:     "relationships_other_users_additional_properties",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Relationships.OtherUsers.AdditionalProperties"),
			},
			{
				Name:     "relationships_roles_additional_properties",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Relationships.Roles.AdditionalProperties"),
			},
			{
				Name:     "relationships_additional_properties",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Relationships.AdditionalProperties"),
			},
			{
				Name:        "type",
				Description: "Users resource type.",
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

func fetchUsers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	logger := c.Logger()
	// TODO: multiplexing
	thisConfig := c.MultiPlexedAccount.V2Config
	thisConfig.SetUnstableOperationEnabled("ListIncidents", true)
	apiClient := datadog.NewAPIClient(&thisConfig)

	resp, r, err := apiClient.UsersApi.ListUsers(c.MultiPlexedAccount.V2Context)
	logger.Debug(r.Status)
	if err != nil {
		return diag.FromError(err, diag.ACCESS)
	}
	res <- resp.Data
	return nil
}
