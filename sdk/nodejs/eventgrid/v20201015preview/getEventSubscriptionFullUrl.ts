// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

export function getEventSubscriptionFullUrl(args: GetEventSubscriptionFullUrlArgs, opts?: pulumi.InvokeOptions): Promise<GetEventSubscriptionFullUrlResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:eventgrid/v20201015preview:getEventSubscriptionFullUrl", {
        "eventSubscriptionName": args.eventSubscriptionName,
        "scope": args.scope,
    }, opts);
}

export interface GetEventSubscriptionFullUrlArgs {
    /**
     * Name of the event subscription.
     */
    readonly eventSubscriptionName: string;
    /**
     * The scope of the event subscription. The scope can be a subscription, or a resource group, or a top level resource belonging to a resource provider namespace, or an EventGrid topic. For example, use '/subscriptions/{subscriptionId}/' for a subscription, '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for a resource group, and '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}' for a resource, and '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/topics/{topicName}' for an EventGrid topic.
     */
    readonly scope: string;
}

/**
 * Full endpoint url of an event subscription
 */
export interface GetEventSubscriptionFullUrlResult {
    /**
     * The URL that represents the endpoint of the destination of an event subscription.
     */
    readonly endpointUrl?: string;
}
