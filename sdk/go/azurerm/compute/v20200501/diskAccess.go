// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200501

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// disk access resource.
type DiskAccess struct {
	pulumi.CustomResourceState

	// Resource location
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource name
	Name       pulumi.StringOutput                `pulumi:"name"`
	Properties DiskAccessPropertiesResponseOutput `pulumi:"properties"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewDiskAccess registers a new resource with the given unique name, arguments, and options.
func NewDiskAccess(ctx *pulumi.Context,
	name string, args *DiskAccessArgs, opts ...pulumi.ResourceOption) (*DiskAccess, error) {
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
		args = &DiskAccessArgs{}
	}
	var resource DiskAccess
	err := ctx.RegisterResource("azurerm:compute/v20200501:DiskAccess", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDiskAccess gets an existing DiskAccess resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDiskAccess(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DiskAccessState, opts ...pulumi.ResourceOption) (*DiskAccess, error) {
	var resource DiskAccess
	err := ctx.ReadResource("azurerm:compute/v20200501:DiskAccess", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DiskAccess resources.
type diskAccessState struct {
	// Resource location
	Location *string `pulumi:"location"`
	// Resource name
	Name       *string                       `pulumi:"name"`
	Properties *DiskAccessPropertiesResponse `pulumi:"properties"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type *string `pulumi:"type"`
}

type DiskAccessState struct {
	// Resource location
	Location pulumi.StringPtrInput
	// Resource name
	Name       pulumi.StringPtrInput
	Properties DiskAccessPropertiesResponsePtrInput
	// Resource tags
	Tags pulumi.StringMapInput
	// Resource type
	Type pulumi.StringPtrInput
}

func (DiskAccessState) ElementType() reflect.Type {
	return reflect.TypeOf((*diskAccessState)(nil)).Elem()
}

type diskAccessArgs struct {
	// Resource location
	Location string `pulumi:"location"`
	// The name of the disk access resource that is being created. The name can't be changed after the disk encryption set is created. Supported characters for the name are a-z, A-Z, 0-9 and _. The maximum name length is 80 characters.
	Name string `pulumi:"name"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a DiskAccess resource.
type DiskAccessArgs struct {
	// Resource location
	Location pulumi.StringInput
	// The name of the disk access resource that is being created. The name can't be changed after the disk encryption set is created. Supported characters for the name are a-z, A-Z, 0-9 and _. The maximum name length is 80 characters.
	Name pulumi.StringInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Resource tags
	Tags pulumi.StringMapInput
}

func (DiskAccessArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*diskAccessArgs)(nil)).Elem()
}
