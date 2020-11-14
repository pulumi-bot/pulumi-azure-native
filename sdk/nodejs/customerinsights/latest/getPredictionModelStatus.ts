// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

export function getPredictionModelStatus(args: GetPredictionModelStatusArgs, opts?: pulumi.InvokeOptions): Promise<GetPredictionModelStatusResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:customerinsights/latest:getPredictionModelStatus", {
        "hubName": args.hubName,
        "predictionName": args.predictionName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetPredictionModelStatusArgs {
    /**
     * The name of the hub.
     */
    readonly hubName: string;
    /**
     * The name of the Prediction.
     */
    readonly predictionName: string;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: string;
}

/**
 * The prediction model status.
 */
export interface GetPredictionModelStatusResult {
    /**
     * The model status message.
     */
    readonly message: string;
    /**
     * Version of the model.
     */
    readonly modelVersion: string;
    /**
     * The prediction GUID ID.
     */
    readonly predictionGuidId: string;
    /**
     * The prediction name.
     */
    readonly predictionName: string;
    /**
     * The signals used.
     */
    readonly signalsUsed: number;
    /**
     * Prediction model life cycle.  When prediction is in PendingModelConfirmation status, it is allowed to update the status to PendingFeaturing or Active through API.
     */
    readonly status: string;
    /**
     * The hub name.
     */
    readonly tenantId: string;
    /**
     * Count of the test set.
     */
    readonly testSetCount: number;
    /**
     * The training accuracy.
     */
    readonly trainingAccuracy: number;
    /**
     * Count of the training set.
     */
    readonly trainingSetCount: number;
    /**
     * Count of the validation set.
     */
    readonly validationSetCount: number;
}
