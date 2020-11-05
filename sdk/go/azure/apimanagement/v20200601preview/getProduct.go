// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200601preview

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupProduct(ctx *pulumi.Context, args *LookupProductArgs, opts ...pulumi.InvokeOption) (*LookupProductResult, error) {
	var rv LookupProductResult
	err := ctx.Invoke("azure-nextgen:apimanagement/v20200601preview:getProduct", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupProductArgs struct {
	// Product identifier. Must be unique in the current API Management service instance.
	ProductId string `pulumi:"productId"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
}

// Product details.
type LookupProductResult struct {
	// whether subscription approval is required. If false, new subscriptions will be approved automatically enabling developers to call the product’s APIs immediately after subscribing. If true, administrators must manually approve the subscription before the developer can any of the product’s APIs. Can be present only if subscriptionRequired property is present and has a value of false.
	ApprovalRequired *bool `pulumi:"approvalRequired"`
	// Product description. May include HTML formatting tags.
	Description *string `pulumi:"description"`
	// Product name.
	DisplayName string `pulumi:"displayName"`
	// Resource name.
	Name string `pulumi:"name"`
	// whether product is published or not. Published products are discoverable by users of developer portal. Non published products are visible only to administrators. Default state of Product is notPublished.
	State *string `pulumi:"state"`
	// Whether a product subscription is required for accessing APIs included in this product. If true, the product is referred to as "protected" and a valid subscription key is required for a request to an API included in the product to succeed. If false, the product is referred to as "open" and requests to an API included in the product can be made without a subscription key. If property is omitted when creating a new product it's value is assumed to be true.
	SubscriptionRequired *bool `pulumi:"subscriptionRequired"`
	// Whether the number of subscriptions a user can have to this product at the same time. Set to null or omit to allow unlimited per user subscriptions. Can be present only if subscriptionRequired property is present and has a value of false.
	SubscriptionsLimit *int `pulumi:"subscriptionsLimit"`
	// Product terms of use. Developers trying to subscribe to the product will be presented and required to accept these terms before they can complete the subscription process.
	Terms *string `pulumi:"terms"`
	// Resource type for API Management resource.
	Type string `pulumi:"type"`
}
