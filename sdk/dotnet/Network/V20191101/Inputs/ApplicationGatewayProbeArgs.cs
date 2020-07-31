// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20191101.Inputs
{

    /// <summary>
    /// Probe of the application gateway.
    /// </summary>
    public sealed class ApplicationGatewayProbeArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Resource ID.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Name of the probe that is unique within an Application Gateway.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Properties of the application gateway probe.
        /// </summary>
        [Input("properties")]
        public Input<Inputs.ApplicationGatewayProbePropertiesFormatArgs>? Properties { get; set; }

        public ApplicationGatewayProbeArgs()
        {
        }
    }
}