// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ContainerService.Inputs
{

    /// <summary>
    /// Information about a service principal identity for the cluster to use for manipulating Azure APIs. Either secret or keyVaultSecretRef must be specified.
    /// </summary>
    public sealed class ContainerServiceServicePrincipalProfileArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID for the service principal.
        /// </summary>
        [Input("clientId", required: true)]
        public Input<string> ClientId { get; set; } = null!;

        /// <summary>
        /// Reference to a secret stored in Azure Key Vault.
        /// </summary>
        [Input("keyVaultSecretRef")]
        public Input<Inputs.KeyVaultSecretRefArgs>? KeyVaultSecretRef { get; set; }

        /// <summary>
        /// The secret password associated with the service principal in plain text.
        /// </summary>
        [Input("secret")]
        public Input<string>? Secret { get; set; }

        public ContainerServiceServicePrincipalProfileArgs()
        {
        }
    }
}