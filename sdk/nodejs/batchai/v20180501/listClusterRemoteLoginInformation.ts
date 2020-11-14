// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

export function listClusterRemoteLoginInformation(args: ListClusterRemoteLoginInformationArgs, opts?: pulumi.InvokeOptions): Promise<ListClusterRemoteLoginInformationResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:batchai/v20180501:listClusterRemoteLoginInformation", {
        "clusterName": args.clusterName,
        "resourceGroupName": args.resourceGroupName,
        "workspaceName": args.workspaceName,
    }, opts);
}

export interface ListClusterRemoteLoginInformationArgs {
    /**
     * The name of the cluster within the specified resource group. Cluster names can only contain a combination of alphanumeric characters along with dash (-) and underscore (_). The name must be from 1 through 64 characters long.
     */
    readonly clusterName: string;
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
export interface ListClusterRemoteLoginInformationResult {
    /**
     * The continuation token.
     */
    readonly nextLink: string;
    /**
     * The collection of returned remote login details.
     */
    readonly value: outputs.batchai.v20180501.RemoteLoginInformationResponse[];
}
