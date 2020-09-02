// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20170228preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A key property for the reference data set. A reference data set can have multiple key properties.
type ReferenceDataSetKeyProperty struct {
	// The name of the key property.
	Name *string `pulumi:"name"`
	// The type of the key property.
	Type *string `pulumi:"type"`
}

// ReferenceDataSetKeyPropertyInput is an input type that accepts ReferenceDataSetKeyPropertyArgs and ReferenceDataSetKeyPropertyOutput values.
// You can construct a concrete instance of `ReferenceDataSetKeyPropertyInput` via:
//
//          ReferenceDataSetKeyPropertyArgs{...}
type ReferenceDataSetKeyPropertyInput interface {
	pulumi.Input

	ToReferenceDataSetKeyPropertyOutput() ReferenceDataSetKeyPropertyOutput
	ToReferenceDataSetKeyPropertyOutputWithContext(context.Context) ReferenceDataSetKeyPropertyOutput
}

// A key property for the reference data set. A reference data set can have multiple key properties.
type ReferenceDataSetKeyPropertyArgs struct {
	// The name of the key property.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The type of the key property.
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (ReferenceDataSetKeyPropertyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ReferenceDataSetKeyProperty)(nil)).Elem()
}

func (i ReferenceDataSetKeyPropertyArgs) ToReferenceDataSetKeyPropertyOutput() ReferenceDataSetKeyPropertyOutput {
	return i.ToReferenceDataSetKeyPropertyOutputWithContext(context.Background())
}

func (i ReferenceDataSetKeyPropertyArgs) ToReferenceDataSetKeyPropertyOutputWithContext(ctx context.Context) ReferenceDataSetKeyPropertyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReferenceDataSetKeyPropertyOutput)
}

// ReferenceDataSetKeyPropertyArrayInput is an input type that accepts ReferenceDataSetKeyPropertyArray and ReferenceDataSetKeyPropertyArrayOutput values.
// You can construct a concrete instance of `ReferenceDataSetKeyPropertyArrayInput` via:
//
//          ReferenceDataSetKeyPropertyArray{ ReferenceDataSetKeyPropertyArgs{...} }
type ReferenceDataSetKeyPropertyArrayInput interface {
	pulumi.Input

	ToReferenceDataSetKeyPropertyArrayOutput() ReferenceDataSetKeyPropertyArrayOutput
	ToReferenceDataSetKeyPropertyArrayOutputWithContext(context.Context) ReferenceDataSetKeyPropertyArrayOutput
}

type ReferenceDataSetKeyPropertyArray []ReferenceDataSetKeyPropertyInput

func (ReferenceDataSetKeyPropertyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ReferenceDataSetKeyProperty)(nil)).Elem()
}

func (i ReferenceDataSetKeyPropertyArray) ToReferenceDataSetKeyPropertyArrayOutput() ReferenceDataSetKeyPropertyArrayOutput {
	return i.ToReferenceDataSetKeyPropertyArrayOutputWithContext(context.Background())
}

func (i ReferenceDataSetKeyPropertyArray) ToReferenceDataSetKeyPropertyArrayOutputWithContext(ctx context.Context) ReferenceDataSetKeyPropertyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReferenceDataSetKeyPropertyArrayOutput)
}

// A key property for the reference data set. A reference data set can have multiple key properties.
type ReferenceDataSetKeyPropertyOutput struct{ *pulumi.OutputState }

func (ReferenceDataSetKeyPropertyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReferenceDataSetKeyProperty)(nil)).Elem()
}

func (o ReferenceDataSetKeyPropertyOutput) ToReferenceDataSetKeyPropertyOutput() ReferenceDataSetKeyPropertyOutput {
	return o
}

func (o ReferenceDataSetKeyPropertyOutput) ToReferenceDataSetKeyPropertyOutputWithContext(ctx context.Context) ReferenceDataSetKeyPropertyOutput {
	return o
}

// The name of the key property.
func (o ReferenceDataSetKeyPropertyOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReferenceDataSetKeyProperty) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// The type of the key property.
func (o ReferenceDataSetKeyPropertyOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReferenceDataSetKeyProperty) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ReferenceDataSetKeyPropertyArrayOutput struct{ *pulumi.OutputState }

func (ReferenceDataSetKeyPropertyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ReferenceDataSetKeyProperty)(nil)).Elem()
}

func (o ReferenceDataSetKeyPropertyArrayOutput) ToReferenceDataSetKeyPropertyArrayOutput() ReferenceDataSetKeyPropertyArrayOutput {
	return o
}

func (o ReferenceDataSetKeyPropertyArrayOutput) ToReferenceDataSetKeyPropertyArrayOutputWithContext(ctx context.Context) ReferenceDataSetKeyPropertyArrayOutput {
	return o
}

func (o ReferenceDataSetKeyPropertyArrayOutput) Index(i pulumi.IntInput) ReferenceDataSetKeyPropertyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ReferenceDataSetKeyProperty {
		return vs[0].([]ReferenceDataSetKeyProperty)[vs[1].(int)]
	}).(ReferenceDataSetKeyPropertyOutput)
}

// A key property for the reference data set. A reference data set can have multiple key properties.
type ReferenceDataSetKeyPropertyResponse struct {
	// The name of the key property.
	Name *string `pulumi:"name"`
	// The type of the key property.
	Type *string `pulumi:"type"`
}

// ReferenceDataSetKeyPropertyResponseInput is an input type that accepts ReferenceDataSetKeyPropertyResponseArgs and ReferenceDataSetKeyPropertyResponseOutput values.
// You can construct a concrete instance of `ReferenceDataSetKeyPropertyResponseInput` via:
//
//          ReferenceDataSetKeyPropertyResponseArgs{...}
type ReferenceDataSetKeyPropertyResponseInput interface {
	pulumi.Input

	ToReferenceDataSetKeyPropertyResponseOutput() ReferenceDataSetKeyPropertyResponseOutput
	ToReferenceDataSetKeyPropertyResponseOutputWithContext(context.Context) ReferenceDataSetKeyPropertyResponseOutput
}

// A key property for the reference data set. A reference data set can have multiple key properties.
type ReferenceDataSetKeyPropertyResponseArgs struct {
	// The name of the key property.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The type of the key property.
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (ReferenceDataSetKeyPropertyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ReferenceDataSetKeyPropertyResponse)(nil)).Elem()
}

func (i ReferenceDataSetKeyPropertyResponseArgs) ToReferenceDataSetKeyPropertyResponseOutput() ReferenceDataSetKeyPropertyResponseOutput {
	return i.ToReferenceDataSetKeyPropertyResponseOutputWithContext(context.Background())
}

func (i ReferenceDataSetKeyPropertyResponseArgs) ToReferenceDataSetKeyPropertyResponseOutputWithContext(ctx context.Context) ReferenceDataSetKeyPropertyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReferenceDataSetKeyPropertyResponseOutput)
}

// ReferenceDataSetKeyPropertyResponseArrayInput is an input type that accepts ReferenceDataSetKeyPropertyResponseArray and ReferenceDataSetKeyPropertyResponseArrayOutput values.
// You can construct a concrete instance of `ReferenceDataSetKeyPropertyResponseArrayInput` via:
//
//          ReferenceDataSetKeyPropertyResponseArray{ ReferenceDataSetKeyPropertyResponseArgs{...} }
type ReferenceDataSetKeyPropertyResponseArrayInput interface {
	pulumi.Input

	ToReferenceDataSetKeyPropertyResponseArrayOutput() ReferenceDataSetKeyPropertyResponseArrayOutput
	ToReferenceDataSetKeyPropertyResponseArrayOutputWithContext(context.Context) ReferenceDataSetKeyPropertyResponseArrayOutput
}

type ReferenceDataSetKeyPropertyResponseArray []ReferenceDataSetKeyPropertyResponseInput

func (ReferenceDataSetKeyPropertyResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ReferenceDataSetKeyPropertyResponse)(nil)).Elem()
}

func (i ReferenceDataSetKeyPropertyResponseArray) ToReferenceDataSetKeyPropertyResponseArrayOutput() ReferenceDataSetKeyPropertyResponseArrayOutput {
	return i.ToReferenceDataSetKeyPropertyResponseArrayOutputWithContext(context.Background())
}

func (i ReferenceDataSetKeyPropertyResponseArray) ToReferenceDataSetKeyPropertyResponseArrayOutputWithContext(ctx context.Context) ReferenceDataSetKeyPropertyResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReferenceDataSetKeyPropertyResponseArrayOutput)
}

// A key property for the reference data set. A reference data set can have multiple key properties.
type ReferenceDataSetKeyPropertyResponseOutput struct{ *pulumi.OutputState }

func (ReferenceDataSetKeyPropertyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReferenceDataSetKeyPropertyResponse)(nil)).Elem()
}

func (o ReferenceDataSetKeyPropertyResponseOutput) ToReferenceDataSetKeyPropertyResponseOutput() ReferenceDataSetKeyPropertyResponseOutput {
	return o
}

func (o ReferenceDataSetKeyPropertyResponseOutput) ToReferenceDataSetKeyPropertyResponseOutputWithContext(ctx context.Context) ReferenceDataSetKeyPropertyResponseOutput {
	return o
}

// The name of the key property.
func (o ReferenceDataSetKeyPropertyResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReferenceDataSetKeyPropertyResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// The type of the key property.
func (o ReferenceDataSetKeyPropertyResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ReferenceDataSetKeyPropertyResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ReferenceDataSetKeyPropertyResponseArrayOutput struct{ *pulumi.OutputState }

func (ReferenceDataSetKeyPropertyResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ReferenceDataSetKeyPropertyResponse)(nil)).Elem()
}

func (o ReferenceDataSetKeyPropertyResponseArrayOutput) ToReferenceDataSetKeyPropertyResponseArrayOutput() ReferenceDataSetKeyPropertyResponseArrayOutput {
	return o
}

func (o ReferenceDataSetKeyPropertyResponseArrayOutput) ToReferenceDataSetKeyPropertyResponseArrayOutputWithContext(ctx context.Context) ReferenceDataSetKeyPropertyResponseArrayOutput {
	return o
}

func (o ReferenceDataSetKeyPropertyResponseArrayOutput) Index(i pulumi.IntInput) ReferenceDataSetKeyPropertyResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ReferenceDataSetKeyPropertyResponse {
		return vs[0].([]ReferenceDataSetKeyPropertyResponse)[vs[1].(int)]
	}).(ReferenceDataSetKeyPropertyResponseOutput)
}

// The sku determines the capacity of the environment, the SLA (in queries-per-minute and total capacity), and the billing rate.
type Sku struct {
	// The capacity of the sku. This value can be changed to support scale out of environments after they have been created.
	Capacity int `pulumi:"capacity"`
	// The name of this SKU.
	Name string `pulumi:"name"`
}

// SkuInput is an input type that accepts SkuArgs and SkuOutput values.
// You can construct a concrete instance of `SkuInput` via:
//
//          SkuArgs{...}
type SkuInput interface {
	pulumi.Input

	ToSkuOutput() SkuOutput
	ToSkuOutputWithContext(context.Context) SkuOutput
}

// The sku determines the capacity of the environment, the SLA (in queries-per-minute and total capacity), and the billing rate.
type SkuArgs struct {
	// The capacity of the sku. This value can be changed to support scale out of environments after they have been created.
	Capacity pulumi.IntInput `pulumi:"capacity"`
	// The name of this SKU.
	Name pulumi.StringInput `pulumi:"name"`
}

func (SkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Sku)(nil)).Elem()
}

func (i SkuArgs) ToSkuOutput() SkuOutput {
	return i.ToSkuOutputWithContext(context.Background())
}

func (i SkuArgs) ToSkuOutputWithContext(ctx context.Context) SkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuOutput)
}

func (i SkuArgs) ToSkuPtrOutput() SkuPtrOutput {
	return i.ToSkuPtrOutputWithContext(context.Background())
}

func (i SkuArgs) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuOutput).ToSkuPtrOutputWithContext(ctx)
}

// SkuPtrInput is an input type that accepts SkuArgs, SkuPtr and SkuPtrOutput values.
// You can construct a concrete instance of `SkuPtrInput` via:
//
//          SkuArgs{...}
//
//  or:
//
//          nil
type SkuPtrInput interface {
	pulumi.Input

	ToSkuPtrOutput() SkuPtrOutput
	ToSkuPtrOutputWithContext(context.Context) SkuPtrOutput
}

type skuPtrType SkuArgs

func SkuPtr(v *SkuArgs) SkuPtrInput {
	return (*skuPtrType)(v)
}

func (*skuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Sku)(nil)).Elem()
}

func (i *skuPtrType) ToSkuPtrOutput() SkuPtrOutput {
	return i.ToSkuPtrOutputWithContext(context.Background())
}

func (i *skuPtrType) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuPtrOutput)
}

// The sku determines the capacity of the environment, the SLA (in queries-per-minute and total capacity), and the billing rate.
type SkuOutput struct{ *pulumi.OutputState }

func (SkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Sku)(nil)).Elem()
}

func (o SkuOutput) ToSkuOutput() SkuOutput {
	return o
}

func (o SkuOutput) ToSkuOutputWithContext(ctx context.Context) SkuOutput {
	return o
}

func (o SkuOutput) ToSkuPtrOutput() SkuPtrOutput {
	return o.ToSkuPtrOutputWithContext(context.Background())
}

func (o SkuOutput) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return o.ApplyT(func(v Sku) *Sku {
		return &v
	}).(SkuPtrOutput)
}

// The capacity of the sku. This value can be changed to support scale out of environments after they have been created.
func (o SkuOutput) Capacity() pulumi.IntOutput {
	return o.ApplyT(func(v Sku) int { return v.Capacity }).(pulumi.IntOutput)
}

// The name of this SKU.
func (o SkuOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v Sku) string { return v.Name }).(pulumi.StringOutput)
}

type SkuPtrOutput struct{ *pulumi.OutputState }

func (SkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Sku)(nil)).Elem()
}

func (o SkuPtrOutput) ToSkuPtrOutput() SkuPtrOutput {
	return o
}

func (o SkuPtrOutput) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return o
}

func (o SkuPtrOutput) Elem() SkuOutput {
	return o.ApplyT(func(v *Sku) Sku { return *v }).(SkuOutput)
}

// The capacity of the sku. This value can be changed to support scale out of environments after they have been created.
func (o SkuPtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Sku) *int {
		if v == nil {
			return nil
		}
		return &v.Capacity
	}).(pulumi.IntPtrOutput)
}

// The name of this SKU.
func (o SkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

// The sku determines the capacity of the environment, the SLA (in queries-per-minute and total capacity), and the billing rate.
type SkuResponse struct {
	// The capacity of the sku. This value can be changed to support scale out of environments after they have been created.
	Capacity int `pulumi:"capacity"`
	// The name of this SKU.
	Name string `pulumi:"name"`
}

// SkuResponseInput is an input type that accepts SkuResponseArgs and SkuResponseOutput values.
// You can construct a concrete instance of `SkuResponseInput` via:
//
//          SkuResponseArgs{...}
type SkuResponseInput interface {
	pulumi.Input

	ToSkuResponseOutput() SkuResponseOutput
	ToSkuResponseOutputWithContext(context.Context) SkuResponseOutput
}

// The sku determines the capacity of the environment, the SLA (in queries-per-minute and total capacity), and the billing rate.
type SkuResponseArgs struct {
	// The capacity of the sku. This value can be changed to support scale out of environments after they have been created.
	Capacity pulumi.IntInput `pulumi:"capacity"`
	// The name of this SKU.
	Name pulumi.StringInput `pulumi:"name"`
}

func (SkuResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuResponse)(nil)).Elem()
}

func (i SkuResponseArgs) ToSkuResponseOutput() SkuResponseOutput {
	return i.ToSkuResponseOutputWithContext(context.Background())
}

func (i SkuResponseArgs) ToSkuResponseOutputWithContext(ctx context.Context) SkuResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuResponseOutput)
}

func (i SkuResponseArgs) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return i.ToSkuResponsePtrOutputWithContext(context.Background())
}

func (i SkuResponseArgs) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuResponseOutput).ToSkuResponsePtrOutputWithContext(ctx)
}

// SkuResponsePtrInput is an input type that accepts SkuResponseArgs, SkuResponsePtr and SkuResponsePtrOutput values.
// You can construct a concrete instance of `SkuResponsePtrInput` via:
//
//          SkuResponseArgs{...}
//
//  or:
//
//          nil
type SkuResponsePtrInput interface {
	pulumi.Input

	ToSkuResponsePtrOutput() SkuResponsePtrOutput
	ToSkuResponsePtrOutputWithContext(context.Context) SkuResponsePtrOutput
}

type skuResponsePtrType SkuResponseArgs

func SkuResponsePtr(v *SkuResponseArgs) SkuResponsePtrInput {
	return (*skuResponsePtrType)(v)
}

func (*skuResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuResponse)(nil)).Elem()
}

func (i *skuResponsePtrType) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return i.ToSkuResponsePtrOutputWithContext(context.Background())
}

func (i *skuResponsePtrType) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuResponsePtrOutput)
}

// The sku determines the capacity of the environment, the SLA (in queries-per-minute and total capacity), and the billing rate.
type SkuResponseOutput struct{ *pulumi.OutputState }

func (SkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuResponse)(nil)).Elem()
}

func (o SkuResponseOutput) ToSkuResponseOutput() SkuResponseOutput {
	return o
}

func (o SkuResponseOutput) ToSkuResponseOutputWithContext(ctx context.Context) SkuResponseOutput {
	return o
}

func (o SkuResponseOutput) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return o.ToSkuResponsePtrOutputWithContext(context.Background())
}

func (o SkuResponseOutput) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return o.ApplyT(func(v SkuResponse) *SkuResponse {
		return &v
	}).(SkuResponsePtrOutput)
}

// The capacity of the sku. This value can be changed to support scale out of environments after they have been created.
func (o SkuResponseOutput) Capacity() pulumi.IntOutput {
	return o.ApplyT(func(v SkuResponse) int { return v.Capacity }).(pulumi.IntOutput)
}

// The name of this SKU.
func (o SkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

type SkuResponsePtrOutput struct{ *pulumi.OutputState }

func (SkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuResponse)(nil)).Elem()
}

func (o SkuResponsePtrOutput) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return o
}

func (o SkuResponsePtrOutput) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return o
}

func (o SkuResponsePtrOutput) Elem() SkuResponseOutput {
	return o.ApplyT(func(v *SkuResponse) SkuResponse { return *v }).(SkuResponseOutput)
}

// The capacity of the sku. This value can be changed to support scale out of environments after they have been created.
func (o SkuResponsePtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *int {
		if v == nil {
			return nil
		}
		return &v.Capacity
	}).(pulumi.IntPtrOutput)
}

// The name of this SKU.
func (o SkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ReferenceDataSetKeyPropertyOutput{})
	pulumi.RegisterOutputType(ReferenceDataSetKeyPropertyArrayOutput{})
	pulumi.RegisterOutputType(ReferenceDataSetKeyPropertyResponseOutput{})
	pulumi.RegisterOutputType(ReferenceDataSetKeyPropertyResponseArrayOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuPtrOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SkuResponsePtrOutput{})
}
