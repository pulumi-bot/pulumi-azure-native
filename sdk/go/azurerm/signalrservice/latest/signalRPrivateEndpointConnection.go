// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// A private endpoint connection to SignalR resource
type SignalRPrivateEndpointConnection struct {
	pulumi.CustomResourceState

	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Private endpoint associated with the private endpoint connection
	PrivateEndpoint PrivateEndpointResponsePtrOutput `pulumi:"privateEndpoint"`
	// Connection state
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateResponsePtrOutput `pulumi:"privateLinkServiceConnectionState"`
	// Provisioning state of the private endpoint connection
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The type of the resource - e.g. "Microsoft.SignalRService/SignalR"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewSignalRPrivateEndpointConnection registers a new resource with the given unique name, arguments, and options.
func NewSignalRPrivateEndpointConnection(ctx *pulumi.Context,
	name string, args *SignalRPrivateEndpointConnectionArgs, opts ...pulumi.ResourceOption) (*SignalRPrivateEndpointConnection, error) {
	if args == nil || args.PrivateEndpointConnectionName == nil {
		return nil, errors.New("missing required argument 'PrivateEndpointConnectionName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ResourceName == nil {
		return nil, errors.New("missing required argument 'ResourceName'")
	}
	if args == nil {
		args = &SignalRPrivateEndpointConnectionArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:signalrservice/v20200501:SignalRPrivateEndpointConnection"),
		},
	})
	opts = append(opts, aliases)
	var resource SignalRPrivateEndpointConnection
	err := ctx.RegisterResource("azurerm:signalrservice/latest:SignalRPrivateEndpointConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSignalRPrivateEndpointConnection gets an existing SignalRPrivateEndpointConnection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSignalRPrivateEndpointConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SignalRPrivateEndpointConnectionState, opts ...pulumi.ResourceOption) (*SignalRPrivateEndpointConnection, error) {
	var resource SignalRPrivateEndpointConnection
	err := ctx.ReadResource("azurerm:signalrservice/latest:SignalRPrivateEndpointConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SignalRPrivateEndpointConnection resources.
type signalRPrivateEndpointConnectionState struct {
	// The name of the resource.
	Name *string `pulumi:"name"`
	// Private endpoint associated with the private endpoint connection
	PrivateEndpoint *PrivateEndpointResponse `pulumi:"privateEndpoint"`
	// Connection state
	PrivateLinkServiceConnectionState *PrivateLinkServiceConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
	// Provisioning state of the private endpoint connection
	ProvisioningState *string `pulumi:"provisioningState"`
	// The type of the resource - e.g. "Microsoft.SignalRService/SignalR"
	Type *string `pulumi:"type"`
}

type SignalRPrivateEndpointConnectionState struct {
	// The name of the resource.
	Name pulumi.StringPtrInput
	// Private endpoint associated with the private endpoint connection
	PrivateEndpoint PrivateEndpointResponsePtrInput
	// Connection state
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateResponsePtrInput
	// Provisioning state of the private endpoint connection
	ProvisioningState pulumi.StringPtrInput
	// The type of the resource - e.g. "Microsoft.SignalRService/SignalR"
	Type pulumi.StringPtrInput
}

func (SignalRPrivateEndpointConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*signalRPrivateEndpointConnectionState)(nil)).Elem()
}

type signalRPrivateEndpointConnectionArgs struct {
	// Private endpoint associated with the private endpoint connection
	PrivateEndpoint *PrivateEndpoint `pulumi:"privateEndpoint"`
	// The name of the private endpoint connection associated with the SignalR resource.
	PrivateEndpointConnectionName string `pulumi:"privateEndpointConnectionName"`
	// Connection state
	PrivateLinkServiceConnectionState *PrivateLinkServiceConnectionState `pulumi:"privateLinkServiceConnectionState"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the SignalR resource.
	ResourceName string `pulumi:"resourceName"`
}

// The set of arguments for constructing a SignalRPrivateEndpointConnection resource.
type SignalRPrivateEndpointConnectionArgs struct {
	// Private endpoint associated with the private endpoint connection
	PrivateEndpoint PrivateEndpointPtrInput
	// The name of the private endpoint connection associated with the SignalR resource.
	PrivateEndpointConnectionName pulumi.StringInput
	// Connection state
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStatePtrInput
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput
	// The name of the SignalR resource.
	ResourceName pulumi.StringInput
}

func (SignalRPrivateEndpointConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*signalRPrivateEndpointConnectionArgs)(nil)).Elem()
}