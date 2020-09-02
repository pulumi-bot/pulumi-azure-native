// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Search.V20191001Preview.Inputs
{

    /// <summary>
    /// Describes the properties of an existing Private Endpoint connection to the Azure Cognitive Search service.
    /// </summary>
    public sealed class PrivateEndpointConnectionPropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The private endpoint resource from Microsoft.Network provider.
        /// </summary>
        [Input("privateEndpoint")]
        public Input<Inputs.PrivateEndpointConnectionPropertiesPrivateEndpointArgs>? PrivateEndpoint { get; set; }

        /// <summary>
        /// Describes the current state of an existing Private Link Service connection to the Azure Private Endpoint.
        /// </summary>
        [Input("privateLinkServiceConnectionState")]
        public Input<Inputs.PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionStateArgs>? PrivateLinkServiceConnectionState { get; set; }

        public PrivateEndpointConnectionPropertiesArgs()
        {
        }
    }
}
