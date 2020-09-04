// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Blueprint.V20181101Preview.Outputs
{

    [OutputType]
    public sealed class AssignmentStatusResponseResult
    {
        /// <summary>
        /// Last modified time of this blueprint definition.
        /// </summary>
        public readonly string LastModified;
        /// <summary>
        /// List of resources that were created by the blueprint assignment.
        /// </summary>
        public readonly ImmutableArray<string> ManagedResources;
        /// <summary>
        /// Creation time of this blueprint definition.
        /// </summary>
        public readonly string TimeCreated;

        [OutputConstructor]
        private AssignmentStatusResponseResult(
            string lastModified,

            ImmutableArray<string> managedResources,

            string timeCreated)
        {
            LastModified = lastModified;
            ManagedResources = managedResources;
            TimeCreated = timeCreated;
        }
    }
}
