// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20191212

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type ErrorDetail struct {
	// The error's code.
	Code string `pulumi:"code"`
	// Additional error details.
	Details []ErrorDetail `pulumi:"details"`
	// A human readable error message.
	Message string `pulumi:"message"`
	// Indicates which property in the request is responsible for the error.
	Target *string `pulumi:"target"`
}

// ErrorDetailInput is an input type that accepts ErrorDetailArgs and ErrorDetailOutput values.
// You can construct a concrete instance of `ErrorDetailInput` via:
//
//          ErrorDetailArgs{...}
type ErrorDetailInput interface {
	pulumi.Input

	ToErrorDetailOutput() ErrorDetailOutput
	ToErrorDetailOutputWithContext(context.Context) ErrorDetailOutput
}

type ErrorDetailArgs struct {
	// The error's code.
	Code pulumi.StringInput `pulumi:"code"`
	// Additional error details.
	Details ErrorDetailArrayInput `pulumi:"details"`
	// A human readable error message.
	Message pulumi.StringInput `pulumi:"message"`
	// Indicates which property in the request is responsible for the error.
	Target pulumi.StringPtrInput `pulumi:"target"`
}

func (ErrorDetailArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ErrorDetail)(nil)).Elem()
}

func (i ErrorDetailArgs) ToErrorDetailOutput() ErrorDetailOutput {
	return i.ToErrorDetailOutputWithContext(context.Background())
}

func (i ErrorDetailArgs) ToErrorDetailOutputWithContext(ctx context.Context) ErrorDetailOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ErrorDetailOutput)
}

// ErrorDetailArrayInput is an input type that accepts ErrorDetailArray and ErrorDetailArrayOutput values.
// You can construct a concrete instance of `ErrorDetailArrayInput` via:
//
//          ErrorDetailArray{ ErrorDetailArgs{...} }
type ErrorDetailArrayInput interface {
	pulumi.Input

	ToErrorDetailArrayOutput() ErrorDetailArrayOutput
	ToErrorDetailArrayOutputWithContext(context.Context) ErrorDetailArrayOutput
}

type ErrorDetailArray []ErrorDetailInput

func (ErrorDetailArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ErrorDetail)(nil)).Elem()
}

func (i ErrorDetailArray) ToErrorDetailArrayOutput() ErrorDetailArrayOutput {
	return i.ToErrorDetailArrayOutputWithContext(context.Background())
}

func (i ErrorDetailArray) ToErrorDetailArrayOutputWithContext(ctx context.Context) ErrorDetailArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ErrorDetailArrayOutput)
}

type ErrorDetailOutput struct{ *pulumi.OutputState }

func (ErrorDetailOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ErrorDetail)(nil)).Elem()
}

func (o ErrorDetailOutput) ToErrorDetailOutput() ErrorDetailOutput {
	return o
}

func (o ErrorDetailOutput) ToErrorDetailOutputWithContext(ctx context.Context) ErrorDetailOutput {
	return o
}

// The error's code.
func (o ErrorDetailOutput) Code() pulumi.StringOutput {
	return o.ApplyT(func(v ErrorDetail) string { return v.Code }).(pulumi.StringOutput)
}

// Additional error details.
func (o ErrorDetailOutput) Details() ErrorDetailArrayOutput {
	return o.ApplyT(func(v ErrorDetail) []ErrorDetail { return v.Details }).(ErrorDetailArrayOutput)
}

// A human readable error message.
func (o ErrorDetailOutput) Message() pulumi.StringOutput {
	return o.ApplyT(func(v ErrorDetail) string { return v.Message }).(pulumi.StringOutput)
}

// Indicates which property in the request is responsible for the error.
func (o ErrorDetailOutput) Target() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ErrorDetail) *string { return v.Target }).(pulumi.StringPtrOutput)
}

type ErrorDetailArrayOutput struct{ *pulumi.OutputState }

func (ErrorDetailArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ErrorDetail)(nil)).Elem()
}

func (o ErrorDetailArrayOutput) ToErrorDetailArrayOutput() ErrorDetailArrayOutput {
	return o
}

func (o ErrorDetailArrayOutput) ToErrorDetailArrayOutputWithContext(ctx context.Context) ErrorDetailArrayOutput {
	return o
}

func (o ErrorDetailArrayOutput) Index(i pulumi.IntInput) ErrorDetailOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ErrorDetail {
		return vs[0].([]ErrorDetail)[vs[1].(int)]
	}).(ErrorDetailOutput)
}

// Describes a hybrid machine.
type MachineType struct {
	Identity map[string]interface{} `pulumi:"identity"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Hybrid Compute Machine properties
	Properties map[string]interface{} `pulumi:"properties"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type string `pulumi:"type"`
}

// MachineTypeInput is an input type that accepts MachineTypeArgs and MachineTypeOutput values.
// You can construct a concrete instance of `MachineTypeInput` via:
//
//          MachineTypeArgs{...}
type MachineTypeInput interface {
	pulumi.Input

	ToMachineTypeOutput() MachineTypeOutput
	ToMachineTypeOutputWithContext(context.Context) MachineTypeOutput
}

// Describes a hybrid machine.
type MachineTypeArgs struct {
	Identity pulumi.MapInput `pulumi:"identity"`
	// The geo-location where the resource lives
	Location pulumi.StringInput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringInput `pulumi:"name"`
	// Hybrid Compute Machine properties
	Properties pulumi.MapInput `pulumi:"properties"`
	// Resource tags.
	Tags pulumi.StringMapInput `pulumi:"tags"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringInput `pulumi:"type"`
}

func (MachineTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MachineType)(nil)).Elem()
}

func (i MachineTypeArgs) ToMachineTypeOutput() MachineTypeOutput {
	return i.ToMachineTypeOutputWithContext(context.Background())
}

func (i MachineTypeArgs) ToMachineTypeOutputWithContext(ctx context.Context) MachineTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineTypeOutput)
}

// Describes a hybrid machine.
type MachineTypeOutput struct{ *pulumi.OutputState }

func (MachineTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MachineType)(nil)).Elem()
}

func (o MachineTypeOutput) ToMachineTypeOutput() MachineTypeOutput {
	return o
}

func (o MachineTypeOutput) ToMachineTypeOutputWithContext(ctx context.Context) MachineTypeOutput {
	return o
}

func (o MachineTypeOutput) Identity() pulumi.MapOutput {
	return o.ApplyT(func(v MachineType) map[string]interface{} { return v.Identity }).(pulumi.MapOutput)
}

// The geo-location where the resource lives
func (o MachineTypeOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v MachineType) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o MachineTypeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v MachineType) string { return v.Name }).(pulumi.StringOutput)
}

// Hybrid Compute Machine properties
func (o MachineTypeOutput) Properties() pulumi.MapOutput {
	return o.ApplyT(func(v MachineType) map[string]interface{} { return v.Properties }).(pulumi.MapOutput)
}

// Resource tags.
func (o MachineTypeOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v MachineType) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
func (o MachineTypeOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v MachineType) string { return v.Type }).(pulumi.StringOutput)
}

// Describes a Machine Extension.
type MachineExtensionType struct {
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Describes Machine Extension Properties.
	Properties map[string]interface{} `pulumi:"properties"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type string `pulumi:"type"`
}

// MachineExtensionTypeInput is an input type that accepts MachineExtensionTypeArgs and MachineExtensionTypeOutput values.
// You can construct a concrete instance of `MachineExtensionTypeInput` via:
//
//          MachineExtensionTypeArgs{...}
type MachineExtensionTypeInput interface {
	pulumi.Input

	ToMachineExtensionTypeOutput() MachineExtensionTypeOutput
	ToMachineExtensionTypeOutputWithContext(context.Context) MachineExtensionTypeOutput
}

// Describes a Machine Extension.
type MachineExtensionTypeArgs struct {
	// The geo-location where the resource lives
	Location pulumi.StringInput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringInput `pulumi:"name"`
	// Describes Machine Extension Properties.
	Properties pulumi.MapInput `pulumi:"properties"`
	// Resource tags.
	Tags pulumi.StringMapInput `pulumi:"tags"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringInput `pulumi:"type"`
}

func (MachineExtensionTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MachineExtensionType)(nil)).Elem()
}

func (i MachineExtensionTypeArgs) ToMachineExtensionTypeOutput() MachineExtensionTypeOutput {
	return i.ToMachineExtensionTypeOutputWithContext(context.Background())
}

func (i MachineExtensionTypeArgs) ToMachineExtensionTypeOutputWithContext(ctx context.Context) MachineExtensionTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineExtensionTypeOutput)
}

// Describes a Machine Extension.
type MachineExtensionTypeOutput struct{ *pulumi.OutputState }

func (MachineExtensionTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MachineExtensionType)(nil)).Elem()
}

func (o MachineExtensionTypeOutput) ToMachineExtensionTypeOutput() MachineExtensionTypeOutput {
	return o
}

func (o MachineExtensionTypeOutput) ToMachineExtensionTypeOutputWithContext(ctx context.Context) MachineExtensionTypeOutput {
	return o
}

// The geo-location where the resource lives
func (o MachineExtensionTypeOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v MachineExtensionType) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o MachineExtensionTypeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v MachineExtensionType) string { return v.Name }).(pulumi.StringOutput)
}

// Describes Machine Extension Properties.
func (o MachineExtensionTypeOutput) Properties() pulumi.MapOutput {
	return o.ApplyT(func(v MachineExtensionType) map[string]interface{} { return v.Properties }).(pulumi.MapOutput)
}

// Resource tags.
func (o MachineExtensionTypeOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v MachineExtensionType) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
func (o MachineExtensionTypeOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v MachineExtensionType) string { return v.Type }).(pulumi.StringOutput)
}

// Describes the Machine Extension Instance View.
type MachineExtensionInstanceView struct {
	// The machine extension name.
	Name *string `pulumi:"name"`
	// Instance view status.
	Status *MachineExtensionInstanceViewProperties `pulumi:"status"`
	// Specifies the type of the extension; an example is "CustomScriptExtension".
	Type *string `pulumi:"type"`
	// Specifies the version of the script handler.
	TypeHandlerVersion *string `pulumi:"typeHandlerVersion"`
}

// MachineExtensionInstanceViewInput is an input type that accepts MachineExtensionInstanceViewArgs and MachineExtensionInstanceViewOutput values.
// You can construct a concrete instance of `MachineExtensionInstanceViewInput` via:
//
//          MachineExtensionInstanceViewArgs{...}
type MachineExtensionInstanceViewInput interface {
	pulumi.Input

	ToMachineExtensionInstanceViewOutput() MachineExtensionInstanceViewOutput
	ToMachineExtensionInstanceViewOutputWithContext(context.Context) MachineExtensionInstanceViewOutput
}

// Describes the Machine Extension Instance View.
type MachineExtensionInstanceViewArgs struct {
	// The machine extension name.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// Instance view status.
	Status MachineExtensionInstanceViewPropertiesPtrInput `pulumi:"status"`
	// Specifies the type of the extension; an example is "CustomScriptExtension".
	Type pulumi.StringPtrInput `pulumi:"type"`
	// Specifies the version of the script handler.
	TypeHandlerVersion pulumi.StringPtrInput `pulumi:"typeHandlerVersion"`
}

func (MachineExtensionInstanceViewArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MachineExtensionInstanceView)(nil)).Elem()
}

func (i MachineExtensionInstanceViewArgs) ToMachineExtensionInstanceViewOutput() MachineExtensionInstanceViewOutput {
	return i.ToMachineExtensionInstanceViewOutputWithContext(context.Background())
}

func (i MachineExtensionInstanceViewArgs) ToMachineExtensionInstanceViewOutputWithContext(ctx context.Context) MachineExtensionInstanceViewOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineExtensionInstanceViewOutput)
}

// MachineExtensionInstanceViewArrayInput is an input type that accepts MachineExtensionInstanceViewArray and MachineExtensionInstanceViewArrayOutput values.
// You can construct a concrete instance of `MachineExtensionInstanceViewArrayInput` via:
//
//          MachineExtensionInstanceViewArray{ MachineExtensionInstanceViewArgs{...} }
type MachineExtensionInstanceViewArrayInput interface {
	pulumi.Input

	ToMachineExtensionInstanceViewArrayOutput() MachineExtensionInstanceViewArrayOutput
	ToMachineExtensionInstanceViewArrayOutputWithContext(context.Context) MachineExtensionInstanceViewArrayOutput
}

type MachineExtensionInstanceViewArray []MachineExtensionInstanceViewInput

func (MachineExtensionInstanceViewArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MachineExtensionInstanceView)(nil)).Elem()
}

func (i MachineExtensionInstanceViewArray) ToMachineExtensionInstanceViewArrayOutput() MachineExtensionInstanceViewArrayOutput {
	return i.ToMachineExtensionInstanceViewArrayOutputWithContext(context.Background())
}

func (i MachineExtensionInstanceViewArray) ToMachineExtensionInstanceViewArrayOutputWithContext(ctx context.Context) MachineExtensionInstanceViewArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineExtensionInstanceViewArrayOutput)
}

// Describes the Machine Extension Instance View.
type MachineExtensionInstanceViewOutput struct{ *pulumi.OutputState }

func (MachineExtensionInstanceViewOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MachineExtensionInstanceView)(nil)).Elem()
}

func (o MachineExtensionInstanceViewOutput) ToMachineExtensionInstanceViewOutput() MachineExtensionInstanceViewOutput {
	return o
}

func (o MachineExtensionInstanceViewOutput) ToMachineExtensionInstanceViewOutputWithContext(ctx context.Context) MachineExtensionInstanceViewOutput {
	return o
}

// The machine extension name.
func (o MachineExtensionInstanceViewOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MachineExtensionInstanceView) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// Instance view status.
func (o MachineExtensionInstanceViewOutput) Status() MachineExtensionInstanceViewPropertiesPtrOutput {
	return o.ApplyT(func(v MachineExtensionInstanceView) *MachineExtensionInstanceViewProperties { return v.Status }).(MachineExtensionInstanceViewPropertiesPtrOutput)
}

// Specifies the type of the extension; an example is "CustomScriptExtension".
func (o MachineExtensionInstanceViewOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MachineExtensionInstanceView) *string { return v.Type }).(pulumi.StringPtrOutput)
}

// Specifies the version of the script handler.
func (o MachineExtensionInstanceViewOutput) TypeHandlerVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MachineExtensionInstanceView) *string { return v.TypeHandlerVersion }).(pulumi.StringPtrOutput)
}

type MachineExtensionInstanceViewArrayOutput struct{ *pulumi.OutputState }

func (MachineExtensionInstanceViewArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MachineExtensionInstanceView)(nil)).Elem()
}

func (o MachineExtensionInstanceViewArrayOutput) ToMachineExtensionInstanceViewArrayOutput() MachineExtensionInstanceViewArrayOutput {
	return o
}

func (o MachineExtensionInstanceViewArrayOutput) ToMachineExtensionInstanceViewArrayOutputWithContext(ctx context.Context) MachineExtensionInstanceViewArrayOutput {
	return o
}

func (o MachineExtensionInstanceViewArrayOutput) Index(i pulumi.IntInput) MachineExtensionInstanceViewOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MachineExtensionInstanceView {
		return vs[0].([]MachineExtensionInstanceView)[vs[1].(int)]
	}).(MachineExtensionInstanceViewOutput)
}

// Instance view status.
type MachineExtensionInstanceViewProperties struct {
	// The status code.
	Code *string `pulumi:"code"`
	// The short localizable label for the status.
	DisplayStatus *string `pulumi:"displayStatus"`
	// The level code.
	Level *string `pulumi:"level"`
	// The detailed status message, including for alerts and error messages.
	Message *string `pulumi:"message"`
	// The time of the status.
	Time *string `pulumi:"time"`
}

// MachineExtensionInstanceViewPropertiesInput is an input type that accepts MachineExtensionInstanceViewPropertiesArgs and MachineExtensionInstanceViewPropertiesOutput values.
// You can construct a concrete instance of `MachineExtensionInstanceViewPropertiesInput` via:
//
//          MachineExtensionInstanceViewPropertiesArgs{...}
type MachineExtensionInstanceViewPropertiesInput interface {
	pulumi.Input

	ToMachineExtensionInstanceViewPropertiesOutput() MachineExtensionInstanceViewPropertiesOutput
	ToMachineExtensionInstanceViewPropertiesOutputWithContext(context.Context) MachineExtensionInstanceViewPropertiesOutput
}

// Instance view status.
type MachineExtensionInstanceViewPropertiesArgs struct {
	// The status code.
	Code pulumi.StringPtrInput `pulumi:"code"`
	// The short localizable label for the status.
	DisplayStatus pulumi.StringPtrInput `pulumi:"displayStatus"`
	// The level code.
	Level pulumi.StringPtrInput `pulumi:"level"`
	// The detailed status message, including for alerts and error messages.
	Message pulumi.StringPtrInput `pulumi:"message"`
	// The time of the status.
	Time pulumi.StringPtrInput `pulumi:"time"`
}

func (MachineExtensionInstanceViewPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MachineExtensionInstanceViewProperties)(nil)).Elem()
}

func (i MachineExtensionInstanceViewPropertiesArgs) ToMachineExtensionInstanceViewPropertiesOutput() MachineExtensionInstanceViewPropertiesOutput {
	return i.ToMachineExtensionInstanceViewPropertiesOutputWithContext(context.Background())
}

func (i MachineExtensionInstanceViewPropertiesArgs) ToMachineExtensionInstanceViewPropertiesOutputWithContext(ctx context.Context) MachineExtensionInstanceViewPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineExtensionInstanceViewPropertiesOutput)
}

func (i MachineExtensionInstanceViewPropertiesArgs) ToMachineExtensionInstanceViewPropertiesPtrOutput() MachineExtensionInstanceViewPropertiesPtrOutput {
	return i.ToMachineExtensionInstanceViewPropertiesPtrOutputWithContext(context.Background())
}

func (i MachineExtensionInstanceViewPropertiesArgs) ToMachineExtensionInstanceViewPropertiesPtrOutputWithContext(ctx context.Context) MachineExtensionInstanceViewPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineExtensionInstanceViewPropertiesOutput).ToMachineExtensionInstanceViewPropertiesPtrOutputWithContext(ctx)
}

// MachineExtensionInstanceViewPropertiesPtrInput is an input type that accepts MachineExtensionInstanceViewPropertiesArgs, MachineExtensionInstanceViewPropertiesPtr and MachineExtensionInstanceViewPropertiesPtrOutput values.
// You can construct a concrete instance of `MachineExtensionInstanceViewPropertiesPtrInput` via:
//
//          MachineExtensionInstanceViewPropertiesArgs{...}
//
//  or:
//
//          nil
type MachineExtensionInstanceViewPropertiesPtrInput interface {
	pulumi.Input

	ToMachineExtensionInstanceViewPropertiesPtrOutput() MachineExtensionInstanceViewPropertiesPtrOutput
	ToMachineExtensionInstanceViewPropertiesPtrOutputWithContext(context.Context) MachineExtensionInstanceViewPropertiesPtrOutput
}

type machineExtensionInstanceViewPropertiesPtrType MachineExtensionInstanceViewPropertiesArgs

func MachineExtensionInstanceViewPropertiesPtr(v *MachineExtensionInstanceViewPropertiesArgs) MachineExtensionInstanceViewPropertiesPtrInput {
	return (*machineExtensionInstanceViewPropertiesPtrType)(v)
}

func (*machineExtensionInstanceViewPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MachineExtensionInstanceViewProperties)(nil)).Elem()
}

func (i *machineExtensionInstanceViewPropertiesPtrType) ToMachineExtensionInstanceViewPropertiesPtrOutput() MachineExtensionInstanceViewPropertiesPtrOutput {
	return i.ToMachineExtensionInstanceViewPropertiesPtrOutputWithContext(context.Background())
}

func (i *machineExtensionInstanceViewPropertiesPtrType) ToMachineExtensionInstanceViewPropertiesPtrOutputWithContext(ctx context.Context) MachineExtensionInstanceViewPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MachineExtensionInstanceViewPropertiesPtrOutput)
}

// Instance view status.
type MachineExtensionInstanceViewPropertiesOutput struct{ *pulumi.OutputState }

func (MachineExtensionInstanceViewPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MachineExtensionInstanceViewProperties)(nil)).Elem()
}

func (o MachineExtensionInstanceViewPropertiesOutput) ToMachineExtensionInstanceViewPropertiesOutput() MachineExtensionInstanceViewPropertiesOutput {
	return o
}

func (o MachineExtensionInstanceViewPropertiesOutput) ToMachineExtensionInstanceViewPropertiesOutputWithContext(ctx context.Context) MachineExtensionInstanceViewPropertiesOutput {
	return o
}

func (o MachineExtensionInstanceViewPropertiesOutput) ToMachineExtensionInstanceViewPropertiesPtrOutput() MachineExtensionInstanceViewPropertiesPtrOutput {
	return o.ToMachineExtensionInstanceViewPropertiesPtrOutputWithContext(context.Background())
}

func (o MachineExtensionInstanceViewPropertiesOutput) ToMachineExtensionInstanceViewPropertiesPtrOutputWithContext(ctx context.Context) MachineExtensionInstanceViewPropertiesPtrOutput {
	return o.ApplyT(func(v MachineExtensionInstanceViewProperties) *MachineExtensionInstanceViewProperties {
		return &v
	}).(MachineExtensionInstanceViewPropertiesPtrOutput)
}

// The status code.
func (o MachineExtensionInstanceViewPropertiesOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MachineExtensionInstanceViewProperties) *string { return v.Code }).(pulumi.StringPtrOutput)
}

// The short localizable label for the status.
func (o MachineExtensionInstanceViewPropertiesOutput) DisplayStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MachineExtensionInstanceViewProperties) *string { return v.DisplayStatus }).(pulumi.StringPtrOutput)
}

// The level code.
func (o MachineExtensionInstanceViewPropertiesOutput) Level() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MachineExtensionInstanceViewProperties) *string { return v.Level }).(pulumi.StringPtrOutput)
}

// The detailed status message, including for alerts and error messages.
func (o MachineExtensionInstanceViewPropertiesOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MachineExtensionInstanceViewProperties) *string { return v.Message }).(pulumi.StringPtrOutput)
}

// The time of the status.
func (o MachineExtensionInstanceViewPropertiesOutput) Time() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MachineExtensionInstanceViewProperties) *string { return v.Time }).(pulumi.StringPtrOutput)
}

type MachineExtensionInstanceViewPropertiesPtrOutput struct{ *pulumi.OutputState }

func (MachineExtensionInstanceViewPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MachineExtensionInstanceViewProperties)(nil)).Elem()
}

func (o MachineExtensionInstanceViewPropertiesPtrOutput) ToMachineExtensionInstanceViewPropertiesPtrOutput() MachineExtensionInstanceViewPropertiesPtrOutput {
	return o
}

func (o MachineExtensionInstanceViewPropertiesPtrOutput) ToMachineExtensionInstanceViewPropertiesPtrOutputWithContext(ctx context.Context) MachineExtensionInstanceViewPropertiesPtrOutput {
	return o
}

func (o MachineExtensionInstanceViewPropertiesPtrOutput) Elem() MachineExtensionInstanceViewPropertiesOutput {
	return o.ApplyT(func(v *MachineExtensionInstanceViewProperties) MachineExtensionInstanceViewProperties { return *v }).(MachineExtensionInstanceViewPropertiesOutput)
}

// The status code.
func (o MachineExtensionInstanceViewPropertiesPtrOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineExtensionInstanceViewProperties) *string {
		if v == nil {
			return nil
		}
		return v.Code
	}).(pulumi.StringPtrOutput)
}

// The short localizable label for the status.
func (o MachineExtensionInstanceViewPropertiesPtrOutput) DisplayStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineExtensionInstanceViewProperties) *string {
		if v == nil {
			return nil
		}
		return v.DisplayStatus
	}).(pulumi.StringPtrOutput)
}

// The level code.
func (o MachineExtensionInstanceViewPropertiesPtrOutput) Level() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineExtensionInstanceViewProperties) *string {
		if v == nil {
			return nil
		}
		return v.Level
	}).(pulumi.StringPtrOutput)
}

// The detailed status message, including for alerts and error messages.
func (o MachineExtensionInstanceViewPropertiesPtrOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineExtensionInstanceViewProperties) *string {
		if v == nil {
			return nil
		}
		return v.Message
	}).(pulumi.StringPtrOutput)
}

// The time of the status.
func (o MachineExtensionInstanceViewPropertiesPtrOutput) Time() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MachineExtensionInstanceViewProperties) *string {
		if v == nil {
			return nil
		}
		return v.Time
	}).(pulumi.StringPtrOutput)
}

// Metadata pertaining to the geographic location of the resource.
type LocationData struct {
	// The city or locality where the resource is located.
	City *string `pulumi:"city"`
	// The country or region where the resource is located
	CountryOrRegion *string `pulumi:"countryOrRegion"`
	// The district, state, or province where the resource is located.
	District *string `pulumi:"district"`
	// A canonical name for the geographic or physical location.
	Name string `pulumi:"name"`
}

// LocationDataInput is an input type that accepts LocationDataArgs and LocationDataOutput values.
// You can construct a concrete instance of `LocationDataInput` via:
//
//          LocationDataArgs{...}
type LocationDataInput interface {
	pulumi.Input

	ToLocationDataOutput() LocationDataOutput
	ToLocationDataOutputWithContext(context.Context) LocationDataOutput
}

// Metadata pertaining to the geographic location of the resource.
type LocationDataArgs struct {
	// The city or locality where the resource is located.
	City pulumi.StringPtrInput `pulumi:"city"`
	// The country or region where the resource is located
	CountryOrRegion pulumi.StringPtrInput `pulumi:"countryOrRegion"`
	// The district, state, or province where the resource is located.
	District pulumi.StringPtrInput `pulumi:"district"`
	// A canonical name for the geographic or physical location.
	Name pulumi.StringInput `pulumi:"name"`
}

func (LocationDataArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LocationData)(nil)).Elem()
}

func (i LocationDataArgs) ToLocationDataOutput() LocationDataOutput {
	return i.ToLocationDataOutputWithContext(context.Background())
}

func (i LocationDataArgs) ToLocationDataOutputWithContext(ctx context.Context) LocationDataOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocationDataOutput)
}

func (i LocationDataArgs) ToLocationDataPtrOutput() LocationDataPtrOutput {
	return i.ToLocationDataPtrOutputWithContext(context.Background())
}

func (i LocationDataArgs) ToLocationDataPtrOutputWithContext(ctx context.Context) LocationDataPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocationDataOutput).ToLocationDataPtrOutputWithContext(ctx)
}

// LocationDataPtrInput is an input type that accepts LocationDataArgs, LocationDataPtr and LocationDataPtrOutput values.
// You can construct a concrete instance of `LocationDataPtrInput` via:
//
//          LocationDataArgs{...}
//
//  or:
//
//          nil
type LocationDataPtrInput interface {
	pulumi.Input

	ToLocationDataPtrOutput() LocationDataPtrOutput
	ToLocationDataPtrOutputWithContext(context.Context) LocationDataPtrOutput
}

type locationDataPtrType LocationDataArgs

func LocationDataPtr(v *LocationDataArgs) LocationDataPtrInput {
	return (*locationDataPtrType)(v)
}

func (*locationDataPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LocationData)(nil)).Elem()
}

func (i *locationDataPtrType) ToLocationDataPtrOutput() LocationDataPtrOutput {
	return i.ToLocationDataPtrOutputWithContext(context.Background())
}

func (i *locationDataPtrType) ToLocationDataPtrOutputWithContext(ctx context.Context) LocationDataPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocationDataPtrOutput)
}

// Metadata pertaining to the geographic location of the resource.
type LocationDataOutput struct{ *pulumi.OutputState }

func (LocationDataOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LocationData)(nil)).Elem()
}

func (o LocationDataOutput) ToLocationDataOutput() LocationDataOutput {
	return o
}

func (o LocationDataOutput) ToLocationDataOutputWithContext(ctx context.Context) LocationDataOutput {
	return o
}

func (o LocationDataOutput) ToLocationDataPtrOutput() LocationDataPtrOutput {
	return o.ToLocationDataPtrOutputWithContext(context.Background())
}

func (o LocationDataOutput) ToLocationDataPtrOutputWithContext(ctx context.Context) LocationDataPtrOutput {
	return o.ApplyT(func(v LocationData) *LocationData {
		return &v
	}).(LocationDataPtrOutput)
}

// The city or locality where the resource is located.
func (o LocationDataOutput) City() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LocationData) *string { return v.City }).(pulumi.StringPtrOutput)
}

// The country or region where the resource is located
func (o LocationDataOutput) CountryOrRegion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LocationData) *string { return v.CountryOrRegion }).(pulumi.StringPtrOutput)
}

// The district, state, or province where the resource is located.
func (o LocationDataOutput) District() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LocationData) *string { return v.District }).(pulumi.StringPtrOutput)
}

// A canonical name for the geographic or physical location.
func (o LocationDataOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LocationData) string { return v.Name }).(pulumi.StringOutput)
}

type LocationDataPtrOutput struct{ *pulumi.OutputState }

func (LocationDataPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LocationData)(nil)).Elem()
}

func (o LocationDataPtrOutput) ToLocationDataPtrOutput() LocationDataPtrOutput {
	return o
}

func (o LocationDataPtrOutput) ToLocationDataPtrOutputWithContext(ctx context.Context) LocationDataPtrOutput {
	return o
}

func (o LocationDataPtrOutput) Elem() LocationDataOutput {
	return o.ApplyT(func(v *LocationData) LocationData { return *v }).(LocationDataOutput)
}

// The city or locality where the resource is located.
func (o LocationDataPtrOutput) City() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocationData) *string {
		if v == nil {
			return nil
		}
		return v.City
	}).(pulumi.StringPtrOutput)
}

// The country or region where the resource is located
func (o LocationDataPtrOutput) CountryOrRegion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocationData) *string {
		if v == nil {
			return nil
		}
		return v.CountryOrRegion
	}).(pulumi.StringPtrOutput)
}

// The district, state, or province where the resource is located.
func (o LocationDataPtrOutput) District() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocationData) *string {
		if v == nil {
			return nil
		}
		return v.District
	}).(pulumi.StringPtrOutput)
}

// A canonical name for the geographic or physical location.
func (o LocationDataPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LocationData) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ErrorDetailOutput{})
	pulumi.RegisterOutputType(ErrorDetailArrayOutput{})
	pulumi.RegisterOutputType(MachineTypeOutput{})
	pulumi.RegisterOutputType(MachineExtensionTypeOutput{})
	pulumi.RegisterOutputType(MachineExtensionInstanceViewOutput{})
	pulumi.RegisterOutputType(MachineExtensionInstanceViewArrayOutput{})
	pulumi.RegisterOutputType(MachineExtensionInstanceViewPropertiesOutput{})
	pulumi.RegisterOutputType(MachineExtensionInstanceViewPropertiesPtrOutput{})
	pulumi.RegisterOutputType(LocationDataOutput{})
	pulumi.RegisterOutputType(LocationDataPtrOutput{})
}
