// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Storage.Inputs
{

    /// <summary>
    /// The Storage Account ObjectReplicationPolicy properties.
    /// </summary>
    public sealed class ObjectReplicationPolicyPropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Required. Destination account name.
        /// </summary>
        [Input("destinationAccount", required: true)]
        public Input<string> DestinationAccount { get; set; } = null!;

        [Input("rules")]
        private InputList<Inputs.ObjectReplicationPolicyRuleArgs>? _rules;

        /// <summary>
        /// The storage account object replication rules.
        /// </summary>
        public InputList<Inputs.ObjectReplicationPolicyRuleArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.ObjectReplicationPolicyRuleArgs>());
            set => _rules = value;
        }

        /// <summary>
        /// Required. Source account name.
        /// </summary>
        [Input("sourceAccount", required: true)]
        public Input<string> SourceAccount { get; set; } = null!;

        public ObjectReplicationPolicyPropertiesArgs()
        {
        }
    }
}