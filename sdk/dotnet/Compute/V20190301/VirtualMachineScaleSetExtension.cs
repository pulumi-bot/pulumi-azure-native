// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Compute.V20190301
{
    /// <summary>
    /// Describes a Virtual Machine Scale Set Extension.
    /// </summary>
    public partial class VirtualMachineScaleSetExtension : Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the extension.
        /// </summary>
        [Output("name")]
        public Output<string?> Name { get; private set; } = null!;

        /// <summary>
        /// Describes the properties of a Virtual Machine Scale Set Extension.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.VirtualMachineScaleSetExtensionPropertiesResponseResult> Properties { get; private set; } = null!;


        /// <summary>
        /// Create a VirtualMachineScaleSetExtension resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VirtualMachineScaleSetExtension(string name, VirtualMachineScaleSetExtensionArgs args, CustomResourceOptions? options = null)
            : base("azurerm:compute/v20190301:VirtualMachineScaleSetExtension", name, args ?? new VirtualMachineScaleSetExtensionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VirtualMachineScaleSetExtension(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:compute/v20190301:VirtualMachineScaleSetExtension", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing VirtualMachineScaleSetExtension resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VirtualMachineScaleSetExtension Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new VirtualMachineScaleSetExtension(name, id, options);
        }
    }

    public sealed class VirtualMachineScaleSetExtensionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicates whether the extension should use a newer minor version if one is available at deployment time. Once deployed, however, the extension will not upgrade minor versions unless redeployed, even with this property set to true.
        /// </summary>
        [Input("autoUpgradeMinorVersion")]
        public Input<bool>? AutoUpgradeMinorVersion { get; set; }

        /// <summary>
        /// If a value is provided and is different from the previous value, the extension handler will be forced to update even if the extension configuration has not changed.
        /// </summary>
        [Input("forceUpdateTag")]
        public Input<string>? ForceUpdateTag { get; set; }

        /// <summary>
        /// The name of the VM scale set extension.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("protectedSettings")]
        private InputMap<object>? _protectedSettings;

        /// <summary>
        /// The extension can contain either protectedSettings or protectedSettingsFromKeyVault or no protected settings at all.
        /// </summary>
        public InputMap<object> ProtectedSettings
        {
            get => _protectedSettings ?? (_protectedSettings = new InputMap<object>());
            set => _protectedSettings = value;
        }

        [Input("provisionAfterExtensions")]
        private InputList<string>? _provisionAfterExtensions;

        /// <summary>
        /// Collection of extension names after which this extension needs to be provisioned.
        /// </summary>
        public InputList<string> ProvisionAfterExtensions
        {
            get => _provisionAfterExtensions ?? (_provisionAfterExtensions = new InputList<string>());
            set => _provisionAfterExtensions = value;
        }

        /// <summary>
        /// The name of the extension handler publisher.
        /// </summary>
        [Input("publisher")]
        public Input<string>? Publisher { get; set; }

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("settings")]
        private InputMap<object>? _settings;

        /// <summary>
        /// Json formatted public settings for the extension.
        /// </summary>
        public InputMap<object> Settings
        {
            get => _settings ?? (_settings = new InputMap<object>());
            set => _settings = value;
        }

        /// <summary>
        /// Specifies the type of the extension; an example is "CustomScriptExtension".
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// Specifies the version of the script handler.
        /// </summary>
        [Input("typeHandlerVersion")]
        public Input<string>? TypeHandlerVersion { get; set; }

        /// <summary>
        /// The name of the VM scale set where the extension should be create or updated.
        /// </summary>
        [Input("vmScaleSetName", required: true)]
        public Input<string> VmScaleSetName { get; set; } = null!;

        public VirtualMachineScaleSetExtensionArgs()
        {
        }
    }
}
