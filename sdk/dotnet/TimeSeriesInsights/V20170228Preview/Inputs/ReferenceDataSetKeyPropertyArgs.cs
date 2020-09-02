// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.TimeSeriesInsights.V20170228Preview.Inputs
{

    /// <summary>
    /// A key property for the reference data set. A reference data set can have multiple key properties.
    /// </summary>
    public sealed class ReferenceDataSetKeyPropertyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the key property.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The type of the key property.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public ReferenceDataSetKeyPropertyArgs()
        {
        }
    }
}
