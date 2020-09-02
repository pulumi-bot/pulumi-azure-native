// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180301preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The details of a management group used during creation.
type CreateManagementGroupDetails struct {
	// (Optional) The ID of the parent management group used during creation.
	Parent *CreateParentGroupInfo `pulumi:"parent"`
}

// CreateManagementGroupDetailsInput is an input type that accepts CreateManagementGroupDetailsArgs and CreateManagementGroupDetailsOutput values.
// You can construct a concrete instance of `CreateManagementGroupDetailsInput` via:
//
//          CreateManagementGroupDetailsArgs{...}
type CreateManagementGroupDetailsInput interface {
	pulumi.Input

	ToCreateManagementGroupDetailsOutput() CreateManagementGroupDetailsOutput
	ToCreateManagementGroupDetailsOutputWithContext(context.Context) CreateManagementGroupDetailsOutput
}

// The details of a management group used during creation.
type CreateManagementGroupDetailsArgs struct {
	// (Optional) The ID of the parent management group used during creation.
	Parent CreateParentGroupInfoPtrInput `pulumi:"parent"`
}

func (CreateManagementGroupDetailsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CreateManagementGroupDetails)(nil)).Elem()
}

func (i CreateManagementGroupDetailsArgs) ToCreateManagementGroupDetailsOutput() CreateManagementGroupDetailsOutput {
	return i.ToCreateManagementGroupDetailsOutputWithContext(context.Background())
}

func (i CreateManagementGroupDetailsArgs) ToCreateManagementGroupDetailsOutputWithContext(ctx context.Context) CreateManagementGroupDetailsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CreateManagementGroupDetailsOutput)
}

func (i CreateManagementGroupDetailsArgs) ToCreateManagementGroupDetailsPtrOutput() CreateManagementGroupDetailsPtrOutput {
	return i.ToCreateManagementGroupDetailsPtrOutputWithContext(context.Background())
}

func (i CreateManagementGroupDetailsArgs) ToCreateManagementGroupDetailsPtrOutputWithContext(ctx context.Context) CreateManagementGroupDetailsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CreateManagementGroupDetailsOutput).ToCreateManagementGroupDetailsPtrOutputWithContext(ctx)
}

// CreateManagementGroupDetailsPtrInput is an input type that accepts CreateManagementGroupDetailsArgs, CreateManagementGroupDetailsPtr and CreateManagementGroupDetailsPtrOutput values.
// You can construct a concrete instance of `CreateManagementGroupDetailsPtrInput` via:
//
//          CreateManagementGroupDetailsArgs{...}
//
//  or:
//
//          nil
type CreateManagementGroupDetailsPtrInput interface {
	pulumi.Input

	ToCreateManagementGroupDetailsPtrOutput() CreateManagementGroupDetailsPtrOutput
	ToCreateManagementGroupDetailsPtrOutputWithContext(context.Context) CreateManagementGroupDetailsPtrOutput
}

type createManagementGroupDetailsPtrType CreateManagementGroupDetailsArgs

func CreateManagementGroupDetailsPtr(v *CreateManagementGroupDetailsArgs) CreateManagementGroupDetailsPtrInput {
	return (*createManagementGroupDetailsPtrType)(v)
}

func (*createManagementGroupDetailsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CreateManagementGroupDetails)(nil)).Elem()
}

func (i *createManagementGroupDetailsPtrType) ToCreateManagementGroupDetailsPtrOutput() CreateManagementGroupDetailsPtrOutput {
	return i.ToCreateManagementGroupDetailsPtrOutputWithContext(context.Background())
}

func (i *createManagementGroupDetailsPtrType) ToCreateManagementGroupDetailsPtrOutputWithContext(ctx context.Context) CreateManagementGroupDetailsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CreateManagementGroupDetailsPtrOutput)
}

// The details of a management group used during creation.
type CreateManagementGroupDetailsOutput struct{ *pulumi.OutputState }

func (CreateManagementGroupDetailsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CreateManagementGroupDetails)(nil)).Elem()
}

func (o CreateManagementGroupDetailsOutput) ToCreateManagementGroupDetailsOutput() CreateManagementGroupDetailsOutput {
	return o
}

func (o CreateManagementGroupDetailsOutput) ToCreateManagementGroupDetailsOutputWithContext(ctx context.Context) CreateManagementGroupDetailsOutput {
	return o
}

func (o CreateManagementGroupDetailsOutput) ToCreateManagementGroupDetailsPtrOutput() CreateManagementGroupDetailsPtrOutput {
	return o.ToCreateManagementGroupDetailsPtrOutputWithContext(context.Background())
}

func (o CreateManagementGroupDetailsOutput) ToCreateManagementGroupDetailsPtrOutputWithContext(ctx context.Context) CreateManagementGroupDetailsPtrOutput {
	return o.ApplyT(func(v CreateManagementGroupDetails) *CreateManagementGroupDetails {
		return &v
	}).(CreateManagementGroupDetailsPtrOutput)
}

// (Optional) The ID of the parent management group used during creation.
func (o CreateManagementGroupDetailsOutput) Parent() CreateParentGroupInfoPtrOutput {
	return o.ApplyT(func(v CreateManagementGroupDetails) *CreateParentGroupInfo { return v.Parent }).(CreateParentGroupInfoPtrOutput)
}

type CreateManagementGroupDetailsPtrOutput struct{ *pulumi.OutputState }

func (CreateManagementGroupDetailsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CreateManagementGroupDetails)(nil)).Elem()
}

func (o CreateManagementGroupDetailsPtrOutput) ToCreateManagementGroupDetailsPtrOutput() CreateManagementGroupDetailsPtrOutput {
	return o
}

func (o CreateManagementGroupDetailsPtrOutput) ToCreateManagementGroupDetailsPtrOutputWithContext(ctx context.Context) CreateManagementGroupDetailsPtrOutput {
	return o
}

func (o CreateManagementGroupDetailsPtrOutput) Elem() CreateManagementGroupDetailsOutput {
	return o.ApplyT(func(v *CreateManagementGroupDetails) CreateManagementGroupDetails { return *v }).(CreateManagementGroupDetailsOutput)
}

// (Optional) The ID of the parent management group used during creation.
func (o CreateManagementGroupDetailsPtrOutput) Parent() CreateParentGroupInfoPtrOutput {
	return o.ApplyT(func(v *CreateManagementGroupDetails) *CreateParentGroupInfo {
		if v == nil {
			return nil
		}
		return v.Parent
	}).(CreateParentGroupInfoPtrOutput)
}

// (Optional) The ID of the parent management group used during creation.
type CreateParentGroupInfo struct {
	// The fully qualified ID for the parent management group.  For example, /providers/Microsoft.Management/managementGroups/0000000-0000-0000-0000-000000000000
	Id *string `pulumi:"id"`
}

// CreateParentGroupInfoInput is an input type that accepts CreateParentGroupInfoArgs and CreateParentGroupInfoOutput values.
// You can construct a concrete instance of `CreateParentGroupInfoInput` via:
//
//          CreateParentGroupInfoArgs{...}
type CreateParentGroupInfoInput interface {
	pulumi.Input

	ToCreateParentGroupInfoOutput() CreateParentGroupInfoOutput
	ToCreateParentGroupInfoOutputWithContext(context.Context) CreateParentGroupInfoOutput
}

// (Optional) The ID of the parent management group used during creation.
type CreateParentGroupInfoArgs struct {
	// The fully qualified ID for the parent management group.  For example, /providers/Microsoft.Management/managementGroups/0000000-0000-0000-0000-000000000000
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (CreateParentGroupInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CreateParentGroupInfo)(nil)).Elem()
}

func (i CreateParentGroupInfoArgs) ToCreateParentGroupInfoOutput() CreateParentGroupInfoOutput {
	return i.ToCreateParentGroupInfoOutputWithContext(context.Background())
}

func (i CreateParentGroupInfoArgs) ToCreateParentGroupInfoOutputWithContext(ctx context.Context) CreateParentGroupInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CreateParentGroupInfoOutput)
}

func (i CreateParentGroupInfoArgs) ToCreateParentGroupInfoPtrOutput() CreateParentGroupInfoPtrOutput {
	return i.ToCreateParentGroupInfoPtrOutputWithContext(context.Background())
}

func (i CreateParentGroupInfoArgs) ToCreateParentGroupInfoPtrOutputWithContext(ctx context.Context) CreateParentGroupInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CreateParentGroupInfoOutput).ToCreateParentGroupInfoPtrOutputWithContext(ctx)
}

// CreateParentGroupInfoPtrInput is an input type that accepts CreateParentGroupInfoArgs, CreateParentGroupInfoPtr and CreateParentGroupInfoPtrOutput values.
// You can construct a concrete instance of `CreateParentGroupInfoPtrInput` via:
//
//          CreateParentGroupInfoArgs{...}
//
//  or:
//
//          nil
type CreateParentGroupInfoPtrInput interface {
	pulumi.Input

	ToCreateParentGroupInfoPtrOutput() CreateParentGroupInfoPtrOutput
	ToCreateParentGroupInfoPtrOutputWithContext(context.Context) CreateParentGroupInfoPtrOutput
}

type createParentGroupInfoPtrType CreateParentGroupInfoArgs

func CreateParentGroupInfoPtr(v *CreateParentGroupInfoArgs) CreateParentGroupInfoPtrInput {
	return (*createParentGroupInfoPtrType)(v)
}

func (*createParentGroupInfoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CreateParentGroupInfo)(nil)).Elem()
}

func (i *createParentGroupInfoPtrType) ToCreateParentGroupInfoPtrOutput() CreateParentGroupInfoPtrOutput {
	return i.ToCreateParentGroupInfoPtrOutputWithContext(context.Background())
}

func (i *createParentGroupInfoPtrType) ToCreateParentGroupInfoPtrOutputWithContext(ctx context.Context) CreateParentGroupInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CreateParentGroupInfoPtrOutput)
}

// (Optional) The ID of the parent management group used during creation.
type CreateParentGroupInfoOutput struct{ *pulumi.OutputState }

func (CreateParentGroupInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CreateParentGroupInfo)(nil)).Elem()
}

func (o CreateParentGroupInfoOutput) ToCreateParentGroupInfoOutput() CreateParentGroupInfoOutput {
	return o
}

func (o CreateParentGroupInfoOutput) ToCreateParentGroupInfoOutputWithContext(ctx context.Context) CreateParentGroupInfoOutput {
	return o
}

func (o CreateParentGroupInfoOutput) ToCreateParentGroupInfoPtrOutput() CreateParentGroupInfoPtrOutput {
	return o.ToCreateParentGroupInfoPtrOutputWithContext(context.Background())
}

func (o CreateParentGroupInfoOutput) ToCreateParentGroupInfoPtrOutputWithContext(ctx context.Context) CreateParentGroupInfoPtrOutput {
	return o.ApplyT(func(v CreateParentGroupInfo) *CreateParentGroupInfo {
		return &v
	}).(CreateParentGroupInfoPtrOutput)
}

// The fully qualified ID for the parent management group.  For example, /providers/Microsoft.Management/managementGroups/0000000-0000-0000-0000-000000000000
func (o CreateParentGroupInfoOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CreateParentGroupInfo) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type CreateParentGroupInfoPtrOutput struct{ *pulumi.OutputState }

func (CreateParentGroupInfoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CreateParentGroupInfo)(nil)).Elem()
}

func (o CreateParentGroupInfoPtrOutput) ToCreateParentGroupInfoPtrOutput() CreateParentGroupInfoPtrOutput {
	return o
}

func (o CreateParentGroupInfoPtrOutput) ToCreateParentGroupInfoPtrOutputWithContext(ctx context.Context) CreateParentGroupInfoPtrOutput {
	return o
}

func (o CreateParentGroupInfoPtrOutput) Elem() CreateParentGroupInfoOutput {
	return o.ApplyT(func(v *CreateParentGroupInfo) CreateParentGroupInfo { return *v }).(CreateParentGroupInfoOutput)
}

// The fully qualified ID for the parent management group.  For example, /providers/Microsoft.Management/managementGroups/0000000-0000-0000-0000-000000000000
func (o CreateParentGroupInfoPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CreateParentGroupInfo) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

// The child information of a management group.
type ManagementGroupChildInfoResponse struct {
	// The list of children.
	Children []ManagementGroupChildInfoResponse `pulumi:"children"`
	// The friendly name of the child resource.
	DisplayName *string `pulumi:"displayName"`
	// The fully qualified ID for the child resource (management group or subscription).  For example, /providers/Microsoft.Management/managementGroups/0000000-0000-0000-0000-000000000000
	Id *string `pulumi:"id"`
	// The name of the child entity.
	Name *string `pulumi:"name"`
	// The roles definitions associated with the management group.
	Roles []string `pulumi:"roles"`
	// The fully qualified resource type which includes provider namespace (e.g. /providers/Microsoft.Management/managementGroups)
	Type *string `pulumi:"type"`
}

// ManagementGroupChildInfoResponseInput is an input type that accepts ManagementGroupChildInfoResponseArgs and ManagementGroupChildInfoResponseOutput values.
// You can construct a concrete instance of `ManagementGroupChildInfoResponseInput` via:
//
//          ManagementGroupChildInfoResponseArgs{...}
type ManagementGroupChildInfoResponseInput interface {
	pulumi.Input

	ToManagementGroupChildInfoResponseOutput() ManagementGroupChildInfoResponseOutput
	ToManagementGroupChildInfoResponseOutputWithContext(context.Context) ManagementGroupChildInfoResponseOutput
}

// The child information of a management group.
type ManagementGroupChildInfoResponseArgs struct {
	// The list of children.
	Children ManagementGroupChildInfoResponseArrayInput `pulumi:"children"`
	// The friendly name of the child resource.
	DisplayName pulumi.StringPtrInput `pulumi:"displayName"`
	// The fully qualified ID for the child resource (management group or subscription).  For example, /providers/Microsoft.Management/managementGroups/0000000-0000-0000-0000-000000000000
	Id pulumi.StringPtrInput `pulumi:"id"`
	// The name of the child entity.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The roles definitions associated with the management group.
	Roles pulumi.StringArrayInput `pulumi:"roles"`
	// The fully qualified resource type which includes provider namespace (e.g. /providers/Microsoft.Management/managementGroups)
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (ManagementGroupChildInfoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementGroupChildInfoResponse)(nil)).Elem()
}

func (i ManagementGroupChildInfoResponseArgs) ToManagementGroupChildInfoResponseOutput() ManagementGroupChildInfoResponseOutput {
	return i.ToManagementGroupChildInfoResponseOutputWithContext(context.Background())
}

func (i ManagementGroupChildInfoResponseArgs) ToManagementGroupChildInfoResponseOutputWithContext(ctx context.Context) ManagementGroupChildInfoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementGroupChildInfoResponseOutput)
}

// ManagementGroupChildInfoResponseArrayInput is an input type that accepts ManagementGroupChildInfoResponseArray and ManagementGroupChildInfoResponseArrayOutput values.
// You can construct a concrete instance of `ManagementGroupChildInfoResponseArrayInput` via:
//
//          ManagementGroupChildInfoResponseArray{ ManagementGroupChildInfoResponseArgs{...} }
type ManagementGroupChildInfoResponseArrayInput interface {
	pulumi.Input

	ToManagementGroupChildInfoResponseArrayOutput() ManagementGroupChildInfoResponseArrayOutput
	ToManagementGroupChildInfoResponseArrayOutputWithContext(context.Context) ManagementGroupChildInfoResponseArrayOutput
}

type ManagementGroupChildInfoResponseArray []ManagementGroupChildInfoResponseInput

func (ManagementGroupChildInfoResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ManagementGroupChildInfoResponse)(nil)).Elem()
}

func (i ManagementGroupChildInfoResponseArray) ToManagementGroupChildInfoResponseArrayOutput() ManagementGroupChildInfoResponseArrayOutput {
	return i.ToManagementGroupChildInfoResponseArrayOutputWithContext(context.Background())
}

func (i ManagementGroupChildInfoResponseArray) ToManagementGroupChildInfoResponseArrayOutputWithContext(ctx context.Context) ManagementGroupChildInfoResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementGroupChildInfoResponseArrayOutput)
}

// The child information of a management group.
type ManagementGroupChildInfoResponseOutput struct{ *pulumi.OutputState }

func (ManagementGroupChildInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementGroupChildInfoResponse)(nil)).Elem()
}

func (o ManagementGroupChildInfoResponseOutput) ToManagementGroupChildInfoResponseOutput() ManagementGroupChildInfoResponseOutput {
	return o
}

func (o ManagementGroupChildInfoResponseOutput) ToManagementGroupChildInfoResponseOutputWithContext(ctx context.Context) ManagementGroupChildInfoResponseOutput {
	return o
}

// The list of children.
func (o ManagementGroupChildInfoResponseOutput) Children() ManagementGroupChildInfoResponseArrayOutput {
	return o.ApplyT(func(v ManagementGroupChildInfoResponse) []ManagementGroupChildInfoResponse { return v.Children }).(ManagementGroupChildInfoResponseArrayOutput)
}

// The friendly name of the child resource.
func (o ManagementGroupChildInfoResponseOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagementGroupChildInfoResponse) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// The fully qualified ID for the child resource (management group or subscription).  For example, /providers/Microsoft.Management/managementGroups/0000000-0000-0000-0000-000000000000
func (o ManagementGroupChildInfoResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagementGroupChildInfoResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// The name of the child entity.
func (o ManagementGroupChildInfoResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagementGroupChildInfoResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// The roles definitions associated with the management group.
func (o ManagementGroupChildInfoResponseOutput) Roles() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ManagementGroupChildInfoResponse) []string { return v.Roles }).(pulumi.StringArrayOutput)
}

// The fully qualified resource type which includes provider namespace (e.g. /providers/Microsoft.Management/managementGroups)
func (o ManagementGroupChildInfoResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagementGroupChildInfoResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ManagementGroupChildInfoResponseArrayOutput struct{ *pulumi.OutputState }

func (ManagementGroupChildInfoResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ManagementGroupChildInfoResponse)(nil)).Elem()
}

func (o ManagementGroupChildInfoResponseArrayOutput) ToManagementGroupChildInfoResponseArrayOutput() ManagementGroupChildInfoResponseArrayOutput {
	return o
}

func (o ManagementGroupChildInfoResponseArrayOutput) ToManagementGroupChildInfoResponseArrayOutputWithContext(ctx context.Context) ManagementGroupChildInfoResponseArrayOutput {
	return o
}

func (o ManagementGroupChildInfoResponseArrayOutput) Index(i pulumi.IntInput) ManagementGroupChildInfoResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ManagementGroupChildInfoResponse {
		return vs[0].([]ManagementGroupChildInfoResponse)[vs[1].(int)]
	}).(ManagementGroupChildInfoResponseOutput)
}

// The details of a management group.
type ManagementGroupDetailsResponse struct {
	// (Optional) The ID of the parent management group.
	Parent *ParentGroupInfoResponse `pulumi:"parent"`
	// The identity of the principal or process that updated the object.
	UpdatedBy *string `pulumi:"updatedBy"`
	// The date and time when this object was last updated.
	UpdatedTime *string `pulumi:"updatedTime"`
	// The version number of the object.
	Version *float64 `pulumi:"version"`
}

// ManagementGroupDetailsResponseInput is an input type that accepts ManagementGroupDetailsResponseArgs and ManagementGroupDetailsResponseOutput values.
// You can construct a concrete instance of `ManagementGroupDetailsResponseInput` via:
//
//          ManagementGroupDetailsResponseArgs{...}
type ManagementGroupDetailsResponseInput interface {
	pulumi.Input

	ToManagementGroupDetailsResponseOutput() ManagementGroupDetailsResponseOutput
	ToManagementGroupDetailsResponseOutputWithContext(context.Context) ManagementGroupDetailsResponseOutput
}

// The details of a management group.
type ManagementGroupDetailsResponseArgs struct {
	// (Optional) The ID of the parent management group.
	Parent ParentGroupInfoResponsePtrInput `pulumi:"parent"`
	// The identity of the principal or process that updated the object.
	UpdatedBy pulumi.StringPtrInput `pulumi:"updatedBy"`
	// The date and time when this object was last updated.
	UpdatedTime pulumi.StringPtrInput `pulumi:"updatedTime"`
	// The version number of the object.
	Version pulumi.Float64PtrInput `pulumi:"version"`
}

func (ManagementGroupDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementGroupDetailsResponse)(nil)).Elem()
}

func (i ManagementGroupDetailsResponseArgs) ToManagementGroupDetailsResponseOutput() ManagementGroupDetailsResponseOutput {
	return i.ToManagementGroupDetailsResponseOutputWithContext(context.Background())
}

func (i ManagementGroupDetailsResponseArgs) ToManagementGroupDetailsResponseOutputWithContext(ctx context.Context) ManagementGroupDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementGroupDetailsResponseOutput)
}

func (i ManagementGroupDetailsResponseArgs) ToManagementGroupDetailsResponsePtrOutput() ManagementGroupDetailsResponsePtrOutput {
	return i.ToManagementGroupDetailsResponsePtrOutputWithContext(context.Background())
}

func (i ManagementGroupDetailsResponseArgs) ToManagementGroupDetailsResponsePtrOutputWithContext(ctx context.Context) ManagementGroupDetailsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementGroupDetailsResponseOutput).ToManagementGroupDetailsResponsePtrOutputWithContext(ctx)
}

// ManagementGroupDetailsResponsePtrInput is an input type that accepts ManagementGroupDetailsResponseArgs, ManagementGroupDetailsResponsePtr and ManagementGroupDetailsResponsePtrOutput values.
// You can construct a concrete instance of `ManagementGroupDetailsResponsePtrInput` via:
//
//          ManagementGroupDetailsResponseArgs{...}
//
//  or:
//
//          nil
type ManagementGroupDetailsResponsePtrInput interface {
	pulumi.Input

	ToManagementGroupDetailsResponsePtrOutput() ManagementGroupDetailsResponsePtrOutput
	ToManagementGroupDetailsResponsePtrOutputWithContext(context.Context) ManagementGroupDetailsResponsePtrOutput
}

type managementGroupDetailsResponsePtrType ManagementGroupDetailsResponseArgs

func ManagementGroupDetailsResponsePtr(v *ManagementGroupDetailsResponseArgs) ManagementGroupDetailsResponsePtrInput {
	return (*managementGroupDetailsResponsePtrType)(v)
}

func (*managementGroupDetailsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagementGroupDetailsResponse)(nil)).Elem()
}

func (i *managementGroupDetailsResponsePtrType) ToManagementGroupDetailsResponsePtrOutput() ManagementGroupDetailsResponsePtrOutput {
	return i.ToManagementGroupDetailsResponsePtrOutputWithContext(context.Background())
}

func (i *managementGroupDetailsResponsePtrType) ToManagementGroupDetailsResponsePtrOutputWithContext(ctx context.Context) ManagementGroupDetailsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementGroupDetailsResponsePtrOutput)
}

// The details of a management group.
type ManagementGroupDetailsResponseOutput struct{ *pulumi.OutputState }

func (ManagementGroupDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagementGroupDetailsResponse)(nil)).Elem()
}

func (o ManagementGroupDetailsResponseOutput) ToManagementGroupDetailsResponseOutput() ManagementGroupDetailsResponseOutput {
	return o
}

func (o ManagementGroupDetailsResponseOutput) ToManagementGroupDetailsResponseOutputWithContext(ctx context.Context) ManagementGroupDetailsResponseOutput {
	return o
}

func (o ManagementGroupDetailsResponseOutput) ToManagementGroupDetailsResponsePtrOutput() ManagementGroupDetailsResponsePtrOutput {
	return o.ToManagementGroupDetailsResponsePtrOutputWithContext(context.Background())
}

func (o ManagementGroupDetailsResponseOutput) ToManagementGroupDetailsResponsePtrOutputWithContext(ctx context.Context) ManagementGroupDetailsResponsePtrOutput {
	return o.ApplyT(func(v ManagementGroupDetailsResponse) *ManagementGroupDetailsResponse {
		return &v
	}).(ManagementGroupDetailsResponsePtrOutput)
}

// (Optional) The ID of the parent management group.
func (o ManagementGroupDetailsResponseOutput) Parent() ParentGroupInfoResponsePtrOutput {
	return o.ApplyT(func(v ManagementGroupDetailsResponse) *ParentGroupInfoResponse { return v.Parent }).(ParentGroupInfoResponsePtrOutput)
}

// The identity of the principal or process that updated the object.
func (o ManagementGroupDetailsResponseOutput) UpdatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagementGroupDetailsResponse) *string { return v.UpdatedBy }).(pulumi.StringPtrOutput)
}

// The date and time when this object was last updated.
func (o ManagementGroupDetailsResponseOutput) UpdatedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagementGroupDetailsResponse) *string { return v.UpdatedTime }).(pulumi.StringPtrOutput)
}

// The version number of the object.
func (o ManagementGroupDetailsResponseOutput) Version() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ManagementGroupDetailsResponse) *float64 { return v.Version }).(pulumi.Float64PtrOutput)
}

type ManagementGroupDetailsResponsePtrOutput struct{ *pulumi.OutputState }

func (ManagementGroupDetailsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagementGroupDetailsResponse)(nil)).Elem()
}

func (o ManagementGroupDetailsResponsePtrOutput) ToManagementGroupDetailsResponsePtrOutput() ManagementGroupDetailsResponsePtrOutput {
	return o
}

func (o ManagementGroupDetailsResponsePtrOutput) ToManagementGroupDetailsResponsePtrOutputWithContext(ctx context.Context) ManagementGroupDetailsResponsePtrOutput {
	return o
}

func (o ManagementGroupDetailsResponsePtrOutput) Elem() ManagementGroupDetailsResponseOutput {
	return o.ApplyT(func(v *ManagementGroupDetailsResponse) ManagementGroupDetailsResponse { return *v }).(ManagementGroupDetailsResponseOutput)
}

// (Optional) The ID of the parent management group.
func (o ManagementGroupDetailsResponsePtrOutput) Parent() ParentGroupInfoResponsePtrOutput {
	return o.ApplyT(func(v *ManagementGroupDetailsResponse) *ParentGroupInfoResponse {
		if v == nil {
			return nil
		}
		return v.Parent
	}).(ParentGroupInfoResponsePtrOutput)
}

// The identity of the principal or process that updated the object.
func (o ManagementGroupDetailsResponsePtrOutput) UpdatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagementGroupDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return v.UpdatedBy
	}).(pulumi.StringPtrOutput)
}

// The date and time when this object was last updated.
func (o ManagementGroupDetailsResponsePtrOutput) UpdatedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagementGroupDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return v.UpdatedTime
	}).(pulumi.StringPtrOutput)
}

// The version number of the object.
func (o ManagementGroupDetailsResponsePtrOutput) Version() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ManagementGroupDetailsResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.Version
	}).(pulumi.Float64PtrOutput)
}

// (Optional) The ID of the parent management group.
type ParentGroupInfoResponse struct {
	// The friendly name of the parent management group.
	DisplayName *string `pulumi:"displayName"`
	// The fully qualified ID for the parent management group.  For example, /providers/Microsoft.Management/managementGroups/0000000-0000-0000-0000-000000000000
	Id *string `pulumi:"id"`
	// The name of the parent management group
	Name *string `pulumi:"name"`
}

// ParentGroupInfoResponseInput is an input type that accepts ParentGroupInfoResponseArgs and ParentGroupInfoResponseOutput values.
// You can construct a concrete instance of `ParentGroupInfoResponseInput` via:
//
//          ParentGroupInfoResponseArgs{...}
type ParentGroupInfoResponseInput interface {
	pulumi.Input

	ToParentGroupInfoResponseOutput() ParentGroupInfoResponseOutput
	ToParentGroupInfoResponseOutputWithContext(context.Context) ParentGroupInfoResponseOutput
}

// (Optional) The ID of the parent management group.
type ParentGroupInfoResponseArgs struct {
	// The friendly name of the parent management group.
	DisplayName pulumi.StringPtrInput `pulumi:"displayName"`
	// The fully qualified ID for the parent management group.  For example, /providers/Microsoft.Management/managementGroups/0000000-0000-0000-0000-000000000000
	Id pulumi.StringPtrInput `pulumi:"id"`
	// The name of the parent management group
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (ParentGroupInfoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ParentGroupInfoResponse)(nil)).Elem()
}

func (i ParentGroupInfoResponseArgs) ToParentGroupInfoResponseOutput() ParentGroupInfoResponseOutput {
	return i.ToParentGroupInfoResponseOutputWithContext(context.Background())
}

func (i ParentGroupInfoResponseArgs) ToParentGroupInfoResponseOutputWithContext(ctx context.Context) ParentGroupInfoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParentGroupInfoResponseOutput)
}

func (i ParentGroupInfoResponseArgs) ToParentGroupInfoResponsePtrOutput() ParentGroupInfoResponsePtrOutput {
	return i.ToParentGroupInfoResponsePtrOutputWithContext(context.Background())
}

func (i ParentGroupInfoResponseArgs) ToParentGroupInfoResponsePtrOutputWithContext(ctx context.Context) ParentGroupInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParentGroupInfoResponseOutput).ToParentGroupInfoResponsePtrOutputWithContext(ctx)
}

// ParentGroupInfoResponsePtrInput is an input type that accepts ParentGroupInfoResponseArgs, ParentGroupInfoResponsePtr and ParentGroupInfoResponsePtrOutput values.
// You can construct a concrete instance of `ParentGroupInfoResponsePtrInput` via:
//
//          ParentGroupInfoResponseArgs{...}
//
//  or:
//
//          nil
type ParentGroupInfoResponsePtrInput interface {
	pulumi.Input

	ToParentGroupInfoResponsePtrOutput() ParentGroupInfoResponsePtrOutput
	ToParentGroupInfoResponsePtrOutputWithContext(context.Context) ParentGroupInfoResponsePtrOutput
}

type parentGroupInfoResponsePtrType ParentGroupInfoResponseArgs

func ParentGroupInfoResponsePtr(v *ParentGroupInfoResponseArgs) ParentGroupInfoResponsePtrInput {
	return (*parentGroupInfoResponsePtrType)(v)
}

func (*parentGroupInfoResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ParentGroupInfoResponse)(nil)).Elem()
}

func (i *parentGroupInfoResponsePtrType) ToParentGroupInfoResponsePtrOutput() ParentGroupInfoResponsePtrOutput {
	return i.ToParentGroupInfoResponsePtrOutputWithContext(context.Background())
}

func (i *parentGroupInfoResponsePtrType) ToParentGroupInfoResponsePtrOutputWithContext(ctx context.Context) ParentGroupInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParentGroupInfoResponsePtrOutput)
}

// (Optional) The ID of the parent management group.
type ParentGroupInfoResponseOutput struct{ *pulumi.OutputState }

func (ParentGroupInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ParentGroupInfoResponse)(nil)).Elem()
}

func (o ParentGroupInfoResponseOutput) ToParentGroupInfoResponseOutput() ParentGroupInfoResponseOutput {
	return o
}

func (o ParentGroupInfoResponseOutput) ToParentGroupInfoResponseOutputWithContext(ctx context.Context) ParentGroupInfoResponseOutput {
	return o
}

func (o ParentGroupInfoResponseOutput) ToParentGroupInfoResponsePtrOutput() ParentGroupInfoResponsePtrOutput {
	return o.ToParentGroupInfoResponsePtrOutputWithContext(context.Background())
}

func (o ParentGroupInfoResponseOutput) ToParentGroupInfoResponsePtrOutputWithContext(ctx context.Context) ParentGroupInfoResponsePtrOutput {
	return o.ApplyT(func(v ParentGroupInfoResponse) *ParentGroupInfoResponse {
		return &v
	}).(ParentGroupInfoResponsePtrOutput)
}

// The friendly name of the parent management group.
func (o ParentGroupInfoResponseOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ParentGroupInfoResponse) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// The fully qualified ID for the parent management group.  For example, /providers/Microsoft.Management/managementGroups/0000000-0000-0000-0000-000000000000
func (o ParentGroupInfoResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ParentGroupInfoResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// The name of the parent management group
func (o ParentGroupInfoResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ParentGroupInfoResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type ParentGroupInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (ParentGroupInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ParentGroupInfoResponse)(nil)).Elem()
}

func (o ParentGroupInfoResponsePtrOutput) ToParentGroupInfoResponsePtrOutput() ParentGroupInfoResponsePtrOutput {
	return o
}

func (o ParentGroupInfoResponsePtrOutput) ToParentGroupInfoResponsePtrOutputWithContext(ctx context.Context) ParentGroupInfoResponsePtrOutput {
	return o
}

func (o ParentGroupInfoResponsePtrOutput) Elem() ParentGroupInfoResponseOutput {
	return o.ApplyT(func(v *ParentGroupInfoResponse) ParentGroupInfoResponse { return *v }).(ParentGroupInfoResponseOutput)
}

// The friendly name of the parent management group.
func (o ParentGroupInfoResponsePtrOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ParentGroupInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.DisplayName
	}).(pulumi.StringPtrOutput)
}

// The fully qualified ID for the parent management group.  For example, /providers/Microsoft.Management/managementGroups/0000000-0000-0000-0000-000000000000
func (o ParentGroupInfoResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ParentGroupInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

// The name of the parent management group
func (o ParentGroupInfoResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ParentGroupInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(CreateManagementGroupDetailsOutput{})
	pulumi.RegisterOutputType(CreateManagementGroupDetailsPtrOutput{})
	pulumi.RegisterOutputType(CreateParentGroupInfoOutput{})
	pulumi.RegisterOutputType(CreateParentGroupInfoPtrOutput{})
	pulumi.RegisterOutputType(ManagementGroupChildInfoResponseOutput{})
	pulumi.RegisterOutputType(ManagementGroupChildInfoResponseArrayOutput{})
	pulumi.RegisterOutputType(ManagementGroupDetailsResponseOutput{})
	pulumi.RegisterOutputType(ManagementGroupDetailsResponsePtrOutput{})
	pulumi.RegisterOutputType(ParentGroupInfoResponseOutput{})
	pulumi.RegisterOutputType(ParentGroupInfoResponsePtrOutput{})
}
