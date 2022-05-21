package services

import (
	"context"
	"fmt"

	datadog "github.com/DataDog/datadog-api-client-go/api/v2/datadog"
	"github.com/andrewthetechie/cq-provider-datadog/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource incidents --config ./resources/services/v2/incidents.hcl --output .
func Incidents() *schema.Table {
	return &schema.Table{
		Name:        "datadog_incidents",
		Description: "IncidentResponseData Incident data from a response.",
		Resolver:    fetchIncidents,
		Options:     schema.TableCreationOptions{PrimaryKeys: []string{"id"}},
		Columns: []schema.Column{
			{
				Name:        "attributes_created",
				Description: "Timestamp when the incident was created.",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("Attributes.Created"),
			},
			{
				Name:        "attributes_customer_impact_duration",
				Description: "Length of the incident's customer impact in seconds. Equals the difference between `customer_impact_start` and `customer_impact_end`.",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("Attributes.CustomerImpactDuration"),
			},
			{
				Name:        "attributes_customer_impacted",
				Description: "A flag indicating whether the incident caused customer impact.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Attributes.CustomerImpacted"),
			},
			{
				Name:        "attributes_fields",
				Description: "A condensed view of the user-defined fields attached to incidents.",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("Attributes.Fields"),
			},
			{
				Name:        "attributes_modified",
				Description: "Timestamp when the incident was last modified.",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("Attributes.Modified"),
			},
			{
				Name:        "attributes_postmortem_id",
				Description: "The UUID of the postmortem object attached to the incident.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Attributes.PostmortemId"),
			},
			{
				Name:        "attributes_public_id",
				Description: "The monotonically increasing integer ID for the incident.",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("Attributes.PublicId"),
			},
			{
				Name:        "attributes_time_to_detect",
				Description: "The amount of time in seconds to detect the incident. Equals the difference between `customer_impact_start` and `detected`.",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("Attributes.TimeToDetect"),
			},
			{
				Name:        "attributes_time_to_internal_response",
				Description: "The amount of time in seconds to call incident after detection",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("Attributes.TimeToInternalResponse"),
			},
			{
				Name:        "attributes_time_to_repair",
				Description: "The amount of time in seconds to resolve customer impact after detecting the issue",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("Attributes.TimeToRepair"),
			},
			{
				Name:        "attributes_time_to_resolve",
				Description: "The amount of time in seconds to resolve the incident after it was created",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("Attributes.TimeToResolve"),
			},
			{
				Name:        "attributes_title",
				Description: "The title of the incident, which summarizes what happened.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Attributes.Title"),
			},
			{
				Name:     "attributes_additional_properties",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Attributes.AdditionalProperties"),
			},
			{
				Name:        "id",
				Description: "The incident's ID.",
				Type:        schema.TypeString,
			},
			{
				Name:     "relationships_commander_user_additional_properties",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Relationships.CommanderUser.AdditionalProperties"),
			},
			{
				Name:        "relationships_created_by_user_data_id",
				Description: "A unique identifier that represents the user.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Relationships.CreatedByUser.Data.Id"),
			},
			{
				Name:        "relationships_created_by_user_data_type",
				Description: "Users resource type.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Relationships.CreatedByUser.Data.Type"),
			},
			{
				Name:     "relationships_created_by_user_data_additional_properties",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Relationships.CreatedByUser.Data.AdditionalProperties"),
			},
			{
				Name:     "relationships_created_by_user_additional_properties",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Relationships.CreatedByUser.AdditionalProperties"),
			},
			{
				Name:     "relationships_integrations_additional_properties",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Relationships.Integrations.AdditionalProperties"),
			},
			{
				Name:        "relationships_last_modified_by_user_data_id",
				Description: "A unique identifier that represents the user.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Relationships.LastModifiedByUser.Data.Id"),
			},
			{
				Name:        "relationships_last_modified_by_user_data_type",
				Description: "Users resource type.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Relationships.LastModifiedByUser.Data.Type"),
			},
			{
				Name:     "relationships_last_modified_by_user_data_additional_properties",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Relationships.LastModifiedByUser.Data.AdditionalProperties"),
			},
			{
				Name:     "relationships_last_modified_by_user_additional_properties",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Relationships.LastModifiedByUser.AdditionalProperties"),
			},
			{
				Name:        "relationships_postmortem_data_id",
				Description: "A unique identifier that represents the postmortem.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Relationships.Postmortem.Data.Id"),
			},
			{
				Name:        "relationships_postmortem_data_type",
				Description: "Incident postmortem resource type.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Relationships.Postmortem.Data.Type"),
			},
			{
				Name:     "relationships_postmortem_data_additional_properties",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Relationships.Postmortem.Data.AdditionalProperties"),
			},
			{
				Name:     "relationships_postmortem_additional_properties",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Relationships.Postmortem.AdditionalProperties"),
			},
			{
				Name:     "relationships_additional_properties",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Relationships.AdditionalProperties"),
			},
			{
				Name:        "type",
				Description: "Incident resource type.",
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

func fetchIncidents(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	logger := c.Logger()
	logger.Debug("in fetchHosts")
	// TODO: multiplexing
	thisConfig := c.Accounts[0].V2Config
	thisConfig.SetUnstableOperationEnabled("ListIncidents", true)
	apiClient := datadog.NewAPIClient(&thisConfig)

	var step int64 = 1000
	params := datadog.NewListIncidentsOptionalParameters().WithPageSize(step)
	resp, r, err := apiClient.IncidentsApi.ListIncidents(c.Accounts[0].V2Context, *params)
	logger.Debug(r.Status)
	if err != nil {
		return diag.FromError(err, diag.ACCESS)
	}
	res <- resp.Data
	logger.Debug(fmt.Sprintf("Offset: %d, Next Offset: %d Size: %d", *resp.Meta.Pagination.Offset, *resp.Meta.Pagination.NextOffset, *resp.Meta.Pagination.Size))
	for *resp.Meta.Pagination.NextOffset > *resp.Meta.Pagination.Size {
		logger.Debug("Looping host requests")
		params := datadog.NewListIncidentsOptionalParameters().WithPageSize(step).WithPageOffset(*resp.Meta.Pagination.NextOffset)
		resp, r, err := apiClient.IncidentsApi.ListIncidents(c.Accounts[0].V2Context, *params)
		logger.Debug(r.Status)
		logger.Debug(fmt.Sprintf("Offset: %d, Next Offset: %d Size: %d", *resp.Meta.Pagination.Offset, *resp.Meta.Pagination.NextOffset, *resp.Meta.Pagination.Size))
		if err != nil {
			return diag.FromError(err, diag.ACCESS)
		}
		res <- resp.Data
	}
	return nil
}