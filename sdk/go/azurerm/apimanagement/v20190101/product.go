// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190101

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Product details.
type Product struct {
	pulumi.CustomResourceState

	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Product entity contract properties.
	Properties ProductContractPropertiesResponseOutput `pulumi:"properties"`
	// Resource type for API Management resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewProduct registers a new resource with the given unique name, arguments, and options.
func NewProduct(ctx *pulumi.Context,
	name string, args *ProductArgs, opts ...pulumi.ResourceOption) (*Product, error) {
	if args == nil || args.DisplayName == nil {
		return nil, errors.New("missing required argument 'DisplayName'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ServiceName == nil {
		return nil, errors.New("missing required argument 'ServiceName'")
	}
	if args == nil {
		args = &ProductArgs{}
	}
	var resource Product
	err := ctx.RegisterResource("azurerm:apimanagement/v20190101:Product", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProduct gets an existing Product resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProduct(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProductState, opts ...pulumi.ResourceOption) (*Product, error) {
	var resource Product
	err := ctx.ReadResource("azurerm:apimanagement/v20190101:Product", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Product resources.
type productState struct {
	// Resource name.
	Name *string `pulumi:"name"`
	// Product entity contract properties.
	Properties *ProductContractPropertiesResponse `pulumi:"properties"`
	// Resource type for API Management resource.
	Type *string `pulumi:"type"`
}

type ProductState struct {
	// Resource name.
	Name pulumi.StringPtrInput
	// Product entity contract properties.
	Properties ProductContractPropertiesResponsePtrInput
	// Resource type for API Management resource.
	Type pulumi.StringPtrInput
}

func (ProductState) ElementType() reflect.Type {
	return reflect.TypeOf((*productState)(nil)).Elem()
}

type productArgs struct {
	// whether subscription approval is required. If false, new subscriptions will be approved automatically enabling developers to call the product’s APIs immediately after subscribing. If true, administrators must manually approve the subscription before the developer can any of the product’s APIs. Can be present only if subscriptionRequired property is present and has a value of false.
	ApprovalRequired *bool `pulumi:"approvalRequired"`
	// Product description. May include HTML formatting tags.
	Description *string `pulumi:"description"`
	// Product name.
	DisplayName string `pulumi:"displayName"`
	// Product identifier. Must be unique in the current API Management service instance.
	Name string `pulumi:"name"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
	// whether product is published or not. Published products are discoverable by users of developer portal. Non published products are visible only to administrators. Default state of Product is notPublished.
	State *string `pulumi:"state"`
	// Whether a product subscription is required for accessing APIs included in this product. If true, the product is referred to as "protected" and a valid subscription key is required for a request to an API included in the product to succeed. If false, the product is referred to as "open" and requests to an API included in the product can be made without a subscription key. If property is omitted when creating a new product it's value is assumed to be true.
	SubscriptionRequired *bool `pulumi:"subscriptionRequired"`
	// Whether the number of subscriptions a user can have to this product at the same time. Set to null or omit to allow unlimited per user subscriptions. Can be present only if subscriptionRequired property is present and has a value of false.
	SubscriptionsLimit *int `pulumi:"subscriptionsLimit"`
	// Product terms of use. Developers trying to subscribe to the product will be presented and required to accept these terms before they can complete the subscription process.
	Terms *string `pulumi:"terms"`
}

// The set of arguments for constructing a Product resource.
type ProductArgs struct {
	// whether subscription approval is required. If false, new subscriptions will be approved automatically enabling developers to call the product’s APIs immediately after subscribing. If true, administrators must manually approve the subscription before the developer can any of the product’s APIs. Can be present only if subscriptionRequired property is present and has a value of false.
	ApprovalRequired pulumi.BoolPtrInput
	// Product description. May include HTML formatting tags.
	Description pulumi.StringPtrInput
	// Product name.
	DisplayName pulumi.StringInput
	// Product identifier. Must be unique in the current API Management service instance.
	Name pulumi.StringInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The name of the API Management service.
	ServiceName pulumi.StringInput
	// whether product is published or not. Published products are discoverable by users of developer portal. Non published products are visible only to administrators. Default state of Product is notPublished.
	State pulumi.StringPtrInput
	// Whether a product subscription is required for accessing APIs included in this product. If true, the product is referred to as "protected" and a valid subscription key is required for a request to an API included in the product to succeed. If false, the product is referred to as "open" and requests to an API included in the product can be made without a subscription key. If property is omitted when creating a new product it's value is assumed to be true.
	SubscriptionRequired pulumi.BoolPtrInput
	// Whether the number of subscriptions a user can have to this product at the same time. Set to null or omit to allow unlimited per user subscriptions. Can be present only if subscriptionRequired property is present and has a value of false.
	SubscriptionsLimit pulumi.IntPtrInput
	// Product terms of use. Developers trying to subscribe to the product will be presented and required to accept these terms before they can complete the subscription process.
	Terms pulumi.StringPtrInput
}

func (ProductArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*productArgs)(nil)).Elem()
}
