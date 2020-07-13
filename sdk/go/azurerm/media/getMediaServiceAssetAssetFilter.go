// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package media

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupMediaServiceAssetAssetFilter(ctx *pulumi.Context, args *LookupMediaServiceAssetAssetFilterArgs, opts ...pulumi.InvokeOption) (*LookupMediaServiceAssetAssetFilterResult, error) {
	var rv LookupMediaServiceAssetAssetFilterResult
	err := ctx.Invoke("azurerm:media:getMediaServiceAssetAssetFilter", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMediaServiceAssetAssetFilterArgs struct {
	// The Media Services account name.
	AccountName string `pulumi:"accountName"`
	// The Asset name.
	AssetName string `pulumi:"assetName"`
	// The Asset Filter name
	Name string `pulumi:"name"`
	// The name of the resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// An Asset Filter.
type LookupMediaServiceAssetAssetFilterResult struct {
	// The name of the resource
	Name string `pulumi:"name"`
	// The Media Filter properties.
	Properties MediaFilterPropertiesResponse `pulumi:"properties"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type string `pulumi:"type"`
}