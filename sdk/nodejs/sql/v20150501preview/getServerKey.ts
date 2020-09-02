// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getServerKey(args: GetServerKeyArgs, opts?: pulumi.InvokeOptions): Promise<GetServerKeyResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:sql/v20150501preview:getServerKey", {
        "keyName": args.keyName,
        "resourceGroupName": args.resourceGroupName,
        "serverName": args.serverName,
    }, opts);
}

export interface GetServerKeyArgs {
    /**
     * The name of the server key to be retrieved.
     */
    readonly keyName: string;
    /**
     * The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
     */
    readonly resourceGroupName: string;
    /**
     * The name of the server.
     */
    readonly serverName: string;
}

/**
 * A server key.
 */
export interface GetServerKeyResult {
    /**
     * The server key creation date.
     */
    readonly creationDate?: string;
    /**
     * Kind of encryption protector. This is metadata used for the Azure portal experience.
     */
    readonly kind?: string;
    /**
     * Resource location.
     */
    readonly location: string;
    /**
     * Resource name.
     */
    readonly name: string;
    /**
     * The server key type like 'ServiceManaged', 'AzureKeyVault'.
     */
    readonly serverKeyType: string;
    /**
     * Subregion of the server key.
     */
    readonly subregion: string;
    /**
     * Thumbprint of the server key.
     */
    readonly thumbprint?: string;
    /**
     * Resource type.
     */
    readonly type: string;
    /**
     * The URI of the server key.
     */
    readonly uri?: string;
}
