// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.ContainerInstance.V20181001.Inputs
{

    /// <summary>
    /// Container group log analytics information.
    /// </summary>
    public sealed class LogAnalyticsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The log type to be used.
        /// </summary>
        [Input("logType")]
        public InputUnion<string, Pulumi.AzureNextGen.ContainerInstance.V20181001.LogAnalyticsLogType>? LogType { get; set; }

        [Input("metadata")]
        private InputMap<string>? _metadata;

        /// <summary>
        /// Metadata for log analytics.
        /// </summary>
        public InputMap<string> Metadata
        {
            get => _metadata ?? (_metadata = new InputMap<string>());
            set => _metadata = value;
        }

        /// <summary>
        /// The workspace id for log analytics
        /// </summary>
        [Input("workspaceId", required: true)]
        public Input<string> WorkspaceId { get; set; } = null!;

        /// <summary>
        /// The workspace key for log analytics
        /// </summary>
        [Input("workspaceKey", required: true)]
        public Input<string> WorkspaceKey { get; set; } = null!;

        public LogAnalyticsArgs()
        {
        }
    }
}
