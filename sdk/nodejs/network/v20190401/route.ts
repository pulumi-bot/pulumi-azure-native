// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Route resource.
 */
export class Route extends pulumi.CustomResource {
    /**
     * Get an existing Route resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Route {
        return new Route(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:network/v20190401:Route';

    /**
     * Returns true if the given object is an instance of Route.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Route {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Route.__pulumiType;
    }

    /**
     * The destination CIDR to which the route applies.
     */
    public readonly addressPrefix!: pulumi.Output<string | undefined>;
    /**
     * A unique read-only string that changes whenever the resource is updated.
     */
    public readonly etag!: pulumi.Output<string | undefined>;
    /**
     * The name of the resource that is unique within a resource group. This name can be used to access the resource.
     */
    public readonly name!: pulumi.Output<string | undefined>;
    /**
     * The IP address packets should be forwarded to. Next hop values are only allowed in routes where the next hop type is VirtualAppliance.
     */
    public readonly nextHopIpAddress!: pulumi.Output<string | undefined>;
    /**
     * The type of Azure hop the packet should be sent to.
     */
    public readonly nextHopType!: pulumi.Output<string>;
    /**
     * The provisioning state of the resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
     */
    public readonly provisioningState!: pulumi.Output<string | undefined>;

    /**
     * Create a Route resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RouteArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.nextHopType === undefined) {
                throw new Error("Missing required property 'nextHopType'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.routeName === undefined) {
                throw new Error("Missing required property 'routeName'");
            }
            if (!args || args.routeTableName === undefined) {
                throw new Error("Missing required property 'routeTableName'");
            }
            inputs["addressPrefix"] = args ? args.addressPrefix : undefined;
            inputs["etag"] = args ? args.etag : undefined;
            inputs["id"] = args ? args.id : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["nextHopIpAddress"] = args ? args.nextHopIpAddress : undefined;
            inputs["nextHopType"] = args ? args.nextHopType : undefined;
            inputs["provisioningState"] = args ? args.provisioningState : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["routeName"] = args ? args.routeName : undefined;
            inputs["routeTableName"] = args ? args.routeTableName : undefined;
        } else {
            inputs["addressPrefix"] = undefined /*out*/;
            inputs["etag"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["nextHopIpAddress"] = undefined /*out*/;
            inputs["nextHopType"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:network/latest:Route" }, { type: "azurerm:network/v20150501preview:Route" }, { type: "azurerm:network/v20150615:Route" }, { type: "azurerm:network/v20160330:Route" }, { type: "azurerm:network/v20160601:Route" }, { type: "azurerm:network/v20160901:Route" }, { type: "azurerm:network/v20161201:Route" }, { type: "azurerm:network/v20170301:Route" }, { type: "azurerm:network/v20170601:Route" }, { type: "azurerm:network/v20170801:Route" }, { type: "azurerm:network/v20170901:Route" }, { type: "azurerm:network/v20171001:Route" }, { type: "azurerm:network/v20171101:Route" }, { type: "azurerm:network/v20180101:Route" }, { type: "azurerm:network/v20180201:Route" }, { type: "azurerm:network/v20180401:Route" }, { type: "azurerm:network/v20180601:Route" }, { type: "azurerm:network/v20180701:Route" }, { type: "azurerm:network/v20180801:Route" }, { type: "azurerm:network/v20181001:Route" }, { type: "azurerm:network/v20181101:Route" }, { type: "azurerm:network/v20181201:Route" }, { type: "azurerm:network/v20190201:Route" }, { type: "azurerm:network/v20190601:Route" }, { type: "azurerm:network/v20190701:Route" }, { type: "azurerm:network/v20190801:Route" }, { type: "azurerm:network/v20190901:Route" }, { type: "azurerm:network/v20191101:Route" }, { type: "azurerm:network/v20191201:Route" }, { type: "azurerm:network/v20200301:Route" }, { type: "azurerm:network/v20200401:Route" }, { type: "azurerm:network/v20200501:Route" }, { type: "azurerm:network/v20200601:Route" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(Route.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Route resource.
 */
export interface RouteArgs {
    /**
     * The destination CIDR to which the route applies.
     */
    readonly addressPrefix?: pulumi.Input<string>;
    /**
     * A unique read-only string that changes whenever the resource is updated.
     */
    readonly etag?: pulumi.Input<string>;
    /**
     * Resource ID.
     */
    readonly id?: pulumi.Input<string>;
    /**
     * The name of the resource that is unique within a resource group. This name can be used to access the resource.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The IP address packets should be forwarded to. Next hop values are only allowed in routes where the next hop type is VirtualAppliance.
     */
    readonly nextHopIpAddress?: pulumi.Input<string>;
    /**
     * The type of Azure hop the packet should be sent to.
     */
    readonly nextHopType: pulumi.Input<string>;
    /**
     * The provisioning state of the resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
     */
    readonly provisioningState?: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The name of the route.
     */
    readonly routeName: pulumi.Input<string>;
    /**
     * The name of the route table.
     */
    readonly routeTableName: pulumi.Input<string>;
}
