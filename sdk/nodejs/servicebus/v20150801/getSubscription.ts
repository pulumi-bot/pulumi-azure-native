// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getSubscription(args: GetSubscriptionArgs, opts?: pulumi.InvokeOptions): Promise<GetSubscriptionResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:servicebus/v20150801:getSubscription", {
        "name": args.name,
        "namespaceName": args.namespaceName,
        "resourceGroupName": args.resourceGroupName,
        "topicName": args.topicName,
    }, opts);
}

export interface GetSubscriptionArgs {
    /**
     * The subscription name.
     */
    readonly name: string;
    /**
     * The namespace name
     */
    readonly namespaceName: string;
    /**
     * Name of the Resource group within the Azure subscription.
     */
    readonly resourceGroupName: string;
    /**
     * The topic name.
     */
    readonly topicName: string;
}

/**
 * Description of subscription resource.
 */
export interface GetSubscriptionResult {
    /**
     * Resource location.
     */
    readonly location?: string;
    /**
     * Resource name
     */
    readonly name: string;
    /**
     * Description of Subscription Resource.
     */
    readonly properties: outputs.servicebus.v20150801.SubscriptionPropertiesResponse;
    /**
     * Resource type
     */
    readonly type: string;
}