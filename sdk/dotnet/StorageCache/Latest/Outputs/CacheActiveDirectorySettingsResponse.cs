// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.StorageCache.Latest.Outputs
{

    [OutputType]
    public sealed class CacheActiveDirectorySettingsResponse
    {
        /// <summary>
        /// The NetBIOS name to assign to the HPC Cache when it joins the Active Directory domain as a server. Length must 1-15 characters from the class [-0-9a-zA-Z].
        /// </summary>
        public readonly string CacheNetBiosName;
        /// <summary>
        /// Active Directory admin credentials used to join the HPC Cache to a domain.
        /// </summary>
        public readonly Outputs.CacheActiveDirectorySettingsResponseCredentials? Credentials;
        /// <summary>
        /// True if the HPC Cache is joined to the Active Directory domain.
        /// </summary>
        public readonly string DomainJoined;
        /// <summary>
        /// The fully qualified domain name of the Active Directory domain controller.
        /// </summary>
        public readonly string DomainName;
        /// <summary>
        /// The Active Directory domain's NetBIOS name.
        /// </summary>
        public readonly string DomainNetBiosName;
        /// <summary>
        /// Primary DNS IP address used to resolve the Active Directory domain controller's fully qualified domain name.
        /// </summary>
        public readonly string PrimaryDnsIpAddress;
        /// <summary>
        /// Secondary DNS IP address used to resolve the Active Directory domain controller's fully qualified domain name.
        /// </summary>
        public readonly string? SecondaryDnsIpAddress;

        [OutputConstructor]
        private CacheActiveDirectorySettingsResponse(
            string cacheNetBiosName,

            Outputs.CacheActiveDirectorySettingsResponseCredentials? credentials,

            string domainJoined,

            string domainName,

            string domainNetBiosName,

            string primaryDnsIpAddress,

            string? secondaryDnsIpAddress)
        {
            CacheNetBiosName = cacheNetBiosName;
            Credentials = credentials;
            DomainJoined = domainJoined;
            DomainName = domainName;
            DomainNetBiosName = domainNetBiosName;
            PrimaryDnsIpAddress = primaryDnsIpAddress;
            SecondaryDnsIpAddress = secondaryDnsIpAddress;
        }
    }
}
