// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Storage.Outputs
{

    [OutputType]
    public sealed class ManagementPolicyPropertiesResponse
    {
        /// <summary>
        /// Returns the date and time the ManagementPolicies was last modified.
        /// </summary>
        public readonly string LastModifiedTime;
        /// <summary>
        /// The Storage Account ManagementPolicy, in JSON format. See more details in: https://docs.microsoft.com/en-us/azure/storage/common/storage-lifecycle-managment-concepts.
        /// </summary>
        public readonly Outputs.ManagementPolicySchemaResponse Policy;

        [OutputConstructor]
        private ManagementPolicyPropertiesResponse(
            string lastModifiedTime,

            Outputs.ManagementPolicySchemaResponse policy)
        {
            LastModifiedTime = lastModifiedTime;
            Policy = policy;
        }
    }
}