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
    /// Information about the container service backing the cluster
    /// </summary>
    public sealed class AcsClusterPropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The number of agent nodes in the Container Service. This can be changed to scale the cluster.
        /// </summary>
        [Input("agentCount")]
        public Input<int>? AgentCount { get; set; }

        /// <summary>
        /// The Azure VM size of the agent VM nodes. This cannot be changed once the cluster is created.
        /// </summary>
        [Input("agentVmSize")]
        public Input<string>? AgentVmSize { get; set; }

        /// <summary>
        /// Orchestrator specific properties
        /// </summary>
        [Input("orchestratorProperties", required: true)]
        public Input<Inputs.KubernetesClusterPropertiesArgs> OrchestratorProperties { get; set; } = null!;

        /// <summary>
        /// Type of orchestrator. It cannot be changed once the cluster is created.
        /// </summary>
        [Input("orchestratorType", required: true)]
        public Input<string> OrchestratorType { get; set; } = null!;

        [Input("systemServices")]
        private InputList<string>? _systemServices;

        /// <summary>
        /// The system services deployed to the cluster
        /// </summary>
        public InputList<string> SystemServices
        {
            get => _systemServices ?? (_systemServices = new InputList<string>());
            set => _systemServices = value;
        }

        public AcsClusterPropertiesArgs()
        {
        }
    }
}
