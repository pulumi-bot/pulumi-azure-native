// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Media.V20180701.Inputs
{

    /// <summary>
    /// The Asset properties.
    /// </summary>
    public sealed class AssetPropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The alternate ID of the Asset.
        /// </summary>
        [Input("alternateId")]
        public Input<string>? AlternateId { get; set; }

        /// <summary>
        /// The name of the asset blob container.
        /// </summary>
        [Input("container")]
        public Input<string>? Container { get; set; }

        /// <summary>
        /// The Asset description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The name of the storage account.
        /// </summary>
        [Input("storageAccountName")]
        public Input<string>? StorageAccountName { get; set; }

        public AssetPropertiesArgs()
        {
        }
    }
}