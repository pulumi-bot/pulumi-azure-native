// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20180401.Inputs
{

    /// <summary>
    /// Network rule collection resource
    /// </summary>
    public sealed class AzureFirewallNetworkRuleCollectionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The action type of a rule collection
        /// </summary>
        [Input("action")]
        public Input<Inputs.AzureFirewallRCActionArgs>? Action { get; set; }

        /// <summary>
        /// Resource ID.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Gets name of the resource that is unique within a resource group. This name can be used to access the resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Priority of the network rule collection resource.
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        /// <summary>
        /// The provisioning state of the resource.
        /// </summary>
        [Input("provisioningState")]
        public Input<string>? ProvisioningState { get; set; }

        [Input("rules")]
        private InputList<Inputs.AzureFirewallNetworkRuleArgs>? _rules;

        /// <summary>
        /// Collection of rules used by a network rule collection.
        /// </summary>
        public InputList<Inputs.AzureFirewallNetworkRuleArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.AzureFirewallNetworkRuleArgs>());
            set => _rules = value;
        }

        public AzureFirewallNetworkRuleCollectionArgs()
        {
        }
    }
}
