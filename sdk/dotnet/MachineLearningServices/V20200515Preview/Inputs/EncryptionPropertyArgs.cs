// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.MachineLearningServices.V20200515Preview.Inputs
{

    public sealed class EncryptionPropertyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Customer Key vault properties.
        /// </summary>
        [Input("keyVaultProperties", required: true)]
        public Input<Inputs.KeyVaultPropertiesArgs> KeyVaultProperties { get; set; } = null!;

        /// <summary>
        /// Indicates whether or not the encryption is enabled for the workspace.
        /// </summary>
        [Input("status", required: true)]
        public InputUnion<string, Pulumi.AzureNextGen.MachineLearningServices.V20200515Preview.EncryptionStatus> Status { get; set; } = null!;

        public EncryptionPropertyArgs()
        {
        }
    }
}
