// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Storage.V20181101.Inputs
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
