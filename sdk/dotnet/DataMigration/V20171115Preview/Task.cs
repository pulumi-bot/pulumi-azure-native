// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DataMigration.V20171115Preview
{
    /// <summary>
    /// A task resource
    /// </summary>
    public partial class Task : Pulumi.CustomResource
    {
        /// <summary>
        /// HTTP strong entity tag value. This is ignored if submitted.
        /// </summary>
        [Output("etag")]
        public Output<string?> Etag { get; private set; } = null!;

        /// <summary>
        /// Resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Custom task properties
        /// </summary>
        [Output("properties")]
        public Output<Outputs.ProjectTaskPropertiesResponseResult> Properties { get; private set; } = null!;

        /// <summary>
        /// Resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Task resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Task(string name, TaskArgs args, CustomResourceOptions? options = null)
            : base("azurerm:datamigration/v20171115preview:Task", name, args ?? new TaskArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Task(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:datamigration/v20171115preview:Task", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:datamigration/latest:Task"},
                    new Pulumi.Alias { Type = "azurerm:datamigration/v20180315preview:Task"},
                    new Pulumi.Alias { Type = "azurerm:datamigration/v20180331preview:Task"},
                    new Pulumi.Alias { Type = "azurerm:datamigration/v20180419:Task"},
                    new Pulumi.Alias { Type = "azurerm:datamigration/v20180715preview:Task"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Task resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Task Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Task(name, id, options);
        }
    }

    public sealed class TaskArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// HTTP strong entity tag value. This is ignored if submitted.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// Name of the resource group
        /// </summary>
        [Input("groupName", required: true)]
        public Input<string> GroupName { get; set; } = null!;

        /// <summary>
        /// Name of the project
        /// </summary>
        [Input("projectName", required: true)]
        public Input<string> ProjectName { get; set; } = null!;

        /// <summary>
        /// Custom task properties
        /// </summary>
        [Input("properties")]
        public Input<Inputs.ProjectTaskPropertiesArgs>? Properties { get; set; }

        /// <summary>
        /// Name of the service
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        /// <summary>
        /// Name of the Task
        /// </summary>
        [Input("taskName", required: true)]
        public Input<string> TaskName { get; set; } = null!;

        public TaskArgs()
        {
        }
    }
}
