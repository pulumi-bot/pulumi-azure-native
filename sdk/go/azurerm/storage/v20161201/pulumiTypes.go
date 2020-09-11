// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20161201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The custom domain assigned to this storage account. This can be set via Update.
type CustomDomain struct {
	// Gets or sets the custom domain name assigned to the storage account. Name is the CNAME source.
	Name string `pulumi:"name"`
	// Indicates whether indirect CName validation is enabled. Default value is false. This should only be set on updates.
	UseSubDomainName *bool `pulumi:"useSubDomainName"`
}

// CustomDomainInput is an input type that accepts CustomDomainArgs and CustomDomainOutput values.
// You can construct a concrete instance of `CustomDomainInput` via:
//
//          CustomDomainArgs{...}
type CustomDomainInput interface {
	pulumi.Input

	ToCustomDomainOutput() CustomDomainOutput
	ToCustomDomainOutputWithContext(context.Context) CustomDomainOutput
}

// The custom domain assigned to this storage account. This can be set via Update.
type CustomDomainArgs struct {
	// Gets or sets the custom domain name assigned to the storage account. Name is the CNAME source.
	Name pulumi.StringInput `pulumi:"name"`
	// Indicates whether indirect CName validation is enabled. Default value is false. This should only be set on updates.
	UseSubDomainName pulumi.BoolPtrInput `pulumi:"useSubDomainName"`
}

func (CustomDomainArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomDomain)(nil)).Elem()
}

func (i CustomDomainArgs) ToCustomDomainOutput() CustomDomainOutput {
	return i.ToCustomDomainOutputWithContext(context.Background())
}

func (i CustomDomainArgs) ToCustomDomainOutputWithContext(ctx context.Context) CustomDomainOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomDomainOutput)
}

func (i CustomDomainArgs) ToCustomDomainPtrOutput() CustomDomainPtrOutput {
	return i.ToCustomDomainPtrOutputWithContext(context.Background())
}

func (i CustomDomainArgs) ToCustomDomainPtrOutputWithContext(ctx context.Context) CustomDomainPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomDomainOutput).ToCustomDomainPtrOutputWithContext(ctx)
}

// CustomDomainPtrInput is an input type that accepts CustomDomainArgs, CustomDomainPtr and CustomDomainPtrOutput values.
// You can construct a concrete instance of `CustomDomainPtrInput` via:
//
//          CustomDomainArgs{...}
//
//  or:
//
//          nil
type CustomDomainPtrInput interface {
	pulumi.Input

	ToCustomDomainPtrOutput() CustomDomainPtrOutput
	ToCustomDomainPtrOutputWithContext(context.Context) CustomDomainPtrOutput
}

type customDomainPtrType CustomDomainArgs

func CustomDomainPtr(v *CustomDomainArgs) CustomDomainPtrInput {
	return (*customDomainPtrType)(v)
}

func (*customDomainPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomDomain)(nil)).Elem()
}

func (i *customDomainPtrType) ToCustomDomainPtrOutput() CustomDomainPtrOutput {
	return i.ToCustomDomainPtrOutputWithContext(context.Background())
}

func (i *customDomainPtrType) ToCustomDomainPtrOutputWithContext(ctx context.Context) CustomDomainPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomDomainPtrOutput)
}

// The custom domain assigned to this storage account. This can be set via Update.
type CustomDomainResponse struct {
	// Gets or sets the custom domain name assigned to the storage account. Name is the CNAME source.
	Name string `pulumi:"name"`
	// Indicates whether indirect CName validation is enabled. Default value is false. This should only be set on updates.
	UseSubDomainName *bool `pulumi:"useSubDomainName"`
}

// The custom domain assigned to this storage account. This can be set via Update.
type CustomDomainResponseOutput struct{ *pulumi.OutputState }

func (CustomDomainResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomDomainResponse)(nil)).Elem()
}

func (o CustomDomainResponseOutput) ToCustomDomainResponseOutput() CustomDomainResponseOutput {
	return o
}

func (o CustomDomainResponseOutput) ToCustomDomainResponseOutputWithContext(ctx context.Context) CustomDomainResponseOutput {
	return o
}

func (o CustomDomainResponseOutput) ToCustomDomainResponsePtrOutput() CustomDomainResponsePtrOutput {
	return o.ToCustomDomainResponsePtrOutputWithContext(context.Background())
}

func (o CustomDomainResponseOutput) ToCustomDomainResponsePtrOutputWithContext(ctx context.Context) CustomDomainResponsePtrOutput {
	return o.ApplyT(func(v CustomDomainResponse) *CustomDomainResponse {
		return &v
	}).(CustomDomainResponsePtrOutput)
}

// Gets or sets the custom domain name assigned to the storage account. Name is the CNAME source.
func (o CustomDomainResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v CustomDomainResponse) string { return v.Name }).(pulumi.StringOutput)
}

// Indicates whether indirect CName validation is enabled. Default value is false. This should only be set on updates.
func (o CustomDomainResponseOutput) UseSubDomainName() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CustomDomainResponse) *bool { return v.UseSubDomainName }).(pulumi.BoolPtrOutput)
}

type CustomDomainResponsePtrOutput struct{ *pulumi.OutputState }

func (CustomDomainResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomDomainResponse)(nil)).Elem()
}

func (o CustomDomainResponsePtrOutput) ToCustomDomainResponsePtrOutput() CustomDomainResponsePtrOutput {
	return o
}

func (o CustomDomainResponsePtrOutput) ToCustomDomainResponsePtrOutputWithContext(ctx context.Context) CustomDomainResponsePtrOutput {
	return o
}

func (o CustomDomainResponsePtrOutput) Elem() CustomDomainResponseOutput {
	return o.ApplyT(func(v *CustomDomainResponse) CustomDomainResponse { return *v }).(CustomDomainResponseOutput)
}

// Gets or sets the custom domain name assigned to the storage account. Name is the CNAME source.
func (o CustomDomainResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomDomainResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

// Indicates whether indirect CName validation is enabled. Default value is false. This should only be set on updates.
func (o CustomDomainResponsePtrOutput) UseSubDomainName() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CustomDomainResponse) *bool {
		if v == nil {
			return nil
		}
		return v.UseSubDomainName
	}).(pulumi.BoolPtrOutput)
}

// The encryption settings on the storage account.
type Encryption struct {
	// The encryption keySource (provider). Possible values (case-insensitive):  Microsoft.Storage
	KeySource string `pulumi:"keySource"`
	// List of services which support encryption.
	Services *EncryptionServices `pulumi:"services"`
}

// EncryptionInput is an input type that accepts EncryptionArgs and EncryptionOutput values.
// You can construct a concrete instance of `EncryptionInput` via:
//
//          EncryptionArgs{...}
type EncryptionInput interface {
	pulumi.Input

	ToEncryptionOutput() EncryptionOutput
	ToEncryptionOutputWithContext(context.Context) EncryptionOutput
}

// The encryption settings on the storage account.
type EncryptionArgs struct {
	// The encryption keySource (provider). Possible values (case-insensitive):  Microsoft.Storage
	KeySource pulumi.StringInput `pulumi:"keySource"`
	// List of services which support encryption.
	Services EncryptionServicesPtrInput `pulumi:"services"`
}

func (EncryptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Encryption)(nil)).Elem()
}

func (i EncryptionArgs) ToEncryptionOutput() EncryptionOutput {
	return i.ToEncryptionOutputWithContext(context.Background())
}

func (i EncryptionArgs) ToEncryptionOutputWithContext(ctx context.Context) EncryptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionOutput)
}

func (i EncryptionArgs) ToEncryptionPtrOutput() EncryptionPtrOutput {
	return i.ToEncryptionPtrOutputWithContext(context.Background())
}

func (i EncryptionArgs) ToEncryptionPtrOutputWithContext(ctx context.Context) EncryptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionOutput).ToEncryptionPtrOutputWithContext(ctx)
}

// EncryptionPtrInput is an input type that accepts EncryptionArgs, EncryptionPtr and EncryptionPtrOutput values.
// You can construct a concrete instance of `EncryptionPtrInput` via:
//
//          EncryptionArgs{...}
//
//  or:
//
//          nil
type EncryptionPtrInput interface {
	pulumi.Input

	ToEncryptionPtrOutput() EncryptionPtrOutput
	ToEncryptionPtrOutputWithContext(context.Context) EncryptionPtrOutput
}

type encryptionPtrType EncryptionArgs

func EncryptionPtr(v *EncryptionArgs) EncryptionPtrInput {
	return (*encryptionPtrType)(v)
}

func (*encryptionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Encryption)(nil)).Elem()
}

func (i *encryptionPtrType) ToEncryptionPtrOutput() EncryptionPtrOutput {
	return i.ToEncryptionPtrOutputWithContext(context.Background())
}

func (i *encryptionPtrType) ToEncryptionPtrOutputWithContext(ctx context.Context) EncryptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionPtrOutput)
}

// The encryption settings on the storage account.
type EncryptionResponse struct {
	// The encryption keySource (provider). Possible values (case-insensitive):  Microsoft.Storage
	KeySource string `pulumi:"keySource"`
	// List of services which support encryption.
	Services *EncryptionServicesResponse `pulumi:"services"`
}

// The encryption settings on the storage account.
type EncryptionResponseOutput struct{ *pulumi.OutputState }

func (EncryptionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionResponse)(nil)).Elem()
}

func (o EncryptionResponseOutput) ToEncryptionResponseOutput() EncryptionResponseOutput {
	return o
}

func (o EncryptionResponseOutput) ToEncryptionResponseOutputWithContext(ctx context.Context) EncryptionResponseOutput {
	return o
}

func (o EncryptionResponseOutput) ToEncryptionResponsePtrOutput() EncryptionResponsePtrOutput {
	return o.ToEncryptionResponsePtrOutputWithContext(context.Background())
}

func (o EncryptionResponseOutput) ToEncryptionResponsePtrOutputWithContext(ctx context.Context) EncryptionResponsePtrOutput {
	return o.ApplyT(func(v EncryptionResponse) *EncryptionResponse {
		return &v
	}).(EncryptionResponsePtrOutput)
}

// The encryption keySource (provider). Possible values (case-insensitive):  Microsoft.Storage
func (o EncryptionResponseOutput) KeySource() pulumi.StringOutput {
	return o.ApplyT(func(v EncryptionResponse) string { return v.KeySource }).(pulumi.StringOutput)
}

// List of services which support encryption.
func (o EncryptionResponseOutput) Services() EncryptionServicesResponsePtrOutput {
	return o.ApplyT(func(v EncryptionResponse) *EncryptionServicesResponse { return v.Services }).(EncryptionServicesResponsePtrOutput)
}

type EncryptionResponsePtrOutput struct{ *pulumi.OutputState }

func (EncryptionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionResponse)(nil)).Elem()
}

func (o EncryptionResponsePtrOutput) ToEncryptionResponsePtrOutput() EncryptionResponsePtrOutput {
	return o
}

func (o EncryptionResponsePtrOutput) ToEncryptionResponsePtrOutputWithContext(ctx context.Context) EncryptionResponsePtrOutput {
	return o
}

func (o EncryptionResponsePtrOutput) Elem() EncryptionResponseOutput {
	return o.ApplyT(func(v *EncryptionResponse) EncryptionResponse { return *v }).(EncryptionResponseOutput)
}

// The encryption keySource (provider). Possible values (case-insensitive):  Microsoft.Storage
func (o EncryptionResponsePtrOutput) KeySource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionResponse) *string {
		if v == nil {
			return nil
		}
		return &v.KeySource
	}).(pulumi.StringPtrOutput)
}

// List of services which support encryption.
func (o EncryptionResponsePtrOutput) Services() EncryptionServicesResponsePtrOutput {
	return o.ApplyT(func(v *EncryptionResponse) *EncryptionServicesResponse {
		if v == nil {
			return nil
		}
		return v.Services
	}).(EncryptionServicesResponsePtrOutput)
}

// A service that allows server-side encryption to be used.
type EncryptionService struct {
	// A boolean indicating whether or not the service encrypts the data as it is stored.
	Enabled *bool `pulumi:"enabled"`
}

// EncryptionServiceInput is an input type that accepts EncryptionServiceArgs and EncryptionServiceOutput values.
// You can construct a concrete instance of `EncryptionServiceInput` via:
//
//          EncryptionServiceArgs{...}
type EncryptionServiceInput interface {
	pulumi.Input

	ToEncryptionServiceOutput() EncryptionServiceOutput
	ToEncryptionServiceOutputWithContext(context.Context) EncryptionServiceOutput
}

// A service that allows server-side encryption to be used.
type EncryptionServiceArgs struct {
	// A boolean indicating whether or not the service encrypts the data as it is stored.
	Enabled pulumi.BoolPtrInput `pulumi:"enabled"`
}

func (EncryptionServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionService)(nil)).Elem()
}

func (i EncryptionServiceArgs) ToEncryptionServiceOutput() EncryptionServiceOutput {
	return i.ToEncryptionServiceOutputWithContext(context.Background())
}

func (i EncryptionServiceArgs) ToEncryptionServiceOutputWithContext(ctx context.Context) EncryptionServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionServiceOutput)
}

func (i EncryptionServiceArgs) ToEncryptionServicePtrOutput() EncryptionServicePtrOutput {
	return i.ToEncryptionServicePtrOutputWithContext(context.Background())
}

func (i EncryptionServiceArgs) ToEncryptionServicePtrOutputWithContext(ctx context.Context) EncryptionServicePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionServiceOutput).ToEncryptionServicePtrOutputWithContext(ctx)
}

// EncryptionServicePtrInput is an input type that accepts EncryptionServiceArgs, EncryptionServicePtr and EncryptionServicePtrOutput values.
// You can construct a concrete instance of `EncryptionServicePtrInput` via:
//
//          EncryptionServiceArgs{...}
//
//  or:
//
//          nil
type EncryptionServicePtrInput interface {
	pulumi.Input

	ToEncryptionServicePtrOutput() EncryptionServicePtrOutput
	ToEncryptionServicePtrOutputWithContext(context.Context) EncryptionServicePtrOutput
}

type encryptionServicePtrType EncryptionServiceArgs

func EncryptionServicePtr(v *EncryptionServiceArgs) EncryptionServicePtrInput {
	return (*encryptionServicePtrType)(v)
}

func (*encryptionServicePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionService)(nil)).Elem()
}

func (i *encryptionServicePtrType) ToEncryptionServicePtrOutput() EncryptionServicePtrOutput {
	return i.ToEncryptionServicePtrOutputWithContext(context.Background())
}

func (i *encryptionServicePtrType) ToEncryptionServicePtrOutputWithContext(ctx context.Context) EncryptionServicePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionServicePtrOutput)
}

// A service that allows server-side encryption to be used.
type EncryptionServiceResponse struct {
	// A boolean indicating whether or not the service encrypts the data as it is stored.
	Enabled *bool `pulumi:"enabled"`
	// Gets a rough estimate of the date/time when the encryption was last enabled by the user. Only returned when encryption is enabled. There might be some unencrypted blobs which were written after this time, as it is just a rough estimate.
	LastEnabledTime string `pulumi:"lastEnabledTime"`
}

// A service that allows server-side encryption to be used.
type EncryptionServiceResponseOutput struct{ *pulumi.OutputState }

func (EncryptionServiceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionServiceResponse)(nil)).Elem()
}

func (o EncryptionServiceResponseOutput) ToEncryptionServiceResponseOutput() EncryptionServiceResponseOutput {
	return o
}

func (o EncryptionServiceResponseOutput) ToEncryptionServiceResponseOutputWithContext(ctx context.Context) EncryptionServiceResponseOutput {
	return o
}

func (o EncryptionServiceResponseOutput) ToEncryptionServiceResponsePtrOutput() EncryptionServiceResponsePtrOutput {
	return o.ToEncryptionServiceResponsePtrOutputWithContext(context.Background())
}

func (o EncryptionServiceResponseOutput) ToEncryptionServiceResponsePtrOutputWithContext(ctx context.Context) EncryptionServiceResponsePtrOutput {
	return o.ApplyT(func(v EncryptionServiceResponse) *EncryptionServiceResponse {
		return &v
	}).(EncryptionServiceResponsePtrOutput)
}

// A boolean indicating whether or not the service encrypts the data as it is stored.
func (o EncryptionServiceResponseOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EncryptionServiceResponse) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// Gets a rough estimate of the date/time when the encryption was last enabled by the user. Only returned when encryption is enabled. There might be some unencrypted blobs which were written after this time, as it is just a rough estimate.
func (o EncryptionServiceResponseOutput) LastEnabledTime() pulumi.StringOutput {
	return o.ApplyT(func(v EncryptionServiceResponse) string { return v.LastEnabledTime }).(pulumi.StringOutput)
}

type EncryptionServiceResponsePtrOutput struct{ *pulumi.OutputState }

func (EncryptionServiceResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionServiceResponse)(nil)).Elem()
}

func (o EncryptionServiceResponsePtrOutput) ToEncryptionServiceResponsePtrOutput() EncryptionServiceResponsePtrOutput {
	return o
}

func (o EncryptionServiceResponsePtrOutput) ToEncryptionServiceResponsePtrOutputWithContext(ctx context.Context) EncryptionServiceResponsePtrOutput {
	return o
}

func (o EncryptionServiceResponsePtrOutput) Elem() EncryptionServiceResponseOutput {
	return o.ApplyT(func(v *EncryptionServiceResponse) EncryptionServiceResponse { return *v }).(EncryptionServiceResponseOutput)
}

// A boolean indicating whether or not the service encrypts the data as it is stored.
func (o EncryptionServiceResponsePtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EncryptionServiceResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

// Gets a rough estimate of the date/time when the encryption was last enabled by the user. Only returned when encryption is enabled. There might be some unencrypted blobs which were written after this time, as it is just a rough estimate.
func (o EncryptionServiceResponsePtrOutput) LastEnabledTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionServiceResponse) *string {
		if v == nil {
			return nil
		}
		return &v.LastEnabledTime
	}).(pulumi.StringPtrOutput)
}

// A list of services that support encryption.
type EncryptionServices struct {
	// The encryption function of the blob storage service.
	Blob *EncryptionService `pulumi:"blob"`
	// The encryption function of the file storage service.
	File *EncryptionService `pulumi:"file"`
}

// EncryptionServicesInput is an input type that accepts EncryptionServicesArgs and EncryptionServicesOutput values.
// You can construct a concrete instance of `EncryptionServicesInput` via:
//
//          EncryptionServicesArgs{...}
type EncryptionServicesInput interface {
	pulumi.Input

	ToEncryptionServicesOutput() EncryptionServicesOutput
	ToEncryptionServicesOutputWithContext(context.Context) EncryptionServicesOutput
}

// A list of services that support encryption.
type EncryptionServicesArgs struct {
	// The encryption function of the blob storage service.
	Blob EncryptionServicePtrInput `pulumi:"blob"`
	// The encryption function of the file storage service.
	File EncryptionServicePtrInput `pulumi:"file"`
}

func (EncryptionServicesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionServices)(nil)).Elem()
}

func (i EncryptionServicesArgs) ToEncryptionServicesOutput() EncryptionServicesOutput {
	return i.ToEncryptionServicesOutputWithContext(context.Background())
}

func (i EncryptionServicesArgs) ToEncryptionServicesOutputWithContext(ctx context.Context) EncryptionServicesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionServicesOutput)
}

func (i EncryptionServicesArgs) ToEncryptionServicesPtrOutput() EncryptionServicesPtrOutput {
	return i.ToEncryptionServicesPtrOutputWithContext(context.Background())
}

func (i EncryptionServicesArgs) ToEncryptionServicesPtrOutputWithContext(ctx context.Context) EncryptionServicesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionServicesOutput).ToEncryptionServicesPtrOutputWithContext(ctx)
}

// EncryptionServicesPtrInput is an input type that accepts EncryptionServicesArgs, EncryptionServicesPtr and EncryptionServicesPtrOutput values.
// You can construct a concrete instance of `EncryptionServicesPtrInput` via:
//
//          EncryptionServicesArgs{...}
//
//  or:
//
//          nil
type EncryptionServicesPtrInput interface {
	pulumi.Input

	ToEncryptionServicesPtrOutput() EncryptionServicesPtrOutput
	ToEncryptionServicesPtrOutputWithContext(context.Context) EncryptionServicesPtrOutput
}

type encryptionServicesPtrType EncryptionServicesArgs

func EncryptionServicesPtr(v *EncryptionServicesArgs) EncryptionServicesPtrInput {
	return (*encryptionServicesPtrType)(v)
}

func (*encryptionServicesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionServices)(nil)).Elem()
}

func (i *encryptionServicesPtrType) ToEncryptionServicesPtrOutput() EncryptionServicesPtrOutput {
	return i.ToEncryptionServicesPtrOutputWithContext(context.Background())
}

func (i *encryptionServicesPtrType) ToEncryptionServicesPtrOutputWithContext(ctx context.Context) EncryptionServicesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionServicesPtrOutput)
}

// A list of services that support encryption.
type EncryptionServicesResponse struct {
	// The encryption function of the blob storage service.
	Blob *EncryptionServiceResponse `pulumi:"blob"`
	// The encryption function of the file storage service.
	File *EncryptionServiceResponse `pulumi:"file"`
	// The encryption function of the queue storage service.
	Queue EncryptionServiceResponse `pulumi:"queue"`
	// The encryption function of the table storage service.
	Table EncryptionServiceResponse `pulumi:"table"`
}

// A list of services that support encryption.
type EncryptionServicesResponseOutput struct{ *pulumi.OutputState }

func (EncryptionServicesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionServicesResponse)(nil)).Elem()
}

func (o EncryptionServicesResponseOutput) ToEncryptionServicesResponseOutput() EncryptionServicesResponseOutput {
	return o
}

func (o EncryptionServicesResponseOutput) ToEncryptionServicesResponseOutputWithContext(ctx context.Context) EncryptionServicesResponseOutput {
	return o
}

func (o EncryptionServicesResponseOutput) ToEncryptionServicesResponsePtrOutput() EncryptionServicesResponsePtrOutput {
	return o.ToEncryptionServicesResponsePtrOutputWithContext(context.Background())
}

func (o EncryptionServicesResponseOutput) ToEncryptionServicesResponsePtrOutputWithContext(ctx context.Context) EncryptionServicesResponsePtrOutput {
	return o.ApplyT(func(v EncryptionServicesResponse) *EncryptionServicesResponse {
		return &v
	}).(EncryptionServicesResponsePtrOutput)
}

// The encryption function of the blob storage service.
func (o EncryptionServicesResponseOutput) Blob() EncryptionServiceResponsePtrOutput {
	return o.ApplyT(func(v EncryptionServicesResponse) *EncryptionServiceResponse { return v.Blob }).(EncryptionServiceResponsePtrOutput)
}

// The encryption function of the file storage service.
func (o EncryptionServicesResponseOutput) File() EncryptionServiceResponsePtrOutput {
	return o.ApplyT(func(v EncryptionServicesResponse) *EncryptionServiceResponse { return v.File }).(EncryptionServiceResponsePtrOutput)
}

// The encryption function of the queue storage service.
func (o EncryptionServicesResponseOutput) Queue() EncryptionServiceResponseOutput {
	return o.ApplyT(func(v EncryptionServicesResponse) EncryptionServiceResponse { return v.Queue }).(EncryptionServiceResponseOutput)
}

// The encryption function of the table storage service.
func (o EncryptionServicesResponseOutput) Table() EncryptionServiceResponseOutput {
	return o.ApplyT(func(v EncryptionServicesResponse) EncryptionServiceResponse { return v.Table }).(EncryptionServiceResponseOutput)
}

type EncryptionServicesResponsePtrOutput struct{ *pulumi.OutputState }

func (EncryptionServicesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionServicesResponse)(nil)).Elem()
}

func (o EncryptionServicesResponsePtrOutput) ToEncryptionServicesResponsePtrOutput() EncryptionServicesResponsePtrOutput {
	return o
}

func (o EncryptionServicesResponsePtrOutput) ToEncryptionServicesResponsePtrOutputWithContext(ctx context.Context) EncryptionServicesResponsePtrOutput {
	return o
}

func (o EncryptionServicesResponsePtrOutput) Elem() EncryptionServicesResponseOutput {
	return o.ApplyT(func(v *EncryptionServicesResponse) EncryptionServicesResponse { return *v }).(EncryptionServicesResponseOutput)
}

// The encryption function of the blob storage service.
func (o EncryptionServicesResponsePtrOutput) Blob() EncryptionServiceResponsePtrOutput {
	return o.ApplyT(func(v *EncryptionServicesResponse) *EncryptionServiceResponse {
		if v == nil {
			return nil
		}
		return v.Blob
	}).(EncryptionServiceResponsePtrOutput)
}

// The encryption function of the file storage service.
func (o EncryptionServicesResponsePtrOutput) File() EncryptionServiceResponsePtrOutput {
	return o.ApplyT(func(v *EncryptionServicesResponse) *EncryptionServiceResponse {
		if v == nil {
			return nil
		}
		return v.File
	}).(EncryptionServiceResponsePtrOutput)
}

// The encryption function of the queue storage service.
func (o EncryptionServicesResponsePtrOutput) Queue() EncryptionServiceResponsePtrOutput {
	return o.ApplyT(func(v *EncryptionServicesResponse) *EncryptionServiceResponse {
		if v == nil {
			return nil
		}
		return &v.Queue
	}).(EncryptionServiceResponsePtrOutput)
}

// The encryption function of the table storage service.
func (o EncryptionServicesResponsePtrOutput) Table() EncryptionServiceResponsePtrOutput {
	return o.ApplyT(func(v *EncryptionServicesResponse) *EncryptionServiceResponse {
		if v == nil {
			return nil
		}
		return &v.Table
	}).(EncryptionServiceResponsePtrOutput)
}

// The URIs that are used to perform a retrieval of a public blob, queue, or table object.
type EndpointsResponse struct {
	// Gets the blob endpoint.
	Blob string `pulumi:"blob"`
	// Gets the file endpoint.
	File string `pulumi:"file"`
	// Gets the queue endpoint.
	Queue string `pulumi:"queue"`
	// Gets the table endpoint.
	Table string `pulumi:"table"`
}

// The URIs that are used to perform a retrieval of a public blob, queue, or table object.
type EndpointsResponseOutput struct{ *pulumi.OutputState }

func (EndpointsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointsResponse)(nil)).Elem()
}

func (o EndpointsResponseOutput) ToEndpointsResponseOutput() EndpointsResponseOutput {
	return o
}

func (o EndpointsResponseOutput) ToEndpointsResponseOutputWithContext(ctx context.Context) EndpointsResponseOutput {
	return o
}

func (o EndpointsResponseOutput) ToEndpointsResponsePtrOutput() EndpointsResponsePtrOutput {
	return o.ToEndpointsResponsePtrOutputWithContext(context.Background())
}

func (o EndpointsResponseOutput) ToEndpointsResponsePtrOutputWithContext(ctx context.Context) EndpointsResponsePtrOutput {
	return o.ApplyT(func(v EndpointsResponse) *EndpointsResponse {
		return &v
	}).(EndpointsResponsePtrOutput)
}

// Gets the blob endpoint.
func (o EndpointsResponseOutput) Blob() pulumi.StringOutput {
	return o.ApplyT(func(v EndpointsResponse) string { return v.Blob }).(pulumi.StringOutput)
}

// Gets the file endpoint.
func (o EndpointsResponseOutput) File() pulumi.StringOutput {
	return o.ApplyT(func(v EndpointsResponse) string { return v.File }).(pulumi.StringOutput)
}

// Gets the queue endpoint.
func (o EndpointsResponseOutput) Queue() pulumi.StringOutput {
	return o.ApplyT(func(v EndpointsResponse) string { return v.Queue }).(pulumi.StringOutput)
}

// Gets the table endpoint.
func (o EndpointsResponseOutput) Table() pulumi.StringOutput {
	return o.ApplyT(func(v EndpointsResponse) string { return v.Table }).(pulumi.StringOutput)
}

type EndpointsResponsePtrOutput struct{ *pulumi.OutputState }

func (EndpointsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointsResponse)(nil)).Elem()
}

func (o EndpointsResponsePtrOutput) ToEndpointsResponsePtrOutput() EndpointsResponsePtrOutput {
	return o
}

func (o EndpointsResponsePtrOutput) ToEndpointsResponsePtrOutputWithContext(ctx context.Context) EndpointsResponsePtrOutput {
	return o
}

func (o EndpointsResponsePtrOutput) Elem() EndpointsResponseOutput {
	return o.ApplyT(func(v *EndpointsResponse) EndpointsResponse { return *v }).(EndpointsResponseOutput)
}

// Gets the blob endpoint.
func (o EndpointsResponsePtrOutput) Blob() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EndpointsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Blob
	}).(pulumi.StringPtrOutput)
}

// Gets the file endpoint.
func (o EndpointsResponsePtrOutput) File() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EndpointsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.File
	}).(pulumi.StringPtrOutput)
}

// Gets the queue endpoint.
func (o EndpointsResponsePtrOutput) Queue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EndpointsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Queue
	}).(pulumi.StringPtrOutput)
}

// Gets the table endpoint.
func (o EndpointsResponsePtrOutput) Table() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EndpointsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Table
	}).(pulumi.StringPtrOutput)
}

// The SKU of the storage account.
type Sku struct {
	// Gets or sets the sku name. Required for account creation; optional for update. Note that in older versions, sku name was called accountType.
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

// The SKU of the storage account.
type SkuArgs struct {
	// Gets or sets the sku name. Required for account creation; optional for update. Note that in older versions, sku name was called accountType.
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

// The SKU of the storage account.
type SkuResponse struct {
	// Gets or sets the sku name. Required for account creation; optional for update. Note that in older versions, sku name was called accountType.
	Name string `pulumi:"name"`
	// Gets the sku tier. This is based on the SKU name.
	Tier string `pulumi:"tier"`
}

// The SKU of the storage account.
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

// Gets or sets the sku name. Required for account creation; optional for update. Note that in older versions, sku name was called accountType.
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

// Gets or sets the sku name. Required for account creation; optional for update. Note that in older versions, sku name was called accountType.
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

// An access key for the storage account.
type StorageAccountKeyResponse struct {
	// Name of the key.
	KeyName string `pulumi:"keyName"`
	// Permissions for the key -- read-only or full permissions.
	Permissions string `pulumi:"permissions"`
	// Base 64-encoded value of the key.
	Value string `pulumi:"value"`
}

func init() {
	pulumi.RegisterOutputType(CustomDomainOutput{})
	pulumi.RegisterOutputType(CustomDomainPtrOutput{})
	pulumi.RegisterOutputType(CustomDomainResponseOutput{})
	pulumi.RegisterOutputType(CustomDomainResponsePtrOutput{})
	pulumi.RegisterOutputType(EncryptionOutput{})
	pulumi.RegisterOutputType(EncryptionPtrOutput{})
	pulumi.RegisterOutputType(EncryptionResponseOutput{})
	pulumi.RegisterOutputType(EncryptionResponsePtrOutput{})
	pulumi.RegisterOutputType(EncryptionServiceOutput{})
	pulumi.RegisterOutputType(EncryptionServicePtrOutput{})
	pulumi.RegisterOutputType(EncryptionServiceResponseOutput{})
	pulumi.RegisterOutputType(EncryptionServiceResponsePtrOutput{})
	pulumi.RegisterOutputType(EncryptionServicesOutput{})
	pulumi.RegisterOutputType(EncryptionServicesPtrOutput{})
	pulumi.RegisterOutputType(EncryptionServicesResponseOutput{})
	pulumi.RegisterOutputType(EncryptionServicesResponsePtrOutput{})
	pulumi.RegisterOutputType(EndpointsResponseOutput{})
	pulumi.RegisterOutputType(EndpointsResponsePtrOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuPtrOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SkuResponsePtrOutput{})
	pulumi.RegisterOutputType(StorageAccountKeyResponseOutput{})
	pulumi.RegisterOutputType(StorageAccountKeyResponseArrayOutput{})
}
