// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getBookmarkRelationRelation(args: GetBookmarkRelationRelationArgs, opts?: pulumi.InvokeOptions): Promise<GetBookmarkRelationRelationResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:securityinsights/v20190101preview:getBookmarkRelationRelation", {
        "bookmarkId": args.bookmarkId,
        "operationalInsightsResourceProvider": args.operationalInsightsResourceProvider,
        "relationName": args.relationName,
        "resourceGroupName": args.resourceGroupName,
        "workspaceName": args.workspaceName,
    }, opts);
}

export interface GetBookmarkRelationRelationArgs {
    /**
     * Bookmark ID
     */
    readonly bookmarkId: string;
    /**
     * The namespace of workspaces resource provider- Microsoft.OperationalInsights.
     */
    readonly operationalInsightsResourceProvider: string;
    /**
     * Relation Name
     */
    readonly relationName: string;
    /**
     * The name of the resource group within the user's subscription. The name is case insensitive.
     */
    readonly resourceGroupName: string;
    /**
     * The name of the workspace.
     */
    readonly workspaceName: string;
}

/**
 * Represents a relation between two resources
 */
export interface GetBookmarkRelationRelationResult {
    /**
     * Etag of the azure resource
     */
    readonly etag?: string;
    /**
     * Azure resource name
     */
    readonly name: string;
    /**
     * The resource ID of the related resource
     */
    readonly relatedResourceId: string;
    /**
     * The resource kind of the related resource
     */
    readonly relatedResourceKind: string;
    /**
     * The name of the related resource
     */
    readonly relatedResourceName: string;
    /**
     * The resource type of the related resource
     */
    readonly relatedResourceType: string;
    /**
     * Azure resource type
     */
    readonly type: string;
}
