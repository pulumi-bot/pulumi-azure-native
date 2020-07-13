// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.HDInsight.Inputs
{

    /// <summary>
    /// Gets the application SSH endpoint
    /// </summary>
    public sealed class ApplicationGetEndpointArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The destination port to connect to.
        /// </summary>
        [Input("destinationPort")]
        public Input<int>? DestinationPort { get; set; }

        /// <summary>
        /// The location of the endpoint.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The public port to connect to.
        /// </summary>
        [Input("publicPort")]
        public Input<int>? PublicPort { get; set; }

        public ApplicationGetEndpointArgs()
        {
        }
    }
}