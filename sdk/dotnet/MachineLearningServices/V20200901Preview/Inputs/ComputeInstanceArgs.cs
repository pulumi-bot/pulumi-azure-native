// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.MachineLearningServices.V20200901Preview.Inputs
{

    /// <summary>
    /// An Azure Machine Learning compute instance.
    /// </summary>
    public sealed class ComputeInstanceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Location for the underlying compute
        /// </summary>
        [Input("computeLocation")]
        public Input<string>? ComputeLocation { get; set; }

        /// <summary>
        /// The type of compute
        /// </summary>
        [Input("computeType", required: true)]
        public Input<string> ComputeType { get; set; } = null!;

        /// <summary>
        /// The description of the Machine Learning compute.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Compute Instance properties
        /// </summary>
        [Input("properties")]
        public Input<Inputs.ComputeInstancePropertiesArgs>? Properties { get; set; }

        /// <summary>
        /// ARM resource id of the underlying compute
        /// </summary>
        [Input("resourceId")]
        public Input<string>? ResourceId { get; set; }

        public ComputeInstanceArgs()
        {
        }
    }
}
