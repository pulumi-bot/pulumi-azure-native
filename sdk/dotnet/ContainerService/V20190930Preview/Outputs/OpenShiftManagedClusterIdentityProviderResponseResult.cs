// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ContainerService.V20190930Preview.Outputs
{

    [OutputType]
    public sealed class OpenShiftManagedClusterIdentityProviderResponseResult
    {
        /// <summary>
        /// Name of the provider.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Configuration of the provider.
        /// </summary>
        public readonly Outputs.OpenShiftManagedClusterBaseIdentityProviderResponseResult? Provider;

        [OutputConstructor]
        private OpenShiftManagedClusterIdentityProviderResponseResult(
            string? name,

            Outputs.OpenShiftManagedClusterBaseIdentityProviderResponseResult? provider)
        {
            Name = name;
            Provider = provider;
        }
    }
}
