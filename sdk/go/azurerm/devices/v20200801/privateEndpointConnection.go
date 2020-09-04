// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200801

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The private endpoint connection of an IotHub
type PrivateEndpointConnection struct {
	pulumi.CustomResourceState

	// The resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The properties of a private endpoint connection
	Properties PrivateEndpointConnectionPropertiesResponseOutput `pulumi:"properties"`
	// The resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewPrivateEndpointConnection registers a new resource with the given unique name, arguments, and options.
func NewPrivateEndpointConnection(ctx *pulumi.Context,
	name string, args *PrivateEndpointConnectionArgs, opts ...pulumi.ResourceOption) (*PrivateEndpointConnection, error) {
	if args == nil || args.PrivateEndpointConnectionName == nil {
		return nil, errors.New("missing required argument 'PrivateEndpointConnectionName'")
	}
	if args == nil || args.Properties == nil {
		return nil, errors.New("missing required argument 'Properties'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.ResourceName == nil {
		return nil, errors.New("missing required argument 'ResourceName'")
	}
	if args == nil {
		args = &PrivateEndpointConnectionArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:devices/latest:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azurerm:devices/v20200301:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azurerm:devices/v20200401:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azurerm:devices/v20200615:PrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azurerm:devices/v20200710preview:PrivateEndpointConnection"),
		},
	})
	opts = append(opts, aliases)
	var resource PrivateEndpointConnection
	err := ctx.RegisterResource("azurerm:devices/v20200801:PrivateEndpointConnection", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azurerm:devices/v20200801:PrivateEndpointConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering PrivateEndpointConnection resources.
type privateEndpointConnectionState struct {
	// The resource name.
	Name *string `pulumi:"name"`
	// The properties of a private endpoint connection
	Properties *PrivateEndpointConnectionPropertiesResponse `pulumi:"properties"`
	// The resource type.
	Type *string `pulumi:"type"`
}

type PrivateEndpointConnectionState struct {
	// The resource name.
	Name pulumi.StringPtrInput
	// The properties of a private endpoint connection
	Properties PrivateEndpointConnectionPropertiesResponsePtrInput
	// The resource type.
	Type pulumi.StringPtrInput
}

func (PrivateEndpointConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*privateEndpointConnectionState)(nil)).Elem()
}

type privateEndpointConnectionArgs struct {
	// The name of the private endpoint connection
	PrivateEndpointConnectionName string `pulumi:"privateEndpointConnectionName"`
	// The properties of a private endpoint connection
	Properties PrivateEndpointConnectionProperties `pulumi:"properties"`
	// The name of the resource group that contains the IoT hub.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the IoT hub.
	ResourceName string `pulumi:"resourceName"`
}

// The set of arguments for constructing a PrivateEndpointConnection resource.
type PrivateEndpointConnectionArgs struct {
	// The name of the private endpoint connection
	PrivateEndpointConnectionName pulumi.StringInput
	// The properties of a private endpoint connection
	Properties PrivateEndpointConnectionPropertiesInput
	// The name of the resource group that contains the IoT hub.
	ResourceGroupName pulumi.StringInput
	// The name of the IoT hub.
	ResourceName pulumi.StringInput
}

func (PrivateEndpointConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*privateEndpointConnectionArgs)(nil)).Elem()
}
