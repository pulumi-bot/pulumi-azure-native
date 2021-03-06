// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.BatchAI.V20180301.Outputs
{

    [OutputType]
    public sealed class ImageSourceRegistryResponse
    {
        /// <summary>
        /// Credentials to access a container image in a private repository.
        /// </summary>
        public readonly Outputs.PrivateRegistryCredentialsResponse? Credentials;
        public readonly string Image;
        public readonly string? ServerUrl;

        [OutputConstructor]
        private ImageSourceRegistryResponse(
            Outputs.PrivateRegistryCredentialsResponse? credentials,

            string image,

            string? serverUrl)
        {
            Credentials = credentials;
            Image = image;
            ServerUrl = serverUrl;
        }
    }
}
