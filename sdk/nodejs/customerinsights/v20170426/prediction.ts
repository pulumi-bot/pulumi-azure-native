// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * The prediction resource format.
 *
 * ## Example Usage
 * ### Predictions_CreateOrUpdate
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const prediction = new azurerm.customerinsights.v20170426.Prediction("prediction", {
 *     autoAnalyze: true,
 *     description: {
 *         "en-us": "sdktest",
 *     },
 *     displayName: {
 *         "en-us": "sdktest",
 *     },
 *     grades: [],
 *     hubName: "sdkTestHub",
 *     involvedInteractionTypes: [],
 *     involvedKpiTypes: [],
 *     involvedRelationships: [],
 *     mappings: {
 *         grade: "sdktest_Grade",
 *         reason: "sdktest_Reason",
 *         score: "sdktest_Score",
 *     },
 *     negativeOutcomeExpression: "Customers.FirstName = 'Mike'",
 *     positiveOutcomeExpression: "Customers.FirstName = 'David'",
 *     predictionName: "sdktest",
 *     primaryProfileType: "Customers",
 *     resourceGroupName: "TestHubRG",
 *     scopeExpression: "*",
 *     scoreLabel: "score label",
 * });
 *
 * ```
 */
export class Prediction extends pulumi.CustomResource {
    /**
     * Get an existing Prediction resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Prediction {
        return new Prediction(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:customerinsights/v20170426:Prediction';

    /**
     * Returns true if the given object is an instance of Prediction.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Prediction {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Prediction.__pulumiType;
    }

    /**
     * Whether do auto analyze.
     */
    public readonly autoAnalyze!: pulumi.Output<boolean>;
    /**
     * Description of the prediction.
     */
    public readonly description!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Display name of the prediction.
     */
    public readonly displayName!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The prediction grades.
     */
    public readonly grades!: pulumi.Output<outputs.customerinsights.v20170426.PredictionResponseGrades[] | undefined>;
    /**
     * Interaction types involved in the prediction.
     */
    public readonly involvedInteractionTypes!: pulumi.Output<string[] | undefined>;
    /**
     * KPI types involved in the prediction.
     */
    public readonly involvedKpiTypes!: pulumi.Output<string[] | undefined>;
    /**
     * Relationships involved in the prediction.
     */
    public readonly involvedRelationships!: pulumi.Output<string[] | undefined>;
    /**
     * Definition of the link mapping of prediction.
     */
    public readonly mappings!: pulumi.Output<outputs.customerinsights.v20170426.PredictionResponseMappings>;
    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Negative outcome expression.
     */
    public readonly negativeOutcomeExpression!: pulumi.Output<string>;
    /**
     * Positive outcome expression.
     */
    public readonly positiveOutcomeExpression!: pulumi.Output<string>;
    /**
     * Name of the prediction.
     */
    public readonly predictionName!: pulumi.Output<string | undefined>;
    /**
     * Primary profile type.
     */
    public readonly primaryProfileType!: pulumi.Output<string>;
    /**
     * Provisioning state.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * Scope expression.
     */
    public readonly scopeExpression!: pulumi.Output<string>;
    /**
     * Score label.
     */
    public readonly scoreLabel!: pulumi.Output<string>;
    /**
     * System generated entities.
     */
    public /*out*/ readonly systemGeneratedEntities!: pulumi.Output<outputs.customerinsights.v20170426.PredictionResponseSystemGeneratedEntities>;
    /**
     * The hub name.
     */
    public /*out*/ readonly tenantId!: pulumi.Output<string>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a Prediction resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PredictionArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: PredictionArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as PredictionArgs | undefined;
            if (!args || args.autoAnalyze === undefined) {
                throw new Error("Missing required property 'autoAnalyze'");
            }
            if (!args || args.hubName === undefined) {
                throw new Error("Missing required property 'hubName'");
            }
            if (!args || args.mappings === undefined) {
                throw new Error("Missing required property 'mappings'");
            }
            if (!args || args.negativeOutcomeExpression === undefined) {
                throw new Error("Missing required property 'negativeOutcomeExpression'");
            }
            if (!args || args.positiveOutcomeExpression === undefined) {
                throw new Error("Missing required property 'positiveOutcomeExpression'");
            }
            if (!args || args.predictionName === undefined) {
                throw new Error("Missing required property 'predictionName'");
            }
            if (!args || args.primaryProfileType === undefined) {
                throw new Error("Missing required property 'primaryProfileType'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.scopeExpression === undefined) {
                throw new Error("Missing required property 'scopeExpression'");
            }
            if (!args || args.scoreLabel === undefined) {
                throw new Error("Missing required property 'scoreLabel'");
            }
            inputs["autoAnalyze"] = args ? args.autoAnalyze : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["grades"] = args ? args.grades : undefined;
            inputs["hubName"] = args ? args.hubName : undefined;
            inputs["involvedInteractionTypes"] = args ? args.involvedInteractionTypes : undefined;
            inputs["involvedKpiTypes"] = args ? args.involvedKpiTypes : undefined;
            inputs["involvedRelationships"] = args ? args.involvedRelationships : undefined;
            inputs["mappings"] = args ? args.mappings : undefined;
            inputs["negativeOutcomeExpression"] = args ? args.negativeOutcomeExpression : undefined;
            inputs["positiveOutcomeExpression"] = args ? args.positiveOutcomeExpression : undefined;
            inputs["predictionName"] = args ? args.predictionName : undefined;
            inputs["primaryProfileType"] = args ? args.primaryProfileType : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["scopeExpression"] = args ? args.scopeExpression : undefined;
            inputs["scoreLabel"] = args ? args.scoreLabel : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["systemGeneratedEntities"] = undefined /*out*/;
            inputs["tenantId"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:customerinsights/latest:Prediction" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(Prediction.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Prediction resource.
 */
export interface PredictionArgs {
    /**
     * Whether do auto analyze.
     */
    readonly autoAnalyze: pulumi.Input<boolean>;
    /**
     * Description of the prediction.
     */
    readonly description?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Display name of the prediction.
     */
    readonly displayName?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The prediction grades.
     */
    readonly grades?: pulumi.Input<pulumi.Input<inputs.customerinsights.v20170426.PredictionGrades>[]>;
    /**
     * The name of the hub.
     */
    readonly hubName: pulumi.Input<string>;
    /**
     * Interaction types involved in the prediction.
     */
    readonly involvedInteractionTypes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * KPI types involved in the prediction.
     */
    readonly involvedKpiTypes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Relationships involved in the prediction.
     */
    readonly involvedRelationships?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Definition of the link mapping of prediction.
     */
    readonly mappings: pulumi.Input<inputs.customerinsights.v20170426.PredictionMappings>;
    /**
     * Negative outcome expression.
     */
    readonly negativeOutcomeExpression: pulumi.Input<string>;
    /**
     * Positive outcome expression.
     */
    readonly positiveOutcomeExpression: pulumi.Input<string>;
    /**
     * Name of the prediction.
     */
    readonly predictionName: pulumi.Input<string>;
    /**
     * Primary profile type.
     */
    readonly primaryProfileType: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Scope expression.
     */
    readonly scopeExpression: pulumi.Input<string>;
    /**
     * Score label.
     */
    readonly scoreLabel: pulumi.Input<string>;
}
