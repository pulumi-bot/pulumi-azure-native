// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190809preview

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupCluster(ctx *pulumi.Context, args *LookupClusterArgs, opts ...pulumi.InvokeOption) (*LookupClusterResult, error) {
	var rv LookupClusterResult
	err := ctx.Invoke("azurerm:avs/v20190809preview:getCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupClusterArgs struct {
	// Name of the cluster in the private cloud
	ClusterName string `pulumi:"clusterName"`
	// Name of the private cloud
	PrivateCloudName string `pulumi:"privateCloudName"`
	// Name of the resource group within the Azure subscription
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A cluster resource
type LookupClusterResult struct {
	// Resource name.
	Name string `pulumi:"name"`
	// The properties of a cluster resource
	Properties ClusterPropertiesResponse `pulumi:"properties"`
	// Resource type.
	Type string `pulumi:"type"`
}
