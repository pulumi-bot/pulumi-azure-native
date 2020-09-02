// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180201

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The ImmutabilityPolicy property of a blob container, including Id, resource name, resource type, Etag.
type BlobContainerImmutabilityPolicy struct {
	pulumi.CustomResourceState

	// Resource Etag.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The immutability period for the blobs in the container since the policy creation, in days.
	ImmutabilityPeriodSinceCreationInDays pulumi.IntOutput `pulumi:"immutabilityPeriodSinceCreationInDays"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The ImmutabilityPolicy state of a blob container, possible values include: Locked and Unlocked.
	State pulumi.StringOutput `pulumi:"state"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewBlobContainerImmutabilityPolicy registers a new resource with the given unique name, arguments, and options.
func NewBlobContainerImmutabilityPolicy(ctx *pulumi.Context,
	name string, args *BlobContainerImmutabilityPolicyArgs, opts ...pulumi.ResourceOption) (*BlobContainerImmutabilityPolicy, error) {
	if args == nil || args.AccountName == nil {
		return nil, errors.New("missing required argument 'AccountName'")
	}
	if args == nil || args.ContainerName == nil {
		return nil, errors.New("missing required argument 'ContainerName'")
	}
	if args == nil || args.ImmutabilityPeriodSinceCreationInDays == nil {
		return nil, errors.New("missing required argument 'ImmutabilityPeriodSinceCreationInDays'")
	}
	if args == nil || args.ImmutabilityPolicyName == nil {
		return nil, errors.New("missing required argument 'ImmutabilityPolicyName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &BlobContainerImmutabilityPolicyArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:storage/latest:BlobContainerImmutabilityPolicy"),
		},
		{
			Type: pulumi.String("azurerm:storage/v20180301preview:BlobContainerImmutabilityPolicy"),
		},
		{
			Type: pulumi.String("azurerm:storage/v20180701:BlobContainerImmutabilityPolicy"),
		},
		{
			Type: pulumi.String("azurerm:storage/v20181101:BlobContainerImmutabilityPolicy"),
		},
		{
			Type: pulumi.String("azurerm:storage/v20190401:BlobContainerImmutabilityPolicy"),
		},
		{
			Type: pulumi.String("azurerm:storage/v20190601:BlobContainerImmutabilityPolicy"),
		},
	})
	opts = append(opts, aliases)
	var resource BlobContainerImmutabilityPolicy
	err := ctx.RegisterResource("azurerm:storage/v20180201:BlobContainerImmutabilityPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBlobContainerImmutabilityPolicy gets an existing BlobContainerImmutabilityPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBlobContainerImmutabilityPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BlobContainerImmutabilityPolicyState, opts ...pulumi.ResourceOption) (*BlobContainerImmutabilityPolicy, error) {
	var resource BlobContainerImmutabilityPolicy
	err := ctx.ReadResource("azurerm:storage/v20180201:BlobContainerImmutabilityPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BlobContainerImmutabilityPolicy resources.
type blobContainerImmutabilityPolicyState struct {
	// Resource Etag.
	Etag *string `pulumi:"etag"`
	// The immutability period for the blobs in the container since the policy creation, in days.
	ImmutabilityPeriodSinceCreationInDays *int `pulumi:"immutabilityPeriodSinceCreationInDays"`
	// The name of the resource
	Name *string `pulumi:"name"`
	// The ImmutabilityPolicy state of a blob container, possible values include: Locked and Unlocked.
	State *string `pulumi:"state"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `pulumi:"type"`
}

type BlobContainerImmutabilityPolicyState struct {
	// Resource Etag.
	Etag pulumi.StringPtrInput
	// The immutability period for the blobs in the container since the policy creation, in days.
	ImmutabilityPeriodSinceCreationInDays pulumi.IntPtrInput
	// The name of the resource
	Name pulumi.StringPtrInput
	// The ImmutabilityPolicy state of a blob container, possible values include: Locked and Unlocked.
	State pulumi.StringPtrInput
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringPtrInput
}

func (BlobContainerImmutabilityPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*blobContainerImmutabilityPolicyState)(nil)).Elem()
}

type blobContainerImmutabilityPolicyArgs struct {
	// The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
	AccountName string `pulumi:"accountName"`
	// The name of the blob container within the specified storage account. Blob container names must be between 3 and 63 characters in length and use numbers, lower-case letters and dash (-) only. Every dash (-) character must be immediately preceded and followed by a letter or number.
	ContainerName string `pulumi:"containerName"`
	// The immutability period for the blobs in the container since the policy creation, in days.
	ImmutabilityPeriodSinceCreationInDays int `pulumi:"immutabilityPeriodSinceCreationInDays"`
	// The name of the blob container immutabilityPolicy within the specified storage account. ImmutabilityPolicy Name must be 'default'
	ImmutabilityPolicyName string `pulumi:"immutabilityPolicyName"`
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a BlobContainerImmutabilityPolicy resource.
type BlobContainerImmutabilityPolicyArgs struct {
	// The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.
	AccountName pulumi.StringInput
	// The name of the blob container within the specified storage account. Blob container names must be between 3 and 63 characters in length and use numbers, lower-case letters and dash (-) only. Every dash (-) character must be immediately preceded and followed by a letter or number.
	ContainerName pulumi.StringInput
	// The immutability period for the blobs in the container since the policy creation, in days.
	ImmutabilityPeriodSinceCreationInDays pulumi.IntInput
	// The name of the blob container immutabilityPolicy within the specified storage account. ImmutabilityPolicy Name must be 'default'
	ImmutabilityPolicyName pulumi.StringInput
	// The name of the resource group within the user's subscription. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
}

func (BlobContainerImmutabilityPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*blobContainerImmutabilityPolicyArgs)(nil)).Elem()
}
