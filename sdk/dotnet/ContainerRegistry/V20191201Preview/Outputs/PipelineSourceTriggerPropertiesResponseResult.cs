// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ContainerRegistry.V20191201Preview.Outputs
{

    [OutputType]
    public sealed class PipelineSourceTriggerPropertiesResponseResult
    {
        /// <summary>
        /// The current status of the source trigger.
        /// </summary>
        public readonly string Status;

        [OutputConstructor]
        private PipelineSourceTriggerPropertiesResponseResult(string status)
        {
            Status = status;
        }
    }
}
