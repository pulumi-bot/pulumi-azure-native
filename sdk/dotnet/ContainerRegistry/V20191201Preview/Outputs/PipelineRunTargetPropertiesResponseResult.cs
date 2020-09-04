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
    public sealed class PipelineRunTargetPropertiesResponseResult
    {
        /// <summary>
        /// The name of the target.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The type of the target.
        /// </summary>
        public readonly string? Type;

        [OutputConstructor]
        private PipelineRunTargetPropertiesResponseResult(
            string? name,

            string? type)
        {
            Name = name;
            Type = type;
        }
    }
}
