// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20191101

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The management group details.
type ManagementGroup struct {
	pulumi.CustomResourceState

	// The name of the management group. For example, 00000000-0000-0000-0000-000000000000
	Name pulumi.StringOutput `pulumi:"name"`
	// The generic properties of a management group.
	Properties ManagementGroupPropertiesResponseOutput `pulumi:"properties"`
	// The type of the resource.  For example, Microsoft.Management/managementGroups
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewManagementGroup registers a new resource with the given unique name, arguments, and options.
func NewManagementGroup(ctx *pulumi.Context,
	name string, args *ManagementGroupArgs, opts ...pulumi.ResourceOption) (*ManagementGroup, error) {
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil {
		args = &ManagementGroupArgs{}
	}
	var resource ManagementGroup
	err := ctx.RegisterResource("azurerm:management/v20191101:ManagementGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetManagementGroup gets an existing ManagementGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetManagementGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagementGroupState, opts ...pulumi.ResourceOption) (*ManagementGroup, error) {
	var resource ManagementGroup
	err := ctx.ReadResource("azurerm:management/v20191101:ManagementGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ManagementGroup resources.
type managementGroupState struct {
	// The name of the management group. For example, 00000000-0000-0000-0000-000000000000
	Name *string `pulumi:"name"`
	// The generic properties of a management group.
	Properties *ManagementGroupPropertiesResponse `pulumi:"properties"`
	// The type of the resource.  For example, Microsoft.Management/managementGroups
	Type *string `pulumi:"type"`
}

type ManagementGroupState struct {
	// The name of the management group. For example, 00000000-0000-0000-0000-000000000000
	Name pulumi.StringPtrInput
	// The generic properties of a management group.
	Properties ManagementGroupPropertiesResponsePtrInput
	// The type of the resource.  For example, Microsoft.Management/managementGroups
	Type pulumi.StringPtrInput
}

func (ManagementGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*managementGroupState)(nil)).Elem()
}

type managementGroupArgs struct {
	// The details of a management group used during creation.
	Details *CreateManagementGroupDetails `pulumi:"details"`
	// The friendly name of the management group. If no value is passed then this  field will be set to the groupId.
	DisplayName *string `pulumi:"displayName"`
	// Management Group ID.
	Name string `pulumi:"name"`
}

// The set of arguments for constructing a ManagementGroup resource.
type ManagementGroupArgs struct {
	// The details of a management group used during creation.
	Details CreateManagementGroupDetailsPtrInput
	// The friendly name of the management group. If no value is passed then this  field will be set to the groupId.
	DisplayName pulumi.StringPtrInput
	// Management Group ID.
	Name pulumi.StringInput
}

func (ManagementGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managementGroupArgs)(nil)).Elem()
}
