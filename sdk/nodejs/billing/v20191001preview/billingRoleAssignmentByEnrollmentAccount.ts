// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * The role assignment
 */
export class BillingRoleAssignmentByEnrollmentAccount extends pulumi.CustomResource {
    /**
     * Get an existing BillingRoleAssignmentByEnrollmentAccount resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): BillingRoleAssignmentByEnrollmentAccount {
        return new BillingRoleAssignmentByEnrollmentAccount(name, undefined, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'azurerm:billing/v20191001preview:BillingRoleAssignmentByEnrollmentAccount';

    /**
     * Returns true if the given object is an instance of BillingRoleAssignmentByEnrollmentAccount.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BillingRoleAssignmentByEnrollmentAccount {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BillingRoleAssignmentByEnrollmentAccount.__pulumiType;
    }

    /**
     * The principal Id of the user who created the role assignment.
     */
    public /*out*/ readonly createdByPrincipalId!: pulumi.Output<string>;
    /**
     * The tenant Id of the user who created the role assignment.
     */
    public /*out*/ readonly createdByPrincipalTenantId!: pulumi.Output<string>;
    /**
     * The email address of the user who created the role assignment. This is supported only for billing accounts with agreement type Enterprise Agreement.
     */
    public /*out*/ readonly createdByUserEmailAddress!: pulumi.Output<string>;
    /**
     * The date the role assignment was created.
     */
    public /*out*/ readonly createdOn!: pulumi.Output<string>;
    /**
     * Resource name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * The principal id of the user to whom the role was assigned.
     */
    public readonly principalId!: pulumi.Output<string | undefined>;
    /**
     * The principal tenant id of the user to whom the role was assigned.
     */
    public readonly principalTenantId!: pulumi.Output<string | undefined>;
    /**
     * The ID of the role definition.
     */
    public readonly roleDefinitionId!: pulumi.Output<string | undefined>;
    /**
     * The scope at which the role was assigned.
     */
    public /*out*/ readonly scope!: pulumi.Output<string>;
    /**
     * Resource type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * The authentication type of the user, whether Organization or MSA, of the user to whom the role was assigned. This is supported only for billing accounts with agreement type Enterprise Agreement.
     */
    public readonly userAuthenticationType!: pulumi.Output<string | undefined>;
    /**
     * The email address of the user to whom the role was assigned. This is supported only for billing accounts with agreement type Enterprise Agreement.
     */
    public readonly userEmailAddress!: pulumi.Output<string | undefined>;

    /**
     * Create a BillingRoleAssignmentByEnrollmentAccount resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BillingRoleAssignmentByEnrollmentAccountArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, state: undefined, opts: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BillingRoleAssignmentByEnrollmentAccountArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (!(opts && opts.id)) {
            const args = argsOrState as BillingRoleAssignmentByEnrollmentAccountArgs | undefined;
            if (!args || args.billingAccountName === undefined) {
                throw new Error("Missing required property 'billingAccountName'");
            }
            if (!args || args.billingRoleAssignmentName === undefined) {
                throw new Error("Missing required property 'billingRoleAssignmentName'");
            }
            if (!args || args.enrollmentAccountName === undefined) {
                throw new Error("Missing required property 'enrollmentAccountName'");
            }
            inputs["billingAccountName"] = args ? args.billingAccountName : undefined;
            inputs["billingRoleAssignmentName"] = args ? args.billingRoleAssignmentName : undefined;
            inputs["enrollmentAccountName"] = args ? args.enrollmentAccountName : undefined;
            inputs["principalId"] = args ? args.principalId : undefined;
            inputs["principalTenantId"] = args ? args.principalTenantId : undefined;
            inputs["roleDefinitionId"] = args ? args.roleDefinitionId : undefined;
            inputs["userAuthenticationType"] = args ? args.userAuthenticationType : undefined;
            inputs["userEmailAddress"] = args ? args.userEmailAddress : undefined;
            inputs["createdByPrincipalId"] = undefined /*out*/;
            inputs["createdByPrincipalTenantId"] = undefined /*out*/;
            inputs["createdByUserEmailAddress"] = undefined /*out*/;
            inputs["createdOn"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["scope"] = undefined /*out*/;
            inputs["type"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(BillingRoleAssignmentByEnrollmentAccount.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a BillingRoleAssignmentByEnrollmentAccount resource.
 */
export interface BillingRoleAssignmentByEnrollmentAccountArgs {
    /**
     * The ID that uniquely identifies a billing account.
     */
    readonly billingAccountName: pulumi.Input<string>;
    /**
     * The ID that uniquely identifies a role assignment.
     */
    readonly billingRoleAssignmentName: pulumi.Input<string>;
    /**
     * The ID that uniquely identifies an enrollment account.
     */
    readonly enrollmentAccountName: pulumi.Input<string>;
    /**
     * The principal id of the user to whom the role was assigned.
     */
    readonly principalId?: pulumi.Input<string>;
    /**
     * The principal tenant id of the user to whom the role was assigned.
     */
    readonly principalTenantId?: pulumi.Input<string>;
    /**
     * The ID of the role definition.
     */
    readonly roleDefinitionId?: pulumi.Input<string>;
    /**
     * The authentication type of the user, whether Organization or MSA, of the user to whom the role was assigned. This is supported only for billing accounts with agreement type Enterprise Agreement.
     */
    readonly userAuthenticationType?: pulumi.Input<string>;
    /**
     * The email address of the user to whom the role was assigned. This is supported only for billing accounts with agreement type Enterprise Agreement.
     */
    readonly userEmailAddress?: pulumi.Input<string>;
}
