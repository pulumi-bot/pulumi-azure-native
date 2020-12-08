// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Automation.V20170515Preview.Inputs
{

    /// <summary>
    /// Update specific properties of the software update configuration.
    /// </summary>
    public sealed class UpdateConfigurationArgs : Pulumi.ResourceArgs
    {
        [Input("azureVirtualMachines")]
        private InputList<string>? _azureVirtualMachines;

        /// <summary>
        /// List of azure resource Ids for azure virtual machines targeted by the software update configuration.
        /// </summary>
        public InputList<string> AzureVirtualMachines
        {
            get => _azureVirtualMachines ?? (_azureVirtualMachines = new InputList<string>());
            set => _azureVirtualMachines = value;
        }

        /// <summary>
        /// Maximum time allowed for the software update configuration run. Duration needs to be specified using the format PT[n]H[n]M[n]S as per ISO8601
        /// </summary>
        [Input("duration")]
        public Input<string>? Duration { get; set; }

        /// <summary>
        /// Linux specific update configuration.
        /// </summary>
        [Input("linux")]
        public Input<Inputs.LinuxPropertiesArgs>? Linux { get; set; }

        [Input("nonAzureComputerNames")]
        private InputList<string>? _nonAzureComputerNames;

        /// <summary>
        /// List of names of non-azure machines targeted by the software update configuration.
        /// </summary>
        public InputList<string> NonAzureComputerNames
        {
            get => _nonAzureComputerNames ?? (_nonAzureComputerNames = new InputList<string>());
            set => _nonAzureComputerNames = value;
        }

        /// <summary>
        /// operating system of target machines
        /// </summary>
        [Input("operatingSystem", required: true)]
        public Input<Pulumi.AzureNextGen.Automation.V20170515Preview.OperatingSystemType> OperatingSystem { get; set; } = null!;

        /// <summary>
        /// Group targets for the software update configuration.
        /// </summary>
        [Input("targets")]
        public Input<Inputs.TargetPropertiesArgs>? Targets { get; set; }

        /// <summary>
        /// Windows specific update configuration.
        /// </summary>
        [Input("windows")]
        public Input<Inputs.WindowsPropertiesArgs>? Windows { get; set; }

        public UpdateConfigurationArgs()
        {
        }
    }
}
