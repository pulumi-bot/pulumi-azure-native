// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20190901.Inputs
{

    /// <summary>
    /// UrlPathMaps give a url path to the backend mapping information for PathBasedRouting.
    /// </summary>
    public sealed class ApplicationGatewayUrlPathMapArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Resource ID.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Name of the URL path map that is unique within an Application Gateway.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Properties of the application gateway URL path map.
        /// </summary>
        [Input("properties")]
        public Input<Inputs.ApplicationGatewayUrlPathMapPropertiesFormatArgs>? Properties { get; set; }

        public ApplicationGatewayUrlPathMapArgs()
        {
        }
    }
}