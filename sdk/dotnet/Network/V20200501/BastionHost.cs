// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20200501
{
    /// <summary>
    /// Bastion Host resource.
    /// 
    /// ## Example Usage
    /// ### Create Bastion Host
    /// ```csharp
    /// using Pulumi;
    /// using AzureRM = Pulumi.AzureRM;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var bastionHost = new AzureRM.Network.V20200501.BastionHost("bastionHost", new AzureRM.Network.V20200501.BastionHostArgs
    ///         {
    ///             BastionHostName = "bastionhosttenant'",
    ///             IpConfigurations = 
    ///             {
    ///                 new AzureRM.Network.V20200501.Inputs.BastionHostIPConfigurationArgs
    ///                 {
    ///                     Name = "bastionHostIpConfiguration",
    ///                 },
    ///             },
    ///             ResourceGroupName = "rg1",
    ///         });
    ///     }
    /// 
    /// }
    /// 
    /// ```
    /// </summary>
    public partial class BastionHost : Pulumi.CustomResource
    {
        /// <summary>
        /// FQDN for the endpoint on which bastion host is accessible.
        /// </summary>
        [Output("dnsName")]
        public Output<string?> DnsName { get; private set; } = null!;

        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// IP configuration of the Bastion Host resource.
        /// </summary>
        [Output("ipConfigurations")]
        public Output<ImmutableArray<Outputs.BastionHostIPConfigurationResponseResult>> IpConfigurations { get; private set; } = null!;

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
        /// The provisioning state of the bastion host resource.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

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
        /// Create a BastionHost resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public BastionHost(string name, BastionHostArgs args, CustomResourceOptions? options = null)
            : base("azurerm:network/v20200501:BastionHost", name, args ?? new BastionHostArgs(), MakeResourceOptions(options, ""))
        {
        }

        private BastionHost(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:network/v20200501:BastionHost", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:network/latest:BastionHost"},
                    new Pulumi.Alias { Type = "azurerm:network/v20190401:BastionHost"},
                    new Pulumi.Alias { Type = "azurerm:network/v20190601:BastionHost"},
                    new Pulumi.Alias { Type = "azurerm:network/v20190701:BastionHost"},
                    new Pulumi.Alias { Type = "azurerm:network/v20190801:BastionHost"},
                    new Pulumi.Alias { Type = "azurerm:network/v20190901:BastionHost"},
                    new Pulumi.Alias { Type = "azurerm:network/v20191101:BastionHost"},
                    new Pulumi.Alias { Type = "azurerm:network/v20191201:BastionHost"},
                    new Pulumi.Alias { Type = "azurerm:network/v20200301:BastionHost"},
                    new Pulumi.Alias { Type = "azurerm:network/v20200401:BastionHost"},
                    new Pulumi.Alias { Type = "azurerm:network/v20200601:BastionHost"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing BastionHost resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static BastionHost Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new BastionHost(name, id, options);
        }
    }

    public sealed class BastionHostArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the Bastion Host.
        /// </summary>
        [Input("bastionHostName", required: true)]
        public Input<string> BastionHostName { get; set; } = null!;

        /// <summary>
        /// FQDN for the endpoint on which bastion host is accessible.
        /// </summary>
        [Input("dnsName")]
        public Input<string>? DnsName { get; set; }

        /// <summary>
        /// Resource ID.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("ipConfigurations")]
        private InputList<Inputs.BastionHostIPConfigurationArgs>? _ipConfigurations;

        /// <summary>
        /// IP configuration of the Bastion Host resource.
        /// </summary>
        public InputList<Inputs.BastionHostIPConfigurationArgs> IpConfigurations
        {
            get => _ipConfigurations ?? (_ipConfigurations = new InputList<Inputs.BastionHostIPConfigurationArgs>());
            set => _ipConfigurations = value;
        }

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

        public BastionHostArgs()
        {
        }
    }
}
