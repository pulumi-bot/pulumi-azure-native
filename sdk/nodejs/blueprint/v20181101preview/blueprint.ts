// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Represents a Blueprint definition.
 */
export class Blueprint extends pulumi.CustomResource {
    /**
     * Get an existing Blueprint resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Blueprint {
        return new Blueprint(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:blueprint/v20181101preview:Blueprint';

    /**
     * Returns true if the given object is an instance of Blueprint.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Blueprint {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Blueprint.__pulumiType;
    }

    /**
     * Multi-line explain this resource.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * One-liner string explain this resource.
     */
    public readonly displayName!: pulumi.Output<string | undefined>;
    /**
     * Layout view of the blueprint definition for UI reference.
     */
    public readonly layout!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * Name of this resource.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Parameters required by this blueprint definition.
     */
    public readonly parameters!: pulumi.Output<{[key: string]: outputs.blueprint.v20181101preview.ParameterDefinitionResponse} | undefined>;
    /**
     * Resource group placeholders defined by this blueprint definition.
     */
    public readonly resourceGroups!: pulumi.Output<{[key: string]: outputs.blueprint.v20181101preview.ResourceGroupDefinitionResponse} | undefined>;
    /**
     * Status of the blueprint. This field is readonly.
     */
    public /*out*/ readonly status!: pulumi.Output<outputs.blueprint.v20181101preview.BlueprintStatusResponse>;
    /**
     * The scope where this blueprint definition can be assigned.
     */
    public readonly targetScope!: pulumi.Output<string>;
    /**
     * Type of this resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * Published versions of this blueprint definition.
     */
    public readonly versions!: pulumi.Output<{[key: string]: any} | undefined>;

    /**
     * Create a Blueprint resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BlueprintArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.blueprintName === undefined) {
                throw new Error("Missing required property 'blueprintName'");
            }
            if (!args || args.resourceScope === undefined) {
                throw new Error("Missing required property 'resourceScope'");
            }
            if (!args || args.targetScope === undefined) {
                throw new Error("Missing required property 'targetScope'");
            }
            inputs["blueprintName"] = args ? args.blueprintName : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["layout"] = args ? args.layout : undefined;
            inputs["parameters"] = args ? args.parameters : undefined;
            inputs["resourceGroups"] = args ? args.resourceGroups : undefined;
            inputs["resourceScope"] = args ? args.resourceScope : undefined;
            inputs["targetScope"] = args ? args.targetScope : undefined;
            inputs["versions"] = args ? args.versions : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["description"] = undefined /*out*/;
            inputs["displayName"] = undefined /*out*/;
            inputs["layout"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["parameters"] = undefined /*out*/;
            inputs["resourceGroups"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
            inputs["targetScope"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["versions"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Blueprint.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Blueprint resource.
 */
export interface BlueprintArgs {
    /**
     * Name of the blueprint definition.
     */
    readonly blueprintName: pulumi.Input<string>;
    /**
     * Multi-line explain this resource.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * One-liner string explain this resource.
     */
    readonly displayName?: pulumi.Input<string>;
    /**
     * Layout view of the blueprint definition for UI reference.
     */
    readonly layout?: pulumi.Input<{[key: string]: any}>;
    /**
     * Parameters required by this blueprint definition.
     */
    readonly parameters?: pulumi.Input<{[key: string]: pulumi.Input<inputs.blueprint.v20181101preview.ParameterDefinition>}>;
    /**
     * Resource group placeholders defined by this blueprint definition.
     */
    readonly resourceGroups?: pulumi.Input<{[key: string]: pulumi.Input<inputs.blueprint.v20181101preview.ResourceGroupDefinition>}>;
    /**
     * The scope of the resource. Valid scopes are: management group (format: '/providers/Microsoft.Management/managementGroups/{managementGroup}'), subscription (format: '/subscriptions/{subscriptionId}').
     */
    readonly resourceScope: pulumi.Input<string>;
    /**
     * The scope where this blueprint definition can be assigned.
     */
    readonly targetScope: pulumi.Input<string>;
    /**
     * Published versions of this blueprint definition.
     */
    readonly versions?: pulumi.Input<{[key: string]: any}>;
}
