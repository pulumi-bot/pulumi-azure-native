// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20151101

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Policy assignment.
type PolicyAssignment struct {
	pulumi.CustomResourceState

	// Gets or sets the policy assignment name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Gets or sets the policy assignment properties.
	Properties PolicyAssignmentPropertiesResponseOutput `pulumi:"properties"`
}

// NewPolicyAssignment registers a new resource with the given unique name, arguments, and options.
func NewPolicyAssignment(ctx *pulumi.Context,
	name string, args *PolicyAssignmentArgs, opts ...pulumi.ResourceOption) (*PolicyAssignment, error) {
	if args == nil || args.PolicyAssignmentName == nil {
		return nil, errors.New("missing required argument 'PolicyAssignmentName'")
	}
	if args == nil || args.Scope == nil {
		return nil, errors.New("missing required argument 'Scope'")
	}
	if args == nil {
		args = &PolicyAssignmentArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:authorization/latest:PolicyAssignment"),
		},
		{
			Type: pulumi.String("azurerm:authorization/v20151001preview:PolicyAssignment"),
		},
		{
			Type: pulumi.String("azurerm:authorization/v20160401:PolicyAssignment"),
		},
		{
			Type: pulumi.String("azurerm:authorization/v20161201:PolicyAssignment"),
		},
		{
			Type: pulumi.String("azurerm:authorization/v20170601preview:PolicyAssignment"),
		},
		{
			Type: pulumi.String("azurerm:authorization/v20180301:PolicyAssignment"),
		},
		{
			Type: pulumi.String("azurerm:authorization/v20180501:PolicyAssignment"),
		},
		{
			Type: pulumi.String("azurerm:authorization/v20190101:PolicyAssignment"),
		},
		{
			Type: pulumi.String("azurerm:authorization/v20190601:PolicyAssignment"),
		},
		{
			Type: pulumi.String("azurerm:authorization/v20190901:PolicyAssignment"),
		},
		{
			Type: pulumi.String("azurerm:authorization/v20200301:PolicyAssignment"),
		},
	})
	opts = append(opts, aliases)
	var resource PolicyAssignment
	err := ctx.RegisterResource("azurerm:authorization/v20151101:PolicyAssignment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPolicyAssignment gets an existing PolicyAssignment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPolicyAssignment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PolicyAssignmentState, opts ...pulumi.ResourceOption) (*PolicyAssignment, error) {
	var resource PolicyAssignment
	err := ctx.ReadResource("azurerm:authorization/v20151101:PolicyAssignment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PolicyAssignment resources.
type policyAssignmentState struct {
	// Gets or sets the policy assignment name.
	Name *string `pulumi:"name"`
	// Gets or sets the policy assignment properties.
	Properties *PolicyAssignmentPropertiesResponse `pulumi:"properties"`
}

type PolicyAssignmentState struct {
	// Gets or sets the policy assignment name.
	Name pulumi.StringPtrInput
	// Gets or sets the policy assignment properties.
	Properties PolicyAssignmentPropertiesResponsePtrInput
}

func (PolicyAssignmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*policyAssignmentState)(nil)).Elem()
}

type policyAssignmentArgs struct {
	// Policy assignment name.
	PolicyAssignmentName string `pulumi:"policyAssignmentName"`
	// Gets or sets the policy assignment properties.
	Properties *PolicyAssignmentProperties `pulumi:"properties"`
	// Scope.
	Scope string `pulumi:"scope"`
}

// The set of arguments for constructing a PolicyAssignment resource.
type PolicyAssignmentArgs struct {
	// Policy assignment name.
	PolicyAssignmentName pulumi.StringInput
	// Gets or sets the policy assignment properties.
	Properties PolicyAssignmentPropertiesPtrInput
	// Scope.
	Scope pulumi.StringInput
}

func (PolicyAssignmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*policyAssignmentArgs)(nil)).Elem()
}
