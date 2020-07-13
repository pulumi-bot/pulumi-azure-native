// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * RouteTable resource in a virtual hub.
 */
export class VirtualHubHubRouteTable extends pulumi.CustomResource {
    /**
     * Get an existing VirtualHubHubRouteTable resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): VirtualHubHubRouteTable {
        return new VirtualHubHubRouteTable(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:network:VirtualHubHubRouteTable';

    /**
     * Returns true if the given object is an instance of VirtualHubHubRouteTable.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VirtualHubHubRouteTable {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VirtualHubHubRouteTable.__pulumiType;
    }

    /**
     * A unique read-only string that changes whenever the resource is updated.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * The name of the resource that is unique within a resource group. This name can be used to access the resource.
     */
    public readonly name!: pulumi.Output<string | undefined>;
    /**
     * Properties of the RouteTable resource.
     */
    public readonly properties!: pulumi.Output<outputs.network.HubRouteTablePropertiesResponse>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a VirtualHubHubRouteTable resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: VirtualHubHubRouteTableArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.routeTableName === undefined) {
                throw new Error("Missing required property 'routeTableName'");
            }
            if (!args || args.virtualHubName === undefined) {
                throw new Error("Missing required property 'virtualHubName'");
            }
        inputs["id"] = args ? args.id : undefined;
        inputs["name"] = args ? args.name : undefined;
        inputs["properties"] = args ? args.properties : undefined;
        inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
        inputs["routeTableName"] = args ? args.routeTableName : undefined;
        inputs["virtualHubName"] = args ? args.virtualHubName : undefined;
        inputs["etag"] = undefined /*out*/;
        inputs["type"] = undefined /*out*/;
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(VirtualHubHubRouteTable.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a VirtualHubHubRouteTable resource.
 */
export interface VirtualHubHubRouteTableArgs {
    /**
     * Resource ID.
     */
    readonly id?: pulumi.Input<string>;
    /**
     * The name of the resource that is unique within a resource group. This name can be used to access the resource.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Properties of the RouteTable resource.
     */
    readonly properties?: pulumi.Input<inputs.network.HubRouteTableProperties>;
    /**
     * The resource group name of the VirtualHub.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the RouteTable.
     */
    readonly routeTableName: pulumi.Input<string>;
    /**
     * The name of the VirtualHub.
     */
    readonly virtualHubName: pulumi.Input<string>;
}