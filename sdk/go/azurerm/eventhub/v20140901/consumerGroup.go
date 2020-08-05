// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20140901

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Single item in List or Get Consumer group operation
type ConsumerGroup struct {
	pulumi.CustomResourceState

	// Resource location
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Properties supplied to the Create Or Update Consumer Group operation.
	Properties ConsumerGroupPropertiesResponseOutput `pulumi:"properties"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewConsumerGroup registers a new resource with the given unique name, arguments, and options.
func NewConsumerGroup(ctx *pulumi.Context,
	name string, args *ConsumerGroupArgs, opts ...pulumi.ResourceOption) (*ConsumerGroup, error) {
	if args == nil || args.EventHubName == nil {
		return nil, errors.New("missing required argument 'EventHubName'")
	}
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.NamespaceName == nil {
		return nil, errors.New("missing required argument 'NamespaceName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &ConsumerGroupArgs{}
	}
	var resource ConsumerGroup
	err := ctx.RegisterResource("azurerm:eventhub/v20140901:ConsumerGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConsumerGroup gets an existing ConsumerGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConsumerGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConsumerGroupState, opts ...pulumi.ResourceOption) (*ConsumerGroup, error) {
	var resource ConsumerGroup
	err := ctx.ReadResource("azurerm:eventhub/v20140901:ConsumerGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ConsumerGroup resources.
type consumerGroupState struct {
	// Resource location
	Location *string `pulumi:"location"`
	// Resource name
	Name *string `pulumi:"name"`
	// Properties supplied to the Create Or Update Consumer Group operation.
	Properties *ConsumerGroupPropertiesResponse `pulumi:"properties"`
	// Resource type
	Type *string `pulumi:"type"`
}

type ConsumerGroupState struct {
	// Resource location
	Location pulumi.StringPtrInput
	// Resource name
	Name pulumi.StringPtrInput
	// Properties supplied to the Create Or Update Consumer Group operation.
	Properties ConsumerGroupPropertiesResponsePtrInput
	// Resource type
	Type pulumi.StringPtrInput
}

func (ConsumerGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*consumerGroupState)(nil)).Elem()
}

type consumerGroupArgs struct {
	// The Event Hub name
	EventHubName string `pulumi:"eventHubName"`
	// Location of the resource.
	Location string `pulumi:"location"`
	// The consumer group name
	Name string `pulumi:"name"`
	// The Namespace name
	NamespaceName string `pulumi:"namespaceName"`
	// Name of the resource group within the azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// ARM type of the Namespace.
	Type *string `pulumi:"type"`
	// The user metadata.
	UserMetadata *string `pulumi:"userMetadata"`
}

// The set of arguments for constructing a ConsumerGroup resource.
type ConsumerGroupArgs struct {
	// The Event Hub name
	EventHubName pulumi.StringInput
	// Location of the resource.
	Location pulumi.StringInput
	// The consumer group name
	Name pulumi.StringInput
	// The Namespace name
	NamespaceName pulumi.StringInput
	// Name of the resource group within the azure subscription.
	ResourceGroupName pulumi.StringInput
	// ARM type of the Namespace.
	Type pulumi.StringPtrInput
	// The user metadata.
	UserMetadata pulumi.StringPtrInput
}

func (ConsumerGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*consumerGroupArgs)(nil)).Elem()
}
