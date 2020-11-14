// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

export function getQueue(args: GetQueueArgs, opts?: pulumi.InvokeOptions): Promise<GetQueueResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:servicebus/latest:getQueue", {
        "namespaceName": args.namespaceName,
        "queueName": args.queueName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetQueueArgs {
    /**
     * The namespace name
     */
    readonly namespaceName: string;
    /**
     * The queue name.
     */
    readonly queueName: string;
    /**
     * Name of the Resource group within the Azure subscription.
     */
    readonly resourceGroupName: string;
}

/**
 * Description of queue Resource.
 */
export interface GetQueueResult {
    /**
     * Last time a message was sent, or the last time there was a receive request to this queue.
     */
    readonly accessedAt: string;
    /**
     * ISO 8061 timeSpan idle interval after which the queue is automatically deleted. The minimum duration is 5 minutes.
     */
    readonly autoDeleteOnIdle?: string;
    /**
     * Message Count Details.
     */
    readonly countDetails: outputs.servicebus.latest.MessageCountDetailsResponse;
    /**
     * The exact time the message was created.
     */
    readonly createdAt: string;
    /**
     * A value that indicates whether this queue has dead letter support when a message expires.
     */
    readonly deadLetteringOnMessageExpiration?: boolean;
    /**
     * ISO 8601 default message timespan to live value. This is the duration after which the message expires, starting from when the message is sent to Service Bus. This is the default value used when TimeToLive is not set on a message itself.
     */
    readonly defaultMessageTimeToLive?: string;
    /**
     * ISO 8601 timeSpan structure that defines the duration of the duplicate detection history. The default value is 10 minutes.
     */
    readonly duplicateDetectionHistoryTimeWindow?: string;
    /**
     * Value that indicates whether server-side batched operations are enabled.
     */
    readonly enableBatchedOperations?: boolean;
    /**
     * A value that indicates whether Express Entities are enabled. An express queue holds a message in memory temporarily before writing it to persistent storage.
     */
    readonly enableExpress?: boolean;
    /**
     * A value that indicates whether the queue is to be partitioned across multiple message brokers.
     */
    readonly enablePartitioning?: boolean;
    /**
     * Queue/Topic name to forward the Dead Letter message
     */
    readonly forwardDeadLetteredMessagesTo?: string;
    /**
     * Queue/Topic name to forward the messages
     */
    readonly forwardTo?: string;
    /**
     * ISO 8601 timespan duration of a peek-lock; that is, the amount of time that the message is locked for other receivers. The maximum value for LockDuration is 5 minutes; the default value is 1 minute.
     */
    readonly lockDuration?: string;
    /**
     * The maximum delivery count. A message is automatically deadlettered after this number of deliveries. default value is 10.
     */
    readonly maxDeliveryCount?: number;
    /**
     * The maximum size of the queue in megabytes, which is the size of memory allocated for the queue. Default is 1024.
     */
    readonly maxSizeInMegabytes?: number;
    /**
     * The number of messages in the queue.
     */
    readonly messageCount: number;
    /**
     * Resource name
     */
    readonly name: string;
    /**
     * A value indicating if this queue requires duplicate detection.
     */
    readonly requiresDuplicateDetection?: boolean;
    /**
     * A value that indicates whether the queue supports the concept of sessions.
     */
    readonly requiresSession?: boolean;
    /**
     * The size of the queue, in bytes.
     */
    readonly sizeInBytes: number;
    /**
     * Enumerates the possible values for the status of a messaging entity.
     */
    readonly status?: string;
    /**
     * Resource type
     */
    readonly type: string;
    /**
     * The exact time the message was updated.
     */
    readonly updatedAt: string;
}
