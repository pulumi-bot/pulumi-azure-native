// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * The workflow type.
 *
 * ## Example Usage
 * ### Create or update a workflow
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const workflow = new azurerm.logic.v20160601.Workflow("workflow", {
 *     definition: {
 *         `$schema`: "https://schema.management.azure.com/providers/Microsoft.Logic/schemas/2016-06-01/workflowdefinition.json#",
 *         actions: {
 *             Find_pet_by_ID: {
 *                 inputs: {
 *                     host: {
 *                         connection: {
 *                             name: `@parameters('$connections')['test-custom-connector']['connectionId']`,
 *                         },
 *                     },
 *                     method: "get",
 *                     path: "/pet/@{encodeURIComponent('1')}",
 *                 },
 *                 runAfter: {},
 *                 type: "ApiConnection",
 *             },
 *         },
 *         contentVersion: "1.0.0.0",
 *         outputs: {},
 *         parameters: {
 *             `$connections`: {
 *                 defaultValue: {},
 *                 type: "Object",
 *             },
 *         },
 *         triggers: {
 *             manual: {
 *                 inputs: {
 *                     schema: {},
 *                 },
 *                 kind: "Http",
 *                 type: "Request",
 *             },
 *         },
 *     },
 *     location: "brazilsouth",
 *     parameters: {
 *         `$connections`: {
 *             value: {
 *                 "test-custom-connector": {
 *                     connectionId: "/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/test-resource-group/providers/Microsoft.Web/connections/test-custom-connector",
 *                     connectionName: "test-custom-connector",
 *                     id: "/subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/providers/Microsoft.Web/locations/brazilsouth/managedApis/test-custom-connector",
 *                 },
 *             },
 *         },
 *     },
 *     resourceGroupName: "test-resource-group",
 *     tags: {},
 *     workflowName: "test-workflow",
 * });
 *
 * ```
 */
export class Workflow extends pulumi.CustomResource {
    /**
     * Get an existing Workflow resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Workflow {
        return new Workflow(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:logic/v20160601:Workflow';

    /**
     * Returns true if the given object is an instance of Workflow.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Workflow {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Workflow.__pulumiType;
    }

    /**
     * Gets the access endpoint.
     */
    public /*out*/ readonly accessEndpoint!: pulumi.Output<string>;
    /**
     * Gets the changed time.
     */
    public /*out*/ readonly changedTime!: pulumi.Output<string>;
    /**
     * Gets the created time.
     */
    public /*out*/ readonly createdTime!: pulumi.Output<string>;
    /**
     * The definition. See [Schema reference for Workflow Definition Language in Azure Logic Apps](https://aka.ms/logic-apps-workflow-definition-language).
     */
    public readonly definition!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * The integration account.
     */
    public /*out*/ readonly integrationAccount!: pulumi.Output<outputs.logic.v20160601.ResourceReferenceResponse | undefined>;
    /**
     * The resource location.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * Gets the resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The parameters.
     */
    public readonly parameters!: pulumi.Output<{[key: string]: outputs.logic.v20160601.WorkflowParameterResponse} | undefined>;
    /**
     * Gets the provisioning state.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * The sku.
     */
    public readonly sku!: pulumi.Output<outputs.logic.v20160601.SkuResponse | undefined>;
    /**
     * The state.
     */
    public readonly state!: pulumi.Output<string | undefined>;
    /**
     * The resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Gets the resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * Gets the version.
     */
    public /*out*/ readonly version!: pulumi.Output<string>;

    /**
     * Create a Workflow resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WorkflowArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: WorkflowArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as WorkflowArgs | undefined;
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.workflowName === undefined) {
                throw new Error("Missing required property 'workflowName'");
            }
            inputs["definition"] = args ? args.definition : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["parameters"] = args ? args.parameters : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["sku"] = args ? args.sku : undefined;
            inputs["state"] = args ? args.state : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["workflowName"] = args ? args.workflowName : undefined;
            inputs["accessEndpoint"] = undefined /*out*/;
            inputs["changedTime"] = undefined /*out*/;
            inputs["createdTime"] = undefined /*out*/;
            inputs["integrationAccount"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["version"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:logic/latest:Workflow" }, { type: "azurerm:logic/v20150201preview:Workflow" }, { type: "azurerm:logic/v20180701preview:Workflow" }, { type: "azurerm:logic/v20190501:Workflow" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(Workflow.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Workflow resource.
 */
export interface WorkflowArgs {
    /**
     * The definition. See [Schema reference for Workflow Definition Language in Azure Logic Apps](https://aka.ms/logic-apps-workflow-definition-language).
     */
    readonly definition?: pulumi.Input<{[key: string]: any}>;
    /**
     * The resource location.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The parameters.
     */
    readonly parameters?: pulumi.Input<{[key: string]: pulumi.Input<inputs.logic.v20160601.WorkflowParameter>}>;
    /**
     * The resource group name.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The sku.
     */
    readonly sku?: pulumi.Input<inputs.logic.v20160601.Sku>;
    /**
     * The state.
     */
    readonly state?: pulumi.Input<string>;
    /**
     * The resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The workflow name.
     */
    readonly workflowName: pulumi.Input<string>;
}
