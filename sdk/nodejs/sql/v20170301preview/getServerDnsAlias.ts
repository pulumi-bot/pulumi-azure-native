// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getServerDnsAlias(args: GetServerDnsAliasArgs, opts?: pulumi.InvokeOptions): Promise<GetServerDnsAliasResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:sql/v20170301preview:getServerDnsAlias", {
        "dnsAliasName": args.dnsAliasName,
        "resourceGroupName": args.resourceGroupName,
        "serverName": args.serverName,
    }, opts);
}

export interface GetServerDnsAliasArgs {
    /**
     * The name of the server DNS alias.
     */
    readonly dnsAliasName: string;
    /**
     * The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
     */
    readonly resourceGroupName: string;
    /**
     * The name of the server that the alias is pointing to.
     */
    readonly serverName: string;
}

/**
 * A server DNS alias.
 */
export interface GetServerDnsAliasResult {
    /**
     * The fully qualified DNS record for alias
     */
    readonly azureDnsRecord: string;
    /**
     * Resource name.
     */
    readonly name: string;
    /**
     * Resource type.
     */
    readonly type: string;
}
