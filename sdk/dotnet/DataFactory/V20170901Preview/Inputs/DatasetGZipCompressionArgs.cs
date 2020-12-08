// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.DataFactory.V20170901Preview.Inputs
{

    /// <summary>
    /// The GZip compression method used on a dataset.
    /// </summary>
    public sealed class DatasetGZipCompressionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The GZip compression level.
        /// </summary>
        [Input("level")]
        public InputUnion<string, Pulumi.AzureNextGen.DataFactory.V20170901Preview.DatasetCompressionLevel>? Level { get; set; }

        /// <summary>
        /// Type of dataset compression.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public DatasetGZipCompressionArgs()
        {
        }
    }
}
