// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Namespace/NotificationHub Connection String
 * Latest API Version: 2017-04-01.
 */
/** @deprecated The 'latest' version is deprecated. Please migrate to the function in the top-level module: 'azure-native:notificationhubs:listNotificationHubKeys'. */
export function listNotificationHubKeys(args: ListNotificationHubKeysArgs, opts?: pulumi.InvokeOptions): Promise<ListNotificationHubKeysResult> {
    pulumi.log.warn("listNotificationHubKeys is deprecated: The 'latest' version is deprecated. Please migrate to the function in the top-level module: 'azure-native:notificationhubs:listNotificationHubKeys'.")
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-native:notificationhubs/latest:listNotificationHubKeys", {
        "authorizationRuleName": args.authorizationRuleName,
        "namespaceName": args.namespaceName,
        "notificationHubName": args.notificationHubName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface ListNotificationHubKeysArgs {
    /**
     * The connection string of the NotificationHub for the specified authorizationRule.
     */
    readonly authorizationRuleName: string;
    /**
     * The namespace name.
     */
    readonly namespaceName: string;
    /**
     * The notification hub name.
     */
    readonly notificationHubName: string;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: string;
}

/**
 * Namespace/NotificationHub Connection String
 */
export interface ListNotificationHubKeysResult {
    /**
     * KeyName of the created AuthorizationRule
     */
    readonly keyName?: string;
    /**
     * PrimaryConnectionString of the AuthorizationRule.
     */
    readonly primaryConnectionString?: string;
    /**
     * PrimaryKey of the created AuthorizationRule.
     */
    readonly primaryKey?: string;
    /**
     * SecondaryConnectionString of the created AuthorizationRule
     */
    readonly secondaryConnectionString?: string;
    /**
     * SecondaryKey of the created AuthorizationRule
     */
    readonly secondaryKey?: string;
}
