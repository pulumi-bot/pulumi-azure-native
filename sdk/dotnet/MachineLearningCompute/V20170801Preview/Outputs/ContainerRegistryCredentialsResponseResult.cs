// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.MachineLearningCompute.V20170801Preview.Outputs
{

    [OutputType]
    public sealed class ContainerRegistryCredentialsResponseResult
    {
        /// <summary>
        /// The ACR login server name. User name is the first part of the FQDN.
        /// </summary>
        public readonly string LoginServer;
        /// <summary>
        /// The ACR primary password.
        /// </summary>
        public readonly string Password;
        /// <summary>
        /// The ACR secondary password.
        /// </summary>
        public readonly string Password2;
        /// <summary>
        /// The ACR login username.
        /// </summary>
        public readonly string Username;

        [OutputConstructor]
        private ContainerRegistryCredentialsResponseResult(
            string loginServer,

            string password,

            string password2,

            string username)
        {
            LoginServer = loginServer;
            Password = password;
            Password2 = password2;
            Username = username;
        }
    }
}
