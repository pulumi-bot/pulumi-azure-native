// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DevTestLab.V20160515
{
    /// <summary>
    /// A schedule.
    /// </summary>
    public partial class GlobalSchedule : Pulumi.CustomResource
    {
        /// <summary>
        /// The location of the resource.
        /// </summary>
        [Output("location")]
        public Output<string?> Location { get; private set; } = null!;

        /// <summary>
        /// The name of the resource.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The properties of the resource.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.SchedulePropertiesResponseResult> Properties { get; private set; } = null!;

        /// <summary>
        /// The tags of the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The type of the resource.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a GlobalSchedule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GlobalSchedule(string name, GlobalScheduleArgs args, CustomResourceOptions? options = null)
            : base("azurerm:devtestlab/v20160515:GlobalSchedule", name, args ?? new GlobalScheduleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GlobalSchedule(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:devtestlab/v20160515:GlobalSchedule", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing GlobalSchedule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GlobalSchedule Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new GlobalSchedule(name, id, options);
        }
    }

    public sealed class GlobalScheduleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// If the schedule will occur once each day of the week, specify the daily recurrence.
        /// </summary>
        [Input("dailyRecurrence")]
        public Input<Inputs.DayDetailsArgs>? DailyRecurrence { get; set; }

        /// <summary>
        /// If the schedule will occur multiple times a day, specify the hourly recurrence.
        /// </summary>
        [Input("hourlyRecurrence")]
        public Input<Inputs.HourDetailsArgs>? HourlyRecurrence { get; set; }

        /// <summary>
        /// The location of the resource.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the schedule.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Notification settings.
        /// </summary>
        [Input("notificationSettings")]
        public Input<Inputs.NotificationSettingsArgs>? NotificationSettings { get; set; }

        /// <summary>
        /// The provisioning status of the resource.
        /// </summary>
        [Input("provisioningState")]
        public Input<string>? ProvisioningState { get; set; }

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The status of the schedule (i.e. Enabled, Disabled)
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// The tags of the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The resource ID to which the schedule belongs
        /// </summary>
        [Input("targetResourceId")]
        public Input<string>? TargetResourceId { get; set; }

        /// <summary>
        /// The task type of the schedule (e.g. LabVmsShutdownTask, LabVmAutoStart).
        /// </summary>
        [Input("taskType")]
        public Input<string>? TaskType { get; set; }

        /// <summary>
        /// The time zone ID (e.g. Pacific Standard time).
        /// </summary>
        [Input("timeZoneId")]
        public Input<string>? TimeZoneId { get; set; }

        /// <summary>
        /// The unique immutable identifier of a resource (Guid).
        /// </summary>
        [Input("uniqueIdentifier")]
        public Input<string>? UniqueIdentifier { get; set; }

        /// <summary>
        /// If the schedule will occur only some days of the week, specify the weekly recurrence.
        /// </summary>
        [Input("weeklyRecurrence")]
        public Input<Inputs.WeekDetailsArgs>? WeeklyRecurrence { get; set; }

        public GlobalScheduleArgs()
        {
        }
    }
}
