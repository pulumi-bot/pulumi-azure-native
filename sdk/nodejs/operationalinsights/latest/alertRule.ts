// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Alert rule.
 *
 * ## Example Usage
 * ### Creates or updates a Fusion alert rule.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const alertRule = new azurerm.operationalinsights.latest.AlertRule("alertRule", {
 *     etag: "3d00c3ca-0000-0100-0000-5d42d5010000",
 *     kind: "Fusion",
 *     resourceGroupName: "myRg",
 *     ruleId: "myFirstFusionRule",
 *     workspaceName: "myWorkspace",
 * });
 *
 * ```
 * ### Creates or updates a MicrosoftSecurityIncidentCreation rule.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const alertRule = new azurerm.operationalinsights.latest.AlertRule("alertRule", {
 *     etag: "\"260097e0-0000-0d00-0000-5d6fa88f0000\"",
 *     kind: "MicrosoftSecurityIncidentCreation",
 *     resourceGroupName: "myRg",
 *     ruleId: "microsoftSecurityIncidentCreationRuleExample",
 *     workspaceName: "myWorkspace",
 * });
 *
 * ```
 * ### Creates or updates a Scheduled alert rule.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const alertRule = new azurerm.operationalinsights.latest.AlertRule("alertRule", {
 *     etag: "\"0300bf09-0000-0000-0000-5c37296e0000\"",
 *     kind: "Scheduled",
 *     resourceGroupName: "myRg",
 *     ruleId: "73e01a99-5cd7-4139-a149-9f2736ff2ab5",
 *     workspaceName: "myWorkspace",
 * });
 *
 * ```
 */
export class AlertRule extends pulumi.CustomResource {
    /**
     * Get an existing AlertRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): AlertRule {
        return new AlertRule(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:operationalinsights/latest:AlertRule';

    /**
     * Returns true if the given object is an instance of AlertRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AlertRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AlertRule.__pulumiType;
    }

    /**
     * Etag of the azure resource
     */
    public readonly etag!: pulumi.Output<string | undefined>;
    /**
     * The alert rule kind
     */
    public readonly kind!: pulumi.Output<string>;
    /**
     * Azure resource name
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Azure resource type
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a AlertRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AlertRuleArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.kind === undefined) {
                throw new Error("Missing required property 'kind'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.ruleId === undefined) {
                throw new Error("Missing required property 'ruleId'");
            }
            if (!args || args.workspaceName === undefined) {
                throw new Error("Missing required property 'workspaceName'");
            }
            inputs["etag"] = args ? args.etag : undefined;
            inputs["kind"] = args ? args.kind : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["ruleId"] = args ? args.ruleId : undefined;
            inputs["workspaceName"] = args ? args.workspaceName : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["etag"] = undefined /*out*/;
            inputs["kind"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:operationalinsights/v20200101:AlertRule" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(AlertRule.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a AlertRule resource.
 */
export interface AlertRuleArgs {
    /**
     * Etag of the azure resource
     */
    readonly etag?: pulumi.Input<string>;
    /**
     * The alert rule kind
     */
    readonly kind: pulumi.Input<string>;
    /**
     * The name of the resource group within the user's subscription. The name is case insensitive.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Alert rule ID
     */
    readonly ruleId: pulumi.Input<string>;
    /**
     * The name of the workspace.
     */
    readonly workspaceName: pulumi.Input<string>;
}
