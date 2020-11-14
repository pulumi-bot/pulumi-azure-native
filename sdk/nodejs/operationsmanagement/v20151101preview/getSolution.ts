// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

export function getSolution(args: GetSolutionArgs, opts?: pulumi.InvokeOptions): Promise<GetSolutionResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:operationsmanagement/v20151101preview:getSolution", {
        "resourceGroupName": args.resourceGroupName,
        "solutionName": args.solutionName,
    }, opts);
}

export interface GetSolutionArgs {
    /**
     * The name of the resource group to get. The name is case insensitive.
     */
    readonly resourceGroupName: string;
    /**
     * User Solution Name.
     */
    readonly solutionName: string;
}

/**
 * The container for solution.
 */
export interface GetSolutionResult {
    /**
     * Resource location
     */
    readonly location?: string;
    /**
     * Resource name.
     */
    readonly name: string;
    /**
     * Plan for solution object supported by the OperationsManagement resource provider.
     */
    readonly plan?: outputs.operationsmanagement.v20151101preview.SolutionPlanResponse;
    /**
     * Properties for solution object supported by the OperationsManagement resource provider.
     */
    readonly properties: outputs.operationsmanagement.v20151101preview.SolutionPropertiesResponse;
    /**
     * Resource tags
     */
    readonly tags?: {[key: string]: string};
    /**
     * Resource type.
     */
    readonly type: string;
}
