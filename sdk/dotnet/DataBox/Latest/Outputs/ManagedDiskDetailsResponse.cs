// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DataBox.Latest.Outputs
{

    [OutputType]
    public sealed class ManagedDiskDetailsResponse
    {
        /// <summary>
        /// Account Type of the data to be transferred.
        /// Expected value is 'ManagedDisk'.
        /// </summary>
        public readonly string DataAccountType;
        /// <summary>
        /// Resource Group Id of the compute disks.
        /// </summary>
        public readonly string ResourceGroupId;
        /// <summary>
        /// Password for all the shares to be created on the device. Should not be passed for TransferType:ExportFromAzure jobs. If this is not passed, the service will generate password itself. This will not be returned in Get Call. Password Requirements :  Password must be minimum of 12 and maximum of 64 characters. Password must have at least one uppercase alphabet, one number and one special character. Password cannot have the following characters : IilLoO0 Password can have only alphabets, numbers and these characters : @#\-$%^!+=;:_()]+
        /// </summary>
        public readonly string? SharePassword;
        /// <summary>
        /// Resource Id of the storage account that can be used to copy the vhd for staging.
        /// </summary>
        public readonly string StagingStorageAccountId;

        [OutputConstructor]
        private ManagedDiskDetailsResponse(
            string dataAccountType,

            string resourceGroupId,

            string? sharePassword,

            string stagingStorageAccountId)
        {
            DataAccountType = dataAccountType;
            ResourceGroupId = resourceGroupId;
            SharePassword = sharePassword;
            StagingStorageAccountId = stagingStorageAccountId;
        }
    }
}
