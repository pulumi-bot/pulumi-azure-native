// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Represents a published Blueprint.
 */
export class PublishedBlueprint extends pulumi.CustomResource {
    /**
     * Get an existing PublishedBlueprint resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): PublishedBlueprint {
        return new PublishedBlueprint(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:management/v20171111preview:PublishedBlueprint';

    /**
     * Returns true if the given object is an instance of PublishedBlueprint.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is PublishedBlueprint {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === PublishedBlueprint.__pulumiType;
    }

    /**
     * Name of the Blueprint definition.
     */
    public readonly blueprintName!: pulumi.Output<string | undefined>;
    /**
     * Version-specific change notes
     */
    public /*out*/ readonly changeNotes!: pulumi.Output<string | undefined>;
    /**
     * Multi-line explain this resource.
     */
    public /*out*/ readonly description!: pulumi.Output<string | undefined>;
    /**
     * One-liner string explain this resource.
     */
    public /*out*/ readonly displayName!: pulumi.Output<string | undefined>;
    /**
     * Name of this resource.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Parameters required by this Blueprint definition.
     */
    public /*out*/ readonly parameters!: pulumi.Output<{[key: string]: outputs.management.v20171111preview.ParameterDefinitionResponse} | undefined>;
    /**
     * Resource group placeholders defined by this Blueprint definition.
     */
    public /*out*/ readonly resourceGroups!: pulumi.Output<{[key: string]: outputs.management.v20171111preview.ResourceGroupDefinitionResponse} | undefined>;
    /**
     * Status of the Blueprint. This field is readonly.
     */
    public /*out*/ readonly status!: pulumi.Output<outputs.management.v20171111preview.BlueprintStatusResponse>;
    /**
     * The scope where this Blueprint can be applied.
     */
    public /*out*/ readonly targetScope!: pulumi.Output<string | undefined>;
    /**
     * Type of this resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a PublishedBlueprint resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: PublishedBlueprintArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.blueprintName === undefined) {
                throw new Error("Missing required property 'blueprintName'");
            }
            if (!args || args.managementGroupName === undefined) {
                throw new Error("Missing required property 'managementGroupName'");
            }
            if (!args || args.versionId === undefined) {
                throw new Error("Missing required property 'versionId'");
            }
            inputs["blueprintName"] = args ? args.blueprintName : undefined;
            inputs["managementGroupName"] = args ? args.managementGroupName : undefined;
            inputs["versionId"] = args ? args.versionId : undefined;
            inputs["changeNotes"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["displayName"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["parameters"] = undefined /*out*/;
            inputs["resourceGroups"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
            inputs["targetScope"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["blueprintName"] = undefined /*out*/;
            inputs["changeNotes"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["displayName"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["parameters"] = undefined /*out*/;
            inputs["resourceGroups"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
            inputs["targetScope"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(PublishedBlueprint.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a PublishedBlueprint resource.
 */
export interface PublishedBlueprintArgs {
    /**
     * name of the blueprint.
     */
    readonly blueprintName: pulumi.Input<string>;
    /**
     * ManagementGroup where blueprint stores.
     */
    readonly managementGroupName: pulumi.Input<string>;
    /**
     * version of the published blueprint.
     */
    readonly versionId: pulumi.Input<string>;
}
