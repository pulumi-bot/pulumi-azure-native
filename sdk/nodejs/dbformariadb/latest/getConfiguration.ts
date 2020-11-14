// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

export function getConfiguration(args: GetConfigurationArgs, opts?: pulumi.InvokeOptions): Promise<GetConfigurationResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:dbformariadb/latest:getConfiguration", {
        "configurationName": args.configurationName,
        "resourceGroupName": args.resourceGroupName,
        "serverName": args.serverName,
    }, opts);
}

export interface GetConfigurationArgs {
    /**
     * The name of the server configuration.
     */
    readonly configurationName: string;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    readonly resourceGroupName: string;
    /**
     * The name of the server.
     */
    readonly serverName: string;
}

/**
 * Represents a Configuration.
 */
export interface GetConfigurationResult {
    /**
     * Allowed values of the configuration.
     */
    readonly allowedValues: string;
    /**
     * Data type of the configuration.
     */
    readonly dataType: string;
    /**
     * Default value of the configuration.
     */
    readonly defaultValue: string;
    /**
     * Description of the configuration.
     */
    readonly description: string;
    /**
     * The name of the resource
     */
    readonly name: string;
    /**
     * Source of the configuration.
     */
    readonly source?: string;
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    readonly type: string;
    /**
     * Value of the configuration.
     */
    readonly value?: string;
}
