// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200801

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The top level Linked service resource container.
type LinkedService struct {
	pulumi.CustomResourceState

	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The provisioning state of the linked service.
	ProvisioningState pulumi.StringPtrOutput `pulumi:"provisioningState"`
	// The resource id of the resource that will be linked to the workspace. This should be used for linking resources which require read access
	ResourceId pulumi.StringPtrOutput `pulumi:"resourceId"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringOutput `pulumi:"type"`
	// The resource id of the resource that will be linked to the workspace. This should be used for linking resources which require write access
	WriteAccessResourceId pulumi.StringPtrOutput `pulumi:"writeAccessResourceId"`
}

// NewLinkedService registers a new resource with the given unique name, arguments, and options.
func NewLinkedService(ctx *pulumi.Context,
	name string, args *LinkedServiceArgs, opts ...pulumi.ResourceOption) (*LinkedService, error) {
	if args == nil || args.LinkedServiceName == nil {
		return nil, errors.New("missing required argument 'LinkedServiceName'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.WorkspaceName == nil {
		return nil, errors.New("missing required argument 'WorkspaceName'")
	}
	if args == nil {
		args = &LinkedServiceArgs{}
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azurerm:operationalinsights/latest:LinkedService"),
		},
		{
			Type: pulumi.String("azurerm:operationalinsights/v20151101preview:LinkedService"),
		},
		{
			Type: pulumi.String("azurerm:operationalinsights/v20190801preview:LinkedService"),
		},
		{
			Type: pulumi.String("azurerm:operationalinsights/v20200301preview:LinkedService"),
		},
	})
	opts = append(opts, aliases)
	var resource LinkedService
	err := ctx.RegisterResource("azurerm:operationalinsights/v20200801:LinkedService", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLinkedService gets an existing LinkedService resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLinkedService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LinkedServiceState, opts ...pulumi.ResourceOption) (*LinkedService, error) {
	var resource LinkedService
	err := ctx.ReadResource("azurerm:operationalinsights/v20200801:LinkedService", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LinkedService resources.
type linkedServiceState struct {
	// The name of the resource
	Name *string `pulumi:"name"`
	// The provisioning state of the linked service.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The resource id of the resource that will be linked to the workspace. This should be used for linking resources which require read access
	ResourceId *string `pulumi:"resourceId"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string `pulumi:"type"`
	// The resource id of the resource that will be linked to the workspace. This should be used for linking resources which require write access
	WriteAccessResourceId *string `pulumi:"writeAccessResourceId"`
}

type LinkedServiceState struct {
	// The name of the resource
	Name pulumi.StringPtrInput
	// The provisioning state of the linked service.
	ProvisioningState pulumi.StringPtrInput
	// The resource id of the resource that will be linked to the workspace. This should be used for linking resources which require read access
	ResourceId pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringPtrInput
	// The resource id of the resource that will be linked to the workspace. This should be used for linking resources which require write access
	WriteAccessResourceId pulumi.StringPtrInput
}

func (LinkedServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*linkedServiceState)(nil)).Elem()
}

type linkedServiceArgs struct {
	// Name of the linkedServices resource
	LinkedServiceName string `pulumi:"linkedServiceName"`
	// The provisioning state of the linked service.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The resource id of the resource that will be linked to the workspace. This should be used for linking resources which require read access
	ResourceId *string `pulumi:"resourceId"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
	// The resource id of the resource that will be linked to the workspace. This should be used for linking resources which require write access
	WriteAccessResourceId *string `pulumi:"writeAccessResourceId"`
}

// The set of arguments for constructing a LinkedService resource.
type LinkedServiceArgs struct {
	// Name of the linkedServices resource
	LinkedServiceName pulumi.StringInput
	// The provisioning state of the linked service.
	ProvisioningState pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The resource id of the resource that will be linked to the workspace. This should be used for linking resources which require read access
	ResourceId pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The name of the workspace.
	WorkspaceName pulumi.StringInput
	// The resource id of the resource that will be linked to the workspace. This should be used for linking resources which require write access
	WriteAccessResourceId pulumi.StringPtrInput
}

func (LinkedServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*linkedServiceArgs)(nil)).Elem()
}
