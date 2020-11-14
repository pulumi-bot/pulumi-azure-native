// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

export function getIntegrationAccountBatchConfiguration(args: GetIntegrationAccountBatchConfigurationArgs, opts?: pulumi.InvokeOptions): Promise<GetIntegrationAccountBatchConfigurationResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:logic/v20180701preview:getIntegrationAccountBatchConfiguration", {
        "batchConfigurationName": args.batchConfigurationName,
        "integrationAccountName": args.integrationAccountName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetIntegrationAccountBatchConfigurationArgs {
    /**
     * The batch configuration name.
     */
    readonly batchConfigurationName: string;
    /**
     * The integration account name.
     */
    readonly integrationAccountName: string;
    /**
     * The resource group name.
     */
    readonly resourceGroupName: string;
}

/**
 * The batch configuration resource definition.
 */
export interface GetIntegrationAccountBatchConfigurationResult {
    /**
     * The resource location.
     */
    readonly location?: string;
    /**
     * Gets the resource name.
     */
    readonly name: string;
    /**
     * The batch configuration properties.
     */
    readonly properties: outputs.logic.v20180701preview.BatchConfigurationPropertiesResponse;
    /**
     * The resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * Gets the resource type.
     */
    readonly type: string;
}
