// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Settings defined at the Management Group scope.
type HierarchySetting struct {
	pulumi.CustomResourceState

	// Settings that sets the default Management Group under which new subscriptions get added in this tenant. For example, /providers/Microsoft.Management/managementGroups/defaultGroup
	DefaultManagementGroup pulumi.StringPtrOutput `pulumi:"defaultManagementGroup"`
	// The name of the object. In this case, default.
	Name pulumi.StringOutput `pulumi:"name"`
	// Indicates whether RBAC access is required upon group creation under the root Management Group. If set to true, user will require Microsoft.Management/managementGroups/write action on the root Management Group scope in order to create new Groups directly under the root. This will prevent new users from creating new Management Groups, unless they are given access.
	RequireAuthorizationForGroupCreation pulumi.BoolPtrOutput `pulumi:"requireAuthorizationForGroupCreation"`
	// The AAD Tenant ID associated with the hierarchy settings. For example, 00000000-0000-0000-0000-000000000000
	TenantId pulumi.StringPtrOutput `pulumi:"tenantId"`
	// The type of the resource.  For example, Microsoft.Management/managementGroups/settings.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewHierarchySetting registers a new resource with the given unique name, arguments, and options.
func NewHierarchySetting(ctx *pulumi.Context,
	name string, args *HierarchySettingArgs, opts ...pulumi.ResourceOption) (*HierarchySetting, error) {
	if args == nil || args.GroupId == nil {
		return nil, errors.New("missing required argument 'GroupId'")
	}
	if args == nil {
		args = &HierarchySettingArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:management/v20200201:HierarchySetting"),
		},
		{
			Type: pulumi.String("azure-nextgen:management/v20200501:HierarchySetting"),
		},
	})
	opts = append(opts, aliases)
	var resource HierarchySetting
	err := ctx.RegisterResource("azure-nextgen:management/latest:HierarchySetting", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHierarchySetting gets an existing HierarchySetting resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHierarchySetting(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HierarchySettingState, opts ...pulumi.ResourceOption) (*HierarchySetting, error) {
	var resource HierarchySetting
	err := ctx.ReadResource("azure-nextgen:management/latest:HierarchySetting", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HierarchySetting resources.
type hierarchySettingState struct {
	// Settings that sets the default Management Group under which new subscriptions get added in this tenant. For example, /providers/Microsoft.Management/managementGroups/defaultGroup
	DefaultManagementGroup *string `pulumi:"defaultManagementGroup"`
	// The name of the object. In this case, default.
	Name *string `pulumi:"name"`
	// Indicates whether RBAC access is required upon group creation under the root Management Group. If set to true, user will require Microsoft.Management/managementGroups/write action on the root Management Group scope in order to create new Groups directly under the root. This will prevent new users from creating new Management Groups, unless they are given access.
	RequireAuthorizationForGroupCreation *bool `pulumi:"requireAuthorizationForGroupCreation"`
	// The AAD Tenant ID associated with the hierarchy settings. For example, 00000000-0000-0000-0000-000000000000
	TenantId *string `pulumi:"tenantId"`
	// The type of the resource.  For example, Microsoft.Management/managementGroups/settings.
	Type *string `pulumi:"type"`
}

type HierarchySettingState struct {
	// Settings that sets the default Management Group under which new subscriptions get added in this tenant. For example, /providers/Microsoft.Management/managementGroups/defaultGroup
	DefaultManagementGroup pulumi.StringPtrInput
	// The name of the object. In this case, default.
	Name pulumi.StringPtrInput
	// Indicates whether RBAC access is required upon group creation under the root Management Group. If set to true, user will require Microsoft.Management/managementGroups/write action on the root Management Group scope in order to create new Groups directly under the root. This will prevent new users from creating new Management Groups, unless they are given access.
	RequireAuthorizationForGroupCreation pulumi.BoolPtrInput
	// The AAD Tenant ID associated with the hierarchy settings. For example, 00000000-0000-0000-0000-000000000000
	TenantId pulumi.StringPtrInput
	// The type of the resource.  For example, Microsoft.Management/managementGroups/settings.
	Type pulumi.StringPtrInput
}

func (HierarchySettingState) ElementType() reflect.Type {
	return reflect.TypeOf((*hierarchySettingState)(nil)).Elem()
}

type hierarchySettingArgs struct {
	// Settings that sets the default Management Group under which new subscriptions get added in this tenant. For example, /providers/Microsoft.Management/managementGroups/defaultGroup
	DefaultManagementGroup *string `pulumi:"defaultManagementGroup"`
	// Management Group ID.
	GroupId string `pulumi:"groupId"`
	// Indicates whether RBAC access is required upon group creation under the root Management Group. If set to true, user will require Microsoft.Management/managementGroups/write action on the root Management Group scope in order to create new Groups directly under the root. This will prevent new users from creating new Management Groups, unless they are given access.
	RequireAuthorizationForGroupCreation *bool `pulumi:"requireAuthorizationForGroupCreation"`
}

// The set of arguments for constructing a HierarchySetting resource.
type HierarchySettingArgs struct {
	// Settings that sets the default Management Group under which new subscriptions get added in this tenant. For example, /providers/Microsoft.Management/managementGroups/defaultGroup
	DefaultManagementGroup pulumi.StringPtrInput
	// Management Group ID.
	GroupId pulumi.StringInput
	// Indicates whether RBAC access is required upon group creation under the root Management Group. If set to true, user will require Microsoft.Management/managementGroups/write action on the root Management Group scope in order to create new Groups directly under the root. This will prevent new users from creating new Management Groups, unless they are given access.
	RequireAuthorizationForGroupCreation pulumi.BoolPtrInput
}

func (HierarchySettingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hierarchySettingArgs)(nil)).Elem()
}

type HierarchySettingInput interface {
	pulumi.Input

	ToHierarchySettingOutput() HierarchySettingOutput
	ToHierarchySettingOutputWithContext(ctx context.Context) HierarchySettingOutput
}

func (HierarchySetting) ElementType() reflect.Type {
	return reflect.TypeOf((*HierarchySetting)(nil)).Elem()
}

func (i HierarchySetting) ToHierarchySettingOutput() HierarchySettingOutput {
	return i.ToHierarchySettingOutputWithContext(context.Background())
}

func (i HierarchySetting) ToHierarchySettingOutputWithContext(ctx context.Context) HierarchySettingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HierarchySettingOutput)
}

type HierarchySettingOutput struct {
	*pulumi.OutputState
}

func (HierarchySettingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HierarchySettingOutput)(nil)).Elem()
}

func (o HierarchySettingOutput) ToHierarchySettingOutput() HierarchySettingOutput {
	return o
}

func (o HierarchySettingOutput) ToHierarchySettingOutputWithContext(ctx context.Context) HierarchySettingOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(HierarchySettingOutput{})
}
