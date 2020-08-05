// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200401

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Peering is a logical representation of a set of connections to the Microsoft Cloud Edge at a location.
type Peering struct {
	pulumi.CustomResourceState

	// The kind of the peering.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// The location of the resource.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The properties that define a peering.
	Properties PeeringPropertiesResponseOutput `pulumi:"properties"`
	// The SKU that defines the tier and kind of the peering.
	Sku PeeringSkuResponseOutput `pulumi:"sku"`
	// The resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewPeering registers a new resource with the given unique name, arguments, and options.
func NewPeering(ctx *pulumi.Context,
	name string, args *PeeringArgs, opts ...pulumi.ResourceOption) (*Peering, error) {
	if args == nil || args.Kind == nil {
		return nil, errors.New("missing required argument 'Kind'")
	}
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.Sku == nil {
		return nil, errors.New("missing required argument 'Sku'")
	}
	if args == nil {
		args = &PeeringArgs{}
	}
	var resource Peering
	err := ctx.RegisterResource("azurerm:peering/v20200401:Peering", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPeering gets an existing Peering resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPeering(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PeeringState, opts ...pulumi.ResourceOption) (*Peering, error) {
	var resource Peering
	err := ctx.ReadResource("azurerm:peering/v20200401:Peering", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Peering resources.
type peeringState struct {
	// The kind of the peering.
	Kind *string `pulumi:"kind"`
	// The location of the resource.
	Location *string `pulumi:"location"`
	// The name of the resource.
	Name *string `pulumi:"name"`
	// The properties that define a peering.
	Properties *PeeringPropertiesResponse `pulumi:"properties"`
	// The SKU that defines the tier and kind of the peering.
	Sku *PeeringSkuResponse `pulumi:"sku"`
	// The resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource.
	Type *string `pulumi:"type"`
}

type PeeringState struct {
	// The kind of the peering.
	Kind pulumi.StringPtrInput
	// The location of the resource.
	Location pulumi.StringPtrInput
	// The name of the resource.
	Name pulumi.StringPtrInput
	// The properties that define a peering.
	Properties PeeringPropertiesResponsePtrInput
	// The SKU that defines the tier and kind of the peering.
	Sku PeeringSkuResponsePtrInput
	// The resource tags.
	Tags pulumi.StringMapInput
	// The type of the resource.
	Type pulumi.StringPtrInput
}

func (PeeringState) ElementType() reflect.Type {
	return reflect.TypeOf((*peeringState)(nil)).Elem()
}

type peeringArgs struct {
	// The properties that define a direct peering.
	Direct *PeeringPropertiesDirect `pulumi:"direct"`
	// The properties that define an exchange peering.
	Exchange *PeeringPropertiesExchange `pulumi:"exchange"`
	// The kind of the peering.
	Kind string `pulumi:"kind"`
	// The location of the resource.
	Location string `pulumi:"location"`
	// The name of the peering.
	Name string `pulumi:"name"`
	// The location of the peering.
	PeeringLocation *string `pulumi:"peeringLocation"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The SKU that defines the tier and kind of the peering.
	Sku PeeringSku `pulumi:"sku"`
	// The resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Peering resource.
type PeeringArgs struct {
	// The properties that define a direct peering.
	Direct PeeringPropertiesDirectPtrInput
	// The properties that define an exchange peering.
	Exchange PeeringPropertiesExchangePtrInput
	// The kind of the peering.
	Kind pulumi.StringInput
	// The location of the resource.
	Location pulumi.StringInput
	// The name of the peering.
	Name pulumi.StringInput
	// The location of the peering.
	PeeringLocation pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The SKU that defines the tier and kind of the peering.
	Sku PeeringSkuInput
	// The resource tags.
	Tags pulumi.StringMapInput
}

func (PeeringArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*peeringArgs)(nil)).Elem()
}
