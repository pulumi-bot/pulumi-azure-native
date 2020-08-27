// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getNamespaceAuthorizationRule(args: GetNamespaceAuthorizationRuleArgs, opts?: pulumi.InvokeOptions): Promise<GetNamespaceAuthorizationRuleResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:notificationhubs/v20170401:getNamespaceAuthorizationRule", {
        "authorizationRuleName": args.authorizationRuleName,
        "namespaceName": args.namespaceName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetNamespaceAuthorizationRuleArgs {
    /**
     * Authorization rule name.
     */
    readonly authorizationRuleName: string;
    /**
     * The namespace name
     */
    readonly namespaceName: string;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: string;
}

/**
 * Description of a Namespace AuthorizationRules.
 */
export interface GetNamespaceAuthorizationRuleResult {
    /**
     * A string that describes the claim type
     */
    readonly claimType: string;
    /**
     * A string that describes the claim value
     */
    readonly claimValue: string;
    /**
     * The created time for this rule
     */
    readonly createdTime: string;
    /**
     * A string that describes the authorization rule.
     */
    readonly keyName: string;
    /**
     * Resource location
     */
    readonly location?: string;
    /**
     * The last modified time for this rule
     */
    readonly modifiedTime: string;
    /**
     * Resource name
     */
    readonly name: string;
    /**
     * A base64-encoded 256-bit primary key for signing and validating the SAS token.
     */
    readonly primaryKey: string;
    /**
     * The revision number for the rule
     */
    readonly revision: number;
    /**
     * The rights associated with the rule.
     */
    readonly rights?: AccessRights[];
    /**
     * A base64-encoded 256-bit primary key for signing and validating the SAS token.
     */
    readonly secondaryKey: string;
    /**
     * The sku of the created namespace
     */
    readonly sku?: outputs.notificationhubs.v20170401.SkuResponse;
    /**
     * Resource tags
     */
    readonly tags?: {[key: string]: string};
    /**
     * Resource type
     */
    readonly type: string;
}
