// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getEventHub(args: GetEventHubArgs, opts?: pulumi.InvokeOptions): Promise<GetEventHubResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:eventhub/v20150801:getEventHub", {
        "eventHubName": args.eventHubName,
        "namespaceName": args.namespaceName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetEventHubArgs {
    /**
     * The Event Hub name
     */
    readonly eventHubName: string;
    /**
     * The Namespace name
     */
    readonly namespaceName: string;
    /**
     * Name of the resource group within the azure subscription.
     */
    readonly resourceGroupName: string;
}

/**
 * Single item in List or Get Event Hub operation
 */
export interface GetEventHubResult {
    /**
     * Exact time the Event Hub was created.
     */
    readonly createdAt: string;
    /**
     * Resource location
     */
    readonly location?: string;
    /**
     * Number of days to retain the events for this Event Hub.
     */
    readonly messageRetentionInDays?: number;
    /**
     * Resource name
     */
    readonly name: string;
    /**
     * Number of partitions created for the Event Hub.
     */
    readonly partitionCount?: number;
    /**
     * Current number of shards on the Event Hub.
     */
    readonly partitionIds: string[];
    /**
     * Enumerates the possible values for the status of the Event Hub.
     */
    readonly status?: EntityStatus;
    /**
     * Resource type
     */
    readonly type: string;
    /**
     * The exact time the message was updated.
     */
    readonly updatedAt: string;
}
