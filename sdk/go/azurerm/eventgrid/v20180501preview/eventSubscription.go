// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180501preview

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Event Subscription
type EventSubscription struct {
	pulumi.CustomResourceState

	// The DeadLetter destination of the event subscription.
	DeadLetterDestination DeadLetterDestinationResponsePtrOutput `pulumi:"deadLetterDestination"`
	// Information about the destination where events have to be delivered for the event subscription.
	Destination EventSubscriptionDestinationResponsePtrOutput `pulumi:"destination"`
	// The event delivery schema for the event subscription.
	EventDeliverySchema pulumi.StringPtrOutput `pulumi:"eventDeliverySchema"`
	// Information about the filter for the event subscription.
	Filter EventSubscriptionFilterResponsePtrOutput `pulumi:"filter"`
	// List of user defined labels.
	Labels pulumi.StringArrayOutput `pulumi:"labels"`
	// Name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Provisioning state of the event subscription.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The retry policy for events. This can be used to configure maximum number of delivery attempts and time to live for events.
	RetryPolicy RetryPolicyResponsePtrOutput `pulumi:"retryPolicy"`
	// Name of the topic of the event subscription.
	Topic pulumi.StringOutput `pulumi:"topic"`
	// Type of the resource
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewEventSubscription registers a new resource with the given unique name, arguments, and options.
func NewEventSubscription(ctx *pulumi.Context,
	name string, args *EventSubscriptionArgs, opts ...pulumi.ResourceOption) (*EventSubscription, error) {
	if args == nil || args.EventSubscriptionName == nil {
		return nil, errors.New("missing required argument 'EventSubscriptionName'")
	}
	if args == nil || args.Scope == nil {
		return nil, errors.New("missing required argument 'Scope'")
	}
	if args == nil {
		args = &EventSubscriptionArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:eventgrid/latest:EventSubscription"),
		},
		{
			Type: pulumi.String("azurerm:eventgrid/v20170615preview:EventSubscription"),
		},
		{
			Type: pulumi.String("azurerm:eventgrid/v20170915preview:EventSubscription"),
		},
		{
			Type: pulumi.String("azurerm:eventgrid/v20180101:EventSubscription"),
		},
		{
			Type: pulumi.String("azurerm:eventgrid/v20180915preview:EventSubscription"),
		},
		{
			Type: pulumi.String("azurerm:eventgrid/v20190101:EventSubscription"),
		},
		{
			Type: pulumi.String("azurerm:eventgrid/v20190201preview:EventSubscription"),
		},
		{
			Type: pulumi.String("azurerm:eventgrid/v20190601:EventSubscription"),
		},
		{
			Type: pulumi.String("azurerm:eventgrid/v20200101preview:EventSubscription"),
		},
		{
			Type: pulumi.String("azurerm:eventgrid/v20200401preview:EventSubscription"),
		},
		{
			Type: pulumi.String("azurerm:eventgrid/v20200601:EventSubscription"),
		},
	})
	opts = append(opts, aliases)
	var resource EventSubscription
	err := ctx.RegisterResource("azurerm:eventgrid/v20180501preview:EventSubscription", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEventSubscription gets an existing EventSubscription resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEventSubscription(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EventSubscriptionState, opts ...pulumi.ResourceOption) (*EventSubscription, error) {
	var resource EventSubscription
	err := ctx.ReadResource("azurerm:eventgrid/v20180501preview:EventSubscription", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EventSubscription resources.
type eventSubscriptionState struct {
	// The DeadLetter destination of the event subscription.
	DeadLetterDestination *DeadLetterDestinationResponse `pulumi:"deadLetterDestination"`
	// Information about the destination where events have to be delivered for the event subscription.
	Destination *EventSubscriptionDestinationResponse `pulumi:"destination"`
	// The event delivery schema for the event subscription.
	EventDeliverySchema *string `pulumi:"eventDeliverySchema"`
	// Information about the filter for the event subscription.
	Filter *EventSubscriptionFilterResponse `pulumi:"filter"`
	// List of user defined labels.
	Labels []string `pulumi:"labels"`
	// Name of the resource
	Name *string `pulumi:"name"`
	// Provisioning state of the event subscription.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The retry policy for events. This can be used to configure maximum number of delivery attempts and time to live for events.
	RetryPolicy *RetryPolicyResponse `pulumi:"retryPolicy"`
	// Name of the topic of the event subscription.
	Topic *string `pulumi:"topic"`
	// Type of the resource
	Type *string `pulumi:"type"`
}

type EventSubscriptionState struct {
	// The DeadLetter destination of the event subscription.
	DeadLetterDestination DeadLetterDestinationResponsePtrInput
	// Information about the destination where events have to be delivered for the event subscription.
	Destination EventSubscriptionDestinationResponsePtrInput
	// The event delivery schema for the event subscription.
	EventDeliverySchema pulumi.StringPtrInput
	// Information about the filter for the event subscription.
	Filter EventSubscriptionFilterResponsePtrInput
	// List of user defined labels.
	Labels pulumi.StringArrayInput
	// Name of the resource
	Name pulumi.StringPtrInput
	// Provisioning state of the event subscription.
	ProvisioningState pulumi.StringPtrInput
	// The retry policy for events. This can be used to configure maximum number of delivery attempts and time to live for events.
	RetryPolicy RetryPolicyResponsePtrInput
	// Name of the topic of the event subscription.
	Topic pulumi.StringPtrInput
	// Type of the resource
	Type pulumi.StringPtrInput
}

func (EventSubscriptionState) ElementType() reflect.Type {
	return reflect.TypeOf((*eventSubscriptionState)(nil)).Elem()
}

type eventSubscriptionArgs struct {
	// The DeadLetter destination of the event subscription.
	DeadLetterDestination *DeadLetterDestination `pulumi:"deadLetterDestination"`
	// Information about the destination where events have to be delivered for the event subscription.
	Destination *EventSubscriptionDestination `pulumi:"destination"`
	// The event delivery schema for the event subscription.
	EventDeliverySchema *string `pulumi:"eventDeliverySchema"`
	// Name of the event subscription. Event subscription names must be between 3 and 64 characters in length and should use alphanumeric letters only.
	EventSubscriptionName string `pulumi:"eventSubscriptionName"`
	// Information about the filter for the event subscription.
	Filter *EventSubscriptionFilter `pulumi:"filter"`
	// List of user defined labels.
	Labels []string `pulumi:"labels"`
	// The retry policy for events. This can be used to configure maximum number of delivery attempts and time to live for events.
	RetryPolicy *RetryPolicy `pulumi:"retryPolicy"`
	// The identifier of the resource to which the event subscription needs to be created or updated. The scope can be a subscription, or a resource group, or a top level resource belonging to a resource provider namespace, or an EventGrid topic. For example, use '/subscriptions/{subscriptionId}/' for a subscription, '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for a resource group, and '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}' for a resource, and '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/topics/{topicName}' for an EventGrid topic.
	Scope string `pulumi:"scope"`
}

// The set of arguments for constructing a EventSubscription resource.
type EventSubscriptionArgs struct {
	// The DeadLetter destination of the event subscription.
	DeadLetterDestination DeadLetterDestinationPtrInput
	// Information about the destination where events have to be delivered for the event subscription.
	Destination EventSubscriptionDestinationPtrInput
	// The event delivery schema for the event subscription.
	EventDeliverySchema pulumi.StringPtrInput
	// Name of the event subscription. Event subscription names must be between 3 and 64 characters in length and should use alphanumeric letters only.
	EventSubscriptionName pulumi.StringInput
	// Information about the filter for the event subscription.
	Filter EventSubscriptionFilterPtrInput
	// List of user defined labels.
	Labels pulumi.StringArrayInput
	// The retry policy for events. This can be used to configure maximum number of delivery attempts and time to live for events.
	RetryPolicy RetryPolicyPtrInput
	// The identifier of the resource to which the event subscription needs to be created or updated. The scope can be a subscription, or a resource group, or a top level resource belonging to a resource provider namespace, or an EventGrid topic. For example, use '/subscriptions/{subscriptionId}/' for a subscription, '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for a resource group, and '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}' for a resource, and '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/topics/{topicName}' for an EventGrid topic.
	Scope pulumi.StringInput
}

func (EventSubscriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*eventSubscriptionArgs)(nil)).Elem()
}
