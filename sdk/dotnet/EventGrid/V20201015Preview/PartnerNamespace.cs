// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.EventGrid.V20201015Preview
{
    /// <summary>
    /// EventGrid Partner Namespace.
    /// </summary>
    [AzureNextGenResourceType("azure-nextgen:eventgrid/v20201015preview:PartnerNamespace")]
    public partial class PartnerNamespace : Pulumi.CustomResource
    {
        /// <summary>
        /// Endpoint for the partner namespace.
        /// </summary>
        [Output("endpoint")]
        public Output<string> Endpoint { get; private set; } = null!;

        /// <summary>
        /// Location of the resource.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Name of the resource.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The fully qualified ARM Id of the partner registration that should be associated with this partner namespace. This takes the following format:
        /// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/partnerRegistrations/{partnerRegistrationName}.
        /// </summary>
        [Output("partnerRegistrationFullyQualifiedId")]
        public Output<string?> PartnerRegistrationFullyQualifiedId { get; private set; } = null!;

        /// <summary>
        /// Provisioning state of the partner namespace.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// The system metadata relating to this resource.
        /// </summary>
        [Output("systemData")]
        public Output<Outputs.SystemDataResponse> SystemData { get; private set; } = null!;

        /// <summary>
        /// Tags of the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Type of the resource.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a PartnerNamespace resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PartnerNamespace(string name, PartnerNamespaceArgs args, CustomResourceOptions? options = null)
            : base("azure-nextgen:eventgrid/v20201015preview:PartnerNamespace", name, args ?? new PartnerNamespaceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PartnerNamespace(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-nextgen:eventgrid/v20201015preview:PartnerNamespace", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azure-nextgen:eventgrid/v20200401preview:PartnerNamespace"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing PartnerNamespace resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PartnerNamespace Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new PartnerNamespace(name, id, options);
        }
    }

    public sealed class PartnerNamespaceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Location of the resource.
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// Name of the partner namespace.
        /// </summary>
        [Input("partnerNamespaceName", required: true)]
        public Input<string> PartnerNamespaceName { get; set; } = null!;

        /// <summary>
        /// The fully qualified ARM Id of the partner registration that should be associated with this partner namespace. This takes the following format:
        /// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/partnerRegistrations/{partnerRegistrationName}.
        /// </summary>
        [Input("partnerRegistrationFullyQualifiedId")]
        public Input<string>? PartnerRegistrationFullyQualifiedId { get; set; }

        /// <summary>
        /// The name of the resource group within the user's subscription.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Tags of the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public PartnerNamespaceArgs()
        {
        }
    }
}
