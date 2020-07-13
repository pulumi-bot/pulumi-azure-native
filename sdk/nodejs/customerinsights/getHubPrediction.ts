// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getHubPrediction(args: GetHubPredictionArgs, opts?: pulumi.InvokeOptions): Promise<GetHubPredictionResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azurerm:customerinsights:getHubPrediction", {
        "hubName": args.hubName,
        "name": args.name,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetHubPredictionArgs {
    /**
     * The name of the hub.
     */
    readonly hubName: string;
    /**
     * The name of the Prediction.
     */
    readonly name: string;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: string;
}

/**
 * The prediction resource format.
 */
export interface GetHubPredictionResult {
    /**
     * Resource name.
     */
    readonly name: string;
    /**
     * The prediction definition.
     */
    readonly properties: outputs.customerinsights.PredictionResponse;
    /**
     * Resource type.
     */
    readonly type: string;
}