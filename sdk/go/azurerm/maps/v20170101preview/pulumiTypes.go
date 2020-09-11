// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20170101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Additional Map account properties
type MapsAccountPropertiesResponse struct {
	// A unique identifier for the maps account
	XMsClientId *string `pulumi:"xMsClientId"`
}

// Additional Map account properties
type MapsAccountPropertiesResponseOutput struct{ *pulumi.OutputState }

func (MapsAccountPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MapsAccountPropertiesResponse)(nil)).Elem()
}

func (o MapsAccountPropertiesResponseOutput) ToMapsAccountPropertiesResponseOutput() MapsAccountPropertiesResponseOutput {
	return o
}

func (o MapsAccountPropertiesResponseOutput) ToMapsAccountPropertiesResponseOutputWithContext(ctx context.Context) MapsAccountPropertiesResponseOutput {
	return o
}

func (o MapsAccountPropertiesResponseOutput) ToMapsAccountPropertiesResponsePtrOutput() MapsAccountPropertiesResponsePtrOutput {
	return o.ToMapsAccountPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o MapsAccountPropertiesResponseOutput) ToMapsAccountPropertiesResponsePtrOutputWithContext(ctx context.Context) MapsAccountPropertiesResponsePtrOutput {
	return o.ApplyT(func(v MapsAccountPropertiesResponse) *MapsAccountPropertiesResponse {
		return &v
	}).(MapsAccountPropertiesResponsePtrOutput)
}

// A unique identifier for the maps account
func (o MapsAccountPropertiesResponseOutput) XMsClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MapsAccountPropertiesResponse) *string { return v.XMsClientId }).(pulumi.StringPtrOutput)
}

type MapsAccountPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (MapsAccountPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MapsAccountPropertiesResponse)(nil)).Elem()
}

func (o MapsAccountPropertiesResponsePtrOutput) ToMapsAccountPropertiesResponsePtrOutput() MapsAccountPropertiesResponsePtrOutput {
	return o
}

func (o MapsAccountPropertiesResponsePtrOutput) ToMapsAccountPropertiesResponsePtrOutputWithContext(ctx context.Context) MapsAccountPropertiesResponsePtrOutput {
	return o
}

func (o MapsAccountPropertiesResponsePtrOutput) Elem() MapsAccountPropertiesResponseOutput {
	return o.ApplyT(func(v *MapsAccountPropertiesResponse) MapsAccountPropertiesResponse { return *v }).(MapsAccountPropertiesResponseOutput)
}

// A unique identifier for the maps account
func (o MapsAccountPropertiesResponsePtrOutput) XMsClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MapsAccountPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.XMsClientId
	}).(pulumi.StringPtrOutput)
}

// The SKU of the Maps Account.
type Sku struct {
	// The name of the SKU, in standard format (such as S0).
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

// The SKU of the Maps Account.
type SkuArgs struct {
	// The name of the SKU, in standard format (such as S0).
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

// The SKU of the Maps Account.
type SkuResponse struct {
	// The name of the SKU, in standard format (such as S0).
	Name string `pulumi:"name"`
	// Gets the sku tier. This is based on the SKU name.
	Tier string `pulumi:"tier"`
}

// The SKU of the Maps Account.
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

// The name of the SKU, in standard format (such as S0).
func (o SkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

// Gets the sku tier. This is based on the SKU name.
func (o SkuResponseOutput) Tier() pulumi.StringOutput {
	return o.ApplyT(func(v SkuResponse) string { return v.Tier }).(pulumi.StringOutput)
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

// The name of the SKU, in standard format (such as S0).
func (o SkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

// Gets the sku tier. This is based on the SKU name.
func (o SkuResponsePtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Tier
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(MapsAccountPropertiesResponseOutput{})
	pulumi.RegisterOutputType(MapsAccountPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuPtrOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SkuResponsePtrOutput{})
}
