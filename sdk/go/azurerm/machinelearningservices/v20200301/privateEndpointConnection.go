// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200301

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The Private Endpoint Connection resource.
type PrivateEndpointConnection struct {
	pulumi.CustomResourceState

	// The identity of the resource.
	Identity IdentityResponsePtrOutput `pulumi:"identity"`
	// Specifies the location of the resource.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Specifies the name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// Resource properties.
	Properties PrivateEndpointConnectionPropertiesResponseOutput `pulumi:"properties"`
	// The sku of the workspace.
	Sku SkuResponsePtrOutput `pulumi:"sku"`
	// Contains resource tags defined as key/value pairs.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Specifies the type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewPrivateEndpointConnection registers a new resource with the given unique name, arguments, and options.
func NewPrivateEndpointConnection(ctx *pulumi.Context,
	name string, args *PrivateEndpointConnectionArgs, opts ...pulumi.ResourceOption) (*PrivateEndpointConnection, error) {
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.PrivateLinkServiceConnectionState == nil {
		return nil, errors.New("missing required argument 'PrivateLinkServiceConnectionState'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.WorkspaceName == nil {
		return nil, errors.New("missing required argument 'WorkspaceName'")
	}
	if args == nil {
		args = &PrivateEndpointConnectionArgs{}
	}
	var resource PrivateEndpointConnection
	err := ctx.RegisterResource("azurerm:machinelearningservices/v20200301:PrivateEndpointConnection", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azurerm:machinelearningservices/v20200301:PrivateEndpointConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PrivateEndpointConnection resources.
type privateEndpointConnectionState struct {
	// The identity of the resource.
	Identity *IdentityResponse `pulumi:"identity"`
	// Specifies the location of the resource.
	Location *string `pulumi:"location"`
	// Specifies the name of the resource.
	Name *string `pulumi:"name"`
	// Resource properties.
	Properties *PrivateEndpointConnectionPropertiesResponse `pulumi:"properties"`
	// The sku of the workspace.
	Sku *SkuResponse `pulumi:"sku"`
	// Contains resource tags defined as key/value pairs.
	Tags map[string]string `pulumi:"tags"`
	// Specifies the type of the resource.
	Type *string `pulumi:"type"`
}

type PrivateEndpointConnectionState struct {
	// The identity of the resource.
	Identity IdentityResponsePtrInput
	// Specifies the location of the resource.
	Location pulumi.StringPtrInput
	// Specifies the name of the resource.
	Name pulumi.StringPtrInput
	// Resource properties.
	Properties PrivateEndpointConnectionPropertiesResponsePtrInput
	// The sku of the workspace.
	Sku SkuResponsePtrInput
	// Contains resource tags defined as key/value pairs.
	Tags pulumi.StringMapInput
	// Specifies the type of the resource.
	Type pulumi.StringPtrInput
}

func (PrivateEndpointConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*privateEndpointConnectionState)(nil)).Elem()
}

type privateEndpointConnectionArgs struct {
	// The identity of the resource.
	Identity *Identity `pulumi:"identity"`
	// Specifies the location of the resource.
	Location *string `pulumi:"location"`
	// The name of the private endpoint connection associated with the workspace
	Name string `pulumi:"name"`
	// The resource of private end point.
	PrivateEndpoint *PrivateEndpoint `pulumi:"privateEndpoint"`
	// A collection of information about the state of the connection between service consumer and provider.
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionState `pulumi:"privateLinkServiceConnectionState"`
	// The provisioning state of the private endpoint connection resource.
	ProvisioningState *string `pulumi:"provisioningState"`
	// Name of the resource group in which workspace is located.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The sku of the workspace.
	Sku *Sku `pulumi:"sku"`
	// Contains resource tags defined as key/value pairs.
	Tags map[string]string `pulumi:"tags"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a PrivateEndpointConnection resource.
type PrivateEndpointConnectionArgs struct {
	// The identity of the resource.
	Identity IdentityPtrInput
	// Specifies the location of the resource.
	Location pulumi.StringPtrInput
	// The name of the private endpoint connection associated with the workspace
	Name pulumi.StringInput
	// The resource of private end point.
	PrivateEndpoint PrivateEndpointPtrInput
	// A collection of information about the state of the connection between service consumer and provider.
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateInput
	// The provisioning state of the private endpoint connection resource.
	ProvisioningState pulumi.StringPtrInput
	// Name of the resource group in which workspace is located.
	ResourceGroupName pulumi.StringInput
	// The sku of the workspace.
	Sku SkuPtrInput
	// Contains resource tags defined as key/value pairs.
	Tags pulumi.StringMapInput
	// Name of Azure Machine Learning workspace.
	WorkspaceName pulumi.StringInput
}

func (PrivateEndpointConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*privateEndpointConnectionArgs)(nil)).Elem()
}
