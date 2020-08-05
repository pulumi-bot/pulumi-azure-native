// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20180214

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Private endpoint connection resource.
type PrivateEndpointConnection struct {
	pulumi.CustomResourceState

	// Azure location of the key vault resource.
	Location pulumi.StringOutput `pulumi:"location"`
	// Name of the key vault resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Resource properties.
	Properties PrivateEndpointConnectionPropertiesResponseOutput `pulumi:"properties"`
	// Tags assigned to the key vault resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type of the key vault resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewPrivateEndpointConnection registers a new resource with the given unique name, arguments, and options.
func NewPrivateEndpointConnection(ctx *pulumi.Context,
	name string, args *PrivateEndpointConnectionArgs, opts ...pulumi.ResourceOption) (*PrivateEndpointConnection, error) {
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.VaultName == nil {
		return nil, errors.New("missing required argument 'VaultName'")
	}
	if args == nil {
		args = &PrivateEndpointConnectionArgs{}
	}
	var resource PrivateEndpointConnection
	err := ctx.RegisterResource("azurerm:keyvault/v20180214:PrivateEndpointConnection", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azurerm:keyvault/v20180214:PrivateEndpointConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PrivateEndpointConnection resources.
type privateEndpointConnectionState struct {
	// Azure location of the key vault resource.
	Location *string `pulumi:"location"`
	// Name of the key vault resource.
	Name *string `pulumi:"name"`
	// Resource properties.
	Properties *PrivateEndpointConnectionPropertiesResponse `pulumi:"properties"`
	// Tags assigned to the key vault resource.
	Tags map[string]string `pulumi:"tags"`
	// Resource type of the key vault resource.
	Type *string `pulumi:"type"`
}

type PrivateEndpointConnectionState struct {
	// Azure location of the key vault resource.
	Location pulumi.StringPtrInput
	// Name of the key vault resource.
	Name pulumi.StringPtrInput
	// Resource properties.
	Properties PrivateEndpointConnectionPropertiesResponsePtrInput
	// Tags assigned to the key vault resource.
	Tags pulumi.StringMapInput
	// Resource type of the key vault resource.
	Type pulumi.StringPtrInput
}

func (PrivateEndpointConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*privateEndpointConnectionState)(nil)).Elem()
}

type privateEndpointConnectionArgs struct {
	// Name of the private endpoint connection associated with the key vault.
	Name string `pulumi:"name"`
	// Properties of the private endpoint object.
	PrivateEndpoint *PrivateEndpoint `pulumi:"privateEndpoint"`
	// Approval state of the private link connection.
	PrivateLinkServiceConnectionState *PrivateLinkServiceConnectionState `pulumi:"privateLinkServiceConnectionState"`
	// Provisioning state of the private endpoint connection.
	ProvisioningState *string `pulumi:"provisioningState"`
	// Name of the resource group that contains the key vault.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the key vault.
	VaultName string `pulumi:"vaultName"`
}

// The set of arguments for constructing a PrivateEndpointConnection resource.
type PrivateEndpointConnectionArgs struct {
	// Name of the private endpoint connection associated with the key vault.
	Name pulumi.StringInput
	// Properties of the private endpoint object.
	PrivateEndpoint PrivateEndpointPtrInput
	// Approval state of the private link connection.
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStatePtrInput
	// Provisioning state of the private endpoint connection.
	ProvisioningState pulumi.StringPtrInput
	// Name of the resource group that contains the key vault.
	ResourceGroupName pulumi.StringInput
	// The name of the key vault.
	VaultName pulumi.StringInput
}

func (PrivateEndpointConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*privateEndpointConnectionArgs)(nil)).Elem()
}
