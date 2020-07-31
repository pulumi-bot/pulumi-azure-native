// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Storage.V20170601.Inputs
{

    /// <summary>
    /// The parameters used to create the storage account.
    /// </summary>
    public sealed class StorageAccountPropertiesCreateParametersArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Required for storage accounts where kind = BlobStorage. The access tier used for billing.
        /// </summary>
        [Input("accessTier")]
        public Input<string>? AccessTier { get; set; }

        /// <summary>
        /// User domain assigned to the storage account. Name is the CNAME source. Only one custom domain is supported per storage account at this time. To clear the existing custom domain, use an empty string for the custom domain name property.
        /// </summary>
        [Input("customDomain")]
        public Input<Inputs.CustomDomainArgs>? CustomDomain { get; set; }

        /// <summary>
        /// Provides the encryption settings on the account. If left unspecified the account encryption settings will remain the same. The default setting is unencrypted.
        /// </summary>
        [Input("encryption")]
        public Input<Inputs.EncryptionArgs>? Encryption { get; set; }

        /// <summary>
        /// Network rule set
        /// </summary>
        [Input("networkAcls")]
        public Input<Inputs.NetworkRuleSetArgs>? NetworkAcls { get; set; }

        /// <summary>
        /// Allows https traffic only to storage service if sets to true.
        /// </summary>
        [Input("supportsHttpsTrafficOnly")]
        public Input<bool>? SupportsHttpsTrafficOnly { get; set; }

        public StorageAccountPropertiesCreateParametersArgs()
        {
        }
    }
}