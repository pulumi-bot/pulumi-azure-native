// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20200501.Inputs
{

    /// <summary>
    /// Load balancing settings for a backend pool
    /// </summary>
    public sealed class LoadBalancingSettingsModelArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The additional latency in milliseconds for probes to fall into the lowest latency bucket
        /// </summary>
        [Input("additionalLatencyMilliseconds")]
        public Input<int>? AdditionalLatencyMilliseconds { get; set; }

        /// <summary>
        /// Resource ID.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Resource name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Resource status.
        /// </summary>
        [Input("resourceState")]
        public Input<string>? ResourceState { get; set; }

        /// <summary>
        /// The number of samples to consider for load balancing decisions
        /// </summary>
        [Input("sampleSize")]
        public Input<int>? SampleSize { get; set; }

        /// <summary>
        /// The number of samples within the sample period that must succeed
        /// </summary>
        [Input("successfulSamplesRequired")]
        public Input<int>? SuccessfulSamplesRequired { get; set; }

        public LoadBalancingSettingsModelArgs()
        {
        }
    }
}
