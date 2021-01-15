// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Automation.Latest
{
    /// <summary>
    /// Definition of the job schedule.
    /// Latest API Version: 2019-06-01.
    /// </summary>
    [AzureNextGenResourceType("azure-nextgen:automation/latest:JobSchedule")]
    public partial class JobSchedule : Pulumi.CustomResource
    {
        /// <summary>
        /// Gets or sets the id of job schedule.
        /// </summary>
        [Output("jobScheduleId")]
        public Output<string?> JobScheduleId { get; private set; } = null!;

        /// <summary>
        /// Gets the name of the variable.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Gets or sets the parameters of the job schedule.
        /// </summary>
        [Output("parameters")]
        public Output<ImmutableDictionary<string, string>?> Parameters { get; private set; } = null!;

        /// <summary>
        /// Gets or sets the hybrid worker group that the scheduled job should run on.
        /// </summary>
        [Output("runOn")]
        public Output<string?> RunOn { get; private set; } = null!;

        /// <summary>
        /// Gets or sets the runbook.
        /// </summary>
        [Output("runbook")]
        public Output<Outputs.RunbookAssociationPropertyResponse?> Runbook { get; private set; } = null!;

        /// <summary>
        /// Gets or sets the schedule.
        /// </summary>
        [Output("schedule")]
        public Output<Outputs.ScheduleAssociationPropertyResponse?> Schedule { get; private set; } = null!;

        /// <summary>
        /// Resource type
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a JobSchedule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public JobSchedule(string name, JobScheduleArgs args, CustomResourceOptions? options = null)
            : base("azure-nextgen:automation/latest:JobSchedule", name, args ?? new JobScheduleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private JobSchedule(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-nextgen:automation/latest:JobSchedule", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azure-nextgen:automation/v20151031:JobSchedule"},
                    new Pulumi.Alias { Type = "azure-nextgen:automation/v20190601:JobSchedule"},
                    new Pulumi.Alias { Type = "azure-nextgen:automation/v20200113preview:JobSchedule"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing JobSchedule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static JobSchedule Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new JobSchedule(name, id, options);
        }
    }

    public sealed class JobScheduleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the automation account.
        /// </summary>
        [Input("automationAccountName", required: true)]
        public Input<string> AutomationAccountName { get; set; } = null!;

        /// <summary>
        /// The job schedule name.
        /// </summary>
        [Input("jobScheduleId", required: true)]
        public Input<string> JobScheduleId { get; set; } = null!;

        [Input("parameters")]
        private InputMap<string>? _parameters;

        /// <summary>
        /// Gets or sets a list of job properties.
        /// </summary>
        public InputMap<string> Parameters
        {
            get => _parameters ?? (_parameters = new InputMap<string>());
            set => _parameters = value;
        }

        /// <summary>
        /// Name of an Azure Resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Gets or sets the hybrid worker group that the scheduled job should run on.
        /// </summary>
        [Input("runOn")]
        public Input<string>? RunOn { get; set; }

        /// <summary>
        /// Gets or sets the runbook.
        /// </summary>
        [Input("runbook", required: true)]
        public Input<Inputs.RunbookAssociationPropertyArgs> Runbook { get; set; } = null!;

        /// <summary>
        /// Gets or sets the schedule.
        /// </summary>
        [Input("schedule", required: true)]
        public Input<Inputs.ScheduleAssociationPropertyArgs> Schedule { get; set; } = null!;

        public JobScheduleArgs()
        {
        }
    }
}
