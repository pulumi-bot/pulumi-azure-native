// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20160515

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A virtual network.
type VirtualNetwork struct {
	pulumi.CustomResourceState

	// The location of the resource.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The properties of the resource.
	Properties VirtualNetworkPropertiesResponseOutput `pulumi:"properties"`
	// The tags of the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewVirtualNetwork registers a new resource with the given unique name, arguments, and options.
func NewVirtualNetwork(ctx *pulumi.Context,
	name string, args *VirtualNetworkArgs, opts ...pulumi.ResourceOption) (*VirtualNetwork, error) {
	if args == nil || args.LabName == nil {
		return nil, errors.New("missing required argument 'LabName'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &VirtualNetworkArgs{}
	}
	var resource VirtualNetwork
	err := ctx.RegisterResource("azurerm:devtestlab/v20160515:VirtualNetwork", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualNetwork gets an existing VirtualNetwork resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualNetwork(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualNetworkState, opts ...pulumi.ResourceOption) (*VirtualNetwork, error) {
	var resource VirtualNetwork
	err := ctx.ReadResource("azurerm:devtestlab/v20160515:VirtualNetwork", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualNetwork resources.
type virtualNetworkState struct {
	// The location of the resource.
	Location *string `pulumi:"location"`
	// The name of the resource.
	Name *string `pulumi:"name"`
	// The properties of the resource.
	Properties *VirtualNetworkPropertiesResponse `pulumi:"properties"`
	// The tags of the resource.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource.
	Type *string `pulumi:"type"`
}

type VirtualNetworkState struct {
	// The location of the resource.
	Location pulumi.StringPtrInput
	// The name of the resource.
	Name pulumi.StringPtrInput
	// The properties of the resource.
	Properties VirtualNetworkPropertiesResponsePtrInput
	// The tags of the resource.
	Tags pulumi.StringMapInput
	// The type of the resource.
	Type pulumi.StringPtrInput
}

func (VirtualNetworkState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualNetworkState)(nil)).Elem()
}

type virtualNetworkArgs struct {
	// The allowed subnets of the virtual network.
	AllowedSubnets []Subnet `pulumi:"allowedSubnets"`
	// The description of the virtual network.
	Description *string `pulumi:"description"`
	// The Microsoft.Network resource identifier of the virtual network.
	ExternalProviderResourceId *string `pulumi:"externalProviderResourceId"`
	// The external subnet properties.
	ExternalSubnets []ExternalSubnet `pulumi:"externalSubnets"`
	// The name of the lab.
	LabName string `pulumi:"labName"`
	// The location of the resource.
	Location *string `pulumi:"location"`
	// The name of the virtual network.
	Name string `pulumi:"name"`
	// The provisioning status of the resource.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The subnet overrides of the virtual network.
	SubnetOverrides []SubnetOverride `pulumi:"subnetOverrides"`
	// The tags of the resource.
	Tags map[string]string `pulumi:"tags"`
	// The unique immutable identifier of a resource (Guid).
	UniqueIdentifier *string `pulumi:"uniqueIdentifier"`
}

// The set of arguments for constructing a VirtualNetwork resource.
type VirtualNetworkArgs struct {
	// The allowed subnets of the virtual network.
	AllowedSubnets SubnetArrayInput
	// The description of the virtual network.
	Description pulumi.StringPtrInput
	// The Microsoft.Network resource identifier of the virtual network.
	ExternalProviderResourceId pulumi.StringPtrInput
	// The external subnet properties.
	ExternalSubnets ExternalSubnetArrayInput
	// The name of the lab.
	LabName pulumi.StringInput
	// The location of the resource.
	Location pulumi.StringPtrInput
	// The name of the virtual network.
	Name pulumi.StringInput
	// The provisioning status of the resource.
	ProvisioningState pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The subnet overrides of the virtual network.
	SubnetOverrides SubnetOverrideArrayInput
	// The tags of the resource.
	Tags pulumi.StringMapInput
	// The unique immutable identifier of a resource (Guid).
	UniqueIdentifier pulumi.StringPtrInput
}

func (VirtualNetworkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualNetworkArgs)(nil)).Elem()
}
