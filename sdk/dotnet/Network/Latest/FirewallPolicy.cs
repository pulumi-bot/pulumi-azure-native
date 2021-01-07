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
    /// FirewallPolicy Resource.
    /// Latest API Version: 2020-08-01.
    /// </summary>
    [AzureNextGenResourceType("azure-nextgen:network/latest:FirewallPolicy")]
    public partial class FirewallPolicy : Pulumi.CustomResource
    {
        /// <summary>
        /// The parent firewall policy from which rules are inherited.
        /// </summary>
        [Output("basePolicy")]
        public Output<Outputs.SubResourceResponse?> BasePolicy { get; private set; } = null!;

        /// <summary>
        /// List of references to Child Firewall Policies.
        /// </summary>
        [Output("childPolicies")]
        public Output<ImmutableArray<Outputs.SubResourceResponse>> ChildPolicies { get; private set; } = null!;

        /// <summary>
        /// DNS Proxy Settings definition.
        /// </summary>
        [Output("dnsSettings")]
        public Output<Outputs.DnsSettingsResponse?> DnsSettings { get; private set; } = null!;

        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// List of references to Azure Firewalls that this Firewall Policy is associated with.
        /// </summary>
        [Output("firewalls")]
        public Output<ImmutableArray<Outputs.SubResourceResponse>> Firewalls { get; private set; } = null!;

        /// <summary>
        /// The identity of the firewall policy.
        /// </summary>
        [Output("identity")]
        public Output<Outputs.ManagedServiceIdentityResponse?> Identity { get; private set; } = null!;

        /// <summary>
        /// The configuration for Intrusion detection.
        /// </summary>
        [Output("intrusionDetection")]
        public Output<Outputs.FirewallPolicyIntrusionDetectionResponse?> IntrusionDetection { get; private set; } = null!;

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
        /// The provisioning state of the firewall policy resource.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// List of references to FirewallPolicyRuleCollectionGroups.
        /// </summary>
        [Output("ruleCollectionGroups")]
        public Output<ImmutableArray<Outputs.SubResourceResponse>> RuleCollectionGroups { get; private set; } = null!;

        /// <summary>
        /// The Firewall Policy SKU.
        /// </summary>
        [Output("sku")]
        public Output<Outputs.FirewallPolicySkuResponse?> Sku { get; private set; } = null!;

        /// <summary>
        /// Resource tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The operation mode for Threat Intelligence.
        /// </summary>
        [Output("threatIntelMode")]
        public Output<string?> ThreatIntelMode { get; private set; } = null!;

        /// <summary>
        /// ThreatIntel Whitelist for Firewall Policy.
        /// </summary>
        [Output("threatIntelWhitelist")]
        public Output<Outputs.FirewallPolicyThreatIntelWhitelistResponse?> ThreatIntelWhitelist { get; private set; } = null!;

        /// <summary>
        /// TLS Configuration definition.
        /// </summary>
        [Output("transportSecurity")]
        public Output<Outputs.FirewallPolicyTransportSecurityResponse?> TransportSecurity { get; private set; } = null!;

        /// <summary>
        /// Resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a FirewallPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FirewallPolicy(string name, FirewallPolicyArgs args, CustomResourceOptions? options = null)
            : base("azure-nextgen:network/latest:FirewallPolicy", name, args ?? new FirewallPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FirewallPolicy(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-nextgen:network/latest:FirewallPolicy", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20190601:FirewallPolicy"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20190701:FirewallPolicy"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20190801:FirewallPolicy"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20190901:FirewallPolicy"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20191101:FirewallPolicy"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20191201:FirewallPolicy"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20200301:FirewallPolicy"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20200401:FirewallPolicy"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20200501:FirewallPolicy"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20200601:FirewallPolicy"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20200701:FirewallPolicy"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20200801:FirewallPolicy"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing FirewallPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FirewallPolicy Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new FirewallPolicy(name, id, options);
        }
    }

    public sealed class FirewallPolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The parent firewall policy from which rules are inherited.
        /// </summary>
        [Input("basePolicy")]
        public Input<Inputs.SubResourceArgs>? BasePolicy { get; set; }

        /// <summary>
        /// DNS Proxy Settings definition.
        /// </summary>
        [Input("dnsSettings")]
        public Input<Inputs.DnsSettingsArgs>? DnsSettings { get; set; }

        /// <summary>
        /// The name of the Firewall Policy.
        /// </summary>
        [Input("firewallPolicyName", required: true)]
        public Input<string> FirewallPolicyName { get; set; } = null!;

        /// <summary>
        /// Resource ID.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The identity of the firewall policy.
        /// </summary>
        [Input("identity")]
        public Input<Inputs.ManagedServiceIdentityArgs>? Identity { get; set; }

        /// <summary>
        /// The configuration for Intrusion detection.
        /// </summary>
        [Input("intrusionDetection")]
        public Input<Inputs.FirewallPolicyIntrusionDetectionArgs>? IntrusionDetection { get; set; }

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
        /// The Firewall Policy SKU.
        /// </summary>
        [Input("sku")]
        public Input<Inputs.FirewallPolicySkuArgs>? Sku { get; set; }

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
        public InputUnion<string, Pulumi.AzureNextGen.Network.Latest.AzureFirewallThreatIntelMode>? ThreatIntelMode { get; set; }

        /// <summary>
        /// ThreatIntel Whitelist for Firewall Policy.
        /// </summary>
        [Input("threatIntelWhitelist")]
        public Input<Inputs.FirewallPolicyThreatIntelWhitelistArgs>? ThreatIntelWhitelist { get; set; }

        /// <summary>
        /// TLS Configuration definition.
        /// </summary>
        [Input("transportSecurity")]
        public Input<Inputs.FirewallPolicyTransportSecurityArgs>? TransportSecurity { get; set; }

        public FirewallPolicyArgs()
        {
        }
    }
}
