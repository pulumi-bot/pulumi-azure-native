// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Linked workspace.
type LinkedWorkspace struct {
	pulumi.CustomResourceState

	// Friendly name of the linked workspace.
	Name pulumi.StringOutput `pulumi:"name"`
	// LinkedWorkspace specific properties.
	Properties LinkedWorkspacePropsResponseOutput `pulumi:"properties"`
	// Resource type of linked workspace.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewLinkedWorkspace registers a new resource with the given unique name, arguments, and options.
func NewLinkedWorkspace(ctx *pulumi.Context,
	name string, args *LinkedWorkspaceArgs, opts ...pulumi.ResourceOption) (*LinkedWorkspace, error) {
	if args == nil || args.LinkName == nil {
		return nil, errors.New("missing required argument 'LinkName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.WorkspaceName == nil {
		return nil, errors.New("missing required argument 'WorkspaceName'")
	}
	if args == nil {
		args = &LinkedWorkspaceArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:machinelearningservices/v20200301:LinkedWorkspace"),
		},
	})
	opts = append(opts, aliases)
	var resource LinkedWorkspace
	err := ctx.RegisterResource("azure-nextgen:machinelearningservices/latest:LinkedWorkspace", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLinkedWorkspace gets an existing LinkedWorkspace resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLinkedWorkspace(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LinkedWorkspaceState, opts ...pulumi.ResourceOption) (*LinkedWorkspace, error) {
	var resource LinkedWorkspace
	err := ctx.ReadResource("azure-nextgen:machinelearningservices/latest:LinkedWorkspace", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LinkedWorkspace resources.
type linkedWorkspaceState struct {
	// Friendly name of the linked workspace.
	Name *string `pulumi:"name"`
	// LinkedWorkspace specific properties.
	Properties *LinkedWorkspacePropsResponse `pulumi:"properties"`
	// Resource type of linked workspace.
	Type *string `pulumi:"type"`
}

type LinkedWorkspaceState struct {
	// Friendly name of the linked workspace.
	Name pulumi.StringPtrInput
	// LinkedWorkspace specific properties.
	Properties LinkedWorkspacePropsResponsePtrInput
	// Resource type of linked workspace.
	Type pulumi.StringPtrInput
}

func (LinkedWorkspaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*linkedWorkspaceState)(nil)).Elem()
}

type linkedWorkspaceArgs struct {
	// Friendly name of the linked workspace
	LinkName string `pulumi:"linkName"`
	// Friendly name of the linked workspace
	Name *string `pulumi:"name"`
	// LinkedWorkspace specific properties.
	Properties *LinkedWorkspaceProps `pulumi:"properties"`
	// Name of the resource group in which workspace is located.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a LinkedWorkspace resource.
type LinkedWorkspaceArgs struct {
	// Friendly name of the linked workspace
	LinkName pulumi.StringInput
	// Friendly name of the linked workspace
	Name pulumi.StringPtrInput
	// LinkedWorkspace specific properties.
	Properties LinkedWorkspacePropsPtrInput
	// Name of the resource group in which workspace is located.
	ResourceGroupName pulumi.StringInput
	// Name of Azure Machine Learning workspace.
	WorkspaceName pulumi.StringInput
}

func (LinkedWorkspaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*linkedWorkspaceArgs)(nil)).Elem()
}

type LinkedWorkspaceInput interface {
	pulumi.Input

	ToLinkedWorkspaceOutput() LinkedWorkspaceOutput
	ToLinkedWorkspaceOutputWithContext(ctx context.Context) LinkedWorkspaceOutput
}

func (LinkedWorkspace) ElementType() reflect.Type {
	return reflect.TypeOf((*LinkedWorkspace)(nil)).Elem()
}

func (i LinkedWorkspace) ToLinkedWorkspaceOutput() LinkedWorkspaceOutput {
	return i.ToLinkedWorkspaceOutputWithContext(context.Background())
}

func (i LinkedWorkspace) ToLinkedWorkspaceOutputWithContext(ctx context.Context) LinkedWorkspaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinkedWorkspaceOutput)
}

type LinkedWorkspaceOutput struct {
	*pulumi.OutputState
}

func (LinkedWorkspaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LinkedWorkspaceOutput)(nil)).Elem()
}

func (o LinkedWorkspaceOutput) ToLinkedWorkspaceOutput() LinkedWorkspaceOutput {
	return o
}

func (o LinkedWorkspaceOutput) ToLinkedWorkspaceOutputWithContext(ctx context.Context) LinkedWorkspaceOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(LinkedWorkspaceOutput{})
}
