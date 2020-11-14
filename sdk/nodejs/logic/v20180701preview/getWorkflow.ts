// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

export function getWorkflow(args: GetWorkflowArgs, opts?: pulumi.InvokeOptions): Promise<GetWorkflowResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:logic/v20180701preview:getWorkflow", {
        "resourceGroupName": args.resourceGroupName,
        "workflowName": args.workflowName,
    }, opts);
}

export interface GetWorkflowArgs {
    /**
     * The resource group name.
     */
    readonly resourceGroupName: string;
    /**
     * The workflow name.
     */
    readonly workflowName: string;
}

/**
 * The workflow type.
 */
export interface GetWorkflowResult {
    /**
     * Gets the access endpoint.
     */
    readonly accessEndpoint: string;
    /**
     * Gets the changed time.
     */
    readonly changedTime: string;
    /**
     * Gets the created time.
     */
    readonly createdTime: string;
    /**
     * The definition. See [Schema reference for Workflow Definition Language in Azure Logic Apps](https://aka.ms/logic-apps-workflow-definition-language).
     */
    readonly definition?: any;
    /**
     * The integration account.
     */
    readonly integrationAccount?: outputs.logic.v20180701preview.ResourceReferenceResponse;
    /**
     * The resource location.
     */
    readonly location?: string;
    /**
     * Gets the resource name.
     */
    readonly name: string;
    /**
     * The parameters.
     */
    readonly parameters?: {[key: string]: outputs.logic.v20180701preview.WorkflowParameterResponse};
    /**
     * Gets the provisioning state.
     */
    readonly provisioningState: string;
    /**
     * The sku.
     */
    readonly sku?: outputs.logic.v20180701preview.SkuResponse;
    /**
     * The state.
     */
    readonly state?: string;
    /**
     * The resource tags.
     */
    readonly tags?: {[key: string]: string};
    /**
     * Gets the resource type.
     */
    readonly type: string;
    /**
     * Gets the version.
     */
    readonly version: string;
}
