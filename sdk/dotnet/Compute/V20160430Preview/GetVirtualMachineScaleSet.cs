// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Compute.V20160430Preview
{
    public static class GetVirtualMachineScaleSet
    {
        public static Task<GetVirtualMachineScaleSetResult> InvokeAsync(GetVirtualMachineScaleSetArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetVirtualMachineScaleSetResult>("azurerm:compute/v20160430preview:getVirtualMachineScaleSet", args ?? new GetVirtualMachineScaleSetArgs(), options.WithVersion());
    }


    public sealed class GetVirtualMachineScaleSetArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the VM scale set.
        /// </summary>
        [Input("vmScaleSetName", required: true)]
        public string VmScaleSetName { get; set; } = null!;

        public GetVirtualMachineScaleSetArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetVirtualMachineScaleSetResult
    {
        /// <summary>
        /// The identity of the virtual machine scale set, if configured.
        /// </summary>
        public readonly Outputs.VirtualMachineScaleSetIdentityResponseResult? Identity;
        /// <summary>
        /// Resource location
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// Resource name
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Specifies whether the Virtual Machine Scale Set should be overprovisioned.
        /// </summary>
        public readonly bool? OverProvision;
        /// <summary>
        /// Specifies information about the marketplace image used to create the virtual machine. This element is only used for marketplace images. Before you can use a marketplace image from an API, you must enable the image for programmatic use.  In the Azure portal, find the marketplace image that you want to use and then click **Want to deploy programmatically, Get Started -&gt;**. Enter any required information and then click **Save**.
        /// </summary>
        public readonly Outputs.PlanResponseResult? Plan;
        /// <summary>
        /// The provisioning state, which only appears in the response.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// When true this limits the scale set to a single placement group, of max size 100 virtual machines.
        /// </summary>
        public readonly bool? SinglePlacementGroup;
        /// <summary>
        /// The virtual machine scale set sku.
        /// </summary>
        public readonly Outputs.SkuResponseResult? Sku;
        /// <summary>
        /// Resource tags
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Resource type
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The upgrade policy.
        /// </summary>
        public readonly Outputs.UpgradePolicyResponseResult? UpgradePolicy;
        /// <summary>
        /// The virtual machine profile.
        /// </summary>
        public readonly Outputs.VirtualMachineScaleSetVMProfileResponseResult? VirtualMachineProfile;

        [OutputConstructor]
        private GetVirtualMachineScaleSetResult(
            Outputs.VirtualMachineScaleSetIdentityResponseResult? identity,

            string location,

            string name,

            bool? overProvision,

            Outputs.PlanResponseResult? plan,

            string provisioningState,

            bool? singlePlacementGroup,

            Outputs.SkuResponseResult? sku,

            ImmutableDictionary<string, string>? tags,

            string type,

            Outputs.UpgradePolicyResponseResult? upgradePolicy,

            Outputs.VirtualMachineScaleSetVMProfileResponseResult? virtualMachineProfile)
        {
            Identity = identity;
            Location = location;
            Name = name;
            OverProvision = overProvision;
            Plan = plan;
            ProvisioningState = provisioningState;
            SinglePlacementGroup = singlePlacementGroup;
            Sku = sku;
            Tags = tags;
            Type = type;
            UpgradePolicy = upgradePolicy;
            VirtualMachineProfile = virtualMachineProfile;
        }
    }
}
