// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Batch.Latest.Outputs
{

    [OutputType]
    public sealed class ContainerRegistryResponse
    {
        public readonly string Password;
        /// <summary>
        /// If omitted, the default is "docker.io".
        /// </summary>
        public readonly string? RegistryServer;
        public readonly string UserName;

        [OutputConstructor]
        private ContainerRegistryResponse(
            string password,

            string? registryServer,

            string userName)
        {
            Password = password;
            RegistryServer = registryServer;
            UserName = userName;
        }
    }
}
