// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20190801
{
    /// <summary>
    /// Azure Firewall resource.
    /// </summary>
    public partial class AzureFirewall : Pulumi.CustomResource
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
        /// Properties of the azure firewall.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.AzureFirewallPropertiesFormatResponseResult> Properties { get; private set; } = null!;

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
        /// A list of availability zones denoting where the resource needs to come from.
        /// </summary>
        [Output("zones")]
        public Output<ImmutableArray<string>> Zones { get; private set; } = null!;


        /// <summary>
        /// Create a AzureFirewall resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AzureFirewall(string name, AzureFirewallArgs args, CustomResourceOptions? options = null)
            : base("azurerm:network/v20190801:AzureFirewall", name, args ?? new AzureFirewallArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AzureFirewall(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:network/v20190801:AzureFirewall", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing AzureFirewall resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AzureFirewall Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new AzureFirewall(name, id, options);
        }
    }

    public sealed class AzureFirewallArgs : Pulumi.ResourceArgs
    {
        [Input("applicationRuleCollections")]
        private InputList<Inputs.AzureFirewallApplicationRuleCollectionArgs>? _applicationRuleCollections;

        /// <summary>
        /// Collection of application rule collections used by Azure Firewall.
        /// </summary>
        public InputList<Inputs.AzureFirewallApplicationRuleCollectionArgs> ApplicationRuleCollections
        {
            get => _applicationRuleCollections ?? (_applicationRuleCollections = new InputList<Inputs.AzureFirewallApplicationRuleCollectionArgs>());
            set => _applicationRuleCollections = value;
        }

        /// <summary>
        /// The firewallPolicy associated with this azure firewall.
        /// </summary>
        [Input("firewallPolicy")]
        public Input<Inputs.SubResourceArgs>? FirewallPolicy { get; set; }

        /// <summary>
        /// Resource ID.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("ipConfigurations")]
        private InputList<Inputs.AzureFirewallIPConfigurationArgs>? _ipConfigurations;

        /// <summary>
        /// IP configuration of the Azure Firewall resource.
        /// </summary>
        public InputList<Inputs.AzureFirewallIPConfigurationArgs> IpConfigurations
        {
            get => _ipConfigurations ?? (_ipConfigurations = new InputList<Inputs.AzureFirewallIPConfigurationArgs>());
            set => _ipConfigurations = value;
        }

        /// <summary>
        /// Resource location.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the Azure Firewall.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("natRuleCollections")]
        private InputList<Inputs.AzureFirewallNatRuleCollectionArgs>? _natRuleCollections;

        /// <summary>
        /// Collection of NAT rule collections used by Azure Firewall.
        /// </summary>
        public InputList<Inputs.AzureFirewallNatRuleCollectionArgs> NatRuleCollections
        {
            get => _natRuleCollections ?? (_natRuleCollections = new InputList<Inputs.AzureFirewallNatRuleCollectionArgs>());
            set => _natRuleCollections = value;
        }

        [Input("networkRuleCollections")]
        private InputList<Inputs.AzureFirewallNetworkRuleCollectionArgs>? _networkRuleCollections;

        /// <summary>
        /// Collection of network rule collections used by Azure Firewall.
        /// </summary>
        public InputList<Inputs.AzureFirewallNetworkRuleCollectionArgs> NetworkRuleCollections
        {
            get => _networkRuleCollections ?? (_networkRuleCollections = new InputList<Inputs.AzureFirewallNetworkRuleCollectionArgs>());
            set => _networkRuleCollections = value;
        }

        /// <summary>
        /// The provisioning state of the Azure firewall resource.
        /// </summary>
        [Input("provisioningState")]
        public Input<string>? ProvisioningState { get; set; }

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The Azure Firewall Resource SKU.
        /// </summary>
        [Input("sku")]
        public Input<Inputs.AzureFirewallSkuArgs>? Sku { get; set; }

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
        /// The operation mode for Threat Intelligence.
        /// </summary>
        [Input("threatIntelMode")]
        public Input<string>? ThreatIntelMode { get; set; }

        /// <summary>
        /// The virtualHub to which the firewall belongs.
        /// </summary>
        [Input("virtualHub")]
        public Input<Inputs.SubResourceArgs>? VirtualHub { get; set; }

        [Input("zones")]
        private InputList<string>? _zones;

        /// <summary>
        /// A list of availability zones denoting where the resource needs to come from.
        /// </summary>
        public InputList<string> Zones
        {
            get => _zones ?? (_zones = new InputList<string>());
            set => _zones = value;
        }

        public AzureFirewallArgs()
        {
        }
    }
}
