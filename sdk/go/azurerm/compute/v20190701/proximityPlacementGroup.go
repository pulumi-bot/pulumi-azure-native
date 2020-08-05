// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190701

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Specifies information about the proximity placement group.
type ProximityPlacementGroup struct {
	pulumi.CustomResourceState

	// Resource location
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Describes the properties of a Proximity Placement Group.
	Properties ProximityPlacementGroupPropertiesResponseOutput `pulumi:"properties"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewProximityPlacementGroup registers a new resource with the given unique name, arguments, and options.
func NewProximityPlacementGroup(ctx *pulumi.Context,
	name string, args *ProximityPlacementGroupArgs, opts ...pulumi.ResourceOption) (*ProximityPlacementGroup, error) {
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
		args = &ProximityPlacementGroupArgs{}
	}
	var resource ProximityPlacementGroup
	err := ctx.RegisterResource("azurerm:compute/v20190701:ProximityPlacementGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProximityPlacementGroup gets an existing ProximityPlacementGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProximityPlacementGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProximityPlacementGroupState, opts ...pulumi.ResourceOption) (*ProximityPlacementGroup, error) {
	var resource ProximityPlacementGroup
	err := ctx.ReadResource("azurerm:compute/v20190701:ProximityPlacementGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProximityPlacementGroup resources.
type proximityPlacementGroupState struct {
	// Resource location
	Location *string `pulumi:"location"`
	// Resource name
	Name *string `pulumi:"name"`
	// Describes the properties of a Proximity Placement Group.
	Properties *ProximityPlacementGroupPropertiesResponse `pulumi:"properties"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type *string `pulumi:"type"`
}

type ProximityPlacementGroupState struct {
	// Resource location
	Location pulumi.StringPtrInput
	// Resource name
	Name pulumi.StringPtrInput
	// Describes the properties of a Proximity Placement Group.
	Properties ProximityPlacementGroupPropertiesResponsePtrInput
	// Resource tags
	Tags pulumi.StringMapInput
	// Resource type
	Type pulumi.StringPtrInput
}

func (ProximityPlacementGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*proximityPlacementGroupState)(nil)).Elem()
}

type proximityPlacementGroupArgs struct {
	// Describes colocation status of the Proximity Placement Group.
	ColocationStatus *InstanceViewStatus `pulumi:"colocationStatus"`
	// Resource location
	Location string `pulumi:"location"`
	// The name of the proximity placement group.
	Name string `pulumi:"name"`
	// Specifies the type of the proximity placement group. <br><br> Possible values are: <br><br> **Standard** : Co-locate resources within an Azure region or Availability Zone. <br><br> **Ultra** : For future use.
	ProximityPlacementGroupType *string `pulumi:"proximityPlacementGroupType"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a ProximityPlacementGroup resource.
type ProximityPlacementGroupArgs struct {
	// Describes colocation status of the Proximity Placement Group.
	ColocationStatus InstanceViewStatusPtrInput
	// Resource location
	Location pulumi.StringInput
	// The name of the proximity placement group.
	Name pulumi.StringInput
	// Specifies the type of the proximity placement group. <br><br> Possible values are: <br><br> **Standard** : Co-locate resources within an Azure region or Availability Zone. <br><br> **Ultra** : For future use.
	ProximityPlacementGroupType pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Resource tags
	Tags pulumi.StringMapInput
}

func (ProximityPlacementGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*proximityPlacementGroupArgs)(nil)).Elem()
}
