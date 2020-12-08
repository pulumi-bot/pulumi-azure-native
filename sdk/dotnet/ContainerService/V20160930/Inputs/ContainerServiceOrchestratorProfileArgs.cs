// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.ContainerService.V20160930.Inputs
{

    /// <summary>
    /// Profile for the container service orchestrator.
    /// </summary>
    public sealed class ContainerServiceOrchestratorProfileArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The orchestrator to use to manage container service cluster resources. Valid values are Swarm, DCOS, and Custom.
        /// </summary>
        [Input("orchestratorType", required: true)]
        public Input<Pulumi.AzureNextGen.ContainerService.V20160930.ContainerServiceOchestratorTypes> OrchestratorType { get; set; } = null!;

        public ContainerServiceOrchestratorProfileArgs()
        {
        }
    }
}
