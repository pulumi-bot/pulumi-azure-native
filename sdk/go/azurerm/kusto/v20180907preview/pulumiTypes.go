// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180907preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type AzureSku struct {
	// SKU capacity.
	Capacity *int `pulumi:"capacity"`
	// SKU name.
	Name string `pulumi:"name"`
	// SKU tier.
	Tier string `pulumi:"tier"`
}

// AzureSkuInput is an input type that accepts AzureSkuArgs and AzureSkuOutput values.
// You can construct a concrete instance of `AzureSkuInput` via:
//
//          AzureSkuArgs{...}
type AzureSkuInput interface {
	pulumi.Input

	ToAzureSkuOutput() AzureSkuOutput
	ToAzureSkuOutputWithContext(context.Context) AzureSkuOutput
}

type AzureSkuArgs struct {
	// SKU capacity.
	Capacity pulumi.IntPtrInput `pulumi:"capacity"`
	// SKU name.
	Name pulumi.StringInput `pulumi:"name"`
	// SKU tier.
	Tier pulumi.StringInput `pulumi:"tier"`
}

func (AzureSkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureSku)(nil)).Elem()
}

func (i AzureSkuArgs) ToAzureSkuOutput() AzureSkuOutput {
	return i.ToAzureSkuOutputWithContext(context.Background())
}

func (i AzureSkuArgs) ToAzureSkuOutputWithContext(ctx context.Context) AzureSkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureSkuOutput)
}

func (i AzureSkuArgs) ToAzureSkuPtrOutput() AzureSkuPtrOutput {
	return i.ToAzureSkuPtrOutputWithContext(context.Background())
}

func (i AzureSkuArgs) ToAzureSkuPtrOutputWithContext(ctx context.Context) AzureSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureSkuOutput).ToAzureSkuPtrOutputWithContext(ctx)
}

// AzureSkuPtrInput is an input type that accepts AzureSkuArgs, AzureSkuPtr and AzureSkuPtrOutput values.
// You can construct a concrete instance of `AzureSkuPtrInput` via:
//
//          AzureSkuArgs{...}
//
//  or:
//
//          nil
type AzureSkuPtrInput interface {
	pulumi.Input

	ToAzureSkuPtrOutput() AzureSkuPtrOutput
	ToAzureSkuPtrOutputWithContext(context.Context) AzureSkuPtrOutput
}

type azureSkuPtrType AzureSkuArgs

func AzureSkuPtr(v *AzureSkuArgs) AzureSkuPtrInput {
	return (*azureSkuPtrType)(v)
}

func (*azureSkuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureSku)(nil)).Elem()
}

func (i *azureSkuPtrType) ToAzureSkuPtrOutput() AzureSkuPtrOutput {
	return i.ToAzureSkuPtrOutputWithContext(context.Background())
}

func (i *azureSkuPtrType) ToAzureSkuPtrOutputWithContext(ctx context.Context) AzureSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureSkuPtrOutput)
}

type AzureSkuOutput struct{ *pulumi.OutputState }

func (AzureSkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureSku)(nil)).Elem()
}

func (o AzureSkuOutput) ToAzureSkuOutput() AzureSkuOutput {
	return o
}

func (o AzureSkuOutput) ToAzureSkuOutputWithContext(ctx context.Context) AzureSkuOutput {
	return o
}

func (o AzureSkuOutput) ToAzureSkuPtrOutput() AzureSkuPtrOutput {
	return o.ToAzureSkuPtrOutputWithContext(context.Background())
}

func (o AzureSkuOutput) ToAzureSkuPtrOutputWithContext(ctx context.Context) AzureSkuPtrOutput {
	return o.ApplyT(func(v AzureSku) *AzureSku {
		return &v
	}).(AzureSkuPtrOutput)
}

// SKU capacity.
func (o AzureSkuOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AzureSku) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

// SKU name.
func (o AzureSkuOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v AzureSku) string { return v.Name }).(pulumi.StringOutput)
}

// SKU tier.
func (o AzureSkuOutput) Tier() pulumi.StringOutput {
	return o.ApplyT(func(v AzureSku) string { return v.Tier }).(pulumi.StringOutput)
}

type AzureSkuPtrOutput struct{ *pulumi.OutputState }

func (AzureSkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureSku)(nil)).Elem()
}

func (o AzureSkuPtrOutput) ToAzureSkuPtrOutput() AzureSkuPtrOutput {
	return o
}

func (o AzureSkuPtrOutput) ToAzureSkuPtrOutputWithContext(ctx context.Context) AzureSkuPtrOutput {
	return o
}

func (o AzureSkuPtrOutput) Elem() AzureSkuOutput {
	return o.ApplyT(func(v *AzureSku) AzureSku { return *v }).(AzureSkuOutput)
}

// SKU capacity.
func (o AzureSkuPtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AzureSku) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

// SKU name.
func (o AzureSkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureSku) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

// SKU tier.
func (o AzureSkuPtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureSku) *string {
		if v == nil {
			return nil
		}
		return &v.Tier
	}).(pulumi.StringPtrOutput)
}

type AzureSkuResponse struct {
	// SKU capacity.
	Capacity *int `pulumi:"capacity"`
	// SKU name.
	Name string `pulumi:"name"`
	// SKU tier.
	Tier string `pulumi:"tier"`
}

// AzureSkuResponseInput is an input type that accepts AzureSkuResponseArgs and AzureSkuResponseOutput values.
// You can construct a concrete instance of `AzureSkuResponseInput` via:
//
//          AzureSkuResponseArgs{...}
type AzureSkuResponseInput interface {
	pulumi.Input

	ToAzureSkuResponseOutput() AzureSkuResponseOutput
	ToAzureSkuResponseOutputWithContext(context.Context) AzureSkuResponseOutput
}

type AzureSkuResponseArgs struct {
	// SKU capacity.
	Capacity pulumi.IntPtrInput `pulumi:"capacity"`
	// SKU name.
	Name pulumi.StringInput `pulumi:"name"`
	// SKU tier.
	Tier pulumi.StringInput `pulumi:"tier"`
}

func (AzureSkuResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureSkuResponse)(nil)).Elem()
}

func (i AzureSkuResponseArgs) ToAzureSkuResponseOutput() AzureSkuResponseOutput {
	return i.ToAzureSkuResponseOutputWithContext(context.Background())
}

func (i AzureSkuResponseArgs) ToAzureSkuResponseOutputWithContext(ctx context.Context) AzureSkuResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureSkuResponseOutput)
}

func (i AzureSkuResponseArgs) ToAzureSkuResponsePtrOutput() AzureSkuResponsePtrOutput {
	return i.ToAzureSkuResponsePtrOutputWithContext(context.Background())
}

func (i AzureSkuResponseArgs) ToAzureSkuResponsePtrOutputWithContext(ctx context.Context) AzureSkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureSkuResponseOutput).ToAzureSkuResponsePtrOutputWithContext(ctx)
}

// AzureSkuResponsePtrInput is an input type that accepts AzureSkuResponseArgs, AzureSkuResponsePtr and AzureSkuResponsePtrOutput values.
// You can construct a concrete instance of `AzureSkuResponsePtrInput` via:
//
//          AzureSkuResponseArgs{...}
//
//  or:
//
//          nil
type AzureSkuResponsePtrInput interface {
	pulumi.Input

	ToAzureSkuResponsePtrOutput() AzureSkuResponsePtrOutput
	ToAzureSkuResponsePtrOutputWithContext(context.Context) AzureSkuResponsePtrOutput
}

type azureSkuResponsePtrType AzureSkuResponseArgs

func AzureSkuResponsePtr(v *AzureSkuResponseArgs) AzureSkuResponsePtrInput {
	return (*azureSkuResponsePtrType)(v)
}

func (*azureSkuResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureSkuResponse)(nil)).Elem()
}

func (i *azureSkuResponsePtrType) ToAzureSkuResponsePtrOutput() AzureSkuResponsePtrOutput {
	return i.ToAzureSkuResponsePtrOutputWithContext(context.Background())
}

func (i *azureSkuResponsePtrType) ToAzureSkuResponsePtrOutputWithContext(ctx context.Context) AzureSkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureSkuResponsePtrOutput)
}

type AzureSkuResponseOutput struct{ *pulumi.OutputState }

func (AzureSkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureSkuResponse)(nil)).Elem()
}

func (o AzureSkuResponseOutput) ToAzureSkuResponseOutput() AzureSkuResponseOutput {
	return o
}

func (o AzureSkuResponseOutput) ToAzureSkuResponseOutputWithContext(ctx context.Context) AzureSkuResponseOutput {
	return o
}

func (o AzureSkuResponseOutput) ToAzureSkuResponsePtrOutput() AzureSkuResponsePtrOutput {
	return o.ToAzureSkuResponsePtrOutputWithContext(context.Background())
}

func (o AzureSkuResponseOutput) ToAzureSkuResponsePtrOutputWithContext(ctx context.Context) AzureSkuResponsePtrOutput {
	return o.ApplyT(func(v AzureSkuResponse) *AzureSkuResponse {
		return &v
	}).(AzureSkuResponsePtrOutput)
}

// SKU capacity.
func (o AzureSkuResponseOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AzureSkuResponse) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

// SKU name.
func (o AzureSkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v AzureSkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

// SKU tier.
func (o AzureSkuResponseOutput) Tier() pulumi.StringOutput {
	return o.ApplyT(func(v AzureSkuResponse) string { return v.Tier }).(pulumi.StringOutput)
}

type AzureSkuResponsePtrOutput struct{ *pulumi.OutputState }

func (AzureSkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureSkuResponse)(nil)).Elem()
}

func (o AzureSkuResponsePtrOutput) ToAzureSkuResponsePtrOutput() AzureSkuResponsePtrOutput {
	return o
}

func (o AzureSkuResponsePtrOutput) ToAzureSkuResponsePtrOutputWithContext(ctx context.Context) AzureSkuResponsePtrOutput {
	return o
}

func (o AzureSkuResponsePtrOutput) Elem() AzureSkuResponseOutput {
	return o.ApplyT(func(v *AzureSkuResponse) AzureSkuResponse { return *v }).(AzureSkuResponseOutput)
}

// SKU capacity.
func (o AzureSkuResponsePtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AzureSkuResponse) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

// SKU name.
func (o AzureSkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureSkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

// SKU tier.
func (o AzureSkuResponsePtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureSkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Tier
	}).(pulumi.StringPtrOutput)
}

type DatabasePrincipalResponse struct {
	// Application id - relevant only for application principal type.
	AppId *string `pulumi:"appId"`
	// Database principal email if exists.
	Email *string `pulumi:"email"`
	// Database principal fully qualified name.
	Fqn *string `pulumi:"fqn"`
	// Database principal name.
	Name string `pulumi:"name"`
	// Database principal role.
	Role string `pulumi:"role"`
	// Database principal type.
	Type string `pulumi:"type"`
}

// DatabasePrincipalResponseInput is an input type that accepts DatabasePrincipalResponseArgs and DatabasePrincipalResponseOutput values.
// You can construct a concrete instance of `DatabasePrincipalResponseInput` via:
//
//          DatabasePrincipalResponseArgs{...}
type DatabasePrincipalResponseInput interface {
	pulumi.Input

	ToDatabasePrincipalResponseOutput() DatabasePrincipalResponseOutput
	ToDatabasePrincipalResponseOutputWithContext(context.Context) DatabasePrincipalResponseOutput
}

type DatabasePrincipalResponseArgs struct {
	// Application id - relevant only for application principal type.
	AppId pulumi.StringPtrInput `pulumi:"appId"`
	// Database principal email if exists.
	Email pulumi.StringPtrInput `pulumi:"email"`
	// Database principal fully qualified name.
	Fqn pulumi.StringPtrInput `pulumi:"fqn"`
	// Database principal name.
	Name pulumi.StringInput `pulumi:"name"`
	// Database principal role.
	Role pulumi.StringInput `pulumi:"role"`
	// Database principal type.
	Type pulumi.StringInput `pulumi:"type"`
}

func (DatabasePrincipalResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabasePrincipalResponse)(nil)).Elem()
}

func (i DatabasePrincipalResponseArgs) ToDatabasePrincipalResponseOutput() DatabasePrincipalResponseOutput {
	return i.ToDatabasePrincipalResponseOutputWithContext(context.Background())
}

func (i DatabasePrincipalResponseArgs) ToDatabasePrincipalResponseOutputWithContext(ctx context.Context) DatabasePrincipalResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabasePrincipalResponseOutput)
}

// DatabasePrincipalResponseArrayInput is an input type that accepts DatabasePrincipalResponseArray and DatabasePrincipalResponseArrayOutput values.
// You can construct a concrete instance of `DatabasePrincipalResponseArrayInput` via:
//
//          DatabasePrincipalResponseArray{ DatabasePrincipalResponseArgs{...} }
type DatabasePrincipalResponseArrayInput interface {
	pulumi.Input

	ToDatabasePrincipalResponseArrayOutput() DatabasePrincipalResponseArrayOutput
	ToDatabasePrincipalResponseArrayOutputWithContext(context.Context) DatabasePrincipalResponseArrayOutput
}

type DatabasePrincipalResponseArray []DatabasePrincipalResponseInput

func (DatabasePrincipalResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DatabasePrincipalResponse)(nil)).Elem()
}

func (i DatabasePrincipalResponseArray) ToDatabasePrincipalResponseArrayOutput() DatabasePrincipalResponseArrayOutput {
	return i.ToDatabasePrincipalResponseArrayOutputWithContext(context.Background())
}

func (i DatabasePrincipalResponseArray) ToDatabasePrincipalResponseArrayOutputWithContext(ctx context.Context) DatabasePrincipalResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabasePrincipalResponseArrayOutput)
}

type DatabasePrincipalResponseOutput struct{ *pulumi.OutputState }

func (DatabasePrincipalResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabasePrincipalResponse)(nil)).Elem()
}

func (o DatabasePrincipalResponseOutput) ToDatabasePrincipalResponseOutput() DatabasePrincipalResponseOutput {
	return o
}

func (o DatabasePrincipalResponseOutput) ToDatabasePrincipalResponseOutputWithContext(ctx context.Context) DatabasePrincipalResponseOutput {
	return o
}

// Application id - relevant only for application principal type.
func (o DatabasePrincipalResponseOutput) AppId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatabasePrincipalResponse) *string { return v.AppId }).(pulumi.StringPtrOutput)
}

// Database principal email if exists.
func (o DatabasePrincipalResponseOutput) Email() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatabasePrincipalResponse) *string { return v.Email }).(pulumi.StringPtrOutput)
}

// Database principal fully qualified name.
func (o DatabasePrincipalResponseOutput) Fqn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatabasePrincipalResponse) *string { return v.Fqn }).(pulumi.StringPtrOutput)
}

// Database principal name.
func (o DatabasePrincipalResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v DatabasePrincipalResponse) string { return v.Name }).(pulumi.StringOutput)
}

// Database principal role.
func (o DatabasePrincipalResponseOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v DatabasePrincipalResponse) string { return v.Role }).(pulumi.StringOutput)
}

// Database principal type.
func (o DatabasePrincipalResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v DatabasePrincipalResponse) string { return v.Type }).(pulumi.StringOutput)
}

type DatabasePrincipalResponseArrayOutput struct{ *pulumi.OutputState }

func (DatabasePrincipalResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DatabasePrincipalResponse)(nil)).Elem()
}

func (o DatabasePrincipalResponseArrayOutput) ToDatabasePrincipalResponseArrayOutput() DatabasePrincipalResponseArrayOutput {
	return o
}

func (o DatabasePrincipalResponseArrayOutput) ToDatabasePrincipalResponseArrayOutputWithContext(ctx context.Context) DatabasePrincipalResponseArrayOutput {
	return o
}

func (o DatabasePrincipalResponseArrayOutput) Index(i pulumi.IntInput) DatabasePrincipalResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DatabasePrincipalResponse {
		return vs[0].([]DatabasePrincipalResponse)[vs[1].(int)]
	}).(DatabasePrincipalResponseOutput)
}

type DatabaseStatistics struct {
	// The database size - the total size of compressed data and index in bytes.
	Size *float64 `pulumi:"size"`
}

// DatabaseStatisticsInput is an input type that accepts DatabaseStatisticsArgs and DatabaseStatisticsOutput values.
// You can construct a concrete instance of `DatabaseStatisticsInput` via:
//
//          DatabaseStatisticsArgs{...}
type DatabaseStatisticsInput interface {
	pulumi.Input

	ToDatabaseStatisticsOutput() DatabaseStatisticsOutput
	ToDatabaseStatisticsOutputWithContext(context.Context) DatabaseStatisticsOutput
}

type DatabaseStatisticsArgs struct {
	// The database size - the total size of compressed data and index in bytes.
	Size pulumi.Float64PtrInput `pulumi:"size"`
}

func (DatabaseStatisticsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseStatistics)(nil)).Elem()
}

func (i DatabaseStatisticsArgs) ToDatabaseStatisticsOutput() DatabaseStatisticsOutput {
	return i.ToDatabaseStatisticsOutputWithContext(context.Background())
}

func (i DatabaseStatisticsArgs) ToDatabaseStatisticsOutputWithContext(ctx context.Context) DatabaseStatisticsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseStatisticsOutput)
}

func (i DatabaseStatisticsArgs) ToDatabaseStatisticsPtrOutput() DatabaseStatisticsPtrOutput {
	return i.ToDatabaseStatisticsPtrOutputWithContext(context.Background())
}

func (i DatabaseStatisticsArgs) ToDatabaseStatisticsPtrOutputWithContext(ctx context.Context) DatabaseStatisticsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseStatisticsOutput).ToDatabaseStatisticsPtrOutputWithContext(ctx)
}

// DatabaseStatisticsPtrInput is an input type that accepts DatabaseStatisticsArgs, DatabaseStatisticsPtr and DatabaseStatisticsPtrOutput values.
// You can construct a concrete instance of `DatabaseStatisticsPtrInput` via:
//
//          DatabaseStatisticsArgs{...}
//
//  or:
//
//          nil
type DatabaseStatisticsPtrInput interface {
	pulumi.Input

	ToDatabaseStatisticsPtrOutput() DatabaseStatisticsPtrOutput
	ToDatabaseStatisticsPtrOutputWithContext(context.Context) DatabaseStatisticsPtrOutput
}

type databaseStatisticsPtrType DatabaseStatisticsArgs

func DatabaseStatisticsPtr(v *DatabaseStatisticsArgs) DatabaseStatisticsPtrInput {
	return (*databaseStatisticsPtrType)(v)
}

func (*databaseStatisticsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseStatistics)(nil)).Elem()
}

func (i *databaseStatisticsPtrType) ToDatabaseStatisticsPtrOutput() DatabaseStatisticsPtrOutput {
	return i.ToDatabaseStatisticsPtrOutputWithContext(context.Background())
}

func (i *databaseStatisticsPtrType) ToDatabaseStatisticsPtrOutputWithContext(ctx context.Context) DatabaseStatisticsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseStatisticsPtrOutput)
}

type DatabaseStatisticsOutput struct{ *pulumi.OutputState }

func (DatabaseStatisticsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseStatistics)(nil)).Elem()
}

func (o DatabaseStatisticsOutput) ToDatabaseStatisticsOutput() DatabaseStatisticsOutput {
	return o
}

func (o DatabaseStatisticsOutput) ToDatabaseStatisticsOutputWithContext(ctx context.Context) DatabaseStatisticsOutput {
	return o
}

func (o DatabaseStatisticsOutput) ToDatabaseStatisticsPtrOutput() DatabaseStatisticsPtrOutput {
	return o.ToDatabaseStatisticsPtrOutputWithContext(context.Background())
}

func (o DatabaseStatisticsOutput) ToDatabaseStatisticsPtrOutputWithContext(ctx context.Context) DatabaseStatisticsPtrOutput {
	return o.ApplyT(func(v DatabaseStatistics) *DatabaseStatistics {
		return &v
	}).(DatabaseStatisticsPtrOutput)
}

// The database size - the total size of compressed data and index in bytes.
func (o DatabaseStatisticsOutput) Size() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v DatabaseStatistics) *float64 { return v.Size }).(pulumi.Float64PtrOutput)
}

type DatabaseStatisticsPtrOutput struct{ *pulumi.OutputState }

func (DatabaseStatisticsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseStatistics)(nil)).Elem()
}

func (o DatabaseStatisticsPtrOutput) ToDatabaseStatisticsPtrOutput() DatabaseStatisticsPtrOutput {
	return o
}

func (o DatabaseStatisticsPtrOutput) ToDatabaseStatisticsPtrOutputWithContext(ctx context.Context) DatabaseStatisticsPtrOutput {
	return o
}

func (o DatabaseStatisticsPtrOutput) Elem() DatabaseStatisticsOutput {
	return o.ApplyT(func(v *DatabaseStatistics) DatabaseStatistics { return *v }).(DatabaseStatisticsOutput)
}

// The database size - the total size of compressed data and index in bytes.
func (o DatabaseStatisticsPtrOutput) Size() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *DatabaseStatistics) *float64 {
		if v == nil {
			return nil
		}
		return v.Size
	}).(pulumi.Float64PtrOutput)
}

type DatabaseStatisticsResponse struct {
	// The database size - the total size of compressed data and index in bytes.
	Size *float64 `pulumi:"size"`
}

// DatabaseStatisticsResponseInput is an input type that accepts DatabaseStatisticsResponseArgs and DatabaseStatisticsResponseOutput values.
// You can construct a concrete instance of `DatabaseStatisticsResponseInput` via:
//
//          DatabaseStatisticsResponseArgs{...}
type DatabaseStatisticsResponseInput interface {
	pulumi.Input

	ToDatabaseStatisticsResponseOutput() DatabaseStatisticsResponseOutput
	ToDatabaseStatisticsResponseOutputWithContext(context.Context) DatabaseStatisticsResponseOutput
}

type DatabaseStatisticsResponseArgs struct {
	// The database size - the total size of compressed data and index in bytes.
	Size pulumi.Float64PtrInput `pulumi:"size"`
}

func (DatabaseStatisticsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseStatisticsResponse)(nil)).Elem()
}

func (i DatabaseStatisticsResponseArgs) ToDatabaseStatisticsResponseOutput() DatabaseStatisticsResponseOutput {
	return i.ToDatabaseStatisticsResponseOutputWithContext(context.Background())
}

func (i DatabaseStatisticsResponseArgs) ToDatabaseStatisticsResponseOutputWithContext(ctx context.Context) DatabaseStatisticsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseStatisticsResponseOutput)
}

func (i DatabaseStatisticsResponseArgs) ToDatabaseStatisticsResponsePtrOutput() DatabaseStatisticsResponsePtrOutput {
	return i.ToDatabaseStatisticsResponsePtrOutputWithContext(context.Background())
}

func (i DatabaseStatisticsResponseArgs) ToDatabaseStatisticsResponsePtrOutputWithContext(ctx context.Context) DatabaseStatisticsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseStatisticsResponseOutput).ToDatabaseStatisticsResponsePtrOutputWithContext(ctx)
}

// DatabaseStatisticsResponsePtrInput is an input type that accepts DatabaseStatisticsResponseArgs, DatabaseStatisticsResponsePtr and DatabaseStatisticsResponsePtrOutput values.
// You can construct a concrete instance of `DatabaseStatisticsResponsePtrInput` via:
//
//          DatabaseStatisticsResponseArgs{...}
//
//  or:
//
//          nil
type DatabaseStatisticsResponsePtrInput interface {
	pulumi.Input

	ToDatabaseStatisticsResponsePtrOutput() DatabaseStatisticsResponsePtrOutput
	ToDatabaseStatisticsResponsePtrOutputWithContext(context.Context) DatabaseStatisticsResponsePtrOutput
}

type databaseStatisticsResponsePtrType DatabaseStatisticsResponseArgs

func DatabaseStatisticsResponsePtr(v *DatabaseStatisticsResponseArgs) DatabaseStatisticsResponsePtrInput {
	return (*databaseStatisticsResponsePtrType)(v)
}

func (*databaseStatisticsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseStatisticsResponse)(nil)).Elem()
}

func (i *databaseStatisticsResponsePtrType) ToDatabaseStatisticsResponsePtrOutput() DatabaseStatisticsResponsePtrOutput {
	return i.ToDatabaseStatisticsResponsePtrOutputWithContext(context.Background())
}

func (i *databaseStatisticsResponsePtrType) ToDatabaseStatisticsResponsePtrOutputWithContext(ctx context.Context) DatabaseStatisticsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseStatisticsResponsePtrOutput)
}

type DatabaseStatisticsResponseOutput struct{ *pulumi.OutputState }

func (DatabaseStatisticsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseStatisticsResponse)(nil)).Elem()
}

func (o DatabaseStatisticsResponseOutput) ToDatabaseStatisticsResponseOutput() DatabaseStatisticsResponseOutput {
	return o
}

func (o DatabaseStatisticsResponseOutput) ToDatabaseStatisticsResponseOutputWithContext(ctx context.Context) DatabaseStatisticsResponseOutput {
	return o
}

func (o DatabaseStatisticsResponseOutput) ToDatabaseStatisticsResponsePtrOutput() DatabaseStatisticsResponsePtrOutput {
	return o.ToDatabaseStatisticsResponsePtrOutputWithContext(context.Background())
}

func (o DatabaseStatisticsResponseOutput) ToDatabaseStatisticsResponsePtrOutputWithContext(ctx context.Context) DatabaseStatisticsResponsePtrOutput {
	return o.ApplyT(func(v DatabaseStatisticsResponse) *DatabaseStatisticsResponse {
		return &v
	}).(DatabaseStatisticsResponsePtrOutput)
}

// The database size - the total size of compressed data and index in bytes.
func (o DatabaseStatisticsResponseOutput) Size() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v DatabaseStatisticsResponse) *float64 { return v.Size }).(pulumi.Float64PtrOutput)
}

type DatabaseStatisticsResponsePtrOutput struct{ *pulumi.OutputState }

func (DatabaseStatisticsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseStatisticsResponse)(nil)).Elem()
}

func (o DatabaseStatisticsResponsePtrOutput) ToDatabaseStatisticsResponsePtrOutput() DatabaseStatisticsResponsePtrOutput {
	return o
}

func (o DatabaseStatisticsResponsePtrOutput) ToDatabaseStatisticsResponsePtrOutputWithContext(ctx context.Context) DatabaseStatisticsResponsePtrOutput {
	return o
}

func (o DatabaseStatisticsResponsePtrOutput) Elem() DatabaseStatisticsResponseOutput {
	return o.ApplyT(func(v *DatabaseStatisticsResponse) DatabaseStatisticsResponse { return *v }).(DatabaseStatisticsResponseOutput)
}

// The database size - the total size of compressed data and index in bytes.
func (o DatabaseStatisticsResponsePtrOutput) Size() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *DatabaseStatisticsResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.Size
	}).(pulumi.Float64PtrOutput)
}

type TrustedExternalTenant struct {
	// GUID representing an external tenant.
	Value *string `pulumi:"value"`
}

// TrustedExternalTenantInput is an input type that accepts TrustedExternalTenantArgs and TrustedExternalTenantOutput values.
// You can construct a concrete instance of `TrustedExternalTenantInput` via:
//
//          TrustedExternalTenantArgs{...}
type TrustedExternalTenantInput interface {
	pulumi.Input

	ToTrustedExternalTenantOutput() TrustedExternalTenantOutput
	ToTrustedExternalTenantOutputWithContext(context.Context) TrustedExternalTenantOutput
}

type TrustedExternalTenantArgs struct {
	// GUID representing an external tenant.
	Value pulumi.StringPtrInput `pulumi:"value"`
}

func (TrustedExternalTenantArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TrustedExternalTenant)(nil)).Elem()
}

func (i TrustedExternalTenantArgs) ToTrustedExternalTenantOutput() TrustedExternalTenantOutput {
	return i.ToTrustedExternalTenantOutputWithContext(context.Background())
}

func (i TrustedExternalTenantArgs) ToTrustedExternalTenantOutputWithContext(ctx context.Context) TrustedExternalTenantOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrustedExternalTenantOutput)
}

// TrustedExternalTenantArrayInput is an input type that accepts TrustedExternalTenantArray and TrustedExternalTenantArrayOutput values.
// You can construct a concrete instance of `TrustedExternalTenantArrayInput` via:
//
//          TrustedExternalTenantArray{ TrustedExternalTenantArgs{...} }
type TrustedExternalTenantArrayInput interface {
	pulumi.Input

	ToTrustedExternalTenantArrayOutput() TrustedExternalTenantArrayOutput
	ToTrustedExternalTenantArrayOutputWithContext(context.Context) TrustedExternalTenantArrayOutput
}

type TrustedExternalTenantArray []TrustedExternalTenantInput

func (TrustedExternalTenantArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TrustedExternalTenant)(nil)).Elem()
}

func (i TrustedExternalTenantArray) ToTrustedExternalTenantArrayOutput() TrustedExternalTenantArrayOutput {
	return i.ToTrustedExternalTenantArrayOutputWithContext(context.Background())
}

func (i TrustedExternalTenantArray) ToTrustedExternalTenantArrayOutputWithContext(ctx context.Context) TrustedExternalTenantArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrustedExternalTenantArrayOutput)
}

type TrustedExternalTenantOutput struct{ *pulumi.OutputState }

func (TrustedExternalTenantOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TrustedExternalTenant)(nil)).Elem()
}

func (o TrustedExternalTenantOutput) ToTrustedExternalTenantOutput() TrustedExternalTenantOutput {
	return o
}

func (o TrustedExternalTenantOutput) ToTrustedExternalTenantOutputWithContext(ctx context.Context) TrustedExternalTenantOutput {
	return o
}

// GUID representing an external tenant.
func (o TrustedExternalTenantOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TrustedExternalTenant) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type TrustedExternalTenantArrayOutput struct{ *pulumi.OutputState }

func (TrustedExternalTenantArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TrustedExternalTenant)(nil)).Elem()
}

func (o TrustedExternalTenantArrayOutput) ToTrustedExternalTenantArrayOutput() TrustedExternalTenantArrayOutput {
	return o
}

func (o TrustedExternalTenantArrayOutput) ToTrustedExternalTenantArrayOutputWithContext(ctx context.Context) TrustedExternalTenantArrayOutput {
	return o
}

func (o TrustedExternalTenantArrayOutput) Index(i pulumi.IntInput) TrustedExternalTenantOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TrustedExternalTenant {
		return vs[0].([]TrustedExternalTenant)[vs[1].(int)]
	}).(TrustedExternalTenantOutput)
}

type TrustedExternalTenantResponse struct {
	// GUID representing an external tenant.
	Value *string `pulumi:"value"`
}

// TrustedExternalTenantResponseInput is an input type that accepts TrustedExternalTenantResponseArgs and TrustedExternalTenantResponseOutput values.
// You can construct a concrete instance of `TrustedExternalTenantResponseInput` via:
//
//          TrustedExternalTenantResponseArgs{...}
type TrustedExternalTenantResponseInput interface {
	pulumi.Input

	ToTrustedExternalTenantResponseOutput() TrustedExternalTenantResponseOutput
	ToTrustedExternalTenantResponseOutputWithContext(context.Context) TrustedExternalTenantResponseOutput
}

type TrustedExternalTenantResponseArgs struct {
	// GUID representing an external tenant.
	Value pulumi.StringPtrInput `pulumi:"value"`
}

func (TrustedExternalTenantResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TrustedExternalTenantResponse)(nil)).Elem()
}

func (i TrustedExternalTenantResponseArgs) ToTrustedExternalTenantResponseOutput() TrustedExternalTenantResponseOutput {
	return i.ToTrustedExternalTenantResponseOutputWithContext(context.Background())
}

func (i TrustedExternalTenantResponseArgs) ToTrustedExternalTenantResponseOutputWithContext(ctx context.Context) TrustedExternalTenantResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrustedExternalTenantResponseOutput)
}

// TrustedExternalTenantResponseArrayInput is an input type that accepts TrustedExternalTenantResponseArray and TrustedExternalTenantResponseArrayOutput values.
// You can construct a concrete instance of `TrustedExternalTenantResponseArrayInput` via:
//
//          TrustedExternalTenantResponseArray{ TrustedExternalTenantResponseArgs{...} }
type TrustedExternalTenantResponseArrayInput interface {
	pulumi.Input

	ToTrustedExternalTenantResponseArrayOutput() TrustedExternalTenantResponseArrayOutput
	ToTrustedExternalTenantResponseArrayOutputWithContext(context.Context) TrustedExternalTenantResponseArrayOutput
}

type TrustedExternalTenantResponseArray []TrustedExternalTenantResponseInput

func (TrustedExternalTenantResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TrustedExternalTenantResponse)(nil)).Elem()
}

func (i TrustedExternalTenantResponseArray) ToTrustedExternalTenantResponseArrayOutput() TrustedExternalTenantResponseArrayOutput {
	return i.ToTrustedExternalTenantResponseArrayOutputWithContext(context.Background())
}

func (i TrustedExternalTenantResponseArray) ToTrustedExternalTenantResponseArrayOutputWithContext(ctx context.Context) TrustedExternalTenantResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrustedExternalTenantResponseArrayOutput)
}

type TrustedExternalTenantResponseOutput struct{ *pulumi.OutputState }

func (TrustedExternalTenantResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TrustedExternalTenantResponse)(nil)).Elem()
}

func (o TrustedExternalTenantResponseOutput) ToTrustedExternalTenantResponseOutput() TrustedExternalTenantResponseOutput {
	return o
}

func (o TrustedExternalTenantResponseOutput) ToTrustedExternalTenantResponseOutputWithContext(ctx context.Context) TrustedExternalTenantResponseOutput {
	return o
}

// GUID representing an external tenant.
func (o TrustedExternalTenantResponseOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TrustedExternalTenantResponse) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type TrustedExternalTenantResponseArrayOutput struct{ *pulumi.OutputState }

func (TrustedExternalTenantResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TrustedExternalTenantResponse)(nil)).Elem()
}

func (o TrustedExternalTenantResponseArrayOutput) ToTrustedExternalTenantResponseArrayOutput() TrustedExternalTenantResponseArrayOutput {
	return o
}

func (o TrustedExternalTenantResponseArrayOutput) ToTrustedExternalTenantResponseArrayOutputWithContext(ctx context.Context) TrustedExternalTenantResponseArrayOutput {
	return o
}

func (o TrustedExternalTenantResponseArrayOutput) Index(i pulumi.IntInput) TrustedExternalTenantResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TrustedExternalTenantResponse {
		return vs[0].([]TrustedExternalTenantResponse)[vs[1].(int)]
	}).(TrustedExternalTenantResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(AzureSkuOutput{})
	pulumi.RegisterOutputType(AzureSkuPtrOutput{})
	pulumi.RegisterOutputType(AzureSkuResponseOutput{})
	pulumi.RegisterOutputType(AzureSkuResponsePtrOutput{})
	pulumi.RegisterOutputType(DatabasePrincipalResponseOutput{})
	pulumi.RegisterOutputType(DatabasePrincipalResponseArrayOutput{})
	pulumi.RegisterOutputType(DatabaseStatisticsOutput{})
	pulumi.RegisterOutputType(DatabaseStatisticsPtrOutput{})
	pulumi.RegisterOutputType(DatabaseStatisticsResponseOutput{})
	pulumi.RegisterOutputType(DatabaseStatisticsResponsePtrOutput{})
	pulumi.RegisterOutputType(TrustedExternalTenantOutput{})
	pulumi.RegisterOutputType(TrustedExternalTenantArrayOutput{})
	pulumi.RegisterOutputType(TrustedExternalTenantResponseOutput{})
	pulumi.RegisterOutputType(TrustedExternalTenantResponseArrayOutput{})
}
