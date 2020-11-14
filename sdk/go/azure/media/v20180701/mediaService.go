// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180701

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A Media Services account.
type MediaService struct {
	pulumi.CustomResourceState

	// The Azure Region of the resource.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The Media Services account ID.
	MediaServiceId pulumi.StringOutput `pulumi:"mediaServiceId"`
	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The storage accounts for this resource.
	StorageAccounts StorageAccountResponseArrayOutput `pulumi:"storageAccounts"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewMediaService registers a new resource with the given unique name, arguments, and options.
func NewMediaService(ctx *pulumi.Context,
	name string, args *MediaServiceArgs, opts ...pulumi.ResourceOption) (*MediaService, error) {
	if args == nil || args.AccountName == nil {
		return nil, errors.New("missing required argument 'AccountName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &MediaServiceArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:media/latest:MediaService"),
		},
		{
			Type: pulumi.String("azure-nextgen:media/v20151001:MediaService"),
		},
		{
			Type: pulumi.String("azure-nextgen:media/v20180330preview:MediaService"),
		},
		{
			Type: pulumi.String("azure-nextgen:media/v20180601preview:MediaService"),
		},
		{
			Type: pulumi.String("azure-nextgen:media/v20200501:MediaService"),
		},
	})
	opts = append(opts, aliases)
	var resource MediaService
	err := ctx.RegisterResource("azure-nextgen:media/v20180701:MediaService", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMediaService gets an existing MediaService resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMediaService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MediaServiceState, opts ...pulumi.ResourceOption) (*MediaService, error) {
	var resource MediaService
	err := ctx.ReadResource("azure-nextgen:media/v20180701:MediaService", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MediaService resources.
type mediaServiceState struct {
	// The Azure Region of the resource.
	Location *string `pulumi:"location"`
	// The Media Services account ID.
	MediaServiceId *string `pulumi:"mediaServiceId"`
	// The name of the resource.
	Name *string `pulumi:"name"`
	// The storage accounts for this resource.
	StorageAccounts []StorageAccountResponse `pulumi:"storageAccounts"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource.
	Type *string `pulumi:"type"`
}

type MediaServiceState struct {
	// The Azure Region of the resource.
	Location pulumi.StringPtrInput
	// The Media Services account ID.
	MediaServiceId pulumi.StringPtrInput
	// The name of the resource.
	Name pulumi.StringPtrInput
	// The storage accounts for this resource.
	StorageAccounts StorageAccountResponseArrayInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The type of the resource.
	Type pulumi.StringPtrInput
}

func (MediaServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*mediaServiceState)(nil)).Elem()
}

type mediaServiceArgs struct {
	// The Media Services account name.
	AccountName string `pulumi:"accountName"`
	// The Azure Region of the resource.
	Location *string `pulumi:"location"`
	// The name of the resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The storage accounts for this resource.
	StorageAccounts []StorageAccount `pulumi:"storageAccounts"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a MediaService resource.
type MediaServiceArgs struct {
	// The Media Services account name.
	AccountName pulumi.StringInput
	// The Azure Region of the resource.
	Location pulumi.StringPtrInput
	// The name of the resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput
	// The storage accounts for this resource.
	StorageAccounts StorageAccountArrayInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (MediaServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mediaServiceArgs)(nil)).Elem()
}

type MediaServiceInput interface {
	pulumi.Input

	ToMediaServiceOutput() MediaServiceOutput
	ToMediaServiceOutputWithContext(ctx context.Context) MediaServiceOutput
}

func (MediaService) ElementType() reflect.Type {
	return reflect.TypeOf((*MediaService)(nil)).Elem()
}

func (i MediaService) ToMediaServiceOutput() MediaServiceOutput {
	return i.ToMediaServiceOutputWithContext(context.Background())
}

func (i MediaService) ToMediaServiceOutputWithContext(ctx context.Context) MediaServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MediaServiceOutput)
}

type MediaServiceOutput struct {
	*pulumi.OutputState
}

func (MediaServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MediaServiceOutput)(nil)).Elem()
}

func (o MediaServiceOutput) ToMediaServiceOutput() MediaServiceOutput {
	return o
}

func (o MediaServiceOutput) ToMediaServiceOutputWithContext(ctx context.Context) MediaServiceOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(MediaServiceOutput{})
}
