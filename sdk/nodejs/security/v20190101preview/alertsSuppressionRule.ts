// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Describes the suppression rule
 */
export class AlertsSuppressionRule extends pulumi.CustomResource {
    /**
     * Get an existing AlertsSuppressionRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): AlertsSuppressionRule {
        return new AlertsSuppressionRule(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:security/v20190101preview:AlertsSuppressionRule';

    /**
     * Returns true if the given object is an instance of AlertsSuppressionRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AlertsSuppressionRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AlertsSuppressionRule.__pulumiType;
    }

    /**
     * Type of the alert to automatically suppress. For all alert types, use '*'
     */
    public readonly alertType!: pulumi.Output<string>;
    /**
     * Any comment regarding the rule
     */
    public readonly comment!: pulumi.Output<string | undefined>;
    /**
     * Expiration date of the rule, if value is not provided or provided as null this field will default to the maximum allowed expiration date.
     */
    public readonly expirationDateUtc!: pulumi.Output<string | undefined>;
    /**
     * The last time this rule was modified
     */
    public /*out*/ readonly lastModifiedUtc!: pulumi.Output<string>;
    /**
     * Resource name
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The reason for dismissing the alert
     */
    public readonly reason!: pulumi.Output<string>;
    /**
     * Possible states of the rule
     */
    public readonly state!: pulumi.Output<string>;
    /**
     * The suppression conditions
     */
    public readonly suppressionAlertsScope!: pulumi.Output<outputs.security.v20190101preview.SuppressionAlertsScopeResponse | undefined>;
    /**
     * Resource type
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a AlertsSuppressionRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AlertsSuppressionRuleArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.alertType === undefined) {
                throw new Error("Missing required property 'alertType'");
            }
            if (!args || args.alertsSuppressionRuleName === undefined) {
                throw new Error("Missing required property 'alertsSuppressionRuleName'");
            }
            if (!args || args.reason === undefined) {
                throw new Error("Missing required property 'reason'");
            }
            if (!args || args.state === undefined) {
                throw new Error("Missing required property 'state'");
            }
            inputs["alertType"] = args ? args.alertType : undefined;
            inputs["alertsSuppressionRuleName"] = args ? args.alertsSuppressionRuleName : undefined;
            inputs["comment"] = args ? args.comment : undefined;
            inputs["expirationDateUtc"] = args ? args.expirationDateUtc : undefined;
            inputs["reason"] = args ? args.reason : undefined;
            inputs["state"] = args ? args.state : undefined;
            inputs["suppressionAlertsScope"] = args ? args.suppressionAlertsScope : undefined;
            inputs["lastModifiedUtc"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["alertType"] = undefined /*out*/;
            inputs["comment"] = undefined /*out*/;
            inputs["expirationDateUtc"] = undefined /*out*/;
            inputs["lastModifiedUtc"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["reason"] = undefined /*out*/;
            inputs["state"] = undefined /*out*/;
            inputs["suppressionAlertsScope"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:security/latest:AlertsSuppressionRule" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(AlertsSuppressionRule.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a AlertsSuppressionRule resource.
 */
export interface AlertsSuppressionRuleArgs {
    /**
     * Type of the alert to automatically suppress. For all alert types, use '*'
     */
    readonly alertType: pulumi.Input<string>;
    /**
     * The unique name of the suppression alert rule
     */
    readonly alertsSuppressionRuleName: pulumi.Input<string>;
    /**
     * Any comment regarding the rule
     */
    readonly comment?: pulumi.Input<string>;
    /**
     * Expiration date of the rule, if value is not provided or provided as null this field will default to the maximum allowed expiration date.
     */
    readonly expirationDateUtc?: pulumi.Input<string>;
    /**
     * The reason for dismissing the alert
     */
    readonly reason: pulumi.Input<string>;
    /**
     * Possible states of the rule
     */
    readonly state: pulumi.Input<string>;
    /**
     * The suppression conditions
     */
    readonly suppressionAlertsScope?: pulumi.Input<inputs.security.v20190101preview.SuppressionAlertsScope>;
}
