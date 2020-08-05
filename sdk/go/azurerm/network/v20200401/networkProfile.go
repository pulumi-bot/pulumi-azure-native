// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200401

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Network profile resource.
type NetworkProfile struct {
	pulumi.CustomResourceState

	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Network profile properties.
	Properties NetworkProfilePropertiesFormatResponseOutput `pulumi:"properties"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewNetworkProfile registers a new resource with the given unique name, arguments, and options.
func NewNetworkProfile(ctx *pulumi.Context,
	name string, args *NetworkProfileArgs, opts ...pulumi.ResourceOption) (*NetworkProfile, error) {
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &NetworkProfileArgs{}
	}
	var resource NetworkProfile
	err := ctx.RegisterResource("azurerm:network/v20200401:NetworkProfile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkProfile gets an existing NetworkProfile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkProfileState, opts ...pulumi.ResourceOption) (*NetworkProfile, error) {
	var resource NetworkProfile
	err := ctx.ReadResource("azurerm:network/v20200401:NetworkProfile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkProfile resources.
type networkProfileState struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Resource name.
	Name *string `pulumi:"name"`
	// Network profile properties.
	Properties *NetworkProfilePropertiesFormatResponse `pulumi:"properties"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type *string `pulumi:"type"`
}

type NetworkProfileState struct {
	// A unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// Resource name.
	Name pulumi.StringPtrInput
	// Network profile properties.
	Properties NetworkProfilePropertiesFormatResponsePtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Resource type.
	Type pulumi.StringPtrInput
}

func (NetworkProfileState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkProfileState)(nil)).Elem()
}

type networkProfileArgs struct {
	// List of chid container network interface configurations.
	ContainerNetworkInterfaceConfigurations []ContainerNetworkInterfaceConfiguration `pulumi:"containerNetworkInterfaceConfigurations"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The name of the network profile.
	Name string `pulumi:"name"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a NetworkProfile resource.
type NetworkProfileArgs struct {
	// List of chid container network interface configurations.
	ContainerNetworkInterfaceConfigurations ContainerNetworkInterfaceConfigurationArrayInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// The name of the network profile.
	Name pulumi.StringInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (NetworkProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkProfileArgs)(nil)).Elem()
}
