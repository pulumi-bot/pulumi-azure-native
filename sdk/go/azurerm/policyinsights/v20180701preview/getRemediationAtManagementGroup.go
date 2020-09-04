// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180701preview

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupRemediationAtManagementGroup(ctx *pulumi.Context, args *LookupRemediationAtManagementGroupArgs, opts ...pulumi.InvokeOption) (*LookupRemediationAtManagementGroupResult, error) {
	var rv LookupRemediationAtManagementGroupResult
	err := ctx.Invoke("azurerm:policyinsights/v20180701preview:getRemediationAtManagementGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRemediationAtManagementGroupArgs struct {
	// Management group ID.
	ManagementGroupId string `pulumi:"managementGroupId"`
	// The namespace for Microsoft Management RP; only "Microsoft.Management" is allowed.
	ManagementGroupsNamespace string `pulumi:"managementGroupsNamespace"`
	// The name of the remediation.
	RemediationName string `pulumi:"remediationName"`
}

// The remediation definition.
type LookupRemediationAtManagementGroupResult struct {
	// The time at which the remediation was created.
	CreatedOn string `pulumi:"createdOn"`
	// The deployment status summary for all deployments created by the remediation.
	DeploymentStatus *RemediationDeploymentSummaryResponse `pulumi:"deploymentStatus"`
	// The filters that will be applied to determine which resources to remediate.
	Filters *RemediationFiltersResponse `pulumi:"filters"`
	// The time at which the remediation was last updated.
	LastUpdatedOn string `pulumi:"lastUpdatedOn"`
	// The name of the remediation.
	Name string `pulumi:"name"`
	// The resource ID of the policy assignment that should be remediated.
	PolicyAssignmentId *string `pulumi:"policyAssignmentId"`
	// The policy definition reference ID of the individual definition that should be remediated. Required when the policy assignment being remediated assigns a policy set definition.
	PolicyDefinitionReferenceId *string `pulumi:"policyDefinitionReferenceId"`
	// The status of the remediation.
	ProvisioningState string `pulumi:"provisioningState"`
	// The type of the remediation.
	Type string `pulumi:"type"`
}
