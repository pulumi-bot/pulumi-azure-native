// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getSavedSearch(args: GetSavedSearchArgs, opts?: pulumi.InvokeOptions): Promise<GetSavedSearchResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:operationalinsights/v20200301preview:getSavedSearch", {
        "resourceGroupName": args.resourceGroupName,
        "savedSearchId": args.savedSearchId,
        "workspaceName": args.workspaceName,
    }, opts);
}

export interface GetSavedSearchArgs {
    /**
     * The name of the resource group. The name is case insensitive.
     */
    readonly resourceGroupName: string;
    /**
     * The id of the saved search.
     */
    readonly savedSearchId: string;
    /**
     * The name of the workspace.
     */
    readonly workspaceName: string;
}

/**
 * Value object for saved search results.
 */
export interface GetSavedSearchResult {
    /**
     * The category of the saved search. This helps the user to find a saved search faster. 
     */
    readonly category: string;
    /**
     * Saved search display name.
     */
    readonly displayName: string;
    /**
     * The ETag of the saved search.
     */
    readonly etag?: string;
    /**
     * The function alias if query serves as a function.
     */
    readonly functionAlias?: string;
    /**
     * The optional function parameters if query serves as a function. Value should be in the following format: 'param-name1:type1 = default_value1, param-name2:type2 = default_value2'. For more examples and proper syntax please refer to https://docs.microsoft.com/en-us/azure/kusto/query/functions/user-defined-functions.
     */
    readonly functionParameters?: string;
    /**
     * The name of the resource
     */
    readonly name: string;
    /**
     * The query expression for the saved search.
     */
    readonly query: string;
    /**
     * The tags attached to the saved search.
     */
    readonly tags?: outputs.operationalinsights.v20200301preview.TagResponse[];
    /**
     * The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
     */
    readonly type: string;
    /**
     * The version number of the query language. The current version is 2 and is the default.
     */
    readonly version?: number;
}
