// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200301

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A common class for general resource information.
type LocalNetworkGateway struct {
	pulumi.CustomResourceState

	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Properties of the local network gateway.
	Properties LocalNetworkGatewayPropertiesFormatResponseOutput `pulumi:"properties"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewLocalNetworkGateway registers a new resource with the given unique name, arguments, and options.
func NewLocalNetworkGateway(ctx *pulumi.Context,
	name string, args *LocalNetworkGatewayArgs, opts ...pulumi.ResourceOption) (*LocalNetworkGateway, error) {
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &LocalNetworkGatewayArgs{}
	}
	var resource LocalNetworkGateway
	err := ctx.RegisterResource("azurerm:network/v20200301:LocalNetworkGateway", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLocalNetworkGateway gets an existing LocalNetworkGateway resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLocalNetworkGateway(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LocalNetworkGatewayState, opts ...pulumi.ResourceOption) (*LocalNetworkGateway, error) {
	var resource LocalNetworkGateway
	err := ctx.ReadResource("azurerm:network/v20200301:LocalNetworkGateway", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LocalNetworkGateway resources.
type localNetworkGatewayState struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Resource name.
	Name *string `pulumi:"name"`
	// Properties of the local network gateway.
	Properties *LocalNetworkGatewayPropertiesFormatResponse `pulumi:"properties"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type *string `pulumi:"type"`
}

type LocalNetworkGatewayState struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// Resource name.
	Name pulumi.StringPtrInput
	// Properties of the local network gateway.
	Properties LocalNetworkGatewayPropertiesFormatResponsePtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Resource type.
	Type pulumi.StringPtrInput
}

func (LocalNetworkGatewayState) ElementType() reflect.Type {
	return reflect.TypeOf((*localNetworkGatewayState)(nil)).Elem()
}

type localNetworkGatewayArgs struct {
	// Local network gateway's BGP speaker settings.
	BgpSettings *BgpSettings `pulumi:"bgpSettings"`
	// FQDN of local network gateway.
	Fqdn *string `pulumi:"fqdn"`
	// IP address of local network gateway.
	GatewayIpAddress *string `pulumi:"gatewayIpAddress"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// Local network site address space.
	LocalNetworkAddressSpace *AddressSpace `pulumi:"localNetworkAddressSpace"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The name of the local network gateway.
	Name string `pulumi:"name"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a LocalNetworkGateway resource.
type LocalNetworkGatewayArgs struct {
	// Local network gateway's BGP speaker settings.
	BgpSettings BgpSettingsPtrInput
	// FQDN of local network gateway.
	Fqdn pulumi.StringPtrInput
	// IP address of local network gateway.
	GatewayIpAddress pulumi.StringPtrInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// Local network site address space.
	LocalNetworkAddressSpace AddressSpacePtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// The name of the local network gateway.
	Name pulumi.StringInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (LocalNetworkGatewayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*localNetworkGatewayArgs)(nil)).Elem()
}
