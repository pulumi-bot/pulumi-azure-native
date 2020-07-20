// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function listWorkspaceExperimentJobRemoteLoginInformation(args: ListWorkspaceExperimentJobRemoteLoginInformationArgs, opts?: pulumi.InvokeOptions): Promise<ListWorkspaceExperimentJobRemoteLoginInformationResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:batchai:listWorkspaceExperimentJobRemoteLoginInformation", {
        "experimentName": args.experimentName,
        "jobName": args.jobName,
        "resourceGroupName": args.resourceGroupName,
        "workspaceName": args.workspaceName,
    }, opts);
}

export interface ListWorkspaceExperimentJobRemoteLoginInformationArgs {
    /**
     * The name of the experiment. Experiment names can only contain a combination of alphanumeric characters along with dash (-) and underscore (_). The name must be from 1 through 64 characters long.
     */
    readonly experimentName: string;
    /**
     * The name of the job within the specified resource group. Job names can only contain a combination of alphanumeric characters along with dash (-) and underscore (_). The name must be from 1 through 64 characters long.
     */
    readonly jobName: string;
    /**
     * Name of the resource group to which the resource belongs.
     */
    readonly resourceGroupName: string;
    /**
     * The name of the workspace. Workspace names can only contain a combination of alphanumeric characters along with dash (-) and underscore (_). The name must be from 1 through 64 characters long.
     */
    readonly workspaceName: string;
}

/**
 * Values returned by the List operation.
 */
export interface ListWorkspaceExperimentJobRemoteLoginInformationResult {
    /**
     * The continuation token.
     */
    readonly nextLink: string;
    /**
     * The collection of returned remote login details.
     */
    readonly value: outputs.batchai.RemoteLoginInformationResponse[];
}