// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200701

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-azure-native/sdk/go/azure"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "azure-native:network/v20200701:ApplicationGateway":
		r, err = NewApplicationGateway(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:network/v20200701:ApplicationGatewayPrivateEndpointConnection":
		r, err = NewApplicationGatewayPrivateEndpointConnection(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:network/v20200701:ApplicationSecurityGroup":
		r, err = NewApplicationSecurityGroup(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:network/v20200701:AzureFirewall":
		r, err = NewAzureFirewall(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:network/v20200701:BastionHost":
		r, err = NewBastionHost(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:network/v20200701:ConnectionMonitor":
		r, err = NewConnectionMonitor(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:network/v20200701:CustomIPPrefix":
		r, err = NewCustomIPPrefix(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:network/v20200701:DdosCustomPolicy":
		r, err = NewDdosCustomPolicy(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:network/v20200701:DdosProtectionPlan":
		r, err = NewDdosProtectionPlan(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:network/v20200701:DscpConfiguration":
		r, err = NewDscpConfiguration(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:network/v20200701:ExpressRouteCircuit":
		r, err = NewExpressRouteCircuit(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:network/v20200701:ExpressRouteCircuitAuthorization":
		r, err = NewExpressRouteCircuitAuthorization(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:network/v20200701:ExpressRouteCircuitConnection":
		r, err = NewExpressRouteCircuitConnection(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:network/v20200701:ExpressRouteCircuitPeering":
		r, err = NewExpressRouteCircuitPeering(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:network/v20200701:ExpressRouteConnection":
		r, err = NewExpressRouteConnection(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:network/v20200701:ExpressRouteCrossConnectionPeering":
		r, err = NewExpressRouteCrossConnectionPeering(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:network/v20200701:ExpressRouteGateway":
		r, err = NewExpressRouteGateway(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:network/v20200701:ExpressRoutePort":
		r, err = NewExpressRoutePort(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:network/v20200701:FirewallPolicy":
		r, err = NewFirewallPolicy(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:network/v20200701:FirewallPolicyRuleCollectionGroup":
		r, err = NewFirewallPolicyRuleCollectionGroup(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:network/v20200701:FlowLog":
		r, err = NewFlowLog(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:network/v20200701:HubRouteTable":
		r, err = NewHubRouteTable(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:network/v20200701:HubVirtualNetworkConnection":
		r, err = NewHubVirtualNetworkConnection(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:network/v20200701:InboundNatRule":
		r, err = NewInboundNatRule(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:network/v20200701:IpAllocation":
		r, err = NewIpAllocation(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:network/v20200701:IpGroup":
		r, err = NewIpGroup(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:network/v20200701:LoadBalancer":
		r, err = NewLoadBalancer(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:network/v20200701:LoadBalancerBackendAddressPool":
		r, err = NewLoadBalancerBackendAddressPool(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:network/v20200701:LocalNetworkGateway":
		r, err = NewLocalNetworkGateway(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:network/v20200701:NatGateway":
		r, err = NewNatGateway(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:network/v20200701:NetworkInterface":
		r, err = NewNetworkInterface(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:network/v20200701:NetworkInterfaceTapConfiguration":
		r, err = NewNetworkInterfaceTapConfiguration(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:network/v20200701:NetworkProfile":
		r, err = NewNetworkProfile(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:network/v20200701:NetworkSecurityGroup":
		r, err = NewNetworkSecurityGroup(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:network/v20200701:NetworkVirtualAppliance":
		r, err = NewNetworkVirtualAppliance(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:network/v20200701:NetworkWatcher":
		r, err = NewNetworkWatcher(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:network/v20200701:P2sVpnGateway":
		r, err = NewP2sVpnGateway(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:network/v20200701:PacketCapture":
		r, err = NewPacketCapture(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:network/v20200701:PrivateDnsZoneGroup":
		r, err = NewPrivateDnsZoneGroup(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:network/v20200701:PrivateEndpoint":
		r, err = NewPrivateEndpoint(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:network/v20200701:PrivateLinkService":
		r, err = NewPrivateLinkService(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:network/v20200701:PrivateLinkServicePrivateEndpointConnection":
		r, err = NewPrivateLinkServicePrivateEndpointConnection(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:network/v20200701:PublicIPAddress":
		r, err = NewPublicIPAddress(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:network/v20200701:PublicIPPrefix":
		r, err = NewPublicIPPrefix(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:network/v20200701:Route":
		r, err = NewRoute(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:network/v20200701:RouteFilter":
		r, err = NewRouteFilter(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:network/v20200701:RouteFilterRule":
		r, err = NewRouteFilterRule(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:network/v20200701:RouteTable":
		r, err = NewRouteTable(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:network/v20200701:SecurityPartnerProvider":
		r, err = NewSecurityPartnerProvider(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:network/v20200701:SecurityRule":
		r, err = NewSecurityRule(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:network/v20200701:ServiceEndpointPolicy":
		r, err = NewServiceEndpointPolicy(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:network/v20200701:ServiceEndpointPolicyDefinition":
		r, err = NewServiceEndpointPolicyDefinition(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:network/v20200701:Subnet":
		r, err = NewSubnet(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:network/v20200701:VirtualApplianceSite":
		r, err = NewVirtualApplianceSite(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:network/v20200701:VirtualHub":
		r, err = NewVirtualHub(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:network/v20200701:VirtualHubBgpConnection":
		r, err = NewVirtualHubBgpConnection(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:network/v20200701:VirtualHubIpConfiguration":
		r, err = NewVirtualHubIpConfiguration(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:network/v20200701:VirtualHubRouteTableV2":
		r, err = NewVirtualHubRouteTableV2(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:network/v20200701:VirtualNetwork":
		r, err = NewVirtualNetwork(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:network/v20200701:VirtualNetworkGateway":
		r, err = NewVirtualNetworkGateway(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:network/v20200701:VirtualNetworkGatewayConnection":
		r, err = NewVirtualNetworkGatewayConnection(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:network/v20200701:VirtualNetworkPeering":
		r, err = NewVirtualNetworkPeering(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:network/v20200701:VirtualNetworkTap":
		r, err = NewVirtualNetworkTap(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:network/v20200701:VirtualRouter":
		r, err = NewVirtualRouter(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:network/v20200701:VirtualRouterPeering":
		r, err = NewVirtualRouterPeering(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:network/v20200701:VirtualWan":
		r, err = NewVirtualWan(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:network/v20200701:VpnConnection":
		r, err = NewVpnConnection(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:network/v20200701:VpnGateway":
		r, err = NewVpnGateway(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:network/v20200701:VpnServerConfiguration":
		r, err = NewVpnServerConfiguration(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:network/v20200701:VpnSite":
		r, err = NewVpnSite(ctx, name, nil, pulumi.URN_(urn))
	case "azure-native:network/v20200701:WebApplicationFirewallPolicy":
		r, err = NewWebApplicationFirewallPolicy(ctx, name, nil, pulumi.URN_(urn))
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	return
}

func init() {
	version, err := azure.PkgVersion()
	if err != nil {
		fmt.Println("failed to determine package version. defaulting to v1: %v", err)
	}
	pulumi.RegisterResourceModule(
		"azure-native",
		"network/v20200701",
		&module{version},
	)
}
