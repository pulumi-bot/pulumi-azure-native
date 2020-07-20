// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package machinelearningservices

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Workspace connection.
type WorkspaceConnection struct {
	pulumi.CustomResourceState

	// Friendly name of the workspace connection.
	Name pulumi.StringOutput `pulumi:"name"`
	// Properties of workspace connection.
	Properties WorkspaceConnectionPropsResponseOutput `pulumi:"properties"`
	// Resource type of workspace connection.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewWorkspaceConnection registers a new resource with the given unique name, arguments, and options.
func NewWorkspaceConnection(ctx *pulumi.Context,
	name string, args *WorkspaceConnectionArgs, opts ...pulumi.ResourceOption) (*WorkspaceConnection, error) {
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.WorkspaceName == nil {
		return nil, errors.New("missing required argument 'WorkspaceName'")
	}
	if args == nil {
		args = &WorkspaceConnectionArgs{}
	}
	var resource WorkspaceConnection
	err := ctx.RegisterResource("azurerm:machinelearningservices:WorkspaceConnection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkspaceConnection gets an existing WorkspaceConnection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkspaceConnection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkspaceConnectionState, opts ...pulumi.ResourceOption) (*WorkspaceConnection, error) {
	var resource WorkspaceConnection
	err := ctx.ReadResource("azurerm:machinelearningservices:WorkspaceConnection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WorkspaceConnection resources.
type workspaceConnectionState struct {
	// Friendly name of the workspace connection.
	Name *string `pulumi:"name"`
	// Properties of workspace connection.
	Properties *WorkspaceConnectionPropsResponse `pulumi:"properties"`
	// Resource type of workspace connection.
	Type *string `pulumi:"type"`
}

type WorkspaceConnectionState struct {
	// Friendly name of the workspace connection.
	Name pulumi.StringPtrInput
	// Properties of workspace connection.
	Properties WorkspaceConnectionPropsResponsePtrInput
	// Resource type of workspace connection.
	Type pulumi.StringPtrInput
}

func (WorkspaceConnectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*workspaceConnectionState)(nil)).Elem()
}

type workspaceConnectionArgs struct {
	// Friendly name of the workspace connection
	Name string `pulumi:"name"`
	// Properties of workspace connection.
	Properties *WorkspaceConnectionProps `pulumi:"properties"`
	// Name of the resource group in which workspace is located.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a WorkspaceConnection resource.
type WorkspaceConnectionArgs struct {
	// Friendly name of the workspace connection
	Name pulumi.StringInput
	// Properties of workspace connection.
	Properties WorkspaceConnectionPropsPtrInput
	// Name of the resource group in which workspace is located.
	ResourceGroupName pulumi.StringInput
	// Name of Azure Machine Learning workspace.
	WorkspaceName pulumi.StringInput
}

func (WorkspaceConnectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workspaceConnectionArgs)(nil)).Elem()
}
