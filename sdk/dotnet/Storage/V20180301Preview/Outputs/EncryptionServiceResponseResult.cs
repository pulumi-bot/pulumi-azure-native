// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Storage.V20180301Preview.Outputs
{

    [OutputType]
    public sealed class EncryptionServiceResponseResult
    {
        /// <summary>
        /// A boolean indicating whether or not the service encrypts the data as it is stored.
        /// </summary>
        public readonly bool? Enabled;
        /// <summary>
        /// Gets a rough estimate of the date/time when the encryption was last enabled by the user. Only returned when encryption is enabled. There might be some unencrypted blobs which were written after this time, as it is just a rough estimate.
        /// </summary>
        public readonly string LastEnabledTime;

        [OutputConstructor]
        private EncryptionServiceResponseResult(
            bool? enabled,

            string lastEnabledTime)
        {
            Enabled = enabled;
            LastEnabledTime = lastEnabledTime;
        }
    }
}
