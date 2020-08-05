// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20181001

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// VpnGateway Resource.
type VpnGateway struct {
	pulumi.CustomResourceState

	// Gets a unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Resource location.
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Parameters for VpnGateway
	Properties VpnGatewayPropertiesResponseOutput `pulumi:"properties"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewVpnGateway registers a new resource with the given unique name, arguments, and options.
func NewVpnGateway(ctx *pulumi.Context,
	name string, args *VpnGatewayArgs, opts ...pulumi.ResourceOption) (*VpnGateway, error) {
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
		args = &VpnGatewayArgs{}
	}
	var resource VpnGateway
	err := ctx.RegisterResource("azurerm:network/v20181001:VpnGateway", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVpnGateway gets an existing VpnGateway resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVpnGateway(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VpnGatewayState, opts ...pulumi.ResourceOption) (*VpnGateway, error) {
	var resource VpnGateway
	err := ctx.ReadResource("azurerm:network/v20181001:VpnGateway", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VpnGateway resources.
type vpnGatewayState struct {
	// Gets a unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Resource name.
	Name *string `pulumi:"name"`
	// Parameters for VpnGateway
	Properties *VpnGatewayPropertiesResponse `pulumi:"properties"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type *string `pulumi:"type"`
}

type VpnGatewayState struct {
	// Gets a unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// Resource name.
	Name pulumi.StringPtrInput
	// Parameters for VpnGateway
	Properties VpnGatewayPropertiesResponsePtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Resource type.
	Type pulumi.StringPtrInput
}

func (VpnGatewayState) ElementType() reflect.Type {
	return reflect.TypeOf((*vpnGatewayState)(nil)).Elem()
}

type vpnGatewayArgs struct {
	// Local network gateway's BGP speaker settings.
	BgpSettings *BgpSettings `pulumi:"bgpSettings"`
	// list of all vpn connections to the gateway.
	Connections []VpnConnectionType `pulumi:"connections"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// Resource location.
	Location string `pulumi:"location"`
	// The name of the gateway.
	Name string `pulumi:"name"`
	// The provisioning state of the resource.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The resource group name of the VpnGateway.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The VirtualHub to which the gateway belongs
	VirtualHub *SubResource `pulumi:"virtualHub"`
	// The scale unit for this vpn gateway.
	VpnGatewayScaleUnit *int `pulumi:"vpnGatewayScaleUnit"`
}

// The set of arguments for constructing a VpnGateway resource.
type VpnGatewayArgs struct {
	// Local network gateway's BGP speaker settings.
	BgpSettings BgpSettingsPtrInput
	// list of all vpn connections to the gateway.
	Connections VpnConnectionTypeArrayInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringInput
	// The name of the gateway.
	Name pulumi.StringInput
	// The provisioning state of the resource.
	ProvisioningState pulumi.StringPtrInput
	// The resource group name of the VpnGateway.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The VirtualHub to which the gateway belongs
	VirtualHub SubResourcePtrInput
	// The scale unit for this vpn gateway.
	VpnGatewayScaleUnit pulumi.IntPtrInput
}

func (VpnGatewayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vpnGatewayArgs)(nil)).Elem()
}
