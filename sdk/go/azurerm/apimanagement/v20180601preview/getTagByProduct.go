// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180601preview

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupTagByProduct(ctx *pulumi.Context, args *LookupTagByProductArgs, opts ...pulumi.InvokeOption) (*LookupTagByProductResult, error) {
	var rv LookupTagByProductResult
	err := ctx.Invoke("azurerm:apimanagement/v20180601preview:getTagByProduct", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTagByProductArgs struct {
	// Product identifier. Must be unique in the current API Management service instance.
	ProductId string `pulumi:"productId"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
	// Tag identifier. Must be unique in the current API Management service instance.
	TagId string `pulumi:"tagId"`
}

// Tag Contract details.
type LookupTagByProductResult struct {
	// Tag name.
	DisplayName string `pulumi:"displayName"`
	// Resource name.
	Name string `pulumi:"name"`
	// Resource type for API Management resource.
	Type string `pulumi:"type"`
}
