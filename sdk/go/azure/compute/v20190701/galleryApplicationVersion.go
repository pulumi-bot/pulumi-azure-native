// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190701

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Specifies information about the gallery Application Version that you want to create or update.
type GalleryApplicationVersion struct {
	pulumi.CustomResourceState

	// Resource location
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioning state, which only appears in the response.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The publishing profile of a gallery image version.
	PublishingProfile GalleryApplicationVersionPublishingProfileResponseOutput `pulumi:"publishingProfile"`
	// This is the replication status of the gallery Image Version.
	ReplicationStatus ReplicationStatusResponseOutput `pulumi:"replicationStatus"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewGalleryApplicationVersion registers a new resource with the given unique name, arguments, and options.
func NewGalleryApplicationVersion(ctx *pulumi.Context,
	name string, args *GalleryApplicationVersionArgs, opts ...pulumi.ResourceOption) (*GalleryApplicationVersion, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GalleryApplicationName == nil {
		return nil, errors.New("invalid value for required argument 'GalleryApplicationName'")
	}
	if args.GalleryApplicationVersionName == nil {
		return nil, errors.New("invalid value for required argument 'GalleryApplicationVersionName'")
	}
	if args.GalleryName == nil {
		return nil, errors.New("invalid value for required argument 'GalleryName'")
	}
	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.PublishingProfile == nil {
		return nil, errors.New("invalid value for required argument 'PublishingProfile'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:compute/latest:GalleryApplicationVersion"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20190301:GalleryApplicationVersion"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20191201:GalleryApplicationVersion"),
		},
		{
			Type: pulumi.String("azure-nextgen:compute/v20200930:GalleryApplicationVersion"),
		},
	})
	opts = append(opts, aliases)
	var resource GalleryApplicationVersion
	err := ctx.RegisterResource("azure-nextgen:compute/v20190701:GalleryApplicationVersion", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGalleryApplicationVersion gets an existing GalleryApplicationVersion resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGalleryApplicationVersion(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GalleryApplicationVersionState, opts ...pulumi.ResourceOption) (*GalleryApplicationVersion, error) {
	var resource GalleryApplicationVersion
	err := ctx.ReadResource("azure-nextgen:compute/v20190701:GalleryApplicationVersion", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GalleryApplicationVersion resources.
type galleryApplicationVersionState struct {
	// Resource location
	Location *string `pulumi:"location"`
	// Resource name
	Name *string `pulumi:"name"`
	// The provisioning state, which only appears in the response.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The publishing profile of a gallery image version.
	PublishingProfile *GalleryApplicationVersionPublishingProfileResponse `pulumi:"publishingProfile"`
	// This is the replication status of the gallery Image Version.
	ReplicationStatus *ReplicationStatusResponse `pulumi:"replicationStatus"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type *string `pulumi:"type"`
}

type GalleryApplicationVersionState struct {
	// Resource location
	Location pulumi.StringPtrInput
	// Resource name
	Name pulumi.StringPtrInput
	// The provisioning state, which only appears in the response.
	ProvisioningState pulumi.StringPtrInput
	// The publishing profile of a gallery image version.
	PublishingProfile GalleryApplicationVersionPublishingProfileResponsePtrInput
	// This is the replication status of the gallery Image Version.
	ReplicationStatus ReplicationStatusResponsePtrInput
	// Resource tags
	Tags pulumi.StringMapInput
	// Resource type
	Type pulumi.StringPtrInput
}

func (GalleryApplicationVersionState) ElementType() reflect.Type {
	return reflect.TypeOf((*galleryApplicationVersionState)(nil)).Elem()
}

type galleryApplicationVersionArgs struct {
	// The name of the gallery Application Definition in which the Application Version is to be created.
	GalleryApplicationName string `pulumi:"galleryApplicationName"`
	// The name of the gallery Application Version to be created. Needs to follow semantic version name pattern: The allowed characters are digit and period. Digits must be within the range of a 32-bit integer. Format: <MajorVersion>.<MinorVersion>.<Patch>
	GalleryApplicationVersionName string `pulumi:"galleryApplicationVersionName"`
	// The name of the Shared Application Gallery in which the Application Definition resides.
	GalleryName string `pulumi:"galleryName"`
	// Resource location
	Location string `pulumi:"location"`
	// The publishing profile of a gallery image version.
	PublishingProfile GalleryApplicationVersionPublishingProfile `pulumi:"publishingProfile"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a GalleryApplicationVersion resource.
type GalleryApplicationVersionArgs struct {
	// The name of the gallery Application Definition in which the Application Version is to be created.
	GalleryApplicationName pulumi.StringInput
	// The name of the gallery Application Version to be created. Needs to follow semantic version name pattern: The allowed characters are digit and period. Digits must be within the range of a 32-bit integer. Format: <MajorVersion>.<MinorVersion>.<Patch>
	GalleryApplicationVersionName pulumi.StringInput
	// The name of the Shared Application Gallery in which the Application Definition resides.
	GalleryName pulumi.StringInput
	// Resource location
	Location pulumi.StringInput
	// The publishing profile of a gallery image version.
	PublishingProfile GalleryApplicationVersionPublishingProfileInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Resource tags
	Tags pulumi.StringMapInput
}

func (GalleryApplicationVersionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*galleryApplicationVersionArgs)(nil)).Elem()
}

type GalleryApplicationVersionInput interface {
	pulumi.Input

	ToGalleryApplicationVersionOutput() GalleryApplicationVersionOutput
	ToGalleryApplicationVersionOutputWithContext(ctx context.Context) GalleryApplicationVersionOutput
}

func (*GalleryApplicationVersion) ElementType() reflect.Type {
	return reflect.TypeOf((*GalleryApplicationVersion)(nil))
}

func (i *GalleryApplicationVersion) ToGalleryApplicationVersionOutput() GalleryApplicationVersionOutput {
	return i.ToGalleryApplicationVersionOutputWithContext(context.Background())
}

func (i *GalleryApplicationVersion) ToGalleryApplicationVersionOutputWithContext(ctx context.Context) GalleryApplicationVersionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GalleryApplicationVersionOutput)
}

type GalleryApplicationVersionOutput struct {
	*pulumi.OutputState
}

func (GalleryApplicationVersionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GalleryApplicationVersion)(nil))
}

func (o GalleryApplicationVersionOutput) ToGalleryApplicationVersionOutput() GalleryApplicationVersionOutput {
	return o
}

func (o GalleryApplicationVersionOutput) ToGalleryApplicationVersionOutputWithContext(ctx context.Context) GalleryApplicationVersionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(GalleryApplicationVersionOutput{})
}
