// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * The list Kusto database principals operation response.
 * Latest API Version: 2020-09-18.
 */
/** @deprecated The 'latest' version is deprecated. Please migrate to the function in the top-level module: 'azure-native:kusto:listClusterFollowerDatabases'. */
export function listClusterFollowerDatabases(args: ListClusterFollowerDatabasesArgs, opts?: pulumi.InvokeOptions): Promise<ListClusterFollowerDatabasesResult> {
    pulumi.log.warn("listClusterFollowerDatabases is deprecated: The 'latest' version is deprecated. Please migrate to the function in the top-level module: 'azure-native:kusto:listClusterFollowerDatabases'.")
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-native:kusto/latest:listClusterFollowerDatabases", {
        "clusterName": args.clusterName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface ListClusterFollowerDatabasesArgs {
    /**
     * The name of the Kusto cluster.
     */
    readonly clusterName: string;
    /**
     * The name of the resource group containing the Kusto cluster.
     */
    readonly resourceGroupName: string;
}

/**
 * The list Kusto database principals operation response.
 */
export interface ListClusterFollowerDatabasesResult {
    /**
     * The list of follower database result.
     */
    readonly value?: outputs.kusto.latest.FollowerDatabaseDefinitionResponse[];
}
