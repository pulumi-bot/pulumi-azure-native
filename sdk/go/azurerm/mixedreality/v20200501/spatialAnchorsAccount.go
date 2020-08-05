// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200501

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// SpatialAnchorsAccount Response.
type SpatialAnchorsAccount struct {
	pulumi.CustomResourceState

	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Property bag.
	Properties MixedRealityAccountPropertiesResponseOutput `pulumi:"properties"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewSpatialAnchorsAccount registers a new resource with the given unique name, arguments, and options.
func NewSpatialAnchorsAccount(ctx *pulumi.Context,
	name string, args *SpatialAnchorsAccountArgs, opts ...pulumi.ResourceOption) (*SpatialAnchorsAccount, error) {
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
		args = &SpatialAnchorsAccountArgs{}
	}
	var resource SpatialAnchorsAccount
	err := ctx.RegisterResource("azurerm:mixedreality/v20200501:SpatialAnchorsAccount", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSpatialAnchorsAccount gets an existing SpatialAnchorsAccount resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSpatialAnchorsAccount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SpatialAnchorsAccountState, opts ...pulumi.ResourceOption) (*SpatialAnchorsAccount, error) {
	var resource SpatialAnchorsAccount
	err := ctx.ReadResource("azurerm:mixedreality/v20200501:SpatialAnchorsAccount", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SpatialAnchorsAccount resources.
type spatialAnchorsAccountState struct {
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the resource
	Name *string `pulumi:"name"`
	// Property bag.
	Properties *MixedRealityAccountPropertiesResponse `pulumi:"properties"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `pulumi:"type"`
}

type SpatialAnchorsAccountState struct {
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The name of the resource
	Name pulumi.StringPtrInput
	// Property bag.
	Properties MixedRealityAccountPropertiesResponsePtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringPtrInput
}

func (SpatialAnchorsAccountState) ElementType() reflect.Type {
	return reflect.TypeOf((*spatialAnchorsAccountState)(nil)).Elem()
}

type spatialAnchorsAccountArgs struct {
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// Name of an Mixed Reality Account.
	Name string `pulumi:"name"`
	// Name of an Azure resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a SpatialAnchorsAccount resource.
type SpatialAnchorsAccountArgs struct {
	// The geo-location where the resource lives
	Location pulumi.StringInput
	// Name of an Mixed Reality Account.
	Name pulumi.StringInput
	// Name of an Azure resource group.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (SpatialAnchorsAccountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*spatialAnchorsAccountArgs)(nil)).Elem()
}
