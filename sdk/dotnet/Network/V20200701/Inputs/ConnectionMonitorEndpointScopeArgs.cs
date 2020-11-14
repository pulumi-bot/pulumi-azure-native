// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Network.V20200701.Inputs
{

    /// <summary>
    /// Describes the connection monitor endpoint scope.
    /// </summary>
    public sealed class ConnectionMonitorEndpointScopeArgs : Pulumi.ResourceArgs
    {
        [Input("exclude")]
        private InputList<Inputs.ConnectionMonitorEndpointScopeItemArgs>? _exclude;

        /// <summary>
        /// List of items which needs to be excluded from the endpoint scope.
        /// </summary>
        public InputList<Inputs.ConnectionMonitorEndpointScopeItemArgs> Exclude
        {
            get => _exclude ?? (_exclude = new InputList<Inputs.ConnectionMonitorEndpointScopeItemArgs>());
            set => _exclude = value;
        }

        [Input("include")]
        private InputList<Inputs.ConnectionMonitorEndpointScopeItemArgs>? _include;

        /// <summary>
        /// List of items which needs to be included to the endpoint scope.
        /// </summary>
        public InputList<Inputs.ConnectionMonitorEndpointScopeItemArgs> Include
        {
            get => _include ?? (_include = new InputList<Inputs.ConnectionMonitorEndpointScopeItemArgs>());
            set => _include = value;
        }

        public ConnectionMonitorEndpointScopeArgs()
        {
        }
    }
}
