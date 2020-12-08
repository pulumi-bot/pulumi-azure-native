// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Network.Latest
{
    /// <summary>
    /// Security Partner Provider resource.
    /// </summary>
    public partial class SecurityPartnerProvider : Pulumi.CustomResource
    {
        /// <summary>
        /// The connection status with the Security Partner Provider.
        /// </summary>
        [Output("connectionStatus")]
        public Output<string> ConnectionStatus { get; private set; } = null!;

        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// Resource location.
        /// </summary>
        [Output("location")]
        public Output<string?> Location { get; private set; } = null!;

        /// <summary>
        /// Resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The provisioning state of the Security Partner Provider resource.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// The security provider name.
        /// </summary>
        [Output("securityProviderName")]
        public Output<string?> SecurityProviderName { get; private set; } = null!;

        /// <summary>
        /// Resource tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// The virtualHub to which the Security Partner Provider belongs.
        /// </summary>
        [Output("virtualHub")]
        public Output<Outputs.SubResourceResponse?> VirtualHub { get; private set; } = null!;


        /// <summary>
        /// Create a SecurityPartnerProvider resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SecurityPartnerProvider(string name, SecurityPartnerProviderArgs args, CustomResourceOptions? options = null)
            : base("azure-nextgen:network/latest:SecurityPartnerProvider", name, args ?? new SecurityPartnerProviderArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SecurityPartnerProvider(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-nextgen:network/latest:SecurityPartnerProvider", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20200301:SecurityPartnerProvider"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20200401:SecurityPartnerProvider"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20200501:SecurityPartnerProvider"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20200601:SecurityPartnerProvider"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20200701:SecurityPartnerProvider"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing SecurityPartnerProvider resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SecurityPartnerProvider Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new SecurityPartnerProvider(name, id, options);
        }
    }

    public sealed class SecurityPartnerProviderArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Resource ID.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Resource location.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the Security Partner Provider.
        /// </summary>
        [Input("securityPartnerProviderName", required: true)]
        public Input<string> SecurityPartnerProviderName { get; set; } = null!;

        /// <summary>
        /// The security provider name.
        /// </summary>
        [Input("securityProviderName")]
        public InputUnion<string, Pulumi.AzureNextGen.Network.Latest.SecurityProviderName>? SecurityProviderName { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Resource tags.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// The virtualHub to which the Security Partner Provider belongs.
        /// </summary>
        [Input("virtualHub")]
        public Input<Inputs.SubResourceArgs>? VirtualHub { get; set; }

        public SecurityPartnerProviderArgs()
        {
        }
    }
}
