// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Storage.V20200801Preview.Inputs
{

    /// <summary>
    /// Management policy action for base blob.
    /// </summary>
    public sealed class ManagementPolicyBaseBlobArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The function to delete the blob
        /// </summary>
        [Input("delete")]
        public Input<Inputs.DateAfterModificationArgs>? Delete { get; set; }

        /// <summary>
        /// This property enables auto tiering of a blob from cool to hot on a blob access. This property requires tierToCool.daysAfterLastAccessTimeGreaterThan.
        /// </summary>
        [Input("enableAutoTierToHotFromCool")]
        public Input<bool>? EnableAutoTierToHotFromCool { get; set; }

        /// <summary>
        /// The function to tier blobs to archive storage. Support blobs currently at Hot or Cool tier
        /// </summary>
        [Input("tierToArchive")]
        public Input<Inputs.DateAfterModificationArgs>? TierToArchive { get; set; }

        /// <summary>
        /// The function to tier blobs to cool storage. Support blobs currently at Hot tier
        /// </summary>
        [Input("tierToCool")]
        public Input<Inputs.DateAfterModificationArgs>? TierToCool { get; set; }

        public ManagementPolicyBaseBlobArgs()
        {
        }
    }
}
