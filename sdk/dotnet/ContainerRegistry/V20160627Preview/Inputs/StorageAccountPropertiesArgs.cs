// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ContainerRegistry.V20160627Preview.Inputs
{

    /// <summary>
    /// The properties of a storage account for a container registry.
    /// </summary>
    public sealed class StorageAccountPropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The access key to the storage account.
        /// </summary>
        [Input("accessKey", required: true)]
        public Input<string> AccessKey { get; set; } = null!;

        /// <summary>
        /// The name of the storage account.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public StorageAccountPropertiesArgs()
        {
        }
    }
}
