// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

export function listNamespaceKeys(args: ListNamespaceKeysArgs, opts?: pulumi.InvokeOptions): Promise<ListNamespaceKeysResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:relay/v20160701:listNamespaceKeys", {
        "authorizationRuleName": args.authorizationRuleName,
        "namespaceName": args.namespaceName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface ListNamespaceKeysArgs {
    /**
     * The authorizationRule name.
     */
    readonly authorizationRuleName: string;
    /**
     * The Namespace Name
     */
    readonly namespaceName: string;
    /**
     * Name of the Resource group within the Azure subscription.
     */
    readonly resourceGroupName: string;
}

/**
 * Namespace/Relay Connection String
 */
export interface ListNamespaceKeysResult {
    /**
     * A string that describes the authorization rule
     */
    readonly keyName?: string;
    /**
     * PrimaryConnectionString of the created Namespace AuthorizationRule.
     */
    readonly primaryConnectionString?: string;
    /**
     * A base64-encoded 256-bit primary key for signing and validating the SAS token
     */
    readonly primaryKey?: string;
    /**
     * SecondaryConnectionString of the created Namespace AuthorizationRule
     */
    readonly secondaryConnectionString?: string;
    /**
     * A base64-encoded 256-bit secondary key for signing and validating the SAS token
     */
    readonly secondaryKey?: string;
}
