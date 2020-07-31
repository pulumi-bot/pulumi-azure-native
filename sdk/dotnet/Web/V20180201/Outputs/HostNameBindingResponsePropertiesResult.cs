// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Web.V20180201.Outputs
{

    [OutputType]
    public sealed class HostNameBindingResponsePropertiesResult
    {
        /// <summary>
        /// Azure resource name.
        /// </summary>
        public readonly string? AzureResourceName;
        /// <summary>
        /// Azure resource type.
        /// </summary>
        public readonly string? AzureResourceType;
        /// <summary>
        /// Custom DNS record type.
        /// </summary>
        public readonly string? CustomHostNameDnsRecordType;
        /// <summary>
        /// Fully qualified ARM domain resource URI.
        /// </summary>
        public readonly string? DomainId;
        /// <summary>
        /// Hostname type.
        /// </summary>
        public readonly string? HostNameType;
        /// <summary>
        /// App Service app name.
        /// </summary>
        public readonly string? SiteName;
        /// <summary>
        /// SSL type
        /// </summary>
        public readonly string? SslState;
        /// <summary>
        /// SSL certificate thumbprint
        /// </summary>
        public readonly string? Thumbprint;
        /// <summary>
        /// Virtual IP address assigned to the hostname if IP based SSL is enabled.
        /// </summary>
        public readonly string VirtualIP;

        [OutputConstructor]
        private HostNameBindingResponsePropertiesResult(
            string? azureResourceName,

            string? azureResourceType,

            string? customHostNameDnsRecordType,

            string? domainId,

            string? hostNameType,

            string? siteName,

            string? sslState,

            string? thumbprint,

            string virtualIP)
        {
            AzureResourceName = azureResourceName;
            AzureResourceType = azureResourceType;
            CustomHostNameDnsRecordType = customHostNameDnsRecordType;
            DomainId = domainId;
            HostNameType = hostNameType;
            SiteName = siteName;
            SslState = sslState;
            Thumbprint = thumbprint;
            VirtualIP = virtualIP;
        }
    }
}