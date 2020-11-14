// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

/**
 * Description of queue Resource.
 */
export class Queue extends pulumi.CustomResource {
    /**
     * Get an existing Queue resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Queue {
        return new Queue(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-nextgen:servicebus/v20170401:Queue';

    /**
     * Returns true if the given object is an instance of Queue.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Queue {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Queue.__pulumiType;
    }

    /**
     * Last time a message was sent, or the last time there was a receive request to this queue.
     */
    public /*out*/ readonly accessedAt!: pulumi.Output<string>;
    /**
     * ISO 8061 timeSpan idle interval after which the queue is automatically deleted. The minimum duration is 5 minutes.
     */
    public readonly autoDeleteOnIdle!: pulumi.Output<string | undefined>;
    /**
     * Message Count Details.
     */
    public /*out*/ readonly countDetails!: pulumi.Output<outputs.servicebus.v20170401.MessageCountDetailsResponse>;
    /**
     * The exact time the message was created.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * A value that indicates whether this queue has dead letter support when a message expires.
     */
    public readonly deadLetteringOnMessageExpiration!: pulumi.Output<boolean | undefined>;
    /**
     * ISO 8601 default message timespan to live value. This is the duration after which the message expires, starting from when the message is sent to Service Bus. This is the default value used when TimeToLive is not set on a message itself.
     */
    public readonly defaultMessageTimeToLive!: pulumi.Output<string | undefined>;
    /**
     * ISO 8601 timeSpan structure that defines the duration of the duplicate detection history. The default value is 10 minutes.
     */
    public readonly duplicateDetectionHistoryTimeWindow!: pulumi.Output<string | undefined>;
    /**
     * Value that indicates whether server-side batched operations are enabled.
     */
    public readonly enableBatchedOperations!: pulumi.Output<boolean | undefined>;
    /**
     * A value that indicates whether Express Entities are enabled. An express queue holds a message in memory temporarily before writing it to persistent storage.
     */
    public readonly enableExpress!: pulumi.Output<boolean | undefined>;
    /**
     * A value that indicates whether the queue is to be partitioned across multiple message brokers.
     */
    public readonly enablePartitioning!: pulumi.Output<boolean | undefined>;
    /**
     * Queue/Topic name to forward the Dead Letter message
     */
    public readonly forwardDeadLetteredMessagesTo!: pulumi.Output<string | undefined>;
    /**
     * Queue/Topic name to forward the messages
     */
    public readonly forwardTo!: pulumi.Output<string | undefined>;
    /**
     * ISO 8601 timespan duration of a peek-lock; that is, the amount of time that the message is locked for other receivers. The maximum value for LockDuration is 5 minutes; the default value is 1 minute.
     */
    public readonly lockDuration!: pulumi.Output<string | undefined>;
    /**
     * The maximum delivery count. A message is automatically deadlettered after this number of deliveries. default value is 10.
     */
    public readonly maxDeliveryCount!: pulumi.Output<number | undefined>;
    /**
     * The maximum size of the queue in megabytes, which is the size of memory allocated for the queue. Default is 1024.
     */
    public readonly maxSizeInMegabytes!: pulumi.Output<number | undefined>;
    /**
     * The number of messages in the queue.
     */
    public /*out*/ readonly messageCount!: pulumi.Output<number>;
    /**
     * Resource name
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * A value indicating if this queue requires duplicate detection.
     */
    public readonly requiresDuplicateDetection!: pulumi.Output<boolean | undefined>;
    /**
     * A value that indicates whether the queue supports the concept of sessions.
     */
    public readonly requiresSession!: pulumi.Output<boolean | undefined>;
    /**
     * The size of the queue, in bytes.
     */
    public /*out*/ readonly sizeInBytes!: pulumi.Output<number>;
    /**
     * Enumerates the possible values for the status of a messaging entity.
     */
    public readonly status!: pulumi.Output<string | undefined>;
    /**
     * Resource type
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * The exact time the message was updated.
     */
    public /*out*/ readonly updatedAt!: pulumi.Output<string>;

    /**
     * Create a Queue resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: QueueArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.namespaceName === undefined) {
                throw new Error("Missing required property 'namespaceName'");
            }
            if (!args || args.queueName === undefined) {
                throw new Error("Missing required property 'queueName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            inputs["autoDeleteOnIdle"] = args ? args.autoDeleteOnIdle : undefined;
            inputs["deadLetteringOnMessageExpiration"] = args ? args.deadLetteringOnMessageExpiration : undefined;
            inputs["defaultMessageTimeToLive"] = args ? args.defaultMessageTimeToLive : undefined;
            inputs["duplicateDetectionHistoryTimeWindow"] = args ? args.duplicateDetectionHistoryTimeWindow : undefined;
            inputs["enableBatchedOperations"] = args ? args.enableBatchedOperations : undefined;
            inputs["enableExpress"] = args ? args.enableExpress : undefined;
            inputs["enablePartitioning"] = args ? args.enablePartitioning : undefined;
            inputs["forwardDeadLetteredMessagesTo"] = args ? args.forwardDeadLetteredMessagesTo : undefined;
            inputs["forwardTo"] = args ? args.forwardTo : undefined;
            inputs["lockDuration"] = args ? args.lockDuration : undefined;
            inputs["maxDeliveryCount"] = args ? args.maxDeliveryCount : undefined;
            inputs["maxSizeInMegabytes"] = args ? args.maxSizeInMegabytes : undefined;
            inputs["namespaceName"] = args ? args.namespaceName : undefined;
            inputs["queueName"] = args ? args.queueName : undefined;
            inputs["requiresDuplicateDetection"] = args ? args.requiresDuplicateDetection : undefined;
            inputs["requiresSession"] = args ? args.requiresSession : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["status"] = args ? args.status : undefined;
            inputs["accessedAt"] = undefined /*out*/;
            inputs["countDetails"] = undefined /*out*/;
            inputs["createdAt"] = undefined /*out*/;
            inputs["messageCount"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["sizeInBytes"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["updatedAt"] = undefined /*out*/;
        } else {
            inputs["accessedAt"] = undefined /*out*/;
            inputs["autoDeleteOnIdle"] = undefined /*out*/;
            inputs["countDetails"] = undefined /*out*/;
            inputs["createdAt"] = undefined /*out*/;
            inputs["deadLetteringOnMessageExpiration"] = undefined /*out*/;
            inputs["defaultMessageTimeToLive"] = undefined /*out*/;
            inputs["duplicateDetectionHistoryTimeWindow"] = undefined /*out*/;
            inputs["enableBatchedOperations"] = undefined /*out*/;
            inputs["enableExpress"] = undefined /*out*/;
            inputs["enablePartitioning"] = undefined /*out*/;
            inputs["forwardDeadLetteredMessagesTo"] = undefined /*out*/;
            inputs["forwardTo"] = undefined /*out*/;
            inputs["lockDuration"] = undefined /*out*/;
            inputs["maxDeliveryCount"] = undefined /*out*/;
            inputs["maxSizeInMegabytes"] = undefined /*out*/;
            inputs["messageCount"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["requiresDuplicateDetection"] = undefined /*out*/;
            inputs["requiresSession"] = undefined /*out*/;
            inputs["sizeInBytes"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["updatedAt"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:servicebus/latest:Queue" }, { type: "azure-nextgen:servicebus/v20140901:Queue" }, { type: "azure-nextgen:servicebus/v20150801:Queue" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(Queue.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Queue resource.
 */
export interface QueueArgs {
    /**
     * ISO 8061 timeSpan idle interval after which the queue is automatically deleted. The minimum duration is 5 minutes.
     */
    readonly autoDeleteOnIdle?: pulumi.Input<string>;
    /**
     * A value that indicates whether this queue has dead letter support when a message expires.
     */
    readonly deadLetteringOnMessageExpiration?: pulumi.Input<boolean>;
    /**
     * ISO 8601 default message timespan to live value. This is the duration after which the message expires, starting from when the message is sent to Service Bus. This is the default value used when TimeToLive is not set on a message itself.
     */
    readonly defaultMessageTimeToLive?: pulumi.Input<string>;
    /**
     * ISO 8601 timeSpan structure that defines the duration of the duplicate detection history. The default value is 10 minutes.
     */
    readonly duplicateDetectionHistoryTimeWindow?: pulumi.Input<string>;
    /**
     * Value that indicates whether server-side batched operations are enabled.
     */
    readonly enableBatchedOperations?: pulumi.Input<boolean>;
    /**
     * A value that indicates whether Express Entities are enabled. An express queue holds a message in memory temporarily before writing it to persistent storage.
     */
    readonly enableExpress?: pulumi.Input<boolean>;
    /**
     * A value that indicates whether the queue is to be partitioned across multiple message brokers.
     */
    readonly enablePartitioning?: pulumi.Input<boolean>;
    /**
     * Queue/Topic name to forward the Dead Letter message
     */
    readonly forwardDeadLetteredMessagesTo?: pulumi.Input<string>;
    /**
     * Queue/Topic name to forward the messages
     */
    readonly forwardTo?: pulumi.Input<string>;
    /**
     * ISO 8601 timespan duration of a peek-lock; that is, the amount of time that the message is locked for other receivers. The maximum value for LockDuration is 5 minutes; the default value is 1 minute.
     */
    readonly lockDuration?: pulumi.Input<string>;
    /**
     * The maximum delivery count. A message is automatically deadlettered after this number of deliveries. default value is 10.
     */
    readonly maxDeliveryCount?: pulumi.Input<number>;
    /**
     * The maximum size of the queue in megabytes, which is the size of memory allocated for the queue. Default is 1024.
     */
    readonly maxSizeInMegabytes?: pulumi.Input<number>;
    /**
     * The namespace name
     */
    readonly namespaceName: pulumi.Input<string>;
    /**
     * The queue name.
     */
    readonly queueName: pulumi.Input<string>;
    /**
     * A value indicating if this queue requires duplicate detection.
     */
    readonly requiresDuplicateDetection?: pulumi.Input<boolean>;
    /**
     * A value that indicates whether the queue supports the concept of sessions.
     */
    readonly requiresSession?: pulumi.Input<boolean>;
    /**
     * Name of the Resource group within the Azure subscription.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Enumerates the possible values for the status of a messaging entity.
     */
    readonly status?: pulumi.Input<string>;
}
