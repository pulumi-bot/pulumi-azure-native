// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

export function getPolicyDefinition(args: GetPolicyDefinitionArgs, opts?: pulumi.InvokeOptions): Promise<GetPolicyDefinitionResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:authorization/v20160401:getPolicyDefinition", {
        "policyDefinitionName": args.policyDefinitionName,
    }, opts);
}

export interface GetPolicyDefinitionArgs {
    /**
     * The name of the policy definition to get.
     */
    readonly policyDefinitionName: string;
}

/**
 * The policy definition.
 */
export interface GetPolicyDefinitionResult {
    /**
     * The policy definition description.
     */
    readonly description?: string;
    /**
     * The display name of the policy definition.
     */
    readonly displayName?: string;
    /**
     * The name of the policy definition. If you do not specify a value for name, the value is inferred from the name value in the request URI.
     */
    readonly name?: string;
    /**
     * The policy rule.
     */
    readonly policyRule?: any;
    /**
     * The type of policy definition. Possible values are NotSpecified, BuiltIn, and Custom.
     */
    readonly policyType?: string;
}
