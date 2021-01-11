// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.RecoveryServices.Latest.Inputs
{

    /// <summary>
    /// The properties of the Key Vault which hosts CMK
    /// </summary>
    public sealed class CmkKeyVaultPropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The key uri of the Customer Managed Key
        /// </summary>
        [Input("keyUri")]
        public Input<string>? KeyUri { get; set; }

        public CmkKeyVaultPropertiesArgs()
        {
        }
    }
}
