// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20181203

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The parameters to provide for the EnterpriseKnowledgeGraph.
type EnterpriseKnowledgeGraphProperties struct {
	// The description of the EnterpriseKnowledgeGraph
	Description *string `pulumi:"description"`
	// Specifies the metadata  of the resource.
	Metadata map[string]interface{} `pulumi:"metadata"`
	// The state of EnterpriseKnowledgeGraph provisioning
	ProvisioningState *string `pulumi:"provisioningState"`
}

// EnterpriseKnowledgeGraphPropertiesInput is an input type that accepts EnterpriseKnowledgeGraphPropertiesArgs and EnterpriseKnowledgeGraphPropertiesOutput values.
// You can construct a concrete instance of `EnterpriseKnowledgeGraphPropertiesInput` via:
//
//          EnterpriseKnowledgeGraphPropertiesArgs{...}
type EnterpriseKnowledgeGraphPropertiesInput interface {
	pulumi.Input

	ToEnterpriseKnowledgeGraphPropertiesOutput() EnterpriseKnowledgeGraphPropertiesOutput
	ToEnterpriseKnowledgeGraphPropertiesOutputWithContext(context.Context) EnterpriseKnowledgeGraphPropertiesOutput
}

// The parameters to provide for the EnterpriseKnowledgeGraph.
type EnterpriseKnowledgeGraphPropertiesArgs struct {
	// The description of the EnterpriseKnowledgeGraph
	Description pulumi.StringPtrInput `pulumi:"description"`
	// Specifies the metadata  of the resource.
	Metadata pulumi.MapInput `pulumi:"metadata"`
	// The state of EnterpriseKnowledgeGraph provisioning
	ProvisioningState pulumi.StringPtrInput `pulumi:"provisioningState"`
}

func (EnterpriseKnowledgeGraphPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EnterpriseKnowledgeGraphProperties)(nil)).Elem()
}

func (i EnterpriseKnowledgeGraphPropertiesArgs) ToEnterpriseKnowledgeGraphPropertiesOutput() EnterpriseKnowledgeGraphPropertiesOutput {
	return i.ToEnterpriseKnowledgeGraphPropertiesOutputWithContext(context.Background())
}

func (i EnterpriseKnowledgeGraphPropertiesArgs) ToEnterpriseKnowledgeGraphPropertiesOutputWithContext(ctx context.Context) EnterpriseKnowledgeGraphPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnterpriseKnowledgeGraphPropertiesOutput)
}

func (i EnterpriseKnowledgeGraphPropertiesArgs) ToEnterpriseKnowledgeGraphPropertiesPtrOutput() EnterpriseKnowledgeGraphPropertiesPtrOutput {
	return i.ToEnterpriseKnowledgeGraphPropertiesPtrOutputWithContext(context.Background())
}

func (i EnterpriseKnowledgeGraphPropertiesArgs) ToEnterpriseKnowledgeGraphPropertiesPtrOutputWithContext(ctx context.Context) EnterpriseKnowledgeGraphPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnterpriseKnowledgeGraphPropertiesOutput).ToEnterpriseKnowledgeGraphPropertiesPtrOutputWithContext(ctx)
}

// EnterpriseKnowledgeGraphPropertiesPtrInput is an input type that accepts EnterpriseKnowledgeGraphPropertiesArgs, EnterpriseKnowledgeGraphPropertiesPtr and EnterpriseKnowledgeGraphPropertiesPtrOutput values.
// You can construct a concrete instance of `EnterpriseKnowledgeGraphPropertiesPtrInput` via:
//
//          EnterpriseKnowledgeGraphPropertiesArgs{...}
//
//  or:
//
//          nil
type EnterpriseKnowledgeGraphPropertiesPtrInput interface {
	pulumi.Input

	ToEnterpriseKnowledgeGraphPropertiesPtrOutput() EnterpriseKnowledgeGraphPropertiesPtrOutput
	ToEnterpriseKnowledgeGraphPropertiesPtrOutputWithContext(context.Context) EnterpriseKnowledgeGraphPropertiesPtrOutput
}

type enterpriseKnowledgeGraphPropertiesPtrType EnterpriseKnowledgeGraphPropertiesArgs

func EnterpriseKnowledgeGraphPropertiesPtr(v *EnterpriseKnowledgeGraphPropertiesArgs) EnterpriseKnowledgeGraphPropertiesPtrInput {
	return (*enterpriseKnowledgeGraphPropertiesPtrType)(v)
}

func (*enterpriseKnowledgeGraphPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EnterpriseKnowledgeGraphProperties)(nil)).Elem()
}

func (i *enterpriseKnowledgeGraphPropertiesPtrType) ToEnterpriseKnowledgeGraphPropertiesPtrOutput() EnterpriseKnowledgeGraphPropertiesPtrOutput {
	return i.ToEnterpriseKnowledgeGraphPropertiesPtrOutputWithContext(context.Background())
}

func (i *enterpriseKnowledgeGraphPropertiesPtrType) ToEnterpriseKnowledgeGraphPropertiesPtrOutputWithContext(ctx context.Context) EnterpriseKnowledgeGraphPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnterpriseKnowledgeGraphPropertiesPtrOutput)
}

// The parameters to provide for the EnterpriseKnowledgeGraph.
type EnterpriseKnowledgeGraphPropertiesOutput struct{ *pulumi.OutputState }

func (EnterpriseKnowledgeGraphPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EnterpriseKnowledgeGraphProperties)(nil)).Elem()
}

func (o EnterpriseKnowledgeGraphPropertiesOutput) ToEnterpriseKnowledgeGraphPropertiesOutput() EnterpriseKnowledgeGraphPropertiesOutput {
	return o
}

func (o EnterpriseKnowledgeGraphPropertiesOutput) ToEnterpriseKnowledgeGraphPropertiesOutputWithContext(ctx context.Context) EnterpriseKnowledgeGraphPropertiesOutput {
	return o
}

func (o EnterpriseKnowledgeGraphPropertiesOutput) ToEnterpriseKnowledgeGraphPropertiesPtrOutput() EnterpriseKnowledgeGraphPropertiesPtrOutput {
	return o.ToEnterpriseKnowledgeGraphPropertiesPtrOutputWithContext(context.Background())
}

func (o EnterpriseKnowledgeGraphPropertiesOutput) ToEnterpriseKnowledgeGraphPropertiesPtrOutputWithContext(ctx context.Context) EnterpriseKnowledgeGraphPropertiesPtrOutput {
	return o.ApplyT(func(v EnterpriseKnowledgeGraphProperties) *EnterpriseKnowledgeGraphProperties {
		return &v
	}).(EnterpriseKnowledgeGraphPropertiesPtrOutput)
}

// The description of the EnterpriseKnowledgeGraph
func (o EnterpriseKnowledgeGraphPropertiesOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnterpriseKnowledgeGraphProperties) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Specifies the metadata  of the resource.
func (o EnterpriseKnowledgeGraphPropertiesOutput) Metadata() pulumi.MapOutput {
	return o.ApplyT(func(v EnterpriseKnowledgeGraphProperties) map[string]interface{} { return v.Metadata }).(pulumi.MapOutput)
}

// The state of EnterpriseKnowledgeGraph provisioning
func (o EnterpriseKnowledgeGraphPropertiesOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnterpriseKnowledgeGraphProperties) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

type EnterpriseKnowledgeGraphPropertiesPtrOutput struct{ *pulumi.OutputState }

func (EnterpriseKnowledgeGraphPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EnterpriseKnowledgeGraphProperties)(nil)).Elem()
}

func (o EnterpriseKnowledgeGraphPropertiesPtrOutput) ToEnterpriseKnowledgeGraphPropertiesPtrOutput() EnterpriseKnowledgeGraphPropertiesPtrOutput {
	return o
}

func (o EnterpriseKnowledgeGraphPropertiesPtrOutput) ToEnterpriseKnowledgeGraphPropertiesPtrOutputWithContext(ctx context.Context) EnterpriseKnowledgeGraphPropertiesPtrOutput {
	return o
}

func (o EnterpriseKnowledgeGraphPropertiesPtrOutput) Elem() EnterpriseKnowledgeGraphPropertiesOutput {
	return o.ApplyT(func(v *EnterpriseKnowledgeGraphProperties) EnterpriseKnowledgeGraphProperties { return *v }).(EnterpriseKnowledgeGraphPropertiesOutput)
}

// The description of the EnterpriseKnowledgeGraph
func (o EnterpriseKnowledgeGraphPropertiesPtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnterpriseKnowledgeGraphProperties) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

// Specifies the metadata  of the resource.
func (o EnterpriseKnowledgeGraphPropertiesPtrOutput) Metadata() pulumi.MapOutput {
	return o.ApplyT(func(v *EnterpriseKnowledgeGraphProperties) map[string]interface{} {
		if v == nil {
			return nil
		}
		return v.Metadata
	}).(pulumi.MapOutput)
}

// The state of EnterpriseKnowledgeGraph provisioning
func (o EnterpriseKnowledgeGraphPropertiesPtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnterpriseKnowledgeGraphProperties) *string {
		if v == nil {
			return nil
		}
		return v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

// The parameters to provide for the EnterpriseKnowledgeGraph.
type EnterpriseKnowledgeGraphPropertiesResponse struct {
	// The description of the EnterpriseKnowledgeGraph
	Description *string `pulumi:"description"`
	// Specifies the metadata  of the resource.
	Metadata map[string]interface{} `pulumi:"metadata"`
	// The state of EnterpriseKnowledgeGraph provisioning
	ProvisioningState *string `pulumi:"provisioningState"`
}

// EnterpriseKnowledgeGraphPropertiesResponseInput is an input type that accepts EnterpriseKnowledgeGraphPropertiesResponseArgs and EnterpriseKnowledgeGraphPropertiesResponseOutput values.
// You can construct a concrete instance of `EnterpriseKnowledgeGraphPropertiesResponseInput` via:
//
//          EnterpriseKnowledgeGraphPropertiesResponseArgs{...}
type EnterpriseKnowledgeGraphPropertiesResponseInput interface {
	pulumi.Input

	ToEnterpriseKnowledgeGraphPropertiesResponseOutput() EnterpriseKnowledgeGraphPropertiesResponseOutput
	ToEnterpriseKnowledgeGraphPropertiesResponseOutputWithContext(context.Context) EnterpriseKnowledgeGraphPropertiesResponseOutput
}

// The parameters to provide for the EnterpriseKnowledgeGraph.
type EnterpriseKnowledgeGraphPropertiesResponseArgs struct {
	// The description of the EnterpriseKnowledgeGraph
	Description pulumi.StringPtrInput `pulumi:"description"`
	// Specifies the metadata  of the resource.
	Metadata pulumi.MapInput `pulumi:"metadata"`
	// The state of EnterpriseKnowledgeGraph provisioning
	ProvisioningState pulumi.StringPtrInput `pulumi:"provisioningState"`
}

func (EnterpriseKnowledgeGraphPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EnterpriseKnowledgeGraphPropertiesResponse)(nil)).Elem()
}

func (i EnterpriseKnowledgeGraphPropertiesResponseArgs) ToEnterpriseKnowledgeGraphPropertiesResponseOutput() EnterpriseKnowledgeGraphPropertiesResponseOutput {
	return i.ToEnterpriseKnowledgeGraphPropertiesResponseOutputWithContext(context.Background())
}

func (i EnterpriseKnowledgeGraphPropertiesResponseArgs) ToEnterpriseKnowledgeGraphPropertiesResponseOutputWithContext(ctx context.Context) EnterpriseKnowledgeGraphPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnterpriseKnowledgeGraphPropertiesResponseOutput)
}

func (i EnterpriseKnowledgeGraphPropertiesResponseArgs) ToEnterpriseKnowledgeGraphPropertiesResponsePtrOutput() EnterpriseKnowledgeGraphPropertiesResponsePtrOutput {
	return i.ToEnterpriseKnowledgeGraphPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i EnterpriseKnowledgeGraphPropertiesResponseArgs) ToEnterpriseKnowledgeGraphPropertiesResponsePtrOutputWithContext(ctx context.Context) EnterpriseKnowledgeGraphPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnterpriseKnowledgeGraphPropertiesResponseOutput).ToEnterpriseKnowledgeGraphPropertiesResponsePtrOutputWithContext(ctx)
}

// EnterpriseKnowledgeGraphPropertiesResponsePtrInput is an input type that accepts EnterpriseKnowledgeGraphPropertiesResponseArgs, EnterpriseKnowledgeGraphPropertiesResponsePtr and EnterpriseKnowledgeGraphPropertiesResponsePtrOutput values.
// You can construct a concrete instance of `EnterpriseKnowledgeGraphPropertiesResponsePtrInput` via:
//
//          EnterpriseKnowledgeGraphPropertiesResponseArgs{...}
//
//  or:
//
//          nil
type EnterpriseKnowledgeGraphPropertiesResponsePtrInput interface {
	pulumi.Input

	ToEnterpriseKnowledgeGraphPropertiesResponsePtrOutput() EnterpriseKnowledgeGraphPropertiesResponsePtrOutput
	ToEnterpriseKnowledgeGraphPropertiesResponsePtrOutputWithContext(context.Context) EnterpriseKnowledgeGraphPropertiesResponsePtrOutput
}

type enterpriseKnowledgeGraphPropertiesResponsePtrType EnterpriseKnowledgeGraphPropertiesResponseArgs

func EnterpriseKnowledgeGraphPropertiesResponsePtr(v *EnterpriseKnowledgeGraphPropertiesResponseArgs) EnterpriseKnowledgeGraphPropertiesResponsePtrInput {
	return (*enterpriseKnowledgeGraphPropertiesResponsePtrType)(v)
}

func (*enterpriseKnowledgeGraphPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EnterpriseKnowledgeGraphPropertiesResponse)(nil)).Elem()
}

func (i *enterpriseKnowledgeGraphPropertiesResponsePtrType) ToEnterpriseKnowledgeGraphPropertiesResponsePtrOutput() EnterpriseKnowledgeGraphPropertiesResponsePtrOutput {
	return i.ToEnterpriseKnowledgeGraphPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *enterpriseKnowledgeGraphPropertiesResponsePtrType) ToEnterpriseKnowledgeGraphPropertiesResponsePtrOutputWithContext(ctx context.Context) EnterpriseKnowledgeGraphPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnterpriseKnowledgeGraphPropertiesResponsePtrOutput)
}

// The parameters to provide for the EnterpriseKnowledgeGraph.
type EnterpriseKnowledgeGraphPropertiesResponseOutput struct{ *pulumi.OutputState }

func (EnterpriseKnowledgeGraphPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EnterpriseKnowledgeGraphPropertiesResponse)(nil)).Elem()
}

func (o EnterpriseKnowledgeGraphPropertiesResponseOutput) ToEnterpriseKnowledgeGraphPropertiesResponseOutput() EnterpriseKnowledgeGraphPropertiesResponseOutput {
	return o
}

func (o EnterpriseKnowledgeGraphPropertiesResponseOutput) ToEnterpriseKnowledgeGraphPropertiesResponseOutputWithContext(ctx context.Context) EnterpriseKnowledgeGraphPropertiesResponseOutput {
	return o
}

func (o EnterpriseKnowledgeGraphPropertiesResponseOutput) ToEnterpriseKnowledgeGraphPropertiesResponsePtrOutput() EnterpriseKnowledgeGraphPropertiesResponsePtrOutput {
	return o.ToEnterpriseKnowledgeGraphPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o EnterpriseKnowledgeGraphPropertiesResponseOutput) ToEnterpriseKnowledgeGraphPropertiesResponsePtrOutputWithContext(ctx context.Context) EnterpriseKnowledgeGraphPropertiesResponsePtrOutput {
	return o.ApplyT(func(v EnterpriseKnowledgeGraphPropertiesResponse) *EnterpriseKnowledgeGraphPropertiesResponse {
		return &v
	}).(EnterpriseKnowledgeGraphPropertiesResponsePtrOutput)
}

// The description of the EnterpriseKnowledgeGraph
func (o EnterpriseKnowledgeGraphPropertiesResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnterpriseKnowledgeGraphPropertiesResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Specifies the metadata  of the resource.
func (o EnterpriseKnowledgeGraphPropertiesResponseOutput) Metadata() pulumi.MapOutput {
	return o.ApplyT(func(v EnterpriseKnowledgeGraphPropertiesResponse) map[string]interface{} { return v.Metadata }).(pulumi.MapOutput)
}

// The state of EnterpriseKnowledgeGraph provisioning
func (o EnterpriseKnowledgeGraphPropertiesResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnterpriseKnowledgeGraphPropertiesResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

type EnterpriseKnowledgeGraphPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (EnterpriseKnowledgeGraphPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EnterpriseKnowledgeGraphPropertiesResponse)(nil)).Elem()
}

func (o EnterpriseKnowledgeGraphPropertiesResponsePtrOutput) ToEnterpriseKnowledgeGraphPropertiesResponsePtrOutput() EnterpriseKnowledgeGraphPropertiesResponsePtrOutput {
	return o
}

func (o EnterpriseKnowledgeGraphPropertiesResponsePtrOutput) ToEnterpriseKnowledgeGraphPropertiesResponsePtrOutputWithContext(ctx context.Context) EnterpriseKnowledgeGraphPropertiesResponsePtrOutput {
	return o
}

func (o EnterpriseKnowledgeGraphPropertiesResponsePtrOutput) Elem() EnterpriseKnowledgeGraphPropertiesResponseOutput {
	return o.ApplyT(func(v *EnterpriseKnowledgeGraphPropertiesResponse) EnterpriseKnowledgeGraphPropertiesResponse {
		return *v
	}).(EnterpriseKnowledgeGraphPropertiesResponseOutput)
}

// The description of the EnterpriseKnowledgeGraph
func (o EnterpriseKnowledgeGraphPropertiesResponsePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnterpriseKnowledgeGraphPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

// Specifies the metadata  of the resource.
func (o EnterpriseKnowledgeGraphPropertiesResponsePtrOutput) Metadata() pulumi.MapOutput {
	return o.ApplyT(func(v *EnterpriseKnowledgeGraphPropertiesResponse) map[string]interface{} {
		if v == nil {
			return nil
		}
		return v.Metadata
	}).(pulumi.MapOutput)
}

// The state of EnterpriseKnowledgeGraph provisioning
func (o EnterpriseKnowledgeGraphPropertiesResponsePtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnterpriseKnowledgeGraphPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

// The SKU of the EnterpriseKnowledgeGraph service account.
type Sku struct {
	// The sku name
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

// The SKU of the EnterpriseKnowledgeGraph service account.
type SkuArgs struct {
	// The sku name
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

// The SKU of the EnterpriseKnowledgeGraph service account.
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

// The sku name
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

// The sku name
func (o SkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

// The SKU of the EnterpriseKnowledgeGraph service account.
type SkuResponse struct {
	// The sku name
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

// The SKU of the EnterpriseKnowledgeGraph service account.
type SkuResponseArgs struct {
	// The sku name
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

// The SKU of the EnterpriseKnowledgeGraph service account.
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

// The sku name
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

// The sku name
func (o SkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(EnterpriseKnowledgeGraphPropertiesOutput{})
	pulumi.RegisterOutputType(EnterpriseKnowledgeGraphPropertiesPtrOutput{})
	pulumi.RegisterOutputType(EnterpriseKnowledgeGraphPropertiesResponseOutput{})
	pulumi.RegisterOutputType(EnterpriseKnowledgeGraphPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuPtrOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SkuResponsePtrOutput{})
}
