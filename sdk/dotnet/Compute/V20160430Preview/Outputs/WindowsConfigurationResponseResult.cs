// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Compute.V20160430Preview.Outputs
{

    [OutputType]
    public sealed class WindowsConfigurationResponseResult
    {
        /// <summary>
        /// Specifies additional base-64 encoded XML formatted information that can be included in the Unattend.xml file, which is used by Windows Setup.
        /// </summary>
        public readonly ImmutableArray<Outputs.AdditionalUnattendContentResponseResult> AdditionalUnattendContent;
        /// <summary>
        /// Indicates whether virtual machine is enabled for automatic updates.
        /// </summary>
        public readonly bool? EnableAutomaticUpdates;
        /// <summary>
        /// Indicates whether virtual machine agent should be provisioned on the virtual machine. &lt;br&gt;&lt;br&gt; When this property is not specified in the request body, default behavior is to set it to true.  This will ensure that VM Agent is installed on the VM so that extensions can be added to the VM later.
        /// </summary>
        public readonly bool? ProvisionVMAgent;
        /// <summary>
        /// Specifies the time zone of the virtual machine. e.g. "Pacific Standard Time"
        /// </summary>
        public readonly string? TimeZone;
        /// <summary>
        /// Specifies the Windows Remote Management listeners. This enables remote Windows PowerShell.
        /// </summary>
        public readonly Outputs.WinRMConfigurationResponseResult? WinRM;

        [OutputConstructor]
        private WindowsConfigurationResponseResult(
            ImmutableArray<Outputs.AdditionalUnattendContentResponseResult> additionalUnattendContent,

            bool? enableAutomaticUpdates,

            bool? provisionVMAgent,

            string? timeZone,

            Outputs.WinRMConfigurationResponseResult? winRM)
        {
            AdditionalUnattendContent = additionalUnattendContent;
            EnableAutomaticUpdates = enableAutomaticUpdates;
            ProvisionVMAgent = provisionVMAgent;
            TimeZone = timeZone;
            WinRM = winRM;
        }
    }
}
