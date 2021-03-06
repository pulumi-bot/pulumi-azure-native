// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * The notification registration definition.
 * Latest API Version: 2020-11-20.
 */
/** @deprecated The 'latest' version is deprecated. Please migrate to the function in the top-level module: 'azure-native:providerhub:getNotificationRegistration'. */
export function getNotificationRegistration(args: GetNotificationRegistrationArgs, opts?: pulumi.InvokeOptions): Promise<GetNotificationRegistrationResult> {
    pulumi.log.warn("getNotificationRegistration is deprecated: The 'latest' version is deprecated. Please migrate to the function in the top-level module: 'azure-native:providerhub:getNotificationRegistration'.")
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-native:providerhub/latest:getNotificationRegistration", {
        "notificationRegistrationName": args.notificationRegistrationName,
        "providerNamespace": args.providerNamespace,
    }, opts);
}

export interface GetNotificationRegistrationArgs {
    /**
     * The notification registration.
     */
    readonly notificationRegistrationName: string;
    /**
     * The name of the resource provider hosted within ProviderHub.
     */
    readonly providerNamespace: string;
}

/**
 * The notification registration definition.
 */
export interface GetNotificationRegistrationResult {
    /**
     * Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
     */
    readonly id: string;
    /**
     * The name of the resource
     */
    readonly name: string;
    readonly properties: outputs.providerhub.latest.NotificationRegistrationResponseProperties;
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    readonly type: string;
}
