// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180701preview

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The remediation definition.
type RemediationAtResourceGroup struct {
	pulumi.CustomResourceState

	// The time at which the remediation was created.
	CreatedOn pulumi.StringOutput `pulumi:"createdOn"`
	// The deployment status summary for all deployments created by the remediation.
	DeploymentStatus RemediationDeploymentSummaryResponsePtrOutput `pulumi:"deploymentStatus"`
	// The filters that will be applied to determine which resources to remediate.
	Filters RemediationFiltersResponsePtrOutput `pulumi:"filters"`
	// The time at which the remediation was last updated.
	LastUpdatedOn pulumi.StringOutput `pulumi:"lastUpdatedOn"`
	// The name of the remediation.
	Name pulumi.StringOutput `pulumi:"name"`
	// The resource ID of the policy assignment that should be remediated.
	PolicyAssignmentId pulumi.StringPtrOutput `pulumi:"policyAssignmentId"`
	// The policy definition reference ID of the individual definition that should be remediated. Required when the policy assignment being remediated assigns a policy set definition.
	PolicyDefinitionReferenceId pulumi.StringPtrOutput `pulumi:"policyDefinitionReferenceId"`
	// The status of the remediation.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The type of the remediation.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewRemediationAtResourceGroup registers a new resource with the given unique name, arguments, and options.
func NewRemediationAtResourceGroup(ctx *pulumi.Context,
	name string, args *RemediationAtResourceGroupArgs, opts ...pulumi.ResourceOption) (*RemediationAtResourceGroup, error) {
	if args == nil || args.RemediationName == nil {
		return nil, errors.New("missing required argument 'RemediationName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &RemediationAtResourceGroupArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:policyinsights/latest:RemediationAtResourceGroup"),
		},
		{
			Type: pulumi.String("azurerm:policyinsights/v20190701:RemediationAtResourceGroup"),
		},
	})
	opts = append(opts, aliases)
	var resource RemediationAtResourceGroup
	err := ctx.RegisterResource("azurerm:policyinsights/v20180701preview:RemediationAtResourceGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRemediationAtResourceGroup gets an existing RemediationAtResourceGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRemediationAtResourceGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RemediationAtResourceGroupState, opts ...pulumi.ResourceOption) (*RemediationAtResourceGroup, error) {
	var resource RemediationAtResourceGroup
	err := ctx.ReadResource("azurerm:policyinsights/v20180701preview:RemediationAtResourceGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RemediationAtResourceGroup resources.
type remediationAtResourceGroupState struct {
	// The time at which the remediation was created.
	CreatedOn *string `pulumi:"createdOn"`
	// The deployment status summary for all deployments created by the remediation.
	DeploymentStatus *RemediationDeploymentSummaryResponse `pulumi:"deploymentStatus"`
	// The filters that will be applied to determine which resources to remediate.
	Filters *RemediationFiltersResponse `pulumi:"filters"`
	// The time at which the remediation was last updated.
	LastUpdatedOn *string `pulumi:"lastUpdatedOn"`
	// The name of the remediation.
	Name *string `pulumi:"name"`
	// The resource ID of the policy assignment that should be remediated.
	PolicyAssignmentId *string `pulumi:"policyAssignmentId"`
	// The policy definition reference ID of the individual definition that should be remediated. Required when the policy assignment being remediated assigns a policy set definition.
	PolicyDefinitionReferenceId *string `pulumi:"policyDefinitionReferenceId"`
	// The status of the remediation.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The type of the remediation.
	Type *string `pulumi:"type"`
}

type RemediationAtResourceGroupState struct {
	// The time at which the remediation was created.
	CreatedOn pulumi.StringPtrInput
	// The deployment status summary for all deployments created by the remediation.
	DeploymentStatus RemediationDeploymentSummaryResponsePtrInput
	// The filters that will be applied to determine which resources to remediate.
	Filters RemediationFiltersResponsePtrInput
	// The time at which the remediation was last updated.
	LastUpdatedOn pulumi.StringPtrInput
	// The name of the remediation.
	Name pulumi.StringPtrInput
	// The resource ID of the policy assignment that should be remediated.
	PolicyAssignmentId pulumi.StringPtrInput
	// The policy definition reference ID of the individual definition that should be remediated. Required when the policy assignment being remediated assigns a policy set definition.
	PolicyDefinitionReferenceId pulumi.StringPtrInput
	// The status of the remediation.
	ProvisioningState pulumi.StringPtrInput
	// The type of the remediation.
	Type pulumi.StringPtrInput
}

func (RemediationAtResourceGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*remediationAtResourceGroupState)(nil)).Elem()
}

type remediationAtResourceGroupArgs struct {
	// The deployment status summary for all deployments created by the remediation.
	DeploymentStatus *RemediationDeploymentSummary `pulumi:"deploymentStatus"`
	// The filters that will be applied to determine which resources to remediate.
	Filters *RemediationFilters `pulumi:"filters"`
	// The resource ID of the policy assignment that should be remediated.
	PolicyAssignmentId *string `pulumi:"policyAssignmentId"`
	// The policy definition reference ID of the individual definition that should be remediated. Required when the policy assignment being remediated assigns a policy set definition.
	PolicyDefinitionReferenceId *string `pulumi:"policyDefinitionReferenceId"`
	// The name of the remediation.
	RemediationName string `pulumi:"remediationName"`
	// Resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a RemediationAtResourceGroup resource.
type RemediationAtResourceGroupArgs struct {
	// The deployment status summary for all deployments created by the remediation.
	DeploymentStatus RemediationDeploymentSummaryPtrInput
	// The filters that will be applied to determine which resources to remediate.
	Filters RemediationFiltersPtrInput
	// The resource ID of the policy assignment that should be remediated.
	PolicyAssignmentId pulumi.StringPtrInput
	// The policy definition reference ID of the individual definition that should be remediated. Required when the policy assignment being remediated assigns a policy set definition.
	PolicyDefinitionReferenceId pulumi.StringPtrInput
	// The name of the remediation.
	RemediationName pulumi.StringInput
	// Resource group name.
	ResourceGroupName pulumi.StringInput
}

func (RemediationAtResourceGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*remediationAtResourceGroupArgs)(nil)).Elem()
}
