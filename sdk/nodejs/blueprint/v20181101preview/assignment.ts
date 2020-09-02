// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * Represents a blueprint assignment.
 */
export class Assignment extends pulumi.CustomResource {
    /**
     * Get an existing Assignment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Assignment {
        return new Assignment(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:blueprint/v20181101preview:Assignment';

    /**
     * Returns true if the given object is an instance of Assignment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Assignment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Assignment.__pulumiType;
    }

    /**
     * ID of the published version of a blueprint definition.
     */
    public readonly blueprintId!: pulumi.Output<string | undefined>;
    /**
     * Multi-line explain this resource.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * One-liner string explain this resource.
     */
    public readonly displayName!: pulumi.Output<string | undefined>;
    /**
     * Managed identity for this blueprint assignment.
     */
    public readonly identity!: pulumi.Output<outputs.blueprint.v20181101preview.ManagedServiceIdentityResponse>;
    /**
     * The location of this blueprint assignment.
     */
    public readonly location!: pulumi.Output<string>;
    /**
     * Defines how resources deployed by a blueprint assignment are locked.
     */
    public readonly locks!: pulumi.Output<outputs.blueprint.v20181101preview.AssignmentLockSettingsResponse | undefined>;
    /**
     * Name of this resource.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Blueprint assignment parameter values.
     */
    public readonly parameters!: pulumi.Output<{[key: string]: outputs.blueprint.v20181101preview.ParameterValueResponse}>;
    /**
     * State of the blueprint assignment.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * Names and locations of resource group placeholders.
     */
    public readonly resourceGroups!: pulumi.Output<{[key: string]: outputs.blueprint.v20181101preview.ResourceGroupValueResponse}>;
    /**
     * The target subscription scope of the blueprint assignment (format: '/subscriptions/{subscriptionId}'). For management group level assignments, the property is required.
     */
    public readonly scope!: pulumi.Output<string | undefined>;
    /**
     * Status of blueprint assignment. This field is readonly.
     */
    public /*out*/ readonly status!: pulumi.Output<outputs.blueprint.v20181101preview.AssignmentStatusResponse>;
    /**
     * Type of this resource.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a Assignment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AssignmentArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AssignmentArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as AssignmentArgs | undefined;
            if (!args || args.assignmentName === undefined) {
                throw new Error("Missing required property 'assignmentName'");
            }
            if (!args || args.identity === undefined) {
                throw new Error("Missing required property 'identity'");
            }
            if (!args || args.location === undefined) {
                throw new Error("Missing required property 'location'");
            }
            if (!args || args.parameters === undefined) {
                throw new Error("Missing required property 'parameters'");
            }
            if (!args || args.resourceGroups === undefined) {
                throw new Error("Missing required property 'resourceGroups'");
            }
            if (!args || args.resourceScope === undefined) {
                throw new Error("Missing required property 'resourceScope'");
            }
            inputs["assignmentName"] = args ? args.assignmentName : undefined;
            inputs["blueprintId"] = args ? args.blueprintId : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["identity"] = args ? args.identity : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["locks"] = args ? args.locks : undefined;
            inputs["parameters"] = args ? args.parameters : undefined;
            inputs["resourceGroups"] = args ? args.resourceGroups : undefined;
            inputs["resourceScope"] = args ? args.resourceScope : undefined;
            inputs["scope"] = args ? args.scope : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Assignment.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Assignment resource.
 */
export interface AssignmentArgs {
    /**
     * Name of the blueprint assignment.
     */
    readonly assignmentName: pulumi.Input<string>;
    /**
     * ID of the published version of a blueprint definition.
     */
    readonly blueprintId?: pulumi.Input<string>;
    /**
     * Multi-line explain this resource.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * One-liner string explain this resource.
     */
    readonly displayName?: pulumi.Input<string>;
    /**
     * Managed identity for this blueprint assignment.
     */
    readonly identity: pulumi.Input<inputs.blueprint.v20181101preview.ManagedServiceIdentity>;
    /**
     * The location of this blueprint assignment.
     */
    readonly location: pulumi.Input<string>;
    /**
     * Defines how resources deployed by a blueprint assignment are locked.
     */
    readonly locks?: pulumi.Input<inputs.blueprint.v20181101preview.AssignmentLockSettings>;
    /**
     * Blueprint assignment parameter values.
     */
    readonly parameters: pulumi.Input<{[key: string]: pulumi.Input<inputs.blueprint.v20181101preview.ParameterValue>}>;
    /**
     * Names and locations of resource group placeholders.
     */
    readonly resourceGroups: pulumi.Input<{[key: string]: pulumi.Input<inputs.blueprint.v20181101preview.ResourceGroupValue>}>;
    /**
     * The scope of the resource. Valid scopes are: management group (format: '/providers/Microsoft.Management/managementGroups/{managementGroup}'), subscription (format: '/subscriptions/{subscriptionId}').
     */
    readonly resourceScope: pulumi.Input<string>;
    /**
     * The target subscription scope of the blueprint assignment (format: '/subscriptions/{subscriptionId}'). For management group level assignments, the property is required.
     */
    readonly scope?: pulumi.Input<string>;
}
