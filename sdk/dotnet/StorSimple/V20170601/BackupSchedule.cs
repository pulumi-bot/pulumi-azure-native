// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.StorSimple.V20170601
{
    /// <summary>
    /// The backup schedule.
    /// </summary>
    public partial class BackupSchedule : Pulumi.CustomResource
    {
        /// <summary>
        /// The type of backup which needs to be taken.
        /// </summary>
        [Output("backupType")]
        public Output<string> BackupType { get; private set; } = null!;

        /// <summary>
        /// The Kind of the object. Currently only Series8000 is supported
        /// </summary>
        [Output("kind")]
        public Output<string?> Kind { get; private set; } = null!;

        /// <summary>
        /// The last successful backup run which was triggered for the schedule.
        /// </summary>
        [Output("lastSuccessfulRun")]
        public Output<string> LastSuccessfulRun { get; private set; } = null!;

        /// <summary>
        /// The name of the object.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The number of backups to be retained.
        /// </summary>
        [Output("retentionCount")]
        public Output<int> RetentionCount { get; private set; } = null!;

        /// <summary>
        /// The schedule recurrence.
        /// </summary>
        [Output("scheduleRecurrence")]
        public Output<Outputs.ScheduleRecurrenceResponse> ScheduleRecurrence { get; private set; } = null!;

        /// <summary>
        /// The schedule status.
        /// </summary>
        [Output("scheduleStatus")]
        public Output<string> ScheduleStatus { get; private set; } = null!;

        /// <summary>
        /// The start time of the schedule.
        /// </summary>
        [Output("startTime")]
        public Output<string> StartTime { get; private set; } = null!;

        /// <summary>
        /// The hierarchical type of the object.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a BackupSchedule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BackupSchedule(string name, BackupScheduleArgs args, CustomResourceOptions? options = null)
            : base("azure-nextgen:storsimple/v20170601:BackupSchedule", name, args ?? new BackupScheduleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BackupSchedule(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-nextgen:storsimple/v20170601:BackupSchedule", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azure-nextgen:storsimple/latest:BackupSchedule"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing BackupSchedule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BackupSchedule Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new BackupSchedule(name, id, options);
        }
    }

    public sealed class BackupScheduleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The backup policy name.
        /// </summary>
        [Input("backupPolicyName", required: true)]
        public Input<string> BackupPolicyName { get; set; } = null!;

        /// <summary>
        /// The backup schedule name.
        /// </summary>
        [Input("backupScheduleName", required: true)]
        public Input<string> BackupScheduleName { get; set; } = null!;

        /// <summary>
        /// The type of backup which needs to be taken.
        /// </summary>
        [Input("backupType", required: true)]
        public Input<Pulumi.AzureNextGen.StorSimple.V20170601.BackupType> BackupType { get; set; } = null!;

        /// <summary>
        /// The device name
        /// </summary>
        [Input("deviceName", required: true)]
        public Input<string> DeviceName { get; set; } = null!;

        /// <summary>
        /// The Kind of the object. Currently only Series8000 is supported
        /// </summary>
        [Input("kind")]
        public Input<Pulumi.AzureNextGen.StorSimple.V20170601.Kind>? Kind { get; set; }

        /// <summary>
        /// The manager name
        /// </summary>
        [Input("managerName", required: true)]
        public Input<string> ManagerName { get; set; } = null!;

        /// <summary>
        /// The resource group name
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The number of backups to be retained.
        /// </summary>
        [Input("retentionCount", required: true)]
        public Input<int> RetentionCount { get; set; } = null!;

        /// <summary>
        /// The schedule recurrence.
        /// </summary>
        [Input("scheduleRecurrence", required: true)]
        public Input<Inputs.ScheduleRecurrenceArgs> ScheduleRecurrence { get; set; } = null!;

        /// <summary>
        /// The schedule status.
        /// </summary>
        [Input("scheduleStatus", required: true)]
        public Input<Pulumi.AzureNextGen.StorSimple.V20170601.ScheduleStatus> ScheduleStatus { get; set; } = null!;

        /// <summary>
        /// The start time of the schedule.
        /// </summary>
        [Input("startTime", required: true)]
        public Input<string> StartTime { get; set; } = null!;

        public BackupScheduleArgs()
        {
        }
    }
}
