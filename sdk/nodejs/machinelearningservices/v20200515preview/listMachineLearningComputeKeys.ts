// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function listMachineLearningComputeKeys(args: ListMachineLearningComputeKeysArgs, opts?: pulumi.InvokeOptions): Promise<ListMachineLearningComputeKeysResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:machinelearningservices/v20200515preview:listMachineLearningComputeKeys", {
        "computeName": args.computeName,
        "resourceGroupName": args.resourceGroupName,
        "workspaceName": args.workspaceName,
    }, opts);
}

export interface ListMachineLearningComputeKeysArgs {
    /**
     * Name of the Azure Machine Learning compute.
     */
    readonly computeName: string;
    /**
     * Name of the resource group in which workspace is located.
     */
    readonly resourceGroupName: string;
    /**
     * Name of Azure Machine Learning workspace.
     */
    readonly workspaceName: string;
}

/**
 * Secrets related to a Machine Learning compute. Might differ for every type of compute.
 */
export interface ListMachineLearningComputeKeysResult {
    /**
     * The type of compute
     */
    readonly computeType: string;
}
