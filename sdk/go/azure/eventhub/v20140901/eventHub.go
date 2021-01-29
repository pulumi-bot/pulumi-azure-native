// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20140901

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Single item in List or Get Event Hub operation
type EventHub struct {
	pulumi.CustomResourceState

	// Exact time the Event Hub was created.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Resource location
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Number of days to retain the events for this Event Hub.
	MessageRetentionInDays pulumi.Float64PtrOutput `pulumi:"messageRetentionInDays"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Number of partitions created for the Event Hub.
	PartitionCount pulumi.Float64PtrOutput `pulumi:"partitionCount"`
	// Current number of shards on the Event Hub.
	PartitionIds pulumi.StringArrayOutput `pulumi:"partitionIds"`
	// Enumerates the possible values for the status of the Event Hub.
	Status pulumi.StringPtrOutput `pulumi:"status"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
	// The exact time the message was updated.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
}

// NewEventHub registers a new resource with the given unique name, arguments, and options.
func NewEventHub(ctx *pulumi.Context,
	name string, args *EventHubArgs, opts ...pulumi.ResourceOption) (*EventHub, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EventHubName == nil {
		return nil, errors.New("invalid value for required argument 'EventHubName'")
	}
	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.NamespaceName == nil {
		return nil, errors.New("invalid value for required argument 'NamespaceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:eventhub/latest:EventHub"),
		},
		{
			Type: pulumi.String("azure-nextgen:eventhub/v20150801:EventHub"),
		},
		{
			Type: pulumi.String("azure-nextgen:eventhub/v20170401:EventHub"),
		},
		{
			Type: pulumi.String("azure-nextgen:eventhub/v20180101preview:EventHub"),
		},
	})
	opts = append(opts, aliases)
	var resource EventHub
	err := ctx.RegisterResource("azure-nextgen:eventhub/v20140901:EventHub", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEventHub gets an existing EventHub resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEventHub(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EventHubState, opts ...pulumi.ResourceOption) (*EventHub, error) {
	var resource EventHub
	err := ctx.ReadResource("azure-nextgen:eventhub/v20140901:EventHub", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EventHub resources.
type eventHubState struct {
	// Exact time the Event Hub was created.
	CreatedAt *string `pulumi:"createdAt"`
	// Resource location
	Location *string `pulumi:"location"`
	// Number of days to retain the events for this Event Hub.
	MessageRetentionInDays *float64 `pulumi:"messageRetentionInDays"`
	// Resource name
	Name *string `pulumi:"name"`
	// Number of partitions created for the Event Hub.
	PartitionCount *float64 `pulumi:"partitionCount"`
	// Current number of shards on the Event Hub.
	PartitionIds []string `pulumi:"partitionIds"`
	// Enumerates the possible values for the status of the Event Hub.
	Status *string `pulumi:"status"`
	// Resource type
	Type *string `pulumi:"type"`
	// The exact time the message was updated.
	UpdatedAt *string `pulumi:"updatedAt"`
}

type EventHubState struct {
	// Exact time the Event Hub was created.
	CreatedAt pulumi.StringPtrInput
	// Resource location
	Location pulumi.StringPtrInput
	// Number of days to retain the events for this Event Hub.
	MessageRetentionInDays pulumi.Float64PtrInput
	// Resource name
	Name pulumi.StringPtrInput
	// Number of partitions created for the Event Hub.
	PartitionCount pulumi.Float64PtrInput
	// Current number of shards on the Event Hub.
	PartitionIds pulumi.StringArrayInput
	// Enumerates the possible values for the status of the Event Hub.
	Status pulumi.StringPtrInput
	// Resource type
	Type pulumi.StringPtrInput
	// The exact time the message was updated.
	UpdatedAt pulumi.StringPtrInput
}

func (EventHubState) ElementType() reflect.Type {
	return reflect.TypeOf((*eventHubState)(nil)).Elem()
}

type eventHubArgs struct {
	// The Event Hub name
	EventHubName string `pulumi:"eventHubName"`
	// Location of the resource.
	Location string `pulumi:"location"`
	// Number of days to retain the events for this Event Hub.
	MessageRetentionInDays *float64 `pulumi:"messageRetentionInDays"`
	// Name of the Event Hub.
	Name *string `pulumi:"name"`
	// The Namespace name
	NamespaceName string `pulumi:"namespaceName"`
	// Number of partitions created for the Event Hub.
	PartitionCount *float64 `pulumi:"partitionCount"`
	// Name of the resource group within the azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Enumerates the possible values for the status of the Event Hub.
	Status *string `pulumi:"status"`
	// ARM type of the Namespace.
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a EventHub resource.
type EventHubArgs struct {
	// The Event Hub name
	EventHubName pulumi.StringInput
	// Location of the resource.
	Location pulumi.StringInput
	// Number of days to retain the events for this Event Hub.
	MessageRetentionInDays pulumi.Float64PtrInput
	// Name of the Event Hub.
	Name pulumi.StringPtrInput
	// The Namespace name
	NamespaceName pulumi.StringInput
	// Number of partitions created for the Event Hub.
	PartitionCount pulumi.Float64PtrInput
	// Name of the resource group within the azure subscription.
	ResourceGroupName pulumi.StringInput
	// Enumerates the possible values for the status of the Event Hub.
	Status *EntityStatus
	// ARM type of the Namespace.
	Type pulumi.StringPtrInput
}

func (EventHubArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*eventHubArgs)(nil)).Elem()
}

type EventHubInput interface {
	pulumi.Input

	ToEventHubOutput() EventHubOutput
	ToEventHubOutputWithContext(ctx context.Context) EventHubOutput
}

func (*EventHub) ElementType() reflect.Type {
	return reflect.TypeOf((*EventHub)(nil))
}

func (i *EventHub) ToEventHubOutput() EventHubOutput {
	return i.ToEventHubOutputWithContext(context.Background())
}

func (i *EventHub) ToEventHubOutputWithContext(ctx context.Context) EventHubOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventHubOutput)
}

type EventHubOutput struct {
	*pulumi.OutputState
}

func (EventHubOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventHub)(nil))
}

func (o EventHubOutput) ToEventHubOutput() EventHubOutput {
	return o
}

func (o EventHubOutput) ToEventHubOutputWithContext(ctx context.Context) EventHubOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(EventHubOutput{})
}
