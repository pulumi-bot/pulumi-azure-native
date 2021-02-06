// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Compute.V20160430Preview
{
    /// <summary>
    /// Describes a Virtual Machine Scale Set.
    /// </summary>
    [AzureNextGenResourceType("azure-nextgen:compute/v20160430preview:VirtualMachineScaleSet")]
    public partial class VirtualMachineScaleSet : Pulumi.CustomResource
    {
        /// <summary>
        /// The identity of the virtual machine scale set, if configured.
        /// </summary>
        [Output("identity")]
        public Output<Outputs.VirtualMachineScaleSetIdentityResponse?> Identity { get; private set; } = null!;

        /// <summary>
        /// Resource location
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Resource name
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Specifies whether the Virtual Machine Scale Set should be overprovisioned.
        /// </summary>
        [Output("overProvision")]
        public Output<bool?> OverProvision { get; private set; } = null!;

        /// <summary>
        /// Specifies information about the marketplace image used to create the virtual machine. This element is only used for marketplace images. Before you can use a marketplace image from an API, you must enable the image for programmatic use.  In the Azure portal, find the marketplace image that you want to use and then click **Want to deploy programmatically, Get Started -&gt;**. Enter any required information and then click **Save**.
        /// </summary>
        [Output("plan")]
        public Output<Outputs.PlanResponse?> Plan { get; private set; } = null!;

        /// <summary>
        /// The provisioning state, which only appears in the response.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// When true this limits the scale set to a single placement group, of max size 100 virtual machines.
        /// </summary>
        [Output("singlePlacementGroup")]
        public Output<bool?> SinglePlacementGroup { get; private set; } = null!;

        /// <summary>
        /// The virtual machine scale set sku.
        /// </summary>
        [Output("sku")]
        public Output<Outputs.SkuResponse?> Sku { get; private set; } = null!;

        /// <summary>
        /// Resource tags
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Resource type
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// The upgrade policy.
        /// </summary>
        [Output("upgradePolicy")]
        public Output<Outputs.UpgradePolicyResponse?> UpgradePolicy { get; private set; } = null!;

        /// <summary>
        /// The virtual machine profile.
        /// </summary>
        [Output("virtualMachineProfile")]
        public Output<Outputs.VirtualMachineScaleSetVMProfileResponse?> VirtualMachineProfile { get; private set; } = null!;


        /// <summary>
        /// Create a VirtualMachineScaleSet resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VirtualMachineScaleSet(string name, VirtualMachineScaleSetArgs args, CustomResourceOptions? options = null)
            : base("azure-nextgen:compute/v20160430preview:VirtualMachineScaleSet", name, args ?? new VirtualMachineScaleSetArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VirtualMachineScaleSet(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-nextgen:compute/v20160430preview:VirtualMachineScaleSet", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azure-nextgen:compute/latest:VirtualMachineScaleSet"},
                    new Pulumi.Alias { Type = "azure-nextgen:compute/v20150615:VirtualMachineScaleSet"},
                    new Pulumi.Alias { Type = "azure-nextgen:compute/v20160330:VirtualMachineScaleSet"},
                    new Pulumi.Alias { Type = "azure-nextgen:compute/v20170330:VirtualMachineScaleSet"},
                    new Pulumi.Alias { Type = "azure-nextgen:compute/v20171201:VirtualMachineScaleSet"},
                    new Pulumi.Alias { Type = "azure-nextgen:compute/v20180401:VirtualMachineScaleSet"},
                    new Pulumi.Alias { Type = "azure-nextgen:compute/v20180601:VirtualMachineScaleSet"},
                    new Pulumi.Alias { Type = "azure-nextgen:compute/v20181001:VirtualMachineScaleSet"},
                    new Pulumi.Alias { Type = "azure-nextgen:compute/v20190301:VirtualMachineScaleSet"},
                    new Pulumi.Alias { Type = "azure-nextgen:compute/v20190701:VirtualMachineScaleSet"},
                    new Pulumi.Alias { Type = "azure-nextgen:compute/v20191201:VirtualMachineScaleSet"},
                    new Pulumi.Alias { Type = "azure-nextgen:compute/v20200601:VirtualMachineScaleSet"},
                    new Pulumi.Alias { Type = "azure-nextgen:compute/v20201201:VirtualMachineScaleSet"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing VirtualMachineScaleSet resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VirtualMachineScaleSet Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new VirtualMachineScaleSet(name, id, options);
        }
    }

    public sealed class VirtualMachineScaleSetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The identity of the virtual machine scale set, if configured.
        /// </summary>
        [Input("identity")]
        public Input<Inputs.VirtualMachineScaleSetIdentityArgs>? Identity { get; set; }

        /// <summary>
        /// Resource location
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the VM scale set to create or update.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Specifies whether the Virtual Machine Scale Set should be overprovisioned.
        /// </summary>
        [Input("overProvision")]
        public Input<bool>? OverProvision { get; set; }

        /// <summary>
        /// Specifies information about the marketplace image used to create the virtual machine. This element is only used for marketplace images. Before you can use a marketplace image from an API, you must enable the image for programmatic use.  In the Azure portal, find the marketplace image that you want to use and then click **Want to deploy programmatically, Get Started -&gt;**. Enter any required information and then click **Save**.
        /// </summary>
        [Input("plan")]
        public Input<Inputs.PlanArgs>? Plan { get; set; }

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// When true this limits the scale set to a single placement group, of max size 100 virtual machines.
        /// </summary>
        [Input("singlePlacementGroup")]
        public Input<bool>? SinglePlacementGroup { get; set; }

        /// <summary>
        /// The virtual machine scale set sku.
        /// </summary>
        [Input("sku")]
        public Input<Inputs.SkuArgs>? Sku { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Resource tags
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The upgrade policy.
        /// </summary>
        [Input("upgradePolicy")]
        public Input<Inputs.UpgradePolicyArgs>? UpgradePolicy { get; set; }

        /// <summary>
        /// The virtual machine profile.
        /// </summary>
        [Input("virtualMachineProfile")]
        public Input<Inputs.VirtualMachineScaleSetVMProfileArgs>? VirtualMachineProfile { get; set; }

        public VirtualMachineScaleSetArgs()
        {
        }
    }
}
