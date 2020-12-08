// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Network.V20200501.Inputs
{

    /// <summary>
    /// Allow to exclude some variable satisfy the condition for the WAF check.
    /// </summary>
    public sealed class OwaspCrsExclusionEntryArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The variable to be excluded.
        /// </summary>
        [Input("matchVariable", required: true)]
        public InputUnion<string, Pulumi.AzureNextGen.Network.V20200501.OwaspCrsExclusionEntryMatchVariable> MatchVariable { get; set; } = null!;

        /// <summary>
        /// When matchVariable is a collection, operator used to specify which elements in the collection this exclusion applies to.
        /// </summary>
        [Input("selector", required: true)]
        public Input<string> Selector { get; set; } = null!;

        /// <summary>
        /// When matchVariable is a collection, operate on the selector to specify which elements in the collection this exclusion applies to.
        /// </summary>
        [Input("selectorMatchOperator", required: true)]
        public InputUnion<string, Pulumi.AzureNextGen.Network.V20200501.OwaspCrsExclusionEntrySelectorMatchOperator> SelectorMatchOperator { get; set; } = null!;

        public OwaspCrsExclusionEntryArgs()
        {
        }
    }
}
