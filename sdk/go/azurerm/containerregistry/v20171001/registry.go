// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20171001

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// An object that represents a container registry.
type Registry struct {
	pulumi.CustomResourceState

	// The location of the resource. This cannot be changed after the resource is created.
	Location pulumi.StringOutput `pulumi:"location"`
	// The name of the resource.
	Name pulumi.StringOutput `pulumi:"name"`
	// The properties of the container registry.
	Properties RegistryPropertiesResponseOutput `pulumi:"properties"`
	// The SKU of the container registry.
	Sku SkuResponseOutput `pulumi:"sku"`
	// The tags of the resource.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewRegistry registers a new resource with the given unique name, arguments, and options.
func NewRegistry(ctx *pulumi.Context,
	name string, args *RegistryArgs, opts ...pulumi.ResourceOption) (*Registry, error) {
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil || args.Sku == nil {
		return nil, errors.New("missing required argument 'Sku'")
	}
	if args == nil {
		args = &RegistryArgs{}
	}
	var resource Registry
	err := ctx.RegisterResource("azurerm:containerregistry/v20171001:Registry", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRegistry gets an existing Registry resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRegistry(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegistryState, opts ...pulumi.ResourceOption) (*Registry, error) {
	var resource Registry
	err := ctx.ReadResource("azurerm:containerregistry/v20171001:Registry", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Registry resources.
type registryState struct {
	// The location of the resource. This cannot be changed after the resource is created.
	Location *string `pulumi:"location"`
	// The name of the resource.
	Name *string `pulumi:"name"`
	// The properties of the container registry.
	Properties *RegistryPropertiesResponse `pulumi:"properties"`
	// The SKU of the container registry.
	Sku *SkuResponse `pulumi:"sku"`
	// The tags of the resource.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource.
	Type *string `pulumi:"type"`
}

type RegistryState struct {
	// The location of the resource. This cannot be changed after the resource is created.
	Location pulumi.StringPtrInput
	// The name of the resource.
	Name pulumi.StringPtrInput
	// The properties of the container registry.
	Properties RegistryPropertiesResponsePtrInput
	// The SKU of the container registry.
	Sku SkuResponsePtrInput
	// The tags of the resource.
	Tags pulumi.StringMapInput
	// The type of the resource.
	Type pulumi.StringPtrInput
}

func (RegistryState) ElementType() reflect.Type {
	return reflect.TypeOf((*registryState)(nil)).Elem()
}

type registryArgs struct {
	// The value that indicates whether the admin user is enabled.
	AdminUserEnabled *bool `pulumi:"adminUserEnabled"`
	// The location of the resource. This cannot be changed after the resource is created.
	Location string `pulumi:"location"`
	// The name of the container registry.
	Name string `pulumi:"name"`
	// The network rule set for a container registry.
	NetworkRuleSet *NetworkRuleSet `pulumi:"networkRuleSet"`
	// The name of the resource group to which the container registry belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The SKU of the container registry.
	Sku Sku `pulumi:"sku"`
	// The properties of the storage account for the container registry. Only applicable to Classic SKU.
	StorageAccount *StorageAccountProperties `pulumi:"storageAccount"`
	// The tags of the resource.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Registry resource.
type RegistryArgs struct {
	// The value that indicates whether the admin user is enabled.
	AdminUserEnabled pulumi.BoolPtrInput
	// The location of the resource. This cannot be changed after the resource is created.
	Location pulumi.StringInput
	// The name of the container registry.
	Name pulumi.StringInput
	// The network rule set for a container registry.
	NetworkRuleSet NetworkRuleSetPtrInput
	// The name of the resource group to which the container registry belongs.
	ResourceGroupName pulumi.StringInput
	// The SKU of the container registry.
	Sku SkuInput
	// The properties of the storage account for the container registry. Only applicable to Classic SKU.
	StorageAccount StorageAccountPropertiesPtrInput
	// The tags of the resource.
	Tags pulumi.StringMapInput
}

func (RegistryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*registryArgs)(nil)).Elem()
}
