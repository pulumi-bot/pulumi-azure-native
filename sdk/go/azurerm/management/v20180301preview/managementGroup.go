// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180301preview

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The management group details.
type ManagementGroup struct {
	pulumi.CustomResourceState

	// The list of children.
	Children ManagementGroupChildInfoResponseArrayOutput `pulumi:"children"`
	// The details of a management group.
	Details ManagementGroupDetailsResponsePtrOutput `pulumi:"details"`
	// The friendly name of the management group.
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// The name of the management group. For example, 00000000-0000-0000-0000-000000000000
	Name pulumi.StringOutput `pulumi:"name"`
	// The role definitions associated with the management group.
	Roles pulumi.StringArrayOutput `pulumi:"roles"`
	// The AAD Tenant ID associated with the management group. For example, 00000000-0000-0000-0000-000000000000
	TenantId pulumi.StringPtrOutput `pulumi:"tenantId"`
	// The type of the resource.  For example, /providers/Microsoft.Management/managementGroups
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewManagementGroup registers a new resource with the given unique name, arguments, and options.
func NewManagementGroup(ctx *pulumi.Context,
	name string, args *ManagementGroupArgs, opts ...pulumi.ResourceOption) (*ManagementGroup, error) {
	if args == nil || args.GroupId == nil {
		return nil, errors.New("missing required argument 'GroupId'")
	}
	if args == nil {
		args = &ManagementGroupArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:management/latest:ManagementGroup"),
		},
		{
			Type: pulumi.String("azurerm:management/v20171101preview:ManagementGroup"),
		},
		{
			Type: pulumi.String("azurerm:management/v20180101preview:ManagementGroup"),
		},
		{
			Type: pulumi.String("azurerm:management/v20191101:ManagementGroup"),
		},
		{
			Type: pulumi.String("azurerm:management/v20200201:ManagementGroup"),
		},
		{
			Type: pulumi.String("azurerm:management/v20200501:ManagementGroup"),
		},
	})
	opts = append(opts, aliases)
	var resource ManagementGroup
	err := ctx.RegisterResource("azurerm:management/v20180301preview:ManagementGroup", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azurerm:management/v20180301preview:ManagementGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ManagementGroup resources.
type managementGroupState struct {
	// The list of children.
	Children []ManagementGroupChildInfoResponse `pulumi:"children"`
	// The details of a management group.
	Details *ManagementGroupDetailsResponse `pulumi:"details"`
	// The friendly name of the management group.
	DisplayName *string `pulumi:"displayName"`
	// The name of the management group. For example, 00000000-0000-0000-0000-000000000000
	Name *string `pulumi:"name"`
	// The role definitions associated with the management group.
	Roles []string `pulumi:"roles"`
	// The AAD Tenant ID associated with the management group. For example, 00000000-0000-0000-0000-000000000000
	TenantId *string `pulumi:"tenantId"`
	// The type of the resource.  For example, /providers/Microsoft.Management/managementGroups
	Type *string `pulumi:"type"`
}

type ManagementGroupState struct {
	// The list of children.
	Children ManagementGroupChildInfoResponseArrayInput
	// The details of a management group.
	Details ManagementGroupDetailsResponsePtrInput
	// The friendly name of the management group.
	DisplayName pulumi.StringPtrInput
	// The name of the management group. For example, 00000000-0000-0000-0000-000000000000
	Name pulumi.StringPtrInput
	// The role definitions associated with the management group.
	Roles pulumi.StringArrayInput
	// The AAD Tenant ID associated with the management group. For example, 00000000-0000-0000-0000-000000000000
	TenantId pulumi.StringPtrInput
	// The type of the resource.  For example, /providers/Microsoft.Management/managementGroups
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
	GroupId string `pulumi:"groupId"`
	// The name of the management group. For example, 00000000-0000-0000-0000-000000000000
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a ManagementGroup resource.
type ManagementGroupArgs struct {
	// The details of a management group used during creation.
	Details CreateManagementGroupDetailsPtrInput
	// The friendly name of the management group. If no value is passed then this  field will be set to the groupId.
	DisplayName pulumi.StringPtrInput
	// Management Group ID.
	GroupId pulumi.StringInput
	// The name of the management group. For example, 00000000-0000-0000-0000-000000000000
	Name pulumi.StringPtrInput
}

func (ManagementGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managementGroupArgs)(nil)).Elem()
}
