// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Media.V20151001.Outputs
{

    [OutputType]
    public sealed class MediaServicePropertiesResponseResult
    {
        /// <summary>
        /// Read-only property that lists the Media Services REST API endpoints for this resource. If supplied on a PUT or PATCH, the value will be ignored.
        /// </summary>
        public readonly ImmutableArray<Outputs.ApiEndpointResponseResult> ApiEndpoints;
        /// <summary>
        /// The storage accounts for this resource.
        /// </summary>
        public readonly ImmutableArray<Outputs.StorageAccountResponseResult> StorageAccounts;

        [OutputConstructor]
        private MediaServicePropertiesResponseResult(
            ImmutableArray<Outputs.ApiEndpointResponseResult> apiEndpoints,

            ImmutableArray<Outputs.StorageAccountResponseResult> storageAccounts)
        {
            ApiEndpoints = apiEndpoints;
            StorageAccounts = storageAccounts;
        }
    }
}