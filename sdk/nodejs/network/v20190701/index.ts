// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export * from "./applicationGateway";
export * from "./applicationSecurityGroup";
export * from "./azureFirewall";
export * from "./bastionHost";
export * from "./connectionMonitor";
export * from "./ddosCustomPolicy";
export * from "./ddosProtectionPlan";
export * from "./expressRouteCircuit";
export * from "./expressRouteCircuitAuthorization";
export * from "./expressRouteCircuitConnection";
export * from "./expressRouteCircuitPeering";
export * from "./expressRouteConnection";
export * from "./expressRouteCrossConnectionPeering";
export * from "./expressRouteGateway";
export * from "./expressRoutePort";
export * from "./firewallPolicy";
export * from "./firewallPolicyRuleGroup";
export * from "./getApplicationGateway";
export * from "./getApplicationGatewayBackendHealthOnDemand";
export * from "./getApplicationSecurityGroup";
export * from "./getAzureFirewall";
export * from "./getBastionHost";
export * from "./getConnectionMonitor";
export * from "./getDdosCustomPolicy";
export * from "./getDdosProtectionPlan";
export * from "./getExpressRouteCircuit";
export * from "./getExpressRouteCircuitAuthorization";
export * from "./getExpressRouteCircuitConnection";
export * from "./getExpressRouteCircuitPeering";
export * from "./getExpressRouteConnection";
export * from "./getExpressRouteCrossConnectionPeering";
export * from "./getExpressRouteGateway";
export * from "./getExpressRoutePort";
export * from "./getFirewallPolicy";
export * from "./getFirewallPolicyRuleGroup";
export * from "./getInboundNatRule";
export * from "./getLoadBalancer";
export * from "./getLocalNetworkGateway";
export * from "./getNatGateway";
export * from "./getNetworkInterface";
export * from "./getNetworkInterfaceTapConfiguration";
export * from "./getNetworkProfile";
export * from "./getNetworkSecurityGroup";
export * from "./getNetworkWatcher";
export * from "./getP2sVpnGateway";
export * from "./getP2sVpnGatewayP2sVpnConnectionHealth";
export * from "./getP2sVpnServerConfiguration";
export * from "./getPacketCapture";
export * from "./getPrivateEndpoint";
export * from "./getPrivateLinkService";
export * from "./getPublicIPAddress";
export * from "./getPublicIPPrefix";
export * from "./getRoute";
export * from "./getRouteFilter";
export * from "./getRouteFilterRule";
export * from "./getRouteTable";
export * from "./getSecurityRule";
export * from "./getServiceEndpointPolicy";
export * from "./getServiceEndpointPolicyDefinition";
export * from "./getSubnet";
export * from "./getVirtualHub";
export * from "./getVirtualNetwork";
export * from "./getVirtualNetworkGateway";
export * from "./getVirtualNetworkGatewayAdvertisedRoutes";
export * from "./getVirtualNetworkGatewayBgpPeerStatus";
export * from "./getVirtualNetworkGatewayConnection";
export * from "./getVirtualNetworkGatewayLearnedRoutes";
export * from "./getVirtualNetworkGatewayVpnclientConnectionHealth";
export * from "./getVirtualNetworkGatewayVpnclientIpsecParameters";
export * from "./getVirtualNetworkPeering";
export * from "./getVirtualNetworkTap";
export * from "./getVirtualRouter";
export * from "./getVirtualRouterPeering";
export * from "./getVirtualWan";
export * from "./getVpnConnection";
export * from "./getVpnGateway";
export * from "./getVpnSite";
export * from "./getWebApplicationFirewallPolicy";
export * from "./inboundNatRule";
export * from "./loadBalancer";
export * from "./localNetworkGateway";
export * from "./natGateway";
export * from "./networkInterface";
export * from "./networkInterfaceTapConfiguration";
export * from "./networkProfile";
export * from "./networkSecurityGroup";
export * from "./networkWatcher";
export * from "./p2sVpnGateway";
export * from "./p2sVpnServerConfiguration";
export * from "./packetCapture";
export * from "./privateEndpoint";
export * from "./privateLinkService";
export * from "./publicIPAddress";
export * from "./publicIPPrefix";
export * from "./route";
export * from "./routeFilter";
export * from "./routeFilterRule";
export * from "./routeTable";
export * from "./securityRule";
export * from "./serviceEndpointPolicy";
export * from "./serviceEndpointPolicyDefinition";
export * from "./subnet";
export * from "./virtualHub";
export * from "./virtualNetwork";
export * from "./virtualNetworkGateway";
export * from "./virtualNetworkGatewayConnection";
export * from "./virtualNetworkPeering";
export * from "./virtualNetworkTap";
export * from "./virtualRouter";
export * from "./virtualRouterPeering";
export * from "./virtualWan";
export * from "./vpnConnection";
export * from "./vpnGateway";
export * from "./vpnSite";
export * from "./webApplicationFirewallPolicy";

// Export enums:
export * from "../../types/enums/network/v20190701";

// Import resources to register:
import { ApplicationGateway } from "./applicationGateway";
import { ApplicationSecurityGroup } from "./applicationSecurityGroup";
import { AzureFirewall } from "./azureFirewall";
import { BastionHost } from "./bastionHost";
import { ConnectionMonitor } from "./connectionMonitor";
import { DdosCustomPolicy } from "./ddosCustomPolicy";
import { DdosProtectionPlan } from "./ddosProtectionPlan";
import { ExpressRouteCircuit } from "./expressRouteCircuit";
import { ExpressRouteCircuitAuthorization } from "./expressRouteCircuitAuthorization";
import { ExpressRouteCircuitConnection } from "./expressRouteCircuitConnection";
import { ExpressRouteCircuitPeering } from "./expressRouteCircuitPeering";
import { ExpressRouteConnection } from "./expressRouteConnection";
import { ExpressRouteCrossConnectionPeering } from "./expressRouteCrossConnectionPeering";
import { ExpressRouteGateway } from "./expressRouteGateway";
import { ExpressRoutePort } from "./expressRoutePort";
import { FirewallPolicy } from "./firewallPolicy";
import { FirewallPolicyRuleGroup } from "./firewallPolicyRuleGroup";
import { InboundNatRule } from "./inboundNatRule";
import { LoadBalancer } from "./loadBalancer";
import { LocalNetworkGateway } from "./localNetworkGateway";
import { NatGateway } from "./natGateway";
import { NetworkInterface } from "./networkInterface";
import { NetworkInterfaceTapConfiguration } from "./networkInterfaceTapConfiguration";
import { NetworkProfile } from "./networkProfile";
import { NetworkSecurityGroup } from "./networkSecurityGroup";
import { NetworkWatcher } from "./networkWatcher";
import { P2sVpnGateway } from "./p2sVpnGateway";
import { P2sVpnServerConfiguration } from "./p2sVpnServerConfiguration";
import { PacketCapture } from "./packetCapture";
import { PrivateEndpoint } from "./privateEndpoint";
import { PrivateLinkService } from "./privateLinkService";
import { PublicIPAddress } from "./publicIPAddress";
import { PublicIPPrefix } from "./publicIPPrefix";
import { Route } from "./route";
import { RouteFilter } from "./routeFilter";
import { RouteFilterRule } from "./routeFilterRule";
import { RouteTable } from "./routeTable";
import { SecurityRule } from "./securityRule";
import { ServiceEndpointPolicy } from "./serviceEndpointPolicy";
import { ServiceEndpointPolicyDefinition } from "./serviceEndpointPolicyDefinition";
import { Subnet } from "./subnet";
import { VirtualHub } from "./virtualHub";
import { VirtualNetwork } from "./virtualNetwork";
import { VirtualNetworkGateway } from "./virtualNetworkGateway";
import { VirtualNetworkGatewayConnection } from "./virtualNetworkGatewayConnection";
import { VirtualNetworkPeering } from "./virtualNetworkPeering";
import { VirtualNetworkTap } from "./virtualNetworkTap";
import { VirtualRouter } from "./virtualRouter";
import { VirtualRouterPeering } from "./virtualRouterPeering";
import { VirtualWan } from "./virtualWan";
import { VpnConnection } from "./vpnConnection";
import { VpnGateway } from "./vpnGateway";
import { VpnSite } from "./vpnSite";
import { WebApplicationFirewallPolicy } from "./webApplicationFirewallPolicy";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "azure-native:network/v20190701:ApplicationGateway":
                return new ApplicationGateway(name, <any>undefined, { urn })
            case "azure-native:network/v20190701:ApplicationSecurityGroup":
                return new ApplicationSecurityGroup(name, <any>undefined, { urn })
            case "azure-native:network/v20190701:AzureFirewall":
                return new AzureFirewall(name, <any>undefined, { urn })
            case "azure-native:network/v20190701:BastionHost":
                return new BastionHost(name, <any>undefined, { urn })
            case "azure-native:network/v20190701:ConnectionMonitor":
                return new ConnectionMonitor(name, <any>undefined, { urn })
            case "azure-native:network/v20190701:DdosCustomPolicy":
                return new DdosCustomPolicy(name, <any>undefined, { urn })
            case "azure-native:network/v20190701:DdosProtectionPlan":
                return new DdosProtectionPlan(name, <any>undefined, { urn })
            case "azure-native:network/v20190701:ExpressRouteCircuit":
                return new ExpressRouteCircuit(name, <any>undefined, { urn })
            case "azure-native:network/v20190701:ExpressRouteCircuitAuthorization":
                return new ExpressRouteCircuitAuthorization(name, <any>undefined, { urn })
            case "azure-native:network/v20190701:ExpressRouteCircuitConnection":
                return new ExpressRouteCircuitConnection(name, <any>undefined, { urn })
            case "azure-native:network/v20190701:ExpressRouteCircuitPeering":
                return new ExpressRouteCircuitPeering(name, <any>undefined, { urn })
            case "azure-native:network/v20190701:ExpressRouteConnection":
                return new ExpressRouteConnection(name, <any>undefined, { urn })
            case "azure-native:network/v20190701:ExpressRouteCrossConnectionPeering":
                return new ExpressRouteCrossConnectionPeering(name, <any>undefined, { urn })
            case "azure-native:network/v20190701:ExpressRouteGateway":
                return new ExpressRouteGateway(name, <any>undefined, { urn })
            case "azure-native:network/v20190701:ExpressRoutePort":
                return new ExpressRoutePort(name, <any>undefined, { urn })
            case "azure-native:network/v20190701:FirewallPolicy":
                return new FirewallPolicy(name, <any>undefined, { urn })
            case "azure-native:network/v20190701:FirewallPolicyRuleGroup":
                return new FirewallPolicyRuleGroup(name, <any>undefined, { urn })
            case "azure-native:network/v20190701:InboundNatRule":
                return new InboundNatRule(name, <any>undefined, { urn })
            case "azure-native:network/v20190701:LoadBalancer":
                return new LoadBalancer(name, <any>undefined, { urn })
            case "azure-native:network/v20190701:LocalNetworkGateway":
                return new LocalNetworkGateway(name, <any>undefined, { urn })
            case "azure-native:network/v20190701:NatGateway":
                return new NatGateway(name, <any>undefined, { urn })
            case "azure-native:network/v20190701:NetworkInterface":
                return new NetworkInterface(name, <any>undefined, { urn })
            case "azure-native:network/v20190701:NetworkInterfaceTapConfiguration":
                return new NetworkInterfaceTapConfiguration(name, <any>undefined, { urn })
            case "azure-native:network/v20190701:NetworkProfile":
                return new NetworkProfile(name, <any>undefined, { urn })
            case "azure-native:network/v20190701:NetworkSecurityGroup":
                return new NetworkSecurityGroup(name, <any>undefined, { urn })
            case "azure-native:network/v20190701:NetworkWatcher":
                return new NetworkWatcher(name, <any>undefined, { urn })
            case "azure-native:network/v20190701:P2sVpnGateway":
                return new P2sVpnGateway(name, <any>undefined, { urn })
            case "azure-native:network/v20190701:P2sVpnServerConfiguration":
                return new P2sVpnServerConfiguration(name, <any>undefined, { urn })
            case "azure-native:network/v20190701:PacketCapture":
                return new PacketCapture(name, <any>undefined, { urn })
            case "azure-native:network/v20190701:PrivateEndpoint":
                return new PrivateEndpoint(name, <any>undefined, { urn })
            case "azure-native:network/v20190701:PrivateLinkService":
                return new PrivateLinkService(name, <any>undefined, { urn })
            case "azure-native:network/v20190701:PublicIPAddress":
                return new PublicIPAddress(name, <any>undefined, { urn })
            case "azure-native:network/v20190701:PublicIPPrefix":
                return new PublicIPPrefix(name, <any>undefined, { urn })
            case "azure-native:network/v20190701:Route":
                return new Route(name, <any>undefined, { urn })
            case "azure-native:network/v20190701:RouteFilter":
                return new RouteFilter(name, <any>undefined, { urn })
            case "azure-native:network/v20190701:RouteFilterRule":
                return new RouteFilterRule(name, <any>undefined, { urn })
            case "azure-native:network/v20190701:RouteTable":
                return new RouteTable(name, <any>undefined, { urn })
            case "azure-native:network/v20190701:SecurityRule":
                return new SecurityRule(name, <any>undefined, { urn })
            case "azure-native:network/v20190701:ServiceEndpointPolicy":
                return new ServiceEndpointPolicy(name, <any>undefined, { urn })
            case "azure-native:network/v20190701:ServiceEndpointPolicyDefinition":
                return new ServiceEndpointPolicyDefinition(name, <any>undefined, { urn })
            case "azure-native:network/v20190701:Subnet":
                return new Subnet(name, <any>undefined, { urn })
            case "azure-native:network/v20190701:VirtualHub":
                return new VirtualHub(name, <any>undefined, { urn })
            case "azure-native:network/v20190701:VirtualNetwork":
                return new VirtualNetwork(name, <any>undefined, { urn })
            case "azure-native:network/v20190701:VirtualNetworkGateway":
                return new VirtualNetworkGateway(name, <any>undefined, { urn })
            case "azure-native:network/v20190701:VirtualNetworkGatewayConnection":
                return new VirtualNetworkGatewayConnection(name, <any>undefined, { urn })
            case "azure-native:network/v20190701:VirtualNetworkPeering":
                return new VirtualNetworkPeering(name, <any>undefined, { urn })
            case "azure-native:network/v20190701:VirtualNetworkTap":
                return new VirtualNetworkTap(name, <any>undefined, { urn })
            case "azure-native:network/v20190701:VirtualRouter":
                return new VirtualRouter(name, <any>undefined, { urn })
            case "azure-native:network/v20190701:VirtualRouterPeering":
                return new VirtualRouterPeering(name, <any>undefined, { urn })
            case "azure-native:network/v20190701:VirtualWan":
                return new VirtualWan(name, <any>undefined, { urn })
            case "azure-native:network/v20190701:VpnConnection":
                return new VpnConnection(name, <any>undefined, { urn })
            case "azure-native:network/v20190701:VpnGateway":
                return new VpnGateway(name, <any>undefined, { urn })
            case "azure-native:network/v20190701:VpnSite":
                return new VpnSite(name, <any>undefined, { urn })
            case "azure-native:network/v20190701:WebApplicationFirewallPolicy":
                return new WebApplicationFirewallPolicy(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("azure-native", "network/v20190701", _module)
