// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Cdn.V20191231.Inputs
{

    /// <summary>
    /// The JSON object that contains the properties to determine origin health using real requests/responses.
    /// </summary>
    public sealed class ResponseBasedOriginErrorDetectionParametersArgs : Pulumi.ResourceArgs
    {
        [Input("httpErrorRanges")]
        private InputList<Inputs.HttpErrorRangeParametersArgs>? _httpErrorRanges;

        /// <summary>
        /// The list of Http status code ranges that are considered as server errors for origin and it is marked as unhealthy.
        /// </summary>
        public InputList<Inputs.HttpErrorRangeParametersArgs> HttpErrorRanges
        {
            get => _httpErrorRanges ?? (_httpErrorRanges = new InputList<Inputs.HttpErrorRangeParametersArgs>());
            set => _httpErrorRanges = value;
        }

        /// <summary>
        /// Type of response errors for real user requests for which origin will be deemed unhealthy
        /// </summary>
        [Input("responseBasedDetectedErrorTypes")]
        public Input<Pulumi.AzureNextGen.Cdn.V20191231.ResponseBasedDetectedErrorTypes>? ResponseBasedDetectedErrorTypes { get; set; }

        /// <summary>
        /// The percentage of failed requests in the sample where failover should trigger.
        /// </summary>
        [Input("responseBasedFailoverThresholdPercentage")]
        public Input<int>? ResponseBasedFailoverThresholdPercentage { get; set; }

        public ResponseBasedOriginErrorDetectionParametersArgs()
        {
        }
    }
}
