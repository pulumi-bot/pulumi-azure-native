// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * Resource group information.
 */
export class ResourceGroup extends pulumi.CustomResource {
    /**
     * Get an existing ResourceGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ResourceGroupState, opts?: pulumi.CustomResourceOptions): ResourceGroup {
        return new ResourceGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:core:ResourceGroup';

    /**
     * Returns true if the given object is an instance of ResourceGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ResourceGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ResourceGroup.__pulumiType;
    }

    /**
     * The location of the resource group. It cannot be changed after the resource group has been created. It must be one of the supported Azure locations.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * The ID of the resource that manages this resource group.
     */
    public readonly managedBy!: pulumi.Output<string | undefined>;
    /**
     * The name of the resource group.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The resource group properties.
     */
    public readonly properties!: pulumi.Output<outputs.core.ResourceGroupPropertiesResponse>;
    /**
     * The tags attached to the resource group.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The type of the resource group.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a ResourceGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ResourceGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ResourceGroupArgs | ResourceGroupState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ResourceGroupState | undefined;
            inputs["location"] = state ? state.location : undefined;
            inputs["managedBy"] = state ? state.managedBy : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["properties"] = state ? state.properties : undefined;
            inputs["tags"] = state ? state.tags : undefined;
            inputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as ResourceGroupArgs | undefined;
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            if (!args || args.name === undefined) {
                throw new Error("Missing required property 'name'");
            }
            inputs["location"] = args ? args.location : undefined;
            inputs["managedBy"] = args ? args.managedBy : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["properties"] = args ? args.properties : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ResourceGroup.__pulumiType, name, inputs, opts);
    }
}

/**
 * Resource group information.
 */
export interface ResourceGroupState {
    /**
     * The location of the resource group. It cannot be changed after the resource group has been created. It must be one of the supported Azure locations.
     */
    readonly location: pulumi.Input<string>;
    /**
     * The ID of the resource that manages this resource group.
     */
    readonly managedBy?: pulumi.Input<string>;
    /**
     * The name of the resource group.
     */
    readonly name: pulumi.Input<string>;
    /**
     * The resource group properties.
     */
    readonly properties: pulumi.Input<inputs.core.ResourceGroupPropertiesResponse>;
    /**
     * The tags attached to the resource group.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The type of the resource group.
     */
    readonly type: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ResourceGroup resource.
 */
export interface ResourceGroupArgs {
    /**
     * The location of the resource group. It cannot be changed after the resource group has been created. It must be one of the supported Azure locations.
     */
    readonly location: pulumi.Input<string>;
    /**
     * The ID of the resource that manages this resource group.
     */
    readonly managedBy?: pulumi.Input<string>;
    /**
     * The name of the resource group to create or update. Can include alphanumeric, underscore, parentheses, hyphen, period (except at end), and Unicode characters that match the allowed characters.
     */
    readonly name: pulumi.Input<string>;
    /**
     * The resource group properties.
     */
    readonly properties?: pulumi.Input<inputs.core.ResourceGroupProperties>;
    /**
     * The tags attached to the resource group.
     */
    readonly tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
