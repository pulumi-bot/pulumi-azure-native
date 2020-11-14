// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

export function getPrediction(args: GetPredictionArgs, opts?: pulumi.InvokeOptions): Promise<GetPredictionResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("azure-nextgen:customerinsights/v20170426:getPrediction", {
        "hubName": args.hubName,
        "predictionName": args.predictionName,
        "resourceGroupName": args.resourceGroupName,
    }, opts);
}

export interface GetPredictionArgs {
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
 * The prediction resource format.
 */
export interface GetPredictionResult {
    /**
     * Whether do auto analyze.
     */
    readonly autoAnalyze: boolean;
    /**
     * Description of the prediction.
     */
    readonly description?: {[key: string]: string};
    /**
     * Display name of the prediction.
     */
    readonly displayName?: {[key: string]: string};
    /**
     * The prediction grades.
     */
    readonly grades?: outputs.customerinsights.v20170426.PredictionResponseGrades[];
    /**
     * Interaction types involved in the prediction.
     */
    readonly involvedInteractionTypes?: string[];
    /**
     * KPI types involved in the prediction.
     */
    readonly involvedKpiTypes?: string[];
    /**
     * Relationships involved in the prediction.
     */
    readonly involvedRelationships?: string[];
    /**
     * Definition of the link mapping of prediction.
     */
    readonly mappings: outputs.customerinsights.v20170426.PredictionResponseMappings;
    /**
     * Resource name.
     */
    readonly name: string;
    /**
     * Negative outcome expression.
     */
    readonly negativeOutcomeExpression: string;
    /**
     * Positive outcome expression.
     */
    readonly positiveOutcomeExpression: string;
    /**
     * Name of the prediction.
     */
    readonly predictionName?: string;
    /**
     * Primary profile type.
     */
    readonly primaryProfileType: string;
    /**
     * Provisioning state.
     */
    readonly provisioningState: string;
    /**
     * Scope expression.
     */
    readonly scopeExpression: string;
    /**
     * Score label.
     */
    readonly scoreLabel: string;
    /**
     * System generated entities.
     */
    readonly systemGeneratedEntities: outputs.customerinsights.v20170426.PredictionResponseSystemGeneratedEntities;
    /**
     * The hub name.
     */
    readonly tenantId: string;
    /**
     * Resource type.
     */
    readonly type: string;
}
