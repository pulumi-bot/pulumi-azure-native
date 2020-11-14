// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package latest

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type WorkspaceCollection struct {
	pulumi.CustomResourceState

	// Azure location
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Workspace collection name
	Name pulumi.StringPtrOutput `pulumi:"name"`
	// Properties
	Properties pulumi.AnyOutput          `pulumi:"properties"`
	Sku        AzureSkuResponsePtrOutput `pulumi:"sku"`
	Tags       pulumi.StringMapOutput    `pulumi:"tags"`
	// Resource type
	Type pulumi.StringPtrOutput `pulumi:"type"`
}

// NewWorkspaceCollection registers a new resource with the given unique name, arguments, and options.
func NewWorkspaceCollection(ctx *pulumi.Context,
	name string, args *WorkspaceCollectionArgs, opts ...pulumi.ResourceOption) (*WorkspaceCollection, error) {
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.WorkspaceCollectionName == nil {
		return nil, errors.New("missing required argument 'WorkspaceCollectionName'")
	}
	if args == nil {
		args = &WorkspaceCollectionArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:powerbi/v20160129:WorkspaceCollection"),
		},
	})
	opts = append(opts, aliases)
	var resource WorkspaceCollection
	err := ctx.RegisterResource("azure-nextgen:powerbi/latest:WorkspaceCollection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkspaceCollection gets an existing WorkspaceCollection resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkspaceCollection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkspaceCollectionState, opts ...pulumi.ResourceOption) (*WorkspaceCollection, error) {
	var resource WorkspaceCollection
	err := ctx.ReadResource("azure-nextgen:powerbi/latest:WorkspaceCollection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WorkspaceCollection resources.
type workspaceCollectionState struct {
	// Azure location
	Location *string `pulumi:"location"`
	// Workspace collection name
	Name *string `pulumi:"name"`
	// Properties
	Properties interface{}       `pulumi:"properties"`
	Sku        *AzureSkuResponse `pulumi:"sku"`
	Tags       map[string]string `pulumi:"tags"`
	// Resource type
	Type *string `pulumi:"type"`
}

type WorkspaceCollectionState struct {
	// Azure location
	Location pulumi.StringPtrInput
	// Workspace collection name
	Name pulumi.StringPtrInput
	// Properties
	Properties pulumi.Input
	Sku        AzureSkuResponsePtrInput
	Tags       pulumi.StringMapInput
	// Resource type
	Type pulumi.StringPtrInput
}

func (WorkspaceCollectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*workspaceCollectionState)(nil)).Elem()
}

type workspaceCollectionArgs struct {
	// Azure location
	Location *string `pulumi:"location"`
	// Azure resource group
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	Sku               *AzureSku         `pulumi:"sku"`
	Tags              map[string]string `pulumi:"tags"`
	// Power BI Embedded Workspace Collection name
	WorkspaceCollectionName string `pulumi:"workspaceCollectionName"`
}

// The set of arguments for constructing a WorkspaceCollection resource.
type WorkspaceCollectionArgs struct {
	// Azure location
	Location pulumi.StringPtrInput
	// Azure resource group
	ResourceGroupName pulumi.StringInput
	Sku               AzureSkuPtrInput
	Tags              pulumi.StringMapInput
	// Power BI Embedded Workspace Collection name
	WorkspaceCollectionName pulumi.StringInput
}

func (WorkspaceCollectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workspaceCollectionArgs)(nil)).Elem()
}

type WorkspaceCollectionInput interface {
	pulumi.Input

	ToWorkspaceCollectionOutput() WorkspaceCollectionOutput
	ToWorkspaceCollectionOutputWithContext(ctx context.Context) WorkspaceCollectionOutput
}

func (WorkspaceCollection) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspaceCollection)(nil)).Elem()
}

func (i WorkspaceCollection) ToWorkspaceCollectionOutput() WorkspaceCollectionOutput {
	return i.ToWorkspaceCollectionOutputWithContext(context.Background())
}

func (i WorkspaceCollection) ToWorkspaceCollectionOutputWithContext(ctx context.Context) WorkspaceCollectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceCollectionOutput)
}

type WorkspaceCollectionOutput struct {
	*pulumi.OutputState
}

func (WorkspaceCollectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspaceCollectionOutput)(nil)).Elem()
}

func (o WorkspaceCollectionOutput) ToWorkspaceCollectionOutput() WorkspaceCollectionOutput {
	return o
}

func (o WorkspaceCollectionOutput) ToWorkspaceCollectionOutputWithContext(ctx context.Context) WorkspaceCollectionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(WorkspaceCollectionOutput{})
}
