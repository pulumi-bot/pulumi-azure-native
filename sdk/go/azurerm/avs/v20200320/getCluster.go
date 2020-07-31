// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200320

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupCluster(ctx *pulumi.Context, args *LookupClusterArgs, opts ...pulumi.InvokeOption) (*LookupClusterResult, error) {
	var rv LookupClusterResult
	err := ctx.Invoke("azurerm:avs/v20200320:getCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupClusterArgs struct {
	// Name of the cluster in the private cloud
	Name string `pulumi:"name"`
	// Name of the private cloud
	PrivateCloudName string `pulumi:"privateCloudName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A cluster resource
type LookupClusterResult struct {
	// Resource name.
	Name string `pulumi:"name"`
	// The properties of a cluster resource
	Properties ClusterPropertiesResponse `pulumi:"properties"`
	// The cluster SKU
	Sku SkuResponse `pulumi:"sku"`
	// Resource type.
	Type string `pulumi:"type"`
}