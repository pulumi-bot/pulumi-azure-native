// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190401

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Nat Gateway resource.
type NatGateway struct {
	pulumi.CustomResourceState

	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// Resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Nat Gateway properties.
	Properties NatGatewayPropertiesFormatResponseOutput `pulumi:"properties"`
	// The nat gateway SKU.
	Sku NatGatewaySkuResponsePtrOutput `pulumi:"sku"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
	// A list of availability zones denoting the zone in which Nat Gateway should be deployed.
	Zones pulumi.StringArrayOutput `pulumi:"zones"`
}

// NewNatGateway registers a new resource with the given unique name, arguments, and options.
func NewNatGateway(ctx *pulumi.Context,
	name string, args *NatGatewayArgs, opts ...pulumi.ResourceOption) (*NatGateway, error) {
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &NatGatewayArgs{}
	}
	var resource NatGateway
	err := ctx.RegisterResource("azurerm:network/v20190401:NatGateway", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNatGateway gets an existing NatGateway resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNatGateway(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NatGatewayState, opts ...pulumi.ResourceOption) (*NatGateway, error) {
	var resource NatGateway
	err := ctx.ReadResource("azurerm:network/v20190401:NatGateway", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NatGateway resources.
type natGatewayState struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Resource name.
	Name *string `pulumi:"name"`
	// Nat Gateway properties.
	Properties *NatGatewayPropertiesFormatResponse `pulumi:"properties"`
	// The nat gateway SKU.
	Sku *NatGatewaySkuResponse `pulumi:"sku"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type *string `pulumi:"type"`
	// A list of availability zones denoting the zone in which Nat Gateway should be deployed.
	Zones []string `pulumi:"zones"`
}

type NatGatewayState struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// Resource name.
	Name pulumi.StringPtrInput
	// Nat Gateway properties.
	Properties NatGatewayPropertiesFormatResponsePtrInput
	// The nat gateway SKU.
	Sku NatGatewaySkuResponsePtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Resource type.
	Type pulumi.StringPtrInput
	// A list of availability zones denoting the zone in which Nat Gateway should be deployed.
	Zones pulumi.StringArrayInput
}

func (NatGatewayState) ElementType() reflect.Type {
	return reflect.TypeOf((*natGatewayState)(nil)).Elem()
}

type natGatewayArgs struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// The idle timeout of the nat gateway.
	IdleTimeoutInMinutes *int `pulumi:"idleTimeoutInMinutes"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The name of the nat gateway.
	Name string `pulumi:"name"`
	// The provisioning state of the NatGateway resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
	ProvisioningState *string `pulumi:"provisioningState"`
	// An array of public ip addresses associated with the nat gateway resource.
	PublicIpAddresses []SubResource `pulumi:"publicIpAddresses"`
	// An array of public ip prefixes associated with the nat gateway resource.
	PublicIpPrefixes []SubResource `pulumi:"publicIpPrefixes"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The resource GUID property of the nat gateway resource.
	ResourceGuid *string `pulumi:"resourceGuid"`
	// The nat gateway SKU.
	Sku *NatGatewaySku `pulumi:"sku"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// A list of availability zones denoting the zone in which Nat Gateway should be deployed.
	Zones []string `pulumi:"zones"`
}

// The set of arguments for constructing a NatGateway resource.
type NatGatewayArgs struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// The idle timeout of the nat gateway.
	IdleTimeoutInMinutes pulumi.IntPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// The name of the nat gateway.
	Name pulumi.StringInput
	// The provisioning state of the NatGateway resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
	ProvisioningState pulumi.StringPtrInput
	// An array of public ip addresses associated with the nat gateway resource.
	PublicIpAddresses SubResourceArrayInput
	// An array of public ip prefixes associated with the nat gateway resource.
	PublicIpPrefixes SubResourceArrayInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The resource GUID property of the nat gateway resource.
	ResourceGuid pulumi.StringPtrInput
	// The nat gateway SKU.
	Sku NatGatewaySkuPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// A list of availability zones denoting the zone in which Nat Gateway should be deployed.
	Zones pulumi.StringArrayInput
}

func (NatGatewayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*natGatewayArgs)(nil)).Elem()
}
