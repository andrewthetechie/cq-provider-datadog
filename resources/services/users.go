package services

import (
	"context"
	"fmt"

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

	apiClient := datadog.NewAPIClient(&thisConfig)

	var page int64 = 1
	// it seems like pagination on the users api is totally broken.
	// It also seems its not super reliable. I get different counts for "TotalCount" on the meta depending on when
	// I call the API and that count differs from what I get in the users UI.

	//What I believe is happening is if you use a page size of greater than 1, if the count that would be returned for
	// that page is less than page size, some weird truncation happens. For example, using page size of 10, the last page
	// that would have 3 items on it for my org, is empty instead

	//After multiple experiments, I found the only way to get the API to actually list ALL users is to go page size of 1
	// and step through them one at a time. This makes retrieving users very slow
	params := datadog.NewListUsersOptionalParameters().WithPageNumber(page).WithPageSize(1)
	resp, r, err := apiClient.UsersApi.ListUsers(c.MultiPlexedAccount.V2Context, *params)
	logger.Debug(r.Status)
	if err != nil {
		return diag.FromError(err, diag.ACCESS)
	}
	retrievedUserCount := int64(len(resp.Data))
	res <- resp.Data
	var totalCount int64 = *resp.Meta.Page.TotalCount
	logger.Debug(fmt.Sprintf("TotalCount: %d", totalCount))
	logger.Debug(fmt.Sprintf("Length of Data %d", retrievedUserCount))
	// this pagination is awful, but the DD api is weird here. See above notes
	for retrievedUserCount < totalCount {
		page += 1
		logger.Debug(fmt.Sprintf("Looping user requests. Requesting page %d. Total User Count: %d, Total Retrieved: %d", page, totalCount, retrievedUserCount))
		params := datadog.NewListUsersOptionalParameters().WithPageNumber(page).WithPageSize(1)
		resp, r, err := apiClient.UsersApi.ListUsers(c.MultiPlexedAccount.V2Context, *params)
		thisUserCount := int64(len(resp.Data))
		retrievedUserCount += thisUserCount
		logger.Debug(fmt.Sprintf("Length of Data %d", thisUserCount))
		logger.Debug(fmt.Sprintf("Total users retrieved: %d", retrievedUserCount))
		logger.Debug(r.Status)
		if err != nil {
			return diag.FromError(err, diag.ACCESS)
		}
		res <- resp.Data
		if thisUserCount == 0 {
			break
		}
	}
	return nil
}
