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
    /// Map of filter type and the details to transfer all data. This field is required only if the TransferConfigurationType is given as TransferAll
    /// </summary>
    public sealed class TransferConfigurationTransferAllDetailsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Details to transfer all data.
        /// </summary>
        [Input("include")]
        public Input<Inputs.TransferAllDetailsArgs>? Include { get; set; }

        public TransferConfigurationTransferAllDetailsArgs()
        {
        }
    }
}
