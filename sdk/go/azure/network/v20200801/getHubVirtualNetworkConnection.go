// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200801

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupHubVirtualNetworkConnection(ctx *pulumi.Context, args *LookupHubVirtualNetworkConnectionArgs, opts ...pulumi.InvokeOption) (*LookupHubVirtualNetworkConnectionResult, error) {
	var rv LookupHubVirtualNetworkConnectionResult
	err := ctx.Invoke("azure-nextgen:network/v20200801:getHubVirtualNetworkConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupHubVirtualNetworkConnectionArgs struct {
	// The name of the vpn connection.
	ConnectionName string `pulumi:"connectionName"`
	// The resource group name of the VirtualHub.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the VirtualHub.
	VirtualHubName string `pulumi:"virtualHubName"`
}

// HubVirtualNetworkConnection Resource.
type LookupHubVirtualNetworkConnectionResult struct {
	// Deprecated: VirtualHub to RemoteVnet transit to enabled or not.
	AllowHubToRemoteVnetTransit *bool `pulumi:"allowHubToRemoteVnetTransit"`
	// Deprecated: Allow RemoteVnet to use Virtual Hub's gateways.
	AllowRemoteVnetToUseHubVnetGateways *bool `pulumi:"allowRemoteVnetToUseHubVnetGateways"`
	// Enable internet security.
	EnableInternetSecurity *bool `pulumi:"enableInternetSecurity"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `pulumi:"name"`
	// The provisioning state of the hub virtual network connection resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// Reference to the remote virtual network.
	RemoteVirtualNetwork *SubResourceResponse `pulumi:"remoteVirtualNetwork"`
	// The Routing Configuration indicating the associated and propagated route tables on this connection.
	RoutingConfiguration *RoutingConfigurationResponse `pulumi:"routingConfiguration"`
}
