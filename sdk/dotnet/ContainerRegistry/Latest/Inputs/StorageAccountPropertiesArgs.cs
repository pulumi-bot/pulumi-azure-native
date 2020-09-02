// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ContainerRegistry.Latest.Inputs
{

    /// <summary>
    /// The properties of a storage account for a container registry. Only applicable to Classic SKU.
    /// </summary>
    public sealed class StorageAccountPropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The resource ID of the storage account.
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        public StorageAccountPropertiesArgs()
        {
        }
    }
}