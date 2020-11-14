// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200701

import (
	"context"
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
	// The private IP address of the IP configuration.
	PrivateIPAddress pulumi.StringPtrOutput `pulumi:"privateIPAddress"`
	// The private IP address allocation method.
	PrivateIPAllocationMethod pulumi.StringPtrOutput `pulumi:"privateIPAllocationMethod"`
	// The provisioning state of the IP configuration resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The reference to the public IP resource.
	PublicIPAddress PublicIPAddressResponsePtrOutput `pulumi:"publicIPAddress"`
	// The reference to the subnet resource.
	Subnet SubnetResponsePtrOutput `pulumi:"subnet"`
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
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:network/latest:VirtualHubIpConfiguration"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200501:VirtualHubIpConfiguration"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200601:VirtualHubIpConfiguration"),
		},
	})
	opts = append(opts, aliases)
	var resource VirtualHubIpConfiguration
	err := ctx.RegisterResource("azure-nextgen:network/v20200701:VirtualHubIpConfiguration", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure-nextgen:network/v20200701:VirtualHubIpConfiguration", name, id, state, &resource, opts...)
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
	// The private IP address of the IP configuration.
	PrivateIPAddress *string `pulumi:"privateIPAddress"`
	// The private IP address allocation method.
	PrivateIPAllocationMethod *string `pulumi:"privateIPAllocationMethod"`
	// The provisioning state of the IP configuration resource.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The reference to the public IP resource.
	PublicIPAddress *PublicIPAddressResponse `pulumi:"publicIPAddress"`
	// The reference to the subnet resource.
	Subnet *SubnetResponse `pulumi:"subnet"`
	// Ipconfiguration type.
	Type *string `pulumi:"type"`
}

type VirtualHubIpConfigurationState struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrInput
	// Name of the Ip Configuration.
	Name pulumi.StringPtrInput
	// The private IP address of the IP configuration.
	PrivateIPAddress pulumi.StringPtrInput
	// The private IP address allocation method.
	PrivateIPAllocationMethod pulumi.StringPtrInput
	// The provisioning state of the IP configuration resource.
	ProvisioningState pulumi.StringPtrInput
	// The reference to the public IP resource.
	PublicIPAddress PublicIPAddressResponsePtrInput
	// The reference to the subnet resource.
	Subnet SubnetResponsePtrInput
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
	// The private IP address of the IP configuration.
	PrivateIPAddress *string `pulumi:"privateIPAddress"`
	// The private IP address allocation method.
	PrivateIPAllocationMethod *string `pulumi:"privateIPAllocationMethod"`
	// The reference to the public IP resource.
	PublicIPAddress *PublicIPAddressType `pulumi:"publicIPAddress"`
	// The resource group name of the VirtualHub.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The reference to the subnet resource.
	Subnet *SubnetType `pulumi:"subnet"`
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
	// The private IP address of the IP configuration.
	PrivateIPAddress pulumi.StringPtrInput
	// The private IP address allocation method.
	PrivateIPAllocationMethod pulumi.StringPtrInput
	// The reference to the public IP resource.
	PublicIPAddress PublicIPAddressTypePtrInput
	// The resource group name of the VirtualHub.
	ResourceGroupName pulumi.StringInput
	// The reference to the subnet resource.
	Subnet SubnetTypePtrInput
	// The name of the VirtualHub.
	VirtualHubName pulumi.StringInput
}

func (VirtualHubIpConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualHubIpConfigurationArgs)(nil)).Elem()
}

type VirtualHubIpConfigurationInput interface {
	pulumi.Input

	ToVirtualHubIpConfigurationOutput() VirtualHubIpConfigurationOutput
	ToVirtualHubIpConfigurationOutputWithContext(ctx context.Context) VirtualHubIpConfigurationOutput
}

func (VirtualHubIpConfiguration) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualHubIpConfiguration)(nil)).Elem()
}

func (i VirtualHubIpConfiguration) ToVirtualHubIpConfigurationOutput() VirtualHubIpConfigurationOutput {
	return i.ToVirtualHubIpConfigurationOutputWithContext(context.Background())
}

func (i VirtualHubIpConfiguration) ToVirtualHubIpConfigurationOutputWithContext(ctx context.Context) VirtualHubIpConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualHubIpConfigurationOutput)
}

type VirtualHubIpConfigurationOutput struct {
	*pulumi.OutputState
}

func (VirtualHubIpConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualHubIpConfigurationOutput)(nil)).Elem()
}

func (o VirtualHubIpConfigurationOutput) ToVirtualHubIpConfigurationOutput() VirtualHubIpConfigurationOutput {
	return o
}

func (o VirtualHubIpConfigurationOutput) ToVirtualHubIpConfigurationOutputWithContext(ctx context.Context) VirtualHubIpConfigurationOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(VirtualHubIpConfigurationOutput{})
}
