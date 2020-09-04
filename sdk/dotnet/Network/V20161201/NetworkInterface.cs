// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20161201
{
    /// <summary>
    /// A network interface in a resource group.
    /// </summary>
    public partial class NetworkInterface : Pulumi.CustomResource
    {
        /// <summary>
        /// The DNS settings in network interface.
        /// </summary>
        [Output("dnsSettings")]
        public Output<Outputs.NetworkInterfaceDnsSettingsResponseResult?> DnsSettings { get; private set; } = null!;

        /// <summary>
        /// If the network interface is accelerated networking enabled.
        /// </summary>
        [Output("enableAcceleratedNetworking")]
        public Output<bool?> EnableAcceleratedNetworking { get; private set; } = null!;

        /// <summary>
        /// Indicates whether IP forwarding is enabled on this network interface.
        /// </summary>
        [Output("enableIPForwarding")]
        public Output<bool?> EnableIPForwarding { get; private set; } = null!;

        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        [Output("etag")]
        public Output<string?> Etag { get; private set; } = null!;

        /// <summary>
        /// A list of IPConfigurations of the network interface.
        /// </summary>
        [Output("ipConfigurations")]
        public Output<ImmutableArray<Outputs.NetworkInterfaceIPConfigurationResponseResult>> IpConfigurations { get; private set; } = null!;

        /// <summary>
        /// Resource location.
        /// </summary>
        [Output("location")]
        public Output<string?> Location { get; private set; } = null!;

        /// <summary>
        /// The MAC address of the network interface.
        /// </summary>
        [Output("macAddress")]
        public Output<string?> MacAddress { get; private set; } = null!;

        /// <summary>
        /// Resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The reference of the NetworkSecurityGroup resource.
        /// </summary>
        [Output("networkSecurityGroup")]
        public Output<Outputs.NetworkSecurityGroupResponseResult?> NetworkSecurityGroup { get; private set; } = null!;

        /// <summary>
        /// Gets whether this is a primary network interface on a virtual machine.
        /// </summary>
        [Output("primary")]
        public Output<bool?> Primary { get; private set; } = null!;

        /// <summary>
        /// The provisioning state of the public IP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
        /// </summary>
        [Output("provisioningState")]
        public Output<string?> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// The resource GUID property of the network interface resource.
        /// </summary>
        [Output("resourceGuid")]
        public Output<string?> ResourceGuid { get; private set; } = null!;

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
        /// The reference of a virtual machine.
        /// </summary>
        [Output("virtualMachine")]
        public Output<Outputs.SubResourceResponseResult?> VirtualMachine { get; private set; } = null!;


        /// <summary>
        /// Create a NetworkInterface resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NetworkInterface(string name, NetworkInterfaceArgs args, CustomResourceOptions? options = null)
            : base("azurerm:network/v20161201:NetworkInterface", name, args ?? new NetworkInterfaceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NetworkInterface(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:network/v20161201:NetworkInterface", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:network/latest:NetworkInterface"},
                    new Pulumi.Alias { Type = "azurerm:network/v20150501preview:NetworkInterface"},
                    new Pulumi.Alias { Type = "azurerm:network/v20150615:NetworkInterface"},
                    new Pulumi.Alias { Type = "azurerm:network/v20160330:NetworkInterface"},
                    new Pulumi.Alias { Type = "azurerm:network/v20160601:NetworkInterface"},
                    new Pulumi.Alias { Type = "azurerm:network/v20160901:NetworkInterface"},
                    new Pulumi.Alias { Type = "azurerm:network/v20170301:NetworkInterface"},
                    new Pulumi.Alias { Type = "azurerm:network/v20170601:NetworkInterface"},
                    new Pulumi.Alias { Type = "azurerm:network/v20170801:NetworkInterface"},
                    new Pulumi.Alias { Type = "azurerm:network/v20170901:NetworkInterface"},
                    new Pulumi.Alias { Type = "azurerm:network/v20171001:NetworkInterface"},
                    new Pulumi.Alias { Type = "azurerm:network/v20171101:NetworkInterface"},
                    new Pulumi.Alias { Type = "azurerm:network/v20180101:NetworkInterface"},
                    new Pulumi.Alias { Type = "azurerm:network/v20180201:NetworkInterface"},
                    new Pulumi.Alias { Type = "azurerm:network/v20180401:NetworkInterface"},
                    new Pulumi.Alias { Type = "azurerm:network/v20180601:NetworkInterface"},
                    new Pulumi.Alias { Type = "azurerm:network/v20180701:NetworkInterface"},
                    new Pulumi.Alias { Type = "azurerm:network/v20180801:NetworkInterface"},
                    new Pulumi.Alias { Type = "azurerm:network/v20181001:NetworkInterface"},
                    new Pulumi.Alias { Type = "azurerm:network/v20181101:NetworkInterface"},
                    new Pulumi.Alias { Type = "azurerm:network/v20181201:NetworkInterface"},
                    new Pulumi.Alias { Type = "azurerm:network/v20190201:NetworkInterface"},
                    new Pulumi.Alias { Type = "azurerm:network/v20190401:NetworkInterface"},
                    new Pulumi.Alias { Type = "azurerm:network/v20190601:NetworkInterface"},
                    new Pulumi.Alias { Type = "azurerm:network/v20190701:NetworkInterface"},
                    new Pulumi.Alias { Type = "azurerm:network/v20190801:NetworkInterface"},
                    new Pulumi.Alias { Type = "azurerm:network/v20190901:NetworkInterface"},
                    new Pulumi.Alias { Type = "azurerm:network/v20191101:NetworkInterface"},
                    new Pulumi.Alias { Type = "azurerm:network/v20191201:NetworkInterface"},
                    new Pulumi.Alias { Type = "azurerm:network/v20200301:NetworkInterface"},
                    new Pulumi.Alias { Type = "azurerm:network/v20200401:NetworkInterface"},
                    new Pulumi.Alias { Type = "azurerm:network/v20200501:NetworkInterface"},
                    new Pulumi.Alias { Type = "azurerm:network/v20200601:NetworkInterface"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing NetworkInterface resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NetworkInterface Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new NetworkInterface(name, id, options);
        }
    }

    public sealed class NetworkInterfaceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The DNS settings in network interface.
        /// </summary>
        [Input("dnsSettings")]
        public Input<Inputs.NetworkInterfaceDnsSettingsArgs>? DnsSettings { get; set; }

        /// <summary>
        /// If the network interface is accelerated networking enabled.
        /// </summary>
        [Input("enableAcceleratedNetworking")]
        public Input<bool>? EnableAcceleratedNetworking { get; set; }

        /// <summary>
        /// Indicates whether IP forwarding is enabled on this network interface.
        /// </summary>
        [Input("enableIPForwarding")]
        public Input<bool>? EnableIPForwarding { get; set; }

        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// Resource ID.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("ipConfigurations")]
        private InputList<Inputs.NetworkInterfaceIPConfigurationArgs>? _ipConfigurations;

        /// <summary>
        /// A list of IPConfigurations of the network interface.
        /// </summary>
        public InputList<Inputs.NetworkInterfaceIPConfigurationArgs> IpConfigurations
        {
            get => _ipConfigurations ?? (_ipConfigurations = new InputList<Inputs.NetworkInterfaceIPConfigurationArgs>());
            set => _ipConfigurations = value;
        }

        /// <summary>
        /// Resource location.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The MAC address of the network interface.
        /// </summary>
        [Input("macAddress")]
        public Input<string>? MacAddress { get; set; }

        /// <summary>
        /// The name of the network interface.
        /// </summary>
        [Input("networkInterfaceName", required: true)]
        public Input<string> NetworkInterfaceName { get; set; } = null!;

        /// <summary>
        /// The reference of the NetworkSecurityGroup resource.
        /// </summary>
        [Input("networkSecurityGroup")]
        public Input<Inputs.NetworkSecurityGroupArgs>? NetworkSecurityGroup { get; set; }

        /// <summary>
        /// Gets whether this is a primary network interface on a virtual machine.
        /// </summary>
        [Input("primary")]
        public Input<bool>? Primary { get; set; }

        /// <summary>
        /// The provisioning state of the public IP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
        /// </summary>
        [Input("provisioningState")]
        public Input<string>? ProvisioningState { get; set; }

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The resource GUID property of the network interface resource.
        /// </summary>
        [Input("resourceGuid")]
        public Input<string>? ResourceGuid { get; set; }

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
        /// The reference of a virtual machine.
        /// </summary>
        [Input("virtualMachine")]
        public Input<Inputs.SubResourceArgs>? VirtualMachine { get; set; }

        public NetworkInterfaceArgs()
        {
        }
    }
}
