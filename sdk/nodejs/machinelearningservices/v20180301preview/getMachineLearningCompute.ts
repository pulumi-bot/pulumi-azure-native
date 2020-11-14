// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

export function getMachineLearningCompute(args: GetMachineLearningComputeArgs, opts?: pulumi.InvokeOptions): Promise<GetMachineLearningComputeResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:machinelearningservices/v20180301preview:getMachineLearningCompute", {
        "computeName": args.computeName,
        "resourceGroupName": args.resourceGroupName,
        "workspaceName": args.workspaceName,
    }, opts);
}

export interface GetMachineLearningComputeArgs {
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
 * Machine Learning compute object wrapped into ARM resource envelope.
 */
export interface GetMachineLearningComputeResult {
    /**
     * The identity of the resource.
     */
    readonly identity?: outputs.machinelearningservices.v20180301preview.IdentityResponse;
    /**
     * Specifies the location of the resource.
     */
    readonly location?: string;
    /**
     * Specifies the name of the resource.
     */
    readonly name: string;
    /**
     * Compute properties
     */
    readonly properties: outputs.machinelearningservices.v20180301preview.AKSResponse | outputs.machinelearningservices.v20180301preview.BatchAIResponse | outputs.machinelearningservices.v20180301preview.DataFactoryResponse | outputs.machinelearningservices.v20180301preview.HDInsightResponse | outputs.machinelearningservices.v20180301preview.VirtualMachineResponse;
    /**
     * Contains resource tags defined as key/value pairs.
     */
    readonly tags?: {[key: string]: string};
    /**
     * Specifies the type of the resource.
     */
    readonly type: string;
}
