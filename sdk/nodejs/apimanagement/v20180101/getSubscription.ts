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
    return pulumi.runtime.invoke("azurerm:apimanagement/v20180101:getSubscription", {
        "resourceGroupName": args.resourceGroupName,
        "serviceName": args.serviceName,
        "sid": args.sid,
    }, opts);
}

export interface GetSubscriptionArgs {
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: string;
    /**
     * The name of the API Management service.
     */
    readonly serviceName: string;
    /**
     * Subscription entity Identifier. The entity represents the association between a user and a product in API Management.
     */
    readonly sid: string;
}

/**
 * Subscription details.
 */
export interface GetSubscriptionResult {
    /**
     * Subscription creation date. The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
     */
    readonly createdDate: string;
    /**
     * The name of the subscription, or null if the subscription has no name.
     */
    readonly displayName?: string;
    /**
     * Date when subscription was cancelled or expired. The setting is for audit purposes only and the subscription is not automatically cancelled. The subscription lifecycle can be managed by using the `state` property. The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
     */
    readonly endDate?: string;
    /**
     * Subscription expiration date. The setting is for audit purposes only and the subscription is not automatically expired. The subscription lifecycle can be managed by using the `state` property. The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
     */
    readonly expirationDate?: string;
    /**
     * Resource name.
     */
    readonly name: string;
    /**
     * Upcoming subscription expiration notification date. The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
     */
    readonly notificationDate?: string;
    /**
     * Subscription primary key.
     */
    readonly primaryKey: string;
    /**
     * The product resource identifier of the subscribed product. The value is a valid relative URL in the format of /products/{productId} where {productId} is a product identifier.
     */
    readonly productId: string;
    /**
     * Subscription secondary key.
     */
    readonly secondaryKey: string;
    /**
     * Subscription activation date. The setting is for audit purposes only and the subscription is not automatically activated. The subscription lifecycle can be managed by using the `state` property. The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
     */
    readonly startDate?: string;
    /**
     * Subscription state. Possible states are * active – the subscription is active, * suspended – the subscription is blocked, and the subscriber cannot call any APIs of the product, * submitted – the subscription request has been made by the developer, but has not yet been approved or rejected, * rejected – the subscription request has been denied by an administrator, * cancelled – the subscription has been cancelled by the developer or administrator, * expired – the subscription reached its expiration date and was deactivated.
     */
    readonly state: SubscriptionState;
    /**
     * Optional subscription comment added by an administrator.
     */
    readonly stateComment?: string;
    /**
     * Resource type for API Management resource.
     */
    readonly type: string;
    /**
     * The user resource identifier of the subscription owner. The value is a valid relative URL in the format of /users/{uid} where {uid} is a user identifier.
     */
    readonly userId: string;
}
