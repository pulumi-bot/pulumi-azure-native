// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200201preview

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// An Azure resource which represents Maps Creator product and provides ability to manage private location data.
type Creator struct {
	pulumi.CustomResourceState

	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The Creator resource properties.
	Properties CreatorPropertiesResponseOutput `pulumi:"properties"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewCreator registers a new resource with the given unique name, arguments, and options.
func NewCreator(ctx *pulumi.Context,
	name string, args *CreatorArgs, opts ...pulumi.ResourceOption) (*Creator, error) {
	if args == nil || args.AccountName == nil {
		return nil, errors.New("missing required argument 'AccountName'")
	}
	if args == nil || args.CreatorName == nil {
		return nil, errors.New("missing required argument 'CreatorName'")
	}
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &CreatorArgs{}
	}
	var resource Creator
	err := ctx.RegisterResource("azurerm:maps/v20200201preview:Creator", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCreator gets an existing Creator resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCreator(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CreatorState, opts ...pulumi.ResourceOption) (*Creator, error) {
	var resource Creator
	err := ctx.ReadResource("azurerm:maps/v20200201preview:Creator", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Creator resources.
type creatorState struct {
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the resource
	Name *string `pulumi:"name"`
	// The Creator resource properties.
	Properties *CreatorPropertiesResponse `pulumi:"properties"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `pulumi:"type"`
}

type CreatorState struct {
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The name of the resource
	Name pulumi.StringPtrInput
	// The Creator resource properties.
	Properties CreatorPropertiesResponsePtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringPtrInput
}

func (CreatorState) ElementType() reflect.Type {
	return reflect.TypeOf((*creatorState)(nil)).Elem()
}

type creatorArgs struct {
	// The name of the Maps Account.
	AccountName string `pulumi:"accountName"`
	// The name of the Maps Creator instance.
	CreatorName string `pulumi:"creatorName"`
	// The location of the resource.
	Location string `pulumi:"location"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Gets or sets a list of key value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Creator resource.
type CreatorArgs struct {
	// The name of the Maps Account.
	AccountName pulumi.StringInput
	// The name of the Maps Creator instance.
	CreatorName pulumi.StringInput
	// The location of the resource.
	Location pulumi.StringInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Gets or sets a list of key value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater than 128 characters and value no greater than 256 characters.
	Tags pulumi.StringMapInput
}

func (CreatorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*creatorArgs)(nil)).Elem()
}
