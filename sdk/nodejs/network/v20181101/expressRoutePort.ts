// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * ExpressRoutePort resource definition.
 *
 * ## Example Usage
 * ### ExpressRoutePortCreate
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const expressRoutePort = new azurerm.network.v20181101.ExpressRoutePort("expressRoutePort", {
 *     bandwidthInGbps: 100,
 *     encapsulation: "QinQ",
 *     expressRoutePortName: "portName",
 *     location: "westus",
 *     peeringLocation: "peeringLocationName",
 *     resourceGroupName: "rg1",
 * });
 *
 * ```
 * ### ExpressRoutePortUpdateLink
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const expressRoutePort = new azurerm.network.v20181101.ExpressRoutePort("expressRoutePort", {
 *     bandwidthInGbps: 100,
 *     encapsulation: "QinQ",
 *     expressRoutePortName: "portName",
 *     links: [{
 *         name: "link1",
 *     }],
 *     location: "westus",
 *     peeringLocation: "peeringLocationName",
 *     resourceGroupName: "rg1",
 * });
 *
 * ```
 */
export class ExpressRoutePort extends pulumi.CustomResource {
    /**
     * Get an existing ExpressRoutePort resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ExpressRoutePort {
        return new ExpressRoutePort(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:network/v20181101:ExpressRoutePort';

    /**
     * Returns true if the given object is an instance of ExpressRoutePort.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ExpressRoutePort {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ExpressRoutePort.__pulumiType;
    }

    /**
     * Date of the physical port allocation to be used in Letter of Authorization.
     */
    public /*out*/ readonly allocationDate!: pulumi.Output<string>;
    /**
     * Bandwidth of procured ports in Gbps
     */
    public readonly bandwidthInGbps!: pulumi.Output<number | undefined>;
    /**
     * Reference the ExpressRoute circuit(s) that are provisioned on this ExpressRoutePort resource.
     */
    public /*out*/ readonly circuits!: pulumi.Output<outputs.network.v20181101.SubResourceResponse[]>;
    /**
     * Encapsulation method on physical ports.
     */
    public readonly encapsulation!: pulumi.Output<string | undefined>;
    /**
     * A unique read-only string that changes whenever the resource is updated.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * Ether type of the physical port.
     */
    public /*out*/ readonly etherType!: pulumi.Output<string>;
    /**
     * The set of physical links of the ExpressRoutePort resource
     */
    public readonly links!: pulumi.Output<outputs.network.v20181101.ExpressRouteLinkResponse[] | undefined>;
    /**
     * Resource location.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * Maximum transmission unit of the physical port pair(s)
     */
    public /*out*/ readonly mtu!: pulumi.Output<string>;
    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The name of the peering location that the ExpressRoutePort is mapped to physically.
     */
    public readonly peeringLocation!: pulumi.Output<string | undefined>;
    /**
     * Aggregate Gbps of associated circuit bandwidths.
     */
    public /*out*/ readonly provisionedBandwidthInGbps!: pulumi.Output<number>;
    /**
     * The provisioning state of the ExpressRoutePort resource. Possible values are: 'Succeeded', 'Updating', 'Deleting', and 'Failed'.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * The resource GUID property of the ExpressRoutePort resource.
     */
    public readonly resourceGuid!: pulumi.Output<string | undefined>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a ExpressRoutePort resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ExpressRoutePortArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.expressRoutePortName === undefined) {
                throw new Error("Missing required property 'expressRoutePortName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["bandwidthInGbps"] = args ? args.bandwidthInGbps : undefined;
            inputs["encapsulation"] = args ? args.encapsulation : undefined;
            inputs["expressRoutePortName"] = args ? args.expressRoutePortName : undefined;
            inputs["id"] = args ? args.id : undefined;
            inputs["links"] = args ? args.links : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["peeringLocation"] = args ? args.peeringLocation : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["resourceGuid"] = args ? args.resourceGuid : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["allocationDate"] = undefined /*out*/;
            inputs["circuits"] = undefined /*out*/;
            inputs["etag"] = undefined /*out*/;
            inputs["etherType"] = undefined /*out*/;
            inputs["mtu"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisionedBandwidthInGbps"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["allocationDate"] = undefined /*out*/;
            inputs["bandwidthInGbps"] = undefined /*out*/;
            inputs["circuits"] = undefined /*out*/;
            inputs["encapsulation"] = undefined /*out*/;
            inputs["etag"] = undefined /*out*/;
            inputs["etherType"] = undefined /*out*/;
            inputs["links"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["mtu"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["peeringLocation"] = undefined /*out*/;
            inputs["provisionedBandwidthInGbps"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["resourceGuid"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:network/latest:ExpressRoutePort" }, { type: "azurerm:network/v20180801:ExpressRoutePort" }, { type: "azurerm:network/v20181001:ExpressRoutePort" }, { type: "azurerm:network/v20181201:ExpressRoutePort" }, { type: "azurerm:network/v20190201:ExpressRoutePort" }, { type: "azurerm:network/v20190401:ExpressRoutePort" }, { type: "azurerm:network/v20190601:ExpressRoutePort" }, { type: "azurerm:network/v20190701:ExpressRoutePort" }, { type: "azurerm:network/v20190801:ExpressRoutePort" }, { type: "azurerm:network/v20190901:ExpressRoutePort" }, { type: "azurerm:network/v20191101:ExpressRoutePort" }, { type: "azurerm:network/v20191201:ExpressRoutePort" }, { type: "azurerm:network/v20200301:ExpressRoutePort" }, { type: "azurerm:network/v20200401:ExpressRoutePort" }, { type: "azurerm:network/v20200501:ExpressRoutePort" }, { type: "azurerm:network/v20200601:ExpressRoutePort" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(ExpressRoutePort.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ExpressRoutePort resource.
 */
export interface ExpressRoutePortArgs {
    /**
     * Bandwidth of procured ports in Gbps
     */
    readonly bandwidthInGbps?: pulumi.Input<number>;
    /**
     * Encapsulation method on physical ports.
     */
    readonly encapsulation?: pulumi.Input<string>;
    /**
     * The name of the ExpressRoutePort resource.
     */
    readonly expressRoutePortName: pulumi.Input<string>;
    /**
     * Resource ID.
     */
    readonly id?: pulumi.Input<string>;
    /**
     * The set of physical links of the ExpressRoutePort resource
     */
    readonly links?: pulumi.Input<pulumi.Input<inputs.network.v20181101.ExpressRouteLink>[]>;
    /**
     * Resource location.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name of the peering location that the ExpressRoutePort is mapped to physically.
     */
    readonly peeringLocation?: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The resource GUID property of the ExpressRoutePort resource.
     */
    readonly resourceGuid?: pulumi.Input<string>;
    /**
     * Resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
