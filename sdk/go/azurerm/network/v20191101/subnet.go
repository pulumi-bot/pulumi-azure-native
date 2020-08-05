// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20191101

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Subnet in a virtual network resource.
type Subnet struct {
	pulumi.CustomResourceState

	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// Properties of the subnet.
	Properties SubnetPropertiesFormatResponseOutput `pulumi:"properties"`
}

// NewSubnet registers a new resource with the given unique name, arguments, and options.
func NewSubnet(ctx *pulumi.Context,
	name string, args *SubnetArgs, opts ...pulumi.ResourceOption) (*Subnet, error) {
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.VirtualNetworkName == nil {
		return nil, errors.New("missing required argument 'VirtualNetworkName'")
	}
	if args == nil {
		args = &SubnetArgs{}
	}
	var resource Subnet
	err := ctx.RegisterResource("azurerm:network/v20191101:Subnet", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azurerm:network/v20191101:Subnet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Subnet resources.
type subnetState struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `pulumi:"name"`
	// Properties of the subnet.
	Properties *SubnetPropertiesFormatResponse `pulumi:"properties"`
}

type SubnetState struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrInput
	// The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name pulumi.StringPtrInput
	// Properties of the subnet.
	Properties SubnetPropertiesFormatResponsePtrInput
}

func (SubnetState) ElementType() reflect.Type {
	return reflect.TypeOf((*subnetState)(nil)).Elem()
}

type subnetArgs struct {
	// The address prefix for the subnet.
	AddressPrefix *string `pulumi:"addressPrefix"`
	// List of address prefixes for the subnet.
	AddressPrefixes []string `pulumi:"addressPrefixes"`
	// An array of references to the delegations on the subnet.
	Delegations []Delegation `pulumi:"delegations"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// The name of the subnet.
	Name string `pulumi:"name"`
	// Nat gateway associated with this subnet.
	NatGateway *SubResource `pulumi:"natGateway"`
	// The reference to the NetworkSecurityGroup resource.
	NetworkSecurityGroup *NetworkSecurityGroupType `pulumi:"networkSecurityGroup"`
	// Enable or Disable apply network policies on private end point in the subnet.
	PrivateEndpointNetworkPolicies *string `pulumi:"privateEndpointNetworkPolicies"`
	// Enable or Disable apply network policies on private link service in the subnet.
	PrivateLinkServiceNetworkPolicies *string `pulumi:"privateLinkServiceNetworkPolicies"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The reference to the RouteTable resource.
	RouteTable *RouteTableType `pulumi:"routeTable"`
	// An array of service endpoint policies.
	ServiceEndpointPolicies []ServiceEndpointPolicyType `pulumi:"serviceEndpointPolicies"`
	// An array of service endpoints.
	ServiceEndpoints []ServiceEndpointPropertiesFormat `pulumi:"serviceEndpoints"`
	// The name of the virtual network.
	VirtualNetworkName string `pulumi:"virtualNetworkName"`
}

// The set of arguments for constructing a Subnet resource.
type SubnetArgs struct {
	// The address prefix for the subnet.
	AddressPrefix pulumi.StringPtrInput
	// List of address prefixes for the subnet.
	AddressPrefixes pulumi.StringArrayInput
	// An array of references to the delegations on the subnet.
	Delegations DelegationArrayInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// The name of the subnet.
	Name pulumi.StringInput
	// Nat gateway associated with this subnet.
	NatGateway SubResourcePtrInput
	// The reference to the NetworkSecurityGroup resource.
	NetworkSecurityGroup NetworkSecurityGroupTypePtrInput
	// Enable or Disable apply network policies on private end point in the subnet.
	PrivateEndpointNetworkPolicies pulumi.StringPtrInput
	// Enable or Disable apply network policies on private link service in the subnet.
	PrivateLinkServiceNetworkPolicies pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The reference to the RouteTable resource.
	RouteTable RouteTableTypePtrInput
	// An array of service endpoint policies.
	ServiceEndpointPolicies ServiceEndpointPolicyTypeArrayInput
	// An array of service endpoints.
	ServiceEndpoints ServiceEndpointPropertiesFormatArrayInput
	// The name of the virtual network.
	VirtualNetworkName pulumi.StringInput
}

func (SubnetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*subnetArgs)(nil)).Elem()
}
