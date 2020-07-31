// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ContainerService.V20160930.Outputs
{

    [OutputType]
    public sealed class ContainerServiceOrchestratorProfileResponseResult
    {
        /// <summary>
        /// The orchestrator to use to manage container service cluster resources. Valid values are Swarm, DCOS, and Custom.
        /// </summary>
        public readonly string OrchestratorType;

        [OutputConstructor]
        private ContainerServiceOrchestratorProfileResponseResult(string orchestratorType)
        {
            OrchestratorType = orchestratorType;
        }
    }
}