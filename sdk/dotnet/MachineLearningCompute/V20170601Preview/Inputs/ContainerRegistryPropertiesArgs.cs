// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.MachineLearningCompute.V20170601Preview.Inputs
{

    /// <summary>
    /// Properties of Azure Container Registry.
    /// </summary>
    public sealed class ContainerRegistryPropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// ARM resource ID of the Azure Container Registry used to store Docker images for web services in the cluster. If not provided one will be created. This cannot be changed once the cluster is created.
        /// </summary>
        [Input("resourceId")]
        public Input<string>? ResourceId { get; set; }

        public ContainerRegistryPropertiesArgs()
        {
        }
    }
}
