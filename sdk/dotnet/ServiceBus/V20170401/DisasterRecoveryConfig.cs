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
    /// Single item in List or Get Alias(Disaster Recovery configuration) operation
    /// </summary>
    public partial class DisasterRecoveryConfig : Pulumi.CustomResource
    {
        /// <summary>
        /// Resource name
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Properties required to the Create Or Update Alias(Disaster Recovery configurations)
        /// </summary>
        [Output("properties")]
        public Output<Outputs.ArmDisasterRecoveryResponsePropertiesResult> Properties { get; private set; } = null!;

        /// <summary>
        /// Resource type
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a DisasterRecoveryConfig resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DisasterRecoveryConfig(string name, DisasterRecoveryConfigArgs args, CustomResourceOptions? options = null)
            : base("azurerm:servicebus/v20170401:DisasterRecoveryConfig", name, args ?? new DisasterRecoveryConfigArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DisasterRecoveryConfig(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:servicebus/v20170401:DisasterRecoveryConfig", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing DisasterRecoveryConfig resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DisasterRecoveryConfig Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new DisasterRecoveryConfig(name, id, options);
        }
    }

    public sealed class DisasterRecoveryConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Primary/Secondary eventhub namespace name, which is part of GEO DR pairing
        /// </summary>
        [Input("alternateName")]
        public Input<string>? AlternateName { get; set; }

        /// <summary>
        /// The Disaster Recovery configuration name
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The namespace name
        /// </summary>
        [Input("namespaceName", required: true)]
        public Input<string> NamespaceName { get; set; } = null!;

        /// <summary>
        /// ARM Id of the Primary/Secondary eventhub namespace name, which is part of GEO DR pairing
        /// </summary>
        [Input("partnerNamespace")]
        public Input<string>? PartnerNamespace { get; set; }

        /// <summary>
        /// Name of the Resource group within the Azure subscription.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public DisasterRecoveryConfigArgs()
        {
        }
    }
}
