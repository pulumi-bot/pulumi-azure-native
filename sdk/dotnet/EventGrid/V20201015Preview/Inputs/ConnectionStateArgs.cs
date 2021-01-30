// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.EventGrid.V20201015Preview.Inputs
{

    /// <summary>
    /// ConnectionState information.
    /// </summary>
    public sealed class ConnectionStateArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Actions required (if any).
        /// </summary>
        [Input("actionsRequired")]
        public Input<string>? ActionsRequired { get; set; }

        /// <summary>
        /// Description of the connection state.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Status of the connection.
        /// </summary>
        [Input("status")]
        public InputUnion<string, Pulumi.AzureNextGen.EventGrid.V20201015Preview.PersistedConnectionStatus>? Status { get; set; }

        public ConnectionStateArgs()
        {
        }
    }
}
