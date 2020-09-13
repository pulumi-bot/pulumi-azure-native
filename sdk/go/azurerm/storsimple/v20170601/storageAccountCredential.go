// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20170601

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The storage account credential.
//
// ## Example Usage
// ### StorageAccountCredentialsCreateOrUpdate
//
// ```go
// package main
//
// import (
// 	storsimple "github.com/pulumi/pulumi-azurerm/sdk/go/azurerm/storsimple/v20170601"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := storsimple.NewStorageAccountCredential(ctx, "storageAccountCredential", &storsimple.StorageAccountCredentialArgs{
// 			AccessKey: &storsimple.AsymmetricEncryptedSecretArgs{
// 				EncryptionAlgorithm:      pulumi.String("RSAES_PKCS1_v_1_5"),
// 				EncryptionCertThumbprint: pulumi.String("A872A2DF196AC7682EE24791E7DE2E2A360F5926"),
// 				Value:                    pulumi.String("ATuJSkmrFk4h8r1jrZ4nd3nthLSddcguEO5QLO/NECUtTuB9kL4dNv3/jC4WOvFkeVr3x1UvfhlIeMmJBF1SMr6hR1JzD0xNU/TtQqUeXN7V3jk7I+2l67P9StuHWR6OMd3XOLwvznxOEQtEWpweDiobZU1ZiY03WafcGZFpV5j6tEoHeopoZ1J/GhPtkYmx+TqxzUN6qnir5rP3NSYiZciImP/qu8U9yUV/xpVRv39KvFc2Yr5SpKpMMRUj55XW10UnPer63M6KovF8X9Wi/fNnrZAs1Esl5XddZETGrW/e5B++VMJ6w0Q/uvPR+UBwrOU0804l0SzwdIe3qVVd0Q=="),
// 			},
// 			EndPoint:                     pulumi.String("blob.core.windows.net"),
// 			ManagerName:                  pulumi.String("ManagerForSDKTest1"),
// 			ResourceGroupName:            pulumi.String("ResourceGroupForSDKTest"),
// 			SslStatus:                    pulumi.String("Enabled"),
// 			StorageAccountCredentialName: pulumi.String("SACForTest"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
//
// ```
type StorageAccountCredential struct {
	pulumi.CustomResourceState

	// The details of the storage account password.
	AccessKey AsymmetricEncryptedSecretResponsePtrOutput `pulumi:"accessKey"`
	// The storage endpoint
	EndPoint pulumi.StringOutput `pulumi:"endPoint"`
	// The Kind of the object. Currently only Series8000 is supported
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// The name of the object.
	Name pulumi.StringOutput `pulumi:"name"`
	// Signifies whether SSL needs to be enabled or not.
	SslStatus pulumi.StringOutput `pulumi:"sslStatus"`
	// The hierarchical type of the object.
	Type pulumi.StringOutput `pulumi:"type"`
	// The count of volumes using this storage account credential.
	VolumesCount pulumi.IntOutput `pulumi:"volumesCount"`
}

// NewStorageAccountCredential registers a new resource with the given unique name, arguments, and options.
func NewStorageAccountCredential(ctx *pulumi.Context,
	name string, args *StorageAccountCredentialArgs, opts ...pulumi.ResourceOption) (*StorageAccountCredential, error) {
	if args == nil || args.EndPoint == nil {
		return nil, errors.New("missing required argument 'EndPoint'")
	}
	if args == nil || args.ManagerName == nil {
		return nil, errors.New("missing required argument 'ManagerName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.SslStatus == nil {
		return nil, errors.New("missing required argument 'SslStatus'")
	}
	if args == nil || args.StorageAccountCredentialName == nil {
		return nil, errors.New("missing required argument 'StorageAccountCredentialName'")
	}
	if args == nil {
		args = &StorageAccountCredentialArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:storsimple/latest:StorageAccountCredential"),
		},
		{
			Type: pulumi.String("azurerm:storsimple/v20161001:StorageAccountCredential"),
		},
	})
	opts = append(opts, aliases)
	var resource StorageAccountCredential
	err := ctx.RegisterResource("azurerm:storsimple/v20170601:StorageAccountCredential", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetStorageAccountCredential gets an existing StorageAccountCredential resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetStorageAccountCredential(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StorageAccountCredentialState, opts ...pulumi.ResourceOption) (*StorageAccountCredential, error) {
	var resource StorageAccountCredential
	err := ctx.ReadResource("azurerm:storsimple/v20170601:StorageAccountCredential", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering StorageAccountCredential resources.
type storageAccountCredentialState struct {
	// The details of the storage account password.
	AccessKey *AsymmetricEncryptedSecretResponse `pulumi:"accessKey"`
	// The storage endpoint
	EndPoint *string `pulumi:"endPoint"`
	// The Kind of the object. Currently only Series8000 is supported
	Kind *string `pulumi:"kind"`
	// The name of the object.
	Name *string `pulumi:"name"`
	// Signifies whether SSL needs to be enabled or not.
	SslStatus *string `pulumi:"sslStatus"`
	// The hierarchical type of the object.
	Type *string `pulumi:"type"`
	// The count of volumes using this storage account credential.
	VolumesCount *int `pulumi:"volumesCount"`
}

type StorageAccountCredentialState struct {
	// The details of the storage account password.
	AccessKey AsymmetricEncryptedSecretResponsePtrInput
	// The storage endpoint
	EndPoint pulumi.StringPtrInput
	// The Kind of the object. Currently only Series8000 is supported
	Kind pulumi.StringPtrInput
	// The name of the object.
	Name pulumi.StringPtrInput
	// Signifies whether SSL needs to be enabled or not.
	SslStatus pulumi.StringPtrInput
	// The hierarchical type of the object.
	Type pulumi.StringPtrInput
	// The count of volumes using this storage account credential.
	VolumesCount pulumi.IntPtrInput
}

func (StorageAccountCredentialState) ElementType() reflect.Type {
	return reflect.TypeOf((*storageAccountCredentialState)(nil)).Elem()
}

type storageAccountCredentialArgs struct {
	// The details of the storage account password.
	AccessKey *AsymmetricEncryptedSecret `pulumi:"accessKey"`
	// The storage endpoint
	EndPoint string `pulumi:"endPoint"`
	// The Kind of the object. Currently only Series8000 is supported
	Kind *string `pulumi:"kind"`
	// The manager name
	ManagerName string `pulumi:"managerName"`
	// The resource group name
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Signifies whether SSL needs to be enabled or not.
	SslStatus string `pulumi:"sslStatus"`
	// The storage account credential name.
	StorageAccountCredentialName string `pulumi:"storageAccountCredentialName"`
}

// The set of arguments for constructing a StorageAccountCredential resource.
type StorageAccountCredentialArgs struct {
	// The details of the storage account password.
	AccessKey AsymmetricEncryptedSecretPtrInput
	// The storage endpoint
	EndPoint pulumi.StringInput
	// The Kind of the object. Currently only Series8000 is supported
	Kind pulumi.StringPtrInput
	// The manager name
	ManagerName pulumi.StringInput
	// The resource group name
	ResourceGroupName pulumi.StringInput
	// Signifies whether SSL needs to be enabled or not.
	SslStatus pulumi.StringInput
	// The storage account credential name.
	StorageAccountCredentialName pulumi.StringInput
}

func (StorageAccountCredentialArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*storageAccountCredentialArgs)(nil)).Elem()
}
