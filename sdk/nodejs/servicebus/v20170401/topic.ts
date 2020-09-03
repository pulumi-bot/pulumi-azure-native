// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Description of topic resource.
 *
 * ## Example Usage
 * ### TopicCreate
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const topic = new azurerm.servicebus.v20170401.Topic("topic", {
 *     enableExpress: true,
 *     namespaceName: "sdk-Namespace-1617",
 *     resourceGroupName: "ArunMonocle",
 *     topicName: "sdk-Topics-5488",
 * });
 *
 * ```
 */
export class Topic extends pulumi.CustomResource {
    /**
     * Get an existing Topic resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Topic {
        return new Topic(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:servicebus/v20170401:Topic';

    /**
     * Returns true if the given object is an instance of Topic.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Topic {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Topic.__pulumiType;
    }

    /**
     * Last time the message was sent, or a request was received, for this topic.
     */
    public /*out*/ readonly accessedAt!: pulumi.Output<string>;
    /**
     * ISO 8601 timespan idle interval after which the topic is automatically deleted. The minimum duration is 5 minutes.
     */
    public readonly autoDeleteOnIdle!: pulumi.Output<string | undefined>;
    /**
     * Message count details
     */
    public /*out*/ readonly countDetails!: pulumi.Output<outputs.servicebus.v20170401.MessageCountDetailsResponse>;
    /**
     * Exact time the message was created.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * ISO 8601 Default message timespan to live value. This is the duration after which the message expires, starting from when the message is sent to Service Bus. This is the default value used when TimeToLive is not set on a message itself.
     */
    public readonly defaultMessageTimeToLive!: pulumi.Output<string | undefined>;
    /**
     * ISO8601 timespan structure that defines the duration of the duplicate detection history. The default value is 10 minutes.
     */
    public readonly duplicateDetectionHistoryTimeWindow!: pulumi.Output<string | undefined>;
    /**
     * Value that indicates whether server-side batched operations are enabled.
     */
    public readonly enableBatchedOperations!: pulumi.Output<boolean | undefined>;
    /**
     * Value that indicates whether Express Entities are enabled. An express topic holds a message in memory temporarily before writing it to persistent storage.
     */
    public readonly enableExpress!: pulumi.Output<boolean | undefined>;
    /**
     * Value that indicates whether the topic to be partitioned across multiple message brokers is enabled.
     */
    public readonly enablePartitioning!: pulumi.Output<boolean | undefined>;
    /**
     * Maximum size of the topic in megabytes, which is the size of the memory allocated for the topic. Default is 1024.
     */
    public readonly maxSizeInMegabytes!: pulumi.Output<number | undefined>;
    /**
     * Resource name
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Value indicating if this topic requires duplicate detection.
     */
    public readonly requiresDuplicateDetection!: pulumi.Output<boolean | undefined>;
    /**
     * Size of the topic, in bytes.
     */
    public /*out*/ readonly sizeInBytes!: pulumi.Output<number>;
    /**
     * Enumerates the possible values for the status of a messaging entity.
     */
    public readonly status!: pulumi.Output<string | undefined>;
    /**
     * Number of subscriptions.
     */
    public /*out*/ readonly subscriptionCount!: pulumi.Output<number>;
    /**
     * Value that indicates whether the topic supports ordering.
     */
    public readonly supportOrdering!: pulumi.Output<boolean | undefined>;
    /**
     * Resource type
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * The exact time the message was updated.
     */
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;

    /**
     * Create a Topic resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TopicArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TopicArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as TopicArgs | undefined;
            if (!args || args.namespaceName === undefined) {
                throw new Error("Missing required property 'namespaceName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.topicName === undefined) {
                throw new Error("Missing required property 'topicName'");
            }
            inputs["autoDeleteOnIdle"] = args ? args.autoDeleteOnIdle : undefined;
            inputs["defaultMessageTimeToLive"] = args ? args.defaultMessageTimeToLive : undefined;
            inputs["duplicateDetectionHistoryTimeWindow"] = args ? args.duplicateDetectionHistoryTimeWindow : undefined;
            inputs["enableBatchedOperations"] = args ? args.enableBatchedOperations : undefined;
            inputs["enableExpress"] = args ? args.enableExpress : undefined;
            inputs["enablePartitioning"] = args ? args.enablePartitioning : undefined;
            inputs["maxSizeInMegabytes"] = args ? args.maxSizeInMegabytes : undefined;
            inputs["namespaceName"] = args ? args.namespaceName : undefined;
            inputs["requiresDuplicateDetection"] = args ? args.requiresDuplicateDetection : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["status"] = args ? args.status : undefined;
            inputs["supportOrdering"] = args ? args.supportOrdering : undefined;
            inputs["topicName"] = args ? args.topicName : undefined;
            inputs["accessedAt"] = undefined /*out*/;
            inputs["countDetails"] = undefined /*out*/;
            inputs["createdAt"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["sizeInBytes"] = undefined /*out*/;
            inputs["subscriptionCount"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["updatedAt"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:servicebus/latest:Topic" }, { type: "azurerm:servicebus/v20140901:Topic" }, { type: "azurerm:servicebus/v20150801:Topic" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(Topic.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Topic resource.
 */
export interface TopicArgs {
    /**
     * ISO 8601 timespan idle interval after which the topic is automatically deleted. The minimum duration is 5 minutes.
     */
    readonly autoDeleteOnIdle?: pulumi.Input<string>;
    /**
     * ISO 8601 Default message timespan to live value. This is the duration after which the message expires, starting from when the message is sent to Service Bus. This is the default value used when TimeToLive is not set on a message itself.
     */
    readonly defaultMessageTimeToLive?: pulumi.Input<string>;
    /**
     * ISO8601 timespan structure that defines the duration of the duplicate detection history. The default value is 10 minutes.
     */
    readonly duplicateDetectionHistoryTimeWindow?: pulumi.Input<string>;
    /**
     * Value that indicates whether server-side batched operations are enabled.
     */
    readonly enableBatchedOperations?: pulumi.Input<boolean>;
    /**
     * Value that indicates whether Express Entities are enabled. An express topic holds a message in memory temporarily before writing it to persistent storage.
     */
    readonly enableExpress?: pulumi.Input<boolean>;
    /**
     * Value that indicates whether the topic to be partitioned across multiple message brokers is enabled.
     */
    readonly enablePartitioning?: pulumi.Input<boolean>;
    /**
     * Maximum size of the topic in megabytes, which is the size of the memory allocated for the topic. Default is 1024.
     */
    readonly maxSizeInMegabytes?: pulumi.Input<number>;
    /**
     * The namespace name
     */
    readonly namespaceName: pulumi.Input<string>;
    /**
     * Value indicating if this topic requires duplicate detection.
     */
    readonly requiresDuplicateDetection?: pulumi.Input<boolean>;
    /**
     * Name of the Resource group within the Azure subscription.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Enumerates the possible values for the status of a messaging entity.
     */
    readonly status?: pulumi.Input<string>;
    /**
     * Value that indicates whether the topic supports ordering.
     */
    readonly supportOrdering?: pulumi.Input<boolean>;
    /**
     * The topic name.
     */
    readonly topicName: pulumi.Input<string>;
}
