// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Network.V20191201
{
    /// <summary>
    /// A DDoS custom policy in a resource group.
    /// </summary>
    public partial class DdosCustomPolicy : Pulumi.CustomResource
    {
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
        /// The protocol-specific DDoS policy customization parameters.
        /// </summary>
        [Output("protocolCustomSettings")]
        public Output<ImmutableArray<Outputs.ProtocolCustomSettingsFormatResponse>> ProtocolCustomSettings { get; private set; } = null!;

        /// <summary>
        /// The provisioning state of the DDoS custom policy resource.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// The list of public IPs associated with the DDoS custom policy resource. This list is read-only.
        /// </summary>
        [Output("publicIPAddresses")]
        public Output<ImmutableArray<Outputs.SubResourceResponse>> PublicIPAddresses { get; private set; } = null!;

        /// <summary>
        /// The resource GUID property of the DDoS custom policy resource. It uniquely identifies the resource, even if the user changes its name or migrate the resource across subscriptions or resource groups.
        /// </summary>
        [Output("resourceGuid")]
        public Output<string> ResourceGuid { get; private set; } = null!;

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
        /// Create a DdosCustomPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DdosCustomPolicy(string name, DdosCustomPolicyArgs args, CustomResourceOptions? options = null)
            : base("azure-nextgen:network/v20191201:DdosCustomPolicy", name, args ?? new DdosCustomPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DdosCustomPolicy(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-nextgen:network/v20191201:DdosCustomPolicy", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azure-nextgen:network/latest:DdosCustomPolicy"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20181101:DdosCustomPolicy"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20181201:DdosCustomPolicy"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20190201:DdosCustomPolicy"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20190401:DdosCustomPolicy"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20190601:DdosCustomPolicy"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20190701:DdosCustomPolicy"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20190801:DdosCustomPolicy"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20190901:DdosCustomPolicy"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20191101:DdosCustomPolicy"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20200301:DdosCustomPolicy"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20200401:DdosCustomPolicy"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20200501:DdosCustomPolicy"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20200601:DdosCustomPolicy"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20200701:DdosCustomPolicy"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing DdosCustomPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DdosCustomPolicy Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new DdosCustomPolicy(name, id, options);
        }
    }

    public sealed class DdosCustomPolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the DDoS custom policy.
        /// </summary>
        [Input("ddosCustomPolicyName", required: true)]
        public Input<string> DdosCustomPolicyName { get; set; } = null!;

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

        [Input("protocolCustomSettings")]
        private InputList<Inputs.ProtocolCustomSettingsFormatArgs>? _protocolCustomSettings;

        /// <summary>
        /// The protocol-specific DDoS policy customization parameters.
        /// </summary>
        public InputList<Inputs.ProtocolCustomSettingsFormatArgs> ProtocolCustomSettings
        {
            get => _protocolCustomSettings ?? (_protocolCustomSettings = new InputList<Inputs.ProtocolCustomSettingsFormatArgs>());
            set => _protocolCustomSettings = value;
        }

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

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

        public DdosCustomPolicyArgs()
        {
        }
    }
}
