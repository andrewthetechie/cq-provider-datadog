package services

import (
	"context"

	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
	"github.com/andrewthetechie/cq-provider-datadog/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource synthetics --config ./resources/services/v1/synthetics.hcl --output .
func Synthetics() *schema.Table {
	return &schema.Table{
		Name:        "datadog_synthetics",
		Description: "SyntheticsTestDetails Object containing details about your Synthetic test.",
		Resolver:    fetchSynthetics,
		Options:     schema.TableCreationOptions{PrimaryKeys: []string{"public_id"}},
		Columns: []schema.Column{
			{
				Name:        "config",
				Description: "Configuration object for a Synthetic test.",
				Type:        schema.TypeJSON,
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
				Name:        "locations",
				Description: "Array of locations used to run the test.",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "message",
				Description: "Notification message associated with the test.",
				Type:        schema.TypeString,
			},
			{
				Name:        "monitor_id",
				Description: "The associated monitor ID.",
				Type:        schema.TypeBigInt,
			},
			{
				Name:        "name",
				Description: "Name of the test.",
				Type:        schema.TypeString,
			},
			{
				Name:        "options",
				Description: "Object describing the extra options for a Synthetic test.",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "public_id",
				Description: "The test public ID.",
				Type:        schema.TypeString,
			},
			{
				Name:        "status",
				Description: "Define whether you want to start (`live`) or pause (`paused`) a Synthetic test.",
				Type:        schema.TypeString,
			},
			{
				Name:        "subtype",
				Description: "The subtype of the Synthetic API test, `http`, `ssl`, `tcp`, `dns`, `icmp`, `udp`, `websocket` or `multi`.",
				Type:        schema.TypeString,
			},
			{
				Name:        "tags",
				Description: "Array of tags attached to the test.",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "type",
				Description: "Type of the Synthetic test, either `api` or `browser`.",
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

func fetchSynthetics(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	logger := c.Logger()
	// TODO: multiplexing
	apiClient := datadog.NewAPIClient(&c.Accounts[0].V1Config)
	resp, r, err := apiClient.SyntheticsApi.ListTests(c.Accounts[0].V1Context)
	logger.Debug(r.Status)
	if err != nil {
		return diag.FromError(err, diag.ACCESS)
	}
	res <- resp.Tests
	return nil
}
