// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getTopicAuthorizationRule(args: GetTopicAuthorizationRuleArgs, opts?: pulumi.InvokeOptions): Promise<GetTopicAuthorizationRuleResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:servicebus/v20140901:getTopicAuthorizationRule", {
        "authorizationRuleName": args.authorizationRuleName,
        "namespaceName": args.namespaceName,
        "resourceGroupName": args.resourceGroupName,
        "topicName": args.topicName,
    }, opts);
}

export interface GetTopicAuthorizationRuleArgs {
    /**
     * The authorization rule name.
     */
    readonly authorizationRuleName: string;
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
 * Description of a namespace authorization rule.
 */
export interface GetTopicAuthorizationRuleResult {
    /**
     * Resource location.
     */
    readonly location?: string;
    /**
     * Resource name
     */
    readonly name: string;
    /**
     * The rights associated with the rule.
     */
    readonly rights: AccessRights[];
    /**
     * Resource type
     */
    readonly type: string;
}
