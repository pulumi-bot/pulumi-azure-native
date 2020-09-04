// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.OperationalInsights.V20200301Preview.Inputs
{

    /// <summary>
    /// Describes a storage account connection.
    /// </summary>
    public sealed class StorageAccountArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Azure Resource Manager ID of the storage account resource.
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        /// <summary>
        /// The storage account key.
        /// </summary>
        [Input("key", required: true)]
        public Input<string> Key { get; set; } = null!;

        public StorageAccountArgs()
        {
        }
    }
}
