// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Compute.V20160430Preview.Outputs
{

    [OutputType]
    public sealed class VirtualMachineScaleSetManagedDiskParametersResponseResult
    {
        /// <summary>
        /// Specifies the storage account type for the managed disk. Possible values are: Standard_LRS or Premium_LRS.
        /// </summary>
        public readonly string? StorageAccountType;

        [OutputConstructor]
        private VirtualMachineScaleSetManagedDiskParametersResponseResult(string? storageAccountType)
        {
            StorageAccountType = storageAccountType;
        }
    }
}
