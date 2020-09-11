// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Information about the event hub destination for an event subscription
type EventHubEventSubscriptionDestination struct {
	// Type of the endpoint for the event subscription destination
	EndpointType string `pulumi:"endpointType"`
	// The Azure Resource Id that represents the endpoint of an Event Hub destination of an event subscription.
	ResourceId *string `pulumi:"resourceId"`
}

// EventHubEventSubscriptionDestinationInput is an input type that accepts EventHubEventSubscriptionDestinationArgs and EventHubEventSubscriptionDestinationOutput values.
// You can construct a concrete instance of `EventHubEventSubscriptionDestinationInput` via:
//
//          EventHubEventSubscriptionDestinationArgs{...}
type EventHubEventSubscriptionDestinationInput interface {
	pulumi.Input

	ToEventHubEventSubscriptionDestinationOutput() EventHubEventSubscriptionDestinationOutput
	ToEventHubEventSubscriptionDestinationOutputWithContext(context.Context) EventHubEventSubscriptionDestinationOutput
}

// Information about the event hub destination for an event subscription
type EventHubEventSubscriptionDestinationArgs struct {
	// Type of the endpoint for the event subscription destination
	EndpointType pulumi.StringInput `pulumi:"endpointType"`
	// The Azure Resource Id that represents the endpoint of an Event Hub destination of an event subscription.
	ResourceId pulumi.StringPtrInput `pulumi:"resourceId"`
}

func (EventHubEventSubscriptionDestinationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EventHubEventSubscriptionDestination)(nil)).Elem()
}

func (i EventHubEventSubscriptionDestinationArgs) ToEventHubEventSubscriptionDestinationOutput() EventHubEventSubscriptionDestinationOutput {
	return i.ToEventHubEventSubscriptionDestinationOutputWithContext(context.Background())
}

func (i EventHubEventSubscriptionDestinationArgs) ToEventHubEventSubscriptionDestinationOutputWithContext(ctx context.Context) EventHubEventSubscriptionDestinationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventHubEventSubscriptionDestinationOutput)
}

// Information about the event hub destination for an event subscription
type EventHubEventSubscriptionDestinationResponse struct {
	// Type of the endpoint for the event subscription destination
	EndpointType string `pulumi:"endpointType"`
	// The Azure Resource Id that represents the endpoint of an Event Hub destination of an event subscription.
	ResourceId *string `pulumi:"resourceId"`
}

// Information about the event hub destination for an event subscription
type EventHubEventSubscriptionDestinationResponseOutput struct{ *pulumi.OutputState }

func (EventHubEventSubscriptionDestinationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventHubEventSubscriptionDestinationResponse)(nil)).Elem()
}

func (o EventHubEventSubscriptionDestinationResponseOutput) ToEventHubEventSubscriptionDestinationResponseOutput() EventHubEventSubscriptionDestinationResponseOutput {
	return o
}

func (o EventHubEventSubscriptionDestinationResponseOutput) ToEventHubEventSubscriptionDestinationResponseOutputWithContext(ctx context.Context) EventHubEventSubscriptionDestinationResponseOutput {
	return o
}

// Type of the endpoint for the event subscription destination
func (o EventHubEventSubscriptionDestinationResponseOutput) EndpointType() pulumi.StringOutput {
	return o.ApplyT(func(v EventHubEventSubscriptionDestinationResponse) string { return v.EndpointType }).(pulumi.StringOutput)
}

// The Azure Resource Id that represents the endpoint of an Event Hub destination of an event subscription.
func (o EventHubEventSubscriptionDestinationResponseOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventHubEventSubscriptionDestinationResponse) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

// Filter for the Event Subscription
type EventSubscriptionFilter struct {
	// A list of applicable event types that need to be part of the event subscription.
	// If it is desired to subscribe to all event types, the string "all" needs to be specified as an element in this list.
	IncludedEventTypes []string `pulumi:"includedEventTypes"`
	// Specifies if the SubjectBeginsWith and SubjectEndsWith properties of the filter
	// should be compared in a case sensitive manner.
	IsSubjectCaseSensitive *bool `pulumi:"isSubjectCaseSensitive"`
	// An optional string to filter events for an event subscription based on a resource path prefix.
	// The format of this depends on the publisher of the events.
	// Wildcard characters are not supported in this path.
	SubjectBeginsWith *string `pulumi:"subjectBeginsWith"`
	// An optional string to filter events for an event subscription based on a resource path suffix.
	// Wildcard characters are not supported in this path.
	SubjectEndsWith *string `pulumi:"subjectEndsWith"`
}

// EventSubscriptionFilterInput is an input type that accepts EventSubscriptionFilterArgs and EventSubscriptionFilterOutput values.
// You can construct a concrete instance of `EventSubscriptionFilterInput` via:
//
//          EventSubscriptionFilterArgs{...}
type EventSubscriptionFilterInput interface {
	pulumi.Input

	ToEventSubscriptionFilterOutput() EventSubscriptionFilterOutput
	ToEventSubscriptionFilterOutputWithContext(context.Context) EventSubscriptionFilterOutput
}

// Filter for the Event Subscription
type EventSubscriptionFilterArgs struct {
	// A list of applicable event types that need to be part of the event subscription.
	// If it is desired to subscribe to all event types, the string "all" needs to be specified as an element in this list.
	IncludedEventTypes pulumi.StringArrayInput `pulumi:"includedEventTypes"`
	// Specifies if the SubjectBeginsWith and SubjectEndsWith properties of the filter
	// should be compared in a case sensitive manner.
	IsSubjectCaseSensitive pulumi.BoolPtrInput `pulumi:"isSubjectCaseSensitive"`
	// An optional string to filter events for an event subscription based on a resource path prefix.
	// The format of this depends on the publisher of the events.
	// Wildcard characters are not supported in this path.
	SubjectBeginsWith pulumi.StringPtrInput `pulumi:"subjectBeginsWith"`
	// An optional string to filter events for an event subscription based on a resource path suffix.
	// Wildcard characters are not supported in this path.
	SubjectEndsWith pulumi.StringPtrInput `pulumi:"subjectEndsWith"`
}

func (EventSubscriptionFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EventSubscriptionFilter)(nil)).Elem()
}

func (i EventSubscriptionFilterArgs) ToEventSubscriptionFilterOutput() EventSubscriptionFilterOutput {
	return i.ToEventSubscriptionFilterOutputWithContext(context.Background())
}

func (i EventSubscriptionFilterArgs) ToEventSubscriptionFilterOutputWithContext(ctx context.Context) EventSubscriptionFilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventSubscriptionFilterOutput)
}

func (i EventSubscriptionFilterArgs) ToEventSubscriptionFilterPtrOutput() EventSubscriptionFilterPtrOutput {
	return i.ToEventSubscriptionFilterPtrOutputWithContext(context.Background())
}

func (i EventSubscriptionFilterArgs) ToEventSubscriptionFilterPtrOutputWithContext(ctx context.Context) EventSubscriptionFilterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventSubscriptionFilterOutput).ToEventSubscriptionFilterPtrOutputWithContext(ctx)
}

// EventSubscriptionFilterPtrInput is an input type that accepts EventSubscriptionFilterArgs, EventSubscriptionFilterPtr and EventSubscriptionFilterPtrOutput values.
// You can construct a concrete instance of `EventSubscriptionFilterPtrInput` via:
//
//          EventSubscriptionFilterArgs{...}
//
//  or:
//
//          nil
type EventSubscriptionFilterPtrInput interface {
	pulumi.Input

	ToEventSubscriptionFilterPtrOutput() EventSubscriptionFilterPtrOutput
	ToEventSubscriptionFilterPtrOutputWithContext(context.Context) EventSubscriptionFilterPtrOutput
}

type eventSubscriptionFilterPtrType EventSubscriptionFilterArgs

func EventSubscriptionFilterPtr(v *EventSubscriptionFilterArgs) EventSubscriptionFilterPtrInput {
	return (*eventSubscriptionFilterPtrType)(v)
}

func (*eventSubscriptionFilterPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EventSubscriptionFilter)(nil)).Elem()
}

func (i *eventSubscriptionFilterPtrType) ToEventSubscriptionFilterPtrOutput() EventSubscriptionFilterPtrOutput {
	return i.ToEventSubscriptionFilterPtrOutputWithContext(context.Background())
}

func (i *eventSubscriptionFilterPtrType) ToEventSubscriptionFilterPtrOutputWithContext(ctx context.Context) EventSubscriptionFilterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventSubscriptionFilterPtrOutput)
}

// Filter for the Event Subscription
type EventSubscriptionFilterResponse struct {
	// A list of applicable event types that need to be part of the event subscription.
	// If it is desired to subscribe to all event types, the string "all" needs to be specified as an element in this list.
	IncludedEventTypes []string `pulumi:"includedEventTypes"`
	// Specifies if the SubjectBeginsWith and SubjectEndsWith properties of the filter
	// should be compared in a case sensitive manner.
	IsSubjectCaseSensitive *bool `pulumi:"isSubjectCaseSensitive"`
	// An optional string to filter events for an event subscription based on a resource path prefix.
	// The format of this depends on the publisher of the events.
	// Wildcard characters are not supported in this path.
	SubjectBeginsWith *string `pulumi:"subjectBeginsWith"`
	// An optional string to filter events for an event subscription based on a resource path suffix.
	// Wildcard characters are not supported in this path.
	SubjectEndsWith *string `pulumi:"subjectEndsWith"`
}

// Filter for the Event Subscription
type EventSubscriptionFilterResponseOutput struct{ *pulumi.OutputState }

func (EventSubscriptionFilterResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventSubscriptionFilterResponse)(nil)).Elem()
}

func (o EventSubscriptionFilterResponseOutput) ToEventSubscriptionFilterResponseOutput() EventSubscriptionFilterResponseOutput {
	return o
}

func (o EventSubscriptionFilterResponseOutput) ToEventSubscriptionFilterResponseOutputWithContext(ctx context.Context) EventSubscriptionFilterResponseOutput {
	return o
}

func (o EventSubscriptionFilterResponseOutput) ToEventSubscriptionFilterResponsePtrOutput() EventSubscriptionFilterResponsePtrOutput {
	return o.ToEventSubscriptionFilterResponsePtrOutputWithContext(context.Background())
}

func (o EventSubscriptionFilterResponseOutput) ToEventSubscriptionFilterResponsePtrOutputWithContext(ctx context.Context) EventSubscriptionFilterResponsePtrOutput {
	return o.ApplyT(func(v EventSubscriptionFilterResponse) *EventSubscriptionFilterResponse {
		return &v
	}).(EventSubscriptionFilterResponsePtrOutput)
}

// A list of applicable event types that need to be part of the event subscription.
// If it is desired to subscribe to all event types, the string "all" needs to be specified as an element in this list.
func (o EventSubscriptionFilterResponseOutput) IncludedEventTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v EventSubscriptionFilterResponse) []string { return v.IncludedEventTypes }).(pulumi.StringArrayOutput)
}

// Specifies if the SubjectBeginsWith and SubjectEndsWith properties of the filter
// should be compared in a case sensitive manner.
func (o EventSubscriptionFilterResponseOutput) IsSubjectCaseSensitive() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EventSubscriptionFilterResponse) *bool { return v.IsSubjectCaseSensitive }).(pulumi.BoolPtrOutput)
}

// An optional string to filter events for an event subscription based on a resource path prefix.
// The format of this depends on the publisher of the events.
// Wildcard characters are not supported in this path.
func (o EventSubscriptionFilterResponseOutput) SubjectBeginsWith() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventSubscriptionFilterResponse) *string { return v.SubjectBeginsWith }).(pulumi.StringPtrOutput)
}

// An optional string to filter events for an event subscription based on a resource path suffix.
// Wildcard characters are not supported in this path.
func (o EventSubscriptionFilterResponseOutput) SubjectEndsWith() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventSubscriptionFilterResponse) *string { return v.SubjectEndsWith }).(pulumi.StringPtrOutput)
}

type EventSubscriptionFilterResponsePtrOutput struct{ *pulumi.OutputState }

func (EventSubscriptionFilterResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EventSubscriptionFilterResponse)(nil)).Elem()
}

func (o EventSubscriptionFilterResponsePtrOutput) ToEventSubscriptionFilterResponsePtrOutput() EventSubscriptionFilterResponsePtrOutput {
	return o
}

func (o EventSubscriptionFilterResponsePtrOutput) ToEventSubscriptionFilterResponsePtrOutputWithContext(ctx context.Context) EventSubscriptionFilterResponsePtrOutput {
	return o
}

func (o EventSubscriptionFilterResponsePtrOutput) Elem() EventSubscriptionFilterResponseOutput {
	return o.ApplyT(func(v *EventSubscriptionFilterResponse) EventSubscriptionFilterResponse { return *v }).(EventSubscriptionFilterResponseOutput)
}

// A list of applicable event types that need to be part of the event subscription.
// If it is desired to subscribe to all event types, the string "all" needs to be specified as an element in this list.
func (o EventSubscriptionFilterResponsePtrOutput) IncludedEventTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *EventSubscriptionFilterResponse) []string {
		if v == nil {
			return nil
		}
		return v.IncludedEventTypes
	}).(pulumi.StringArrayOutput)
}

// Specifies if the SubjectBeginsWith and SubjectEndsWith properties of the filter
// should be compared in a case sensitive manner.
func (o EventSubscriptionFilterResponsePtrOutput) IsSubjectCaseSensitive() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EventSubscriptionFilterResponse) *bool {
		if v == nil {
			return nil
		}
		return v.IsSubjectCaseSensitive
	}).(pulumi.BoolPtrOutput)
}

// An optional string to filter events for an event subscription based on a resource path prefix.
// The format of this depends on the publisher of the events.
// Wildcard characters are not supported in this path.
func (o EventSubscriptionFilterResponsePtrOutput) SubjectBeginsWith() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventSubscriptionFilterResponse) *string {
		if v == nil {
			return nil
		}
		return v.SubjectBeginsWith
	}).(pulumi.StringPtrOutput)
}

// An optional string to filter events for an event subscription based on a resource path suffix.
// Wildcard characters are not supported in this path.
func (o EventSubscriptionFilterResponsePtrOutput) SubjectEndsWith() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventSubscriptionFilterResponse) *string {
		if v == nil {
			return nil
		}
		return v.SubjectEndsWith
	}).(pulumi.StringPtrOutput)
}

// Information about the webhook destination for an event subscription
type WebHookEventSubscriptionDestination struct {
	// Type of the endpoint for the event subscription destination
	EndpointType string `pulumi:"endpointType"`
	// The URL that represents the endpoint of the destination of an event subscription.
	EndpointUrl *string `pulumi:"endpointUrl"`
}

// WebHookEventSubscriptionDestinationInput is an input type that accepts WebHookEventSubscriptionDestinationArgs and WebHookEventSubscriptionDestinationOutput values.
// You can construct a concrete instance of `WebHookEventSubscriptionDestinationInput` via:
//
//          WebHookEventSubscriptionDestinationArgs{...}
type WebHookEventSubscriptionDestinationInput interface {
	pulumi.Input

	ToWebHookEventSubscriptionDestinationOutput() WebHookEventSubscriptionDestinationOutput
	ToWebHookEventSubscriptionDestinationOutputWithContext(context.Context) WebHookEventSubscriptionDestinationOutput
}

// Information about the webhook destination for an event subscription
type WebHookEventSubscriptionDestinationArgs struct {
	// Type of the endpoint for the event subscription destination
	EndpointType pulumi.StringInput `pulumi:"endpointType"`
	// The URL that represents the endpoint of the destination of an event subscription.
	EndpointUrl pulumi.StringPtrInput `pulumi:"endpointUrl"`
}

func (WebHookEventSubscriptionDestinationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WebHookEventSubscriptionDestination)(nil)).Elem()
}

func (i WebHookEventSubscriptionDestinationArgs) ToWebHookEventSubscriptionDestinationOutput() WebHookEventSubscriptionDestinationOutput {
	return i.ToWebHookEventSubscriptionDestinationOutputWithContext(context.Background())
}

func (i WebHookEventSubscriptionDestinationArgs) ToWebHookEventSubscriptionDestinationOutputWithContext(ctx context.Context) WebHookEventSubscriptionDestinationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebHookEventSubscriptionDestinationOutput)
}

// Information about the webhook destination for an event subscription
type WebHookEventSubscriptionDestinationResponse struct {
	// The base URL that represents the endpoint of the destination of an event subscription.
	EndpointBaseUrl string `pulumi:"endpointBaseUrl"`
	// Type of the endpoint for the event subscription destination
	EndpointType string `pulumi:"endpointType"`
	// The URL that represents the endpoint of the destination of an event subscription.
	EndpointUrl *string `pulumi:"endpointUrl"`
}

// Information about the webhook destination for an event subscription
type WebHookEventSubscriptionDestinationResponseOutput struct{ *pulumi.OutputState }

func (WebHookEventSubscriptionDestinationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebHookEventSubscriptionDestinationResponse)(nil)).Elem()
}

func (o WebHookEventSubscriptionDestinationResponseOutput) ToWebHookEventSubscriptionDestinationResponseOutput() WebHookEventSubscriptionDestinationResponseOutput {
	return o
}

func (o WebHookEventSubscriptionDestinationResponseOutput) ToWebHookEventSubscriptionDestinationResponseOutputWithContext(ctx context.Context) WebHookEventSubscriptionDestinationResponseOutput {
	return o
}

// The base URL that represents the endpoint of the destination of an event subscription.
func (o WebHookEventSubscriptionDestinationResponseOutput) EndpointBaseUrl() pulumi.StringOutput {
	return o.ApplyT(func(v WebHookEventSubscriptionDestinationResponse) string { return v.EndpointBaseUrl }).(pulumi.StringOutput)
}

// Type of the endpoint for the event subscription destination
func (o WebHookEventSubscriptionDestinationResponseOutput) EndpointType() pulumi.StringOutput {
	return o.ApplyT(func(v WebHookEventSubscriptionDestinationResponse) string { return v.EndpointType }).(pulumi.StringOutput)
}

// The URL that represents the endpoint of the destination of an event subscription.
func (o WebHookEventSubscriptionDestinationResponseOutput) EndpointUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WebHookEventSubscriptionDestinationResponse) *string { return v.EndpointUrl }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(EventHubEventSubscriptionDestinationOutput{})
	pulumi.RegisterOutputType(EventHubEventSubscriptionDestinationResponseOutput{})
	pulumi.RegisterOutputType(EventSubscriptionFilterOutput{})
	pulumi.RegisterOutputType(EventSubscriptionFilterPtrOutput{})
	pulumi.RegisterOutputType(EventSubscriptionFilterResponseOutput{})
	pulumi.RegisterOutputType(EventSubscriptionFilterResponsePtrOutput{})
	pulumi.RegisterOutputType(WebHookEventSubscriptionDestinationOutput{})
	pulumi.RegisterOutputType(WebHookEventSubscriptionDestinationResponseOutput{})
}
