// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ApiManagement.V20191201Preview
{
    /// <summary>
    /// Email Template details.
    /// </summary>
    public partial class EmailTemplate : Pulumi.CustomResource
    {
        /// <summary>
        /// Email Template Body. This should be a valid XDocument
        /// </summary>
        [Output("body")]
        public Output<string> Body { get; private set; } = null!;

        /// <summary>
        /// Description of the Email Template.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Whether the template is the default template provided by Api Management or has been edited.
        /// </summary>
        [Output("isDefault")]
        public Output<bool> IsDefault { get; private set; } = null!;

        /// <summary>
        /// Resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Email Template Parameter values.
        /// </summary>
        [Output("parameters")]
        public Output<ImmutableArray<Outputs.EmailTemplateParametersContractPropertiesResponseResult>> Parameters { get; private set; } = null!;

        /// <summary>
        /// Subject of the Template.
        /// </summary>
        [Output("subject")]
        public Output<string> Subject { get; private set; } = null!;

        /// <summary>
        /// Title of the Template.
        /// </summary>
        [Output("title")]
        public Output<string?> Title { get; private set; } = null!;

        /// <summary>
        /// Resource type for API Management resource.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a EmailTemplate resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EmailTemplate(string name, EmailTemplateArgs args, CustomResourceOptions? options = null)
            : base("azurerm:apimanagement/v20191201preview:EmailTemplate", name, args ?? new EmailTemplateArgs(), MakeResourceOptions(options, ""))
        {
        }

        private EmailTemplate(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:apimanagement/v20191201preview:EmailTemplate", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:apimanagement/latest:EmailTemplate"},
                    new Pulumi.Alias { Type = "azurerm:apimanagement/v20170301:EmailTemplate"},
                    new Pulumi.Alias { Type = "azurerm:apimanagement/v20180101:EmailTemplate"},
                    new Pulumi.Alias { Type = "azurerm:apimanagement/v20180601preview:EmailTemplate"},
                    new Pulumi.Alias { Type = "azurerm:apimanagement/v20190101:EmailTemplate"},
                    new Pulumi.Alias { Type = "azurerm:apimanagement/v20191201:EmailTemplate"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing EmailTemplate resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EmailTemplate Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new EmailTemplate(name, id, options);
        }
    }

    public sealed class EmailTemplateArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Email Template Body. This should be a valid XDocument
        /// </summary>
        [Input("body")]
        public Input<string>? Body { get; set; }

        /// <summary>
        /// Description of the Email Template.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("parameters")]
        private InputList<Inputs.EmailTemplateParametersContractPropertiesArgs>? _parameters;

        /// <summary>
        /// Email Template Parameter values.
        /// </summary>
        public InputList<Inputs.EmailTemplateParametersContractPropertiesArgs> Parameters
        {
            get => _parameters ?? (_parameters = new InputList<Inputs.EmailTemplateParametersContractPropertiesArgs>());
            set => _parameters = value;
        }

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the API Management service.
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        /// <summary>
        /// Subject of the Template.
        /// </summary>
        [Input("subject")]
        public Input<string>? Subject { get; set; }

        /// <summary>
        /// Email Template Name Identifier.
        /// </summary>
        [Input("templateName", required: true)]
        public Input<string> TemplateName { get; set; } = null!;

        /// <summary>
        /// Title of the Template.
        /// </summary>
        [Input("title")]
        public Input<string>? Title { get; set; }

        public EmailTemplateArgs()
        {
        }
    }
}
