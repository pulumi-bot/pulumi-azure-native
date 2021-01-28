// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20190601preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Integration runtime resource type.
type IntegrationRuntime struct {
	pulumi.CustomResourceState

	// Resource Etag.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Integration runtime properties.
	Properties pulumi.AnyOutput `pulumi:"properties"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewIntegrationRuntime registers a new resource with the given unique name, arguments, and options.
func NewIntegrationRuntime(ctx *pulumi.Context,
	name string, args *IntegrationRuntimeArgs, opts ...pulumi.ResourceOption) (*IntegrationRuntime, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.IntegrationRuntimeName == nil {
		return nil, errors.New("invalid value for required argument 'IntegrationRuntimeName'")
	}
	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:synapse/v20201201:IntegrationRuntime"),
		},
	})
	opts = append(opts, aliases)
	var resource IntegrationRuntime
	err := ctx.RegisterResource("azure-nextgen:synapse/v20190601preview:IntegrationRuntime", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIntegrationRuntime gets an existing IntegrationRuntime resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIntegrationRuntime(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IntegrationRuntimeState, opts ...pulumi.ResourceOption) (*IntegrationRuntime, error) {
	var resource IntegrationRuntime
	err := ctx.ReadResource("azure-nextgen:synapse/v20190601preview:IntegrationRuntime", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IntegrationRuntime resources.
type integrationRuntimeState struct {
	// Resource Etag.
	Etag *string `pulumi:"etag"`
	// The name of the resource
	Name *string `pulumi:"name"`
	// Integration runtime properties.
	Properties interface{} `pulumi:"properties"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `pulumi:"type"`
}

type IntegrationRuntimeState struct {
	// Resource Etag.
	Etag pulumi.StringPtrInput
	// The name of the resource
	Name pulumi.StringPtrInput
	// Integration runtime properties.
	Properties pulumi.Input
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringPtrInput
}

func (IntegrationRuntimeState) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationRuntimeState)(nil)).Elem()
}

type integrationRuntimeArgs struct {
	// Integration runtime name
	IntegrationRuntimeName string `pulumi:"integrationRuntimeName"`
	// Integration runtime properties.
	Properties interface{} `pulumi:"properties"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a IntegrationRuntime resource.
type IntegrationRuntimeArgs struct {
	// Integration runtime name
	IntegrationRuntimeName pulumi.StringInput
	// Integration runtime properties.
	Properties pulumi.Input
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The name of the workspace.
	WorkspaceName pulumi.StringInput
}

func (IntegrationRuntimeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*integrationRuntimeArgs)(nil)).Elem()
}

type IntegrationRuntimeInput interface {
	pulumi.Input

	ToIntegrationRuntimeOutput() IntegrationRuntimeOutput
	ToIntegrationRuntimeOutputWithContext(ctx context.Context) IntegrationRuntimeOutput
}

func (*IntegrationRuntime) ElementType() reflect.Type {
	return reflect.TypeOf((*IntegrationRuntime)(nil))
}

func (i *IntegrationRuntime) ToIntegrationRuntimeOutput() IntegrationRuntimeOutput {
	return i.ToIntegrationRuntimeOutputWithContext(context.Background())
}

func (i *IntegrationRuntime) ToIntegrationRuntimeOutputWithContext(ctx context.Context) IntegrationRuntimeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IntegrationRuntimeOutput)
}

type IntegrationRuntimeOutput struct {
	*pulumi.OutputState
}

func (IntegrationRuntimeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IntegrationRuntime)(nil))
}

func (o IntegrationRuntimeOutput) ToIntegrationRuntimeOutput() IntegrationRuntimeOutput {
	return o
}

func (o IntegrationRuntimeOutput) ToIntegrationRuntimeOutputWithContext(ctx context.Context) IntegrationRuntimeOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(IntegrationRuntimeOutput{})
}
