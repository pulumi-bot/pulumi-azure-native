// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190701

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Specifies information about the Shared Image Gallery that you want to create or update.
type Gallery struct {
	pulumi.CustomResourceState

	// Resource location
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Describes the properties of a Shared Image Gallery.
	Properties GalleryPropertiesResponseOutput `pulumi:"properties"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewGallery registers a new resource with the given unique name, arguments, and options.
func NewGallery(ctx *pulumi.Context,
	name string, args *GalleryArgs, opts ...pulumi.ResourceOption) (*Gallery, error) {
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &GalleryArgs{}
	}
	var resource Gallery
	err := ctx.RegisterResource("azurerm:compute/v20190701:Gallery", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGallery gets an existing Gallery resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGallery(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GalleryState, opts ...pulumi.ResourceOption) (*Gallery, error) {
	var resource Gallery
	err := ctx.ReadResource("azurerm:compute/v20190701:Gallery", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Gallery resources.
type galleryState struct {
	// Resource location
	Location *string `pulumi:"location"`
	// Resource name
	Name *string `pulumi:"name"`
	// Describes the properties of a Shared Image Gallery.
	Properties *GalleryPropertiesResponse `pulumi:"properties"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type *string `pulumi:"type"`
}

type GalleryState struct {
	// Resource location
	Location pulumi.StringPtrInput
	// Resource name
	Name pulumi.StringPtrInput
	// Describes the properties of a Shared Image Gallery.
	Properties GalleryPropertiesResponsePtrInput
	// Resource tags
	Tags pulumi.StringMapInput
	// Resource type
	Type pulumi.StringPtrInput
}

func (GalleryState) ElementType() reflect.Type {
	return reflect.TypeOf((*galleryState)(nil)).Elem()
}

type galleryArgs struct {
	// The description of this Shared Image Gallery resource. This property is updatable.
	Description *string `pulumi:"description"`
	// Describes the gallery unique name.
	Identifier *GalleryIdentifier `pulumi:"identifier"`
	// Resource location
	Location string `pulumi:"location"`
	// The name of the Shared Image Gallery. The allowed characters are alphabets and numbers with dots and periods allowed in the middle. The maximum length is 80 characters.
	Name string `pulumi:"name"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Gallery resource.
type GalleryArgs struct {
	// The description of this Shared Image Gallery resource. This property is updatable.
	Description pulumi.StringPtrInput
	// Describes the gallery unique name.
	Identifier GalleryIdentifierPtrInput
	// Resource location
	Location pulumi.StringInput
	// The name of the Shared Image Gallery. The allowed characters are alphabets and numbers with dots and periods allowed in the middle. The maximum length is 80 characters.
	Name pulumi.StringInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Resource tags
	Tags pulumi.StringMapInput
}

func (GalleryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*galleryArgs)(nil)).Elem()
}
