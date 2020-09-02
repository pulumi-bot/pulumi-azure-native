// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ServiceFabric.V20190601Preview.Inputs
{

    /// <summary>
    /// Port range details
    /// </summary>
    public sealed class EndpointRangeDescriptionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// End port of a range of ports
        /// </summary>
        [Input("endPort", required: true)]
        public Input<int> EndPort { get; set; } = null!;

        /// <summary>
        /// Starting port of a range of ports
        /// </summary>
        [Input("startPort", required: true)]
        public Input<int> StartPort { get; set; } = null!;

        public EndpointRangeDescriptionArgs()
        {
        }
    }
}
