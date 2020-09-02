// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ManagedServices.V20190901
{
    /// <summary>
    /// Registration definition.
    /// </summary>
    public partial class RegistrationDefinition : Pulumi.CustomResource
    {
        /// <summary>
        /// Name of the registration definition.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Plan details for the managed services.
        /// </summary>
        [Output("plan")]
        public Output<Outputs.PlanResponseResult?> Plan { get; private set; } = null!;

        /// <summary>
        /// Properties of a registration definition.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.RegistrationDefinitionPropertiesResponseResult> Properties { get; private set; } = null!;

        /// <summary>
        /// Type of the resource.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a RegistrationDefinition resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RegistrationDefinition(string name, RegistrationDefinitionArgs args, CustomResourceOptions? options = null)
            : base("azurerm:managedservices/v20190901:RegistrationDefinition", name, args ?? new RegistrationDefinitionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RegistrationDefinition(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:managedservices/v20190901:RegistrationDefinition", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:managedservices/latest:RegistrationDefinition"},
                    new Pulumi.Alias { Type = "azurerm:managedservices/v20180601preview:RegistrationDefinition"},
                    new Pulumi.Alias { Type = "azurerm:managedservices/v20190401preview:RegistrationDefinition"},
                    new Pulumi.Alias { Type = "azurerm:managedservices/v20190601:RegistrationDefinition"},
                    new Pulumi.Alias { Type = "azurerm:managedservices/v20200201preview:RegistrationDefinition"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing RegistrationDefinition resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RegistrationDefinition Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new RegistrationDefinition(name, id, options);
        }
    }

    public sealed class RegistrationDefinitionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Plan details for the managed services.
        /// </summary>
        [Input("plan")]
        public Input<Inputs.PlanArgs>? Plan { get; set; }

        /// <summary>
        /// Properties of a registration definition.
        /// </summary>
        [Input("properties")]
        public Input<Inputs.RegistrationDefinitionPropertiesArgs>? Properties { get; set; }

        /// <summary>
        /// Guid of the registration definition.
        /// </summary>
        [Input("registrationDefinitionId", required: true)]
        public Input<string> RegistrationDefinitionId { get; set; } = null!;

        /// <summary>
        /// Scope of the resource.
        /// </summary>
        [Input("scope", required: true)]
        public Input<string> Scope { get; set; } = null!;

        public RegistrationDefinitionArgs()
        {
        }
    }
}
