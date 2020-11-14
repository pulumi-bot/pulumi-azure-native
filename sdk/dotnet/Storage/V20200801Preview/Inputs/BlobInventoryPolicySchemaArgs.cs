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
    /// The storage account blob inventory policy rules.
    /// </summary>
    public sealed class BlobInventoryPolicySchemaArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Container name where blob inventory files are stored. Must be pre-created.
        /// </summary>
        [Input("destination", required: true)]
        public Input<string> Destination { get; set; } = null!;

        /// <summary>
        /// Policy is enabled if set to true.
        /// </summary>
        [Input("enabled", required: true)]
        public Input<bool> Enabled { get; set; } = null!;

        [Input("rules", required: true)]
        private InputList<Inputs.BlobInventoryPolicyRuleArgs>? _rules;

        /// <summary>
        /// The storage account blob inventory policy rules. The rule is applied when it is enabled.
        /// </summary>
        public InputList<Inputs.BlobInventoryPolicyRuleArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.BlobInventoryPolicyRuleArgs>());
            set => _rules = value;
        }

        /// <summary>
        /// The valid value is Inventory
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public BlobInventoryPolicySchemaArgs()
        {
        }
    }
}
