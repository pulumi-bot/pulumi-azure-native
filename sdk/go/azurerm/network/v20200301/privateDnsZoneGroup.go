// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200301

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Private dns zone group resource.
type PrivateDnsZoneGroup struct {
	pulumi.CustomResourceState

	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// Properties of the private dns zone group.
	Properties PrivateDnsZoneGroupPropertiesFormatResponseOutput `pulumi:"properties"`
}

// NewPrivateDnsZoneGroup registers a new resource with the given unique name, arguments, and options.
func NewPrivateDnsZoneGroup(ctx *pulumi.Context,
	name string, args *PrivateDnsZoneGroupArgs, opts ...pulumi.ResourceOption) (*PrivateDnsZoneGroup, error) {
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.PrivateEndpointName == nil {
		return nil, errors.New("missing required argument 'PrivateEndpointName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &PrivateDnsZoneGroupArgs{}
	}
	var resource PrivateDnsZoneGroup
	err := ctx.RegisterResource("azurerm:network/v20200301:PrivateDnsZoneGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPrivateDnsZoneGroup gets an existing PrivateDnsZoneGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPrivateDnsZoneGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateDnsZoneGroupState, opts ...pulumi.ResourceOption) (*PrivateDnsZoneGroup, error) {
	var resource PrivateDnsZoneGroup
	err := ctx.ReadResource("azurerm:network/v20200301:PrivateDnsZoneGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PrivateDnsZoneGroup resources.
type privateDnsZoneGroupState struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// Name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `pulumi:"name"`
	// Properties of the private dns zone group.
	Properties *PrivateDnsZoneGroupPropertiesFormatResponse `pulumi:"properties"`
}

type PrivateDnsZoneGroupState struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrInput
	// Name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name pulumi.StringPtrInput
	// Properties of the private dns zone group.
	Properties PrivateDnsZoneGroupPropertiesFormatResponsePtrInput
}

func (PrivateDnsZoneGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*privateDnsZoneGroupState)(nil)).Elem()
}

type privateDnsZoneGroupArgs struct {
	// Resource ID.
	Id *string `pulumi:"id"`
	// The name of the private dns zone group.
	Name string `pulumi:"name"`
	// A collection of private dns zone configurations of the private dns zone group.
	PrivateDnsZoneConfigs []PrivateDnsZoneConfig `pulumi:"privateDnsZoneConfigs"`
	// The name of the private endpoint.
	PrivateEndpointName string `pulumi:"privateEndpointName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a PrivateDnsZoneGroup resource.
type PrivateDnsZoneGroupArgs struct {
	// Resource ID.
	Id pulumi.StringPtrInput
	// The name of the private dns zone group.
	Name pulumi.StringInput
	// A collection of private dns zone configurations of the private dns zone group.
	PrivateDnsZoneConfigs PrivateDnsZoneConfigArrayInput
	// The name of the private endpoint.
	PrivateEndpointName pulumi.StringInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
}

func (PrivateDnsZoneGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*privateDnsZoneGroupArgs)(nil)).Elem()
}
