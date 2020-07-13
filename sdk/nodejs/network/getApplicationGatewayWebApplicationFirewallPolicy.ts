// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getApplicationGatewayWebApplicationFirewallPolicy(args: GetApplicationGatewayWebApplicationFirewallPolicyArgs, opts?: pulumi.InvokeOptions): Promise<GetApplicationGatewayWebApplicationFirewallPolicyResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:network:getApplicationGatewayWebApplicationFirewallPolicy", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetApplicationGatewayWebApplicationFirewallPolicyArgs {
    /**
     * The name of the policy.
     */
    readonly name: string;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: string;
}

/**
 * Defines web application firewall policy.
 */
export interface GetApplicationGatewayWebApplicationFirewallPolicyResult {
    /**
     * A unique read-only string that changes whenever the resource is updated.
     */
    readonly etag: string;
    /**
     * Resource location.
     */
    readonly location?: string;
    /**
     * Resource name.
     */
    readonly name: string;
    /**
     * Properties of the web application firewall policy.
     */
    readonly properties: outputs.network.WebApplicationFirewallPolicyPropertiesFormatResponse;
    /**
     * Resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * Resource type.
     */
    readonly type: string;
}