// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DevTestLab.V20180915.Outputs
{

    [OutputType]
    public sealed class DataDiskStorageTypeInfoResponseResult
    {
        /// <summary>
        /// Disk Lun
        /// </summary>
        public readonly string? Lun;
        /// <summary>
        /// Disk Storage Type
        /// </summary>
        public readonly string? StorageType;

        [OutputConstructor]
        private DataDiskStorageTypeInfoResponseResult(
            string? lun,

            string? storageType)
        {
            Lun = lun;
            StorageType = storageType;
        }
    }
}