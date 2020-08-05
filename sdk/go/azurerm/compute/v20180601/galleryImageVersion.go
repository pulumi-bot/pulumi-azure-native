// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180601

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Specifies information about the gallery Image Version that you want to create or update.
type GalleryImageVersion struct {
	pulumi.CustomResourceState

	// Resource location
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Describes the properties of a gallery Image Version.
	Properties GalleryImageVersionPropertiesResponseOutput `pulumi:"properties"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewGalleryImageVersion registers a new resource with the given unique name, arguments, and options.
func NewGalleryImageVersion(ctx *pulumi.Context,
	name string, args *GalleryImageVersionArgs, opts ...pulumi.ResourceOption) (*GalleryImageVersion, error) {
	if args == nil || args.GalleryImageName == nil {
		return nil, errors.New("missing required argument 'GalleryImageName'")
	}
	if args == nil || args.GalleryName == nil {
		return nil, errors.New("missing required argument 'GalleryName'")
	}
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.PublishingProfile == nil {
		return nil, errors.New("missing required argument 'PublishingProfile'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &GalleryImageVersionArgs{}
	}
	var resource GalleryImageVersion
	err := ctx.RegisterResource("azurerm:compute/v20180601:GalleryImageVersion", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGalleryImageVersion gets an existing GalleryImageVersion resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGalleryImageVersion(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GalleryImageVersionState, opts ...pulumi.ResourceOption) (*GalleryImageVersion, error) {
	var resource GalleryImageVersion
	err := ctx.ReadResource("azurerm:compute/v20180601:GalleryImageVersion", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GalleryImageVersion resources.
type galleryImageVersionState struct {
	// Resource location
	Location *string `pulumi:"location"`
	// Resource name
	Name *string `pulumi:"name"`
	// Describes the properties of a gallery Image Version.
	Properties *GalleryImageVersionPropertiesResponse `pulumi:"properties"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type *string `pulumi:"type"`
}

type GalleryImageVersionState struct {
	// Resource location
	Location pulumi.StringPtrInput
	// Resource name
	Name pulumi.StringPtrInput
	// Describes the properties of a gallery Image Version.
	Properties GalleryImageVersionPropertiesResponsePtrInput
	// Resource tags
	Tags pulumi.StringMapInput
	// Resource type
	Type pulumi.StringPtrInput
}

func (GalleryImageVersionState) ElementType() reflect.Type {
	return reflect.TypeOf((*galleryImageVersionState)(nil)).Elem()
}

type galleryImageVersionArgs struct {
	// The name of the gallery Image Definition in which the Image Version is to be created.
	GalleryImageName string `pulumi:"galleryImageName"`
	// The name of the Shared Image Gallery in which the Image Definition resides.
	GalleryName string `pulumi:"galleryName"`
	// Resource location
	Location string `pulumi:"location"`
	// The name of the gallery Image Version to be created. Needs to follow semantic version name pattern: The allowed characters are digit and period. Digits must be within the range of a 32-bit integer. Format: <MajorVersion>.<MinorVersion>.<Patch>
	Name string `pulumi:"name"`
	// The publishing profile of a gallery Image Version.
	PublishingProfile GalleryImageVersionPublishingProfile `pulumi:"publishingProfile"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a GalleryImageVersion resource.
type GalleryImageVersionArgs struct {
	// The name of the gallery Image Definition in which the Image Version is to be created.
	GalleryImageName pulumi.StringInput
	// The name of the Shared Image Gallery in which the Image Definition resides.
	GalleryName pulumi.StringInput
	// Resource location
	Location pulumi.StringInput
	// The name of the gallery Image Version to be created. Needs to follow semantic version name pattern: The allowed characters are digit and period. Digits must be within the range of a 32-bit integer. Format: <MajorVersion>.<MinorVersion>.<Patch>
	Name pulumi.StringInput
	// The publishing profile of a gallery Image Version.
	PublishingProfile GalleryImageVersionPublishingProfileInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Resource tags
	Tags pulumi.StringMapInput
}

func (GalleryImageVersionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*galleryImageVersionArgs)(nil)).Elem()
}
