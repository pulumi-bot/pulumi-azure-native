// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Compute.V20180601.Inputs
{

    /// <summary>
    /// The disks sku name. Can be Standard_LRS, Premium_LRS, StandardSSD_LRS, or UltraSSD_LRS.
    /// </summary>
    public sealed class DiskSkuArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The sku name.
        /// </summary>
        [Input("name")]
        public InputUnion<string, Pulumi.AzureNextGen.Compute.V20180601.DiskStorageAccountTypes>? Name { get; set; }

        public DiskSkuArgs()
        {
        }
    }
}
