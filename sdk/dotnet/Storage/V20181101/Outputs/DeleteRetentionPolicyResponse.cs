// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Storage.V20181101.Outputs
{

    [OutputType]
    public sealed class DeleteRetentionPolicyResponse
    {
        /// <summary>
        /// Indicates the number of days that the deleted blob should be retained. The minimum specified value can be 1 and the maximum value can be 365.
        /// </summary>
        public readonly int? Days;
        /// <summary>
        /// Indicates whether DeleteRetentionPolicy is enabled for the Blob service.
        /// </summary>
        public readonly bool? Enabled;

        [OutputConstructor]
        private DeleteRetentionPolicyResponse(
            int? days,

            bool? enabled)
        {
            Days = days;
            Enabled = enabled;
        }
    }
}
