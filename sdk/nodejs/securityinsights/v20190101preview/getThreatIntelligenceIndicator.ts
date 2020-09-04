// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getThreatIntelligenceIndicator(args: GetThreatIntelligenceIndicatorArgs, opts?: pulumi.InvokeOptions): Promise<GetThreatIntelligenceIndicatorResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:securityinsights/v20190101preview:getThreatIntelligenceIndicator", {
        "name": args.name,
        "operationalInsightsResourceProvider": args.operationalInsightsResourceProvider,
        "resourceGroupName": args.resourceGroupName,
        "workspaceName": args.workspaceName,
    }, opts);
}

export interface GetThreatIntelligenceIndicatorArgs {
    /**
     * Threat Intelligence Identifier
     */
    readonly name: string;
    /**
     * The namespace of workspaces resource provider- Microsoft.OperationalInsights.
     */
    readonly operationalInsightsResourceProvider: string;
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
 * Threat intelligence resource.
 */
export interface GetThreatIntelligenceIndicatorResult {
    /**
     * Etag of the azure resource
     */
    readonly etag?: string;
    /**
     * The kind of the entity.
     */
    readonly kind: string;
    /**
     * Azure resource name
     */
    readonly name: string;
    /**
     * Azure resource type
     */
    readonly type: string;
}
