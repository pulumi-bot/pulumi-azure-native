// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The storage domain.
// Latest API Version: 2016-10-01.
//
// Deprecated: The 'latest' version is deprecated. Please migrate to the resource in the top-level module: 'azure-native:storsimple:StorageDomain'.
type StorageDomain struct {
	pulumi.CustomResourceState

	// The encryption key used to encrypt the data. This is a user secret.
	EncryptionKey AsymmetricEncryptedSecretResponsePtrOutput `pulumi:"encryptionKey"`
	// The encryption status "Enabled | Disabled".
	EncryptionStatus pulumi.StringOutput `pulumi:"encryptionStatus"`
	// The name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The storage account credentials.
	StorageAccountCredentialIds pulumi.StringArrayOutput `pulumi:"storageAccountCredentialIds"`
	// The type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewStorageDomain registers a new resource with the given unique name, arguments, and options.
func NewStorageDomain(ctx *pulumi.Context,
	name string, args *StorageDomainArgs, opts ...pulumi.ResourceOption) (*StorageDomain, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ManagerName == nil {
		return nil, errors.New("invalid value for required argument 'ManagerName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.StorageAccountCredentialIds == nil {
		return nil, errors.New("invalid value for required argument 'StorageAccountCredentialIds'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:storsimple/latest:StorageDomain"),
		},
		{
			Type: pulumi.String("azure-native:storsimple:StorageDomain"),
		},
		{
			Type: pulumi.String("azure-nextgen:storsimple:StorageDomain"),
		},
		{
			Type: pulumi.String("azure-native:storsimple/v20161001:StorageDomain"),
		},
		{
			Type: pulumi.String("azure-nextgen:storsimple/v20161001:StorageDomain"),
		},
	})
	opts = append(opts, aliases)
	var resource StorageDomain
	err := ctx.RegisterResource("azure-native:storsimple/latest:StorageDomain", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStorageDomain gets an existing StorageDomain resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStorageDomain(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StorageDomainState, opts ...pulumi.ResourceOption) (*StorageDomain, error) {
	var resource StorageDomain
	err := ctx.ReadResource("azure-native:storsimple/latest:StorageDomain", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StorageDomain resources.
type storageDomainState struct {
	// The encryption key used to encrypt the data. This is a user secret.
	EncryptionKey *AsymmetricEncryptedSecretResponse `pulumi:"encryptionKey"`
	// The encryption status "Enabled | Disabled".
	EncryptionStatus *string `pulumi:"encryptionStatus"`
	// The name.
	Name *string `pulumi:"name"`
	// The storage account credentials.
	StorageAccountCredentialIds []string `pulumi:"storageAccountCredentialIds"`
	// The type.
	Type *string `pulumi:"type"`
}

type StorageDomainState struct {
	// The encryption key used to encrypt the data. This is a user secret.
	EncryptionKey AsymmetricEncryptedSecretResponsePtrInput
	// The encryption status "Enabled | Disabled".
	EncryptionStatus pulumi.StringPtrInput
	// The name.
	Name pulumi.StringPtrInput
	// The storage account credentials.
	StorageAccountCredentialIds pulumi.StringArrayInput
	// The type.
	Type pulumi.StringPtrInput
}

func (StorageDomainState) ElementType() reflect.Type {
	return reflect.TypeOf((*storageDomainState)(nil)).Elem()
}

type storageDomainArgs struct {
	// The encryption key used to encrypt the data. This is a user secret.
	EncryptionKey *AsymmetricEncryptedSecret `pulumi:"encryptionKey"`
	// The encryption status "Enabled | Disabled".
	EncryptionStatus string `pulumi:"encryptionStatus"`
	// The manager name
	ManagerName string `pulumi:"managerName"`
	// The resource group name
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The storage account credentials.
	StorageAccountCredentialIds []string `pulumi:"storageAccountCredentialIds"`
	// The storage domain name.
	StorageDomainName *string `pulumi:"storageDomainName"`
}

// The set of arguments for constructing a StorageDomain resource.
type StorageDomainArgs struct {
	// The encryption key used to encrypt the data. This is a user secret.
	EncryptionKey AsymmetricEncryptedSecretPtrInput
	// The encryption status "Enabled | Disabled".
	EncryptionStatus EncryptionStatus
	// The manager name
	ManagerName pulumi.StringInput
	// The resource group name
	ResourceGroupName pulumi.StringInput
	// The storage account credentials.
	StorageAccountCredentialIds pulumi.StringArrayInput
	// The storage domain name.
	StorageDomainName pulumi.StringPtrInput
}

func (StorageDomainArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*storageDomainArgs)(nil)).Elem()
}

type StorageDomainInput interface {
	pulumi.Input

	ToStorageDomainOutput() StorageDomainOutput
	ToStorageDomainOutputWithContext(ctx context.Context) StorageDomainOutput
}

func (*StorageDomain) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageDomain)(nil))
}

func (i *StorageDomain) ToStorageDomainOutput() StorageDomainOutput {
	return i.ToStorageDomainOutputWithContext(context.Background())
}

func (i *StorageDomain) ToStorageDomainOutputWithContext(ctx context.Context) StorageDomainOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageDomainOutput)
}

type StorageDomainOutput struct {
	*pulumi.OutputState
}

func (StorageDomainOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageDomain)(nil))
}

func (o StorageDomainOutput) ToStorageDomainOutput() StorageDomainOutput {
	return o
}

func (o StorageDomainOutput) ToStorageDomainOutputWithContext(ctx context.Context) StorageDomainOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(StorageDomainOutput{})
}
