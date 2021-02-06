// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.TimeSeriesInsights.V20170228Preview
{
    /// <summary>
    /// An environment receives data from one or more event sources. Each event source has associated connection info that allows the Time Series Insights ingress pipeline to connect to and pull data from the event source
    /// </summary>
    [AzureNextGenResourceType("azure-nextgen:timeseriesinsights/v20170228preview:EventSource")]
    public partial class EventSource : Pulumi.CustomResource
    {
        /// <summary>
        /// The kind of the event source.
        /// </summary>
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// Resource location
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Resource name
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Resource tags
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Resource type
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a EventSource resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EventSource(string name, EventSourceArgs args, CustomResourceOptions? options = null)
            : base("azure-nextgen:timeseriesinsights/v20170228preview:EventSource", name, args ?? new EventSourceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private EventSource(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-nextgen:timeseriesinsights/v20170228preview:EventSource", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azure-nextgen:timeseriesinsights/latest:EventSource"},
                    new Pulumi.Alias { Type = "azure-nextgen:timeseriesinsights/v20171115:EventSource"},
                    new Pulumi.Alias { Type = "azure-nextgen:timeseriesinsights/v20180815preview:EventSource"},
                    new Pulumi.Alias { Type = "azure-nextgen:timeseriesinsights/v20200515:EventSource"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing EventSource resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EventSource Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new EventSource(name, id, options);
        }
    }

    public sealed class EventSourceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the Time Series Insights environment associated with the specified resource group.
        /// </summary>
        [Input("environmentName", required: true)]
        public Input<string> EnvironmentName { get; set; } = null!;

        /// <summary>
        /// Name of the event source.
        /// </summary>
        [Input("eventSourceName", required: true)]
        public Input<string> EventSourceName { get; set; } = null!;

        /// <summary>
        /// The kind of the event source.
        /// </summary>
        [Input("kind", required: true)]
        public Input<Pulumi.AzureNextGen.TimeSeriesInsights.V20170228Preview.Kind> Kind { get; set; } = null!;

        /// <summary>
        /// The location of the resource.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Name of an Azure Resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Key-value pairs of additional properties for the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public EventSourceArgs()
        {
        }
    }
}
