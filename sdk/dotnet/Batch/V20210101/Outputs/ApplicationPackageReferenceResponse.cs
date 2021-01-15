// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Batch.V20210101.Outputs
{

    [OutputType]
    public sealed class ApplicationPackageReferenceResponse
    {
        public readonly string Id;
        /// <summary>
        /// If this is omitted, and no default version is specified for this application, the request fails with the error code InvalidApplicationPackageReferences. If you are calling the REST API directly, the HTTP status code is 409.
        /// </summary>
        public readonly string? Version;

        [OutputConstructor]
        private ApplicationPackageReferenceResponse(
            string id,

            string? version)
        {
            Id = id;
            Version = version;
        }
    }
}
