// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.ChangeAnalysis.V20200401Preview.Inputs
{

    /// <summary>
    /// Settings of change notification configuration for a subscription.
    /// </summary>
    public sealed class NotificationSettingsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The state of notifications feature.
        /// </summary>
        [Input("activationState")]
        public InputUnion<string, Pulumi.AzureNextGen.ChangeAnalysis.V20200401Preview.NotificationsState>? ActivationState { get; set; }

        /// <summary>
        /// Configuration properties of an Azure Monitor workspace that receives change notifications.
        /// </summary>
        [Input("azureMonitorWorkspaceProperties")]
        public Input<Inputs.AzureMonitorWorkspacePropertiesArgs>? AzureMonitorWorkspaceProperties { get; set; }

        public NotificationSettingsArgs()
        {
        }
    }
}
