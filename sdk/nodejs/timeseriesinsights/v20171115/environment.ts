// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * An environment is a set of time-series data available for query, and is the top level Azure Time Series Insights resource.
 */
export class Environment extends pulumi.CustomResource {
    /**
     * Get an existing Environment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Environment {
        return new Environment(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:timeseriesinsights/v20171115:Environment';

    /**
     * Returns true if the given object is an instance of Environment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Environment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Environment.__pulumiType;
    }

    /**
     * The time the resource was created.
     */
    public /*out*/ readonly creationTime!: pulumi.Output<string>;
    /**
     * The fully qualified domain name used to access the environment data, e.g. to query the environment's events or upload reference data for the environment.
     */
    public /*out*/ readonly dataAccessFqdn!: pulumi.Output<string>;
    /**
     * An id used to access the environment data, e.g. to query the environment's events or upload reference data for the environment.
     */
    public /*out*/ readonly dataAccessId!: pulumi.Output<string>;
    /**
     * ISO8601 timespan specifying the minimum number of days the environment's events will be available for query.
     */
    public readonly dataRetentionTime!: pulumi.Output<string>;
    /**
     * Resource location
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Resource name
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The list of partition keys according to which the data in the environment will be ordered.
     */
    public readonly partitionKeyProperties!: pulumi.Output<outputs.timeseriesinsights.v20171115.PartitionKeyPropertyResponse[] | undefined>;
    /**
     * Provisioning state of the resource.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string | undefined>;
    /**
     * The sku determines the capacity of the environment, the SLA (in queries-per-minute and total capacity), and the billing rate.
     */
    public readonly sku!: pulumi.Output<outputs.timeseriesinsights.v20171115.SkuResponse | undefined>;
    /**
     * An object that represents the status of the environment, and its internal state in the Time Series Insights service.
     */
    public /*out*/ readonly status!: pulumi.Output<outputs.timeseriesinsights.v20171115.EnvironmentStatusResponse | undefined>;
    /**
     * The behavior the Time Series Insights service should take when the environment's capacity has been exceeded. If "PauseIngress" is specified, new events will not be read from the event source. If "PurgeOldData" is specified, new events will continue to be read and old events will be deleted from the environment. The default behavior is PurgeOldData.
     */
    public readonly storageLimitExceededBehavior!: pulumi.Output<string | undefined>;
    /**
     * Resource tags
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Resource type
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a Environment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EnvironmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EnvironmentArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as EnvironmentArgs | undefined;
            if (!args || args.dataRetentionTime === undefined) {
                throw new Error("Missing required property 'dataRetentionTime'");
            }
            if (!args || args.environmentName === undefined) {
                throw new Error("Missing required property 'environmentName'");
            }
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.sku === undefined) {
                throw new Error("Missing required property 'sku'");
            }
            inputs["dataRetentionTime"] = args ? args.dataRetentionTime : undefined;
            inputs["environmentName"] = args ? args.environmentName : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["partitionKeyProperties"] = args ? args.partitionKeyProperties : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["sku"] = args ? args.sku : undefined;
            inputs["storageLimitExceededBehavior"] = args ? args.storageLimitExceededBehavior : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["creationTime"] = undefined /*out*/;
            inputs["dataAccessFqdn"] = undefined /*out*/;
            inputs["dataAccessId"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:timeseriesinsights/latest:Environment" }, { type: "azurerm:timeseriesinsights/v20170228preview:Environment" }, { type: "azurerm:timeseriesinsights/v20180815preview:Environment" }, { type: "azurerm:timeseriesinsights/v20200515:Environment" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(Environment.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Environment resource.
 */
export interface EnvironmentArgs {
    /**
     * ISO8601 timespan specifying the minimum number of days the environment's events will be available for query.
     */
    readonly dataRetentionTime: pulumi.Input<string>;
    /**
     * Name of the environment
     */
    readonly environmentName: pulumi.Input<string>;
    /**
     * The location of the resource.
     */
    readonly location: pulumi.Input<string>;
    /**
     * The list of partition keys according to which the data in the environment will be ordered.
     */
    readonly partitionKeyProperties?: pulumi.Input<pulumi.Input<inputs.timeseriesinsights.v20171115.PartitionKeyProperty>[]>;
    /**
     * Name of an Azure Resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The sku determines the capacity of the environment, the SLA (in queries-per-minute and total capacity), and the billing rate.
     */
    readonly sku: pulumi.Input<inputs.timeseriesinsights.v20171115.Sku>;
    /**
     * The behavior the Time Series Insights service should take when the environment's capacity has been exceeded. If "PauseIngress" is specified, new events will not be read from the event source. If "PurgeOldData" is specified, new events will continue to be read and old events will be deleted from the environment. The default behavior is PurgeOldData.
     */
    readonly storageLimitExceededBehavior?: pulumi.Input<string>;
    /**
     * Key-value pairs of additional properties for the resource.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
