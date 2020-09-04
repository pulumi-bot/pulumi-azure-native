// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getRemoteRenderingAccount(args: GetRemoteRenderingAccountArgs, opts?: pulumi.InvokeOptions): Promise<GetRemoteRenderingAccountResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:mixedreality/v20200406preview:getRemoteRenderingAccount", {
        "accountName": args.accountName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetRemoteRenderingAccountArgs {
    /**
     * Name of an Mixed Reality Account.
     */
    readonly accountName: string;
    /**
     * Name of an Azure resource group.
     */
    readonly resourceGroupName: string;
}

/**
 * RemoteRenderingAccount Response.
 */
export interface GetRemoteRenderingAccountResult {
    /**
     * Correspond domain name of certain Spatial Anchors Account
     */
    readonly accountDomain: string;
    /**
     * unique id of certain account.
     */
    readonly accountId: string;
    readonly identity?: outputs.mixedreality.v20200406preview.RemoteRenderingAccountResponseIdentity;
    /**
     * The geo-location where the resource lives
     */
    readonly location: string;
    /**
     * The name of the resource
     */
    readonly name: string;
    /**
     * Resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
     */
    readonly type: string;
}
