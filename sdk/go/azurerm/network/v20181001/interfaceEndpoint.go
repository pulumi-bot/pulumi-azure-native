// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20181001

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Interface endpoint resource.
type InterfaceEndpoint struct {
	pulumi.CustomResourceState

	// Gets a unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	// Resource location.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Properties of the interface endpoint.
	Properties InterfaceEndpointPropertiesResponseOutput `pulumi:"properties"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewInterfaceEndpoint registers a new resource with the given unique name, arguments, and options.
func NewInterfaceEndpoint(ctx *pulumi.Context,
	name string, args *InterfaceEndpointArgs, opts ...pulumi.ResourceOption) (*InterfaceEndpoint, error) {
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &InterfaceEndpointArgs{}
	}
	var resource InterfaceEndpoint
	err := ctx.RegisterResource("azurerm:network/v20181001:InterfaceEndpoint", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInterfaceEndpoint gets an existing InterfaceEndpoint resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInterfaceEndpoint(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InterfaceEndpointState, opts ...pulumi.ResourceOption) (*InterfaceEndpoint, error) {
	var resource InterfaceEndpoint
	err := ctx.ReadResource("azurerm:network/v20181001:InterfaceEndpoint", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering InterfaceEndpoint resources.
type interfaceEndpointState struct {
	// Gets a unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// Resource location.
	Location *string `pulumi:"location"`
	// Resource name.
	Name *string `pulumi:"name"`
	// Properties of the interface endpoint.
	Properties *InterfaceEndpointPropertiesResponse `pulumi:"properties"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Resource type.
	Type *string `pulumi:"type"`
}

type InterfaceEndpointState struct {
	// Gets a unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// Resource name.
	Name pulumi.StringPtrInput
	// Properties of the interface endpoint.
	Properties InterfaceEndpointPropertiesResponsePtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Resource type.
	Type pulumi.StringPtrInput
}

func (InterfaceEndpointState) ElementType() reflect.Type {
	return reflect.TypeOf((*interfaceEndpointState)(nil)).Elem()
}

type interfaceEndpointArgs struct {
	// A reference to the service being brought into the virtual network.
	EndpointService *EndpointService `pulumi:"endpointService"`
	// Gets a unique read-only string that changes whenever the resource is updated.
	Etag *string `pulumi:"etag"`
	// A first-party service's FQDN that is mapped to the private IP allocated via this interface endpoint.
	Fqdn *string `pulumi:"fqdn"`
	// Resource ID.
	Id *string `pulumi:"id"`
	// Resource location.
	Location *string `pulumi:"location"`
	// The name of the interface endpoint.
	Name string `pulumi:"name"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The ID of the subnet from which the private IP will be allocated.
	Subnet *SubnetType `pulumi:"subnet"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a InterfaceEndpoint resource.
type InterfaceEndpointArgs struct {
	// A reference to the service being brought into the virtual network.
	EndpointService EndpointServicePtrInput
	// Gets a unique read-only string that changes whenever the resource is updated.
	Etag pulumi.StringPtrInput
	// A first-party service's FQDN that is mapped to the private IP allocated via this interface endpoint.
	Fqdn pulumi.StringPtrInput
	// Resource ID.
	Id pulumi.StringPtrInput
	// Resource location.
	Location pulumi.StringPtrInput
	// The name of the interface endpoint.
	Name pulumi.StringInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The ID of the subnet from which the private IP will be allocated.
	Subnet SubnetTypePtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
}

func (InterfaceEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*interfaceEndpointArgs)(nil)).Elem()
}
