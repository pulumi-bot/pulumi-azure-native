// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Automation
{
    /// <summary>
    /// Definition of the dsc node configuration.
    /// </summary>
    public partial class AutomationAccountNodeConfiguration : Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Gets or sets the configuration properties.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.DscNodeConfigurationPropertiesResponseResult> Properties { get; private set; } = null!;

        /// <summary>
        /// The type of the resource.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a AutomationAccountNodeConfiguration resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AutomationAccountNodeConfiguration(string name, AutomationAccountNodeConfigurationArgs args, CustomResourceOptions? options = null)
            : base("azurerm:automation:AutomationAccountNodeConfiguration", name, args ?? new AutomationAccountNodeConfigurationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AutomationAccountNodeConfiguration(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:automation:AutomationAccountNodeConfiguration", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing AutomationAccountNodeConfiguration resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AutomationAccountNodeConfiguration Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new AutomationAccountNodeConfiguration(name, id, options);
        }
    }

    public sealed class AutomationAccountNodeConfigurationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the automation account.
        /// </summary>
        [Input("automationAccountName", required: true)]
        public Input<string> AutomationAccountName { get; set; } = null!;

        /// <summary>
        /// The Dsc node configuration name.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Node configuration properties
        /// </summary>
        [Input("properties")]
        public Input<Inputs.DscNodeConfigurationCreateOrUpdateParametersPropertiesArgs>? Properties { get; set; }

        /// <summary>
        /// Name of an Azure Resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Gets or sets the tags attached to the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public AutomationAccountNodeConfigurationArgs()
        {
        }
    }
}
