// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Insights.V20170301Preview.Outputs
{

    [OutputType]
    public sealed class ActivityLogAlertActionListResponseResult
    {
        /// <summary>
        /// The list of activity log alerts.
        /// </summary>
        public readonly ImmutableArray<Outputs.ActivityLogAlertActionGroupResponseResult> ActionGroups;

        [OutputConstructor]
        private ActivityLogAlertActionListResponseResult(ImmutableArray<Outputs.ActivityLogAlertActionGroupResponseResult> actionGroups)
        {
            ActionGroups = actionGroups;
        }
    }
}
