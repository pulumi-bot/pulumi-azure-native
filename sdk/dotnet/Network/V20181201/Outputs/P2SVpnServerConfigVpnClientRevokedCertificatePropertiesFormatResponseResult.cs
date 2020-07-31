// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20181201.Outputs
{

    [OutputType]
    public sealed class P2SVpnServerConfigVpnClientRevokedCertificatePropertiesFormatResponseResult
    {
        /// <summary>
        /// The provisioning state of the VPN client revoked certificate resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// The revoked VPN client certificate thumbprint.
        /// </summary>
        public readonly string? Thumbprint;

        [OutputConstructor]
        private P2SVpnServerConfigVpnClientRevokedCertificatePropertiesFormatResponseResult(
            string provisioningState,

            string? thumbprint)
        {
            ProvisioningState = provisioningState;
            Thumbprint = thumbprint;
        }
    }
}