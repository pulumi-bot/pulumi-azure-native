// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Insights.V20191017Preview
{
    /// <summary>
    /// An Application Insights workbook template definition.
    /// </summary>
    public partial class WorkbookTemplate : Pulumi.CustomResource
    {
        /// <summary>
        /// Information about the author of the workbook template.
        /// </summary>
        [Output("author")]
        public Output<string?> Author { get; private set; } = null!;

        /// <summary>
        /// Workbook galleries supported by the template.
        /// </summary>
        [Output("galleries")]
        public Output<ImmutableArray<Outputs.WorkbookTemplateGalleryResponseResult>> Galleries { get; private set; } = null!;

        /// <summary>
        /// Key value pair of localized gallery. Each key is the locale code of languages supported by the Azure portal.
        /// </summary>
        [Output("localized")]
        public Output<ImmutableDictionary<string, ImmutableArray<Outputs.WorkbookTemplateLocalizedGalleryResponseResult>>?> Localized { get; private set; } = null!;

        /// <summary>
        /// Resource location
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Azure resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Priority of the template. Determines which template to open when a workbook gallery is opened in viewer mode.
        /// </summary>
        [Output("priority")]
        public Output<int?> Priority { get; private set; } = null!;

        /// <summary>
        /// Resource tags
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Valid JSON object containing workbook template payload.
        /// </summary>
        [Output("templateData")]
        public Output<ImmutableDictionary<string, object>> TemplateData { get; private set; } = null!;

        /// <summary>
        /// Azure resource type
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a WorkbookTemplate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public WorkbookTemplate(string name, WorkbookTemplateArgs args, CustomResourceOptions? options = null)
            : base("azurerm:insights/v20191017preview:WorkbookTemplate", name, args ?? new WorkbookTemplateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private WorkbookTemplate(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:insights/v20191017preview:WorkbookTemplate", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing WorkbookTemplate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static WorkbookTemplate Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new WorkbookTemplate(name, id, options);
        }
    }

    public sealed class WorkbookTemplateArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Information about the author of the workbook template.
        /// </summary>
        [Input("author")]
        public Input<string>? Author { get; set; }

        [Input("galleries", required: true)]
        private InputList<Inputs.WorkbookTemplateGalleryArgs>? _galleries;

        /// <summary>
        /// Workbook galleries supported by the template.
        /// </summary>
        public InputList<Inputs.WorkbookTemplateGalleryArgs> Galleries
        {
            get => _galleries ?? (_galleries = new InputList<Inputs.WorkbookTemplateGalleryArgs>());
            set => _galleries = value;
        }

        [Input("localized")]
        private InputMap<ImmutableArray<Inputs.WorkbookTemplateLocalizedGalleryArgs>>? _localized;

        /// <summary>
        /// Key value pair of localized gallery. Each key is the locale code of languages supported by the Azure portal.
        /// </summary>
        public InputMap<ImmutableArray<Inputs.WorkbookTemplateLocalizedGalleryArgs>> Localized
        {
            get => _localized ?? (_localized = new InputMap<ImmutableArray<Inputs.WorkbookTemplateLocalizedGalleryArgs>>());
            set => _localized = value;
        }

        /// <summary>
        /// Resource location
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// Priority of the template. Determines which template to open when a workbook gallery is opened in viewer mode.
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        /// <summary>
        /// The name of the resource group. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the Application Insights component resource.
        /// </summary>
        [Input("resourceName", required: true)]
        public Input<string> ResourceName { get; set; } = null!;

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

        [Input("templateData", required: true)]
        private InputMap<object>? _templateData;

        /// <summary>
        /// Valid JSON object containing workbook template payload.
        /// </summary>
        public InputMap<object> TemplateData
        {
            get => _templateData ?? (_templateData = new InputMap<object>());
            set => _templateData = value;
        }

        public WorkbookTemplateArgs()
        {
        }
    }
}
