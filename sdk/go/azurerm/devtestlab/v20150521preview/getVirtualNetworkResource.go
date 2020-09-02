// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20150521preview

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupVirtualNetworkResource(ctx *pulumi.Context, args *LookupVirtualNetworkResourceArgs, opts ...pulumi.InvokeOption) (*LookupVirtualNetworkResourceResult, error) {
	var rv LookupVirtualNetworkResourceResult
	err := ctx.Invoke("azurerm:devtestlab/v20150521preview:getVirtualNetworkResource", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVirtualNetworkResourceArgs struct {
	// The name of the lab.
	LabName string `pulumi:"labName"`
	// The name of the virtual network.
	Name string `pulumi:"name"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// A virtual network.
type LookupVirtualNetworkResourceResult struct {
	// The allowed subnets of the virtual network.
	AllowedSubnets []SubnetResponse `pulumi:"allowedSubnets"`
	// The description of the virtual network.
	Description *string `pulumi:"description"`
	// The Microsoft.Network resource identifier of the virtual network.
	ExternalProviderResourceId *string `pulumi:"externalProviderResourceId"`
	// The location of the resource.
	Location *string `pulumi:"location"`
	// The name of the resource.
	Name *string `pulumi:"name"`
	// The provisioning status of the resource.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The subnet overrides of the virtual network.
	SubnetOverrides []SubnetOverrideResponse `pulumi:"subnetOverrides"`
	// The tags of the resource.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource.
	Type *string `pulumi:"type"`
}
