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
    /// Definition of the connection type.
    /// </summary>
    public partial class AutomationAccountConnectionType : Pulumi.CustomResource
    {
        /// <summary>
        /// Gets the name of the connection type.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Gets or sets the properties of the connection type.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.ConnectionTypePropertiesResponseResult> Properties { get; private set; } = null!;

        /// <summary>
        /// Resource type
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a AutomationAccountConnectionType resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AutomationAccountConnectionType(string name, AutomationAccountConnectionTypeArgs args, CustomResourceOptions? options = null)
            : base("azurerm:automation:AutomationAccountConnectionType", name, args ?? new AutomationAccountConnectionTypeArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AutomationAccountConnectionType(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:automation:AutomationAccountConnectionType", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing AutomationAccountConnectionType resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AutomationAccountConnectionType Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new AutomationAccountConnectionType(name, id, options);
        }
    }

    public sealed class AutomationAccountConnectionTypeArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the automation account.
        /// </summary>
        [Input("automationAccountName", required: true)]
        public Input<string> AutomationAccountName { get; set; } = null!;

        /// <summary>
        /// The parameters supplied to the create or update connection type operation.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Gets or sets the value of the connection type.
        /// </summary>
        [Input("properties", required: true)]
        public Input<Inputs.ConnectionTypeCreateOrUpdatePropertiesArgs> Properties { get; set; } = null!;

        /// <summary>
        /// Name of an Azure Resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public AutomationAccountConnectionTypeArgs()
        {
        }
    }
}
