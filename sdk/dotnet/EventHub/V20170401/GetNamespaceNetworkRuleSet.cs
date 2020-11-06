// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.EventHub.V20170401
{
    public static class GetNamespaceNetworkRuleSet
    {
        public static Task<GetNamespaceNetworkRuleSetResult> InvokeAsync(GetNamespaceNetworkRuleSetArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetNamespaceNetworkRuleSetResult>("azure-nextgen:eventhub/v20170401:getNamespaceNetworkRuleSet", args ?? new GetNamespaceNetworkRuleSetArgs(), options.WithVersion());
    }


    public sealed class GetNamespaceNetworkRuleSetArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Namespace name
        /// </summary>
        [Input("namespaceName", required: true)]
        public string NamespaceName { get; set; } = null!;

        /// <summary>
        /// Name of the resource group within the azure subscription.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetNamespaceNetworkRuleSetArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetNamespaceNetworkRuleSetResult
    {
        /// <summary>
        /// Default Action for Network Rule Set
        /// </summary>
        public readonly string? DefaultAction;
        /// <summary>
        /// List of IpRules
        /// </summary>
        public readonly ImmutableArray<Outputs.NWRuleSetIpRulesResponse> IpRules;
        /// <summary>
        /// Resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// List VirtualNetwork Rules
        /// </summary>
        public readonly ImmutableArray<Outputs.NWRuleSetVirtualNetworkRulesResponse> VirtualNetworkRules;

        [OutputConstructor]
        private GetNamespaceNetworkRuleSetResult(
            string? defaultAction,

            ImmutableArray<Outputs.NWRuleSetIpRulesResponse> ipRules,

            string name,

            string type,

            ImmutableArray<Outputs.NWRuleSetVirtualNetworkRulesResponse> virtualNetworkRules)
        {
            DefaultAction = defaultAction;
            IpRules = ipRules;
            Name = name;
            Type = type;
            VirtualNetworkRules = virtualNetworkRules;
        }
    }
}
