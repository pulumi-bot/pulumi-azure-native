// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20160330
{
    /// <summary>
    /// Subnet in a VirtualNetwork resource
    /// </summary>
    public partial class Subnet : Pulumi.CustomResource
    {
        /// <summary>
        /// Gets or sets Address prefix for the subnet.
        /// </summary>
        [Output("addressPrefix")]
        public Output<string?> AddressPrefix { get; private set; } = null!;

        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated
        /// </summary>
        [Output("etag")]
        public Output<string?> Etag { get; private set; } = null!;

        /// <summary>
        /// Gets array of references to the network interface IP configurations using subnet
        /// </summary>
        [Output("ipConfigurations")]
        public Output<ImmutableArray<Outputs.IPConfigurationResponseResult>> IpConfigurations { get; private set; } = null!;

        /// <summary>
        /// Gets name of the resource that is unique within a resource group. This name can be used to access the resource
        /// </summary>
        [Output("name")]
        public Output<string?> Name { get; private set; } = null!;

        /// <summary>
        /// Gets or sets the reference of the NetworkSecurityGroup resource
        /// </summary>
        [Output("networkSecurityGroup")]
        public Output<Outputs.NetworkSecurityGroupResponseResult?> NetworkSecurityGroup { get; private set; } = null!;

        /// <summary>
        /// Gets or sets Provisioning state of the PublicIP resource Updating/Deleting/Failed
        /// </summary>
        [Output("provisioningState")]
        public Output<string?> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// Gets or sets the reference of the RouteTable resource
        /// </summary>
        [Output("routeTable")]
        public Output<Outputs.RouteTableResponseResult?> RouteTable { get; private set; } = null!;


        /// <summary>
        /// Create a Subnet resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Subnet(string name, SubnetArgs args, CustomResourceOptions? options = null)
            : base("azurerm:network/v20160330:Subnet", name, args ?? new SubnetArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Subnet(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:network/v20160330:Subnet", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:network/latest:Subnet"},
                    new Pulumi.Alias { Type = "azurerm:network/v20150501preview:Subnet"},
                    new Pulumi.Alias { Type = "azurerm:network/v20150615:Subnet"},
                    new Pulumi.Alias { Type = "azurerm:network/v20160601:Subnet"},
                    new Pulumi.Alias { Type = "azurerm:network/v20160901:Subnet"},
                    new Pulumi.Alias { Type = "azurerm:network/v20161201:Subnet"},
                    new Pulumi.Alias { Type = "azurerm:network/v20170301:Subnet"},
                    new Pulumi.Alias { Type = "azurerm:network/v20170601:Subnet"},
                    new Pulumi.Alias { Type = "azurerm:network/v20170801:Subnet"},
                    new Pulumi.Alias { Type = "azurerm:network/v20170901:Subnet"},
                    new Pulumi.Alias { Type = "azurerm:network/v20171001:Subnet"},
                    new Pulumi.Alias { Type = "azurerm:network/v20171101:Subnet"},
                    new Pulumi.Alias { Type = "azurerm:network/v20180101:Subnet"},
                    new Pulumi.Alias { Type = "azurerm:network/v20180201:Subnet"},
                    new Pulumi.Alias { Type = "azurerm:network/v20180401:Subnet"},
                    new Pulumi.Alias { Type = "azurerm:network/v20180601:Subnet"},
                    new Pulumi.Alias { Type = "azurerm:network/v20180701:Subnet"},
                    new Pulumi.Alias { Type = "azurerm:network/v20180801:Subnet"},
                    new Pulumi.Alias { Type = "azurerm:network/v20181001:Subnet"},
                    new Pulumi.Alias { Type = "azurerm:network/v20181101:Subnet"},
                    new Pulumi.Alias { Type = "azurerm:network/v20181201:Subnet"},
                    new Pulumi.Alias { Type = "azurerm:network/v20190201:Subnet"},
                    new Pulumi.Alias { Type = "azurerm:network/v20190401:Subnet"},
                    new Pulumi.Alias { Type = "azurerm:network/v20190601:Subnet"},
                    new Pulumi.Alias { Type = "azurerm:network/v20190701:Subnet"},
                    new Pulumi.Alias { Type = "azurerm:network/v20190801:Subnet"},
                    new Pulumi.Alias { Type = "azurerm:network/v20190901:Subnet"},
                    new Pulumi.Alias { Type = "azurerm:network/v20191101:Subnet"},
                    new Pulumi.Alias { Type = "azurerm:network/v20191201:Subnet"},
                    new Pulumi.Alias { Type = "azurerm:network/v20200301:Subnet"},
                    new Pulumi.Alias { Type = "azurerm:network/v20200401:Subnet"},
                    new Pulumi.Alias { Type = "azurerm:network/v20200501:Subnet"},
                    new Pulumi.Alias { Type = "azurerm:network/v20200601:Subnet"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Subnet resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Subnet Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Subnet(name, id, options);
        }
    }

    public sealed class SubnetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Gets or sets Address prefix for the subnet.
        /// </summary>
        [Input("addressPrefix")]
        public Input<string>? AddressPrefix { get; set; }

        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// Resource Id
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("ipConfigurations")]
        private InputList<Inputs.IPConfigurationArgs>? _ipConfigurations;

        /// <summary>
        /// Gets array of references to the network interface IP configurations using subnet
        /// </summary>
        public InputList<Inputs.IPConfigurationArgs> IpConfigurations
        {
            get => _ipConfigurations ?? (_ipConfigurations = new InputList<Inputs.IPConfigurationArgs>());
            set => _ipConfigurations = value;
        }

        /// <summary>
        /// Gets name of the resource that is unique within a resource group. This name can be used to access the resource
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Gets or sets the reference of the NetworkSecurityGroup resource
        /// </summary>
        [Input("networkSecurityGroup")]
        public Input<Inputs.NetworkSecurityGroupArgs>? NetworkSecurityGroup { get; set; }

        /// <summary>
        /// Gets or sets Provisioning state of the PublicIP resource Updating/Deleting/Failed
        /// </summary>
        [Input("provisioningState")]
        public Input<string>? ProvisioningState { get; set; }

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Gets or sets the reference of the RouteTable resource
        /// </summary>
        [Input("routeTable")]
        public Input<Inputs.RouteTableArgs>? RouteTable { get; set; }

        /// <summary>
        /// The name of the subnet.
        /// </summary>
        [Input("subnetName", required: true)]
        public Input<string> SubnetName { get; set; } = null!;

        /// <summary>
        /// The name of the virtual network.
        /// </summary>
        [Input("virtualNetworkName", required: true)]
        public Input<string> VirtualNetworkName { get; set; } = null!;

        public SubnetArgs()
        {
        }
    }
}
