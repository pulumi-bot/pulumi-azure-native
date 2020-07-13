// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ApiManagement.Outputs
{

    [OutputType]
    public sealed class ApiReleaseContractPropertiesResponseResult
    {
        /// <summary>
        /// Identifier of the API the release belongs to.
        /// </summary>
        public readonly string? ApiId;
        /// <summary>
        /// The time the API was released. The date conforms to the following format: yyyy-MM-ddTHH:mm:ssZ as specified by the ISO 8601 standard.
        /// </summary>
        public readonly string CreatedDateTime;
        /// <summary>
        /// Release Notes
        /// </summary>
        public readonly string? Notes;
        /// <summary>
        /// The time the API release was updated.
        /// </summary>
        public readonly string UpdatedDateTime;

        [OutputConstructor]
        private ApiReleaseContractPropertiesResponseResult(
            string? apiId,

            string createdDateTime,

            string? notes,

            string updatedDateTime)
        {
            ApiId = apiId;
            CreatedDateTime = createdDateTime;
            Notes = notes;
            UpdatedDateTime = updatedDateTime;
        }
    }
}