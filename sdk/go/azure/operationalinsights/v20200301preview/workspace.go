// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200301preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The top level Workspace resource container.
type Workspace struct {
	pulumi.CustomResourceState

	// This is a read-only property. Represents the ID associated with the workspace.
	CustomerId pulumi.StringOutput `pulumi:"customerId"`
	// The ETag of the workspace.
	ETag pulumi.StringPtrOutput `pulumi:"eTag"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// List of linked private link scope resources.
	PrivateLinkScopedResources PrivateLinkScopedResourceResponseArrayOutput `pulumi:"privateLinkScopedResources"`
	// The provisioning state of the workspace.
	ProvisioningState pulumi.StringPtrOutput `pulumi:"provisioningState"`
	// The network access type for accessing Log Analytics ingestion.
	PublicNetworkAccessForIngestion pulumi.StringPtrOutput `pulumi:"publicNetworkAccessForIngestion"`
	// The network access type for accessing Log Analytics query.
	PublicNetworkAccessForQuery pulumi.StringPtrOutput `pulumi:"publicNetworkAccessForQuery"`
	// The workspace data retention in days. -1 means Unlimited retention for the Unlimited Sku. 730 days is the maximum allowed for all other Skus.
	RetentionInDays pulumi.IntPtrOutput `pulumi:"retentionInDays"`
	// The SKU of the workspace.
	Sku WorkspaceSkuResponsePtrOutput `pulumi:"sku"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// The daily volume cap for ingestion.
	WorkspaceCapping WorkspaceCappingResponsePtrOutput `pulumi:"workspaceCapping"`
}

// NewWorkspace registers a new resource with the given unique name, arguments, and options.
func NewWorkspace(ctx *pulumi.Context,
	name string, args *WorkspaceArgs, opts ...pulumi.ResourceOption) (*Workspace, error) {
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.WorkspaceName == nil {
		return nil, errors.New("missing required argument 'WorkspaceName'")
	}
	if args == nil {
		args = &WorkspaceArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:operationalinsights/latest:Workspace"),
		},
		{
			Type: pulumi.String("azure-nextgen:operationalinsights/v20151101preview:Workspace"),
		},
		{
			Type: pulumi.String("azure-nextgen:operationalinsights/v20200801:Workspace"),
		},
	})
	opts = append(opts, aliases)
	var resource Workspace
	err := ctx.RegisterResource("azure-nextgen:operationalinsights/v20200301preview:Workspace", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkspace gets an existing Workspace resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkspace(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkspaceState, opts ...pulumi.ResourceOption) (*Workspace, error) {
	var resource Workspace
	err := ctx.ReadResource("azure-nextgen:operationalinsights/v20200301preview:Workspace", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Workspace resources.
type workspaceState struct {
	// This is a read-only property. Represents the ID associated with the workspace.
	CustomerId *string `pulumi:"customerId"`
	// The ETag of the workspace.
	ETag *string `pulumi:"eTag"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The name of the resource
	Name *string `pulumi:"name"`
	// List of linked private link scope resources.
	PrivateLinkScopedResources []PrivateLinkScopedResourceResponse `pulumi:"privateLinkScopedResources"`
	// The provisioning state of the workspace.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The network access type for accessing Log Analytics ingestion.
	PublicNetworkAccessForIngestion *string `pulumi:"publicNetworkAccessForIngestion"`
	// The network access type for accessing Log Analytics query.
	PublicNetworkAccessForQuery *string `pulumi:"publicNetworkAccessForQuery"`
	// The workspace data retention in days. -1 means Unlimited retention for the Unlimited Sku. 730 days is the maximum allowed for all other Skus.
	RetentionInDays *int `pulumi:"retentionInDays"`
	// The SKU of the workspace.
	Sku *WorkspaceSkuResponse `pulumi:"sku"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `pulumi:"type"`
	// The daily volume cap for ingestion.
	WorkspaceCapping *WorkspaceCappingResponse `pulumi:"workspaceCapping"`
}

type WorkspaceState struct {
	// This is a read-only property. Represents the ID associated with the workspace.
	CustomerId pulumi.StringPtrInput
	// The ETag of the workspace.
	ETag pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The name of the resource
	Name pulumi.StringPtrInput
	// List of linked private link scope resources.
	PrivateLinkScopedResources PrivateLinkScopedResourceResponseArrayInput
	// The provisioning state of the workspace.
	ProvisioningState pulumi.StringPtrInput
	// The network access type for accessing Log Analytics ingestion.
	PublicNetworkAccessForIngestion pulumi.StringPtrInput
	// The network access type for accessing Log Analytics query.
	PublicNetworkAccessForQuery pulumi.StringPtrInput
	// The workspace data retention in days. -1 means Unlimited retention for the Unlimited Sku. 730 days is the maximum allowed for all other Skus.
	RetentionInDays pulumi.IntPtrInput
	// The SKU of the workspace.
	Sku WorkspaceSkuResponsePtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringPtrInput
	// The daily volume cap for ingestion.
	WorkspaceCapping WorkspaceCappingResponsePtrInput
}

func (WorkspaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*workspaceState)(nil)).Elem()
}

type workspaceArgs struct {
	// The ETag of the workspace.
	ETag *string `pulumi:"eTag"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The provisioning state of the workspace.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The network access type for accessing Log Analytics ingestion.
	PublicNetworkAccessForIngestion *string `pulumi:"publicNetworkAccessForIngestion"`
	// The network access type for accessing Log Analytics query.
	PublicNetworkAccessForQuery *string `pulumi:"publicNetworkAccessForQuery"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The workspace data retention in days. -1 means Unlimited retention for the Unlimited Sku. 730 days is the maximum allowed for all other Skus.
	RetentionInDays *int `pulumi:"retentionInDays"`
	// The SKU of the workspace.
	Sku *WorkspaceSku `pulumi:"sku"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The daily volume cap for ingestion.
	WorkspaceCapping *WorkspaceCapping `pulumi:"workspaceCapping"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a Workspace resource.
type WorkspaceArgs struct {
	// The ETag of the workspace.
	ETag pulumi.StringPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringInput
	// The provisioning state of the workspace.
	ProvisioningState pulumi.StringPtrInput
	// The network access type for accessing Log Analytics ingestion.
	PublicNetworkAccessForIngestion pulumi.StringPtrInput
	// The network access type for accessing Log Analytics query.
	PublicNetworkAccessForQuery pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The workspace data retention in days. -1 means Unlimited retention for the Unlimited Sku. 730 days is the maximum allowed for all other Skus.
	RetentionInDays pulumi.IntPtrInput
	// The SKU of the workspace.
	Sku WorkspaceSkuPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The daily volume cap for ingestion.
	WorkspaceCapping WorkspaceCappingPtrInput
	// The name of the workspace.
	WorkspaceName pulumi.StringInput
}

func (WorkspaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workspaceArgs)(nil)).Elem()
}

type WorkspaceInput interface {
	pulumi.Input

	ToWorkspaceOutput() WorkspaceOutput
	ToWorkspaceOutputWithContext(ctx context.Context) WorkspaceOutput
}

func (Workspace) ElementType() reflect.Type {
	return reflect.TypeOf((*Workspace)(nil)).Elem()
}

func (i Workspace) ToWorkspaceOutput() WorkspaceOutput {
	return i.ToWorkspaceOutputWithContext(context.Background())
}

func (i Workspace) ToWorkspaceOutputWithContext(ctx context.Context) WorkspaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceOutput)
}

type WorkspaceOutput struct {
	*pulumi.OutputState
}

func (WorkspaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkspaceOutput)(nil)).Elem()
}

func (o WorkspaceOutput) ToWorkspaceOutput() WorkspaceOutput {
	return o
}

func (o WorkspaceOutput) ToWorkspaceOutputWithContext(ctx context.Context) WorkspaceOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(WorkspaceOutput{})
}
