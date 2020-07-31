// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Relay.V20160701
{
    public static class GetHybridConnectionAuthorizationRule
    {
        public static Task<GetHybridConnectionAuthorizationRuleResult> InvokeAsync(GetHybridConnectionAuthorizationRuleArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetHybridConnectionAuthorizationRuleResult>("azurerm:relay/v20160701:getHybridConnectionAuthorizationRule", args ?? new GetHybridConnectionAuthorizationRuleArgs(), options.WithVersion());
    }


    public sealed class GetHybridConnectionAuthorizationRuleArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The hybrid connection name.
        /// </summary>
        [Input("hybridConnectionName", required: true)]
        public string HybridConnectionName { get; set; } = null!;

        /// <summary>
        /// The authorizationRule name.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The Namespace Name
        /// </summary>
        [Input("namespaceName", required: true)]
        public string NamespaceName { get; set; } = null!;

        /// <summary>
        /// Name of the Resource group within the Azure subscription.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetHybridConnectionAuthorizationRuleArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetHybridConnectionAuthorizationRuleResult
    {
        /// <summary>
        /// Resource name
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Authorization Rule properties
        /// </summary>
        public readonly Outputs.AuthorizationRulePropertiesResponseResult Properties;
        /// <summary>
        /// Resource type
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetHybridConnectionAuthorizationRuleResult(
            string name,

            Outputs.AuthorizationRulePropertiesResponseResult properties,

            string type)
        {
            Name = name;
            Properties = properties;
            Type = type;
        }
    }
}