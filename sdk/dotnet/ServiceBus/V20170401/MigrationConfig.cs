// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ServiceBus.V20170401
{
    /// <summary>
    /// Single item in List or Get Migration Config operation
    /// </summary>
    public partial class MigrationConfig : Pulumi.CustomResource
    {
        /// <summary>
        /// Resource name
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Properties required to the Create Migration Configuration
        /// </summary>
        [Output("properties")]
        public Output<Outputs.MigrationConfigPropertiesResponsePropertiesResult> Properties { get; private set; } = null!;

        /// <summary>
        /// Resource type
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a MigrationConfig resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MigrationConfig(string name, MigrationConfigArgs args, CustomResourceOptions? options = null)
            : base("azurerm:servicebus/v20170401:MigrationConfig", name, args ?? new MigrationConfigArgs(), MakeResourceOptions(options, ""))
        {
        }

        private MigrationConfig(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:servicebus/v20170401:MigrationConfig", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing MigrationConfig resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static MigrationConfig Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new MigrationConfig(name, id, options);
        }
    }

    public sealed class MigrationConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The configuration name. Should always be "$default".
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The namespace name
        /// </summary>
        [Input("namespaceName", required: true)]
        public Input<string> NamespaceName { get; set; } = null!;

        /// <summary>
        /// Name to access Standard Namespace after migration
        /// </summary>
        [Input("postMigrationName", required: true)]
        public Input<string> PostMigrationName { get; set; } = null!;

        /// <summary>
        /// Name of the Resource group within the Azure subscription.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Existing premium Namespace ARM Id name which has no entities, will be used for migration
        /// </summary>
        [Input("targetNamespace", required: true)]
        public Input<string> TargetNamespace { get; set; } = null!;

        public MigrationConfigArgs()
        {
        }
    }
}
