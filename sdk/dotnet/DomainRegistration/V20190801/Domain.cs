// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DomainRegistration.V20190801
{
    /// <summary>
    /// Information about a domain.
    /// </summary>
    public partial class Domain : Pulumi.CustomResource
    {
        /// <summary>
        /// Kind of resource.
        /// </summary>
        [Output("kind")]
        public Output<string?> Kind { get; private set; } = null!;

        /// <summary>
        /// Resource Location.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Resource Name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Domain resource specific properties
        /// </summary>
        [Output("properties")]
        public Output<Outputs.DomainResponsePropertiesResult> Properties { get; private set; } = null!;

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
        /// Create a Domain resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Domain(string name, DomainArgs args, CustomResourceOptions? options = null)
            : base("azurerm:domainregistration/v20190801:Domain", name, args ?? new DomainArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Domain(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:domainregistration/v20190801:Domain", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Domain resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Domain Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Domain(name, id, options);
        }
    }

    public sealed class DomainArgs : Pulumi.ResourceArgs
    {
        [Input("authCode")]
        public Input<string>? AuthCode { get; set; }

        /// <summary>
        /// &lt;code&gt;true&lt;/code&gt; if the domain should be automatically renewed; otherwise, &lt;code&gt;false&lt;/code&gt;.
        /// </summary>
        [Input("autoRenew")]
        public Input<bool>? AutoRenew { get; set; }

        /// <summary>
        /// Legal agreement consent.
        /// </summary>
        [Input("consent", required: true)]
        public Input<Inputs.DomainPurchaseConsentArgs> Consent { get; set; } = null!;

        /// <summary>
        /// Administrative contact.
        /// </summary>
        [Input("contactAdmin", required: true)]
        public Input<Inputs.ContactArgs> ContactAdmin { get; set; } = null!;

        /// <summary>
        /// Billing contact.
        /// </summary>
        [Input("contactBilling", required: true)]
        public Input<Inputs.ContactArgs> ContactBilling { get; set; } = null!;

        /// <summary>
        /// Registrant contact.
        /// </summary>
        [Input("contactRegistrant", required: true)]
        public Input<Inputs.ContactArgs> ContactRegistrant { get; set; } = null!;

        /// <summary>
        /// Technical contact.
        /// </summary>
        [Input("contactTech", required: true)]
        public Input<Inputs.ContactArgs> ContactTech { get; set; } = null!;

        /// <summary>
        /// Current DNS type
        /// </summary>
        [Input("dnsType")]
        public Input<string>? DnsType { get; set; }

        /// <summary>
        /// Azure DNS Zone to use
        /// </summary>
        [Input("dnsZoneId")]
        public Input<string>? DnsZoneId { get; set; }

        /// <summary>
        /// Kind of resource.
        /// </summary>
        [Input("kind")]
        public Input<string>? Kind { get; set; }

        /// <summary>
        /// Resource Location.
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// Name of the domain.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// &lt;code&gt;true&lt;/code&gt; if domain privacy is enabled for this domain; otherwise, &lt;code&gt;false&lt;/code&gt;.
        /// </summary>
        [Input("privacy")]
        public Input<bool>? Privacy { get; set; }

        /// <summary>
        /// Name of the resource group to which the resource belongs.
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

        /// <summary>
        /// Target DNS type (would be used for migration)
        /// </summary>
        [Input("targetDnsType")]
        public Input<string>? TargetDnsType { get; set; }

        public DomainArgs()
        {
        }
    }
}
