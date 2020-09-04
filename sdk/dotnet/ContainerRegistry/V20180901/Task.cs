// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ContainerRegistry.V20180901
{
    /// <summary>
    /// The task that has the ARM resource and task properties.
    /// The task will have all information to schedule a run against it.
    /// 
    /// ## Example Usage
    /// ### Tasks_Create
    /// ```csharp
    /// using Pulumi;
    /// using AzureRM = Pulumi.AzureRM;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var task = new AzureRM.ContainerRegistry.V20180901.Task("task", new AzureRM.ContainerRegistry.V20180901.TaskArgs
    ///         {
    ///             AgentConfiguration = new AzureRM.ContainerRegistry.V20180901.Inputs.AgentPropertiesArgs
    ///             {
    ///                 Cpu = 2,
    ///             },
    ///             Location = "eastus",
    ///             Platform = new AzureRM.ContainerRegistry.V20180901.Inputs.PlatformPropertiesArgs
    ///             {
    ///                 Architecture = "amd64",
    ///                 Os = "Linux",
    ///             },
    ///             RegistryName = "myRegistry",
    ///             ResourceGroupName = "myResourceGroup",
    ///             Status = "Enabled",
    ///             Step = new AzureRM.ContainerRegistry.V20180901.Inputs.TaskStepPropertiesArgs
    ///             {
    ///                 ContextPath = "dockerfiles",
    ///             },
    ///             Tags = 
    ///             {
    ///                 { "testkey", "value" },
    ///             },
    ///             TaskName = "mytTask",
    ///             Trigger = new AzureRM.ContainerRegistry.V20180901.Inputs.TriggerPropertiesArgs
    ///             {
    ///                 BaseImageTrigger = new AzureRM.ContainerRegistry.V20180901.Inputs.BaseImageTriggerArgs
    ///                 {
    ///                     BaseImageTriggerType = "Runtime",
    ///                     Name = "myBaseImageTrigger",
    ///                 },
    ///                 SourceTriggers = 
    ///                 {
    ///                     new AzureRM.ContainerRegistry.V20180901.Inputs.SourceTriggerArgs
    ///                     {
    ///                         Name = "mySourceTrigger",
    ///                         SourceRepository = new AzureRM.ContainerRegistry.V20180901.Inputs.SourcePropertiesArgs
    ///                         {
    ///                             Branch = "master",
    ///                             RepositoryUrl = "https://github.com/Azure/azure-rest-api-specs",
    ///                             SourceControlAuthProperties = new AzureRM.ContainerRegistry.V20180901.Inputs.AuthInfoArgs
    ///                             {
    ///                                 Token = "xxxxx",
    ///                                 TokenType = "PAT",
    ///                             },
    ///                             SourceControlType = "Github",
    ///                         },
    ///                         SourceTriggerEvents = 
    ///                         {
    ///                             "commit",
    ///                         },
    ///                         Status = "Enabled",
    ///                     },
    ///                 },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// 
    /// ```
    /// </summary>
    public partial class Task : Pulumi.CustomResource
    {
        /// <summary>
        /// The machine configuration of the run agent.
        /// </summary>
        [Output("agentConfiguration")]
        public Output<Outputs.AgentPropertiesResponseResult?> AgentConfiguration { get; private set; } = null!;

        /// <summary>
        /// The creation date of task.
        /// </summary>
        [Output("creationDate")]
        public Output<string> CreationDate { get; private set; } = null!;

        /// <summary>
        /// The properties that describes a set of credentials that will be used when this run is invoked.
        /// </summary>
        [Output("credentials")]
        public Output<Outputs.CredentialsResponseResult?> Credentials { get; private set; } = null!;

        /// <summary>
        /// The location of the resource. This cannot be changed after the resource is created.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The name of the resource.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The platform properties against which the run has to happen.
        /// </summary>
        [Output("platform")]
        public Output<Outputs.PlatformPropertiesResponseResult> Platform { get; private set; } = null!;

        /// <summary>
        /// The provisioning state of the task.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// The current status of task.
        /// </summary>
        [Output("status")]
        public Output<string?> Status { get; private set; } = null!;

        /// <summary>
        /// The properties of a task step.
        /// </summary>
        [Output("step")]
        public Output<Outputs.TaskStepPropertiesResponseResult> Step { get; private set; } = null!;

        /// <summary>
        /// The tags of the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Run timeout in seconds.
        /// </summary>
        [Output("timeout")]
        public Output<int?> Timeout { get; private set; } = null!;

        /// <summary>
        /// The properties that describe all triggers for the task.
        /// </summary>
        [Output("trigger")]
        public Output<Outputs.TriggerPropertiesResponseResult?> Trigger { get; private set; } = null!;

        /// <summary>
        /// The type of the resource.
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
            : base("azurerm:containerregistry/v20180901:Task", name, args ?? new TaskArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Task(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:containerregistry/v20180901:Task", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:containerregistry/latest:Task"},
                    new Pulumi.Alias { Type = "azurerm:containerregistry/v20190401:Task"},
                    new Pulumi.Alias { Type = "azurerm:containerregistry/v20190601preview:Task"},
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
        /// The machine configuration of the run agent.
        /// </summary>
        [Input("agentConfiguration")]
        public Input<Inputs.AgentPropertiesArgs>? AgentConfiguration { get; set; }

        /// <summary>
        /// The properties that describes a set of credentials that will be used when this run is invoked.
        /// </summary>
        [Input("credentials")]
        public Input<Inputs.CredentialsArgs>? Credentials { get; set; }

        /// <summary>
        /// The location of the resource. This cannot be changed after the resource is created.
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// The platform properties against which the run has to happen.
        /// </summary>
        [Input("platform", required: true)]
        public Input<Inputs.PlatformPropertiesArgs> Platform { get; set; } = null!;

        /// <summary>
        /// The name of the container registry.
        /// </summary>
        [Input("registryName", required: true)]
        public Input<string> RegistryName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group to which the container registry belongs.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The current status of task.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The properties of a task step.
        /// </summary>
        [Input("step", required: true)]
        public Input<Inputs.TaskStepPropertiesArgs> Step { get; set; } = null!;

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
        /// The name of the container registry task.
        /// </summary>
        [Input("taskName", required: true)]
        public Input<string> TaskName { get; set; } = null!;

        /// <summary>
        /// Run timeout in seconds.
        /// </summary>
        [Input("timeout")]
        public Input<int>? Timeout { get; set; }

        /// <summary>
        /// The properties that describe all triggers for the task.
        /// </summary>
        [Input("trigger")]
        public Input<Inputs.TriggerPropertiesArgs>? Trigger { get; set; }

        public TaskArgs()
        {
        }
    }
}
