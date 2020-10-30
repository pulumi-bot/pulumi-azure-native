// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.DataBox.V20201101.Outputs
{

    [OutputType]
    public sealed class DataImportDetailsResponse
    {
        /// <summary>
        /// Account details of the data to be transferred
        /// </summary>
        public readonly Union<Outputs.ManagedDiskDetailsResponse, Outputs.StorageAccountDetailsResponse> AccountDetails;

        [OutputConstructor]
        private DataImportDetailsResponse(Union<Outputs.ManagedDiskDetailsResponse, Outputs.StorageAccountDetailsResponse> accountDetails)
        {
            AccountDetails = accountDetails;
        }
    }
}
