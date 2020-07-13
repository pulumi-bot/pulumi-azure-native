// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.MachineLearning.Outputs
{

    [OutputType]
    public sealed class BlobLocationResponse
    {
        /// <summary>
        /// Access credentials for the blob, if applicable (e.g. blob specified by storage account connection string + blob URI)
        /// </summary>
        public readonly string? Credentials;
        /// <summary>
        /// The URI from which the blob is accessible from. For example, aml://abc for system assets or https://xyz for user assets or payload.
        /// </summary>
        public readonly string Uri;

        [OutputConstructor]
        private BlobLocationResponse(
            string? credentials,

            string uri)
        {
            Credentials = credentials;
            Uri = uri;
        }
    }
}