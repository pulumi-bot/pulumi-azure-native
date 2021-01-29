// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180101

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
	case "azure-nextgen:network/v20180101:ApplicationGateway":
		r, err = NewApplicationGateway(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/v20180101:ApplicationSecurityGroup":
		r, err = NewApplicationSecurityGroup(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/v20180101:ConnectionMonitor":
		r, err = NewConnectionMonitor(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/v20180101:ExpressRouteCircuit":
		r, err = NewExpressRouteCircuit(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/v20180101:ExpressRouteCircuitAuthorization":
		r, err = NewExpressRouteCircuitAuthorization(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/v20180101:ExpressRouteCircuitPeering":
		r, err = NewExpressRouteCircuitPeering(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/v20180101:InboundNatRule":
		r, err = NewInboundNatRule(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/v20180101:LoadBalancer":
		r, err = NewLoadBalancer(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/v20180101:LocalNetworkGateway":
		r, err = NewLocalNetworkGateway(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/v20180101:NetworkInterface":
		r, err = NewNetworkInterface(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/v20180101:NetworkSecurityGroup":
		r, err = NewNetworkSecurityGroup(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/v20180101:NetworkWatcher":
		r, err = NewNetworkWatcher(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/v20180101:PacketCapture":
		r, err = NewPacketCapture(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/v20180101:PublicIPAddress":
		r, err = NewPublicIPAddress(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/v20180101:Route":
		r, err = NewRoute(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/v20180101:RouteFilter":
		r, err = NewRouteFilter(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/v20180101:RouteFilterRule":
		r, err = NewRouteFilterRule(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/v20180101:RouteTable":
		r, err = NewRouteTable(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/v20180101:SecurityRule":
		r, err = NewSecurityRule(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/v20180101:Subnet":
		r, err = NewSubnet(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/v20180101:VirtualNetwork":
		r, err = NewVirtualNetwork(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/v20180101:VirtualNetworkGateway":
		r, err = NewVirtualNetworkGateway(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/v20180101:VirtualNetworkGatewayConnection":
		r, err = NewVirtualNetworkGatewayConnection(ctx, name, nil, pulumi.URN_(urn))
	case "azure-nextgen:network/v20180101:VirtualNetworkPeering":
		r, err = NewVirtualNetworkPeering(ctx, name, nil, pulumi.URN_(urn))
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
		"network/v20180101",
		&module{version},
	)
}
