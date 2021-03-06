// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Blueprint.Outputs
{

    [OutputType]
    public sealed class BlueprintStatusResponse
    {
        /// <summary>
        /// Last modified time of this blueprint definition.
        /// </summary>
        public readonly string LastModified;
        /// <summary>
        /// Creation time of this blueprint definition.
        /// </summary>
        public readonly string TimeCreated;

        [OutputConstructor]
        private BlueprintStatusResponse(
            string lastModified,

            string timeCreated)
        {
            LastModified = lastModified;
            TimeCreated = timeCreated;
        }
    }
}
