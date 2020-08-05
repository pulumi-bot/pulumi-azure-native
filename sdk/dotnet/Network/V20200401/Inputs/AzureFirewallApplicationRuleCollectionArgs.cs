// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20200401.Inputs
{

    /// <summary>
    /// Application rule collection resource.
    /// </summary>
    public sealed class AzureFirewallApplicationRuleCollectionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The action type of a rule collection.
        /// </summary>
        [Input("action")]
        public Input<Inputs.AzureFirewallRCActionArgs>? Action { get; set; }

        /// <summary>
        /// Resource ID.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The name of the resource that is unique within the Azure firewall. This name can be used to access the resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Priority of the application rule collection resource.
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        [Input("rules")]
        private InputList<Inputs.AzureFirewallApplicationRuleArgs>? _rules;

        /// <summary>
        /// Collection of rules used by a application rule collection.
        /// </summary>
        public InputList<Inputs.AzureFirewallApplicationRuleArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.AzureFirewallApplicationRuleArgs>());
            set => _rules = value;
        }

        public AzureFirewallApplicationRuleCollectionArgs()
        {
        }
    }
}
