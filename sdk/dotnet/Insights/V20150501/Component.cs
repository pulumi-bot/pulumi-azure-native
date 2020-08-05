// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Insights.V20150501
{
    /// <summary>
    /// An Application Insights component definition.
    /// </summary>
    public partial class Component : Pulumi.CustomResource
    {
        /// <summary>
        /// The kind of application that this component refers to, used to customize UI. This value is a freeform string, values should typically be one of the following: web, ios, other, store, java, phone.
        /// </summary>
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// Resource location
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Azure resource name
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Properties that define an Application Insights component resource.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.ApplicationInsightsComponentPropertiesResponseResult> Properties { get; private set; } = null!;

        /// <summary>
        /// Resource tags
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Azure resource type
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Component resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Component(string name, ComponentArgs args, CustomResourceOptions? options = null)
            : base("azurerm:insights/v20150501:Component", name, args ?? new ComponentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Component(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:insights/v20150501:Component", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Component resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Component Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Component(name, id, options);
        }
    }

    public sealed class ComponentArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Type of application being monitored.
        /// </summary>
        [Input("Application_Type", required: true)]
        public Input<string> Application_Type { get; set; } = null!;

        /// <summary>
        /// Disable IP masking.
        /// </summary>
        [Input("DisableIpMasking")]
        public Input<bool>? DisableIpMasking { get; set; }

        /// <summary>
        /// Used by the Application Insights system to determine what kind of flow this component was created by. This is to be set to 'Bluefield' when creating/updating a component via the REST API.
        /// </summary>
        [Input("Flow_Type")]
        public Input<string>? Flow_Type { get; set; }

        /// <summary>
        /// The unique application ID created when a new application is added to HockeyApp, used for communications with HockeyApp.
        /// </summary>
        [Input("HockeyAppId")]
        public Input<string>? HockeyAppId { get; set; }

        /// <summary>
        /// Purge data immediately after 30 days.
        /// </summary>
        [Input("ImmediatePurgeDataOn30Days")]
        public Input<bool>? ImmediatePurgeDataOn30Days { get; set; }

        /// <summary>
        /// Indicates the flow of the ingestion.
        /// </summary>
        [Input("IngestionMode")]
        public Input<string>? IngestionMode { get; set; }

        /// <summary>
        /// Describes what tool created this Application Insights component. Customers using this API should set this to the default 'rest'.
        /// </summary>
        [Input("Request_Source")]
        public Input<string>? Request_Source { get; set; }

        /// <summary>
        /// Retention period in days.
        /// </summary>
        [Input("RetentionInDays")]
        public Input<int>? RetentionInDays { get; set; }

        /// <summary>
        /// Percentage of the data produced by the application being monitored that is being sampled for Application Insights telemetry.
        /// </summary>
        [Input("SamplingPercentage")]
        public Input<double>? SamplingPercentage { get; set; }

        /// <summary>
        /// The kind of application that this component refers to, used to customize UI. This value is a freeform string, values should typically be one of the following: web, ios, other, store, java, phone.
        /// </summary>
        [Input("kind", required: true)]
        public Input<string> Kind { get; set; } = null!;

        /// <summary>
        /// Resource location
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// The name of the Application Insights component resource.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Resource tags
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public ComponentArgs()
        {
        }
    }
}
