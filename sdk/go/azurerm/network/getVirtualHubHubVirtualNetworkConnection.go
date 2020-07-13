// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupVirtualHubHubVirtualNetworkConnection(ctx *pulumi.Context, args *LookupVirtualHubHubVirtualNetworkConnectionArgs, opts ...pulumi.InvokeOption) (*LookupVirtualHubHubVirtualNetworkConnectionResult, error) {
	var rv LookupVirtualHubHubVirtualNetworkConnectionResult
	err := ctx.Invoke("azurerm:network:getVirtualHubHubVirtualNetworkConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVirtualHubHubVirtualNetworkConnectionArgs struct {
	// The name of the vpn connection.
	Name string `pulumi:"name"`
	// The resource group name of the VirtualHub.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the VirtualHub.
	VirtualHubName string `pulumi:"virtualHubName"`
}

// HubVirtualNetworkConnection Resource.
type LookupVirtualHubHubVirtualNetworkConnectionResult struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `pulumi:"name"`
	// Properties of the hub virtual network connection.
	Properties HubVirtualNetworkConnectionPropertiesResponse `pulumi:"properties"`
}