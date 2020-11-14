// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

export function getprivateLinkForAzureAd(args: GetprivateLinkForAzureAdArgs, opts?: pulumi.InvokeOptions): Promise<GetprivateLinkForAzureAdResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:aadiam/v20200301preview:getprivateLinkForAzureAd", {
        "policyName": args.policyName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetprivateLinkForAzureAdArgs {
    /**
     * The name of the private link policy in Azure AD.
     */
    readonly policyName: string;
    /**
     * Name of an Azure resource group.
     */
    readonly resourceGroupName: string;
}

/**
 * PrivateLink Policy configuration object.
 */
export interface GetprivateLinkForAzureAdResult {
    /**
     * Flag indicating whether all tenants are allowed
     */
    readonly allTenants?: boolean;
    /**
     * Name of this resource.
     */
    readonly name?: string;
    /**
     * Guid of the owner tenant
     */
    readonly ownerTenantId?: string;
    /**
     * Name of the resource group
     */
    readonly resourceGroup?: string;
    /**
     * Name of the private link policy resource
     */
    readonly resourceName?: string;
    /**
     * Subscription Identifier
     */
    readonly subscriptionId?: string;
    /**
     * Resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * The list of tenantIds.
     */
    readonly tenants?: string[];
    /**
     * Type of this resource.
     */
    readonly type: string;
}
