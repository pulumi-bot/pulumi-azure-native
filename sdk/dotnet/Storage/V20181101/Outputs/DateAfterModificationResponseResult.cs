// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Storage.V20181101.Outputs
{

    [OutputType]
    public sealed class DateAfterModificationResponseResult
    {
        /// <summary>
        /// Integer value indicating the age in days after last modification
        /// </summary>
        public readonly int DaysAfterModificationGreaterThan;

        [OutputConstructor]
        private DateAfterModificationResponseResult(int daysAfterModificationGreaterThan)
        {
            DaysAfterModificationGreaterThan = daysAfterModificationGreaterThan;
        }
    }
}