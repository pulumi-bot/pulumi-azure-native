// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200801

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func GetP2sVpnGatewayP2sVpnConnectionHealth(ctx *pulumi.Context, args *GetP2sVpnGatewayP2sVpnConnectionHealthArgs, opts ...pulumi.InvokeOption) (*GetP2sVpnGatewayP2sVpnConnectionHealthResult, error) {
	var rv GetP2sVpnGatewayP2sVpnConnectionHealthResult
	err := ctx.Invoke("azure-nextgen:network/v20200801:getP2sVpnGatewayP2sVpnConnectionHealth", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetP2sVpnGatewayP2sVpnConnectionHealthArgs struct {
	// The name of the P2SVpnGateway.
	GatewayName string `pulumi:"gatewayName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// P2SVpnGateway Resource.
type GetP2sVpnGatewayP2sVpnConnectionHealthResult struct {
	// List of all customer specified DNS servers IP addresses.
	CustomDnsServers []string `pulumi:"customDnsServers"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag string `pulumi:"etag"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// Enable Routing Preference property for the Public IP Interface of the P2SVpnGateway.
	IsRoutingPreferenceInternet *bool `pulumi:"isRoutingPreferenceInternet"`
	// Resource location.
	Location string `pulumi:"location"`
	// Resource name.
	Name string `pulumi:"name"`
	// List of all p2s connection configurations of the gateway.
	P2SConnectionConfigurations []P2SConnectionConfigurationResponse `pulumi:"p2SConnectionConfigurations"`
	// The provisioning state of the P2S VPN gateway resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type string `pulumi:"type"`
	// The VirtualHub to which the gateway belongs.
	VirtualHub *SubResourceResponse `pulumi:"virtualHub"`
	// All P2S VPN clients' connection health status.
	VpnClientConnectionHealth VpnClientConnectionHealthResponse `pulumi:"vpnClientConnectionHealth"`
	// The scale unit for this p2s vpn gateway.
	VpnGatewayScaleUnit *int `pulumi:"vpnGatewayScaleUnit"`
	// The VpnServerConfiguration to which the p2sVpnGateway is attached to.
	VpnServerConfiguration *SubResourceResponse `pulumi:"vpnServerConfiguration"`
}
