// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200401preview

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupPartnerTopicEventSubscription(ctx *pulumi.Context, args *LookupPartnerTopicEventSubscriptionArgs, opts ...pulumi.InvokeOption) (*LookupPartnerTopicEventSubscriptionResult, error) {
	var rv LookupPartnerTopicEventSubscriptionResult
	err := ctx.Invoke("azurerm:eventgrid/v20200401preview:getPartnerTopicEventSubscription", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPartnerTopicEventSubscriptionArgs struct {
	// Name of the event subscription to be found. Event subscription names must be between 3 and 100 characters in length and use alphanumeric letters only.
	EventSubscriptionName string `pulumi:"eventSubscriptionName"`
	// Name of the partner topic.
	PartnerTopicName string `pulumi:"partnerTopicName"`
	// The name of the resource group within the user's subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Event Subscription
type LookupPartnerTopicEventSubscriptionResult struct {
	// The dead letter destination of the event subscription. Any event that cannot be delivered to its' destination is sent to the dead letter destination.
	// Uses Azure Event Grid's identity to acquire the authentication tokens being used during delivery / dead-lettering.
	DeadLetterDestination *DeadLetterDestinationResponse `pulumi:"deadLetterDestination"`
	// The dead letter destination of the event subscription. Any event that cannot be delivered to its' destination is sent to the dead letter destination.
	// Uses the managed identity setup on the parent resource (namely, topic or domain) to acquire the authentication tokens being used during delivery / dead-lettering.
	DeadLetterWithResourceIdentity *DeadLetterWithResourceIdentityResponse `pulumi:"deadLetterWithResourceIdentity"`
	// Information about the destination where events have to be delivered for the event subscription.
	// Uses the managed identity setup on the parent resource (namely, topic or domain) to acquire the authentication tokens being used during delivery / dead-lettering.
	DeliveryWithResourceIdentity *DeliveryWithResourceIdentityResponse `pulumi:"deliveryWithResourceIdentity"`
	// Information about the destination where events have to be delivered for the event subscription.
	// Uses Azure Event Grid's identity to acquire the authentication tokens being used during delivery / dead-lettering.
	Destination *EventSubscriptionDestinationResponse `pulumi:"destination"`
	// The event delivery schema for the event subscription.
	EventDeliverySchema *string `pulumi:"eventDeliverySchema"`
	// Expiration time of the event subscription.
	ExpirationTimeUtc *string `pulumi:"expirationTimeUtc"`
	// Information about the filter for the event subscription.
	Filter *EventSubscriptionFilterResponse `pulumi:"filter"`
	// List of user defined labels.
	Labels []string `pulumi:"labels"`
	// Name of the resource
	Name string `pulumi:"name"`
	// Provisioning state of the event subscription.
	ProvisioningState string `pulumi:"provisioningState"`
	// The retry policy for events. This can be used to configure maximum number of delivery attempts and time to live for events.
	RetryPolicy *RetryPolicyResponse `pulumi:"retryPolicy"`
	// Name of the topic of the event subscription.
	Topic string `pulumi:"topic"`
	// Type of the resource
	Type string `pulumi:"type"`
}
