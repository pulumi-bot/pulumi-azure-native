// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20150615

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Describes a Virtual Machine Scale Set.
type VirtualMachineScaleSet struct {
	pulumi.CustomResourceState

	// Resource location
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Describes the properties of a Virtual Machine Scale Set.
	Properties VirtualMachineScaleSetPropertiesResponseOutput `pulumi:"properties"`
	// The virtual machine scale set sku.
	Sku SkuResponsePtrOutput `pulumi:"sku"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewVirtualMachineScaleSet registers a new resource with the given unique name, arguments, and options.
func NewVirtualMachineScaleSet(ctx *pulumi.Context,
	name string, args *VirtualMachineScaleSetArgs, opts ...pulumi.ResourceOption) (*VirtualMachineScaleSet, error) {
	if args == nil || args.Location == nil {
		return nil, errors.New("missing required argument 'Location'")
	}
	if args == nil || args.Name == nil {
		return nil, errors.New("missing required argument 'Name'")
	}
	if args == nil || args.ResourceGroupName == nil {
		return nil, errors.New("missing required argument 'ResourceGroupName'")
	}
	if args == nil {
		args = &VirtualMachineScaleSetArgs{}
	}
	var resource VirtualMachineScaleSet
	err := ctx.RegisterResource("azurerm:compute/v20150615:VirtualMachineScaleSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVirtualMachineScaleSet gets an existing VirtualMachineScaleSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVirtualMachineScaleSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualMachineScaleSetState, opts ...pulumi.ResourceOption) (*VirtualMachineScaleSet, error) {
	var resource VirtualMachineScaleSet
	err := ctx.ReadResource("azurerm:compute/v20150615:VirtualMachineScaleSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualMachineScaleSet resources.
type virtualMachineScaleSetState struct {
	// Resource location
	Location *string `pulumi:"location"`
	// Resource name
	Name *string `pulumi:"name"`
	// Describes the properties of a Virtual Machine Scale Set.
	Properties *VirtualMachineScaleSetPropertiesResponse `pulumi:"properties"`
	// The virtual machine scale set sku.
	Sku *SkuResponse `pulumi:"sku"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type *string `pulumi:"type"`
}

type VirtualMachineScaleSetState struct {
	// Resource location
	Location pulumi.StringPtrInput
	// Resource name
	Name pulumi.StringPtrInput
	// Describes the properties of a Virtual Machine Scale Set.
	Properties VirtualMachineScaleSetPropertiesResponsePtrInput
	// The virtual machine scale set sku.
	Sku SkuResponsePtrInput
	// Resource tags
	Tags pulumi.StringMapInput
	// Resource type
	Type pulumi.StringPtrInput
}

func (VirtualMachineScaleSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualMachineScaleSetState)(nil)).Elem()
}

type virtualMachineScaleSetArgs struct {
	// Resource location
	Location string `pulumi:"location"`
	// The name of the VM scale set to create or update.
	Name string `pulumi:"name"`
	// Specifies whether the Virtual Machine Scale Set should be overprovisioned.
	OverProvision *bool `pulumi:"overProvision"`
	// The provisioning state, which only appears in the response.
	ProvisioningState *string `pulumi:"provisioningState"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The virtual machine scale set sku.
	Sku *Sku `pulumi:"sku"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// The upgrade policy.
	UpgradePolicy *UpgradePolicy `pulumi:"upgradePolicy"`
	// The virtual machine profile.
	VirtualMachineProfile *VirtualMachineScaleSetVMProfile `pulumi:"virtualMachineProfile"`
}

// The set of arguments for constructing a VirtualMachineScaleSet resource.
type VirtualMachineScaleSetArgs struct {
	// Resource location
	Location pulumi.StringInput
	// The name of the VM scale set to create or update.
	Name pulumi.StringInput
	// Specifies whether the Virtual Machine Scale Set should be overprovisioned.
	OverProvision pulumi.BoolPtrInput
	// The provisioning state, which only appears in the response.
	ProvisioningState pulumi.StringPtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// The virtual machine scale set sku.
	Sku SkuPtrInput
	// Resource tags
	Tags pulumi.StringMapInput
	// The upgrade policy.
	UpgradePolicy UpgradePolicyPtrInput
	// The virtual machine profile.
	VirtualMachineProfile VirtualMachineScaleSetVMProfilePtrInput
}

func (VirtualMachineScaleSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualMachineScaleSetArgs)(nil)).Elem()
}
