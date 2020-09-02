// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Cdn.V20171012
{
    /// <summary>
    /// Friendly domain name mapping to the endpoint hostname that the customer provides for branding purposes, e.g. www.contoso.com.
    /// </summary>
    public partial class CustomDomain : Pulumi.CustomResource
    {
        /// <summary>
        /// Provisioning status of Custom Https of the custom domain.
        /// </summary>
        [Output("customHttpsProvisioningState")]
        public Output<string> CustomHttpsProvisioningState { get; private set; } = null!;

        /// <summary>
        /// Provisioning substate shows the progress of custom HTTPS enabling/disabling process step by step.
        /// </summary>
        [Output("customHttpsProvisioningSubstate")]
        public Output<string> CustomHttpsProvisioningSubstate { get; private set; } = null!;

        /// <summary>
        /// The host name of the custom domain. Must be a domain name.
        /// </summary>
        [Output("hostName")]
        public Output<string> HostName { get; private set; } = null!;

        /// <summary>
        /// Resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Provisioning status of the custom domain.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// Resource status of the custom domain.
        /// </summary>
        [Output("resourceState")]
        public Output<string> ResourceState { get; private set; } = null!;

        /// <summary>
        /// Resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// Special validation or data may be required when delivering CDN to some regions due to local compliance reasons. E.g. ICP license number of a custom domain is required to deliver content in China.
        /// </summary>
        [Output("validationData")]
        public Output<string?> ValidationData { get; private set; } = null!;


        /// <summary>
        /// Create a CustomDomain resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CustomDomain(string name, CustomDomainArgs args, CustomResourceOptions? options = null)
            : base("azurerm:cdn/v20171012:CustomDomain", name, args ?? new CustomDomainArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CustomDomain(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:cdn/v20171012:CustomDomain", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:cdn/latest:CustomDomain"},
                    new Pulumi.Alias { Type = "azurerm:cdn/v20150601:CustomDomain"},
                    new Pulumi.Alias { Type = "azurerm:cdn/v20160402:CustomDomain"},
                    new Pulumi.Alias { Type = "azurerm:cdn/v20161002:CustomDomain"},
                    new Pulumi.Alias { Type = "azurerm:cdn/v20170402:CustomDomain"},
                    new Pulumi.Alias { Type = "azurerm:cdn/v20190415:CustomDomain"},
                    new Pulumi.Alias { Type = "azurerm:cdn/v20190615:CustomDomain"},
                    new Pulumi.Alias { Type = "azurerm:cdn/v20190615preview:CustomDomain"},
                    new Pulumi.Alias { Type = "azurerm:cdn/v20191231:CustomDomain"},
                    new Pulumi.Alias { Type = "azurerm:cdn/v20200331:CustomDomain"},
                    new Pulumi.Alias { Type = "azurerm:cdn/v20200415:CustomDomain"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing CustomDomain resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CustomDomain Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new CustomDomain(name, id, options);
        }
    }

    public sealed class CustomDomainArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the custom domain within an endpoint.
        /// </summary>
        [Input("customDomainName", required: true)]
        public Input<string> CustomDomainName { get; set; } = null!;

        /// <summary>
        /// Name of the endpoint under the profile which is unique globally.
        /// </summary>
        [Input("endpointName", required: true)]
        public Input<string> EndpointName { get; set; } = null!;

        /// <summary>
        /// The host name of the custom domain. Must be a domain name.
        /// </summary>
        [Input("hostName", required: true)]
        public Input<string> HostName { get; set; } = null!;

        /// <summary>
        /// Name of the CDN profile which is unique within the resource group.
        /// </summary>
        [Input("profileName", required: true)]
        public Input<string> ProfileName { get; set; } = null!;

        /// <summary>
        /// Name of the Resource group within the Azure subscription.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        public CustomDomainArgs()
        {
        }
    }
}
