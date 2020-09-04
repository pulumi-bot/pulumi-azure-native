// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200401

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The customer's ASN that is registered by the peering service provider.
type RegisteredAsn struct {
	pulumi.CustomResourceState

	// The customer's ASN from which traffic originates.
	Asn pulumi.IntPtrOutput `pulumi:"asn"`
	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The peering service prefix key that is to be shared with the customer.
	PeeringServicePrefixKey pulumi.StringOutput `pulumi:"peeringServicePrefixKey"`
	// The provisioning state of the resource.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewRegisteredAsn registers a new resource with the given unique name, arguments, and options.
func NewRegisteredAsn(ctx *pulumi.Context,
	name string, args *RegisteredAsnArgs, opts ...pulumi.ResourceOption) (*RegisteredAsn, error) {
	if args == nil || args.PeeringName == nil {
		return nil, errors.New("missing required argument 'PeeringName'")
	}
	if args == nil || args.RegisteredAsnName == nil {
		return nil, errors.New("missing required argument 'RegisteredAsnName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &RegisteredAsnArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:peering/latest:RegisteredAsn"),
		},
		{
			Type: pulumi.String("azurerm:peering/v20200101preview:RegisteredAsn"),
		},
	})
	opts = append(opts, aliases)
	var resource RegisteredAsn
	err := ctx.RegisterResource("azurerm:peering/v20200401:RegisteredAsn", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRegisteredAsn gets an existing RegisteredAsn resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegisteredAsn(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegisteredAsnState, opts ...pulumi.ResourceOption) (*RegisteredAsn, error) {
	var resource RegisteredAsn
	err := ctx.ReadResource("azurerm:peering/v20200401:RegisteredAsn", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RegisteredAsn resources.
type registeredAsnState struct {
	// The customer's ASN from which traffic originates.
	Asn *int `pulumi:"asn"`
	// The name of the resource.
	Name *string `pulumi:"name"`
	// The peering service prefix key that is to be shared with the customer.
	PeeringServicePrefixKey *string `pulumi:"peeringServicePrefixKey"`
	// The provisioning state of the resource.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The type of the resource.
	Type *string `pulumi:"type"`
}

type RegisteredAsnState struct {
	// The customer's ASN from which traffic originates.
	Asn pulumi.IntPtrInput
	// The name of the resource.
	Name pulumi.StringPtrInput
	// The peering service prefix key that is to be shared with the customer.
	PeeringServicePrefixKey pulumi.StringPtrInput
	// The provisioning state of the resource.
	ProvisioningState pulumi.StringPtrInput
	// The type of the resource.
	Type pulumi.StringPtrInput
}

func (RegisteredAsnState) ElementType() reflect.Type {
	return reflect.TypeOf((*registeredAsnState)(nil)).Elem()
}

type registeredAsnArgs struct {
	// The customer's ASN from which traffic originates.
	Asn *int `pulumi:"asn"`
	// The name of the peering.
	PeeringName string `pulumi:"peeringName"`
	// The name of the ASN.
	RegisteredAsnName string `pulumi:"registeredAsnName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a RegisteredAsn resource.
type RegisteredAsnArgs struct {
	// The customer's ASN from which traffic originates.
	Asn pulumi.IntPtrInput
	// The name of the peering.
	PeeringName pulumi.StringInput
	// The name of the ASN.
	RegisteredAsnName pulumi.StringInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
}

func (RegisteredAsnArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*registeredAsnArgs)(nil)).Elem()
}
