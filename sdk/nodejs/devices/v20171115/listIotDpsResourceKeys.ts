// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

export function listIotDpsResourceKeys(args: ListIotDpsResourceKeysArgs, opts?: pulumi.InvokeOptions): Promise<ListIotDpsResourceKeysResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:devices/v20171115:listIotDpsResourceKeys", {
        "provisioningServiceName": args.provisioningServiceName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface ListIotDpsResourceKeysArgs {
    /**
     * The provisioning service name to get the shared access keys for.
     */
    readonly provisioningServiceName: string;
    /**
     * resource group name
     */
    readonly resourceGroupName: string;
}

/**
 * List of shared access keys.
 */
export interface ListIotDpsResourceKeysResult {
    /**
     * The next link.
     */
    readonly nextLink: string;
    /**
     * The list of shared access policies.
     */
    readonly value?: outputs.devices.v20171115.SharedAccessSignatureAuthorizationRuleAccessRightsDescriptionResponse[];
}
