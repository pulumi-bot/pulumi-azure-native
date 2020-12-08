// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Media.V20180601Preview.Inputs
{

    /// <summary>
    /// Class to specify one track property condition
    /// </summary>
    public sealed class TrackPropertyConditionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Track property condition operation
        /// </summary>
        [Input("operation", required: true)]
        public InputUnion<string, Pulumi.AzureNextGen.Media.V20180601Preview.TrackPropertyCompareOperation> Operation { get; set; } = null!;

        /// <summary>
        /// Track property type
        /// </summary>
        [Input("property", required: true)]
        public InputUnion<string, Pulumi.AzureNextGen.Media.V20180601Preview.TrackPropertyType> Property { get; set; } = null!;

        /// <summary>
        /// Track property value
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public TrackPropertyConditionArgs()
        {
        }
    }
}
