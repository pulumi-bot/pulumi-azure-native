// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearning.V20160501Preview.Outputs
{

    [OutputType]
    public sealed class GraphParameterLinkResponse
    {
        /// <summary>
        /// The graph node's identifier
        /// </summary>
        public readonly string NodeId;
        /// <summary>
        /// The identifier of the node parameter that the global parameter maps to.
        /// </summary>
        public readonly string ParameterKey;

        [OutputConstructor]
        private GraphParameterLinkResponse(
            string nodeId,

            string parameterKey)
        {
            NodeId = nodeId;
            ParameterKey = parameterKey;
        }
    }
}
