// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200901preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The private endpoint connection of a provisioning service
type IotDpsResourcePrivateEndpointConnection struct {
	pulumi.CustomResourceState

	// The resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// The properties of a private endpoint connection
	Properties PrivateEndpointConnectionPropertiesResponseOutput `pulumi:"properties"`
	// The resource type.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewIotDpsResourcePrivateEndpointConnection registers a new resource with the given unique name, arguments, and options.
func NewIotDpsResourcePrivateEndpointConnection(ctx *pulumi.Context,
	name string, args *IotDpsResourcePrivateEndpointConnectionArgs, opts ...pulumi.ResourceOption) (*IotDpsResourcePrivateEndpointConnection, error) {
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
		args = &IotDpsResourcePrivateEndpointConnectionArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:devices/latest:IotDpsResourcePrivateEndpointConnection"),
		},
		{
			Type: pulumi.String("azure-nextgen:devices/v20200301:IotDpsResourcePrivateEndpointConnection"),
		},
	})
	opts = append(opts, aliases)
	var resource IotDpsResourcePrivateEndpointConnection
	err := ctx.RegisterResource("azure-nextgen:devices/v20200901preview:IotDpsResourcePrivateEndpointConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIotDpsResourcePrivateEndpointConnection gets an existing IotDpsResourcePrivateEndpointConnection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIotDpsResourcePrivateEndpointConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IotDpsResourcePrivateEndpointConnectionState, opts ...pulumi.ResourceOption) (*IotDpsResourcePrivateEndpointConnection, error) {
	var resource IotDpsResourcePrivateEndpointConnection
	err := ctx.ReadResource("azure-nextgen:devices/v20200901preview:IotDpsResourcePrivateEndpointConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IotDpsResourcePrivateEndpointConnection resources.
type iotDpsResourcePrivateEndpointConnectionState struct {
	// The resource name.
	Name *string `pulumi:"name"`
	// The properties of a private endpoint connection
	Properties *PrivateEndpointConnectionPropertiesResponse `pulumi:"properties"`
	// The resource type.
	Type *string `pulumi:"type"`
}

type IotDpsResourcePrivateEndpointConnectionState struct {
	// The resource name.
	Name pulumi.StringPtrInput
	// The properties of a private endpoint connection
	Properties PrivateEndpointConnectionPropertiesResponsePtrInput
	// The resource type.
	Type pulumi.StringPtrInput
}

func (IotDpsResourcePrivateEndpointConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*iotDpsResourcePrivateEndpointConnectionState)(nil)).Elem()
}

type iotDpsResourcePrivateEndpointConnectionArgs struct {
	// The name of the private endpoint connection
	PrivateEndpointConnectionName string `pulumi:"privateEndpointConnectionName"`
	// The properties of a private endpoint connection
	Properties PrivateEndpointConnectionProperties `pulumi:"properties"`
	// The name of the resource group that contains the provisioning service.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the provisioning service.
	ResourceName string `pulumi:"resourceName"`
}

// The set of arguments for constructing a IotDpsResourcePrivateEndpointConnection resource.
type IotDpsResourcePrivateEndpointConnectionArgs struct {
	// The name of the private endpoint connection
	PrivateEndpointConnectionName pulumi.StringInput
	// The properties of a private endpoint connection
	Properties PrivateEndpointConnectionPropertiesInput
	// The name of the resource group that contains the provisioning service.
	ResourceGroupName pulumi.StringInput
	// The name of the provisioning service.
	ResourceName pulumi.StringInput
}

func (IotDpsResourcePrivateEndpointConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iotDpsResourcePrivateEndpointConnectionArgs)(nil)).Elem()
}

type IotDpsResourcePrivateEndpointConnectionInput interface {
	pulumi.Input

	ToIotDpsResourcePrivateEndpointConnectionOutput() IotDpsResourcePrivateEndpointConnectionOutput
	ToIotDpsResourcePrivateEndpointConnectionOutputWithContext(ctx context.Context) IotDpsResourcePrivateEndpointConnectionOutput
}

func (IotDpsResourcePrivateEndpointConnection) ElementType() reflect.Type {
	return reflect.TypeOf((*IotDpsResourcePrivateEndpointConnection)(nil)).Elem()
}

func (i IotDpsResourcePrivateEndpointConnection) ToIotDpsResourcePrivateEndpointConnectionOutput() IotDpsResourcePrivateEndpointConnectionOutput {
	return i.ToIotDpsResourcePrivateEndpointConnectionOutputWithContext(context.Background())
}

func (i IotDpsResourcePrivateEndpointConnection) ToIotDpsResourcePrivateEndpointConnectionOutputWithContext(ctx context.Context) IotDpsResourcePrivateEndpointConnectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IotDpsResourcePrivateEndpointConnectionOutput)
}

type IotDpsResourcePrivateEndpointConnectionOutput struct {
	*pulumi.OutputState
}

func (IotDpsResourcePrivateEndpointConnectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IotDpsResourcePrivateEndpointConnectionOutput)(nil)).Elem()
}

func (o IotDpsResourcePrivateEndpointConnectionOutput) ToIotDpsResourcePrivateEndpointConnectionOutput() IotDpsResourcePrivateEndpointConnectionOutput {
	return o
}

func (o IotDpsResourcePrivateEndpointConnectionOutput) ToIotDpsResourcePrivateEndpointConnectionOutputWithContext(ctx context.Context) IotDpsResourcePrivateEndpointConnectionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(IotDpsResourcePrivateEndpointConnectionOutput{})
}
