// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20181201preview

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The properties of the EventHubConsumerGroupInfo object.
type IotHubResourceEventHubConsumerGroup struct {
	pulumi.CustomResourceState

	// The etag.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The Event Hub-compatible consumer group name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The tags.
	Properties pulumi.StringMapOutput `pulumi:"properties"`
	// the resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewIotHubResourceEventHubConsumerGroup registers a new resource with the given unique name, arguments, and options.
func NewIotHubResourceEventHubConsumerGroup(ctx *pulumi.Context,
	name string, args *IotHubResourceEventHubConsumerGroupArgs, opts ...pulumi.ResourceOption) (*IotHubResourceEventHubConsumerGroup, error) {
	if args == nil || args.EventHubEndpointName == nil {
		return nil, errors.New("missing required argument 'EventHubEndpointName'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ResourceName == nil {
		return nil, errors.New("missing required argument 'ResourceName'")
	}
	if args == nil {
		args = &IotHubResourceEventHubConsumerGroupArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:devices/latest:IotHubResourceEventHubConsumerGroup"),
		},
		{
			Type: pulumi.String("azurerm:devices/v20160203:IotHubResourceEventHubConsumerGroup"),
		},
		{
			Type: pulumi.String("azurerm:devices/v20170119:IotHubResourceEventHubConsumerGroup"),
		},
		{
			Type: pulumi.String("azurerm:devices/v20170701:IotHubResourceEventHubConsumerGroup"),
		},
		{
			Type: pulumi.String("azurerm:devices/v20180122:IotHubResourceEventHubConsumerGroup"),
		},
		{
			Type: pulumi.String("azurerm:devices/v20180401:IotHubResourceEventHubConsumerGroup"),
		},
		{
			Type: pulumi.String("azurerm:devices/v20190322:IotHubResourceEventHubConsumerGroup"),
		},
		{
			Type: pulumi.String("azurerm:devices/v20190322preview:IotHubResourceEventHubConsumerGroup"),
		},
		{
			Type: pulumi.String("azurerm:devices/v20190701preview:IotHubResourceEventHubConsumerGroup"),
		},
		{
			Type: pulumi.String("azurerm:devices/v20191104:IotHubResourceEventHubConsumerGroup"),
		},
		{
			Type: pulumi.String("azurerm:devices/v20200301:IotHubResourceEventHubConsumerGroup"),
		},
		{
			Type: pulumi.String("azurerm:devices/v20200401:IotHubResourceEventHubConsumerGroup"),
		},
		{
			Type: pulumi.String("azurerm:devices/v20200615:IotHubResourceEventHubConsumerGroup"),
		},
		{
			Type: pulumi.String("azurerm:devices/v20200710preview:IotHubResourceEventHubConsumerGroup"),
		},
		{
			Type: pulumi.String("azurerm:devices/v20200801:IotHubResourceEventHubConsumerGroup"),
		},
	})
	opts = append(opts, aliases)
	var resource IotHubResourceEventHubConsumerGroup
	err := ctx.RegisterResource("azurerm:devices/v20181201preview:IotHubResourceEventHubConsumerGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIotHubResourceEventHubConsumerGroup gets an existing IotHubResourceEventHubConsumerGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIotHubResourceEventHubConsumerGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IotHubResourceEventHubConsumerGroupState, opts ...pulumi.ResourceOption) (*IotHubResourceEventHubConsumerGroup, error) {
	var resource IotHubResourceEventHubConsumerGroup
	err := ctx.ReadResource("azurerm:devices/v20181201preview:IotHubResourceEventHubConsumerGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IotHubResourceEventHubConsumerGroup resources.
type iotHubResourceEventHubConsumerGroupState struct {
	// The etag.
	Etag *string `pulumi:"etag"`
	// The Event Hub-compatible consumer group name.
	Name *string `pulumi:"name"`
	// The tags.
	Properties map[string]string `pulumi:"properties"`
	// the resource type.
	Type *string `pulumi:"type"`
}

type IotHubResourceEventHubConsumerGroupState struct {
	// The etag.
	Etag pulumi.StringPtrInput
	// The Event Hub-compatible consumer group name.
	Name pulumi.StringPtrInput
	// The tags.
	Properties pulumi.StringMapInput
	// the resource type.
	Type pulumi.StringPtrInput
}

func (IotHubResourceEventHubConsumerGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*iotHubResourceEventHubConsumerGroupState)(nil)).Elem()
}

type iotHubResourceEventHubConsumerGroupArgs struct {
	// The name of the Event Hub-compatible endpoint in the IoT hub.
	EventHubEndpointName string `pulumi:"eventHubEndpointName"`
	// The name of the consumer group to add.
	Name string `pulumi:"name"`
	// The name of the resource group that contains the IoT hub.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the IoT hub.
	ResourceName string `pulumi:"resourceName"`
}

// The set of arguments for constructing a IotHubResourceEventHubConsumerGroup resource.
type IotHubResourceEventHubConsumerGroupArgs struct {
	// The name of the Event Hub-compatible endpoint in the IoT hub.
	EventHubEndpointName pulumi.StringInput
	// The name of the consumer group to add.
	Name pulumi.StringInput
	// The name of the resource group that contains the IoT hub.
	ResourceGroupName pulumi.StringInput
	// The name of the IoT hub.
	ResourceName pulumi.StringInput
}

func (IotHubResourceEventHubConsumerGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iotHubResourceEventHubConsumerGroupArgs)(nil)).Elem()
}
