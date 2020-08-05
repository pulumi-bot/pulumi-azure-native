// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20200601

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Describes a Virtual Machine Scale Set.
type VirtualMachineScaleSet struct {
	pulumi.CustomResourceState

	// The identity of the virtual machine scale set, if configured.
	Identity VirtualMachineScaleSetIdentityResponsePtrOutput `pulumi:"identity"`
	// Resource location
	Location pulumi.StringOutput `pulumi:"location"`
	// Resource name
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies information about the marketplace image used to create the virtual machine. This element is only used for marketplace images. Before you can use a marketplace image from an API, you must enable the image for programmatic use.  In the Azure portal, find the marketplace image that you want to use and then click **Want to deploy programmatically, Get Started ->**. Enter any required information and then click **Save**.
	Plan PlanResponsePtrOutput `pulumi:"plan"`
	// Describes the properties of a Virtual Machine Scale Set.
	Properties VirtualMachineScaleSetPropertiesResponseOutput `pulumi:"properties"`
	// The virtual machine scale set sku.
	Sku SkuResponsePtrOutput `pulumi:"sku"`
	// Resource tags
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Resource type
	Type pulumi.StringOutput `pulumi:"type"`
	// The virtual machine scale set zones. NOTE: Availability zones can only be set when you create the scale set
	Zones pulumi.StringArrayOutput `pulumi:"zones"`
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
	err := ctx.RegisterResource("azurerm:compute/v20200601:VirtualMachineScaleSet", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azurerm:compute/v20200601:VirtualMachineScaleSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VirtualMachineScaleSet resources.
type virtualMachineScaleSetState struct {
	// The identity of the virtual machine scale set, if configured.
	Identity *VirtualMachineScaleSetIdentityResponse `pulumi:"identity"`
	// Resource location
	Location *string `pulumi:"location"`
	// Resource name
	Name *string `pulumi:"name"`
	// Specifies information about the marketplace image used to create the virtual machine. This element is only used for marketplace images. Before you can use a marketplace image from an API, you must enable the image for programmatic use.  In the Azure portal, find the marketplace image that you want to use and then click **Want to deploy programmatically, Get Started ->**. Enter any required information and then click **Save**.
	Plan *PlanResponse `pulumi:"plan"`
	// Describes the properties of a Virtual Machine Scale Set.
	Properties *VirtualMachineScaleSetPropertiesResponse `pulumi:"properties"`
	// The virtual machine scale set sku.
	Sku *SkuResponse `pulumi:"sku"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type *string `pulumi:"type"`
	// The virtual machine scale set zones. NOTE: Availability zones can only be set when you create the scale set
	Zones []string `pulumi:"zones"`
}

type VirtualMachineScaleSetState struct {
	// The identity of the virtual machine scale set, if configured.
	Identity VirtualMachineScaleSetIdentityResponsePtrInput
	// Resource location
	Location pulumi.StringPtrInput
	// Resource name
	Name pulumi.StringPtrInput
	// Specifies information about the marketplace image used to create the virtual machine. This element is only used for marketplace images. Before you can use a marketplace image from an API, you must enable the image for programmatic use.  In the Azure portal, find the marketplace image that you want to use and then click **Want to deploy programmatically, Get Started ->**. Enter any required information and then click **Save**.
	Plan PlanResponsePtrInput
	// Describes the properties of a Virtual Machine Scale Set.
	Properties VirtualMachineScaleSetPropertiesResponsePtrInput
	// The virtual machine scale set sku.
	Sku SkuResponsePtrInput
	// Resource tags
	Tags pulumi.StringMapInput
	// Resource type
	Type pulumi.StringPtrInput
	// The virtual machine scale set zones. NOTE: Availability zones can only be set when you create the scale set
	Zones pulumi.StringArrayInput
}

func (VirtualMachineScaleSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualMachineScaleSetState)(nil)).Elem()
}

type virtualMachineScaleSetArgs struct {
	// Specifies additional capabilities enabled or disabled on the Virtual Machines in the Virtual Machine Scale Set. For instance: whether the Virtual Machines have the capability to support attaching managed data disks with UltraSSD_LRS storage account type.
	AdditionalCapabilities *AdditionalCapabilities `pulumi:"additionalCapabilities"`
	// Policy for automatic repairs.
	AutomaticRepairsPolicy *AutomaticRepairsPolicy `pulumi:"automaticRepairsPolicy"`
	// When Overprovision is enabled, extensions are launched only on the requested number of VMs which are finally kept. This property will hence ensure that the extensions do not run on the extra overprovisioned VMs.
	DoNotRunExtensionsOnOverprovisionedVMs *bool `pulumi:"doNotRunExtensionsOnOverprovisionedVMs"`
	// Specifies information about the dedicated host group that the virtual machine scale set resides in. <br><br>Minimum api-version: 2020-06-01.
	HostGroup *SubResource `pulumi:"hostGroup"`
	// The identity of the virtual machine scale set, if configured.
	Identity *VirtualMachineScaleSetIdentity `pulumi:"identity"`
	// Resource location
	Location string `pulumi:"location"`
	// The name of the VM scale set to create or update.
	Name string `pulumi:"name"`
	// Specifies whether the Virtual Machine Scale Set should be overprovisioned.
	Overprovision *bool `pulumi:"overprovision"`
	// Specifies information about the marketplace image used to create the virtual machine. This element is only used for marketplace images. Before you can use a marketplace image from an API, you must enable the image for programmatic use.  In the Azure portal, find the marketplace image that you want to use and then click **Want to deploy programmatically, Get Started ->**. Enter any required information and then click **Save**.
	Plan *Plan `pulumi:"plan"`
	// Fault Domain count for each placement group.
	PlatformFaultDomainCount *int `pulumi:"platformFaultDomainCount"`
	// Specifies information about the proximity placement group that the virtual machine scale set should be assigned to. <br><br>Minimum api-version: 2018-04-01.
	ProximityPlacementGroup *SubResource `pulumi:"proximityPlacementGroup"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Specifies the scale-in policy that decides which virtual machines are chosen for removal when a Virtual Machine Scale Set is scaled-in.
	ScaleInPolicy *ScaleInPolicy `pulumi:"scaleInPolicy"`
	// When true this limits the scale set to a single placement group, of max size 100 virtual machines. NOTE: If singlePlacementGroup is true, it may be modified to false. However, if singlePlacementGroup is false, it may not be modified to true.
	SinglePlacementGroup *bool `pulumi:"singlePlacementGroup"`
	// The virtual machine scale set sku.
	Sku *Sku `pulumi:"sku"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// The upgrade policy.
	UpgradePolicy *UpgradePolicy `pulumi:"upgradePolicy"`
	// The virtual machine profile.
	VirtualMachineProfile *VirtualMachineScaleSetVMProfile `pulumi:"virtualMachineProfile"`
	// Whether to force strictly even Virtual Machine distribution cross x-zones in case there is zone outage.
	ZoneBalance *bool `pulumi:"zoneBalance"`
	// The virtual machine scale set zones. NOTE: Availability zones can only be set when you create the scale set
	Zones []string `pulumi:"zones"`
}

// The set of arguments for constructing a VirtualMachineScaleSet resource.
type VirtualMachineScaleSetArgs struct {
	// Specifies additional capabilities enabled or disabled on the Virtual Machines in the Virtual Machine Scale Set. For instance: whether the Virtual Machines have the capability to support attaching managed data disks with UltraSSD_LRS storage account type.
	AdditionalCapabilities AdditionalCapabilitiesPtrInput
	// Policy for automatic repairs.
	AutomaticRepairsPolicy AutomaticRepairsPolicyPtrInput
	// When Overprovision is enabled, extensions are launched only on the requested number of VMs which are finally kept. This property will hence ensure that the extensions do not run on the extra overprovisioned VMs.
	DoNotRunExtensionsOnOverprovisionedVMs pulumi.BoolPtrInput
	// Specifies information about the dedicated host group that the virtual machine scale set resides in. <br><br>Minimum api-version: 2020-06-01.
	HostGroup SubResourcePtrInput
	// The identity of the virtual machine scale set, if configured.
	Identity VirtualMachineScaleSetIdentityPtrInput
	// Resource location
	Location pulumi.StringInput
	// The name of the VM scale set to create or update.
	Name pulumi.StringInput
	// Specifies whether the Virtual Machine Scale Set should be overprovisioned.
	Overprovision pulumi.BoolPtrInput
	// Specifies information about the marketplace image used to create the virtual machine. This element is only used for marketplace images. Before you can use a marketplace image from an API, you must enable the image for programmatic use.  In the Azure portal, find the marketplace image that you want to use and then click **Want to deploy programmatically, Get Started ->**. Enter any required information and then click **Save**.
	Plan PlanPtrInput
	// Fault Domain count for each placement group.
	PlatformFaultDomainCount pulumi.IntPtrInput
	// Specifies information about the proximity placement group that the virtual machine scale set should be assigned to. <br><br>Minimum api-version: 2018-04-01.
	ProximityPlacementGroup SubResourcePtrInput
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput
	// Specifies the scale-in policy that decides which virtual machines are chosen for removal when a Virtual Machine Scale Set is scaled-in.
	ScaleInPolicy ScaleInPolicyPtrInput
	// When true this limits the scale set to a single placement group, of max size 100 virtual machines. NOTE: If singlePlacementGroup is true, it may be modified to false. However, if singlePlacementGroup is false, it may not be modified to true.
	SinglePlacementGroup pulumi.BoolPtrInput
	// The virtual machine scale set sku.
	Sku SkuPtrInput
	// Resource tags
	Tags pulumi.StringMapInput
	// The upgrade policy.
	UpgradePolicy UpgradePolicyPtrInput
	// The virtual machine profile.
	VirtualMachineProfile VirtualMachineScaleSetVMProfilePtrInput
	// Whether to force strictly even Virtual Machine distribution cross x-zones in case there is zone outage.
	ZoneBalance pulumi.BoolPtrInput
	// The virtual machine scale set zones. NOTE: Availability zones can only be set when you create the scale set
	Zones pulumi.StringArrayInput
}

func (VirtualMachineScaleSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualMachineScaleSetArgs)(nil)).Elem()
}
