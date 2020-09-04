// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * A common class for general resource information.
 *
 * ## Example Usage
 * ### UpdateVirtualNetworkGateway
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const virtualNetworkGateway = new azurerm.network.v20200601.VirtualNetworkGateway("virtualNetworkGateway", {
 *     activeActive: false,
 *     bgpSettings: {
 *         asn: 65515,
 *         bgpPeeringAddress: "10.0.1.30",
 *         peerWeight: 0,
 *     },
 *     customRoutes: {
 *         addressPrefixes: ["101.168.0.6/32"],
 *     },
 *     enableBgp: false,
 *     enableDnsForwarding: true,
 *     gatewayType: "Vpn",
 *     ipConfigurations: [{
 *         name: "gwipconfig1",
 *     }],
 *     location: "centralus",
 *     resourceGroupName: "rg1",
 *     sku: {
 *         name: "VpnGw1",
 *         tier: "VpnGw1",
 *     },
 *     virtualNetworkGatewayName: "vpngw",
 *     vpnClientConfiguration: {
 *         radiusServers: [{
 *             radiusServerAddress: "10.2.0.0",
 *             radiusServerScore: 20,
 *             radiusServerSecret: "radiusServerSecret",
 *         }],
 *         vpnClientProtocols: ["OpenVPN"],
 *         vpnClientRevokedCertificates: [],
 *         vpnClientRootCertificates: [],
 *     },
 *     vpnType: "RouteBased",
 * });
 *
 * ```
 */
export class VirtualNetworkGateway extends pulumi.CustomResource {
    /**
     * Get an existing VirtualNetworkGateway resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): VirtualNetworkGateway {
        return new VirtualNetworkGateway(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:network/v20200601:VirtualNetworkGateway';

    /**
     * Returns true if the given object is an instance of VirtualNetworkGateway.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VirtualNetworkGateway {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VirtualNetworkGateway.__pulumiType;
    }

    /**
     * ActiveActive flag.
     */
    public readonly activeActive!: pulumi.Output<boolean | undefined>;
    /**
     * Virtual network gateway's BGP speaker settings.
     */
    public readonly bgpSettings!: pulumi.Output<outputs.network.v20200601.BgpSettingsResponse | undefined>;
    /**
     * The reference to the address space resource which represents the custom routes address space specified by the customer for virtual network gateway and VpnClient.
     */
    public readonly customRoutes!: pulumi.Output<outputs.network.v20200601.AddressSpaceResponse | undefined>;
    /**
     * Whether BGP is enabled for this virtual network gateway or not.
     */
    public readonly enableBgp!: pulumi.Output<boolean | undefined>;
    /**
     * Whether dns forwarding is enabled or not.
     */
    public readonly enableDnsForwarding!: pulumi.Output<boolean | undefined>;
    /**
     * Whether private IP needs to be enabled on this gateway for connections or not.
     */
    public readonly enablePrivateIpAddress!: pulumi.Output<boolean | undefined>;
    /**
     * A unique read-only string that changes whenever the resource is updated.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * The reference to the LocalNetworkGateway resource which represents local network site having default routes. Assign Null value in case of removing existing default site setting.
     */
    public readonly gatewayDefaultSite!: pulumi.Output<outputs.network.v20200601.SubResourceResponse | undefined>;
    /**
     * The type of this virtual network gateway.
     */
    public readonly gatewayType!: pulumi.Output<string | undefined>;
    /**
     * The IP address allocated by the gateway to which dns requests can be sent.
     */
    public /*out*/ readonly inboundDnsForwardingEndpoint!: pulumi.Output<string>;
    /**
     * IP configurations for virtual network gateway.
     */
    public readonly ipConfigurations!: pulumi.Output<outputs.network.v20200601.VirtualNetworkGatewayIPConfigurationResponse[] | undefined>;
    /**
     * Resource location.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The provisioning state of the virtual network gateway resource.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * The resource GUID property of the virtual network gateway resource.
     */
    public /*out*/ readonly resourceGuid!: pulumi.Output<string>;
    /**
     * The reference to the VirtualNetworkGatewaySku resource which represents the SKU selected for Virtual network gateway.
     */
    public readonly sku!: pulumi.Output<outputs.network.v20200601.VirtualNetworkGatewaySkuResponse | undefined>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * The reference to the VpnClientConfiguration resource which represents the P2S VpnClient configurations.
     */
    public readonly vpnClientConfiguration!: pulumi.Output<outputs.network.v20200601.VpnClientConfigurationResponse | undefined>;
    /**
     * The generation for this VirtualNetworkGateway. Must be None if gatewayType is not VPN.
     */
    public readonly vpnGatewayGeneration!: pulumi.Output<string | undefined>;
    /**
     * The type of this virtual network gateway.
     */
    public readonly vpnType!: pulumi.Output<string | undefined>;

    /**
     * Create a VirtualNetworkGateway resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VirtualNetworkGatewayArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.virtualNetworkGatewayName === undefined) {
                throw new Error("Missing required property 'virtualNetworkGatewayName'");
            }
            inputs["activeActive"] = args ? args.activeActive : undefined;
            inputs["bgpSettings"] = args ? args.bgpSettings : undefined;
            inputs["customRoutes"] = args ? args.customRoutes : undefined;
            inputs["enableBgp"] = args ? args.enableBgp : undefined;
            inputs["enableDnsForwarding"] = args ? args.enableDnsForwarding : undefined;
            inputs["enablePrivateIpAddress"] = args ? args.enablePrivateIpAddress : undefined;
            inputs["gatewayDefaultSite"] = args ? args.gatewayDefaultSite : undefined;
            inputs["gatewayType"] = args ? args.gatewayType : undefined;
            inputs["id"] = args ? args.id : undefined;
            inputs["ipConfigurations"] = args ? args.ipConfigurations : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["sku"] = args ? args.sku : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["virtualNetworkGatewayName"] = args ? args.virtualNetworkGatewayName : undefined;
            inputs["vpnClientConfiguration"] = args ? args.vpnClientConfiguration : undefined;
            inputs["vpnGatewayGeneration"] = args ? args.vpnGatewayGeneration : undefined;
            inputs["vpnType"] = args ? args.vpnType : undefined;
            inputs["etag"] = undefined /*out*/;
            inputs["inboundDnsForwardingEndpoint"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["resourceGuid"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["activeActive"] = undefined /*out*/;
            inputs["bgpSettings"] = undefined /*out*/;
            inputs["customRoutes"] = undefined /*out*/;
            inputs["enableBgp"] = undefined /*out*/;
            inputs["enableDnsForwarding"] = undefined /*out*/;
            inputs["enablePrivateIpAddress"] = undefined /*out*/;
            inputs["etag"] = undefined /*out*/;
            inputs["gatewayDefaultSite"] = undefined /*out*/;
            inputs["gatewayType"] = undefined /*out*/;
            inputs["inboundDnsForwardingEndpoint"] = undefined /*out*/;
            inputs["ipConfigurations"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["resourceGuid"] = undefined /*out*/;
            inputs["sku"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["vpnClientConfiguration"] = undefined /*out*/;
            inputs["vpnGatewayGeneration"] = undefined /*out*/;
            inputs["vpnType"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:network/latest:VirtualNetworkGateway" }, { type: "azurerm:network/v20150615:VirtualNetworkGateway" }, { type: "azurerm:network/v20160330:VirtualNetworkGateway" }, { type: "azurerm:network/v20160601:VirtualNetworkGateway" }, { type: "azurerm:network/v20160901:VirtualNetworkGateway" }, { type: "azurerm:network/v20161201:VirtualNetworkGateway" }, { type: "azurerm:network/v20170301:VirtualNetworkGateway" }, { type: "azurerm:network/v20170601:VirtualNetworkGateway" }, { type: "azurerm:network/v20170801:VirtualNetworkGateway" }, { type: "azurerm:network/v20170901:VirtualNetworkGateway" }, { type: "azurerm:network/v20171001:VirtualNetworkGateway" }, { type: "azurerm:network/v20171101:VirtualNetworkGateway" }, { type: "azurerm:network/v20180101:VirtualNetworkGateway" }, { type: "azurerm:network/v20180201:VirtualNetworkGateway" }, { type: "azurerm:network/v20180401:VirtualNetworkGateway" }, { type: "azurerm:network/v20180601:VirtualNetworkGateway" }, { type: "azurerm:network/v20180701:VirtualNetworkGateway" }, { type: "azurerm:network/v20180801:VirtualNetworkGateway" }, { type: "azurerm:network/v20181001:VirtualNetworkGateway" }, { type: "azurerm:network/v20181101:VirtualNetworkGateway" }, { type: "azurerm:network/v20181201:VirtualNetworkGateway" }, { type: "azurerm:network/v20190201:VirtualNetworkGateway" }, { type: "azurerm:network/v20190401:VirtualNetworkGateway" }, { type: "azurerm:network/v20190601:VirtualNetworkGateway" }, { type: "azurerm:network/v20190701:VirtualNetworkGateway" }, { type: "azurerm:network/v20190801:VirtualNetworkGateway" }, { type: "azurerm:network/v20190901:VirtualNetworkGateway" }, { type: "azurerm:network/v20191101:VirtualNetworkGateway" }, { type: "azurerm:network/v20191201:VirtualNetworkGateway" }, { type: "azurerm:network/v20200301:VirtualNetworkGateway" }, { type: "azurerm:network/v20200401:VirtualNetworkGateway" }, { type: "azurerm:network/v20200501:VirtualNetworkGateway" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(VirtualNetworkGateway.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a VirtualNetworkGateway resource.
 */
export interface VirtualNetworkGatewayArgs {
    /**
     * ActiveActive flag.
     */
    readonly activeActive?: pulumi.Input<boolean>;
    /**
     * Virtual network gateway's BGP speaker settings.
     */
    readonly bgpSettings?: pulumi.Input<inputs.network.v20200601.BgpSettings>;
    /**
     * The reference to the address space resource which represents the custom routes address space specified by the customer for virtual network gateway and VpnClient.
     */
    readonly customRoutes?: pulumi.Input<inputs.network.v20200601.AddressSpace>;
    /**
     * Whether BGP is enabled for this virtual network gateway or not.
     */
    readonly enableBgp?: pulumi.Input<boolean>;
    /**
     * Whether dns forwarding is enabled or not.
     */
    readonly enableDnsForwarding?: pulumi.Input<boolean>;
    /**
     * Whether private IP needs to be enabled on this gateway for connections or not.
     */
    readonly enablePrivateIpAddress?: pulumi.Input<boolean>;
    /**
     * The reference to the LocalNetworkGateway resource which represents local network site having default routes. Assign Null value in case of removing existing default site setting.
     */
    readonly gatewayDefaultSite?: pulumi.Input<inputs.network.v20200601.SubResource>;
    /**
     * The type of this virtual network gateway.
     */
    readonly gatewayType?: pulumi.Input<string>;
    /**
     * Resource ID.
     */
    readonly id?: pulumi.Input<string>;
    /**
     * IP configurations for virtual network gateway.
     */
    readonly ipConfigurations?: pulumi.Input<pulumi.Input<inputs.network.v20200601.VirtualNetworkGatewayIPConfiguration>[]>;
    /**
     * Resource location.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The reference to the VirtualNetworkGatewaySku resource which represents the SKU selected for Virtual network gateway.
     */
    readonly sku?: pulumi.Input<inputs.network.v20200601.VirtualNetworkGatewaySku>;
    /**
     * Resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The name of the virtual network gateway.
     */
    readonly virtualNetworkGatewayName: pulumi.Input<string>;
    /**
     * The reference to the VpnClientConfiguration resource which represents the P2S VpnClient configurations.
     */
    readonly vpnClientConfiguration?: pulumi.Input<inputs.network.v20200601.VpnClientConfiguration>;
    /**
     * The generation for this VirtualNetworkGateway. Must be None if gatewayType is not VPN.
     */
    readonly vpnGatewayGeneration?: pulumi.Input<string>;
    /**
     * The type of this virtual network gateway.
     */
    readonly vpnType?: pulumi.Input<string>;
}
