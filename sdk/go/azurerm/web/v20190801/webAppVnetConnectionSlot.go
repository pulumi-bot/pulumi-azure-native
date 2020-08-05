// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190801

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Virtual Network information contract.
type WebAppVnetConnectionSlot struct {
	pulumi.CustomResourceState

	// Kind of resource.
	Kind pulumi.StringPtrOutput `pulumi:"kind"`
	// Resource Name.
	Name pulumi.StringOutput `pulumi:"name"`
	// VnetInfo resource specific properties
	Properties VnetInfoResponsePropertiesOutput `pulumi:"properties"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewWebAppVnetConnectionSlot registers a new resource with the given unique name, arguments, and options.
func NewWebAppVnetConnectionSlot(ctx *pulumi.Context,
	name string, args *WebAppVnetConnectionSlotArgs, opts ...pulumi.ResourceOption) (*WebAppVnetConnectionSlot, error) {
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.Slot == nil {
		return nil, errors.New("missing required argument 'Slot'")
	}
	if args == nil {
		args = &WebAppVnetConnectionSlotArgs{}
	}
	var resource WebAppVnetConnectionSlot
	err := ctx.RegisterResource("azurerm:web/v20190801:WebAppVnetConnectionSlot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWebAppVnetConnectionSlot gets an existing WebAppVnetConnectionSlot resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWebAppVnetConnectionSlot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAppVnetConnectionSlotState, opts ...pulumi.ResourceOption) (*WebAppVnetConnectionSlot, error) {
	var resource WebAppVnetConnectionSlot
	err := ctx.ReadResource("azurerm:web/v20190801:WebAppVnetConnectionSlot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WebAppVnetConnectionSlot resources.
type webAppVnetConnectionSlotState struct {
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Resource Name.
	Name *string `pulumi:"name"`
	// VnetInfo resource specific properties
	Properties *VnetInfoResponseProperties `pulumi:"properties"`
	// Resource type.
	Type *string `pulumi:"type"`
}

type WebAppVnetConnectionSlotState struct {
	// Kind of resource.
	Kind pulumi.StringPtrInput
	// Resource Name.
	Name pulumi.StringPtrInput
	// VnetInfo resource specific properties
	Properties VnetInfoResponsePropertiesPtrInput
	// Resource type.
	Type pulumi.StringPtrInput
}

func (WebAppVnetConnectionSlotState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppVnetConnectionSlotState)(nil)).Elem()
}

type webAppVnetConnectionSlotArgs struct {
	// A certificate file (.cer) blob containing the public key of the private key used to authenticate a
	// Point-To-Site VPN connection.
	CertBlob *string `pulumi:"certBlob"`
	// DNS servers to be used by this Virtual Network. This should be a comma-separated list of IP addresses.
	DnsServers *string `pulumi:"dnsServers"`
	// Flag that is used to denote if this is VNET injection
	IsSwift *bool `pulumi:"isSwift"`
	// Kind of resource.
	Kind *string `pulumi:"kind"`
	// Name of an existing Virtual Network.
	Name string `pulumi:"name"`
	// Name of the resource group to which the resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the deployment slot. If a slot is not specified, the API will add or update connections for the production slot.
	Slot string `pulumi:"slot"`
	// The Virtual Network's resource ID.
	VnetResourceId *string `pulumi:"vnetResourceId"`
}

// The set of arguments for constructing a WebAppVnetConnectionSlot resource.
type WebAppVnetConnectionSlotArgs struct {
	// A certificate file (.cer) blob containing the public key of the private key used to authenticate a
	// Point-To-Site VPN connection.
	CertBlob pulumi.StringPtrInput
	// DNS servers to be used by this Virtual Network. This should be a comma-separated list of IP addresses.
	DnsServers pulumi.StringPtrInput
	// Flag that is used to denote if this is VNET injection
	IsSwift pulumi.BoolPtrInput
	// Kind of resource.
	Kind pulumi.StringPtrInput
	// Name of an existing Virtual Network.
	Name pulumi.StringInput
	// Name of the resource group to which the resource belongs.
	ResourceGroupName pulumi.StringInput
	// Name of the deployment slot. If a slot is not specified, the API will add or update connections for the production slot.
	Slot pulumi.StringInput
	// The Virtual Network's resource ID.
	VnetResourceId pulumi.StringPtrInput
}

func (WebAppVnetConnectionSlotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppVnetConnectionSlotArgs)(nil)).Elem()
}
