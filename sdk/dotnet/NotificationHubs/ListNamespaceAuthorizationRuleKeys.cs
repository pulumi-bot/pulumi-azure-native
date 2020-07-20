// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.NotificationHubs
{
    public static class ListNamespaceAuthorizationRuleKeys
    {
        public static Task<ListNamespaceAuthorizationRuleKeysResult> InvokeAsync(ListNamespaceAuthorizationRuleKeysArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<ListNamespaceAuthorizationRuleKeysResult>("azurerm:notificationhubs:listNamespaceAuthorizationRuleKeys", args ?? new ListNamespaceAuthorizationRuleKeysArgs(), options.WithVersion());
    }


    public sealed class ListNamespaceAuthorizationRuleKeysArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The connection string of the namespace for the specified authorizationRule.
        /// </summary>
        [Input("authorizationRuleName", required: true)]
        public string AuthorizationRuleName { get; set; } = null!;

        /// <summary>
        /// The namespace name.
        /// </summary>
        [Input("namespaceName", required: true)]
        public string NamespaceName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public ListNamespaceAuthorizationRuleKeysArgs()
        {
        }
    }


    [OutputType]
    public sealed class ListNamespaceAuthorizationRuleKeysResult
    {
        /// <summary>
        /// Link to the next set of results. Not empty if Value contains incomplete list of AuthorizationRules
        /// </summary>
        public readonly string? NextLink;
        /// <summary>
        /// Result of the List AuthorizationRules operation.
        /// </summary>
        public readonly ImmutableArray<Outputs.SharedAccessAuthorizationRuleResourceResponseResult> Value;

        [OutputConstructor]
        private ListNamespaceAuthorizationRuleKeysResult(
            string? nextLink,

            ImmutableArray<Outputs.SharedAccessAuthorizationRuleResourceResponseResult> value)
        {
            NextLink = nextLink;
            Value = value;
        }
    }
}