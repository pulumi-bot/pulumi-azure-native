// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// IpConfigurations.
type VirtualHubIpConfiguration struct {
	pulumi.CustomResourceState

	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Name of the Ip Configuration.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// The properties of the Virtual Hub IPConfigurations.
	Properties HubIPConfigurationPropertiesFormatResponseOutput `pulumi:"properties"`
	// Ipconfiguration type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewVirtualHubIpConfiguration registers a new resource with the given unique name, arguments, and options.
func NewVirtualHubIpConfiguration(ctx *pulumi.Context,
	name string, args *VirtualHubIpConfigurationArgs, opts ...pulumi.ResourceOption) (*VirtualHubIpConfiguration, error) {
	if args == nil || args.IpConfigName == nil {
		return nil, errors.New("missing required argument 'IpConfigName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.VirtualHubName == nil {
		return nil, errors.New("missing required argument 'VirtualHubName'")
	}
	if args == nil {
		args = &VirtualHubIpConfigurationArgs{}
	}
	var resource VirtualHubIpConfiguration
	err := ctx.RegisterResource("azurerm:network:VirtualHubIpConfiguration", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualHubIpConfiguration gets an existing VirtualHubIpConfiguration resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualHubIpConfiguration(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualHubIpConfigurationState, opts ...pulumi.ResourceOption) (*VirtualHubIpConfiguration, error) {
	var resource VirtualHubIpConfiguration
	err := ctx.ReadResource("azurerm:network:VirtualHubIpConfiguration", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualHubIpConfiguration resources.
type virtualHubIpConfigurationState struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// Name of the Ip Configuration.
	Name *string `pulumi:"name"`
	// The properties of the Virtual Hub IPConfigurations.
	Properties *HubIPConfigurationPropertiesFormatResponse `pulumi:"properties"`
	// Ipconfiguration type.
	Type *string `pulumi:"type"`
}

type VirtualHubIpConfigurationState struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrInput
	// Name of the Ip Configuration.
	Name pulumi.StringPtrInput
	// The properties of the Virtual Hub IPConfigurations.
	Properties HubIPConfigurationPropertiesFormatResponsePtrInput
	// Ipconfiguration type.
	Type pulumi.StringPtrInput
}

func (VirtualHubIpConfigurationState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualHubIpConfigurationState)(nil)).Elem()
}

type virtualHubIpConfigurationArgs struct {
	// Resource ID.
	Id *string `pulumi:"id"`
	// The name of the ipconfig.
	IpConfigName string `pulumi:"ipConfigName"`
	// Name of the Ip Configuration.
	Name *string `pulumi:"name"`
	// The properties of the Virtual Hub IPConfigurations.
	Properties *HubIPConfigurationPropertiesFormat `pulumi:"properties"`
	// The resource group name of the VirtualHub.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the VirtualHub.
	VirtualHubName string `pulumi:"virtualHubName"`
}

// The set of arguments for constructing a VirtualHubIpConfiguration resource.
type VirtualHubIpConfigurationArgs struct {
	// Resource ID.
	Id pulumi.StringPtrInput
	// The name of the ipconfig.
	IpConfigName pulumi.StringInput
	// Name of the Ip Configuration.
	Name pulumi.StringPtrInput
	// The properties of the Virtual Hub IPConfigurations.
	Properties HubIPConfigurationPropertiesFormatPtrInput
	// The resource group name of the VirtualHub.
	ResourceGroupName pulumi.StringInput
	// The name of the VirtualHub.
	VirtualHubName pulumi.StringInput
}

func (VirtualHubIpConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualHubIpConfigurationArgs)(nil)).Elem()
}