// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

export function getPolicy(args: GetPolicyArgs, opts?: pulumi.InvokeOptions): Promise<GetPolicyResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:network/v20180801:getPolicy", {
        "policyName": args.policyName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetPolicyArgs {
    /**
     * The name of the resource group.
     */
    readonly policyName: string;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: string;
}

/**
 * Defines web application firewall policy.
 */
export interface GetPolicyResult {
    /**
     * Describes custom rules inside the policy
     */
    readonly customRules?: outputs.network.v20180801.CustomRulesResponse;
    /**
     * Gets a unique read-only string that changes whenever the resource is updated.
     */
    readonly etag?: string;
    /**
     * Resource location.
     */
    readonly location?: string;
    /**
     * Describes managed rules inside the policy
     */
    readonly managedRules?: outputs.network.v20180801.ManagedRuleSetsResponse;
    /**
     * Resource name.
     */
    readonly name: string;
    /**
     * Describes  policySettings for policy
     */
    readonly policySettings?: outputs.network.v20180801.PolicySettingsResponse;
    /**
     * Provisioning state of the WebApplicationFirewallPolicy.
     */
    readonly provisioningState: string;
    readonly resourceState: string;
    /**
     * Resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * Resource type.
     */
    readonly type: string;
}
