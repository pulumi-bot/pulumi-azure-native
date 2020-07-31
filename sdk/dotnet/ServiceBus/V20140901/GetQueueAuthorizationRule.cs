// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ServiceBus.V20140901
{
    public static class GetQueueAuthorizationRule
    {
        public static Task<GetQueueAuthorizationRuleResult> InvokeAsync(GetQueueAuthorizationRuleArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetQueueAuthorizationRuleResult>("azurerm:servicebus/v20140901:getQueueAuthorizationRule", args ?? new GetQueueAuthorizationRuleArgs(), options.WithVersion());
    }


    public sealed class GetQueueAuthorizationRuleArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The authorization rule name.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The namespace name
        /// </summary>
        [Input("namespaceName", required: true)]
        public string NamespaceName { get; set; } = null!;

        /// <summary>
        /// The queue name.
        /// </summary>
        [Input("queueName", required: true)]
        public string QueueName { get; set; } = null!;

        /// <summary>
        /// Name of the Resource group within the Azure subscription.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetQueueAuthorizationRuleArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetQueueAuthorizationRuleResult
    {
        /// <summary>
        /// Resource location.
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// Resource name
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// SharedAccessAuthorizationRule properties.
        /// </summary>
        public readonly Outputs.SharedAccessAuthorizationRuleGetPropertiesResponseResult Properties;
        /// <summary>
        /// Resource type
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetQueueAuthorizationRuleResult(
            string? location,

            string name,

            Outputs.SharedAccessAuthorizationRuleGetPropertiesResponseResult properties,

            string type)
        {
            Location = location;
            Name = name;
            Properties = properties;
            Type = type;
        }
    }
}