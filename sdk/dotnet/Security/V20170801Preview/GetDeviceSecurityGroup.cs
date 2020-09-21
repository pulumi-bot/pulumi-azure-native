// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Security.V20170801Preview
{
    public static class GetDeviceSecurityGroup
    {
        public static Task<GetDeviceSecurityGroupResult> InvokeAsync(GetDeviceSecurityGroupArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDeviceSecurityGroupResult>("azure-nextgen:security/v20170801preview:getDeviceSecurityGroup", args ?? new GetDeviceSecurityGroupArgs(), options.WithVersion());
    }


    public sealed class GetDeviceSecurityGroupArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the device security group. Note that the name of the device security group is case insensitive.
        /// </summary>
        [Input("deviceSecurityGroupName", required: true)]
        public string DeviceSecurityGroupName { get; set; } = null!;

        /// <summary>
        /// The identifier of the resource.
        /// </summary>
        [Input("resourceId", required: true)]
        public string ResourceId { get; set; } = null!;

        public GetDeviceSecurityGroupArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDeviceSecurityGroupResult
    {
        /// <summary>
        /// The allow-list custom alert rules.
        /// </summary>
        public readonly ImmutableArray<Outputs.AllowlistCustomAlertRuleResponse> AllowlistRules;
        /// <summary>
        /// The deny-list custom alert rules.
        /// </summary>
        public readonly ImmutableArray<Outputs.DenylistCustomAlertRuleResponse> DenylistRules;
        /// <summary>
        /// Resource name
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The list of custom alert threshold rules.
        /// </summary>
        public readonly ImmutableArray<Outputs.ThresholdCustomAlertRuleResponse> ThresholdRules;
        /// <summary>
        /// The list of custom alert time-window rules.
        /// </summary>
        public readonly ImmutableArray<Outputs.TimeWindowCustomAlertRuleResponse> TimeWindowRules;
        /// <summary>
        /// Resource type
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetDeviceSecurityGroupResult(
            ImmutableArray<Outputs.AllowlistCustomAlertRuleResponse> allowlistRules,

            ImmutableArray<Outputs.DenylistCustomAlertRuleResponse> denylistRules,

            string name,

            ImmutableArray<Outputs.ThresholdCustomAlertRuleResponse> thresholdRules,

            ImmutableArray<Outputs.TimeWindowCustomAlertRuleResponse> timeWindowRules,

            string type)
        {
            AllowlistRules = allowlistRules;
            DenylistRules = denylistRules;
            Name = name;
            ThresholdRules = thresholdRules;
            TimeWindowRules = timeWindowRules;
            Type = type;
        }
    }
}
