// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Network
{
    public static class GetUserRule
    {
        /// <summary>
        /// Network security admin rule.
        /// API Version: 2021-02-01-preview.
        /// </summary>
        public static Task<GetUserRuleResult> InvokeAsync(GetUserRuleArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetUserRuleResult>("azure-native:network:getUserRule", args ?? new GetUserRuleArgs(), options.WithVersion());
    }


    public sealed class GetUserRuleArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the network manager security Configuration.
        /// </summary>
        [Input("configurationName", required: true)]
        public string ConfigurationName { get; set; } = null!;

        /// <summary>
        /// The name of the network manager.
        /// </summary>
        [Input("networkManagerName", required: true)]
        public string NetworkManagerName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the rule.
        /// </summary>
        [Input("ruleName", required: true)]
        public string RuleName { get; set; } = null!;

        public GetUserRuleArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetUserRuleResult
    {
        /// <summary>
        /// A description for this rule. Restricted to 140 chars.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The destination address prefixes. CIDR or destination IP ranges.
        /// </summary>
        public readonly ImmutableArray<Outputs.AddressPrefixItemResponse> Destination;
        /// <summary>
        /// The destination port ranges.
        /// </summary>
        public readonly ImmutableArray<string> DestinationPortRanges;
        /// <summary>
        /// Indicates if the traffic matched against the rule in inbound or outbound.
        /// </summary>
        public readonly string Direction;
        /// <summary>
        /// A friendly name for the rule.
        /// </summary>
        public readonly string? DisplayName;
        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// Resource ID.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Network protocol this rule applies to.
        /// </summary>
        public readonly string Protocol;
        /// <summary>
        /// The provisioning state of the security Configuration resource.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// The CIDR or source IP ranges.
        /// </summary>
        public readonly ImmutableArray<Outputs.AddressPrefixItemResponse> Source;
        /// <summary>
        /// The source port ranges.
        /// </summary>
        public readonly ImmutableArray<string> SourcePortRanges;
        /// <summary>
        /// The system metadata related to this resource.
        /// </summary>
        public readonly Outputs.SystemDataResponse SystemData;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetUserRuleResult(
            string? description,

            ImmutableArray<Outputs.AddressPrefixItemResponse> destination,

            ImmutableArray<string> destinationPortRanges,

            string direction,

            string? displayName,

            string etag,

            string id,

            string name,

            string protocol,

            string provisioningState,

            ImmutableArray<Outputs.AddressPrefixItemResponse> source,

            ImmutableArray<string> sourcePortRanges,

            Outputs.SystemDataResponse systemData,

            string type)
        {
            Description = description;
            Destination = destination;
            DestinationPortRanges = destinationPortRanges;
            Direction = direction;
            DisplayName = displayName;
            Etag = etag;
            Id = id;
            Name = name;
            Protocol = protocol;
            ProvisioningState = provisioningState;
            Source = source;
            SourcePortRanges = sourcePortRanges;
            SystemData = systemData;
            Type = type;
        }
    }
}
