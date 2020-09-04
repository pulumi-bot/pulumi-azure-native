// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ApiManagement.V20180601Preview
{
    /// <summary>
    /// Diagnostic details.
    /// </summary>
    public partial class Diagnostic : Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies for what type of messages sampling settings should not apply.
        /// </summary>
        [Output("alwaysLog")]
        public Output<string?> AlwaysLog { get; private set; } = null!;

        /// <summary>
        /// Diagnostic settings for incoming/outgoing HTTP messages to the Backend
        /// </summary>
        [Output("backend")]
        public Output<Outputs.PipelineDiagnosticSettingsResponseResult?> Backend { get; private set; } = null!;

        /// <summary>
        /// Whether to process Correlation Headers coming to Api Management Service. Only applicable to Application Insights diagnostics. Default is true.
        /// </summary>
        [Output("enableHttpCorrelationHeaders")]
        public Output<bool?> EnableHttpCorrelationHeaders { get; private set; } = null!;

        /// <summary>
        /// Diagnostic settings for incoming/outgoing HTTP messages to the Gateway.
        /// </summary>
        [Output("frontend")]
        public Output<Outputs.PipelineDiagnosticSettingsResponseResult?> Frontend { get; private set; } = null!;

        /// <summary>
        /// Resource Id of a target logger.
        /// </summary>
        [Output("loggerId")]
        public Output<string> LoggerId { get; private set; } = null!;

        /// <summary>
        /// Resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Sampling settings for Diagnostic.
        /// </summary>
        [Output("sampling")]
        public Output<Outputs.SamplingSettingsResponseResult?> Sampling { get; private set; } = null!;

        /// <summary>
        /// Resource type for API Management resource.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Diagnostic resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Diagnostic(string name, DiagnosticArgs args, CustomResourceOptions? options = null)
            : base("azurerm:apimanagement/v20180601preview:Diagnostic", name, args ?? new DiagnosticArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Diagnostic(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:apimanagement/v20180601preview:Diagnostic", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:apimanagement/latest:Diagnostic"},
                    new Pulumi.Alias { Type = "azurerm:apimanagement/v20170301:Diagnostic"},
                    new Pulumi.Alias { Type = "azurerm:apimanagement/v20180101:Diagnostic"},
                    new Pulumi.Alias { Type = "azurerm:apimanagement/v20190101:Diagnostic"},
                    new Pulumi.Alias { Type = "azurerm:apimanagement/v20191201:Diagnostic"},
                    new Pulumi.Alias { Type = "azurerm:apimanagement/v20191201preview:Diagnostic"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Diagnostic resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Diagnostic Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Diagnostic(name, id, options);
        }
    }

    public sealed class DiagnosticArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies for what type of messages sampling settings should not apply.
        /// </summary>
        [Input("alwaysLog")]
        public Input<string>? AlwaysLog { get; set; }

        /// <summary>
        /// Diagnostic settings for incoming/outgoing HTTP messages to the Backend
        /// </summary>
        [Input("backend")]
        public Input<Inputs.PipelineDiagnosticSettingsArgs>? Backend { get; set; }

        /// <summary>
        /// Diagnostic identifier. Must be unique in the current API Management service instance.
        /// </summary>
        [Input("diagnosticId", required: true)]
        public Input<string> DiagnosticId { get; set; } = null!;

        /// <summary>
        /// Whether to process Correlation Headers coming to Api Management Service. Only applicable to Application Insights diagnostics. Default is true.
        /// </summary>
        [Input("enableHttpCorrelationHeaders")]
        public Input<bool>? EnableHttpCorrelationHeaders { get; set; }

        /// <summary>
        /// Diagnostic settings for incoming/outgoing HTTP messages to the Gateway.
        /// </summary>
        [Input("frontend")]
        public Input<Inputs.PipelineDiagnosticSettingsArgs>? Frontend { get; set; }

        /// <summary>
        /// Resource Id of a target logger.
        /// </summary>
        [Input("loggerId", required: true)]
        public Input<string> LoggerId { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Sampling settings for Diagnostic.
        /// </summary>
        [Input("sampling")]
        public Input<Inputs.SamplingSettingsArgs>? Sampling { get; set; }

        /// <summary>
        /// The name of the API Management service.
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        public DiagnosticArgs()
        {
        }
    }
}
