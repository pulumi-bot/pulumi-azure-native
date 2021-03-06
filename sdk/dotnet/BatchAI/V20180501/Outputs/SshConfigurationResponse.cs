// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.BatchAI.V20180501.Outputs
{

    [OutputType]
    public sealed class SshConfigurationResponse
    {
        /// <summary>
        /// List of source IP ranges to allow SSH connection from. The default value is '*' (all source IPs are allowed). Maximum number of IP ranges that can be specified is 400.
        /// </summary>
        public readonly ImmutableArray<string> PublicIPsToAllow;
        /// <summary>
        /// Settings for administrator user account to be created on a node. The account can be used to establish SSH connection to the node.
        /// </summary>
        public readonly Outputs.UserAccountSettingsResponse UserAccountSettings;

        [OutputConstructor]
        private SshConfigurationResponse(
            ImmutableArray<string> publicIPsToAllow,

            Outputs.UserAccountSettingsResponse userAccountSettings)
        {
            PublicIPsToAllow = publicIPsToAllow;
            UserAccountSettings = userAccountSettings;
        }
    }
}
