// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.CostManagement.V20191001.Inputs
{

    /// <summary>
    /// The configuration for sorting in the query.
    /// </summary>
    public sealed class QuerySortingConfigurationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the column to use in sorting.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The sorting direction
        /// </summary>
        [Input("querySortingDirection")]
        public InputUnion<string, Pulumi.AzureNextGen.CostManagement.V20191001.SortDirection>? QuerySortingDirection { get; set; }

        public QuerySortingConfigurationArgs()
        {
        }
    }
}
