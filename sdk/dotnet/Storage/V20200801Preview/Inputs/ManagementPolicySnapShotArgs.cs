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
    /// Management policy action for snapshot.
    /// </summary>
    public sealed class ManagementPolicySnapShotArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The function to delete the blob snapshot
        /// </summary>
        [Input("delete")]
        public Input<Inputs.DateAfterCreationArgs>? Delete { get; set; }

        /// <summary>
        /// The function to tier blob snapshot to archive storage. Support blob snapshot currently at Hot or Cool tier
        /// </summary>
        [Input("tierToArchive")]
        public Input<Inputs.DateAfterCreationArgs>? TierToArchive { get; set; }

        /// <summary>
        /// The function to tier blob snapshot to cool storage. Support blob snapshot currently at Hot tier
        /// </summary>
        [Input("tierToCool")]
        public Input<Inputs.DateAfterCreationArgs>? TierToCool { get; set; }

        public ManagementPolicySnapShotArgs()
        {
        }
    }
}
