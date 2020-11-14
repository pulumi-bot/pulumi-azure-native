// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20161201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Subnet in a virtual network resource.
type Subnet struct {
	pulumi.CustomResourceState

	// The address prefix for the subnet.
	AddressPrefix pulumi.StringPtrOutput `pulumi:"addressPrefix"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// Gets an array of references to the network interface IP configurations using subnet.
	IpConfigurations IPConfigurationResponseArrayOutput `pulumi:"ipConfigurations"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// The reference of the NetworkSecurityGroup resource.
	NetworkSecurityGroup NetworkSecurityGroupResponsePtrOutput `pulumi:"networkSecurityGroup"`
	// The provisioning state of the resource.
	ProvisioningState pulumi.StringPtrOutput `pulumi:"provisioningState"`
	// Gets an array of references to the external resources using subnet.
	ResourceNavigationLinks ResourceNavigationLinkResponseArrayOutput `pulumi:"resourceNavigationLinks"`
	// The reference of the RouteTable resource.
	RouteTable RouteTableResponsePtrOutput `pulumi:"routeTable"`
}

// NewSubnet registers a new resource with the given unique name, arguments, and options.
func NewSubnet(ctx *pulumi.Context,
	name string, args *SubnetArgs, opts ...pulumi.ResourceOption) (*Subnet, error) {
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.SubnetName == nil {
		return nil, errors.New("missing required argument 'SubnetName'")
	}
	if args == nil || args.VirtualNetworkName == nil {
		return nil, errors.New("missing required argument 'VirtualNetworkName'")
	}
	if args == nil {
		args = &SubnetArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:network/latest:Subnet"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20150501preview:Subnet"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20150615:Subnet"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20160330:Subnet"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20160601:Subnet"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20160901:Subnet"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20170301:Subnet"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20170601:Subnet"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20170801:Subnet"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20170901:Subnet"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20171001:Subnet"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20171101:Subnet"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20180101:Subnet"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20180201:Subnet"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20180401:Subnet"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20180601:Subnet"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20180701:Subnet"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20180801:Subnet"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20181001:Subnet"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20181101:Subnet"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20181201:Subnet"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190201:Subnet"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190401:Subnet"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190601:Subnet"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190701:Subnet"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190801:Subnet"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20190901:Subnet"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20191101:Subnet"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20191201:Subnet"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200301:Subnet"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200401:Subnet"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200501:Subnet"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200601:Subnet"),
		},
		{
			Type: pulumi.String("azure-nextgen:network/v20200701:Subnet"),
		},
	})
	opts = append(opts, aliases)
	var resource Subnet
	err := ctx.RegisterResource("azure-nextgen:network/v20161201:Subnet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSubnet gets an existing Subnet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSubnet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SubnetState, opts ...pulumi.ResourceOption) (*Subnet, error) {
	var resource Subnet
	err := ctx.ReadResource("azure-nextgen:network/v20161201:Subnet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Subnet resources.
type subnetState struct {
	// The address prefix for the subnet.
	AddressPrefix *string `pulumi:"addressPrefix"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// Gets an array of references to the network interface IP configurations using subnet.
	IpConfigurations []IPConfigurationResponse `pulumi:"ipConfigurations"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `pulumi:"name"`
	// The reference of the NetworkSecurityGroup resource.
	NetworkSecurityGroup *NetworkSecurityGroupResponse `pulumi:"networkSecurityGroup"`
	// The provisioning state of the resource.
	ProvisioningState *string `pulumi:"provisioningState"`
	// Gets an array of references to the external resources using subnet.
	ResourceNavigationLinks []ResourceNavigationLinkResponse `pulumi:"resourceNavigationLinks"`
	// The reference of the RouteTable resource.
	RouteTable *RouteTableResponse `pulumi:"routeTable"`
}

type SubnetState struct {
	// The address prefix for the subnet.
	AddressPrefix pulumi.StringPtrInput
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrInput
	// Gets an array of references to the network interface IP configurations using subnet.
	IpConfigurations IPConfigurationResponseArrayInput
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name pulumi.StringPtrInput
	// The reference of the NetworkSecurityGroup resource.
	NetworkSecurityGroup NetworkSecurityGroupResponsePtrInput
	// The provisioning state of the resource.
	ProvisioningState pulumi.StringPtrInput
	// Gets an array of references to the external resources using subnet.
	ResourceNavigationLinks ResourceNavigationLinkResponseArrayInput
	// The reference of the RouteTable resource.
	RouteTable RouteTableResponsePtrInput
}

func (SubnetState) ElementType() reflect.Type {
	return reflect.TypeOf((*subnetState)(nil)).Elem()
}

type subnetArgs struct {
	// The address prefix for the subnet.
	AddressPrefix *string `pulumi:"addressPrefix"`
	// A unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `pulumi:"name"`
	// The reference of the NetworkSecurityGroup resource.
	NetworkSecurityGroup *NetworkSecurityGroupType `pulumi:"networkSecurityGroup"`
	// The provisioning state of the resource.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Gets an array of references to the external resources using subnet.
	ResourceNavigationLinks []ResourceNavigationLink `pulumi:"resourceNavigationLinks"`
	// The reference of the RouteTable resource.
	RouteTable *RouteTableType `pulumi:"routeTable"`
	// The name of the subnet.
	SubnetName string `pulumi:"subnetName"`
	// The name of the virtual network.
	VirtualNetworkName string `pulumi:"virtualNetworkName"`
}

// The set of arguments for constructing a Subnet resource.
type SubnetArgs struct {
	// The address prefix for the subnet.
	AddressPrefix pulumi.StringPtrInput
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name pulumi.StringPtrInput
	// The reference of the NetworkSecurityGroup resource.
	NetworkSecurityGroup NetworkSecurityGroupTypePtrInput
	// The provisioning state of the resource.
	ProvisioningState pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Gets an array of references to the external resources using subnet.
	ResourceNavigationLinks ResourceNavigationLinkArrayInput
	// The reference of the RouteTable resource.
	RouteTable RouteTableTypePtrInput
	// The name of the subnet.
	SubnetName pulumi.StringInput
	// The name of the virtual network.
	VirtualNetworkName pulumi.StringInput
}

func (SubnetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*subnetArgs)(nil)).Elem()
}

type SubnetInput interface {
	pulumi.Input

	ToSubnetOutput() SubnetOutput
	ToSubnetOutputWithContext(ctx context.Context) SubnetOutput
}

func (Subnet) ElementType() reflect.Type {
	return reflect.TypeOf((*Subnet)(nil)).Elem()
}

func (i Subnet) ToSubnetOutput() SubnetOutput {
	return i.ToSubnetOutputWithContext(context.Background())
}

func (i Subnet) ToSubnetOutputWithContext(ctx context.Context) SubnetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubnetOutput)
}

type SubnetOutput struct {
	*pulumi.OutputState
}

func (SubnetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubnetOutput)(nil)).Elem()
}

func (o SubnetOutput) ToSubnetOutput() SubnetOutput {
	return o
}

func (o SubnetOutput) ToSubnetOutputWithContext(ctx context.Context) SubnetOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SubnetOutput{})
}
