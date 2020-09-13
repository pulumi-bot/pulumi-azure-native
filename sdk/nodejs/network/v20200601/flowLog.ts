// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * A flow log resource.
 */
export class FlowLog extends pulumi.CustomResource {
    /**
     * Get an existing FlowLog resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): FlowLog {
        return new FlowLog(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:network/v20200601:FlowLog';

    /**
     * Returns true if the given object is an instance of FlowLog.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FlowLog {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FlowLog.__pulumiType;
    }

    /**
     * Flag to enable/disable flow logging.
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * A unique read-only string that changes whenever the resource is updated.
     */
    public /*out*/ readonly etag!: pulumi.Output<string>;
    /**
     * Parameters that define the configuration of traffic analytics.
     */
    public readonly flowAnalyticsConfiguration!: pulumi.Output<outputs.network.v20200601.TrafficAnalyticsPropertiesResponse | undefined>;
    /**
     * Parameters that define the flow log format.
     */
    public readonly format!: pulumi.Output<outputs.network.v20200601.FlowLogFormatParametersResponse | undefined>;
    /**
     * Resource location.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The provisioning state of the flow log.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * Parameters that define the retention policy for flow log.
     */
    public readonly retentionPolicy!: pulumi.Output<outputs.network.v20200601.RetentionPolicyParametersResponse | undefined>;
    /**
     * ID of the storage account which is used to store the flow log.
     */
    public readonly storageId!: pulumi.Output<string>;
    /**
     * Resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Guid of network security group to which flow log will be applied.
     */
    public /*out*/ readonly targetResourceGuid!: pulumi.Output<string>;
    /**
     * ID of network security group to which flow log will be applied.
     */
    public readonly targetResourceId!: pulumi.Output<string>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a FlowLog resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FlowLogArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.flowLogName === undefined) {
                throw new Error("Missing required property 'flowLogName'");
            }
            if (!args || args.networkWatcherName === undefined) {
                throw new Error("Missing required property 'networkWatcherName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.storageId === undefined) {
                throw new Error("Missing required property 'storageId'");
            }
            if (!args || args.targetResourceId === undefined) {
                throw new Error("Missing required property 'targetResourceId'");
            }
            inputs["enabled"] = args ? args.enabled : undefined;
            inputs["flowAnalyticsConfiguration"] = args ? args.flowAnalyticsConfiguration : undefined;
            inputs["flowLogName"] = args ? args.flowLogName : undefined;
            inputs["format"] = args ? args.format : undefined;
            inputs["id"] = args ? args.id : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["networkWatcherName"] = args ? args.networkWatcherName : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["retentionPolicy"] = args ? args.retentionPolicy : undefined;
            inputs["storageId"] = args ? args.storageId : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["targetResourceId"] = args ? args.targetResourceId : undefined;
            inputs["etag"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["targetResourceGuid"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["enabled"] = undefined /*out*/;
            inputs["etag"] = undefined /*out*/;
            inputs["flowAnalyticsConfiguration"] = undefined /*out*/;
            inputs["format"] = undefined /*out*/;
            inputs["location"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["retentionPolicy"] = undefined /*out*/;
            inputs["storageId"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["targetResourceGuid"] = undefined /*out*/;
            inputs["targetResourceId"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:network/latest:FlowLog" }, { type: "azurerm:network/v20191101:FlowLog" }, { type: "azurerm:network/v20191201:FlowLog" }, { type: "azurerm:network/v20200301:FlowLog" }, { type: "azurerm:network/v20200401:FlowLog" }, { type: "azurerm:network/v20200501:FlowLog" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(FlowLog.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a FlowLog resource.
 */
export interface FlowLogArgs {
    /**
     * Flag to enable/disable flow logging.
     */
    readonly enabled?: pulumi.Input<boolean>;
    /**
     * Parameters that define the configuration of traffic analytics.
     */
    readonly flowAnalyticsConfiguration?: pulumi.Input<inputs.network.v20200601.TrafficAnalyticsProperties>;
    /**
     * The name of the flow log.
     */
    readonly flowLogName: pulumi.Input<string>;
    /**
     * Parameters that define the flow log format.
     */
    readonly format?: pulumi.Input<inputs.network.v20200601.FlowLogFormatParameters>;
    /**
     * Resource ID.
     */
    readonly id?: pulumi.Input<string>;
    /**
     * Resource location.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The name of the network watcher.
     */
    readonly networkWatcherName: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Parameters that define the retention policy for flow log.
     */
    readonly retentionPolicy?: pulumi.Input<inputs.network.v20200601.RetentionPolicyParameters>;
    /**
     * ID of the storage account which is used to store the flow log.
     */
    readonly storageId: pulumi.Input<string>;
    /**
     * Resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * ID of network security group to which flow log will be applied.
     */
    readonly targetResourceId: pulumi.Input<string>;
}
