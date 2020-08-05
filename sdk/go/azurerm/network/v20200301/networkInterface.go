// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200301

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A network interface in a resource group.
type NetworkInterface struct {
	pulumi.CustomResourceState

	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Properties of the network interface.
	Properties NetworkInterfacePropertiesFormatResponseOutput `pulumi:"properties"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewNetworkInterface registers a new resource with the given unique name, arguments, and options.
func NewNetworkInterface(ctx *pulumi.Context,
	name string, args *NetworkInterfaceArgs, opts ...pulumi.ResourceOption) (*NetworkInterface, error) {
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &NetworkInterfaceArgs{}
	}
	var resource NetworkInterface
	err := ctx.RegisterResource("azurerm:network/v20200301:NetworkInterface", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkInterface gets an existing NetworkInterface resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkInterface(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkInterfaceState, opts ...pulumi.ResourceOption) (*NetworkInterface, error) {
	var resource NetworkInterface
	err := ctx.ReadResource("azurerm:network/v20200301:NetworkInterface", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkInterface resources.
type networkInterfaceState struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Resource name.
	Name *string `pulumi:"name"`
	// Properties of the network interface.
	Properties *NetworkInterfacePropertiesFormatResponse `pulumi:"properties"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type *string `pulumi:"type"`
}

type NetworkInterfaceState struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// Resource name.
	Name pulumi.StringPtrInput
	// Properties of the network interface.
	Properties NetworkInterfacePropertiesFormatResponsePtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Resource type.
	Type pulumi.StringPtrInput
}

func (NetworkInterfaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkInterfaceState)(nil)).Elem()
}

type networkInterfaceArgs struct {
	// The DNS settings in network interface.
	DnsSettings *NetworkInterfaceDnsSettings `pulumi:"dnsSettings"`
	// If the network interface is accelerated networking enabled.
	EnableAcceleratedNetworking *bool `pulumi:"enableAcceleratedNetworking"`
	// Indicates whether IP forwarding is enabled on this network interface.
	EnableIPForwarding *bool `pulumi:"enableIPForwarding"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// A list of IPConfigurations of the network interface.
	IpConfigurations []NetworkInterfaceIPConfiguration `pulumi:"ipConfigurations"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The name of the network interface.
	Name string `pulumi:"name"`
	// The reference to the NetworkSecurityGroup resource.
	NetworkSecurityGroup *NetworkSecurityGroupType `pulumi:"networkSecurityGroup"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a NetworkInterface resource.
type NetworkInterfaceArgs struct {
	// The DNS settings in network interface.
	DnsSettings NetworkInterfaceDnsSettingsPtrInput
	// If the network interface is accelerated networking enabled.
	EnableAcceleratedNetworking pulumi.BoolPtrInput
	// Indicates whether IP forwarding is enabled on this network interface.
	EnableIPForwarding pulumi.BoolPtrInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// A list of IPConfigurations of the network interface.
	IpConfigurations NetworkInterfaceIPConfigurationArrayInput
	// Resource location.
	Location pulumi.StringPtrInput
	// The name of the network interface.
	Name pulumi.StringInput
	// The reference to the NetworkSecurityGroup resource.
	NetworkSecurityGroup NetworkSecurityGroupTypePtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (NetworkInterfaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkInterfaceArgs)(nil)).Elem()
}
