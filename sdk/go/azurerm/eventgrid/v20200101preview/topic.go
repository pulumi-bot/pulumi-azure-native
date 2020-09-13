// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200101preview

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// EventGrid Topic
//
// ## Example Usage
// ### Topics_CreateOrUpdate
//
// ```go
// package main
//
// import (
// 	eventgrid "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/eventgrid/v20200101preview"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := eventgrid.NewTopic(ctx, "topic", &eventgrid.TopicArgs{
// 			Location:          pulumi.String("westus2"),
// 			ResourceGroupName: pulumi.String("examplerg"),
// 			Tags: pulumi.StringMap{
// 				"tag1": pulumi.String("value1"),
// 				"tag2": pulumi.String("value2"),
// 			},
// 			TopicName: pulumi.String("exampletopic1"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
type Topic struct {
	pulumi.CustomResourceState

	// Endpoint for the topic.
	Endpoint pulumi.StringOutput `pulumi:"endpoint"`
	// This determines the format that Event Grid should expect for incoming events published to the topic.
	InputSchema pulumi.StringPtrOutput `pulumi:"inputSchema"`
	// This enables publishing using custom event schemas. An InputSchemaMapping can be specified to map various properties of a source schema to various required properties of the EventGridEvent schema.
	InputSchemaMapping JsonInputSchemaMappingResponsePtrOutput `pulumi:"inputSchemaMapping"`
	// Location of the resource
	Location pulumi.StringOutput `pulumi:"location"`
	// Metric resource id for the topic.
	MetricResourceId pulumi.StringOutput `pulumi:"metricResourceId"`
	// Name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Provisioning state of the topic.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Tags of the resource
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Type of the resource
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewTopic registers a new resource with the given unique name, arguments, and options.
func NewTopic(ctx *pulumi.Context,
	name string, args *TopicArgs, opts ...pulumi.ResourceOption) (*Topic, error) {
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.TopicName == nil {
		return nil, errors.New("missing required argument 'TopicName'")
	}
	if args == nil {
		args = &TopicArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:eventgrid/latest:Topic"),
		},
		{
			Type: pulumi.String("azurerm:eventgrid/v20170615preview:Topic"),
		},
		{
			Type: pulumi.String("azurerm:eventgrid/v20170915preview:Topic"),
		},
		{
			Type: pulumi.String("azurerm:eventgrid/v20180101:Topic"),
		},
		{
			Type: pulumi.String("azurerm:eventgrid/v20180501preview:Topic"),
		},
		{
			Type: pulumi.String("azurerm:eventgrid/v20180915preview:Topic"),
		},
		{
			Type: pulumi.String("azurerm:eventgrid/v20190101:Topic"),
		},
		{
			Type: pulumi.String("azurerm:eventgrid/v20190201preview:Topic"),
		},
		{
			Type: pulumi.String("azurerm:eventgrid/v20190601:Topic"),
		},
		{
			Type: pulumi.String("azurerm:eventgrid/v20200401preview:Topic"),
		},
		{
			Type: pulumi.String("azurerm:eventgrid/v20200601:Topic"),
		},
	})
	opts = append(opts, aliases)
	var resource Topic
	err := ctx.RegisterResource("azurerm:eventgrid/v20200101preview:Topic", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTopic gets an existing Topic resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTopic(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TopicState, opts ...pulumi.ResourceOption) (*Topic, error) {
	var resource Topic
	err := ctx.ReadResource("azurerm:eventgrid/v20200101preview:Topic", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Topic resources.
type topicState struct {
	// Endpoint for the topic.
	Endpoint *string `pulumi:"endpoint"`
	// This determines the format that Event Grid should expect for incoming events published to the topic.
	InputSchema *string `pulumi:"inputSchema"`
	// This enables publishing using custom event schemas. An InputSchemaMapping can be specified to map various properties of a source schema to various required properties of the EventGridEvent schema.
	InputSchemaMapping *JsonInputSchemaMappingResponse `pulumi:"inputSchemaMapping"`
	// Location of the resource
	Location *string `pulumi:"location"`
	// Metric resource id for the topic.
	MetricResourceId *string `pulumi:"metricResourceId"`
	// Name of the resource
	Name *string `pulumi:"name"`
	// Provisioning state of the topic.
	ProvisioningState *string `pulumi:"provisioningState"`
	// Tags of the resource
	Tags map[string]string `pulumi:"tags"`
	// Type of the resource
	Type *string `pulumi:"type"`
}

type TopicState struct {
	// Endpoint for the topic.
	Endpoint pulumi.StringPtrInput
	// This determines the format that Event Grid should expect for incoming events published to the topic.
	InputSchema pulumi.StringPtrInput
	// This enables publishing using custom event schemas. An InputSchemaMapping can be specified to map various properties of a source schema to various required properties of the EventGridEvent schema.
	InputSchemaMapping JsonInputSchemaMappingResponsePtrInput
	// Location of the resource
	Location pulumi.StringPtrInput
	// Metric resource id for the topic.
	MetricResourceId pulumi.StringPtrInput
	// Name of the resource
	Name pulumi.StringPtrInput
	// Provisioning state of the topic.
	ProvisioningState pulumi.StringPtrInput
	// Tags of the resource
	Tags pulumi.StringMapInput
	// Type of the resource
	Type pulumi.StringPtrInput
}

func (TopicState) ElementType() reflect.Type {
	return reflect.TypeOf((*topicState)(nil)).Elem()
}

type topicArgs struct {
	// This determines the format that Event Grid should expect for incoming events published to the topic.
	InputSchema *string `pulumi:"inputSchema"`
	// This enables publishing using custom event schemas. An InputSchemaMapping can be specified to map various properties of a source schema to various required properties of the EventGridEvent schema.
	InputSchemaMapping *JsonInputSchemaMapping `pulumi:"inputSchemaMapping"`
	// Location of the resource
	Location string `pulumi:"location"`
	// The name of the resource group within the user's subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Tags of the resource
	Tags map[string]string `pulumi:"tags"`
	// Name of the topic
	TopicName string `pulumi:"topicName"`
}

// The set of arguments for constructing a Topic resource.
type TopicArgs struct {
	// This determines the format that Event Grid should expect for incoming events published to the topic.
	InputSchema pulumi.StringPtrInput
	// This enables publishing using custom event schemas. An InputSchemaMapping can be specified to map various properties of a source schema to various required properties of the EventGridEvent schema.
	InputSchemaMapping JsonInputSchemaMappingPtrInput
	// Location of the resource
	Location pulumi.StringInput
	// The name of the resource group within the user's subscription.
	ResourceGroupName pulumi.StringInput
	// Tags of the resource
	Tags pulumi.StringMapInput
	// Name of the topic
	TopicName pulumi.StringInput
}

func (TopicArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*topicArgs)(nil)).Elem()
}
