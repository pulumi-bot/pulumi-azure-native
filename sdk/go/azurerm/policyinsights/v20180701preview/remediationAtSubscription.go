// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180701preview

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The remediation definition.
//
// ## Example Usage
// ### Create remediation at subscription scope
//
// ```go
// package main
//
// import (
// 	policyinsights "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/policyinsights/v20180701preview"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := policyinsights.NewRemediationAtSubscription(ctx, "remediationAtSubscription", &policyinsights.RemediationAtSubscriptionArgs{
// 			PolicyAssignmentId: pulumi.String("/subscriptions/35ee058e-5fa0-414c-8145-3ebb8d09b6e2/providers/microsoft.authorization/policyassignments/b101830944f246d8a14088c5"),
// 			RemediationName:    pulumi.String("storageRemediation"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
// ### Create remediation at subscription scope with all properties
//
// ```go
// package main
//
// import (
// 	policyinsights "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/policyinsights/v20180701preview"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := policyinsights.NewRemediationAtSubscription(ctx, "remediationAtSubscription", &policyinsights.RemediationAtSubscriptionArgs{
// 			Filters: &policyinsights.RemediationFiltersArgs{
// 				Locations: pulumi.StringArray{
// 					pulumi.String("eastus"),
// 					pulumi.String("westus"),
// 				},
// 			},
// 			PolicyAssignmentId:          pulumi.String("/subscriptions/35ee058e-5fa0-414c-8145-3ebb8d09b6e2/providers/microsoft.authorization/policyassignments/b101830944f246d8a14088c5"),
// 			PolicyDefinitionReferenceId: pulumi.String("8c8fa9e4"),
// 			RemediationName:             pulumi.String("storageRemediation"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
type RemediationAtSubscription struct {
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

// NewRemediationAtSubscription registers a new resource with the given unique name, arguments, and options.
func NewRemediationAtSubscription(ctx *pulumi.Context,
	name string, args *RemediationAtSubscriptionArgs, opts ...pulumi.ResourceOption) (*RemediationAtSubscription, error) {
	if args == nil || args.RemediationName == nil {
		return nil, errors.New("missing required argument 'RemediationName'")
	}
	if args == nil {
		args = &RemediationAtSubscriptionArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:policyinsights/latest:RemediationAtSubscription"),
		},
		{
			Type: pulumi.String("azurerm:policyinsights/v20190701:RemediationAtSubscription"),
		},
	})
	opts = append(opts, aliases)
	var resource RemediationAtSubscription
	err := ctx.RegisterResource("azurerm:policyinsights/v20180701preview:RemediationAtSubscription", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRemediationAtSubscription gets an existing RemediationAtSubscription resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRemediationAtSubscription(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RemediationAtSubscriptionState, opts ...pulumi.ResourceOption) (*RemediationAtSubscription, error) {
	var resource RemediationAtSubscription
	err := ctx.ReadResource("azurerm:policyinsights/v20180701preview:RemediationAtSubscription", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RemediationAtSubscription resources.
type remediationAtSubscriptionState struct {
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

type RemediationAtSubscriptionState struct {
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

func (RemediationAtSubscriptionState) ElementType() reflect.Type {
	return reflect.TypeOf((*remediationAtSubscriptionState)(nil)).Elem()
}

type remediationAtSubscriptionArgs struct {
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
}

// The set of arguments for constructing a RemediationAtSubscription resource.
type RemediationAtSubscriptionArgs struct {
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
}

func (RemediationAtSubscriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*remediationAtSubscriptionArgs)(nil)).Elem()
}
