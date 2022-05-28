package services

import (
	"context"
	"fmt"

	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"github.com/andrewthetechie/cq-provider-datadog/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource notebooks --config resources/services/notebooks.hcl --output .
func Notebooks() *schema.Table {
	return &schema.Table{
		Name:        "datadog_notebooks",
		Description: "NotebooksResponseData The data for a notebook in get all response.",
		Resolver:    fetchNotebooks,
		Options:     schema.TableCreationOptions{PrimaryKeys: []string{"id"}},
		Columns: []schema.Column{
			{
				Name:        "attributes_author_created_at",
				Description: "Creation time of the user.",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("Attributes.Author.CreatedAt"),
			},
			{
				Name:        "attributes_author_disabled",
				Description: "Whether the user is disabled.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Attributes.Author.Disabled"),
			},
			{
				Name:        "attributes_author_email",
				Description: "Email of the user.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Attributes.Author.Email"),
			},
			{
				Name:        "attributes_author_handle",
				Description: "Handle of the user.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Attributes.Author.Handle"),
			},
			{
				Name:        "attributes_author_icon",
				Description: "URL of the user's icon.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Attributes.Author.Icon"),
			},
			{
				Name:        "attributes_author_status",
				Description: "Status of the user.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Attributes.Author.Status"),
			},
			{
				Name:        "attributes_author_verified",
				Description: "Whether the user is verified.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Attributes.Author.Verified"),
			},
			{
				Name:     "attributes_author_additional_properties",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Attributes.Author.AdditionalProperties"),
			},
			{
				Name:        "attributes_created",
				Description: "UTC time stamp for when the notebook was created.",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("Attributes.Created"),
			},
			{
				Name:        "attributes_metadata_is_template",
				Description: "Whether or not the notebook is a template.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Attributes.Metadata.IsTemplate"),
			},
			{
				Name:        "attributes_metadata_take_snapshots",
				Description: "Whether or not the notebook takes snapshot image backups of the notebook's fixed-time graphs.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Attributes.Metadata.TakeSnapshots"),
			},
			{
				Name:     "attributes_metadata_additional_properties",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Attributes.Metadata.AdditionalProperties"),
			},
			{
				Name:        "attributes_modified",
				Description: "UTC time stamp for when the notebook was last modified.",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("Attributes.Modified"),
			},
			{
				Name:        "attributes_name",
				Description: "The name of the notebook.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Attributes.Name"),
			},
			{
				Name:        "attributes_status",
				Description: "Publication status of the notebook",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Attributes.Status"),
			},
			{
				Name:        "attributes_time_notebook_relative_time_live_span",
				Description: "The available timeframes depend on the widget you are using.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Attributes.Time.NotebookRelativeTime.LiveSpan"),
			},
			{
				Name:     "attributes_time_notebook_relative_time_additional_properties",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Attributes.Time.NotebookRelativeTime.AdditionalProperties"),
			},
			{
				Name:        "attributes_time_notebook_absolute_time_end",
				Description: "The end time.",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("Attributes.Time.NotebookAbsoluteTime.End"),
			},
			{
				Name:        "attributes_time_notebook_absolute_time_live",
				Description: "Indicates whether the timeframe should be shifted to end at the current time.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("Attributes.Time.NotebookAbsoluteTime.Live"),
			},
			{
				Name:        "attributes_time_notebook_absolute_time_start",
				Description: "The start time.",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("Attributes.Time.NotebookAbsoluteTime.Start"),
			},
			{
				Name:     "attributes_time_notebook_absolute_time_additional_properties",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Attributes.Time.NotebookAbsoluteTime.AdditionalProperties"),
			},
			{
				Name:     "attributes_additional_properties",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Attributes.AdditionalProperties"),
			},
			{
				Name:        "id",
				Description: "Unique notebook ID, assigned when you create the notebook.",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "type",
				Description: "Type of the Notebook resource.",
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

func fetchNotebooks(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	logger := c.Logger()
	logger.Debug("in fetchNotebooks")
	// TODO: multiplexing
	apiClient := datadog.NewAPIClient(&c.Accounts[0].V1Config)
	var step int64 = 250
	params := datadog.NewListNotebooksOptionalParameters().WithCount(step)
	resp, r, err := apiClient.NotebooksApi.ListNotebooks(c.Accounts[0].V1Context, *params)
	logger.Debug(r.Status)
	if err != nil {
		return diag.FromError(err, diag.ACCESS)
	}
	var retrievedCount int64 = *resp.Meta.Page.TotalFilteredCount
	var totalCount int64 = *resp.Meta.Page.TotalCount
	res <- resp.Data
	logger.Debug(fmt.Sprintf("TotalCount: %d, RetrievedCount: %d", totalCount, retrievedCount))
	for retrievedCount < totalCount {
		logger.Debug("Looping notebook requests")
		params := datadog.NewListNotebooksOptionalParameters().WithCount(step).WithStart(retrievedCount)
		resp, r, err := apiClient.NotebooksApi.ListNotebooks(c.Accounts[0].V1Context, *params)
		logger.Debug(r.Status)
		if err != nil {
			return diag.FromError(err, diag.ACCESS)
		}
		res <- resp.Data
		retrievedCount += *resp.Meta.Page.TotalFilteredCount
	}
	logger.Debug(fmt.Sprintf("Total Retrieved: %d", retrievedCount))
	return nil
}
