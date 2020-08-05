// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20170601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Customer subscription.
type CustomerSubscriptionType struct {
	// The entity tag used for optimistic concurrency when modifying the resource.
	Etag *string `pulumi:"etag"`
	// Name of the resource.
	Name string `pulumi:"name"`
	// Customer subscription properties.
	Properties CustomerSubscriptionPropertiesResponse `pulumi:"properties"`
	// Type of Resource.
	Type string `pulumi:"type"`
}

// CustomerSubscriptionTypeInput is an input type that accepts CustomerSubscriptionTypeArgs and CustomerSubscriptionTypeOutput values.
// You can construct a concrete instance of `CustomerSubscriptionTypeInput` via:
//
//          CustomerSubscriptionTypeArgs{...}
type CustomerSubscriptionTypeInput interface {
	pulumi.Input

	ToCustomerSubscriptionTypeOutput() CustomerSubscriptionTypeOutput
	ToCustomerSubscriptionTypeOutputWithContext(context.Context) CustomerSubscriptionTypeOutput
}

// Customer subscription.
type CustomerSubscriptionTypeArgs struct {
	// The entity tag used for optimistic concurrency when modifying the resource.
	Etag pulumi.StringPtrInput `pulumi:"etag"`
	// Name of the resource.
	Name pulumi.StringInput `pulumi:"name"`
	// Customer subscription properties.
	Properties CustomerSubscriptionPropertiesResponseInput `pulumi:"properties"`
	// Type of Resource.
	Type pulumi.StringInput `pulumi:"type"`
}

func (CustomerSubscriptionTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomerSubscriptionType)(nil)).Elem()
}

func (i CustomerSubscriptionTypeArgs) ToCustomerSubscriptionTypeOutput() CustomerSubscriptionTypeOutput {
	return i.ToCustomerSubscriptionTypeOutputWithContext(context.Background())
}

func (i CustomerSubscriptionTypeArgs) ToCustomerSubscriptionTypeOutputWithContext(ctx context.Context) CustomerSubscriptionTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomerSubscriptionTypeOutput)
}

// Customer subscription.
type CustomerSubscriptionTypeOutput struct{ *pulumi.OutputState }

func (CustomerSubscriptionTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomerSubscriptionType)(nil)).Elem()
}

func (o CustomerSubscriptionTypeOutput) ToCustomerSubscriptionTypeOutput() CustomerSubscriptionTypeOutput {
	return o
}

func (o CustomerSubscriptionTypeOutput) ToCustomerSubscriptionTypeOutputWithContext(ctx context.Context) CustomerSubscriptionTypeOutput {
	return o
}

// The entity tag used for optimistic concurrency when modifying the resource.
func (o CustomerSubscriptionTypeOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomerSubscriptionType) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

// Name of the resource.
func (o CustomerSubscriptionTypeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v CustomerSubscriptionType) string { return v.Name }).(pulumi.StringOutput)
}

// Customer subscription properties.
func (o CustomerSubscriptionTypeOutput) Properties() CustomerSubscriptionPropertiesResponseOutput {
	return o.ApplyT(func(v CustomerSubscriptionType) CustomerSubscriptionPropertiesResponse { return v.Properties }).(CustomerSubscriptionPropertiesResponseOutput)
}

// Type of Resource.
func (o CustomerSubscriptionTypeOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v CustomerSubscriptionType) string { return v.Type }).(pulumi.StringOutput)
}

// Customer subscription properties.
type CustomerSubscriptionPropertiesResponse struct {
	// Tenant Id.
	TenantId *string `pulumi:"tenantId"`
}

// CustomerSubscriptionPropertiesResponseInput is an input type that accepts CustomerSubscriptionPropertiesResponseArgs and CustomerSubscriptionPropertiesResponseOutput values.
// You can construct a concrete instance of `CustomerSubscriptionPropertiesResponseInput` via:
//
//          CustomerSubscriptionPropertiesResponseArgs{...}
type CustomerSubscriptionPropertiesResponseInput interface {
	pulumi.Input

	ToCustomerSubscriptionPropertiesResponseOutput() CustomerSubscriptionPropertiesResponseOutput
	ToCustomerSubscriptionPropertiesResponseOutputWithContext(context.Context) CustomerSubscriptionPropertiesResponseOutput
}

// Customer subscription properties.
type CustomerSubscriptionPropertiesResponseArgs struct {
	// Tenant Id.
	TenantId pulumi.StringPtrInput `pulumi:"tenantId"`
}

func (CustomerSubscriptionPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomerSubscriptionPropertiesResponse)(nil)).Elem()
}

func (i CustomerSubscriptionPropertiesResponseArgs) ToCustomerSubscriptionPropertiesResponseOutput() CustomerSubscriptionPropertiesResponseOutput {
	return i.ToCustomerSubscriptionPropertiesResponseOutputWithContext(context.Background())
}

func (i CustomerSubscriptionPropertiesResponseArgs) ToCustomerSubscriptionPropertiesResponseOutputWithContext(ctx context.Context) CustomerSubscriptionPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomerSubscriptionPropertiesResponseOutput)
}

func (i CustomerSubscriptionPropertiesResponseArgs) ToCustomerSubscriptionPropertiesResponsePtrOutput() CustomerSubscriptionPropertiesResponsePtrOutput {
	return i.ToCustomerSubscriptionPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i CustomerSubscriptionPropertiesResponseArgs) ToCustomerSubscriptionPropertiesResponsePtrOutputWithContext(ctx context.Context) CustomerSubscriptionPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomerSubscriptionPropertiesResponseOutput).ToCustomerSubscriptionPropertiesResponsePtrOutputWithContext(ctx)
}

// CustomerSubscriptionPropertiesResponsePtrInput is an input type that accepts CustomerSubscriptionPropertiesResponseArgs, CustomerSubscriptionPropertiesResponsePtr and CustomerSubscriptionPropertiesResponsePtrOutput values.
// You can construct a concrete instance of `CustomerSubscriptionPropertiesResponsePtrInput` via:
//
//          CustomerSubscriptionPropertiesResponseArgs{...}
//
//  or:
//
//          nil
type CustomerSubscriptionPropertiesResponsePtrInput interface {
	pulumi.Input

	ToCustomerSubscriptionPropertiesResponsePtrOutput() CustomerSubscriptionPropertiesResponsePtrOutput
	ToCustomerSubscriptionPropertiesResponsePtrOutputWithContext(context.Context) CustomerSubscriptionPropertiesResponsePtrOutput
}

type customerSubscriptionPropertiesResponsePtrType CustomerSubscriptionPropertiesResponseArgs

func CustomerSubscriptionPropertiesResponsePtr(v *CustomerSubscriptionPropertiesResponseArgs) CustomerSubscriptionPropertiesResponsePtrInput {
	return (*customerSubscriptionPropertiesResponsePtrType)(v)
}

func (*customerSubscriptionPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomerSubscriptionPropertiesResponse)(nil)).Elem()
}

func (i *customerSubscriptionPropertiesResponsePtrType) ToCustomerSubscriptionPropertiesResponsePtrOutput() CustomerSubscriptionPropertiesResponsePtrOutput {
	return i.ToCustomerSubscriptionPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *customerSubscriptionPropertiesResponsePtrType) ToCustomerSubscriptionPropertiesResponsePtrOutputWithContext(ctx context.Context) CustomerSubscriptionPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomerSubscriptionPropertiesResponsePtrOutput)
}

// Customer subscription properties.
type CustomerSubscriptionPropertiesResponseOutput struct{ *pulumi.OutputState }

func (CustomerSubscriptionPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomerSubscriptionPropertiesResponse)(nil)).Elem()
}

func (o CustomerSubscriptionPropertiesResponseOutput) ToCustomerSubscriptionPropertiesResponseOutput() CustomerSubscriptionPropertiesResponseOutput {
	return o
}

func (o CustomerSubscriptionPropertiesResponseOutput) ToCustomerSubscriptionPropertiesResponseOutputWithContext(ctx context.Context) CustomerSubscriptionPropertiesResponseOutput {
	return o
}

func (o CustomerSubscriptionPropertiesResponseOutput) ToCustomerSubscriptionPropertiesResponsePtrOutput() CustomerSubscriptionPropertiesResponsePtrOutput {
	return o.ToCustomerSubscriptionPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o CustomerSubscriptionPropertiesResponseOutput) ToCustomerSubscriptionPropertiesResponsePtrOutputWithContext(ctx context.Context) CustomerSubscriptionPropertiesResponsePtrOutput {
	return o.ApplyT(func(v CustomerSubscriptionPropertiesResponse) *CustomerSubscriptionPropertiesResponse {
		return &v
	}).(CustomerSubscriptionPropertiesResponsePtrOutput)
}

// Tenant Id.
func (o CustomerSubscriptionPropertiesResponseOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomerSubscriptionPropertiesResponse) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

type CustomerSubscriptionPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (CustomerSubscriptionPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomerSubscriptionPropertiesResponse)(nil)).Elem()
}

func (o CustomerSubscriptionPropertiesResponsePtrOutput) ToCustomerSubscriptionPropertiesResponsePtrOutput() CustomerSubscriptionPropertiesResponsePtrOutput {
	return o
}

func (o CustomerSubscriptionPropertiesResponsePtrOutput) ToCustomerSubscriptionPropertiesResponsePtrOutputWithContext(ctx context.Context) CustomerSubscriptionPropertiesResponsePtrOutput {
	return o
}

func (o CustomerSubscriptionPropertiesResponsePtrOutput) Elem() CustomerSubscriptionPropertiesResponseOutput {
	return o.ApplyT(func(v *CustomerSubscriptionPropertiesResponse) CustomerSubscriptionPropertiesResponse { return *v }).(CustomerSubscriptionPropertiesResponseOutput)
}

// Tenant Id.
func (o CustomerSubscriptionPropertiesResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomerSubscriptionPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.TenantId
	}).(pulumi.StringPtrOutput)
}

// Data disk image.
type DataDiskImageResponse struct {
	// The LUN.
	Lun int `pulumi:"lun"`
	// SAS key for source blob.
	SourceBlobSasUri string `pulumi:"sourceBlobSasUri"`
}

// DataDiskImageResponseInput is an input type that accepts DataDiskImageResponseArgs and DataDiskImageResponseOutput values.
// You can construct a concrete instance of `DataDiskImageResponseInput` via:
//
//          DataDiskImageResponseArgs{...}
type DataDiskImageResponseInput interface {
	pulumi.Input

	ToDataDiskImageResponseOutput() DataDiskImageResponseOutput
	ToDataDiskImageResponseOutputWithContext(context.Context) DataDiskImageResponseOutput
}

// Data disk image.
type DataDiskImageResponseArgs struct {
	// The LUN.
	Lun pulumi.IntInput `pulumi:"lun"`
	// SAS key for source blob.
	SourceBlobSasUri pulumi.StringInput `pulumi:"sourceBlobSasUri"`
}

func (DataDiskImageResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataDiskImageResponse)(nil)).Elem()
}

func (i DataDiskImageResponseArgs) ToDataDiskImageResponseOutput() DataDiskImageResponseOutput {
	return i.ToDataDiskImageResponseOutputWithContext(context.Background())
}

func (i DataDiskImageResponseArgs) ToDataDiskImageResponseOutputWithContext(ctx context.Context) DataDiskImageResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataDiskImageResponseOutput)
}

// DataDiskImageResponseArrayInput is an input type that accepts DataDiskImageResponseArray and DataDiskImageResponseArrayOutput values.
// You can construct a concrete instance of `DataDiskImageResponseArrayInput` via:
//
//          DataDiskImageResponseArray{ DataDiskImageResponseArgs{...} }
type DataDiskImageResponseArrayInput interface {
	pulumi.Input

	ToDataDiskImageResponseArrayOutput() DataDiskImageResponseArrayOutput
	ToDataDiskImageResponseArrayOutputWithContext(context.Context) DataDiskImageResponseArrayOutput
}

type DataDiskImageResponseArray []DataDiskImageResponseInput

func (DataDiskImageResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataDiskImageResponse)(nil)).Elem()
}

func (i DataDiskImageResponseArray) ToDataDiskImageResponseArrayOutput() DataDiskImageResponseArrayOutput {
	return i.ToDataDiskImageResponseArrayOutputWithContext(context.Background())
}

func (i DataDiskImageResponseArray) ToDataDiskImageResponseArrayOutputWithContext(ctx context.Context) DataDiskImageResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataDiskImageResponseArrayOutput)
}

// Data disk image.
type DataDiskImageResponseOutput struct{ *pulumi.OutputState }

func (DataDiskImageResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataDiskImageResponse)(nil)).Elem()
}

func (o DataDiskImageResponseOutput) ToDataDiskImageResponseOutput() DataDiskImageResponseOutput {
	return o
}

func (o DataDiskImageResponseOutput) ToDataDiskImageResponseOutputWithContext(ctx context.Context) DataDiskImageResponseOutput {
	return o
}

// The LUN.
func (o DataDiskImageResponseOutput) Lun() pulumi.IntOutput {
	return o.ApplyT(func(v DataDiskImageResponse) int { return v.Lun }).(pulumi.IntOutput)
}

// SAS key for source blob.
func (o DataDiskImageResponseOutput) SourceBlobSasUri() pulumi.StringOutput {
	return o.ApplyT(func(v DataDiskImageResponse) string { return v.SourceBlobSasUri }).(pulumi.StringOutput)
}

type DataDiskImageResponseArrayOutput struct{ *pulumi.OutputState }

func (DataDiskImageResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DataDiskImageResponse)(nil)).Elem()
}

func (o DataDiskImageResponseArrayOutput) ToDataDiskImageResponseArrayOutput() DataDiskImageResponseArrayOutput {
	return o
}

func (o DataDiskImageResponseArrayOutput) ToDataDiskImageResponseArrayOutputWithContext(ctx context.Context) DataDiskImageResponseArrayOutput {
	return o
}

func (o DataDiskImageResponseArrayOutput) Index(i pulumi.IntInput) DataDiskImageResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DataDiskImageResponse {
		return vs[0].([]DataDiskImageResponse)[vs[1].(int)]
	}).(DataDiskImageResponseOutput)
}

// Product information.
type ExtendedProductPropertiesResponse struct {
	// Specifies kind of compute role included in the package.
	ComputeRole string `pulumi:"computeRole"`
	// List of attached data disks.
	DataDiskImages []DataDiskImageResponse `pulumi:"dataDiskImages"`
	// Specifies if product is a Virtual Machine Extension.
	IsSystemExtension bool `pulumi:"isSystemExtension"`
	// OS disk image used by product.
	OsDiskImage OsDiskImageResponse `pulumi:"osDiskImage"`
	// Specifies a download location where content can be downloaded from.
	SourceBlob UriResponse `pulumi:"sourceBlob"`
	// Indicates if specified product supports multiple extensions.
	SupportMultipleExtensions bool `pulumi:"supportMultipleExtensions"`
	// Specifies product version.
	Version string `pulumi:"version"`
	// Specifies operating system used by the product.
	VmOsType string `pulumi:"vmOsType"`
	// Indicates if virtual machine Scale Set is enabled in the specified product.
	VmScaleSetEnabled bool `pulumi:"vmScaleSetEnabled"`
}

// ExtendedProductPropertiesResponseInput is an input type that accepts ExtendedProductPropertiesResponseArgs and ExtendedProductPropertiesResponseOutput values.
// You can construct a concrete instance of `ExtendedProductPropertiesResponseInput` via:
//
//          ExtendedProductPropertiesResponseArgs{...}
type ExtendedProductPropertiesResponseInput interface {
	pulumi.Input

	ToExtendedProductPropertiesResponseOutput() ExtendedProductPropertiesResponseOutput
	ToExtendedProductPropertiesResponseOutputWithContext(context.Context) ExtendedProductPropertiesResponseOutput
}

// Product information.
type ExtendedProductPropertiesResponseArgs struct {
	// Specifies kind of compute role included in the package.
	ComputeRole pulumi.StringInput `pulumi:"computeRole"`
	// List of attached data disks.
	DataDiskImages DataDiskImageResponseArrayInput `pulumi:"dataDiskImages"`
	// Specifies if product is a Virtual Machine Extension.
	IsSystemExtension pulumi.BoolInput `pulumi:"isSystemExtension"`
	// OS disk image used by product.
	OsDiskImage OsDiskImageResponseInput `pulumi:"osDiskImage"`
	// Specifies a download location where content can be downloaded from.
	SourceBlob UriResponseInput `pulumi:"sourceBlob"`
	// Indicates if specified product supports multiple extensions.
	SupportMultipleExtensions pulumi.BoolInput `pulumi:"supportMultipleExtensions"`
	// Specifies product version.
	Version pulumi.StringInput `pulumi:"version"`
	// Specifies operating system used by the product.
	VmOsType pulumi.StringInput `pulumi:"vmOsType"`
	// Indicates if virtual machine Scale Set is enabled in the specified product.
	VmScaleSetEnabled pulumi.BoolInput `pulumi:"vmScaleSetEnabled"`
}

func (ExtendedProductPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ExtendedProductPropertiesResponse)(nil)).Elem()
}

func (i ExtendedProductPropertiesResponseArgs) ToExtendedProductPropertiesResponseOutput() ExtendedProductPropertiesResponseOutput {
	return i.ToExtendedProductPropertiesResponseOutputWithContext(context.Background())
}

func (i ExtendedProductPropertiesResponseArgs) ToExtendedProductPropertiesResponseOutputWithContext(ctx context.Context) ExtendedProductPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtendedProductPropertiesResponseOutput)
}

// Product information.
type ExtendedProductPropertiesResponseOutput struct{ *pulumi.OutputState }

func (ExtendedProductPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExtendedProductPropertiesResponse)(nil)).Elem()
}

func (o ExtendedProductPropertiesResponseOutput) ToExtendedProductPropertiesResponseOutput() ExtendedProductPropertiesResponseOutput {
	return o
}

func (o ExtendedProductPropertiesResponseOutput) ToExtendedProductPropertiesResponseOutputWithContext(ctx context.Context) ExtendedProductPropertiesResponseOutput {
	return o
}

// Specifies kind of compute role included in the package.
func (o ExtendedProductPropertiesResponseOutput) ComputeRole() pulumi.StringOutput {
	return o.ApplyT(func(v ExtendedProductPropertiesResponse) string { return v.ComputeRole }).(pulumi.StringOutput)
}

// List of attached data disks.
func (o ExtendedProductPropertiesResponseOutput) DataDiskImages() DataDiskImageResponseArrayOutput {
	return o.ApplyT(func(v ExtendedProductPropertiesResponse) []DataDiskImageResponse { return v.DataDiskImages }).(DataDiskImageResponseArrayOutput)
}

// Specifies if product is a Virtual Machine Extension.
func (o ExtendedProductPropertiesResponseOutput) IsSystemExtension() pulumi.BoolOutput {
	return o.ApplyT(func(v ExtendedProductPropertiesResponse) bool { return v.IsSystemExtension }).(pulumi.BoolOutput)
}

// OS disk image used by product.
func (o ExtendedProductPropertiesResponseOutput) OsDiskImage() OsDiskImageResponseOutput {
	return o.ApplyT(func(v ExtendedProductPropertiesResponse) OsDiskImageResponse { return v.OsDiskImage }).(OsDiskImageResponseOutput)
}

// Specifies a download location where content can be downloaded from.
func (o ExtendedProductPropertiesResponseOutput) SourceBlob() UriResponseOutput {
	return o.ApplyT(func(v ExtendedProductPropertiesResponse) UriResponse { return v.SourceBlob }).(UriResponseOutput)
}

// Indicates if specified product supports multiple extensions.
func (o ExtendedProductPropertiesResponseOutput) SupportMultipleExtensions() pulumi.BoolOutput {
	return o.ApplyT(func(v ExtendedProductPropertiesResponse) bool { return v.SupportMultipleExtensions }).(pulumi.BoolOutput)
}

// Specifies product version.
func (o ExtendedProductPropertiesResponseOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v ExtendedProductPropertiesResponse) string { return v.Version }).(pulumi.StringOutput)
}

// Specifies operating system used by the product.
func (o ExtendedProductPropertiesResponseOutput) VmOsType() pulumi.StringOutput {
	return o.ApplyT(func(v ExtendedProductPropertiesResponse) string { return v.VmOsType }).(pulumi.StringOutput)
}

// Indicates if virtual machine Scale Set is enabled in the specified product.
func (o ExtendedProductPropertiesResponseOutput) VmScaleSetEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v ExtendedProductPropertiesResponse) bool { return v.VmScaleSetEnabled }).(pulumi.BoolOutput)
}

// OS disk image.
type OsDiskImageResponse struct {
	// OS operating system type.
	OperatingSystem string `pulumi:"operatingSystem"`
	// SAS key for source blob.
	SourceBlobSasUri string `pulumi:"sourceBlobSasUri"`
}

// OsDiskImageResponseInput is an input type that accepts OsDiskImageResponseArgs and OsDiskImageResponseOutput values.
// You can construct a concrete instance of `OsDiskImageResponseInput` via:
//
//          OsDiskImageResponseArgs{...}
type OsDiskImageResponseInput interface {
	pulumi.Input

	ToOsDiskImageResponseOutput() OsDiskImageResponseOutput
	ToOsDiskImageResponseOutputWithContext(context.Context) OsDiskImageResponseOutput
}

// OS disk image.
type OsDiskImageResponseArgs struct {
	// OS operating system type.
	OperatingSystem pulumi.StringInput `pulumi:"operatingSystem"`
	// SAS key for source blob.
	SourceBlobSasUri pulumi.StringInput `pulumi:"sourceBlobSasUri"`
}

func (OsDiskImageResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OsDiskImageResponse)(nil)).Elem()
}

func (i OsDiskImageResponseArgs) ToOsDiskImageResponseOutput() OsDiskImageResponseOutput {
	return i.ToOsDiskImageResponseOutputWithContext(context.Background())
}

func (i OsDiskImageResponseArgs) ToOsDiskImageResponseOutputWithContext(ctx context.Context) OsDiskImageResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OsDiskImageResponseOutput)
}

// OS disk image.
type OsDiskImageResponseOutput struct{ *pulumi.OutputState }

func (OsDiskImageResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OsDiskImageResponse)(nil)).Elem()
}

func (o OsDiskImageResponseOutput) ToOsDiskImageResponseOutput() OsDiskImageResponseOutput {
	return o
}

func (o OsDiskImageResponseOutput) ToOsDiskImageResponseOutputWithContext(ctx context.Context) OsDiskImageResponseOutput {
	return o
}

// OS operating system type.
func (o OsDiskImageResponseOutput) OperatingSystem() pulumi.StringOutput {
	return o.ApplyT(func(v OsDiskImageResponse) string { return v.OperatingSystem }).(pulumi.StringOutput)
}

// SAS key for source blob.
func (o OsDiskImageResponseOutput) SourceBlobSasUri() pulumi.StringOutput {
	return o.ApplyT(func(v OsDiskImageResponse) string { return v.SourceBlobSasUri }).(pulumi.StringOutput)
}

// Registration information.
type RegistrationType struct {
	// The entity tag used for optimistic concurrency when modifying the resource.
	Etag *string `pulumi:"etag"`
	// Location of the resource.
	Location string `pulumi:"location"`
	// Name of the resource.
	Name string `pulumi:"name"`
	// Registration resource.
	Properties RegistrationPropertiesResponse `pulumi:"properties"`
	// Custom tags for the resource.
	Tags map[string]string `pulumi:"tags"`
	// Type of Resource.
	Type string `pulumi:"type"`
}

// RegistrationTypeInput is an input type that accepts RegistrationTypeArgs and RegistrationTypeOutput values.
// You can construct a concrete instance of `RegistrationTypeInput` via:
//
//          RegistrationTypeArgs{...}
type RegistrationTypeInput interface {
	pulumi.Input

	ToRegistrationTypeOutput() RegistrationTypeOutput
	ToRegistrationTypeOutputWithContext(context.Context) RegistrationTypeOutput
}

// Registration information.
type RegistrationTypeArgs struct {
	// The entity tag used for optimistic concurrency when modifying the resource.
	Etag pulumi.StringPtrInput `pulumi:"etag"`
	// Location of the resource.
	Location pulumi.StringInput `pulumi:"location"`
	// Name of the resource.
	Name pulumi.StringInput `pulumi:"name"`
	// Registration resource.
	Properties RegistrationPropertiesResponseInput `pulumi:"properties"`
	// Custom tags for the resource.
	Tags pulumi.StringMapInput `pulumi:"tags"`
	// Type of Resource.
	Type pulumi.StringInput `pulumi:"type"`
}

func (RegistrationTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RegistrationType)(nil)).Elem()
}

func (i RegistrationTypeArgs) ToRegistrationTypeOutput() RegistrationTypeOutput {
	return i.ToRegistrationTypeOutputWithContext(context.Background())
}

func (i RegistrationTypeArgs) ToRegistrationTypeOutputWithContext(ctx context.Context) RegistrationTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistrationTypeOutput)
}

// Registration information.
type RegistrationTypeOutput struct{ *pulumi.OutputState }

func (RegistrationTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RegistrationType)(nil)).Elem()
}

func (o RegistrationTypeOutput) ToRegistrationTypeOutput() RegistrationTypeOutput {
	return o
}

func (o RegistrationTypeOutput) ToRegistrationTypeOutputWithContext(ctx context.Context) RegistrationTypeOutput {
	return o
}

// The entity tag used for optimistic concurrency when modifying the resource.
func (o RegistrationTypeOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RegistrationType) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

// Location of the resource.
func (o RegistrationTypeOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v RegistrationType) string { return v.Location }).(pulumi.StringOutput)
}

// Name of the resource.
func (o RegistrationTypeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v RegistrationType) string { return v.Name }).(pulumi.StringOutput)
}

// Registration resource.
func (o RegistrationTypeOutput) Properties() RegistrationPropertiesResponseOutput {
	return o.ApplyT(func(v RegistrationType) RegistrationPropertiesResponse { return v.Properties }).(RegistrationPropertiesResponseOutput)
}

// Custom tags for the resource.
func (o RegistrationTypeOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v RegistrationType) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Type of Resource.
func (o RegistrationTypeOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v RegistrationType) string { return v.Type }).(pulumi.StringOutput)
}

// Properties portion of the registration resource.
type RegistrationPropertiesResponse struct {
	// Specifies the billing mode for the Azure Stack registration.
	BillingModel *string `pulumi:"billingModel"`
	// The identifier of the registered Azure Stack.
	CloudId *string `pulumi:"cloudId"`
	// The object identifier associated with the Azure Stack connecting to Azure.
	ObjectId *string `pulumi:"objectId"`
}

// RegistrationPropertiesResponseInput is an input type that accepts RegistrationPropertiesResponseArgs and RegistrationPropertiesResponseOutput values.
// You can construct a concrete instance of `RegistrationPropertiesResponseInput` via:
//
//          RegistrationPropertiesResponseArgs{...}
type RegistrationPropertiesResponseInput interface {
	pulumi.Input

	ToRegistrationPropertiesResponseOutput() RegistrationPropertiesResponseOutput
	ToRegistrationPropertiesResponseOutputWithContext(context.Context) RegistrationPropertiesResponseOutput
}

// Properties portion of the registration resource.
type RegistrationPropertiesResponseArgs struct {
	// Specifies the billing mode for the Azure Stack registration.
	BillingModel pulumi.StringPtrInput `pulumi:"billingModel"`
	// The identifier of the registered Azure Stack.
	CloudId pulumi.StringPtrInput `pulumi:"cloudId"`
	// The object identifier associated with the Azure Stack connecting to Azure.
	ObjectId pulumi.StringPtrInput `pulumi:"objectId"`
}

func (RegistrationPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RegistrationPropertiesResponse)(nil)).Elem()
}

func (i RegistrationPropertiesResponseArgs) ToRegistrationPropertiesResponseOutput() RegistrationPropertiesResponseOutput {
	return i.ToRegistrationPropertiesResponseOutputWithContext(context.Background())
}

func (i RegistrationPropertiesResponseArgs) ToRegistrationPropertiesResponseOutputWithContext(ctx context.Context) RegistrationPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistrationPropertiesResponseOutput)
}

func (i RegistrationPropertiesResponseArgs) ToRegistrationPropertiesResponsePtrOutput() RegistrationPropertiesResponsePtrOutput {
	return i.ToRegistrationPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i RegistrationPropertiesResponseArgs) ToRegistrationPropertiesResponsePtrOutputWithContext(ctx context.Context) RegistrationPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistrationPropertiesResponseOutput).ToRegistrationPropertiesResponsePtrOutputWithContext(ctx)
}

// RegistrationPropertiesResponsePtrInput is an input type that accepts RegistrationPropertiesResponseArgs, RegistrationPropertiesResponsePtr and RegistrationPropertiesResponsePtrOutput values.
// You can construct a concrete instance of `RegistrationPropertiesResponsePtrInput` via:
//
//          RegistrationPropertiesResponseArgs{...}
//
//  or:
//
//          nil
type RegistrationPropertiesResponsePtrInput interface {
	pulumi.Input

	ToRegistrationPropertiesResponsePtrOutput() RegistrationPropertiesResponsePtrOutput
	ToRegistrationPropertiesResponsePtrOutputWithContext(context.Context) RegistrationPropertiesResponsePtrOutput
}

type registrationPropertiesResponsePtrType RegistrationPropertiesResponseArgs

func RegistrationPropertiesResponsePtr(v *RegistrationPropertiesResponseArgs) RegistrationPropertiesResponsePtrInput {
	return (*registrationPropertiesResponsePtrType)(v)
}

func (*registrationPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RegistrationPropertiesResponse)(nil)).Elem()
}

func (i *registrationPropertiesResponsePtrType) ToRegistrationPropertiesResponsePtrOutput() RegistrationPropertiesResponsePtrOutput {
	return i.ToRegistrationPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *registrationPropertiesResponsePtrType) ToRegistrationPropertiesResponsePtrOutputWithContext(ctx context.Context) RegistrationPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistrationPropertiesResponsePtrOutput)
}

// Properties portion of the registration resource.
type RegistrationPropertiesResponseOutput struct{ *pulumi.OutputState }

func (RegistrationPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RegistrationPropertiesResponse)(nil)).Elem()
}

func (o RegistrationPropertiesResponseOutput) ToRegistrationPropertiesResponseOutput() RegistrationPropertiesResponseOutput {
	return o
}

func (o RegistrationPropertiesResponseOutput) ToRegistrationPropertiesResponseOutputWithContext(ctx context.Context) RegistrationPropertiesResponseOutput {
	return o
}

func (o RegistrationPropertiesResponseOutput) ToRegistrationPropertiesResponsePtrOutput() RegistrationPropertiesResponsePtrOutput {
	return o.ToRegistrationPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o RegistrationPropertiesResponseOutput) ToRegistrationPropertiesResponsePtrOutputWithContext(ctx context.Context) RegistrationPropertiesResponsePtrOutput {
	return o.ApplyT(func(v RegistrationPropertiesResponse) *RegistrationPropertiesResponse {
		return &v
	}).(RegistrationPropertiesResponsePtrOutput)
}

// Specifies the billing mode for the Azure Stack registration.
func (o RegistrationPropertiesResponseOutput) BillingModel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RegistrationPropertiesResponse) *string { return v.BillingModel }).(pulumi.StringPtrOutput)
}

// The identifier of the registered Azure Stack.
func (o RegistrationPropertiesResponseOutput) CloudId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RegistrationPropertiesResponse) *string { return v.CloudId }).(pulumi.StringPtrOutput)
}

// The object identifier associated with the Azure Stack connecting to Azure.
func (o RegistrationPropertiesResponseOutput) ObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RegistrationPropertiesResponse) *string { return v.ObjectId }).(pulumi.StringPtrOutput)
}

type RegistrationPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (RegistrationPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RegistrationPropertiesResponse)(nil)).Elem()
}

func (o RegistrationPropertiesResponsePtrOutput) ToRegistrationPropertiesResponsePtrOutput() RegistrationPropertiesResponsePtrOutput {
	return o
}

func (o RegistrationPropertiesResponsePtrOutput) ToRegistrationPropertiesResponsePtrOutputWithContext(ctx context.Context) RegistrationPropertiesResponsePtrOutput {
	return o
}

func (o RegistrationPropertiesResponsePtrOutput) Elem() RegistrationPropertiesResponseOutput {
	return o.ApplyT(func(v *RegistrationPropertiesResponse) RegistrationPropertiesResponse { return *v }).(RegistrationPropertiesResponseOutput)
}

// Specifies the billing mode for the Azure Stack registration.
func (o RegistrationPropertiesResponsePtrOutput) BillingModel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RegistrationPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.BillingModel
	}).(pulumi.StringPtrOutput)
}

// The identifier of the registered Azure Stack.
func (o RegistrationPropertiesResponsePtrOutput) CloudId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RegistrationPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.CloudId
	}).(pulumi.StringPtrOutput)
}

// The object identifier associated with the Azure Stack connecting to Azure.
func (o RegistrationPropertiesResponsePtrOutput) ObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RegistrationPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ObjectId
	}).(pulumi.StringPtrOutput)
}

// The URI.
type UriResponse struct {
	// The URI.
	Uri string `pulumi:"uri"`
}

// UriResponseInput is an input type that accepts UriResponseArgs and UriResponseOutput values.
// You can construct a concrete instance of `UriResponseInput` via:
//
//          UriResponseArgs{...}
type UriResponseInput interface {
	pulumi.Input

	ToUriResponseOutput() UriResponseOutput
	ToUriResponseOutputWithContext(context.Context) UriResponseOutput
}

// The URI.
type UriResponseArgs struct {
	// The URI.
	Uri pulumi.StringInput `pulumi:"uri"`
}

func (UriResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UriResponse)(nil)).Elem()
}

func (i UriResponseArgs) ToUriResponseOutput() UriResponseOutput {
	return i.ToUriResponseOutputWithContext(context.Background())
}

func (i UriResponseArgs) ToUriResponseOutputWithContext(ctx context.Context) UriResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UriResponseOutput)
}

// The URI.
type UriResponseOutput struct{ *pulumi.OutputState }

func (UriResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UriResponse)(nil)).Elem()
}

func (o UriResponseOutput) ToUriResponseOutput() UriResponseOutput {
	return o
}

func (o UriResponseOutput) ToUriResponseOutputWithContext(ctx context.Context) UriResponseOutput {
	return o
}

// The URI.
func (o UriResponseOutput) Uri() pulumi.StringOutput {
	return o.ApplyT(func(v UriResponse) string { return v.Uri }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(CustomerSubscriptionTypeOutput{})
	pulumi.RegisterOutputType(CustomerSubscriptionPropertiesResponseOutput{})
	pulumi.RegisterOutputType(CustomerSubscriptionPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(DataDiskImageResponseOutput{})
	pulumi.RegisterOutputType(DataDiskImageResponseArrayOutput{})
	pulumi.RegisterOutputType(ExtendedProductPropertiesResponseOutput{})
	pulumi.RegisterOutputType(OsDiskImageResponseOutput{})
	pulumi.RegisterOutputType(RegistrationTypeOutput{})
	pulumi.RegisterOutputType(RegistrationPropertiesResponseOutput{})
	pulumi.RegisterOutputType(RegistrationPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(UriResponseOutput{})
}
