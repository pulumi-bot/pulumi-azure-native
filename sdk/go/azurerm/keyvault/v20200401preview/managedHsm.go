// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200401preview

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Resource information with extended details.
type ManagedHsm struct {
	pulumi.CustomResourceState

	// The supported Azure location where the managed HSM Pool should be created.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// The name of the managed HSM Pool.
	Name pulumi.StringOutput `pulumi:"name"`
	// Properties of the managed HSM
	Properties ManagedHsmPropertiesResponseOutput `pulumi:"properties"`
	// SKU details
	Sku ManagedHsmSkuResponsePtrOutput `pulumi:"sku"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The resource type of the managed HSM Pool.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewManagedHsm registers a new resource with the given unique name, arguments, and options.
func NewManagedHsm(ctx *pulumi.Context,
	name string, args *ManagedHsmArgs, opts ...pulumi.ResourceOption) (*ManagedHsm, error) {
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &ManagedHsmArgs{}
	}
	var resource ManagedHsm
	err := ctx.RegisterResource("azurerm:keyvault/v20200401preview:ManagedHsm", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetManagedHsm gets an existing ManagedHsm resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetManagedHsm(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagedHsmState, opts ...pulumi.ResourceOption) (*ManagedHsm, error) {
	var resource ManagedHsm
	err := ctx.ReadResource("azurerm:keyvault/v20200401preview:ManagedHsm", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ManagedHsm resources.
type managedHsmState struct {
	// The supported Azure location where the managed HSM Pool should be created.
	Location *string `pulumi:"location"`
	// The name of the managed HSM Pool.
	Name *string `pulumi:"name"`
	// Properties of the managed HSM
	Properties *ManagedHsmPropertiesResponse `pulumi:"properties"`
	// SKU details
	Sku *ManagedHsmSkuResponse `pulumi:"sku"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// The resource type of the managed HSM Pool.
	Type *string `pulumi:"type"`
}

type ManagedHsmState struct {
	// The supported Azure location where the managed HSM Pool should be created.
	Location pulumi.StringPtrInput
	// The name of the managed HSM Pool.
	Name pulumi.StringPtrInput
	// Properties of the managed HSM
	Properties ManagedHsmPropertiesResponsePtrInput
	// SKU details
	Sku ManagedHsmSkuResponsePtrInput
	// Resource tags
	Tags pulumi.StringMapInput
	// The resource type of the managed HSM Pool.
	Type pulumi.StringPtrInput
}

func (ManagedHsmState) ElementType() reflect.Type {
	return reflect.TypeOf((*managedHsmState)(nil)).Elem()
}

type managedHsmArgs struct {
	// The supported Azure location where the managed HSM Pool should be created.
	Location *string `pulumi:"location"`
	// Name of the managed HSM Pool
	Name string `pulumi:"name"`
	// Properties of the managed HSM
	Properties *ManagedHsmProperties `pulumi:"properties"`
	// Name of the resource group that contains the managed HSM pool.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// SKU details
	Sku *ManagedHsmSku `pulumi:"sku"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a ManagedHsm resource.
type ManagedHsmArgs struct {
	// The supported Azure location where the managed HSM Pool should be created.
	Location pulumi.StringPtrInput
	// Name of the managed HSM Pool
	Name pulumi.StringInput
	// Properties of the managed HSM
	Properties ManagedHsmPropertiesPtrInput
	// Name of the resource group that contains the managed HSM pool.
	ResourceGroupName pulumi.StringInput
	// SKU details
	Sku ManagedHsmSkuPtrInput
	// Resource tags
	Tags pulumi.StringMapInput
}

func (ManagedHsmArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managedHsmArgs)(nil)).Elem()
}
