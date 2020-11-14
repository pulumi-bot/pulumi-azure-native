// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// An environment receives data from one or more event sources. Each event source has associated connection info that allows the Time Series Insights ingress pipeline to connect to and pull data from the event source
type EventSource struct {
	pulumi.CustomResourceState

	// The kind of the event source.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// Resource location
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewEventSource registers a new resource with the given unique name, arguments, and options.
func NewEventSource(ctx *pulumi.Context,
	name string, args *EventSourceArgs, opts ...pulumi.ResourceOption) (*EventSource, error) {
	if args == nil || args.EnvironmentName == nil {
		return nil, errors.New("missing required argument 'EnvironmentName'")
	}
	if args == nil || args.EventSourceName == nil {
		return nil, errors.New("missing required argument 'EventSourceName'")
	}
	if args == nil || args.Kind == nil {
		return nil, errors.New("missing required argument 'Kind'")
	}
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &EventSourceArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:timeseriesinsights/v20170228preview:EventSource"),
		},
		{
			Type: pulumi.String("azure-nextgen:timeseriesinsights/v20171115:EventSource"),
		},
		{
			Type: pulumi.String("azure-nextgen:timeseriesinsights/v20180815preview:EventSource"),
		},
		{
			Type: pulumi.String("azure-nextgen:timeseriesinsights/v20200515:EventSource"),
		},
	})
	opts = append(opts, aliases)
	var resource EventSource
	err := ctx.RegisterResource("azure-nextgen:timeseriesinsights/latest:EventSource", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEventSource gets an existing EventSource resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEventSource(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EventSourceState, opts ...pulumi.ResourceOption) (*EventSource, error) {
	var resource EventSource
	err := ctx.ReadResource("azure-nextgen:timeseriesinsights/latest:EventSource", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EventSource resources.
type eventSourceState struct {
	// The kind of the event source.
	Kind *string `pulumi:"kind"`
	// Resource location
	Location *string `pulumi:"location"`
	// Resource name
	Name *string `pulumi:"name"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type *string `pulumi:"type"`
}

type EventSourceState struct {
	// The kind of the event source.
	Kind pulumi.StringPtrInput
	// Resource location
	Location pulumi.StringPtrInput
	// Resource name
	Name pulumi.StringPtrInput
	// Resource tags
	Tags pulumi.StringMapInput
	// Resource type
	Type pulumi.StringPtrInput
}

func (EventSourceState) ElementType() reflect.Type {
	return reflect.TypeOf((*eventSourceState)(nil)).Elem()
}

type eventSourceArgs struct {
	// The name of the Time Series Insights environment associated with the specified resource group.
	EnvironmentName string `pulumi:"environmentName"`
	// Name of the event source.
	EventSourceName string `pulumi:"eventSourceName"`
	// The kind of the event source.
	Kind string `pulumi:"kind"`
	// An object that represents the local timestamp property. It contains the format of local timestamp that needs to be used and the corresponding timezone offset information. If a value isn't specified for localTimestamp, or if null, then the local timestamp will not be ingressed with the events.
	LocalTimestamp *LocalTimestamp `pulumi:"localTimestamp"`
	// The location of the resource.
	Location string `pulumi:"location"`
	// Name of an Azure Resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Key-value pairs of additional properties for the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a EventSource resource.
type EventSourceArgs struct {
	// The name of the Time Series Insights environment associated with the specified resource group.
	EnvironmentName pulumi.StringInput
	// Name of the event source.
	EventSourceName pulumi.StringInput
	// The kind of the event source.
	Kind pulumi.StringInput
	// An object that represents the local timestamp property. It contains the format of local timestamp that needs to be used and the corresponding timezone offset information. If a value isn't specified for localTimestamp, or if null, then the local timestamp will not be ingressed with the events.
	LocalTimestamp LocalTimestampPtrInput
	// The location of the resource.
	Location pulumi.StringInput
	// Name of an Azure Resource group.
	ResourceGroupName pulumi.StringInput
	// Key-value pairs of additional properties for the resource.
	Tags pulumi.StringMapInput
}

func (EventSourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*eventSourceArgs)(nil)).Elem()
}

type EventSourceInput interface {
	pulumi.Input

	ToEventSourceOutput() EventSourceOutput
	ToEventSourceOutputWithContext(ctx context.Context) EventSourceOutput
}

func (EventSource) ElementType() reflect.Type {
	return reflect.TypeOf((*EventSource)(nil)).Elem()
}

func (i EventSource) ToEventSourceOutput() EventSourceOutput {
	return i.ToEventSourceOutputWithContext(context.Background())
}

func (i EventSource) ToEventSourceOutputWithContext(ctx context.Context) EventSourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventSourceOutput)
}

type EventSourceOutput struct {
	*pulumi.OutputState
}

func (EventSourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventSourceOutput)(nil)).Elem()
}

func (o EventSourceOutput) ToEventSourceOutput() EventSourceOutput {
	return o
}

func (o EventSourceOutput) ToEventSourceOutputWithContext(ctx context.Context) EventSourceOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(EventSourceOutput{})
}
