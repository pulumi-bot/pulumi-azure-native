// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getWorkspaceStorageInsightConfig(args: GetWorkspaceStorageInsightConfigArgs, opts?: pulumi.InvokeOptions): Promise<GetWorkspaceStorageInsightConfigResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:operationalinsights:getWorkspaceStorageInsightConfig", {
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
        "workspaceName": args.workspaceName,
    }, opts);
}

export interface GetWorkspaceStorageInsightConfigArgs {
    /**
     * Name of the storageInsightsConfigs resource
     */
    readonly name: string;
    /**
     * The Resource Group name.
     */
    readonly resourceGroupName: string;
    /**
     * The Log Analytics Workspace name.
     */
    readonly workspaceName: string;
}

/**
 * The top level storage insight resource container.
 */
export interface GetWorkspaceStorageInsightConfigResult {
    /**
     * The ETag of the storage insight.
     */
    readonly eTag?: string;
    /**
     * Resource name.
     */
    readonly name: string;
    /**
     * Storage insight properties.
     */
    readonly properties: outputs.operationalinsights.StorageInsightPropertiesResponse;
    /**
     * Resource tags
     */
    readonly tags?: {[key: string]: string};
    /**
     * Resource type.
     */
    readonly type: string;
}