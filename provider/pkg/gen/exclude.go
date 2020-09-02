package gen

import (
	"github.com/pulumi/pulumi/pkg/v2/codegen"
)

// excludeResources maintains a list of resources that we are skipping
// generating examples for at the moment due to codegen issues.
var excludeResources = codegen.NewStringSet(
	"azurerm:batch/v20170901:Certificate",
	"azurerm:batch/v20181201:Certificate",
	"azurerm:batch/v20190401:Certificate",
	"azurerm:batch/v20190801:Certificate",
	"azurerm:batch/v20200301:Certificate",
	"azurerm:batch/v20200501:Certificate",
	"azurerm:batch/latest:Certificate",
	"azurerm:batch/v20170901:Pool",
	"azurerm:batch/v20181201:Pool",
	"azurerm:batch/v20190401:Pool",
	"azurerm:batch/v20190801:Pool",
	"azurerm:batch/v20200301:Pool",
	"azurerm:batch/v20200501:Pool",
	"azurerm:batch/latest:Pool",
	"azurerm:costmanagement/v20180531:ReportConfig",
	"azurerm:costmanagement/latest:ReportConfig",
	"azurerm:costmanagement/v20190101:Export",
	"azurerm:costmanagement/v20191001:Export",
	"azurerm:costmanagement/v20180531:ReportConfigByResourceGroupName",
	"azurerm:costmanagement/latest:ReportConfigByResourceGroupName",
	"azurerm:costmanagement/v20190901:Export",
	"azurerm:costmanagement/v20191101:Export",
	"azurerm:costmanagement/v20200601:Export",
	"azurerm:costmanagement/latest:Export",
	"azurerm:costmanagement/v20191101:ViewByScope",
	"azurerm:costmanagement/v20200601:ViewByScope",
	"azurerm:costmanagement/latest:ViewByScope",
	"azurerm:costmanagement/v20191101:View",
	"azurerm:costmanagement/v20200601:View",
	"azurerm:costmanagement/latest:View",

	"azurerm:databox/v20200401:Job",
	"azurerm:databox/latest:Job",

	"azurerm:datamigration/v20180419:Task",
	"azurerm:datamigration/latest:Task",

	"azurerm:hybridcompute/v20191212:Machine",
	"azurerm:hybridcompute/v20200802:Machine",
	"azurerm:hybridcompute/latest:Machine",

	"azurerm:machinelearning/v20170101:WebService",
	"azurerm:machinelearning/latest:WebService",

	"azurerm:management/v20191001:DeploymentAtManagementGroupScope",
	"azurerm:management/v20200601:DeploymentAtManagementGroupScope",
	"azurerm:management/latest:DeploymentAtManagementGroupScope",
	"azurerm:management/v20191101:ManagementGroup",
	"azurerm:management/v20200201:ManagementGroup",
	"azurerm:management/v20200501:ManagementGroup",
	"azurerm:management/latest:ManagementGroup",

	"azurerm:network/v20161201:RouteFilter",
	"azurerm:network/v20170301:RouteFilter",
	"azurerm:network/v20170601:RouteFilter",
	"azurerm:network/v20170601:InboundNatRule",
	"azurerm:network/v20170601:LoadBalancer",
	"azurerm:network/v20170601:NetworkInterface",
	"azurerm:network/v20170601:NetworkSecurityGroup",
	"azurerm:network/v20170601:PublicIPAddress",
	"azurerm:network/v20170601:RouteTable",
	"azurerm:network/v20170601:Subnet",
	"azurerm:network/v20170601:VirtualNetwork",
	"azurerm:network/v20170801:InboundNatRule",
	"azurerm:network/v20170801:LoadBalancer",
	"azurerm:network/v20170801:NetworkInterface",
	"azurerm:network/v20170801:NetworkSecurityGroup",
	"azurerm:network/v20170801:PublicIPAddress",
	"azurerm:network/v20170801:RouteTable",
	"azurerm:network/v20170801:Subnet",
	"azurerm:network/v20170801:VirtualNetwork",
	"azurerm:network/v20170801:RouteFilter",
	"azurerm:network/v20170901:InboundNatRule",
	"azurerm:network/v20170901:LoadBalancer",
	"azurerm:network/v20170901:NetworkInterface",
	"azurerm:network/v20170901:NetworkSecurityGroup",
	"azurerm:network/v20170901:PublicIPAddress",
	"azurerm:network/v20170901:RouteTable",
	"azurerm:network/v20170901:Subnet",
	"azurerm:network/v20170901:VirtualNetwork",
	"azurerm:network/v20170901:RouteFilter",
	"azurerm:network/v20171001:InboundNatRule",
	"azurerm:network/v20171001:LoadBalancer",
	"azurerm:network/v20171001:NetworkInterface",
	"azurerm:network/v20171001:NetworkSecurityGroup",
	"azurerm:network/v20171001:PublicIPAddress",
	"azurerm:network/v20171001:RouteTable",
	"azurerm:network/v20171001:Subnet",
	"azurerm:network/v20171001:VirtualNetwork",
	"azurerm:network/v20171001:RouteFilter",
	"azurerm:network/v20171101:InboundNatRule",
	"azurerm:network/v20171101:LoadBalancer",
	"azurerm:network/v20171101:NetworkInterface",
	"azurerm:network/v20171101:NetworkSecurityGroup",
	"azurerm:network/v20171101:PublicIPAddress",
	"azurerm:network/v20171101:RouteTable",
	"azurerm:network/v20171101:Subnet",
	"azurerm:network/v20171101:VirtualNetwork",
	"azurerm:network/v20171101:RouteFilter",
	"azurerm:network/v20180101:ApplicationGateway",
	"azurerm:network/v20180101:InboundNatRule",
	"azurerm:network/v20180101:LoadBalancer",
	"azurerm:network/v20180101:NetworkInterface",
	"azurerm:network/v20180101:NetworkSecurityGroup",
	"azurerm:network/v20180101:PublicIPAddress",
	"azurerm:network/v20180101:RouteTable",
	"azurerm:network/v20180101:Subnet",
	"azurerm:network/v20180101:VirtualNetwork",
	"azurerm:network/v20180101:RouteFilter",
	"azurerm:network/v20180101:ApplicationGateway",
	"azurerm:network/v20180201:InboundNatRule",
	"azurerm:network/v20180201:LoadBalancer",
	"azurerm:network/v20180201:NetworkInterface",
	"azurerm:network/v20180201:NetworkSecurityGroup",
	"azurerm:network/v20180201:PublicIPAddress",
	"azurerm:network/v20180201:RouteTable",
	"azurerm:network/v20180201:Subnet",
	"azurerm:network/v20180201:VirtualNetwork",
	"azurerm:network/v20180201:RouteFilter",
	"azurerm:network/v20180201:ApplicationGateway",
	"azurerm:network/v20180201:ExpressRouteCrossConnectionPeering",
	"azurerm:network/v20180401:ApplicationGateway",
	"azurerm:network/v20180401:InboundNatRule",
	"azurerm:network/v20180401:LoadBalancer",
	"azurerm:network/v20180401:NetworkInterface",
	"azurerm:network/v20180401:NetworkSecurityGroup",
	"azurerm:network/v20180401:PublicIPAddress",
	"azurerm:network/v20180401:RouteTable",
	"azurerm:network/v20180401:Subnet",
	"azurerm:network/v20180401:VirtualNetwork",
	"azurerm:network/v20180401:RouteFilter",
	"azurerm:network/v20180401:ExpressRouteCircuit",
	"azurerm:network/v20180401:ExpressRouteCircuitPeering",
	"azurerm:network/v20180401:ExpressRouteCrossConnectionPeering",
	"azurerm:network/v20180601:ApplicationGateway",
	"azurerm:network/v20180601:InboundNatRule",
	"azurerm:network/v20180601:LoadBalancer",
	"azurerm:network/v20180601:NetworkInterface",
	"azurerm:network/v20180601:NetworkSecurityGroup",
	"azurerm:network/v20180601:PublicIPAddress",
	"azurerm:network/v20180601:RouteTable",
	"azurerm:network/v20180601:Subnet",
	"azurerm:network/v20180601:VirtualNetwork",
	"azurerm:network/v20180601:RouteFilter",
	"azurerm:network/v20180601:ExpressRouteCircuit",
	"azurerm:network/v20180601:ExpressRouteCircuitPeering",
	"azurerm:network/v20180601:ExpressRouteCrossConnectionPeering",
	"azurerm:network/v20180701:ApplicationGateway",
	"azurerm:network/v20180701:InboundNatRule",
	"azurerm:network/v20180701:LoadBalancer",
	"azurerm:network/v20180701:NetworkInterface",
	"azurerm:network/v20180701:NetworkSecurityGroup",
	"azurerm:network/v20180701:PublicIPAddress",
	"azurerm:network/v20180701:RouteTable",
	"azurerm:network/v20180701:Subnet",
	"azurerm:network/v20180701:VirtualNetwork",
	"azurerm:network/v20180701:RouteFilter",
	"azurerm:network/v20180701:ExpressRouteCircuit",
	"azurerm:network/v20180701:ExpressRouteCircuitPeering",
	"azurerm:network/v20180701:ExpressRouteCrossConnectionPeering",
	"azurerm:network/v20180801:ApplicationGateway",
	"azurerm:network/v20180801:InboundNatRule",
	"azurerm:network/v20180801:LoadBalancer",
	"azurerm:network/v20180801:NetworkInterface",
	"azurerm:network/v20180801:NetworkSecurityGroup",
	"azurerm:network/v20180801:PublicIPAddress",
	"azurerm:network/v20180801:RouteTable",
	"azurerm:network/v20180801:Subnet",
	"azurerm:network/v20180801:VirtualNetwork",
	"azurerm:network/v20180801:RouteFilter",
	"azurerm:network/v20180801:ExpressRouteCircuit",
	"azurerm:network/v20180801:ExpressRouteCircuitPeering",
	"azurerm:network/v20180801:ExpressRouteCrossConnectionPeering",
	"azurerm:network/v20180801:InterfaceEndpoint",
	"azurerm:network/v20180801:NetworkInterfaceTapConfiguration",
	"azurerm:network/v20180801:NetworkProfile",
	"azurerm:network/v20180801:ServiceEndpointPolicy",
	"azurerm:network/v20180801:VirtualNetworkTap",
	"azurerm:network/v20181001:ApplicationGateway",
	"azurerm:network/v20181001:InboundNatRule",
	"azurerm:network/v20181001:LoadBalancer",
	"azurerm:network/v20181001:NetworkInterface",
	"azurerm:network/v20181001:NetworkSecurityGroup",
	"azurerm:network/v20181001:PublicIPAddress",
	"azurerm:network/v20181001:RouteTable",
	"azurerm:network/v20181001:Subnet",
	"azurerm:network/v20181001:VirtualNetwork",
	"azurerm:network/v20181001:RouteFilter",
	"azurerm:network/v20181001:ExpressRouteCircuit",
	"azurerm:network/v20181001:ExpressRouteCircuitPeering",
	"azurerm:network/v20181001:ExpressRouteCrossConnectionPeering",
	"azurerm:network/v20181001:InterfaceEndpoint",
	"azurerm:network/v20181001:NetworkInterfaceTapConfiguration",
	"azurerm:network/v20181001:NetworkProfile",
	"azurerm:network/v20181001:ServiceEndpointPolicy",
	"azurerm:network/v20181001:VirtualNetworkTap",

	"azurerm:network/v20181101:ApplicationGateway",
	"azurerm:network/v20181101:InboundNatRule",
	"azurerm:network/v20181101:LoadBalancer",
	"azurerm:network/v20181101:NetworkInterface",
	"azurerm:network/v20181101:NetworkSecurityGroup",
	"azurerm:network/v20181101:PublicIPAddress",
	"azurerm:network/v20181101:RouteTable",
	"azurerm:network/v20181101:Subnet",
	"azurerm:network/v20181101:VirtualNetwork",
	"azurerm:network/v20181101:RouteFilter",
	"azurerm:network/v20181101:ExpressRouteCircuit",
	"azurerm:network/v20181101:ExpressRouteCircuitPeering",
	"azurerm:network/v20181101:ExpressRouteCrossConnectionPeering",
	"azurerm:network/v20181101:InterfaceEndpoint",
	"azurerm:network/v20181101:NetworkInterfaceTapConfiguration",
	"azurerm:network/v20181101:NetworkProfile",
	"azurerm:network/v20181101:ServiceEndpointPolicy",
	"azurerm:network/v20181101:VirtualNetworkTap",

	"azurerm:network/v20181201:ApplicationGateway",
	"azurerm:network/v20181201:InboundNatRule",
	"azurerm:network/v20181201:LoadBalancer",
	"azurerm:network/v20181201:NetworkInterface",
	"azurerm:network/v20181201:NetworkSecurityGroup",
	"azurerm:network/v20181201:PublicIPAddress",
	"azurerm:network/v20181201:RouteTable",
	"azurerm:network/v20181201:Subnet",
	"azurerm:network/v20181201:VirtualNetwork",
	"azurerm:network/v20181201:RouteFilter",
	"azurerm:network/v20181201:ExpressRouteCircuit",
	"azurerm:network/v20181201:ExpressRouteCircuitPeering",
	"azurerm:network/v20181201:ExpressRouteCrossConnectionPeering",
	"azurerm:network/v20181201:InterfaceEndpoint",
	"azurerm:network/v20181201:NetworkInterfaceTapConfiguration",
	"azurerm:network/v20181201:NetworkProfile",
	"azurerm:network/v20181201:ServiceEndpointPolicy",
	"azurerm:network/v20181201:VirtualNetworkTap",
	"azurerm:network/v20181201:WebApplicationFirewallPolicy",

	"azurerm:network/v20190201:ApplicationGateway",
	"azurerm:network/v20190201:InboundNatRule",
	"azurerm:network/v20190201:LoadBalancer",
	"azurerm:network/v20190201:NetworkInterface",
	"azurerm:network/v20190201:NetworkSecurityGroup",
	"azurerm:network/v20190201:PublicIPAddress",
	"azurerm:network/v20190201:RouteTable",
	"azurerm:network/v20190201:Subnet",
	"azurerm:network/v20190201:VirtualNetwork",
	"azurerm:network/v20190201:RouteFilter",
	"azurerm:network/v20190201:ExpressRouteCircuit",
	"azurerm:network/v20190201:ExpressRouteCircuitPeering",
	"azurerm:network/v20190201:ExpressRouteCrossConnectionPeering",
	"azurerm:network/v20190201:InterfaceEndpoint",
	"azurerm:network/v20190201:NetworkInterfaceTapConfiguration",
	"azurerm:network/v20190201:NetworkProfile",
	"azurerm:network/v20190201:ServiceEndpointPolicy",
	"azurerm:network/v20190201:VirtualNetworkTap",
	"azurerm:network/v20190201:WebApplicationFirewallPolicy",

	"azurerm:network/v20190401:ApplicationGateway",
	"azurerm:network/v20190401:InboundNatRule",
	"azurerm:network/v20190401:LoadBalancer",
	"azurerm:network/v20190401:NetworkInterface",
	"azurerm:network/v20190401:NetworkSecurityGroup",
	"azurerm:network/v20190401:PublicIPAddress",
	"azurerm:network/v20190401:PrivateEndpoint",
	"azurerm:network/v20190401:RouteTable",
	"azurerm:network/v20190401:Subnet",
	"azurerm:network/v20190401:VirtualNetwork",
	"azurerm:network/v20190401:RouteFilter",
	"azurerm:network/v20190401:ExpressRouteCircuit",
	"azurerm:network/v20190401:ExpressRouteCircuitPeering",
	"azurerm:network/v20190401:ExpressRouteCrossConnectionPeering",
	"azurerm:network/v20190401:InterfaceEndpoint",
	"azurerm:network/v20190401:NetworkInterfaceTapConfiguration",
	"azurerm:network/v20190401:NetworkProfile",
	"azurerm:network/v20190401:ServiceEndpointPolicy",
	"azurerm:network/v20190401:VirtualNetworkTap",
	"azurerm:network/v20190401:WebApplicationFirewallPolicy",
	"azurerm:network/v20190401:PrivateLinkService",

	"azurerm:network/v20190601:ApplicationGateway",
	"azurerm:network/v20190601:InboundNatRule",
	"azurerm:network/v20190601:LoadBalancer",
	"azurerm:network/v20190601:NetworkInterface",
	"azurerm:network/v20190601:NetworkSecurityGroup",
	"azurerm:network/v20190601:PublicIPAddress",
	"azurerm:network/v20190601:PrivateEndpoint",
	"azurerm:network/v20190601:RouteTable",
	"azurerm:network/v20190601:Subnet",
	"azurerm:network/v20190601:VirtualNetwork",
	"azurerm:network/v20190601:RouteFilter",
	"azurerm:network/v20190601:ExpressRouteCircuit",
	"azurerm:network/v20190601:ExpressRouteCircuitPeering",
	"azurerm:network/v20190601:ExpressRouteCrossConnectionPeering",
	"azurerm:network/v20190601:InterfaceEndpoint",
	"azurerm:network/v20190601:NetworkInterfaceTapConfiguration",
	"azurerm:network/v20190601:NetworkProfile",
	"azurerm:network/v20190601:ServiceEndpointPolicy",
	"azurerm:network/v20190601:VirtualNetworkTap",
	"azurerm:network/v20190601:WebApplicationFirewallPolicy",
	"azurerm:network/v20190601:PrivateLinkService",

	"azurerm:network/v20190701:ApplicationGateway",
	"azurerm:network/v20190701:InboundNatRule",
	"azurerm:network/v20190701:LoadBalancer",
	"azurerm:network/v20190701:NetworkInterface",
	"azurerm:network/v20190701:NetworkSecurityGroup",
	"azurerm:network/v20190701:PublicIPAddress",
	"azurerm:network/v20190701:PrivateEndpoint",
	"azurerm:network/v20190701:RouteTable",
	"azurerm:network/v20190701:Subnet",
	"azurerm:network/v20190701:VirtualNetwork",
	"azurerm:network/v20190701:RouteFilter",
	"azurerm:network/v20190701:ExpressRouteCircuit",
	"azurerm:network/v20190701:ExpressRouteCircuitPeering",
	"azurerm:network/v20190701:ExpressRouteCrossConnectionPeering",
	"azurerm:network/v20190701:InterfaceEndpoint",
	"azurerm:network/v20190701:NetworkInterfaceTapConfiguration",
	"azurerm:network/v20190701:NetworkProfile",
	"azurerm:network/v20190701:ServiceEndpointPolicy",
	"azurerm:network/v20190701:VirtualNetworkTap",
	"azurerm:network/v20190701:WebApplicationFirewallPolicy",
	"azurerm:network/v20190701:PrivateLinkService",

	"azurerm:network/v20190801:ApplicationGateway",
	"azurerm:network/v20190801:InboundNatRule",
	"azurerm:network/v20190801:LoadBalancer",
	"azurerm:network/v20190801:NetworkInterface",
	"azurerm:network/v20190801:NetworkSecurityGroup",
	"azurerm:network/v20190801:PublicIPAddress",
	"azurerm:network/v20190801:PrivateEndpoint",
	"azurerm:network/v20190801:RouteTable",
	"azurerm:network/v20190801:Subnet",
	"azurerm:network/v20190801:VirtualNetwork",
	"azurerm:network/v20190801:RouteFilter",
	"azurerm:network/v20190801:ExpressRouteCircuit",
	"azurerm:network/v20190801:ExpressRouteCircuitPeering",
	"azurerm:network/v20190801:ExpressRouteCrossConnectionPeering",
	"azurerm:network/v20190801:InterfaceEndpoint",
	"azurerm:network/v20190801:NetworkInterfaceTapConfiguration",
	"azurerm:network/v20190801:NetworkProfile",
	"azurerm:network/v20190801:ServiceEndpointPolicy",
	"azurerm:network/v20190801:VirtualNetworkTap",
	"azurerm:network/v20190801:WebApplicationFirewallPolicy",
	"azurerm:network/v20190801:PrivateLinkService",

	"azurerm:network/v20190901:ApplicationGateway",
	"azurerm:network/v20190901:InboundNatRule",
	"azurerm:network/v20190901:LoadBalancer",
	"azurerm:network/v20190901:NetworkInterface",
	"azurerm:network/v20190901:NetworkSecurityGroup",
	"azurerm:network/v20190901:PublicIPAddress",
	"azurerm:network/v20190901:PrivateEndpoint",
	"azurerm:network/v20190901:RouteTable",
	"azurerm:network/v20190901:Subnet",
	"azurerm:network/v20190901:VirtualNetwork",
	"azurerm:network/v20190901:RouteFilter",
	"azurerm:network/v20190901:ExpressRouteCircuit",
	"azurerm:network/v20190901:ExpressRouteCircuitPeering",
	"azurerm:network/v20190901:ExpressRouteCrossConnectionPeering",
	"azurerm:network/v20190901:InterfaceEndpoint",
	"azurerm:network/v20190901:NetworkInterfaceTapConfiguration",
	"azurerm:network/v20190901:NetworkProfile",
	"azurerm:network/v20190901:ServiceEndpointPolicy",
	"azurerm:network/v20190901:VirtualNetworkTap",
	"azurerm:network/v20190901:WebApplicationFirewallPolicy",
	"azurerm:network/v20190901:PrivateLinkService",
	"azurerm:network/v20190901:PrivateLinkServicePrivateEndpointConnection",

	"azurerm:network/v20191101:ApplicationGateway",
	"azurerm:network/v20191101:InboundNatRule",
	"azurerm:network/v20191101:LoadBalancer",
	"azurerm:network/v20191101:NetworkInterface",
	"azurerm:network/v20191101:NetworkSecurityGroup",
	"azurerm:network/v20191101:PublicIPAddress",
	"azurerm:network/v20191101:PrivateEndpoint",
	"azurerm:network/v20191101:RouteTable",
	"azurerm:network/v20191101:Subnet",
	"azurerm:network/v20191101:VirtualNetwork",
	"azurerm:network/v20191101:RouteFilter",
	"azurerm:network/v20191101:ExpressRouteCircuit",
	"azurerm:network/v20191101:ExpressRouteCircuitPeering",
	"azurerm:network/v20191101:ExpressRouteCrossConnectionPeering",
	"azurerm:network/v20191101:InterfaceEndpoint",
	"azurerm:network/v20191101:NetworkInterfaceTapConfiguration",
	"azurerm:network/v20191101:NetworkProfile",
	"azurerm:network/v20191101:ServiceEndpointPolicy",
	"azurerm:network/v20191101:VirtualNetworkTap",
	"azurerm:network/v20191101:WebApplicationFirewallPolicy",
	"azurerm:network/v20191101:PrivateLinkService",
	"azurerm:network/v20191101:PrivateLinkServicePrivateEndpointConnection",

	"azurerm:network/v20191201:ApplicationGateway",
	"azurerm:network/v20191201:InboundNatRule",
	"azurerm:network/v20191201:LoadBalancer",
	"azurerm:network/v20191201:NetworkInterface",
	"azurerm:network/v20191201:NetworkSecurityGroup",
	"azurerm:network/v20191201:PublicIPAddress",
	"azurerm:network/v20191201:PrivateEndpoint",
	"azurerm:network/v20191201:RouteTable",
	"azurerm:network/v20191201:Subnet",
	"azurerm:network/v20191201:VirtualNetwork",
	"azurerm:network/v20191201:RouteFilter",
	"azurerm:network/v20191201:ExpressRouteCircuit",
	"azurerm:network/v20191201:ExpressRouteCircuitPeering",
	"azurerm:network/v20191201:ExpressRouteCrossConnectionPeering",
	"azurerm:network/v20191201:InterfaceEndpoint",
	"azurerm:network/v20191201:NetworkInterfaceTapConfiguration",
	"azurerm:network/v20191201:NetworkProfile",
	"azurerm:network/v20191201:ServiceEndpointPolicy",
	"azurerm:network/v20191201:VirtualNetworkTap",
	"azurerm:network/v20191201:WebApplicationFirewallPolicy",
	"azurerm:network/v20191201:PrivateLinkService",
	"azurerm:network/v20191201:PrivateLinkServicePrivateEndpointConnection",

	"azurerm:network/v20200301:ApplicationGateway",
	"azurerm:network/v20200301:InboundNatRule",
	"azurerm:network/v20200301:LoadBalancer",
	"azurerm:network/v20200301:NetworkInterface",
	"azurerm:network/v20200301:NetworkSecurityGroup",
	"azurerm:network/v20200301:PublicIPAddress",
	"azurerm:network/v20200301:PrivateEndpoint",
	"azurerm:network/v20200301:RouteTable",
	"azurerm:network/v20200301:Subnet",
	"azurerm:network/v20200301:VirtualNetwork",
	"azurerm:network/v20200301:RouteFilter",
	"azurerm:network/v20200301:ExpressRouteCircuit",
	"azurerm:network/v20200301:ExpressRouteCircuitPeering",
	"azurerm:network/v20200301:ExpressRouteCrossConnectionPeering",
	"azurerm:network/v20200301:InterfaceEndpoint",
	"azurerm:network/v20200301:NetworkInterfaceTapConfiguration",
	"azurerm:network/v20200301:NetworkProfile",
	"azurerm:network/v20200301:ServiceEndpointPolicy",
	"azurerm:network/v20200301:VirtualNetworkTap",
	"azurerm:network/v20200301:WebApplicationFirewallPolicy",
	"azurerm:network/v20200301:PrivateLinkService",
	"azurerm:network/v20200301:PrivateLinkServicePrivateEndpointConnection",

	"azurerm:network/v20200401:ApplicationGateway",
	"azurerm:network/v20200401:InboundNatRule",
	"azurerm:network/v20200401:LoadBalancer",
	"azurerm:network/v20200401:NetworkInterface",
	"azurerm:network/v20200401:NetworkSecurityGroup",
	"azurerm:network/v20200401:PublicIPAddress",
	"azurerm:network/v20200401:PrivateEndpoint",
	"azurerm:network/v20200401:RouteTable",
	"azurerm:network/v20200401:Subnet",
	"azurerm:network/v20200401:VirtualNetwork",
	"azurerm:network/v20200401:RouteFilter",
	"azurerm:network/v20200401:ExpressRouteCircuit",
	"azurerm:network/v20200401:ExpressRouteCircuitPeering",
	"azurerm:network/v20200401:ExpressRouteCrossConnectionPeering",
	"azurerm:network/v20200401:InterfaceEndpoint",
	"azurerm:network/v20200401:NetworkInterfaceTapConfiguration",
	"azurerm:network/v20200401:NetworkProfile",
	"azurerm:network/v20200401:ServiceEndpointPolicy",
	"azurerm:network/v20200401:VirtualNetworkTap",
	"azurerm:network/v20200401:WebApplicationFirewallPolicy",
	"azurerm:network/v20200401:PrivateLinkService",
	"azurerm:network/v20200401:PrivateLinkServicePrivateEndpointConnection",
	"azurerm:network/v20200401:LoadBalancerBackendAddressPool",

	"azurerm:network/v20200501:ApplicationGateway",
	"azurerm:network/v20200501:InboundNatRule",
	"azurerm:network/v20200501:LoadBalancer",
	"azurerm:network/v20200501:NetworkInterface",
	"azurerm:network/v20200501:NetworkSecurityGroup",
	"azurerm:network/v20200501:PublicIPAddress",
	"azurerm:network/v20200501:PrivateEndpoint",
	"azurerm:network/v20200501:RouteTable",
	"azurerm:network/v20200501:Subnet",
	"azurerm:network/v20200501:VirtualNetwork",
	"azurerm:network/v20200501:RouteFilter",
	"azurerm:network/v20200501:ExpressRouteCircuit",
	"azurerm:network/v20200501:ExpressRouteCircuitPeering",
	"azurerm:network/v20200501:ExpressRouteCrossConnectionPeering",
	"azurerm:network/v20200501:InterfaceEndpoint",
	"azurerm:network/v20200501:NetworkInterfaceTapConfiguration",
	"azurerm:network/v20200501:NetworkProfile",
	"azurerm:network/v20200501:ServiceEndpointPolicy",
	"azurerm:network/v20200501:VirtualNetworkTap",
	"azurerm:network/v20200501:WebApplicationFirewallPolicy",
	"azurerm:network/v20200501:PrivateLinkService",
	"azurerm:network/v20200501:PrivateLinkServicePrivateEndpointConnection",
	"azurerm:network/v20200501:LoadBalancerBackendAddressPool",
	"azurerm:network/v20200501:VirtualHubIpConfiguration",
	"azurerm:network/v20200501:VpnSite",
	"azurerm:network/v20200501:ApplicationGatewayPrivateEndpointConnection",

	"azurerm:network/v20200601:ApplicationGateway",
	"azurerm:network/v20200601:InboundNatRule",
	"azurerm:network/v20200601:LoadBalancer",
	"azurerm:network/v20200601:NetworkInterface",
	"azurerm:network/v20200601:NetworkSecurityGroup",
	"azurerm:network/v20200601:PublicIPAddress",
	"azurerm:network/v20200601:PrivateEndpoint",
	"azurerm:network/v20200601:RouteTable",
	"azurerm:network/v20200601:Subnet",
	"azurerm:network/v20200601:VirtualNetwork",
	"azurerm:network/v20200601:RouteFilter",
	"azurerm:network/v20200601:ExpressRouteCircuit",
	"azurerm:network/v20200601:ExpressRouteCircuitPeering",
	"azurerm:network/v20200601:ExpressRouteCrossConnectionPeering",
	"azurerm:network/v20200601:InterfaceEndpoint",
	"azurerm:network/v20200601:NetworkInterfaceTapConfiguration",
	"azurerm:network/v20200601:NetworkProfile",
	"azurerm:network/v20200601:ServiceEndpointPolicy",
	"azurerm:network/v20200601:VirtualNetworkTap",
	"azurerm:network/v20200601:WebApplicationFirewallPolicy",
	"azurerm:network/v20200601:PrivateLinkService",
	"azurerm:network/v20200601:PrivateLinkServicePrivateEndpointConnection",
	"azurerm:network/v20200601:LoadBalancerBackendAddressPool",
	"azurerm:network/v20200601:VirtualHubIpConfiguration",
	"azurerm:network/v20200601:VpnSite",
	"azurerm:network/v20200601:ApplicationGatewayPrivateEndpointConnection",
	"azurerm:network/v20200601:DscpConfiguration",

	"azurerm:network/latest:ApplicationGateway",
	"azurerm:network/latest:InboundNatRule",
	"azurerm:network/latest:LoadBalancer",
	"azurerm:network/latest:NetworkInterface",
	"azurerm:network/latest:NetworkSecurityGroup",
	"azurerm:network/latest:PublicIPAddress",
	"azurerm:network/latest:PrivateEndpoint",
	"azurerm:network/latest:RouteTable",
	"azurerm:network/latest:Subnet",
	"azurerm:network/latest:VirtualNetwork",
	"azurerm:network/latest:RouteFilter",
	"azurerm:network/latest:ExpressRouteCircuit",
	"azurerm:network/latest:ExpressRouteCircuitPeering",
	"azurerm:network/latest:ExpressRouteCrossConnectionPeering",
	"azurerm:network/latest:InterfaceEndpoint",
	"azurerm:network/latest:NetworkInterfaceTapConfiguration",
	"azurerm:network/latest:NetworkProfile",
	"azurerm:network/latest:ServiceEndpointPolicy",
	"azurerm:network/latest:VirtualNetworkTap",
	"azurerm:network/latest:WebApplicationFirewallPolicy",
	"azurerm:network/latest:PrivateLinkService",
	"azurerm:network/latest:PrivateLinkServicePrivateEndpointConnection",
	"azurerm:network/latest:LoadBalancerBackendAddressPool",
	"azurerm:network/latest:VirtualHubIpConfiguration",
	"azurerm:network/latest:VpnSite",
	"azurerm:network/latest:ApplicationGatewayPrivateEndpointConnection",
	"azurerm:network/latest:DscpConfiguration",

	"azurerm:powerbidedicated/v20171001:CapacityDetails",
	"azurerm:powerbidedicated/latest:CapacityDetails",
	"azurerm:recoveryservices/v20160810:ReplicationFabric",
	"azurerm:recoveryservices/latest:ReplicationFabric",
	"azurerm:recoveryservices/v20160810:ReplicationProtectedItem",
	"azurerm:recoveryservices/latest:ReplicationProtectedItem",
	"azurerm:recoveryservices/v20160810:ReplicationProtectionContainerMapping",
	"azurerm:recoveryservices/latest:ReplicationProtectionContainerMapping",
	"azurerm:resources/v20191001:Deployment",
	"azurerm:resources/v20191001:DeploymentAtScope",
	"azurerm:resources/v20191001:DeploymentAtTenantScope",
	"azurerm:resources/v20200601:Deployment",
	"azurerm:resources/v20200601:DeploymentAtScope",
	"azurerm:resources/v20200601:DeploymentAtTenantScope",
	"azurerm:resources/v20200601:DeploymentAtSubscriptionScope",
	"azurerm:resources/latest:Deployment",
	"azurerm:resources/latest:DeploymentAtScope",
	"azurerm:resources/latest:DeploymentAtTenantScope",
	"azurerm:resources/latest:DeploymentAtSubscriptionScope",

	// recursive reference in schema

	"azurerm:botservice/v20200602:BotConnection", // malformed body

	"azurerm:hybridcompute/v20181120:GuestConfigurationHCRPAssignment",
	"azurerm:hybridcompute/v20200625:GuestConfigurationHCRPAssignment", // python name mismatch
	"azurerm:hybridcompute/latest:GuestConfigurationHCRPAssignment", // python name mismatch
)
