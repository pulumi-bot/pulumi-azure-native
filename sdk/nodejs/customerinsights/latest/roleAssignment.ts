// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

/**
 * The Role Assignment resource format.
 */
export class RoleAssignment extends pulumi.CustomResource {
    /**
     * Get an existing RoleAssignment resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): RoleAssignment {
        return new RoleAssignment(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azure-nextgen:customerinsights/latest:RoleAssignment';

    /**
     * Returns true if the given object is an instance of RoleAssignment.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RoleAssignment {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RoleAssignment.__pulumiType;
    }

    /**
     * The name of the metadata object.
     */
    public readonly assignmentName!: pulumi.Output<string>;
    /**
     * Widget types set for the assignment.
     */
    public readonly conflationPolicies!: pulumi.Output<outputs.customerinsights.latest.ResourceSetDescriptionResponse | undefined>;
    /**
     * Connectors set for the assignment.
     */
    public readonly connectors!: pulumi.Output<outputs.customerinsights.latest.ResourceSetDescriptionResponse | undefined>;
    /**
     * Localized description for the metadata.
     */
    public readonly description!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Localized display names for the metadata.
     */
    public readonly displayName!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * Interactions set for the assignment.
     */
    public readonly interactions!: pulumi.Output<outputs.customerinsights.latest.ResourceSetDescriptionResponse | undefined>;
    /**
     * Kpis set for the assignment.
     */
    public readonly kpis!: pulumi.Output<outputs.customerinsights.latest.ResourceSetDescriptionResponse | undefined>;
    /**
     * Links set for the assignment.
     */
    public readonly links!: pulumi.Output<outputs.customerinsights.latest.ResourceSetDescriptionResponse | undefined>;
    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The principals being assigned to.
     */
    public readonly principals!: pulumi.Output<outputs.customerinsights.latest.AssignmentPrincipalResponse[]>;
    /**
     * Profiles set for the assignment.
     */
    public readonly profiles!: pulumi.Output<outputs.customerinsights.latest.ResourceSetDescriptionResponse | undefined>;
    /**
     * Provisioning state.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * The Role assignments set for the relationship links.
     */
    public readonly relationshipLinks!: pulumi.Output<outputs.customerinsights.latest.ResourceSetDescriptionResponse | undefined>;
    /**
     * The Role assignments set for the relationships.
     */
    public readonly relationships!: pulumi.Output<outputs.customerinsights.latest.ResourceSetDescriptionResponse | undefined>;
    /**
     * Type of roles.
     */
    public readonly role!: pulumi.Output<string>;
    /**
     * The Role assignments set for the assignment.
     */
    public readonly roleAssignments!: pulumi.Output<outputs.customerinsights.latest.ResourceSetDescriptionResponse | undefined>;
    /**
     * Sas Policies set for the assignment.
     */
    public readonly sasPolicies!: pulumi.Output<outputs.customerinsights.latest.ResourceSetDescriptionResponse | undefined>;
    /**
     * The Role assignments set for the assignment.
     */
    public readonly segments!: pulumi.Output<outputs.customerinsights.latest.ResourceSetDescriptionResponse | undefined>;
    /**
     * The hub name.
     */
    public /*out*/ readonly tenantId!: pulumi.Output<string>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * Views set for the assignment.
     */
    public readonly views!: pulumi.Output<outputs.customerinsights.latest.ResourceSetDescriptionResponse | undefined>;
    /**
     * Widget types set for the assignment.
     */
    public readonly widgetTypes!: pulumi.Output<outputs.customerinsights.latest.ResourceSetDescriptionResponse | undefined>;

    /**
     * Create a RoleAssignment resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RoleAssignmentArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            if (!args || args.assignmentName === undefined) {
                throw new Error("Missing required property 'assignmentName'");
            }
            if (!args || args.hubName === undefined) {
                throw new Error("Missing required property 'hubName'");
            }
            if (!args || args.principals === undefined) {
                throw new Error("Missing required property 'principals'");
            }
            if (!args || args.resourceGroupName === undefined) {
                throw new Error("Missing required property 'resourceGroupName'");
            }
            if (!args || args.role === undefined) {
                throw new Error("Missing required property 'role'");
            }
            inputs["assignmentName"] = args ? args.assignmentName : undefined;
            inputs["conflationPolicies"] = args ? args.conflationPolicies : undefined;
            inputs["connectors"] = args ? args.connectors : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["hubName"] = args ? args.hubName : undefined;
            inputs["interactions"] = args ? args.interactions : undefined;
            inputs["kpis"] = args ? args.kpis : undefined;
            inputs["links"] = args ? args.links : undefined;
            inputs["principals"] = args ? args.principals : undefined;
            inputs["profiles"] = args ? args.profiles : undefined;
            inputs["relationshipLinks"] = args ? args.relationshipLinks : undefined;
            inputs["relationships"] = args ? args.relationships : undefined;
            inputs["resourceGroupName"] = args ? args.resourceGroupName : undefined;
            inputs["role"] = args ? args.role : undefined;
            inputs["roleAssignments"] = args ? args.roleAssignments : undefined;
            inputs["sasPolicies"] = args ? args.sasPolicies : undefined;
            inputs["segments"] = args ? args.segments : undefined;
            inputs["views"] = args ? args.views : undefined;
            inputs["widgetTypes"] = args ? args.widgetTypes : undefined;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["tenantId"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        } else {
            inputs["assignmentName"] = undefined /*out*/;
            inputs["conflationPolicies"] = undefined /*out*/;
            inputs["connectors"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["displayName"] = undefined /*out*/;
            inputs["interactions"] = undefined /*out*/;
            inputs["kpis"] = undefined /*out*/;
            inputs["links"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["principals"] = undefined /*out*/;
            inputs["profiles"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["relationshipLinks"] = undefined /*out*/;
            inputs["relationships"] = undefined /*out*/;
            inputs["role"] = undefined /*out*/;
            inputs["roleAssignments"] = undefined /*out*/;
            inputs["sasPolicies"] = undefined /*out*/;
            inputs["segments"] = undefined /*out*/;
            inputs["tenantId"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
            inputs["views"] = undefined /*out*/;
            inputs["widgetTypes"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azure-nextgen:customerinsights/v20170101:RoleAssignment" }, { type: "azure-nextgen:customerinsights/v20170426:RoleAssignment" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(RoleAssignment.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a RoleAssignment resource.
 */
export interface RoleAssignmentArgs {
    /**
     * The assignment name
     */
    readonly assignmentName: pulumi.Input<string>;
    /**
     * Widget types set for the assignment.
     */
    readonly conflationPolicies?: pulumi.Input<inputs.customerinsights.latest.ResourceSetDescription>;
    /**
     * Connectors set for the assignment.
     */
    readonly connectors?: pulumi.Input<inputs.customerinsights.latest.ResourceSetDescription>;
    /**
     * Localized description for the metadata.
     */
    readonly description?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Localized display names for the metadata.
     */
    readonly displayName?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The name of the hub.
     */
    readonly hubName: pulumi.Input<string>;
    /**
     * Interactions set for the assignment.
     */
    readonly interactions?: pulumi.Input<inputs.customerinsights.latest.ResourceSetDescription>;
    /**
     * Kpis set for the assignment.
     */
    readonly kpis?: pulumi.Input<inputs.customerinsights.latest.ResourceSetDescription>;
    /**
     * Links set for the assignment.
     */
    readonly links?: pulumi.Input<inputs.customerinsights.latest.ResourceSetDescription>;
    /**
     * The principals being assigned to.
     */
    readonly principals: pulumi.Input<pulumi.Input<inputs.customerinsights.latest.AssignmentPrincipal>[]>;
    /**
     * Profiles set for the assignment.
     */
    readonly profiles?: pulumi.Input<inputs.customerinsights.latest.ResourceSetDescription>;
    /**
     * The Role assignments set for the relationship links.
     */
    readonly relationshipLinks?: pulumi.Input<inputs.customerinsights.latest.ResourceSetDescription>;
    /**
     * The Role assignments set for the relationships.
     */
    readonly relationships?: pulumi.Input<inputs.customerinsights.latest.ResourceSetDescription>;
    /**
     * The name of the resource group.
     */
    readonly resourceGroupName: pulumi.Input<string>;
    /**
     * Type of roles.
     */
    readonly role: pulumi.Input<string>;
    /**
     * The Role assignments set for the assignment.
     */
    readonly roleAssignments?: pulumi.Input<inputs.customerinsights.latest.ResourceSetDescription>;
    /**
     * Sas Policies set for the assignment.
     */
    readonly sasPolicies?: pulumi.Input<inputs.customerinsights.latest.ResourceSetDescription>;
    /**
     * The Role assignments set for the assignment.
     */
    readonly segments?: pulumi.Input<inputs.customerinsights.latest.ResourceSetDescription>;
    /**
     * Views set for the assignment.
     */
    readonly views?: pulumi.Input<inputs.customerinsights.latest.ResourceSetDescription>;
    /**
     * Widget types set for the assignment.
     */
    readonly widgetTypes?: pulumi.Input<inputs.customerinsights.latest.ResourceSetDescription>;
}
