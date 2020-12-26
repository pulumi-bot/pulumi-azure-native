// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.StorageCache.Latest.Inputs
{

    /// <summary>
    /// Active Directory settings used to join a cache to a domain.
    /// </summary>
    public sealed class CacheActiveDirectorySettingsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The NetBIOS name to assign to the HPC Cache when it joins the Active Directory domain as a server. Length must 1-15 characters from the class [-0-9a-zA-Z].
        /// </summary>
        [Input("cacheNetBiosName", required: true)]
        public Input<string> CacheNetBiosName { get; set; } = null!;

        /// <summary>
        /// Active Directory admin credentials used to join the HPC Cache to a domain.
        /// </summary>
        [Input("credentials")]
        public Input<Inputs.CacheActiveDirectorySettingsCredentialsArgs>? Credentials { get; set; }

        /// <summary>
        /// The fully qualified domain name of the Active Directory domain controller.
        /// </summary>
        [Input("domainName", required: true)]
        public Input<string> DomainName { get; set; } = null!;

        /// <summary>
        /// The Active Directory domain's NetBIOS name.
        /// </summary>
        [Input("domainNetBiosName", required: true)]
        public Input<string> DomainNetBiosName { get; set; } = null!;

        /// <summary>
        /// Primary DNS IP address used to resolve the Active Directory domain controller's fully qualified domain name.
        /// </summary>
        [Input("primaryDnsIpAddress", required: true)]
        public Input<string> PrimaryDnsIpAddress { get; set; } = null!;

        /// <summary>
        /// Secondary DNS IP address used to resolve the Active Directory domain controller's fully qualified domain name.
        /// </summary>
        [Input("secondaryDnsIpAddress")]
        public Input<string>? SecondaryDnsIpAddress { get; set; }

        public CacheActiveDirectorySettingsArgs()
        {
        }
    }
}
