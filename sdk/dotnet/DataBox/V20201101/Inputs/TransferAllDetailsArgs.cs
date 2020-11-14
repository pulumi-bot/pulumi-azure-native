// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.DataBox.V20201101.Inputs
{

    /// <summary>
    /// Details to transfer all data.
    /// </summary>
    public sealed class TransferAllDetailsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Type of the account of data
        /// </summary>
        [Input("dataAccountType", required: true)]
        public Input<string> DataAccountType { get; set; } = null!;

        /// <summary>
        /// To indicate if all Azure blobs have to be transferred
        /// </summary>
        [Input("transferAllBlobs")]
        public Input<bool>? TransferAllBlobs { get; set; }

        /// <summary>
        /// To indicate if all Azure Files have to be transferred
        /// </summary>
        [Input("transferAllFiles")]
        public Input<bool>? TransferAllFiles { get; set; }

        public TransferAllDetailsArgs()
        {
        }
    }
}
