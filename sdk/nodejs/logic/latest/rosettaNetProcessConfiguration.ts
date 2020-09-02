// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * The integration account RosettaNet process configuration.
 *
 * ## Create or update an RosettaNetProcessConfiguration
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const rosettaNetProcessConfiguration = new azurerm.logic.latest.RosettaNetProcessConfiguration("rosettaNetProcessConfiguration", {
 *     activitySettings: {
 *         acknowledgmentOfReceiptSettings: {
 *             isNonRepudiationRequired: false,
 *             timeToAcknowledgeInSeconds: 600,
 *         },
 *         activityBehavior: {
 *             actionType: "DoubleAction",
 *             isAuthorizationRequired: false,
 *             isSecuredTransportRequired: false,
 *             nonRepudiationOfOriginAndContent: false,
 *             persistentConfidentialityScope: "None",
 *             responseType: "Async",
 *             retryCount: 2,
 *             timeToPerformInSeconds: 7200,
 *         },
 *         activityType: "RequestResponse",
 *     },
 *     description: "Test description",
 *     initiatorRoleSettings: {
 *         action: "Purchase Order Request",
 *         businessDocument: {
 *             description: "A request to accept a purchase order for fulfillment..",
 *             name: "Purchase Order Request",
 *             version: "V02.02.00",
 *         },
 *         description: "This partner role creates a demand for a product or service.",
 *         role: "Buyer",
 *         roleType: "Functional",
 *         service: "Buyer Service",
 *         serviceClassification: "Business Service",
 *     },
 *     integrationAccountName: "testia123",
 *     processCode: "3A4",
 *     processName: "Request Purchase Order",
 *     processVersion: "V02.02.00",
 *     resourceGroupName: "testrg123",
 *     responderRoleSettings: {
 *         action: "Purchase Order Confirmation Action",
 *         businessDocument: {
 *             description: "Formally confirms the status of line item(s) in a Purchase Order. A Purchase Order line item may have one of the following states: accepted, rejected, or pending.",
 *             name: "Purchase Order Confirmation",
 *             version: "V02.02.00",
 *         },
 *         description: "An organization that sells products to partners in the supply chain.",
 *         role: "Seller",
 *         roleType: "Organizational",
 *         service: "Seller Service",
 *         serviceClassification: "Business Service",
 *     },
 *     rosettaNetProcessConfigurationName: "3A4",
 * });
 *
 * ```
 */
export class RosettaNetProcessConfiguration extends pulumi.CustomResource {
    /**
     * Get an existing RosettaNetProcessConfiguration resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): RosettaNetProcessConfiguration {
        return new RosettaNetProcessConfiguration(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:logic/latest:RosettaNetProcessConfiguration';

    /**
     * Returns true if the given object is an instance of RosettaNetProcessConfiguration.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RosettaNetProcessConfiguration {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RosettaNetProcessConfiguration.__pulumiType;
    }

    /**
     * The RosettaNet process configuration activity settings.
     */
    public readonly activitySettings!: pulumi.Output<outputs.logic.latest.RosettaNetPipActivitySettingsResponse>;
    /**
     * The changed time.
     */
    public /*out*/ readonly changedTime!: pulumi.Output<string>;
    /**
     * The created time.
     */
    public /*out*/ readonly createdTime!: pulumi.Output<string>;
    /**
     * The integration account RosettaNet ProcessConfiguration properties.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The RosettaNet initiator role settings.
     */
    public readonly initiatorRoleSettings!: pulumi.Output<outputs.logic.latest.RosettaNetPipRoleSettingsResponse>;
    /**
     * The resource location.
     */
    public readonly location!: pulumi.Output<string | undefined>;
    /**
     * The metadata.
     */
    public readonly metadata!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Gets the resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The integration account RosettaNet process code.
     */
    public readonly processCode!: pulumi.Output<string>;
    /**
     * The integration account RosettaNet process name.
     */
    public readonly processName!: pulumi.Output<string>;
    /**
     * The integration account RosettaNet process version.
     */
    public readonly processVersion!: pulumi.Output<string>;
    /**
     * The RosettaNet responder role settings.
     */
    public readonly responderRoleSettings!: pulumi.Output<outputs.logic.latest.RosettaNetPipRoleSettingsResponse>;
    /**
     * The resource tags.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Gets the resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a RosettaNetProcessConfiguration resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RosettaNetProcessConfigurationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RosettaNetProcessConfigurationArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as RosettaNetProcessConfigurationArgs | undefined;
            if (!args || args.activitySettings === undefined) {
                throw new Error("Missing required property 'activitySettings'");
            }
            if (!args || args.initiatorRoleSettings === undefined) {
                throw new Error("Missing required property 'initiatorRoleSettings'");
            }
            if (!args || args.integrationAccountName === undefined) {
                throw new Error("Missing required property 'integrationAccountName'");
            }
            if (!args || args.processCode === undefined) {
                throw new Error("Missing required property 'processCode'");
            }
            if (!args || args.processName === undefined) {
                throw new Error("Missing required property 'processName'");
            }
            if (!args || args.processVersion === undefined) {
                throw new Error("Missing required property 'processVersion'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.responderRoleSettings === undefined) {
                throw new Error("Missing required property 'responderRoleSettings'");
            }
            if (!args || args.rosettaNetProcessConfigurationName === undefined) {
                throw new Error("Missing required property 'rosettaNetProcessConfigurationName'");
            }
            inputs["activitySettings"] = args ? args.activitySettings : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["initiatorRoleSettings"] = args ? args.initiatorRoleSettings : undefined;
            inputs["integrationAccountName"] = args ? args.integrationAccountName : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["metadata"] = args ? args.metadata : undefined;
            inputs["processCode"] = args ? args.processCode : undefined;
            inputs["processName"] = args ? args.processName : undefined;
            inputs["processVersion"] = args ? args.processVersion : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["responderRoleSettings"] = args ? args.responderRoleSettings : undefined;
            inputs["rosettaNetProcessConfigurationName"] = args ? args.rosettaNetProcessConfigurationName : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["changedTime"] = undefined /*out*/;
            inputs["createdTime"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:logic/v20160601:RosettaNetProcessConfiguration" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(RosettaNetProcessConfiguration.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a RosettaNetProcessConfiguration resource.
 */
export interface RosettaNetProcessConfigurationArgs {
    /**
     * The RosettaNet process configuration activity settings.
     */
    readonly activitySettings: pulumi.Input<inputs.logic.latest.RosettaNetPipActivitySettings>;
    /**
     * The integration account RosettaNet ProcessConfiguration properties.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The RosettaNet initiator role settings.
     */
    readonly initiatorRoleSettings: pulumi.Input<inputs.logic.latest.RosettaNetPipRoleSettings>;
    /**
     * The integration account name.
     */
    readonly integrationAccountName: pulumi.Input<string>;
    /**
     * The resource location.
     */
    readonly location?: pulumi.Input<string>;
    /**
     * The metadata.
     */
    readonly metadata?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The integration account RosettaNet process code.
     */
    readonly processCode: pulumi.Input<string>;
    /**
     * The integration account RosettaNet process name.
     */
    readonly processName: pulumi.Input<string>;
    /**
     * The integration account RosettaNet process version.
     */
    readonly processVersion: pulumi.Input<string>;
    /**
     * The resource group name.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * The RosettaNet responder role settings.
     */
    readonly responderRoleSettings: pulumi.Input<inputs.logic.latest.RosettaNetPipRoleSettings>;
    /**
     * The integration account RosettaNet ProcessConfiguration name.
     */
    readonly rosettaNetProcessConfigurationName: pulumi.Input<string>;
    /**
     * The resource tags.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}