// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * VirtualHubRouteTableV2 Resource.
 *
 * ## Example Usage
 * ### VirtualHubRouteTableV2Put
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const virtualHubRouteTableV2 = new azurerm.network.v20191201.VirtualHubRouteTableV2("virtualHubRouteTableV2", {
 *     attachedConnections: ["All_Vnets"],
 *     resourceGroupName: "rg1",
 *     routeTableName: "virtualHubRouteTable1a",
 *     routes: [
 *         {
 *             destinationType: "CIDR",
 *             destinations: [
 *                 "20.10.0.0/16",
 *                 "20.20.0.0/16",
 *             ],
 *             nextHopType: "IPAddress",
 *             nextHops: ["10.0.0.68"],
 *         },
 *         {
 *             destinationType: "CIDR",
 *             destinations: ["0.0.0.0/0"],
 *             nextHopType: "IPAddress",
 *             nextHops: ["10.0.0.68"],
 *         },
 *     ],
 *     virtualHubName: "virtualHub1",
 * });
 *
 * ```
 */
export class VirtualHubRouteTableV2 extends pulumi.CustomResource {
    /**
     * Get an existing VirtualHubRouteTableV2 resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): VirtualHubRouteTableV2 {
        return new VirtualHubRouteTableV2(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:network/v20191201:VirtualHubRouteTableV2';

    /**
     * Returns true if the given object is an instance of VirtualHubRouteTableV2.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VirtualHubRouteTableV2 {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VirtualHubRouteTableV2.__pulumiType;
    }

    /**
     * List of all connections attached to this route table v2.
     */
    public readonly attachedConnections!: pulumi.Output<string[] | undefined>;
    /**
     * A unique read-only string that changes whenever the resource is updated.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * The name of the resource that is unique within a resource group. This name can be used to access the resource.
     */
    public readonly name!: pulumi.Output<string | undefined>;
    /**
     * The provisioning state of the virtual hub route table v2 resource.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * List of all routes.
     */
    public readonly routes!: pulumi.Output<outputs.network.v20191201.VirtualHubRouteV2Response[] | undefined>;

    /**
     * Create a VirtualHubRouteTableV2 resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VirtualHubRouteTableV2Args, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VirtualHubRouteTableV2Args, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as VirtualHubRouteTableV2Args | undefined;
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.routeTableName === undefined) {
                throw new Error("Missing required property 'routeTableName'");
            }
            if (!args || args.virtualHubName === undefined) {
                throw new Error("Missing required property 'virtualHubName'");
            }
            inputs["attachedConnections"] = args ? args.attachedConnections : undefined;
            inputs["id"] = args ? args.id : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["routeTableName"] = args ? args.routeTableName : undefined;
            inputs["routes"] = args ? args.routes : undefined;
            inputs["virtualHubName"] = args ? args.virtualHubName : undefined;
            inputs["etag"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:network/latest:VirtualHubRouteTableV2" }, { type: "azurerm:network/v20190901:VirtualHubRouteTableV2" }, { type: "azurerm:network/v20191101:VirtualHubRouteTableV2" }, { type: "azurerm:network/v20200301:VirtualHubRouteTableV2" }, { type: "azurerm:network/v20200401:VirtualHubRouteTableV2" }, { type: "azurerm:network/v20200501:VirtualHubRouteTableV2" }, { type: "azurerm:network/v20200601:VirtualHubRouteTableV2" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(VirtualHubRouteTableV2.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a VirtualHubRouteTableV2 resource.
 */
export interface VirtualHubRouteTableV2Args {
    /**
     * List of all connections attached to this route table v2.
     */
    readonly attachedConnections?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Resource ID.
     */
    readonly id?: pulumi.Input<string>;
    /**
     * The name of the resource that is unique within a resource group. This name can be used to access the resource.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The resource group name of the VirtualHub.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the VirtualHubRouteTableV2.
     */
    readonly routeTableName: pulumi.Input<string>;
    /**
     * List of all routes.
     */
    readonly routes?: pulumi.Input<pulumi.Input<inputs.network.v20191201.VirtualHubRouteV2>[]>;
    /**
     * The name of the VirtualHub.
     */
    readonly virtualHubName: pulumi.Input<string>;
}
