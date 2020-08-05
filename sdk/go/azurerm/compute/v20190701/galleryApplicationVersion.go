// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190701

import (
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
	// Describes the properties of a gallery Image Version.
	Properties GalleryApplicationVersionPropertiesResponseOutput `pulumi:"properties"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewGalleryApplicationVersion registers a new resource with the given unique name, arguments, and options.
func NewGalleryApplicationVersion(ctx *pulumi.Context,
	name string, args *GalleryApplicationVersionArgs, opts ...pulumi.ResourceOption) (*GalleryApplicationVersion, error) {
	if args == nil || args.GalleryApplicationName == nil {
		return nil, errors.New("missing required argument 'GalleryApplicationName'")
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
		args = &GalleryApplicationVersionArgs{}
	}
	var resource GalleryApplicationVersion
	err := ctx.RegisterResource("azurerm:compute/v20190701:GalleryApplicationVersion", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azurerm:compute/v20190701:GalleryApplicationVersion", name, id, state, &resource, opts...)
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
	// Describes the properties of a gallery Image Version.
	Properties *GalleryApplicationVersionPropertiesResponse `pulumi:"properties"`
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
	// Describes the properties of a gallery Image Version.
	Properties GalleryApplicationVersionPropertiesResponsePtrInput
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
	// The name of the Shared Application Gallery in which the Application Definition resides.
	GalleryName string `pulumi:"galleryName"`
	// Resource location
	Location string `pulumi:"location"`
	// The name of the gallery Application Version to be created. Needs to follow semantic version name pattern: The allowed characters are digit and period. Digits must be within the range of a 32-bit integer. Format: <MajorVersion>.<MinorVersion>.<Patch>
	Name string `pulumi:"name"`
	// The publishing profile of a gallery Image Version.
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
	// The name of the Shared Application Gallery in which the Application Definition resides.
	GalleryName pulumi.StringInput
	// Resource location
	Location pulumi.StringInput
	// The name of the gallery Application Version to be created. Needs to follow semantic version name pattern: The allowed characters are digit and period. Digits must be within the range of a 32-bit integer. Format: <MajorVersion>.<MinorVersion>.<Patch>
	Name pulumi.StringInput
	// The publishing profile of a gallery Image Version.
	PublishingProfile GalleryApplicationVersionPublishingProfileInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Resource tags
	Tags pulumi.StringMapInput
}

func (GalleryApplicationVersionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*galleryApplicationVersionArgs)(nil)).Elem()
}
