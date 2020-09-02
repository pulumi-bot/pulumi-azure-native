// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as utilities from "../../utilities";

/**
 * The remediation definition.
 *
 * ## Example Usage
 * ### Create remediation at management group scope
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as azurerm from "@pulumi/azurerm";
 *
 * const remediationAtManagementGroup = new azurerm.policyinsights.v20190701.RemediationAtManagementGroup("remediationAtManagementGroup", {
 *     managementGroupId: "financeMg",
 *     managementGroupsNamespace: "Microsoft.Management",
 *     policyAssignmentId: "/providers/microsoft.management/managementGroups/financeMg/providers/microsoft.authorization/policyassignments/b101830944f246d8a14088c5",
 *     remediationName: "storageRemediation",
 * });
 *
 * ```
 */
export class RemediationAtManagementGroup extends pulumi.CustomResource {
    /**
     * Get an existing RemediationAtManagementGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): RemediationAtManagementGroup {
        return new RemediationAtManagementGroup(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:policyinsights/v20190701:RemediationAtManagementGroup';

    /**
     * Returns true if the given object is an instance of RemediationAtManagementGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RemediationAtManagementGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RemediationAtManagementGroup.__pulumiType;
    }

    /**
     * The time at which the remediation was created.
     */
    public /*out*/ readonly createdOn!: pulumi.Output<string>;
    /**
     * The deployment status summary for all deployments created by the remediation.
     */
    public /*out*/ readonly deploymentStatus!: pulumi.Output<outputs.policyinsights.v20190701.RemediationDeploymentSummaryResponse>;
    /**
     * The filters that will be applied to determine which resources to remediate.
     */
    public readonly filters!: pulumi.Output<outputs.policyinsights.v20190701.RemediationFiltersResponse | undefined>;
    /**
     * The time at which the remediation was last updated.
     */
    public /*out*/ readonly lastUpdatedOn!: pulumi.Output<string>;
    /**
     * The name of the remediation.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The resource ID of the policy assignment that should be remediated.
     */
    public readonly policyAssignmentId!: pulumi.Output<string | undefined>;
    /**
     * The policy definition reference ID of the individual definition that should be remediated. Required when the policy assignment being remediated assigns a policy set definition.
     */
    public readonly policyDefinitionReferenceId!: pulumi.Output<string | undefined>;
    /**
     * The status of the remediation.
     */
    public /*out*/ readonly provisioningState!: pulumi.Output<string>;
    /**
     * The way resources to remediate are discovered. Defaults to ExistingNonCompliant if not specified.
     */
    public readonly resourceDiscoveryMode!: pulumi.Output<string | undefined>;
    /**
     * The type of the remediation.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;

    /**
     * Create a RemediationAtManagementGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RemediationAtManagementGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RemediationAtManagementGroupArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as RemediationAtManagementGroupArgs | undefined;
            if (!args || args.managementGroupId === undefined) {
                throw new Error("Missing required property 'managementGroupId'");
            }
            if (!args || args.managementGroupsNamespace === undefined) {
                throw new Error("Missing required property 'managementGroupsNamespace'");
            }
            if (!args || args.remediationName === undefined) {
                throw new Error("Missing required property 'remediationName'");
            }
            inputs["filters"] = args ? args.filters : undefined;
            inputs["managementGroupId"] = args ? args.managementGroupId : undefined;
            inputs["managementGroupsNamespace"] = args ? args.managementGroupsNamespace : undefined;
            inputs["policyAssignmentId"] = args ? args.policyAssignmentId : undefined;
            inputs["policyDefinitionReferenceId"] = args ? args.policyDefinitionReferenceId : undefined;
            inputs["remediationName"] = args ? args.remediationName : undefined;
            inputs["resourceDiscoveryMode"] = args ? args.resourceDiscoveryMode : undefined;
            inputs["createdOn"] = undefined /*out*/;
            inputs["deploymentStatus"] = undefined /*out*/;
            inputs["lastUpdatedOn"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["provisioningState"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        const aliasOpts = { aliases: [{ type: "azurerm:policyinsights/latest:RemediationAtManagementGroup" }] };
        opts = opts ? pulumi.mergeOptions(opts, aliasOpts) : aliasOpts;
        super(RemediationAtManagementGroup.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a RemediationAtManagementGroup resource.
 */
export interface RemediationAtManagementGroupArgs {
    /**
     * The filters that will be applied to determine which resources to remediate.
     */
    readonly filters?: pulumi.Input<inputs.policyinsights.v20190701.RemediationFilters>;
    /**
     * Management group ID.
     */
    readonly managementGroupId: pulumi.Input<string>;
    /**
     * The namespace for Microsoft Management RP; only "Microsoft.Management" is allowed.
     */
    readonly managementGroupsNamespace: pulumi.Input<string>;
    /**
     * The resource ID of the policy assignment that should be remediated.
     */
    readonly policyAssignmentId?: pulumi.Input<string>;
    /**
     * The policy definition reference ID of the individual definition that should be remediated. Required when the policy assignment being remediated assigns a policy set definition.
     */
    readonly policyDefinitionReferenceId?: pulumi.Input<string>;
    /**
     * The name of the remediation.
     */
    readonly remediationName: pulumi.Input<string>;
    /**
     * The way resources to remediate are discovered. Defaults to ExistingNonCompliant if not specified.
     */
    readonly resourceDiscoveryMode?: pulumi.Input<string>;
}
