// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Batch.V20210101.Inputs
{

    /// <summary>
    /// The properties related to the auto-storage account.
    /// </summary>
    public sealed class AutoStorageBasePropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The resource ID of the storage account to be used for auto-storage account.
        /// </summary>
        [Input("storageAccountId", required: true)]
        public Input<string> StorageAccountId { get; set; } = null!;

        public AutoStorageBasePropertiesArgs()
        {
        }
    }
}
