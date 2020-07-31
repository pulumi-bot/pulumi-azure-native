// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20200501.Outputs
{

    [OutputType]
    public sealed class VpnLinkProviderPropertiesResponseResult
    {
        /// <summary>
        /// Name of the link provider.
        /// </summary>
        public readonly string? LinkProviderName;
        /// <summary>
        /// Link speed.
        /// </summary>
        public readonly int? LinkSpeedInMbps;

        [OutputConstructor]
        private VpnLinkProviderPropertiesResponseResult(
            string? linkProviderName,

            int? linkSpeedInMbps)
        {
            LinkProviderName = linkProviderName;
            LinkSpeedInMbps = linkSpeedInMbps;
        }
    }
}