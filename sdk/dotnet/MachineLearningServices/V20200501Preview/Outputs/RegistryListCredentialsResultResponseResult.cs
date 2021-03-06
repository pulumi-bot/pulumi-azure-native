// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.MachineLearningServices.V20200501Preview.Outputs
{

    [OutputType]
    public sealed class RegistryListCredentialsResultResponseResult
    {
        public readonly string Location;
        public readonly ImmutableArray<Outputs.PasswordResponseResult> Passwords;
        public readonly string Username;

        [OutputConstructor]
        private RegistryListCredentialsResultResponseResult(
            string location,

            ImmutableArray<Outputs.PasswordResponseResult> passwords,

            string username)
        {
            Location = location;
            Passwords = passwords;
            Username = username;
        }
    }
}
