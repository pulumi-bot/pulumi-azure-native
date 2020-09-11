// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20150320

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Describes a storage account connection.
type StorageAccount struct {
	// The Azure Resource Manager ID of the storage account resource.
	Id string `pulumi:"id"`
	// The storage account key.
	Key string `pulumi:"key"`
}

// StorageAccountInput is an input type that accepts StorageAccountArgs and StorageAccountOutput values.
// You can construct a concrete instance of `StorageAccountInput` via:
//
//          StorageAccountArgs{...}
type StorageAccountInput interface {
	pulumi.Input

	ToStorageAccountOutput() StorageAccountOutput
	ToStorageAccountOutputWithContext(context.Context) StorageAccountOutput
}

// Describes a storage account connection.
type StorageAccountArgs struct {
	// The Azure Resource Manager ID of the storage account resource.
	Id pulumi.StringInput `pulumi:"id"`
	// The storage account key.
	Key pulumi.StringInput `pulumi:"key"`
}

func (StorageAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccount)(nil)).Elem()
}

func (i StorageAccountArgs) ToStorageAccountOutput() StorageAccountOutput {
	return i.ToStorageAccountOutputWithContext(context.Background())
}

func (i StorageAccountArgs) ToStorageAccountOutputWithContext(ctx context.Context) StorageAccountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageAccountOutput)
}

func (i StorageAccountArgs) ToStorageAccountPtrOutput() StorageAccountPtrOutput {
	return i.ToStorageAccountPtrOutputWithContext(context.Background())
}

func (i StorageAccountArgs) ToStorageAccountPtrOutputWithContext(ctx context.Context) StorageAccountPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageAccountOutput).ToStorageAccountPtrOutputWithContext(ctx)
}

// StorageAccountPtrInput is an input type that accepts StorageAccountArgs, StorageAccountPtr and StorageAccountPtrOutput values.
// You can construct a concrete instance of `StorageAccountPtrInput` via:
//
//          StorageAccountArgs{...}
//
//  or:
//
//          nil
type StorageAccountPtrInput interface {
	pulumi.Input

	ToStorageAccountPtrOutput() StorageAccountPtrOutput
	ToStorageAccountPtrOutputWithContext(context.Context) StorageAccountPtrOutput
}

type storageAccountPtrType StorageAccountArgs

func StorageAccountPtr(v *StorageAccountArgs) StorageAccountPtrInput {
	return (*storageAccountPtrType)(v)
}

func (*storageAccountPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageAccount)(nil)).Elem()
}

func (i *storageAccountPtrType) ToStorageAccountPtrOutput() StorageAccountPtrOutput {
	return i.ToStorageAccountPtrOutputWithContext(context.Background())
}

func (i *storageAccountPtrType) ToStorageAccountPtrOutputWithContext(ctx context.Context) StorageAccountPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageAccountPtrOutput)
}

// Describes a storage account connection.
type StorageAccountResponse struct {
	// The Azure Resource Manager ID of the storage account resource.
	Id string `pulumi:"id"`
	// The storage account key.
	Key string `pulumi:"key"`
}

// Describes a storage account connection.
type StorageAccountResponseOutput struct{ *pulumi.OutputState }

func (StorageAccountResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccountResponse)(nil)).Elem()
}

func (o StorageAccountResponseOutput) ToStorageAccountResponseOutput() StorageAccountResponseOutput {
	return o
}

func (o StorageAccountResponseOutput) ToStorageAccountResponseOutputWithContext(ctx context.Context) StorageAccountResponseOutput {
	return o
}

func (o StorageAccountResponseOutput) ToStorageAccountResponsePtrOutput() StorageAccountResponsePtrOutput {
	return o.ToStorageAccountResponsePtrOutputWithContext(context.Background())
}

func (o StorageAccountResponseOutput) ToStorageAccountResponsePtrOutputWithContext(ctx context.Context) StorageAccountResponsePtrOutput {
	return o.ApplyT(func(v StorageAccountResponse) *StorageAccountResponse {
		return &v
	}).(StorageAccountResponsePtrOutput)
}

// The Azure Resource Manager ID of the storage account resource.
func (o StorageAccountResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v StorageAccountResponse) string { return v.Id }).(pulumi.StringOutput)
}

// The storage account key.
func (o StorageAccountResponseOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v StorageAccountResponse) string { return v.Key }).(pulumi.StringOutput)
}

type StorageAccountResponsePtrOutput struct{ *pulumi.OutputState }

func (StorageAccountResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageAccountResponse)(nil)).Elem()
}

func (o StorageAccountResponsePtrOutput) ToStorageAccountResponsePtrOutput() StorageAccountResponsePtrOutput {
	return o
}

func (o StorageAccountResponsePtrOutput) ToStorageAccountResponsePtrOutputWithContext(ctx context.Context) StorageAccountResponsePtrOutput {
	return o
}

func (o StorageAccountResponsePtrOutput) Elem() StorageAccountResponseOutput {
	return o.ApplyT(func(v *StorageAccountResponse) StorageAccountResponse { return *v }).(StorageAccountResponseOutput)
}

// The Azure Resource Manager ID of the storage account resource.
func (o StorageAccountResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageAccountResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

// The storage account key.
func (o StorageAccountResponsePtrOutput) Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageAccountResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Key
	}).(pulumi.StringPtrOutput)
}

// The status of the storage insight.
type StorageInsightStatusResponse struct {
	// Description of the state of the storage insight.
	Description *string `pulumi:"description"`
	// The state of the storage insight connection to the workspace
	State string `pulumi:"state"`
}

// The status of the storage insight.
type StorageInsightStatusResponseOutput struct{ *pulumi.OutputState }

func (StorageInsightStatusResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageInsightStatusResponse)(nil)).Elem()
}

func (o StorageInsightStatusResponseOutput) ToStorageInsightStatusResponseOutput() StorageInsightStatusResponseOutput {
	return o
}

func (o StorageInsightStatusResponseOutput) ToStorageInsightStatusResponseOutputWithContext(ctx context.Context) StorageInsightStatusResponseOutput {
	return o
}

func (o StorageInsightStatusResponseOutput) ToStorageInsightStatusResponsePtrOutput() StorageInsightStatusResponsePtrOutput {
	return o.ToStorageInsightStatusResponsePtrOutputWithContext(context.Background())
}

func (o StorageInsightStatusResponseOutput) ToStorageInsightStatusResponsePtrOutputWithContext(ctx context.Context) StorageInsightStatusResponsePtrOutput {
	return o.ApplyT(func(v StorageInsightStatusResponse) *StorageInsightStatusResponse {
		return &v
	}).(StorageInsightStatusResponsePtrOutput)
}

// Description of the state of the storage insight.
func (o StorageInsightStatusResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageInsightStatusResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The state of the storage insight connection to the workspace
func (o StorageInsightStatusResponseOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v StorageInsightStatusResponse) string { return v.State }).(pulumi.StringOutput)
}

type StorageInsightStatusResponsePtrOutput struct{ *pulumi.OutputState }

func (StorageInsightStatusResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageInsightStatusResponse)(nil)).Elem()
}

func (o StorageInsightStatusResponsePtrOutput) ToStorageInsightStatusResponsePtrOutput() StorageInsightStatusResponsePtrOutput {
	return o
}

func (o StorageInsightStatusResponsePtrOutput) ToStorageInsightStatusResponsePtrOutputWithContext(ctx context.Context) StorageInsightStatusResponsePtrOutput {
	return o
}

func (o StorageInsightStatusResponsePtrOutput) Elem() StorageInsightStatusResponseOutput {
	return o.ApplyT(func(v *StorageInsightStatusResponse) StorageInsightStatusResponse { return *v }).(StorageInsightStatusResponseOutput)
}

// Description of the state of the storage insight.
func (o StorageInsightStatusResponsePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageInsightStatusResponse) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

// The state of the storage insight connection to the workspace
func (o StorageInsightStatusResponsePtrOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageInsightStatusResponse) *string {
		if v == nil {
			return nil
		}
		return &v.State
	}).(pulumi.StringPtrOutput)
}

// A tag of a saved search.
type Tag struct {
	// The tag name.
	Name string `pulumi:"name"`
	// The tag value.
	Value string `pulumi:"value"`
}

// TagInput is an input type that accepts TagArgs and TagOutput values.
// You can construct a concrete instance of `TagInput` via:
//
//          TagArgs{...}
type TagInput interface {
	pulumi.Input

	ToTagOutput() TagOutput
	ToTagOutputWithContext(context.Context) TagOutput
}

// A tag of a saved search.
type TagArgs struct {
	// The tag name.
	Name pulumi.StringInput `pulumi:"name"`
	// The tag value.
	Value pulumi.StringInput `pulumi:"value"`
}

func (TagArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Tag)(nil)).Elem()
}

func (i TagArgs) ToTagOutput() TagOutput {
	return i.ToTagOutputWithContext(context.Background())
}

func (i TagArgs) ToTagOutputWithContext(ctx context.Context) TagOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagOutput)
}

// TagArrayInput is an input type that accepts TagArray and TagArrayOutput values.
// You can construct a concrete instance of `TagArrayInput` via:
//
//          TagArray{ TagArgs{...} }
type TagArrayInput interface {
	pulumi.Input

	ToTagArrayOutput() TagArrayOutput
	ToTagArrayOutputWithContext(context.Context) TagArrayOutput
}

type TagArray []TagInput

func (TagArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Tag)(nil)).Elem()
}

func (i TagArray) ToTagArrayOutput() TagArrayOutput {
	return i.ToTagArrayOutputWithContext(context.Background())
}

func (i TagArray) ToTagArrayOutputWithContext(ctx context.Context) TagArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagArrayOutput)
}

// A tag of a saved search.
type TagResponse struct {
	// The tag name.
	Name string `pulumi:"name"`
	// The tag value.
	Value string `pulumi:"value"`
}

// A tag of a saved search.
type TagResponseOutput struct{ *pulumi.OutputState }

func (TagResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TagResponse)(nil)).Elem()
}

func (o TagResponseOutput) ToTagResponseOutput() TagResponseOutput {
	return o
}

func (o TagResponseOutput) ToTagResponseOutputWithContext(ctx context.Context) TagResponseOutput {
	return o
}

// The tag name.
func (o TagResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v TagResponse) string { return v.Name }).(pulumi.StringOutput)
}

// The tag value.
func (o TagResponseOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v TagResponse) string { return v.Value }).(pulumi.StringOutput)
}

type TagResponseArrayOutput struct{ *pulumi.OutputState }

func (TagResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TagResponse)(nil)).Elem()
}

func (o TagResponseArrayOutput) ToTagResponseArrayOutput() TagResponseArrayOutput {
	return o
}

func (o TagResponseArrayOutput) ToTagResponseArrayOutputWithContext(ctx context.Context) TagResponseArrayOutput {
	return o
}

func (o TagResponseArrayOutput) Index(i pulumi.IntInput) TagResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TagResponse {
		return vs[0].([]TagResponse)[vs[1].(int)]
	}).(TagResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(StorageAccountOutput{})
	pulumi.RegisterOutputType(StorageAccountPtrOutput{})
	pulumi.RegisterOutputType(StorageAccountResponseOutput{})
	pulumi.RegisterOutputType(StorageAccountResponsePtrOutput{})
	pulumi.RegisterOutputType(StorageInsightStatusResponseOutput{})
	pulumi.RegisterOutputType(StorageInsightStatusResponsePtrOutput{})
	pulumi.RegisterOutputType(TagOutput{})
	pulumi.RegisterOutputType(TagArrayOutput{})
	pulumi.RegisterOutputType(TagResponseOutput{})
	pulumi.RegisterOutputType(TagResponseArrayOutput{})
}
