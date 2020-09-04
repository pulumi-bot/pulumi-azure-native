// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * A rules engine configuration containing a list of rules that will run to modify the runtime behavior of the request and response.
 *
 * ## Example Usage
 * ### Create or update a specific Rules Engine Configuration
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const rulesEngine = new azurerm.network.latest.RulesEngine("rulesEngine", {
 *     frontDoorName: "frontDoor1",
 *     resourceGroupName: "rg1",
 *     rules: [
 *         {
 *             action: {
 *                 routeConfigurationOverride: {
 *                     odataType: "#Microsoft.Azure.FrontDoor.Models.FrontdoorRedirectConfiguration",
 *                 },
 *             },
 *             matchConditions: [{
 *                 rulesEngineMatchValue: ["CH"],
 *                 rulesEngineMatchVariable: "RemoteAddr",
 *                 rulesEngineOperator: "GeoMatch",
 *             }],
 *             matchProcessingBehavior: "Stop",
 *             name: "Rule1",
 *             priority: 1,
 *         },
 *         {
 *             action: {
 *                 responseHeaderActions: [{
 *                     headerActionType: "Overwrite",
 *                     headerName: "Cache-Control",
 *                     value: "public, max-age=31536000",
 *                 }],
 *             },
 *             matchConditions: [{
 *                 rulesEngineMatchValue: ["jpg"],
 *                 rulesEngineMatchVariable: "RequestFilenameExtension",
 *                 rulesEngineOperator: "Equal",
 *                 transforms: ["Lowercase"],
 *             }],
 *             name: "Rule2",
 *             priority: 2,
 *         },
 *         {
 *             action: {
 *                 routeConfigurationOverride: {
 *                     odataType: "#Microsoft.Azure.FrontDoor.Models.FrontdoorForwardingConfiguration",
 *                 },
 *             },
 *             matchConditions: [{
 *                 negateCondition: false,
 *                 rulesEngineMatchValue: ["allowoverride"],
 *                 rulesEngineMatchVariable: "RequestHeader",
 *                 rulesEngineOperator: "Equal",
 *                 selector: "Rules-Engine-Route-Forward",
 *                 transforms: ["Lowercase"],
 *             }],
 *             name: "Rule3",
 *             priority: 3,
 *         },
 *     ],
 *     rulesEngineName: "rulesEngine1",
 * });
 *
 * ```
 */
export class RulesEngine extends pulumi.CustomResource {
    /**
     * Get an existing RulesEngine resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): RulesEngine {
        return new RulesEngine(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:network/latest:RulesEngine';

    /**
     * Returns true if the given object is an instance of RulesEngine.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RulesEngine {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RulesEngine.__pulumiType;
    }

    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Resource status.
     */
    public readonly resourceState!: pulumi.Output<string | undefined>;
    /**
     * A list of rules that define a particular Rules Engine Configuration.
     */
    public readonly rules!: pulumi.Output<outputs.network.latest.RulesEngineRuleResponse[] | undefined>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a RulesEngine resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RulesEngineArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.frontDoorName === undefined) {
                throw new Error("Missing required property 'frontDoorName'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.rulesEngineName === undefined) {
                throw new Error("Missing required property 'rulesEngineName'");
            }
            inputs["frontDoorName"] = args ? args.frontDoorName : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["resourceState"] = args ? args.resourceState : undefined;
            inputs["rules"] = args ? args.rules : undefined;
            inputs["rulesEngineName"] = args ? args.rulesEngineName : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["name"] = undefined /*out*/;
            inputs["resourceState"] = undefined /*out*/;
            inputs["rules"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:network/v20200101:RulesEngine" }, { type: "azurerm:network/v20200401:RulesEngine" }, { type: "azurerm:network/v20200501:RulesEngine" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(RulesEngine.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a RulesEngine resource.
 */
export interface RulesEngineArgs {
    /**
     * Name of the Front Door which is globally unique.
     */
    readonly frontDoorName: pulumi.Input<string>;
    /**
     * Name of the Resource group within the Azure subscription.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Resource status.
     */
    readonly resourceState?: pulumi.Input<string>;
    /**
     * A list of rules that define a particular Rules Engine Configuration.
     */
    readonly rules?: pulumi.Input<pulumi.Input<inputs.network.latest.RulesEngineRule>[]>;
    /**
     * Name of the Rules Engine which is unique within the Front Door.
     */
    readonly rulesEngineName: pulumi.Input<string>;
}
