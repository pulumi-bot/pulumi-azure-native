// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Kubernetes.V20210401Preview.Outputs
{

    [OutputType]
    public sealed class CredentialResultResponseResult
    {
        /// <summary>
        /// The name of the credential.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Base64-encoded Kubernetes configuration file.
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private CredentialResultResponseResult(
            string name,

            string value)
        {
            Name = name;
            Value = value;
        }
    }
}
