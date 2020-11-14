// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

export function listTenantAccessGitSecrets(args: ListTenantAccessGitSecretsArgs, opts?: pulumi.InvokeOptions): Promise<ListTenantAccessGitSecretsResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:apimanagement/v20191201:listTenantAccessGitSecrets", {
        "accessName": args.accessName,
        "resourceGroupName": args.resourceGroupName,
        "serviceName": args.serviceName,
    }, opts);
}

export interface ListTenantAccessGitSecretsArgs {
    /**
     * The identifier of the Access configuration.
     */
    readonly accessName: string;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: string;
    /**
     * The name of the API Management service.
     */
    readonly serviceName: string;
}

/**
 * Tenant access information contract of the API Management service.
 */
export interface ListTenantAccessGitSecretsResult {
    /**
     * Determines whether direct access is enabled.
     */
    readonly enabled?: boolean;
    /**
     * Primary access key. This property will not be filled on 'GET' operations! Use '/listSecrets' POST request to get the value.
     */
    readonly primaryKey?: string;
    /**
     * Secondary access key. This property will not be filled on 'GET' operations! Use '/listSecrets' POST request to get the value.
     */
    readonly secondaryKey?: string;
}
