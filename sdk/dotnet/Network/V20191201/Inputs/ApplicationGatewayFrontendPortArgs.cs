// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20191201.Inputs
{

    /// <summary>
    /// Frontend port of an application gateway.
    /// </summary>
    public sealed class ApplicationGatewayFrontendPortArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Resource ID.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Name of the frontend port that is unique within an Application Gateway.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Frontend port.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        public ApplicationGatewayFrontendPortArgs()
        {
        }
    }
}
