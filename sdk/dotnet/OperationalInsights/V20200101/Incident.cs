// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.OperationalInsights.V20200101
{
    /// <summary>
    /// Represents an incident in Azure Security Insights.
    /// </summary>
    public partial class Incident : Pulumi.CustomResource
    {
        /// <summary>
        /// Etag of the azure resource
        /// </summary>
        [Output("etag")]
        public Output<string?> Etag { get; private set; } = null!;

        /// <summary>
        /// Azure resource name
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Incident properties
        /// </summary>
        [Output("properties")]
        public Output<Outputs.IncidentPropertiesResponseResult> Properties { get; private set; } = null!;

        /// <summary>
        /// Azure resource type
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Incident resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Incident(string name, IncidentArgs args, CustomResourceOptions? options = null)
            : base("azurerm:operationalinsights/v20200101:Incident", name, args ?? new IncidentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Incident(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:operationalinsights/v20200101:Incident", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Incident resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Incident Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Incident(name, id, options);
        }
    }

    public sealed class IncidentArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The reason the incident was closed
        /// </summary>
        [Input("classification")]
        public Input<string>? Classification { get; set; }

        /// <summary>
        /// Describes the reason the incident was closed
        /// </summary>
        [Input("classificationComment")]
        public Input<string>? ClassificationComment { get; set; }

        /// <summary>
        /// The classification reason the incident was closed with
        /// </summary>
        [Input("classificationReason")]
        public Input<string>? ClassificationReason { get; set; }

        /// <summary>
        /// The description of the incident
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Etag of the azure resource
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// The time of the first activity in the incident
        /// </summary>
        [Input("firstActivityTimeUtc")]
        public Input<string>? FirstActivityTimeUtc { get; set; }

        [Input("labels")]
        private InputList<Inputs.IncidentLabelArgs>? _labels;

        /// <summary>
        /// List of labels relevant to this incident
        /// </summary>
        public InputList<Inputs.IncidentLabelArgs> Labels
        {
            get => _labels ?? (_labels = new InputList<Inputs.IncidentLabelArgs>());
            set => _labels = value;
        }

        /// <summary>
        /// The time of the last activity in the incident
        /// </summary>
        [Input("lastActivityTimeUtc")]
        public Input<string>? LastActivityTimeUtc { get; set; }

        /// <summary>
        /// Incident ID
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Describes a user that the incident is assigned to
        /// </summary>
        [Input("owner")]
        public Input<Inputs.IncidentOwnerInfoArgs>? Owner { get; set; }

        /// <summary>
        /// The name of the resource group within the user's subscription. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The severity of the incident
        /// </summary>
        [Input("severity", required: true)]
        public Input<string> Severity { get; set; } = null!;

        /// <summary>
        /// The status of the incident
        /// </summary>
        [Input("status", required: true)]
        public Input<string> Status { get; set; } = null!;

        /// <summary>
        /// The title of the incident
        /// </summary>
        [Input("title", required: true)]
        public Input<string> Title { get; set; } = null!;

        /// <summary>
        /// The name of the workspace.
        /// </summary>
        [Input("workspaceName", required: true)]
        public Input<string> WorkspaceName { get; set; } = null!;

        public IncidentArgs()
        {
        }
    }
}
