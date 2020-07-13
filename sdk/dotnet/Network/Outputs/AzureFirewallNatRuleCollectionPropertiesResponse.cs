// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.Outputs
{

    [OutputType]
    public sealed class AzureFirewallNatRuleCollectionPropertiesResponse
    {
        /// <summary>
        /// The action type of a NAT rule collection.
        /// </summary>
        public readonly Outputs.AzureFirewallNatRCActionResponse? Action;
        /// <summary>
        /// Priority of the NAT rule collection resource.
        /// </summary>
        public readonly int? Priority;
        /// <summary>
        /// The provisioning state of the NAT rule collection resource.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Collection of rules used by a NAT rule collection.
        /// </summary>
        public readonly ImmutableArray<Outputs.AzureFirewallNatRuleResponse> Rules;

        [OutputConstructor]
        private AzureFirewallNatRuleCollectionPropertiesResponse(
            Outputs.AzureFirewallNatRCActionResponse? action,

            int? priority,

            string provisioningState,

            ImmutableArray<Outputs.AzureFirewallNatRuleResponse> rules)
        {
            Action = action;
            Priority = priority;
            ProvisioningState = provisioningState;
            Rules = rules;
        }
    }
}