// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getSystemTopic(args: GetSystemTopicArgs, opts?: pulumi.InvokeOptions): Promise<GetSystemTopicResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:eventgrid/v20200401preview:getSystemTopic", {
        "resourceGroupName": args.resourceGroupName,
        "systemTopicName": args.systemTopicName,
    }, opts);
}

export interface GetSystemTopicArgs {
    /**
     * The name of the resource group within the user's subscription.
     */
    readonly resourceGroupName: string;
    /**
     * Name of the system topic.
     */
    readonly systemTopicName: string;
}

/**
 * EventGrid System Topic.
 */
export interface GetSystemTopicResult {
    /**
     * Location of the resource.
     */
    readonly location: string;
    /**
     * Metric resource id for the system topic.
     */
    readonly metricResourceId: string;
    /**
     * Name of the resource
     */
    readonly name: string;
    /**
     * Provisioning state of the system topic.
     */
    readonly provisioningState: string;
    /**
     * Source for the system topic.
     */
    readonly source?: string;
    /**
     * Tags of the resource.
     */
    readonly tags?: {[key: string]: string};
    /**
     * TopicType for the system topic.
     */
    readonly topicType?: string;
    /**
     * Type of the resource
     */
    readonly type: string;
}
