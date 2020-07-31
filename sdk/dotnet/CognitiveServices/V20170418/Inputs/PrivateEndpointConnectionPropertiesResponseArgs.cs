// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.CognitiveServices.V20170418.Inputs
{

    /// <summary>
    /// Properties of the PrivateEndpointConnectProperties.
    /// </summary>
    public sealed class PrivateEndpointConnectionPropertiesResponseArgs : Pulumi.InvokeArgs
    {
        [Input("groupIds")]
        private List<string>? _groupIds;

        /// <summary>
        /// The private link resource group ids.
        /// </summary>
        public List<string> GroupIds
        {
            get => _groupIds ?? (_groupIds = new List<string>());
            set => _groupIds = value;
        }

        /// <summary>
        /// The resource of private end point.
        /// </summary>
        [Input("privateEndpoint")]
        public Inputs.PrivateEndpointResponseArgs? PrivateEndpoint { get; set; }

        /// <summary>
        /// A collection of information about the state of the connection between service consumer and provider.
        /// </summary>
        [Input("privateLinkServiceConnectionState", required: true)]
        public Inputs.PrivateLinkServiceConnectionStateResponseArgs PrivateLinkServiceConnectionState { get; set; } = null!;

        public PrivateEndpointConnectionPropertiesResponseArgs()
        {
        }
    }
}