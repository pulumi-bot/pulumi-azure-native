// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package network

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// P2SVpnGateway Resource.
type P2svpnGateway struct {
	pulumi.CustomResourceState

	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Resource location.
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Properties of the P2SVpnGateway.
	Properties P2SVpnGatewayPropertiesResponseOutput `pulumi:"properties"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewP2svpnGateway registers a new resource with the given unique name, arguments, and options.
func NewP2svpnGateway(ctx *pulumi.Context,
	name string, args *P2svpnGatewayArgs, opts ...pulumi.ResourceOption) (*P2svpnGateway, error) {
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &P2svpnGatewayArgs{}
	}
	var resource P2svpnGateway
	err := ctx.RegisterResource("azurerm:network:P2svpnGateway", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetP2svpnGateway gets an existing P2svpnGateway resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetP2svpnGateway(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *P2svpnGatewayState, opts ...pulumi.ResourceOption) (*P2svpnGateway, error) {
	var resource P2svpnGateway
	err := ctx.ReadResource("azurerm:network:P2svpnGateway", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering P2svpnGateway resources.
type p2svpnGatewayState struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Resource name.
	Name *string `pulumi:"name"`
	// Properties of the P2SVpnGateway.
	Properties *P2SVpnGatewayPropertiesResponse `pulumi:"properties"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type *string `pulumi:"type"`
}

type P2svpnGatewayState struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// Resource name.
	Name pulumi.StringPtrInput
	// Properties of the P2SVpnGateway.
	Properties P2SVpnGatewayPropertiesResponsePtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Resource type.
	Type pulumi.StringPtrInput
}

func (P2svpnGatewayState) ElementType() reflect.Type {
	return reflect.TypeOf((*p2svpnGatewayState)(nil)).Elem()
}

type p2svpnGatewayArgs struct {
	// Resource ID.
	Id *string `pulumi:"id"`
	// Resource location.
	Location string `pulumi:"location"`
	// The name of the gateway.
	Name string `pulumi:"name"`
	// Properties of the P2SVpnGateway.
	Properties *P2SVpnGatewayProperties `pulumi:"properties"`
	// The resource group name of the P2SVpnGateway.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a P2svpnGateway resource.
type P2svpnGatewayArgs struct {
	// Resource ID.
	Id pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringInput
	// The name of the gateway.
	Name pulumi.StringInput
	// Properties of the P2SVpnGateway.
	Properties P2SVpnGatewayPropertiesPtrInput
	// The resource group name of the P2SVpnGateway.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (P2svpnGatewayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*p2svpnGatewayArgs)(nil)).Elem()
}