// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

export function getEventSubscription(args: GetEventSubscriptionArgs, opts?: pulumi.InvokeOptions): Promise<GetEventSubscriptionResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:eventgrid/v20170915preview:getEventSubscription", {
        "eventSubscriptionName": args.eventSubscriptionName,
        "scope": args.scope,
    }, opts);
}

export interface GetEventSubscriptionArgs {
    /**
     * Name of the event subscription
     */
    readonly eventSubscriptionName: string;
    /**
     * The scope of the event subscription. The scope can be a subscription, or a resource group, or a top level resource belonging to a resource provider namespace, or an EventGrid topic. For example, use '/subscriptions/{subscriptionId}/' for a subscription, '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for a resource group, and '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}' for a resource, and '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/topics/{topicName}' for an EventGrid topic.
     */
    readonly scope: string;
}

/**
 * Event Subscription
 */
export interface GetEventSubscriptionResult {
    /**
     * Information about the destination where events have to be delivered for the event subscription.
     */
    readonly destination?: outputs.eventgrid.v20170915preview.EventHubEventSubscriptionDestinationResponse | outputs.eventgrid.v20170915preview.WebHookEventSubscriptionDestinationResponse;
    /**
     * Information about the filter for the event subscription.
     */
    readonly filter?: outputs.eventgrid.v20170915preview.EventSubscriptionFilterResponse;
    /**
     * List of user defined labels.
     */
    readonly labels?: string[];
    /**
     * Name of the resource
     */
    readonly name: string;
    /**
     * Provisioning state of the event subscription.
     */
    readonly provisioningState: string;
    /**
     * Name of the topic of the event subscription.
     */
    readonly topic: string;
    /**
     * Type of the resource
     */
    readonly type: string;
}
