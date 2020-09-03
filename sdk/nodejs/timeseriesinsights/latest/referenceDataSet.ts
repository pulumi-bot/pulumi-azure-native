// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * A reference data set provides metadata about the events in an environment. Metadata in the reference data set will be joined with events as they are read from event sources. The metadata that makes up the reference data set is uploaded or modified through the Time Series Insights data plane APIs.
 *
 * ## Example Usage
 * ### ReferenceDataSetsCreate
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const referenceDataSet = new azurerm.timeseriesinsights.latest.ReferenceDataSet("referenceDataSet", {
 *     environmentName: "env1",
 *     keyProperties: [
 *         {
 *             name: "DeviceId1",
 *             type: "String",
 *         },
 *         {
 *             name: "DeviceFloor",
 *             type: "Double",
 *         },
 *     ],
 *     location: "West US",
 *     referenceDataSetName: "rds1",
 *     resourceGroupName: "rg1",
 * });
 *
 * ```
 */
export class ReferenceDataSet extends pulumi.CustomResource {
    /**
     * Get an existing ReferenceDataSet resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ReferenceDataSet {
        return new ReferenceDataSet(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:timeseriesinsights/latest:ReferenceDataSet';

    /**
     * Returns true if the given object is an instance of ReferenceDataSet.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ReferenceDataSet {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ReferenceDataSet.__pulumiType;
    }

    /**
     * The time the resource was created.
     */
    public /*out*/ readonly creationTime!: pulumi.Output<string>;
    /**
     * The reference data set key comparison behavior can be set using this property. By default, the value is 'Ordinal' - which means case sensitive key comparison will be performed while joining reference data with events or while adding new reference data. When 'OrdinalIgnoreCase' is set, case insensitive comparison will be used.
     */
    public readonly dataStringComparisonBehavior!: pulumi.Output<string | undefined>;
    /**
     * The list of key properties for the reference data set.
     */
    public readonly keyProperties!: pulumi.Output<outputs.timeseriesinsights.latest.ReferenceDataSetKeyPropertyResponse[]>;
    /**
     * Resource location
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Resource name
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Provisioning state of the resource.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string | undefined>;
    /**
     * Resource tags
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Resource type
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a ReferenceDataSet resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ReferenceDataSetArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ReferenceDataSetArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as ReferenceDataSetArgs | undefined;
            if (!args || args.environmentName === undefined) {
                throw new Error("Missing required property 'environmentName'");
            }
            if (!args || args.keyProperties === undefined) {
                throw new Error("Missing required property 'keyProperties'");
            }
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            if (!args || args.referenceDataSetName === undefined) {
                throw new Error("Missing required property 'referenceDataSetName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["dataStringComparisonBehavior"] = args ? args.dataStringComparisonBehavior : undefined;
            inputs["environmentName"] = args ? args.environmentName : undefined;
            inputs["keyProperties"] = args ? args.keyProperties : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["referenceDataSetName"] = args ? args.referenceDataSetName : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["creationTime"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:timeseriesinsights/v20170228preview:ReferenceDataSet" }, { type: "azurerm:timeseriesinsights/v20171115:ReferenceDataSet" }, { type: "azurerm:timeseriesinsights/v20180815preview:ReferenceDataSet" }, { type: "azurerm:timeseriesinsights/v20200515:ReferenceDataSet" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(ReferenceDataSet.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a ReferenceDataSet resource.
 */
export interface ReferenceDataSetArgs {
    /**
     * The reference data set key comparison behavior can be set using this property. By default, the value is 'Ordinal' - which means case sensitive key comparison will be performed while joining reference data with events or while adding new reference data. When 'OrdinalIgnoreCase' is set, case insensitive comparison will be used.
     */
    readonly dataStringComparisonBehavior?: pulumi.Input<string>;
    /**
     * The name of the Time Series Insights environment associated with the specified resource group.
     */
    readonly environmentName: pulumi.Input<string>;
    /**
     * The list of key properties for the reference data set.
     */
    readonly keyProperties: pulumi.Input<pulumi.Input<inputs.timeseriesinsights.latest.ReferenceDataSetKeyProperty>[]>;
    /**
     * The location of the resource.
     */
    readonly location: pulumi.Input<string>;
    /**
     * Name of the reference data set.
     */
    readonly referenceDataSetName: pulumi.Input<string>;
    /**
     * Name of an Azure Resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Key-value pairs of additional properties for the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
