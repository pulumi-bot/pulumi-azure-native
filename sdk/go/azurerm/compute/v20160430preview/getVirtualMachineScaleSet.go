// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v20160430preview

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupVirtualMachineScaleSet(ctx *pulumi.Context, args *LookupVirtualMachineScaleSetArgs, opts ...pulumi.InvokeOption) (*LookupVirtualMachineScaleSetResult, error) {
	var rv LookupVirtualMachineScaleSetResult
	err := ctx.Invoke("azurerm:compute/v20160430preview:getVirtualMachineScaleSet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVirtualMachineScaleSetArgs struct {
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the VM scale set.
	VmScaleSetName string `pulumi:"vmScaleSetName"`
}

// Describes a Virtual Machine Scale Set.
type LookupVirtualMachineScaleSetResult struct {
	// The identity of the virtual machine scale set, if configured.
	Identity *VirtualMachineScaleSetIdentityResponse `pulumi:"identity"`
	// Resource location
	Location string `pulumi:"location"`
	// Resource name
	Name string `pulumi:"name"`
	// Specifies whether the Virtual Machine Scale Set should be overprovisioned.
	OverProvision *bool `pulumi:"overProvision"`
	// Specifies information about the marketplace image used to create the virtual machine. This element is only used for marketplace images. Before you can use a marketplace image from an API, you must enable the image for programmatic use.  In the Azure portal, find the marketplace image that you want to use and then click **Want to deploy programmatically, Get Started ->**. Enter any required information and then click **Save**.
	Plan *PlanResponse `pulumi:"plan"`
	// The provisioning state, which only appears in the response.
	ProvisioningState string `pulumi:"provisioningState"`
	// When true this limits the scale set to a single placement group, of max size 100 virtual machines.
	SinglePlacementGroup *bool `pulumi:"singlePlacementGroup"`
	// The virtual machine scale set sku.
	Sku *SkuResponse `pulumi:"sku"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type string `pulumi:"type"`
	// The upgrade policy.
	UpgradePolicy *UpgradePolicyResponse `pulumi:"upgradePolicy"`
	// The virtual machine profile.
	VirtualMachineProfile *VirtualMachineScaleSetVMProfileResponse `pulumi:"virtualMachineProfile"`
}
