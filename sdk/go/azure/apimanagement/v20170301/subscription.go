// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20170301

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Subscription details.
type Subscription struct {
	pulumi.CustomResourceState

	// Subscription creation date. The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
	CreatedDate pulumi.StringOutput `pulumi:"createdDate"`
	// The name of the subscription, or null if the subscription has no name.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// Date when subscription was cancelled or expired. The setting is for audit purposes only and the subscription is not automatically cancelled. The subscription lifecycle can be managed by using the `state` property. The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
	EndDate pulumi.StringPtrOutput `pulumi:"endDate"`
	// Subscription expiration date. The setting is for audit purposes only and the subscription is not automatically expired. The subscription lifecycle can be managed by using the `state` property. The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
	ExpirationDate pulumi.StringPtrOutput `pulumi:"expirationDate"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Upcoming subscription expiration notification date. The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
	NotificationDate pulumi.StringPtrOutput `pulumi:"notificationDate"`
	// Subscription primary key.
	PrimaryKey pulumi.StringOutput `pulumi:"primaryKey"`
	// The product resource identifier of the subscribed product. The value is a valid relative URL in the format of /products/{productId} where {productId} is a product identifier.
	ProductId pulumi.StringOutput `pulumi:"productId"`
	// Subscription secondary key.
	SecondaryKey pulumi.StringOutput `pulumi:"secondaryKey"`
	// Subscription activation date. The setting is for audit purposes only and the subscription is not automatically activated. The subscription lifecycle can be managed by using the `state` property. The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
	StartDate pulumi.StringPtrOutput `pulumi:"startDate"`
	// Subscription state. Possible states are * active – the subscription is active, * suspended – the subscription is blocked, and the subscriber cannot call any APIs of the product, * submitted – the subscription request has been made by the developer, but has not yet been approved or rejected, * rejected – the subscription request has been denied by an administrator, * cancelled – the subscription has been cancelled by the developer or administrator, * expired – the subscription reached its expiration date and was deactivated.
	State pulumi.StringOutput `pulumi:"state"`
	// Optional subscription comment added by an administrator.
	StateComment pulumi.StringPtrOutput `pulumi:"stateComment"`
	// Resource type for API Management resource.
	Type pulumi.StringOutput `pulumi:"type"`
	// The user resource identifier of the subscription owner. The value is a valid relative URL in the format of /users/{uid} where {uid} is a user identifier.
	UserId pulumi.StringOutput `pulumi:"userId"`
}

// NewSubscription registers a new resource with the given unique name, arguments, and options.
func NewSubscription(ctx *pulumi.Context,
	name string, args *SubscriptionArgs, opts ...pulumi.ResourceOption) (*Subscription, error) {
	if args == nil || args.DisplayName == nil {
		return nil, errors.New("missing required argument 'DisplayName'")
	}
	if args == nil || args.ProductId == nil {
		return nil, errors.New("missing required argument 'ProductId'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ServiceName == nil {
		return nil, errors.New("missing required argument 'ServiceName'")
	}
	if args == nil || args.Sid == nil {
		return nil, errors.New("missing required argument 'Sid'")
	}
	if args == nil || args.UserId == nil {
		return nil, errors.New("missing required argument 'UserId'")
	}
	if args == nil {
		args = &SubscriptionArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:apimanagement/latest:Subscription"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20160707:Subscription"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20161010:Subscription"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20180101:Subscription"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20180601preview:Subscription"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20190101:Subscription"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20191201:Subscription"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20191201preview:Subscription"),
		},
		{
			Type: pulumi.String("azure-nextgen:apimanagement/v20200601preview:Subscription"),
		},
	})
	opts = append(opts, aliases)
	var resource Subscription
	err := ctx.RegisterResource("azure-nextgen:apimanagement/v20170301:Subscription", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSubscription gets an existing Subscription resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSubscription(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SubscriptionState, opts ...pulumi.ResourceOption) (*Subscription, error) {
	var resource Subscription
	err := ctx.ReadResource("azure-nextgen:apimanagement/v20170301:Subscription", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Subscription resources.
type subscriptionState struct {
	// Subscription creation date. The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
	CreatedDate *string `pulumi:"createdDate"`
	// The name of the subscription, or null if the subscription has no name.
	DisplayName *string `pulumi:"displayName"`
	// Date when subscription was cancelled or expired. The setting is for audit purposes only and the subscription is not automatically cancelled. The subscription lifecycle can be managed by using the `state` property. The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
	EndDate *string `pulumi:"endDate"`
	// Subscription expiration date. The setting is for audit purposes only and the subscription is not automatically expired. The subscription lifecycle can be managed by using the `state` property. The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
	ExpirationDate *string `pulumi:"expirationDate"`
	// Resource name.
	Name *string `pulumi:"name"`
	// Upcoming subscription expiration notification date. The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
	NotificationDate *string `pulumi:"notificationDate"`
	// Subscription primary key.
	PrimaryKey *string `pulumi:"primaryKey"`
	// The product resource identifier of the subscribed product. The value is a valid relative URL in the format of /products/{productId} where {productId} is a product identifier.
	ProductId *string `pulumi:"productId"`
	// Subscription secondary key.
	SecondaryKey *string `pulumi:"secondaryKey"`
	// Subscription activation date. The setting is for audit purposes only and the subscription is not automatically activated. The subscription lifecycle can be managed by using the `state` property. The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
	StartDate *string `pulumi:"startDate"`
	// Subscription state. Possible states are * active – the subscription is active, * suspended – the subscription is blocked, and the subscriber cannot call any APIs of the product, * submitted – the subscription request has been made by the developer, but has not yet been approved or rejected, * rejected – the subscription request has been denied by an administrator, * cancelled – the subscription has been cancelled by the developer or administrator, * expired – the subscription reached its expiration date and was deactivated.
	State *string `pulumi:"state"`
	// Optional subscription comment added by an administrator.
	StateComment *string `pulumi:"stateComment"`
	// Resource type for API Management resource.
	Type *string `pulumi:"type"`
	// The user resource identifier of the subscription owner. The value is a valid relative URL in the format of /users/{uid} where {uid} is a user identifier.
	UserId *string `pulumi:"userId"`
}

type SubscriptionState struct {
	// Subscription creation date. The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
	CreatedDate pulumi.StringPtrInput
	// The name of the subscription, or null if the subscription has no name.
	DisplayName pulumi.StringPtrInput
	// Date when subscription was cancelled or expired. The setting is for audit purposes only and the subscription is not automatically cancelled. The subscription lifecycle can be managed by using the `state` property. The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
	EndDate pulumi.StringPtrInput
	// Subscription expiration date. The setting is for audit purposes only and the subscription is not automatically expired. The subscription lifecycle can be managed by using the `state` property. The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
	ExpirationDate pulumi.StringPtrInput
	// Resource name.
	Name pulumi.StringPtrInput
	// Upcoming subscription expiration notification date. The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
	NotificationDate pulumi.StringPtrInput
	// Subscription primary key.
	PrimaryKey pulumi.StringPtrInput
	// The product resource identifier of the subscribed product. The value is a valid relative URL in the format of /products/{productId} where {productId} is a product identifier.
	ProductId pulumi.StringPtrInput
	// Subscription secondary key.
	SecondaryKey pulumi.StringPtrInput
	// Subscription activation date. The setting is for audit purposes only and the subscription is not automatically activated. The subscription lifecycle can be managed by using the `state` property. The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
	StartDate pulumi.StringPtrInput
	// Subscription state. Possible states are * active – the subscription is active, * suspended – the subscription is blocked, and the subscriber cannot call any APIs of the product, * submitted – the subscription request has been made by the developer, but has not yet been approved or rejected, * rejected – the subscription request has been denied by an administrator, * cancelled – the subscription has been cancelled by the developer or administrator, * expired – the subscription reached its expiration date and was deactivated.
	State pulumi.StringPtrInput
	// Optional subscription comment added by an administrator.
	StateComment pulumi.StringPtrInput
	// Resource type for API Management resource.
	Type pulumi.StringPtrInput
	// The user resource identifier of the subscription owner. The value is a valid relative URL in the format of /users/{uid} where {uid} is a user identifier.
	UserId pulumi.StringPtrInput
}

func (SubscriptionState) ElementType() reflect.Type {
	return reflect.TypeOf((*subscriptionState)(nil)).Elem()
}

type subscriptionArgs struct {
	// Subscription name.
	DisplayName string `pulumi:"displayName"`
	// Primary subscription key. If not specified during request key will be generated automatically.
	PrimaryKey *string `pulumi:"primaryKey"`
	// Product (product id path) for which subscription is being created in form /products/{productId}
	ProductId string `pulumi:"productId"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Secondary subscription key. If not specified during request key will be generated automatically.
	SecondaryKey *string `pulumi:"secondaryKey"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
	// Subscription entity Identifier. The entity represents the association between a user and a product in API Management.
	Sid string `pulumi:"sid"`
	// Initial subscription state. If no value is specified, subscription is created with Submitted state. Possible states are * active – the subscription is active, * suspended – the subscription is blocked, and the subscriber cannot call any APIs of the product, * submitted – the subscription request has been made by the developer, but has not yet been approved or rejected, * rejected – the subscription request has been denied by an administrator, * cancelled – the subscription has been cancelled by the developer or administrator, * expired – the subscription reached its expiration date and was deactivated.
	State *string `pulumi:"state"`
	// User (user id path) for whom subscription is being created in form /users/{uid}
	UserId string `pulumi:"userId"`
}

// The set of arguments for constructing a Subscription resource.
type SubscriptionArgs struct {
	// Subscription name.
	DisplayName pulumi.StringInput
	// Primary subscription key. If not specified during request key will be generated automatically.
	PrimaryKey pulumi.StringPtrInput
	// Product (product id path) for which subscription is being created in form /products/{productId}
	ProductId pulumi.StringInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Secondary subscription key. If not specified during request key will be generated automatically.
	SecondaryKey pulumi.StringPtrInput
	// The name of the API Management service.
	ServiceName pulumi.StringInput
	// Subscription entity Identifier. The entity represents the association between a user and a product in API Management.
	Sid pulumi.StringInput
	// Initial subscription state. If no value is specified, subscription is created with Submitted state. Possible states are * active – the subscription is active, * suspended – the subscription is blocked, and the subscriber cannot call any APIs of the product, * submitted – the subscription request has been made by the developer, but has not yet been approved or rejected, * rejected – the subscription request has been denied by an administrator, * cancelled – the subscription has been cancelled by the developer or administrator, * expired – the subscription reached its expiration date and was deactivated.
	State pulumi.StringPtrInput
	// User (user id path) for whom subscription is being created in form /users/{uid}
	UserId pulumi.StringInput
}

func (SubscriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*subscriptionArgs)(nil)).Elem()
}
