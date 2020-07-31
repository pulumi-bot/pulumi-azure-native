// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Compute.V20160330.Outputs
{

    [OutputType]
    public sealed class SshConfigurationResponseResult
    {
        /// <summary>
        /// The list of SSH public keys used to authenticate with linux based VMs.
        /// </summary>
        public readonly ImmutableArray<Outputs.SshPublicKeyResponseResult> PublicKeys;

        [OutputConstructor]
        private SshConfigurationResponseResult(ImmutableArray<Outputs.SshPublicKeyResponseResult> publicKeys)
        {
            PublicKeys = publicKeys;
        }
    }
}