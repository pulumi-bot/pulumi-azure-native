// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DataMigration.V20180315Preview.Inputs
{

    /// <summary>
    /// Defines the connection properties of a server
    /// </summary>
    public sealed class ConnectionInfoArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Password credential.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// Type of connection info
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// User name
        /// </summary>
        [Input("userName")]
        public Input<string>? UserName { get; set; }

        public ConnectionInfoArgs()
        {
        }
    }
}
