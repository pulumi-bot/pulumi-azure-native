// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Network.V20150504Preview.Outputs
{

    [OutputType]
    public sealed class ZonePropertiesResponse
    {
        /// <summary>
        /// Gets or sets the maximum number of record sets that can be created in this zone.
        /// </summary>
        public readonly double? MaxNumberOfRecordSets;
        /// <summary>
        /// The maximum number of records per record set that can be created in this DNS zone.  This is a read-only property and any attempt to set this value will be ignored.
        /// </summary>
        public readonly double MaxNumberOfRecordsPerRecordSet;
        /// <summary>
        /// Gets or sets the current number of record sets in this zone.
        /// </summary>
        public readonly double? NumberOfRecordSets;

        [OutputConstructor]
        private ZonePropertiesResponse(
            double? maxNumberOfRecordSets,

            double maxNumberOfRecordsPerRecordSet,

            double? numberOfRecordSets)
        {
            MaxNumberOfRecordSets = maxNumberOfRecordSets;
            MaxNumberOfRecordsPerRecordSet = maxNumberOfRecordsPerRecordSet;
            NumberOfRecordSets = numberOfRecordSets;
        }
    }
}
