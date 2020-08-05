// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200601

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type PrivateEndpointConnection struct {
	pulumi.CustomResourceState

	// Name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Properties of the PrivateEndpointConnection.
	Properties PrivateEndpointConnectionPropertiesResponseOutput `pulumi:"properties"`
	// Type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewPrivateEndpointConnection registers a new resource with the given unique name, arguments, and options.
func NewPrivateEndpointConnection(ctx *pulumi.Context,
	name string, args *PrivateEndpointConnectionArgs, opts ...pulumi.ResourceOption) (*PrivateEndpointConnection, error) {
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ParentName == nil {
		return nil, errors.New("missing required argument 'ParentName'")
	}
	if args == nil || args.ParentType == nil {
		return nil, errors.New("missing required argument 'ParentType'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &PrivateEndpointConnectionArgs{}
	}
	var resource PrivateEndpointConnection
	err := ctx.RegisterResource("azurerm:eventgrid/v20200601:PrivateEndpointConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetPrivateEndpointConnection gets an existing PrivateEndpointConnection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetPrivateEndpointConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrivateEndpointConnectionState, opts ...pulumi.ResourceOption) (*PrivateEndpointConnection, error) {
	var resource PrivateEndpointConnection
	err := ctx.ReadResource("azurerm:eventgrid/v20200601:PrivateEndpointConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PrivateEndpointConnection resources.
type privateEndpointConnectionState struct {
	// Name of the resource.
	Name *string `pulumi:"name"`
	// Properties of the PrivateEndpointConnection.
	Properties *PrivateEndpointConnectionPropertiesResponse `pulumi:"properties"`
	// Type of the resource.
	Type *string `pulumi:"type"`
}

type PrivateEndpointConnectionState struct {
	// Name of the resource.
	Name pulumi.StringPtrInput
	// Properties of the PrivateEndpointConnection.
	Properties PrivateEndpointConnectionPropertiesResponsePtrInput
	// Type of the resource.
	Type pulumi.StringPtrInput
}

func (PrivateEndpointConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*privateEndpointConnectionState)(nil)).Elem()
}

type privateEndpointConnectionArgs struct {
	// GroupIds from the private link service resource.
	GroupIds []string `pulumi:"groupIds"`
	// The name of the private endpoint connection connection.
	Name string `pulumi:"name"`
	// The name of the parent resource (namely, either, the topic name or domain name).
	ParentName string `pulumi:"parentName"`
	// The type of the parent resource. This can be either \'topics\' or \'domains\'.
	ParentType string `pulumi:"parentType"`
	// The Private Endpoint resource for this Connection.
	PrivateEndpoint *PrivateEndpoint `pulumi:"privateEndpoint"`
	// Details about the state of the connection.
	PrivateLinkServiceConnectionState *ConnectionState `pulumi:"privateLinkServiceConnectionState"`
	// Provisioning state of the Private Endpoint Connection.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The name of the resource group within the user's subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The set of arguments for constructing a PrivateEndpointConnection resource.
type PrivateEndpointConnectionArgs struct {
	// GroupIds from the private link service resource.
	GroupIds pulumi.StringArrayInput
	// The name of the private endpoint connection connection.
	Name pulumi.StringInput
	// The name of the parent resource (namely, either, the topic name or domain name).
	ParentName pulumi.StringInput
	// The type of the parent resource. This can be either \'topics\' or \'domains\'.
	ParentType pulumi.StringInput
	// The Private Endpoint resource for this Connection.
	PrivateEndpoint PrivateEndpointPtrInput
	// Details about the state of the connection.
	PrivateLinkServiceConnectionState ConnectionStatePtrInput
	// Provisioning state of the Private Endpoint Connection.
	ProvisioningState pulumi.StringPtrInput
	// The name of the resource group within the user's subscription.
	ResourceGroupName pulumi.StringInput
}

func (PrivateEndpointConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*privateEndpointConnectionArgs)(nil)).Elem()
}
