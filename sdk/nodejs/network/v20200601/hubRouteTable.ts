// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * RouteTable resource in a virtual hub.
 *
 * ## Example Usage
 * ### RouteTablePut
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const hubRouteTable = new azurerm.network.v20200601.HubRouteTable("hubRouteTable", {
 *     labels: [
 *         "label1",
 *         "label2",
 *     ],
 *     resourceGroupName: "rg1",
 *     routeTableName: "hubRouteTable1",
 *     routes: [{
 *         destinationType: "CIDR",
 *         destinations: [
 *             "10.0.0.0/8",
 *             "20.0.0.0/8",
 *             "30.0.0.0/8",
 *         ],
 *         name: "route1",
 *         nextHop: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/azureFirewalls/azureFirewall1",
 *         nextHopType: "ResourceId",
 *     }],
 *     virtualHubName: "virtualHub1",
 * });
 *
 * ```
 */
export class HubRouteTable extends pulumi.CustomResource {
    /**
     * Get an existing HubRouteTable resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): HubRouteTable {
        return new HubRouteTable(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:network/v20200601:HubRouteTable';

    /**
     * Returns true if the given object is an instance of HubRouteTable.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is HubRouteTable {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === HubRouteTable.__pulumiType;
    }

    /**
     * List of all connections associated with this route table.
     */
    public /*out*/ readonly associatedConnections!: pulumi.Output<outputs.network.v20200601.SubResourceResponse[]>;
    /**
     * A unique read-only string that changes whenever the resource is updated.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * List of labels associated with this route table.
     */
    public readonly labels!: pulumi.Output<string[] | undefined>;
    /**
     * The name of the resource that is unique within a resource group. This name can be used to access the resource.
     */
    public readonly name!: pulumi.Output<string | undefined>;
    /**
     * List of all connections that advertise to this route table.
     */
    public /*out*/ readonly propagatingConnections!: pulumi.Output<outputs.network.v20200601.SubResourceResponse[]>;
    /**
     * The provisioning state of the RouteTable resource.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * List of all routes.
     */
    public readonly routes!: pulumi.Output<outputs.network.v20200601.HubRouteResponse[] | undefined>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a HubRouteTable resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: HubRouteTableArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: HubRouteTableArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as HubRouteTableArgs | undefined;
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
            inputs["labels"] = args ? args.labels : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["routeTableName"] = args ? args.routeTableName : undefined;
            inputs["routes"] = args ? args.routes : undefined;
            inputs["virtualHubName"] = args ? args.virtualHubName : undefined;
            inputs["associatedConnections"] = undefined /*out*/;
            inputs["etag"] = undefined /*out*/;
            inputs["propagatingConnections"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:network/latest:HubRouteTable" }, { type: "azurerm:network/v20200401:HubRouteTable" }, { type: "azurerm:network/v20200501:HubRouteTable" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(HubRouteTable.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a HubRouteTable resource.
 */
export interface HubRouteTableArgs {
    /**
     * Resource ID.
     */
    readonly id?: pulumi.Input<string>;
    /**
     * List of labels associated with this route table.
     */
    readonly labels?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the resource that is unique within a resource group. This name can be used to access the resource.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The resource group name of the VirtualHub.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the RouteTable.
     */
    readonly routeTableName: pulumi.Input<string>;
    /**
     * List of all routes.
     */
    readonly routes?: pulumi.Input<pulumi.Input<inputs.network.v20200601.HubRoute>[]>;
    /**
     * The name of the VirtualHub.
     */
    readonly virtualHubName: pulumi.Input<string>;
}
