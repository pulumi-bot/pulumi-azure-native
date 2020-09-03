// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Automation.V20151031
{
    /// <summary>
    /// Definition of the watcher type.
    /// 
    /// ## Example Usage
    /// ### Create or update watcher
    /// ```csharp
    /// using Pulumi;
    /// using AzureRM = Pulumi.AzureRM;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var watcher = new AzureRM.Automation.V20151031.Watcher("watcher", new AzureRM.Automation.V20151031.WatcherArgs
    ///         {
    ///             AutomationAccountName = "MyTestAutomationAccount",
    ///             Description = "This is a test watcher.",
    ///             ExecutionFrequencyInSeconds = 60,
    ///             ResourceGroupName = "rg",
    ///             ScriptName = "MyTestWatcherRunbook",
    ///             ScriptRunOn = "MyTestHybridWorkerGroup",
    ///             Tags = ,
    ///             WatcherName = "MyTestWatcher",
    ///         });
    ///     }
    /// 
    /// }
    /// 
    /// ```
    /// </summary>
    public partial class Watcher : Pulumi.CustomResource
    {
        /// <summary>
        /// Gets or sets the creation time.
        /// </summary>
        [Output("creationTime")]
        public Output<string> CreationTime { get; private set; } = null!;

        /// <summary>
        /// Gets or sets the description.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Gets or sets the etag of the resource.
        /// </summary>
        [Output("etag")]
        public Output<string?> Etag { get; private set; } = null!;

        /// <summary>
        /// Gets or sets the frequency at which the watcher is invoked.
        /// </summary>
        [Output("executionFrequencyInSeconds")]
        public Output<int?> ExecutionFrequencyInSeconds { get; private set; } = null!;

        /// <summary>
        /// Details of the user who last modified the watcher.
        /// </summary>
        [Output("lastModifiedBy")]
        public Output<string> LastModifiedBy { get; private set; } = null!;

        /// <summary>
        /// Gets or sets the last modified time.
        /// </summary>
        [Output("lastModifiedTime")]
        public Output<string> LastModifiedTime { get; private set; } = null!;

        /// <summary>
        /// The Azure Region where the resource lives
        /// </summary>
        [Output("location")]
        public Output<string?> Location { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Gets or sets the name of the script the watcher is attached to, i.e. the name of an existing runbook.
        /// </summary>
        [Output("scriptName")]
        public Output<string?> ScriptName { get; private set; } = null!;

        /// <summary>
        /// Gets or sets the parameters of the script.
        /// </summary>
        [Output("scriptParameters")]
        public Output<ImmutableDictionary<string, string>?> ScriptParameters { get; private set; } = null!;

        /// <summary>
        /// Gets or sets the name of the hybrid worker group the watcher will run on.
        /// </summary>
        [Output("scriptRunOn")]
        public Output<string?> ScriptRunOn { get; private set; } = null!;

        /// <summary>
        /// Gets the current status of the watcher.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Resource tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The type of the resource.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Watcher resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Watcher(string name, WatcherArgs args, CustomResourceOptions? options = null)
            : base("azurerm:automation/v20151031:Watcher", name, args ?? new WatcherArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Watcher(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:automation/v20151031:Watcher", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:automation/latest:Watcher"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Watcher resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Watcher Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Watcher(name, id, options);
        }
    }

    public sealed class WatcherArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the automation account.
        /// </summary>
        [Input("automationAccountName", required: true)]
        public Input<string> AutomationAccountName { get; set; } = null!;

        /// <summary>
        /// Gets or sets the description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Gets or sets the etag of the resource.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// Gets or sets the frequency at which the watcher is invoked.
        /// </summary>
        [Input("executionFrequencyInSeconds")]
        public Input<int>? ExecutionFrequencyInSeconds { get; set; }

        /// <summary>
        /// The Azure Region where the resource lives
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Name of an Azure Resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Gets or sets the name of the script the watcher is attached to, i.e. the name of an existing runbook.
        /// </summary>
        [Input("scriptName")]
        public Input<string>? ScriptName { get; set; }

        [Input("scriptParameters")]
        private InputMap<string>? _scriptParameters;

        /// <summary>
        /// Gets or sets the parameters of the script.
        /// </summary>
        public InputMap<string> ScriptParameters
        {
            get => _scriptParameters ?? (_scriptParameters = new InputMap<string>());
            set => _scriptParameters = value;
        }

        /// <summary>
        /// Gets or sets the name of the hybrid worker group the watcher will run on.
        /// </summary>
        [Input("scriptRunOn")]
        public Input<string>? ScriptRunOn { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Resource tags.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The watcher name.
        /// </summary>
        [Input("watcherName", required: true)]
        public Input<string> WatcherName { get; set; } = null!;

        public WatcherArgs()
        {
        }
    }
}
