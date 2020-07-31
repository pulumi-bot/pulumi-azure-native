// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20150819

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupService(ctx *pulumi.Context, args *LookupServiceArgs, opts ...pulumi.InvokeOption) (*LookupServiceResult, error) {
	var rv LookupServiceResult
	err := ctx.Invoke("azurerm:search/v20150819:getService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServiceArgs struct {
	// The name of the Azure Cognitive Search service associated with the specified resource group.
	Name string `pulumi:"name"`
	// The name of the resource group within the current subscription. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Describes an Azure Cognitive Search service and its current state.
type LookupServiceResult struct {
	// The identity of the resource.
	Identity *IdentityResponse `pulumi:"identity"`
	// The geographic location of the resource. This must be one of the supported and registered Azure Geo Regions (for example, West US, East US, Southeast Asia, and so forth). This property is required when creating a new resource.
	Location *string `pulumi:"location"`
	// The name of the resource.
	Name string `pulumi:"name"`
	// Properties of the Search service.
	Properties SearchServicePropertiesResponse `pulumi:"properties"`
	// The SKU of the Search Service, which determines price tier and capacity limits. This property is required when creating a new Search Service.
	Sku *SkuResponse `pulumi:"sku"`
	// Tags to help categorize the resource in the Azure portal.
	Tags map[string]string `pulumi:"tags"`
	// The resource type.
	Type string `pulumi:"type"`
}