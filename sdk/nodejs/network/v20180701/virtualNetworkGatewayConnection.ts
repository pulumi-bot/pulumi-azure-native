// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * A common class for general resource information
 *
 * ## Example Usage
 * ### CreateVirtualNetworkGatewayConnection_S2S
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const virtualNetworkGatewayConnection = new azurerm.network.v20180701.VirtualNetworkGatewayConnection("virtualNetworkGatewayConnection", {
 *     connectionType: "IPsec",
 *     enableBgp: false,
 *     ipsecPolicies: [],
 *     localNetworkGateway2: {
 *         etag: "W/\"00000000-0000-0000-0000-000000000000\"",
 *         id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/localNetworkGateways/localgw",
 *         location: "centralus",
 *         tags: {},
 *     },
 *     location: "centralus",
 *     resourceGroupName: "rg1",
 *     routingWeight: 0,
 *     sharedKey: "Abc123",
 *     usePolicyBasedTrafficSelectors: false,
 *     virtualNetworkGateway1: {
 *         etag: "W/\"00000000-0000-0000-0000-000000000000\"",
 *         id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/vpngw",
 *         location: "centralus",
 *         tags: {},
 *     },
 *     virtualNetworkGatewayConnectionName: "connS2S",
 * });
 *
 * ```
 */
export class VirtualNetworkGatewayConnection extends pulumi.CustomResource {
    /**
     * Get an existing VirtualNetworkGatewayConnection resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): VirtualNetworkGatewayConnection {
        return new VirtualNetworkGatewayConnection(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:network/v20180701:VirtualNetworkGatewayConnection';

    /**
     * Returns true if the given object is an instance of VirtualNetworkGatewayConnection.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VirtualNetworkGatewayConnection {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VirtualNetworkGatewayConnection.__pulumiType;
    }

    /**
     * The authorizationKey.
     */
    public readonly authorizationKey!: pulumi.Output<string | undefined>;
    /**
     * Virtual network Gateway connection status. Possible values are 'Unknown', 'Connecting', 'Connected' and 'NotConnected'.
     */
    public /*out*/ readonly connectionStatus!: pulumi.Output<string>;
    /**
     * Gateway connection type. Possible values are: 'IPsec','Vnet2Vnet','ExpressRoute', and 'VPNClient.
     */
    public readonly connectionType!: pulumi.Output<string>;
    /**
     * The egress bytes transferred in this connection.
     */
    public /*out*/ readonly egressBytesTransferred!: pulumi.Output<number>;
    /**
     * EnableBgp flag
     */
    public readonly enableBgp!: pulumi.Output<boolean | undefined>;
    /**
     * Gets a unique read-only string that changes whenever the resource is updated.
     */
    public readonly etag!: pulumi.Output<string | undefined>;
    /**
     * Bypass ExpressRoute Gateway for data forwarding
     */
    public readonly expressRouteGatewayBypass!: pulumi.Output<boolean | undefined>;
    /**
     * The ingress bytes transferred in this connection.
     */
    public /*out*/ readonly ingressBytesTransferred!: pulumi.Output<number>;
    /**
     * The IPSec Policies to be considered by this connection.
     */
    public readonly ipsecPolicies!: pulumi.Output<outputs.network.v20180701.IpsecPolicyResponse[] | undefined>;
    /**
     * The reference to local network gateway resource.
     */
    public readonly localNetworkGateway2!: pulumi.Output<outputs.network.v20180701.LocalNetworkGatewayResponse | undefined>;
    /**
     * Resource location.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The reference to peerings resource.
     */
    public readonly peer!: pulumi.Output<outputs.network.v20180701.SubResourceResponse | undefined>;
    /**
     * The provisioning state of the VirtualNetworkGatewayConnection resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * The resource GUID property of the VirtualNetworkGatewayConnection resource.
     */
    public readonly resourceGuid!: pulumi.Output<string | undefined>;
    /**
     * The routing weight.
     */
    public readonly routingWeight!: pulumi.Output<number | undefined>;
    /**
     * The IPSec shared key.
     */
    public readonly sharedKey!: pulumi.Output<string | undefined>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Collection of all tunnels' connection health status.
     */
    public /*out*/ readonly tunnelConnectionStatus!: pulumi.Output<outputs.network.v20180701.TunnelConnectionHealthResponse[]>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * Enable policy-based traffic selectors.
     */
    public readonly usePolicyBasedTrafficSelectors!: pulumi.Output<boolean | undefined>;
    /**
     * The reference to virtual network gateway resource.
     */
    public readonly virtualNetworkGateway1!: pulumi.Output<outputs.network.v20180701.VirtualNetworkGatewayResponse>;
    /**
     * The reference to virtual network gateway resource.
     */
    public readonly virtualNetworkGateway2!: pulumi.Output<outputs.network.v20180701.VirtualNetworkGatewayResponse | undefined>;

    /**
     * Create a VirtualNetworkGatewayConnection resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VirtualNetworkGatewayConnectionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VirtualNetworkGatewayConnectionArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as VirtualNetworkGatewayConnectionArgs | undefined;
            if (!args || args.connectionType === undefined) {
                throw new Error("Missing required property 'connectionType'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.virtualNetworkGateway1 === undefined) {
                throw new Error("Missing required property 'virtualNetworkGateway1'");
            }
            if (!args || args.virtualNetworkGatewayConnectionName === undefined) {
                throw new Error("Missing required property 'virtualNetworkGatewayConnectionName'");
            }
            inputs["authorizationKey"] = args ? args.authorizationKey : undefined;
            inputs["connectionType"] = args ? args.connectionType : undefined;
            inputs["enableBgp"] = args ? args.enableBgp : undefined;
            inputs["etag"] = args ? args.etag : undefined;
            inputs["expressRouteGatewayBypass"] = args ? args.expressRouteGatewayBypass : undefined;
            inputs["id"] = args ? args.id : undefined;
            inputs["ipsecPolicies"] = args ? args.ipsecPolicies : undefined;
            inputs["localNetworkGateway2"] = args ? args.localNetworkGateway2 : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["peer"] = args ? args.peer : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["resourceGuid"] = args ? args.resourceGuid : undefined;
            inputs["routingWeight"] = args ? args.routingWeight : undefined;
            inputs["sharedKey"] = args ? args.sharedKey : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["usePolicyBasedTrafficSelectors"] = args ? args.usePolicyBasedTrafficSelectors : undefined;
            inputs["virtualNetworkGateway1"] = args ? args.virtualNetworkGateway1 : undefined;
            inputs["virtualNetworkGateway2"] = args ? args.virtualNetworkGateway2 : undefined;
            inputs["virtualNetworkGatewayConnectionName"] = args ? args.virtualNetworkGatewayConnectionName : undefined;
            inputs["connectionStatus"] = undefined /*out*/;
            inputs["egressBytesTransferred"] = undefined /*out*/;
            inputs["ingressBytesTransferred"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["tunnelConnectionStatus"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:network/latest:VirtualNetworkGatewayConnection" }, { type: "azurerm:network/v20150615:VirtualNetworkGatewayConnection" }, { type: "azurerm:network/v20160330:VirtualNetworkGatewayConnection" }, { type: "azurerm:network/v20160601:VirtualNetworkGatewayConnection" }, { type: "azurerm:network/v20160901:VirtualNetworkGatewayConnection" }, { type: "azurerm:network/v20161201:VirtualNetworkGatewayConnection" }, { type: "azurerm:network/v20170301:VirtualNetworkGatewayConnection" }, { type: "azurerm:network/v20170601:VirtualNetworkGatewayConnection" }, { type: "azurerm:network/v20170801:VirtualNetworkGatewayConnection" }, { type: "azurerm:network/v20170901:VirtualNetworkGatewayConnection" }, { type: "azurerm:network/v20171001:VirtualNetworkGatewayConnection" }, { type: "azurerm:network/v20171101:VirtualNetworkGatewayConnection" }, { type: "azurerm:network/v20180101:VirtualNetworkGatewayConnection" }, { type: "azurerm:network/v20180201:VirtualNetworkGatewayConnection" }, { type: "azurerm:network/v20180401:VirtualNetworkGatewayConnection" }, { type: "azurerm:network/v20180601:VirtualNetworkGatewayConnection" }, { type: "azurerm:network/v20180801:VirtualNetworkGatewayConnection" }, { type: "azurerm:network/v20181001:VirtualNetworkGatewayConnection" }, { type: "azurerm:network/v20181101:VirtualNetworkGatewayConnection" }, { type: "azurerm:network/v20181201:VirtualNetworkGatewayConnection" }, { type: "azurerm:network/v20190201:VirtualNetworkGatewayConnection" }, { type: "azurerm:network/v20190401:VirtualNetworkGatewayConnection" }, { type: "azurerm:network/v20190601:VirtualNetworkGatewayConnection" }, { type: "azurerm:network/v20190701:VirtualNetworkGatewayConnection" }, { type: "azurerm:network/v20190801:VirtualNetworkGatewayConnection" }, { type: "azurerm:network/v20190901:VirtualNetworkGatewayConnection" }, { type: "azurerm:network/v20191101:VirtualNetworkGatewayConnection" }, { type: "azurerm:network/v20191201:VirtualNetworkGatewayConnection" }, { type: "azurerm:network/v20200301:VirtualNetworkGatewayConnection" }, { type: "azurerm:network/v20200401:VirtualNetworkGatewayConnection" }, { type: "azurerm:network/v20200501:VirtualNetworkGatewayConnection" }, { type: "azurerm:network/v20200601:VirtualNetworkGatewayConnection" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(VirtualNetworkGatewayConnection.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a VirtualNetworkGatewayConnection resource.
 */
export interface VirtualNetworkGatewayConnectionArgs {
    /**
     * The authorizationKey.
     */
    readonly authorizationKey?: pulumi.Input<string>;
    /**
     * Gateway connection type. Possible values are: 'IPsec','Vnet2Vnet','ExpressRoute', and 'VPNClient.
     */
    readonly connectionType: pulumi.Input<string>;
    /**
     * EnableBgp flag
     */
    readonly enableBgp?: pulumi.Input<boolean>;
    /**
     * Gets a unique read-only string that changes whenever the resource is updated.
     */
    readonly etag?: pulumi.Input<string>;
    /**
     * Bypass ExpressRoute Gateway for data forwarding
     */
    readonly expressRouteGatewayBypass?: pulumi.Input<boolean>;
    /**
     * Resource ID.
     */
    readonly id?: pulumi.Input<string>;
    /**
     * The IPSec Policies to be considered by this connection.
     */
    readonly ipsecPolicies?: pulumi.Input<pulumi.Input<inputs.network.v20180701.IpsecPolicy>[]>;
    /**
     * The reference to local network gateway resource.
     */
    readonly localNetworkGateway2?: pulumi.Input<inputs.network.v20180701.LocalNetworkGateway>;
    /**
     * Resource location.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The reference to peerings resource.
     */
    readonly peer?: pulumi.Input<inputs.network.v20180701.SubResource>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The resource GUID property of the VirtualNetworkGatewayConnection resource.
     */
    readonly resourceGuid?: pulumi.Input<string>;
    /**
     * The routing weight.
     */
    readonly routingWeight?: pulumi.Input<number>;
    /**
     * The IPSec shared key.
     */
    readonly sharedKey?: pulumi.Input<string>;
    /**
     * Resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Enable policy-based traffic selectors.
     */
    readonly usePolicyBasedTrafficSelectors?: pulumi.Input<boolean>;
    /**
     * The reference to virtual network gateway resource.
     */
    readonly virtualNetworkGateway1: pulumi.Input<inputs.network.v20180701.VirtualNetworkGateway>;
    /**
     * The reference to virtual network gateway resource.
     */
    readonly virtualNetworkGateway2?: pulumi.Input<inputs.network.v20180701.VirtualNetworkGateway>;
    /**
     * The name of the virtual network gateway connection.
     */
    readonly virtualNetworkGatewayConnectionName: pulumi.Input<string>;
}
