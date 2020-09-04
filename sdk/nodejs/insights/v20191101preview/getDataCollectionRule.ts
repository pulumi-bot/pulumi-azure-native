// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

export function getDataCollectionRule(args: GetDataCollectionRuleArgs, opts?: pulumi.InvokeOptions): Promise<GetDataCollectionRuleResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:insights/v20191101preview:getDataCollectionRule", {
        "dataCollectionRuleName": args.dataCollectionRuleName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetDataCollectionRuleArgs {
    /**
     * The name of the data collection rule. The name is case insensitive.
     */
    readonly dataCollectionRuleName: string;
    /**
     * The name of the resource group. The name is case insensitive.
     */
    readonly resourceGroupName: string;
}

/**
 * Definition of ARM tracked top level resource.
 */
export interface GetDataCollectionRuleResult {
    /**
     * The specification of data flows.
     */
    readonly dataFlows: outputs.insights.v20191101preview.DataFlowResponse[];
    /**
     * The specification of data sources. 
     * This property is optional and can be omitted if the rule is meant to be used via direct calls to the provisioned endpoint.
     */
    readonly dataSources?: outputs.insights.v20191101preview.DataCollectionRuleResponseDataSources;
    /**
     * Description of the data collection rule.
     */
    readonly description?: string;
    /**
     * The specification of destinations.
     */
    readonly destinations: outputs.insights.v20191101preview.DataCollectionRuleResponseDestinations;
    /**
     * Resource entity tag (ETag).
     */
    readonly etag: string;
    /**
     * The geo-location where the resource lives.
     */
    readonly location: string;
    /**
     * The name of the resource.
     */
    readonly name: string;
    /**
     * The resource provisioning state.
     */
    readonly provisioningState: string;
    /**
     * Resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * The type of the resource.
     */
    readonly type: string;
}
