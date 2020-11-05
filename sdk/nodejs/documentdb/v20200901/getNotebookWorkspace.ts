// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getNotebookWorkspace(args: GetNotebookWorkspaceArgs, opts?: pulumi.InvokeOptions): Promise<GetNotebookWorkspaceResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:documentdb/v20200901:getNotebookWorkspace", {
        "accountName": args.accountName,
        "notebookWorkspaceName": args.notebookWorkspaceName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetNotebookWorkspaceArgs {
    /**
     * Cosmos DB database account name.
     */
    readonly accountName: string;
    /**
     * The name of the notebook workspace resource.
     */
    readonly notebookWorkspaceName: string;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    readonly resourceGroupName: string;
}

/**
 * A notebook workspace resource
 */
export interface GetNotebookWorkspaceResult {
    /**
     * The name of the database account.
     */
    readonly name: string;
    /**
     * Specifies the endpoint of Notebook server.
     */
    readonly notebookServerEndpoint: string;
    /**
     * Status of the notebook workspace. Possible values are: Creating, Online, Deleting, Failed, Updating.
     */
    readonly status: string;
    /**
     * The type of Azure resource.
     */
    readonly type: string;
}
