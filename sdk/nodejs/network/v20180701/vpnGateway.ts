// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * VpnGateway Resource.
 *
 * ## Example Usage
 * ### VpnGatewayPut
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const vpnGateway = new azurerm.network.v20180701.VpnGateway("vpnGateway", {
 *     bgpSettings: {
 *         asn: 65515,
 *         bgpPeeringAddress: "10.0.1.30",
 *         peerWeight: 0,
 *     },
 *     connections: [{
 *         name: "vpnConnection1",
 *     }],
 *     gatewayName: "gateway1",
 *     location: "West US",
 *     policies: {
 *         allowBranchToBranchTraffic: true,
 *         allowVnetToVnetTraffic: false,
 *     },
 *     resourceGroupName: "rg1",
 *     tags: {
 *         key1: "value1",
 *     },
 *     virtualHub: {
 *         id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1",
 *     },
 * });
 *
 * ```
 */
export class VpnGateway extends pulumi.CustomResource {
    /**
     * Get an existing VpnGateway resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): VpnGateway {
        return new VpnGateway(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:network/v20180701:VpnGateway';

    /**
     * Returns true if the given object is an instance of VpnGateway.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VpnGateway {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VpnGateway.__pulumiType;
    }

    /**
     * Local network gateway's BGP speaker settings.
     */
    public readonly bgpSettings!: pulumi.Output<outputs.network.v20180701.BgpSettingsResponse | undefined>;
    /**
     * list of all vpn connections to the gateway.
     */
    public readonly connections!: pulumi.Output<outputs.network.v20180701.VpnConnectionResponse[] | undefined>;
    /**
     * Gets a unique read-only string that changes whenever the resource is updated.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * Resource location.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The policies applied to this vpn gateway.
     */
    public readonly policies!: pulumi.Output<outputs.network.v20180701.PoliciesResponse | undefined>;
    /**
     * The provisioning state of the resource.
     */
    public readonly provisioningState!: pulumi.Output<string | undefined>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * The VirtualHub to which the gateway belongs
     */
    public readonly virtualHub!: pulumi.Output<outputs.network.v20180701.SubResourceResponse | undefined>;

    /**
     * Create a VpnGateway resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VpnGatewayArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VpnGatewayArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as VpnGatewayArgs | undefined;
            if (!args || args.gatewayName === undefined) {
                throw new Error("Missing required property 'gatewayName'");
            }
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["bgpSettings"] = args ? args.bgpSettings : undefined;
            inputs["connections"] = args ? args.connections : undefined;
            inputs["gatewayName"] = args ? args.gatewayName : undefined;
            inputs["id"] = args ? args.id : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["policies"] = args ? args.policies : undefined;
            inputs["provisioningState"] = args ? args.provisioningState : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["virtualHub"] = args ? args.virtualHub : undefined;
            inputs["etag"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:network/latest:VpnGateway" }, { type: "azurerm:network/v20180401:VpnGateway" }, { type: "azurerm:network/v20180601:VpnGateway" }, { type: "azurerm:network/v20180801:VpnGateway" }, { type: "azurerm:network/v20181001:VpnGateway" }, { type: "azurerm:network/v20181101:VpnGateway" }, { type: "azurerm:network/v20181201:VpnGateway" }, { type: "azurerm:network/v20190201:VpnGateway" }, { type: "azurerm:network/v20190401:VpnGateway" }, { type: "azurerm:network/v20190601:VpnGateway" }, { type: "azurerm:network/v20190701:VpnGateway" }, { type: "azurerm:network/v20190801:VpnGateway" }, { type: "azurerm:network/v20190901:VpnGateway" }, { type: "azurerm:network/v20191101:VpnGateway" }, { type: "azurerm:network/v20191201:VpnGateway" }, { type: "azurerm:network/v20200301:VpnGateway" }, { type: "azurerm:network/v20200401:VpnGateway" }, { type: "azurerm:network/v20200501:VpnGateway" }, { type: "azurerm:network/v20200601:VpnGateway" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(VpnGateway.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a VpnGateway resource.
 */
export interface VpnGatewayArgs {
    /**
     * Local network gateway's BGP speaker settings.
     */
    readonly bgpSettings?: pulumi.Input<inputs.network.v20180701.BgpSettings>;
    /**
     * list of all vpn connections to the gateway.
     */
    readonly connections?: pulumi.Input<pulumi.Input<inputs.network.v20180701.VpnConnection>[]>;
    /**
     * The name of the gateway.
     */
    readonly gatewayName: pulumi.Input<string>;
    /**
     * Resource ID.
     */
    readonly id?: pulumi.Input<string>;
    /**
     * Resource location.
     */
    readonly location: pulumi.Input<string>;
    /**
     * The policies applied to this vpn gateway.
     */
    readonly policies?: pulumi.Input<inputs.network.v20180701.Policies>;
    /**
     * The provisioning state of the resource.
     */
    readonly provisioningState?: pulumi.Input<string>;
    /**
     * The resource group name of the VpnGateway.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The VirtualHub to which the gateway belongs
     */
    readonly virtualHub?: pulumi.Input<inputs.network.v20180701.SubResource>;
}
