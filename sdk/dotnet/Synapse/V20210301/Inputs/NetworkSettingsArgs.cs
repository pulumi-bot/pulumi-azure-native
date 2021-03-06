// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Synapse.V20210301.Inputs
{

    /// <summary>
    /// Network Settings
    /// </summary>
    public sealed class NetworkSettingsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enable or Disable pubic network access to workspace
        /// </summary>
        [Input("publicNetworkAccess")]
        public InputUnion<string, Pulumi.AzureNative.Synapse.V20210301.WorkspacePublicNetworkAccess>? PublicNetworkAccess { get; set; }

        public NetworkSettingsArgs()
        {
        }
    }
}
