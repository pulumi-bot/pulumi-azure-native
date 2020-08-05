// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190701

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Virtual Router Peering resource
type VirtualRouterPeering struct {
	pulumi.CustomResourceState

	// Gets a unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Gets name of the peering unique to VirtualRouter. This name can be used to access the resource.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// The properties of the Virtual Router Peering.
	Properties VirtualRouterPeeringPropertiesResponseOutput `pulumi:"properties"`
	// Peering type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewVirtualRouterPeering registers a new resource with the given unique name, arguments, and options.
func NewVirtualRouterPeering(ctx *pulumi.Context,
	name string, args *VirtualRouterPeeringArgs, opts ...pulumi.ResourceOption) (*VirtualRouterPeering, error) {
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.VirtualRouterName == nil {
		return nil, errors.New("missing required argument 'VirtualRouterName'")
	}
	if args == nil {
		args = &VirtualRouterPeeringArgs{}
	}
	var resource VirtualRouterPeering
	err := ctx.RegisterResource("azurerm:network/v20190701:VirtualRouterPeering", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualRouterPeering gets an existing VirtualRouterPeering resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualRouterPeering(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualRouterPeeringState, opts ...pulumi.ResourceOption) (*VirtualRouterPeering, error) {
	var resource VirtualRouterPeering
	err := ctx.ReadResource("azurerm:network/v20190701:VirtualRouterPeering", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualRouterPeering resources.
type virtualRouterPeeringState struct {
	// Gets a unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// Gets name of the peering unique to VirtualRouter. This name can be used to access the resource.
	Name *string `pulumi:"name"`
	// The properties of the Virtual Router Peering.
	Properties *VirtualRouterPeeringPropertiesResponse `pulumi:"properties"`
	// Peering type.
	Type *string `pulumi:"type"`
}

type VirtualRouterPeeringState struct {
	// Gets a unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrInput
	// Gets name of the peering unique to VirtualRouter. This name can be used to access the resource.
	Name pulumi.StringPtrInput
	// The properties of the Virtual Router Peering.
	Properties VirtualRouterPeeringPropertiesResponsePtrInput
	// Peering type.
	Type pulumi.StringPtrInput
}

func (VirtualRouterPeeringState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualRouterPeeringState)(nil)).Elem()
}

type virtualRouterPeeringArgs struct {
	// Resource ID.
	Id *string `pulumi:"id"`
	// The name of the Virtual Router Peering.
	Name string `pulumi:"name"`
	// Peer ASN.
	PeerAsn *int `pulumi:"peerAsn"`
	// Peer IP.
	PeerIp *string `pulumi:"peerIp"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Virtual Router.
	VirtualRouterName string `pulumi:"virtualRouterName"`
}

// The set of arguments for constructing a VirtualRouterPeering resource.
type VirtualRouterPeeringArgs struct {
	// Resource ID.
	Id pulumi.StringPtrInput
	// The name of the Virtual Router Peering.
	Name pulumi.StringInput
	// Peer ASN.
	PeerAsn pulumi.IntPtrInput
	// Peer IP.
	PeerIp pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The name of the Virtual Router.
	VirtualRouterName pulumi.StringInput
}

func (VirtualRouterPeeringArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualRouterPeeringArgs)(nil)).Elem()
}
