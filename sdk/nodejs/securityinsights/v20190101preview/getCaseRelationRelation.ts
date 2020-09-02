// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getCaseRelationRelation(args: GetCaseRelationRelationArgs, opts?: pulumi.InvokeOptions): Promise<GetCaseRelationRelationResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:securityinsights/v20190101preview:getCaseRelationRelation", {
        "caseId": args.caseId,
        "operationalInsightsResourceProvider": args.operationalInsightsResourceProvider,
        "relationName": args.relationName,
        "resourceGroupName": args.resourceGroupName,
        "workspaceName": args.workspaceName,
    }, opts);
}

export interface GetCaseRelationRelationArgs {
    /**
     * Case ID
     */
    readonly caseId: string;
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
 * Represents a case relation
 */
export interface GetCaseRelationRelationResult {
    /**
     * The case related bookmark id
     */
    readonly bookmarkId: string;
    /**
     * The case related bookmark name
     */
    readonly bookmarkName?: string;
    /**
     * The case identifier
     */
    readonly caseIdentifier: string;
    /**
     * ETag for relation
     */
    readonly etag?: string;
    /**
     * The type of relation node
     */
    readonly kind: string;
    /**
     * Azure resource name
     */
    readonly name: string;
    /**
     * Name of relation
     */
    readonly relationName: string;
    /**
     * Azure resource type
     */
    readonly type: string;
}
