// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-azure-nextgen/sdk/go/azure"
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
	case "azure-nextgen:network/latest:ApplicationGateway":
		r, err = NewApplicationGateway(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/latest:ApplicationGatewayPrivateEndpointConnection":
		r, err = NewApplicationGatewayPrivateEndpointConnection(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/latest:ApplicationSecurityGroup":
		r, err = NewApplicationSecurityGroup(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/latest:AzureFirewall":
		r, err = NewAzureFirewall(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/latest:BastionHost":
		r, err = NewBastionHost(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/latest:ConnectionMonitor":
		r, err = NewConnectionMonitor(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/latest:CustomIPPrefix":
		r, err = NewCustomIPPrefix(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/latest:DdosCustomPolicy":
		r, err = NewDdosCustomPolicy(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/latest:DdosProtectionPlan":
		r, err = NewDdosProtectionPlan(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/latest:DscpConfiguration":
		r, err = NewDscpConfiguration(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/latest:Endpoint":
		r, err = NewEndpoint(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/latest:Experiment":
		r, err = NewExperiment(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/latest:ExpressRouteCircuit":
		r, err = NewExpressRouteCircuit(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/latest:ExpressRouteCircuitAuthorization":
		r, err = NewExpressRouteCircuitAuthorization(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/latest:ExpressRouteCircuitConnection":
		r, err = NewExpressRouteCircuitConnection(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/latest:ExpressRouteCircuitPeering":
		r, err = NewExpressRouteCircuitPeering(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/latest:ExpressRouteConnection":
		r, err = NewExpressRouteConnection(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/latest:ExpressRouteCrossConnectionPeering":
		r, err = NewExpressRouteCrossConnectionPeering(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/latest:ExpressRouteGateway":
		r, err = NewExpressRouteGateway(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/latest:ExpressRoutePort":
		r, err = NewExpressRoutePort(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/latest:FirewallPolicy":
		r, err = NewFirewallPolicy(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/latest:FirewallPolicyRuleCollectionGroup":
		r, err = NewFirewallPolicyRuleCollectionGroup(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/latest:FirewallPolicyRuleGroup":
		r, err = NewFirewallPolicyRuleGroup(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/latest:FlowLog":
		r, err = NewFlowLog(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/latest:FrontDoor":
		r, err = NewFrontDoor(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/latest:HubRouteTable":
		r, err = NewHubRouteTable(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/latest:HubVirtualNetworkConnection":
		r, err = NewHubVirtualNetworkConnection(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/latest:InboundNatRule":
		r, err = NewInboundNatRule(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/latest:InterfaceEndpoint":
		r, err = NewInterfaceEndpoint(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/latest:IpAllocation":
		r, err = NewIpAllocation(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/latest:IpGroup":
		r, err = NewIpGroup(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/latest:LoadBalancer":
		r, err = NewLoadBalancer(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/latest:LoadBalancerBackendAddressPool":
		r, err = NewLoadBalancerBackendAddressPool(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/latest:LocalNetworkGateway":
		r, err = NewLocalNetworkGateway(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/latest:NatGateway":
		r, err = NewNatGateway(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/latest:NatRule":
		r, err = NewNatRule(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/latest:NetworkExperimentProfile":
		r, err = NewNetworkExperimentProfile(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/latest:NetworkInterface":
		r, err = NewNetworkInterface(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/latest:NetworkInterfaceTapConfiguration":
		r, err = NewNetworkInterfaceTapConfiguration(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/latest:NetworkProfile":
		r, err = NewNetworkProfile(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/latest:NetworkSecurityGroup":
		r, err = NewNetworkSecurityGroup(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/latest:NetworkVirtualAppliance":
		r, err = NewNetworkVirtualAppliance(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/latest:NetworkWatcher":
		r, err = NewNetworkWatcher(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/latest:P2sVpnGateway":
		r, err = NewP2sVpnGateway(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/latest:P2sVpnServerConfiguration":
		r, err = NewP2sVpnServerConfiguration(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/latest:PacketCapture":
		r, err = NewPacketCapture(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/latest:Policy":
		r, err = NewPolicy(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/latest:PrivateDnsZoneGroup":
		r, err = NewPrivateDnsZoneGroup(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/latest:PrivateEndpoint":
		r, err = NewPrivateEndpoint(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/latest:PrivateLinkService":
		r, err = NewPrivateLinkService(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/latest:PrivateLinkServicePrivateEndpointConnection":
		r, err = NewPrivateLinkServicePrivateEndpointConnection(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/latest:PrivateZone":
		r, err = NewPrivateZone(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/latest:Profile":
		r, err = NewProfile(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/latest:PublicIPAddress":
		r, err = NewPublicIPAddress(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/latest:PublicIPPrefix":
		r, err = NewPublicIPPrefix(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/latest:RecordSet":
		r, err = NewRecordSet(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/latest:Route":
		r, err = NewRoute(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/latest:RouteFilter":
		r, err = NewRouteFilter(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/latest:RouteFilterRule":
		r, err = NewRouteFilterRule(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/latest:RouteTable":
		r, err = NewRouteTable(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/latest:RulesEngine":
		r, err = NewRulesEngine(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/latest:SecurityPartnerProvider":
		r, err = NewSecurityPartnerProvider(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/latest:SecurityRule":
		r, err = NewSecurityRule(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/latest:ServiceEndpointPolicy":
		r, err = NewServiceEndpointPolicy(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/latest:ServiceEndpointPolicyDefinition":
		r, err = NewServiceEndpointPolicyDefinition(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/latest:Subnet":
		r, err = NewSubnet(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/latest:TrafficManagerUserMetricsKey":
		r, err = NewTrafficManagerUserMetricsKey(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/latest:VirtualApplianceSite":
		r, err = NewVirtualApplianceSite(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/latest:VirtualHub":
		r, err = NewVirtualHub(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/latest:VirtualHubBgpConnection":
		r, err = NewVirtualHubBgpConnection(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/latest:VirtualHubIpConfiguration":
		r, err = NewVirtualHubIpConfiguration(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/latest:VirtualHubRouteTableV2":
		r, err = NewVirtualHubRouteTableV2(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/latest:VirtualNetwork":
		r, err = NewVirtualNetwork(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/latest:VirtualNetworkGateway":
		r, err = NewVirtualNetworkGateway(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/latest:VirtualNetworkGatewayConnection":
		r, err = NewVirtualNetworkGatewayConnection(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/latest:VirtualNetworkLink":
		r, err = NewVirtualNetworkLink(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/latest:VirtualNetworkPeering":
		r, err = NewVirtualNetworkPeering(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/latest:VirtualNetworkTap":
		r, err = NewVirtualNetworkTap(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/latest:VirtualRouter":
		r, err = NewVirtualRouter(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/latest:VirtualRouterPeering":
		r, err = NewVirtualRouterPeering(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/latest:VirtualWan":
		r, err = NewVirtualWan(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/latest:VpnConnection":
		r, err = NewVpnConnection(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/latest:VpnGateway":
		r, err = NewVpnGateway(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/latest:VpnServerConfiguration":
		r, err = NewVpnServerConfiguration(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/latest:VpnSite":
		r, err = NewVpnSite(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/latest:WebApplicationFirewallPolicy":
		r, err = NewWebApplicationFirewallPolicy(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/latest:Zone":
		r, err = NewZone(ctx, name, nil, pulumi.URN_(urn))
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
		"azure-nextgen",
		"network/latest",
		&module{version},
	)
}
