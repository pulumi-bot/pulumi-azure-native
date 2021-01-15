// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

export function getDscNodeConfiguration(args: GetDscNodeConfigurationArgs, opts?: pulumi.InvokeOptions): Promise<GetDscNodeConfigurationResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:automation/v20190601:getDscNodeConfiguration", {
        "automationAccountName": args.automationAccountName,
        "nodeConfigurationName": args.nodeConfigurationName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetDscNodeConfigurationArgs {
    /**
     * The name of the automation account.
     */
    readonly automationAccountName: string;
    /**
     * The Dsc node configuration name.
     */
    readonly nodeConfigurationName: string;
    /**
     * Name of an Azure Resource group.
     */
    readonly resourceGroupName: string;
}

/**
 * Definition of the dsc node configuration.
 */
export interface GetDscNodeConfigurationResult {
    /**
     * Gets or sets the configuration of the node.
     */
    readonly configuration?: outputs.automation.v20190601.DscConfigurationAssociationPropertyResponse;
    /**
     * Gets or sets creation time.
     */
    readonly creationTime?: string;
    /**
     * Fully qualified resource Id for the resource
     */
    readonly id: string;
    /**
     * If a new build version of NodeConfiguration is required.
     */
    readonly incrementNodeConfigurationBuild?: boolean;
    /**
     * Gets or sets the last modified time.
     */
    readonly lastModifiedTime?: string;
    /**
     * The name of the resource
     */
    readonly name: string;
    /**
     * Number of nodes with this node configuration assigned
     */
    readonly nodeCount?: number;
    /**
     * Source of node configuration.
     */
    readonly source?: string;
    /**
     * The type of the resource.
     */
    readonly type: string;
}
