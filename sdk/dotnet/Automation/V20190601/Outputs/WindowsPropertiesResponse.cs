// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Automation.V20190601.Outputs
{

    [OutputType]
    public sealed class WindowsPropertiesResponse
    {
        /// <summary>
        /// KB numbers excluded from the software update configuration.
        /// </summary>
        public readonly ImmutableArray<string> ExcludedKbNumbers;
        /// <summary>
        /// KB numbers included from the software update configuration.
        /// </summary>
        public readonly ImmutableArray<string> IncludedKbNumbers;
        /// <summary>
        /// Update classification included in the software update configuration. A comma separated string with required values
        /// </summary>
        public readonly string? IncludedUpdateClassifications;
        /// <summary>
        /// Reboot setting for the software update configuration.
        /// </summary>
        public readonly string? RebootSetting;

        [OutputConstructor]
        private WindowsPropertiesResponse(
            ImmutableArray<string> excludedKbNumbers,

            ImmutableArray<string> includedKbNumbers,

            string? includedUpdateClassifications,

            string? rebootSetting)
        {
            ExcludedKbNumbers = excludedKbNumbers;
            IncludedKbNumbers = includedKbNumbers;
            IncludedUpdateClassifications = includedUpdateClassifications;
            RebootSetting = rebootSetting;
        }
    }
}
