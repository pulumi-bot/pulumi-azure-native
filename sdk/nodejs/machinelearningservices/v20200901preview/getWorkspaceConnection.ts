// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

export function getWorkspaceConnection(args: GetWorkspaceConnectionArgs, opts?: pulumi.InvokeOptions): Promise<GetWorkspaceConnectionResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:machinelearningservices/v20200901preview:getWorkspaceConnection", {
        "connectionName": args.connectionName,
        "resourceGroupName": args.resourceGroupName,
        "workspaceName": args.workspaceName,
    }, opts);
}

export interface GetWorkspaceConnectionArgs {
    /**
     * Friendly name of the workspace connection
     */
    readonly connectionName: string;
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
 * Workspace connection.
 */
export interface GetWorkspaceConnectionResult {
    /**
     * Authorization type of the workspace connection.
     */
    readonly authType?: string;
    /**
     * Category of the workspace connection.
     */
    readonly category?: string;
    /**
     * Friendly name of the workspace connection.
     */
    readonly name: string;
    /**
     * Target of the workspace connection.
     */
    readonly target?: string;
    /**
     * Resource type of workspace connection.
     */
    readonly type: string;
    /**
     * Value details of the workspace connection.
     */
    readonly value?: string;
}
