// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Compute.V20160430Preview.Inputs
{

    /// <summary>
    /// Describes a Virtual Machine Scale Set Extension.
    /// </summary>
    public sealed class VirtualMachineScaleSetExtensionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicates whether the extension should use a newer minor version if one is available at deployment time. Once deployed, however, the extension will not upgrade minor versions unless redeployed, even with this property set to true.
        /// </summary>
        [Input("autoUpgradeMinorVersion")]
        public Input<bool>? AutoUpgradeMinorVersion { get; set; }

        /// <summary>
        /// The name of the extension.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

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

        /// <summary>
        /// The name of the extension handler publisher.
        /// </summary>
        [Input("publisher")]
        public Input<string>? Publisher { get; set; }

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

        public VirtualMachineScaleSetExtensionArgs()
        {
        }
    }
}
