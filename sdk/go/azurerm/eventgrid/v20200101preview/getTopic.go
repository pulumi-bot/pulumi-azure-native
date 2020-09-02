// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200101preview

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupTopic(ctx *pulumi.Context, args *LookupTopicArgs, opts ...pulumi.InvokeOption) (*LookupTopicResult, error) {
	var rv LookupTopicResult
	err := ctx.Invoke("azurerm:eventgrid/v20200101preview:getTopic", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTopicArgs struct {
	// The name of the resource group within the user's subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the topic
	TopicName string `pulumi:"topicName"`
}

// EventGrid Topic
type LookupTopicResult struct {
	// Endpoint for the topic.
	Endpoint string `pulumi:"endpoint"`
	// This determines the format that Event Grid should expect for incoming events published to the topic.
	InputSchema *string `pulumi:"inputSchema"`
	// This enables publishing using custom event schemas. An InputSchemaMapping can be specified to map various properties of a source schema to various required properties of the EventGridEvent schema.
	InputSchemaMapping *InputSchemaMappingResponse `pulumi:"inputSchemaMapping"`
	// Location of the resource
	Location string `pulumi:"location"`
	// Metric resource id for the topic.
	MetricResourceId string `pulumi:"metricResourceId"`
	// Name of the resource
	Name string `pulumi:"name"`
	// Provisioning state of the topic.
	ProvisioningState string `pulumi:"provisioningState"`
	// Tags of the resource
	Tags map[string]string `pulumi:"tags"`
	// Type of the resource
	Type string `pulumi:"type"`
}
