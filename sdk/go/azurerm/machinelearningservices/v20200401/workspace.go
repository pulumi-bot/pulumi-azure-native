// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200401

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// An object that represents a machine learning workspace.
type Workspace struct {
	pulumi.CustomResourceState

	// The identity of the resource.
	Identity IdentityResponsePtrOutput `pulumi:"identity"`
	// Specifies the location of the resource.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Specifies the name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The properties of the machine learning workspace.
	Properties WorkspacePropertiesResponseOutput `pulumi:"properties"`
	// The sku of the workspace.
	Sku SkuResponsePtrOutput `pulumi:"sku"`
	// Contains resource tags defined as key/value pairs.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Specifies the type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewWorkspace registers a new resource with the given unique name, arguments, and options.
func NewWorkspace(ctx *pulumi.Context,
	name string, args *WorkspaceArgs, opts ...pulumi.ResourceOption) (*Workspace, error) {
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &WorkspaceArgs{}
	}
	var resource Workspace
	err := ctx.RegisterResource("azurerm:machinelearningservices/v20200401:Workspace", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azurerm:machinelearningservices/v20200401:Workspace", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Workspace resources.
type workspaceState struct {
	// The identity of the resource.
	Identity *IdentityResponse `pulumi:"identity"`
	// Specifies the location of the resource.
	Location *string `pulumi:"location"`
	// Specifies the name of the resource.
	Name *string `pulumi:"name"`
	// The properties of the machine learning workspace.
	Properties *WorkspacePropertiesResponse `pulumi:"properties"`
	// The sku of the workspace.
	Sku *SkuResponse `pulumi:"sku"`
	// Contains resource tags defined as key/value pairs.
	Tags map[string]string `pulumi:"tags"`
	// Specifies the type of the resource.
	Type *string `pulumi:"type"`
}

type WorkspaceState struct {
	// The identity of the resource.
	Identity IdentityResponsePtrInput
	// Specifies the location of the resource.
	Location pulumi.StringPtrInput
	// Specifies the name of the resource.
	Name pulumi.StringPtrInput
	// The properties of the machine learning workspace.
	Properties WorkspacePropertiesResponsePtrInput
	// The sku of the workspace.
	Sku SkuResponsePtrInput
	// Contains resource tags defined as key/value pairs.
	Tags pulumi.StringMapInput
	// Specifies the type of the resource.
	Type pulumi.StringPtrInput
}

func (WorkspaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*workspaceState)(nil)).Elem()
}

type workspaceArgs struct {
	// The flag to indicate whether to allow public access when behind VNet.
	AllowPublicAccessWhenBehindVnet *bool `pulumi:"allowPublicAccessWhenBehindVnet"`
	// ARM id of the application insights associated with this workspace. This cannot be changed once the workspace has been created
	ApplicationInsights *string `pulumi:"applicationInsights"`
	// ARM id of the container registry associated with this workspace. This cannot be changed once the workspace has been created
	ContainerRegistry *string `pulumi:"containerRegistry"`
	// The description of this workspace.
	Description *string `pulumi:"description"`
	// Url for the discovery service to identify regional endpoints for machine learning experimentation services
	DiscoveryUrl *string `pulumi:"discoveryUrl"`
	// The encryption settings of Azure ML workspace.
	Encryption *EncryptionProperty `pulumi:"encryption"`
	// The friendly name for this workspace. This name in mutable
	FriendlyName *string `pulumi:"friendlyName"`
	// The flag to signal HBI data in the workspace and reduce diagnostic data collected by the service
	HbiWorkspace *bool `pulumi:"hbiWorkspace"`
	// The identity of the resource.
	Identity *Identity `pulumi:"identity"`
	// The compute name for image build
	ImageBuildCompute *string `pulumi:"imageBuildCompute"`
	// ARM id of the key vault associated with this workspace. This cannot be changed once the workspace has been created
	KeyVault *string `pulumi:"keyVault"`
	// Specifies the location of the resource.
	Location *string `pulumi:"location"`
	// Name of Azure Machine Learning workspace.
	Name string `pulumi:"name"`
	// Name of the resource group in which workspace is located.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The list of shared private link resources in this workspace.
	SharedPrivateLinkResources []SharedPrivateLinkResource `pulumi:"sharedPrivateLinkResources"`
	// The sku of the workspace.
	Sku *Sku `pulumi:"sku"`
	// ARM id of the storage account associated with this workspace. This cannot be changed once the workspace has been created
	StorageAccount *string `pulumi:"storageAccount"`
	// Contains resource tags defined as key/value pairs.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Workspace resource.
type WorkspaceArgs struct {
	// The flag to indicate whether to allow public access when behind VNet.
	AllowPublicAccessWhenBehindVnet pulumi.BoolPtrInput
	// ARM id of the application insights associated with this workspace. This cannot be changed once the workspace has been created
	ApplicationInsights pulumi.StringPtrInput
	// ARM id of the container registry associated with this workspace. This cannot be changed once the workspace has been created
	ContainerRegistry pulumi.StringPtrInput
	// The description of this workspace.
	Description pulumi.StringPtrInput
	// Url for the discovery service to identify regional endpoints for machine learning experimentation services
	DiscoveryUrl pulumi.StringPtrInput
	// The encryption settings of Azure ML workspace.
	Encryption EncryptionPropertyPtrInput
	// The friendly name for this workspace. This name in mutable
	FriendlyName pulumi.StringPtrInput
	// The flag to signal HBI data in the workspace and reduce diagnostic data collected by the service
	HbiWorkspace pulumi.BoolPtrInput
	// The identity of the resource.
	Identity IdentityPtrInput
	// The compute name for image build
	ImageBuildCompute pulumi.StringPtrInput
	// ARM id of the key vault associated with this workspace. This cannot be changed once the workspace has been created
	KeyVault pulumi.StringPtrInput
	// Specifies the location of the resource.
	Location pulumi.StringPtrInput
	// Name of Azure Machine Learning workspace.
	Name pulumi.StringInput
	// Name of the resource group in which workspace is located.
	ResourceGroupName pulumi.StringInput
	// The list of shared private link resources in this workspace.
	SharedPrivateLinkResources SharedPrivateLinkResourceArrayInput
	// The sku of the workspace.
	Sku SkuPtrInput
	// ARM id of the storage account associated with this workspace. This cannot be changed once the workspace has been created
	StorageAccount pulumi.StringPtrInput
	// Contains resource tags defined as key/value pairs.
	Tags pulumi.StringMapInput
}

func (WorkspaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workspaceArgs)(nil)).Elem()
}
