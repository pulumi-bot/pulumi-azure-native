// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ContainerRegistry.V20190601Preview.Outputs
{

    [OutputType]
    public sealed class AgentPropertiesResponseResult
    {
        /// <summary>
        /// The CPU configuration in terms of number of cores required for the run.
        /// </summary>
        public readonly int? Cpu;

        [OutputConstructor]
        private AgentPropertiesResponseResult(int? cpu)
        {
            Cpu = cpu;
        }
    }
}
