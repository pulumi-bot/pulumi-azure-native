// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Media.V20180330Preview.Inputs
{

    /// <summary>
    /// Class to specify configurations of FairPlay in Streaming Policy
    /// </summary>
    public sealed class StreamingPolicyFairPlayConfigurationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// All license to be persistent or not
        /// </summary>
        [Input("allowPersistentLicense", required: true)]
        public Input<bool> AllowPersistentLicense { get; set; } = null!;

        /// <summary>
        /// The template for a customer service to deliver keys to end users.  Not needed when using Azure Media Services for issuing keys.
        /// </summary>
        [Input("customLicenseAcquisitionUrlTemplate")]
        public Input<string>? CustomLicenseAcquisitionUrlTemplate { get; set; }

        public StreamingPolicyFairPlayConfigurationArgs()
        {
        }
    }
}
