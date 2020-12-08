// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Network.V20200601.Inputs
{

    /// <summary>
    /// Describes a connection monitor output destination.
    /// </summary>
    public sealed class ConnectionMonitorOutputArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Connection monitor output destination type. Currently, only "Workspace" is supported.
        /// </summary>
        [Input("type")]
        public InputUnion<string, Pulumi.AzureNextGen.Network.V20200601.OutputType>? Type { get; set; }

        /// <summary>
        /// Describes the settings for producing output into a log analytics workspace.
        /// </summary>
        [Input("workspaceSettings")]
        public Input<Inputs.ConnectionMonitorWorkspaceSettingsArgs>? WorkspaceSettings { get; set; }

        public ConnectionMonitorOutputArgs()
        {
        }
    }
}
