// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Management.Latest
{
    public static class GetManagementGroupSubscription
    {
        public static Task<GetManagementGroupSubscriptionResult> InvokeAsync(GetManagementGroupSubscriptionArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetManagementGroupSubscriptionResult>("azure-nextgen:management/latest:getManagementGroupSubscription", args ?? new GetManagementGroupSubscriptionArgs(), options.WithVersion());
    }


    public sealed class GetManagementGroupSubscriptionArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Management Group ID.
        /// </summary>
        [Input("groupId", required: true)]
        public string GroupId { get; set; } = null!;

        /// <summary>
        /// Subscription ID.
        /// </summary>
        [Input("subscriptionId")]
        public string? SubscriptionId { get; set; }

        public GetManagementGroupSubscriptionArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetManagementGroupSubscriptionResult
    {
        /// <summary>
        /// The friendly name of the subscription.
        /// </summary>
        public readonly string? DisplayName;
        /// <summary>
        /// The fully qualified ID for the subscription.  For example, /providers/Microsoft.Management/managementGroups/0000000-0000-0000-0000-000000000000/subscriptions/0000000-0000-0000-0000-000000000001
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The stringified id of the subscription. For example, 00000000-0000-0000-0000-000000000000
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The ID of the parent management group.
        /// </summary>
        public readonly Outputs.DescendantParentGroupInfoResponse? Parent;
        /// <summary>
        /// The state of the subscription.
        /// </summary>
        public readonly string? State;
        /// <summary>
        /// The AAD Tenant ID associated with the subscription. For example, 00000000-0000-0000-0000-000000000000
        /// </summary>
        public readonly string? Tenant;
        /// <summary>
        /// The type of the resource.  For example, Microsoft.Management/managementGroups/subscriptions
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetManagementGroupSubscriptionResult(
            string? displayName,

            string id,

            string name,

            Outputs.DescendantParentGroupInfoResponse? parent,

            string? state,

            string? tenant,

            string type)
        {
            DisplayName = displayName;
            Id = id;
            Name = name;
            Parent = parent;
            State = state;
            Tenant = tenant;
            Type = type;
        }
    }
}
