// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

export function getSqlPoolWorkloadGroup(args: GetSqlPoolWorkloadGroupArgs, opts?: pulumi.InvokeOptions): Promise<GetSqlPoolWorkloadGroupResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:synapse/v20201201:getSqlPoolWorkloadGroup", {
        "resourceGroupName": args.resourceGroupName,
        "sqlPoolName": args.sqlPoolName,
        "workloadGroupName": args.workloadGroupName,
        "workspaceName": args.workspaceName,
    }, opts);
}

export interface GetSqlPoolWorkloadGroupArgs {
    /**
     * The name of the resource group. The name is case insensitive.
     */
    readonly resourceGroupName: string;
    /**
     * SQL pool name
     */
    readonly sqlPoolName: string;
    /**
     * The name of the workload group.
     */
    readonly workloadGroupName: string;
    /**
     * The name of the workspace
     */
    readonly workspaceName: string;
}

/**
 * Workload group operations for a sql pool
 */
export interface GetSqlPoolWorkloadGroupResult {
    /**
     * Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
     */
    readonly id: string;
    /**
     * The workload group importance level.
     */
    readonly importance?: string;
    /**
     * The workload group cap percentage resource.
     */
    readonly maxResourcePercent: number;
    /**
     * The workload group request maximum grant percentage.
     */
    readonly maxResourcePercentPerRequest?: number;
    /**
     * The workload group minimum percentage resource.
     */
    readonly minResourcePercent: number;
    /**
     * The workload group request minimum grant percentage.
     */
    readonly minResourcePercentPerRequest: number;
    /**
     * The name of the resource
     */
    readonly name: string;
    /**
     * The workload group query execution timeout.
     */
    readonly queryExecutionTimeout?: number;
    /**
     * The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
     */
    readonly type: string;
}
