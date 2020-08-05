// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190401

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Tap configuration in a Network Interface.
type NetworkInterfaceTapConfiguration struct {
	pulumi.CustomResourceState

	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// Properties of the Virtual Network Tap configuration.
	Properties NetworkInterfaceTapConfigurationPropertiesFormatResponseOutput `pulumi:"properties"`
	// Sub Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewNetworkInterfaceTapConfiguration registers a new resource with the given unique name, arguments, and options.
func NewNetworkInterfaceTapConfiguration(ctx *pulumi.Context,
	name string, args *NetworkInterfaceTapConfigurationArgs, opts ...pulumi.ResourceOption) (*NetworkInterfaceTapConfiguration, error) {
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.NetworkInterfaceName == nil {
		return nil, errors.New("missing required argument 'NetworkInterfaceName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &NetworkInterfaceTapConfigurationArgs{}
	}
	var resource NetworkInterfaceTapConfiguration
	err := ctx.RegisterResource("azurerm:network/v20190401:NetworkInterfaceTapConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkInterfaceTapConfiguration gets an existing NetworkInterfaceTapConfiguration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkInterfaceTapConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkInterfaceTapConfigurationState, opts ...pulumi.ResourceOption) (*NetworkInterfaceTapConfiguration, error) {
	var resource NetworkInterfaceTapConfiguration
	err := ctx.ReadResource("azurerm:network/v20190401:NetworkInterfaceTapConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkInterfaceTapConfiguration resources.
type networkInterfaceTapConfigurationState struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `pulumi:"name"`
	// Properties of the Virtual Network Tap configuration.
	Properties *NetworkInterfaceTapConfigurationPropertiesFormatResponse `pulumi:"properties"`
	// Sub Resource type.
	Type *string `pulumi:"type"`
}

type NetworkInterfaceTapConfigurationState struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrInput
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name pulumi.StringPtrInput
	// Properties of the Virtual Network Tap configuration.
	Properties NetworkInterfaceTapConfigurationPropertiesFormatResponsePtrInput
	// Sub Resource type.
	Type pulumi.StringPtrInput
}

func (NetworkInterfaceTapConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkInterfaceTapConfigurationState)(nil)).Elem()
}

type networkInterfaceTapConfigurationArgs struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// The name of the tap configuration.
	Name string `pulumi:"name"`
	// The name of the network interface.
	NetworkInterfaceName string `pulumi:"networkInterfaceName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The reference of the Virtual Network Tap resource.
	VirtualNetworkTap *VirtualNetworkTapType `pulumi:"virtualNetworkTap"`
}

// The set of arguments for constructing a NetworkInterfaceTapConfiguration resource.
type NetworkInterfaceTapConfigurationArgs struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// The name of the tap configuration.
	Name pulumi.StringInput
	// The name of the network interface.
	NetworkInterfaceName pulumi.StringInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The reference of the Virtual Network Tap resource.
	VirtualNetworkTap VirtualNetworkTapTypePtrInput
}

func (NetworkInterfaceTapConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkInterfaceTapConfigurationArgs)(nil)).Elem()
}
