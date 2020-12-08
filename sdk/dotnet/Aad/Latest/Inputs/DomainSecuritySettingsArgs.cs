// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Aad.Latest.Inputs
{

    /// <summary>
    /// Domain Security Settings
    /// </summary>
    public sealed class DomainSecuritySettingsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A flag to determine whether or not NtlmV1 is enabled or disabled.
        /// </summary>
        [Input("ntlmV1")]
        public InputUnion<string, Pulumi.AzureNextGen.Aad.Latest.NtlmV1>? NtlmV1 { get; set; }

        /// <summary>
        /// A flag to determine whether or not SyncKerberosPasswords is enabled or disabled.
        /// </summary>
        [Input("syncKerberosPasswords")]
        public InputUnion<string, Pulumi.AzureNextGen.Aad.Latest.SyncKerberosPasswords>? SyncKerberosPasswords { get; set; }

        /// <summary>
        /// A flag to determine whether or not SyncNtlmPasswords is enabled or disabled.
        /// </summary>
        [Input("syncNtlmPasswords")]
        public InputUnion<string, Pulumi.AzureNextGen.Aad.Latest.SyncNtlmPasswords>? SyncNtlmPasswords { get; set; }

        /// <summary>
        /// A flag to determine whether or not SyncOnPremPasswords is enabled or disabled.
        /// </summary>
        [Input("syncOnPremPasswords")]
        public InputUnion<string, Pulumi.AzureNextGen.Aad.Latest.SyncOnPremPasswords>? SyncOnPremPasswords { get; set; }

        /// <summary>
        /// A flag to determine whether or not TlsV1 is enabled or disabled.
        /// </summary>
        [Input("tlsV1")]
        public InputUnion<string, Pulumi.AzureNextGen.Aad.Latest.TlsV1>? TlsV1 { get; set; }

        public DomainSecuritySettingsArgs()
        {
        }
    }
}
