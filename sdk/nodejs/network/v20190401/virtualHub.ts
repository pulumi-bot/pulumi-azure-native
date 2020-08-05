// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * VirtualHub Resource.
 */
export class VirtualHub extends pulumi.CustomResource {
    /**
     * Get an existing VirtualHub resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): VirtualHub {
        return new VirtualHub(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:network/v20190401:VirtualHub';

    /**
     * Returns true if the given object is an instance of VirtualHub.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VirtualHub {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VirtualHub.__pulumiType;
    }

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
    public readonly name!: pulumi.Output<string>;
    /**
     * Properties of the virtual hub.
     */
    public /*out*/ readonly properties!: pulumi.Output<outputs.network.v20190401.VirtualHubPropertiesResponse>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a VirtualHub resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VirtualHubArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VirtualHubArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as VirtualHubArgs | undefined;
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["addressPrefix"] = args ? args.addressPrefix : undefined;
            inputs["expressRouteGateway"] = args ? args.expressRouteGateway : undefined;
            inputs["id"] = args ? args.id : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["p2SVpnGateway"] = args ? args.p2SVpnGateway : undefined;
            inputs["provisioningState"] = args ? args.provisioningState : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["routeTable"] = args ? args.routeTable : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["virtualNetworkConnections"] = args ? args.virtualNetworkConnections : undefined;
            inputs["virtualWan"] = args ? args.virtualWan : undefined;
            inputs["vpnGateway"] = args ? args.vpnGateway : undefined;
            inputs["etag"] = undefined /*out*/;
            inputs["properties"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(VirtualHub.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a VirtualHub resource.
 */
export interface VirtualHubArgs {
    /**
     * Address-prefix for this VirtualHub.
     */
    readonly addressPrefix?: pulumi.Input<string>;
    /**
     * The expressRouteGateway associated with this VirtualHub.
     */
    readonly expressRouteGateway?: pulumi.Input<inputs.network.v20190401.SubResource>;
    /**
     * Resource ID.
     */
    readonly id?: pulumi.Input<string>;
    /**
     * Resource location.
     */
    readonly location: pulumi.Input<string>;
    /**
     * The name of the VirtualHub.
     */
    readonly name: pulumi.Input<string>;
    /**
     * The P2SVpnGateway associated with this VirtualHub.
     */
    readonly p2SVpnGateway?: pulumi.Input<inputs.network.v20190401.SubResource>;
    /**
     * The provisioning state of the resource.
     */
    readonly provisioningState?: pulumi.Input<string>;
    /**
     * The resource group name of the VirtualHub.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The routeTable associated with this virtual hub.
     */
    readonly routeTable?: pulumi.Input<inputs.network.v20190401.VirtualHubRouteTable>;
    /**
     * Resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * List of all vnet connections with this VirtualHub.
     */
    readonly virtualNetworkConnections?: pulumi.Input<pulumi.Input<inputs.network.v20190401.HubVirtualNetworkConnection>[]>;
    /**
     * The VirtualWAN to which the VirtualHub belongs.
     */
    readonly virtualWan?: pulumi.Input<inputs.network.v20190401.SubResource>;
    /**
     * The VpnGateway associated with this VirtualHub.
     */
    readonly vpnGateway?: pulumi.Input<inputs.network.v20190401.SubResource>;
}
