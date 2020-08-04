// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.HDInsight.V20180601Preview.Inputs
{

    /// <summary>
    /// The load-based autoscale request parameters
    /// </summary>
    public sealed class AutoscaleCapacityArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The maximum instance count of the cluster
        /// </summary>
        [Input("maxInstanceCount")]
        public Input<int>? MaxInstanceCount { get; set; }

        /// <summary>
        /// The minimum instance count of the cluster
        /// </summary>
        [Input("minInstanceCount")]
        public Input<int>? MinInstanceCount { get; set; }

        public AutoscaleCapacityArgs()
        {
        }
    }
}