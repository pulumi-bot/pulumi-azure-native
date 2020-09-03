// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20180801
{
    /// <summary>
    /// P2SVpnServerConfiguration Resource.
    /// 
    /// ## Example Usage
    /// ### P2SVpnServerConfigurationPut
    /// ```csharp
    /// using Pulumi;
    /// using AzureRM = Pulumi.AzureRM;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var p2sVpnServerConfiguration = new AzureRM.Network.V20180801.P2sVpnServerConfiguration("p2sVpnServerConfiguration", new AzureRM.Network.V20180801.P2sVpnServerConfigurationArgs
    ///         {
    ///             P2SVpnServerConfigRadiusClientRootCertificates = 
    ///             {
    ///                 new AzureRM.Network.V20180801.Inputs.P2SVpnServerConfigRadiusClientRootCertificateArgs
    ///                 {
    ///                     Name = "p2sVpnServerConfigRadiusClientRootCert1",
    ///                 },
    ///             },
    ///             P2SVpnServerConfigRadiusServerRootCertificates = 
    ///             {
    ///                 new AzureRM.Network.V20180801.Inputs.P2SVpnServerConfigRadiusServerRootCertificateArgs
    ///                 {
    ///                     Name = "p2sVpnServerConfigRadiusServerRootCert1",
    ///                 },
    ///             },
    ///             P2SVpnServerConfigVpnClientRevokedCertificates = 
    ///             {
    ///                 new AzureRM.Network.V20180801.Inputs.P2SVpnServerConfigVpnClientRevokedCertificateArgs
    ///                 {
    ///                     Name = "p2sVpnServerConfigVpnClientRevokedCert1",
    ///                 },
    ///             },
    ///             P2SVpnServerConfigVpnClientRootCertificates = 
    ///             {
    ///                 new AzureRM.Network.V20180801.Inputs.P2SVpnServerConfigVpnClientRootCertificateArgs
    ///                 {
    ///                     Name = "p2sVpnServerConfigVpnClientRootCert1",
    ///                 },
    ///             },
    ///             P2SVpnServerConfigurationName = "p2sVpnServerConfiguration1",
    ///             RadiusServerAddress = "8.9.9.9",
    ///             RadiusServerSecret = "123_abc",
    ///             ResourceGroupName = "rg1",
    ///             VirtualWanName = "virtualWan1",
    ///             VpnClientIpsecPolicies = 
    ///             {
    ///                 new AzureRM.Network.V20180801.Inputs.IpsecPolicyArgs
    ///                 {
    ///                     DhGroup = "DHGroup14",
    ///                     IkeEncryption = "AES256",
    ///                     IkeIntegrity = "SHA384",
    ///                     IpsecEncryption = "AES256",
    ///                     IpsecIntegrity = "SHA256",
    ///                     PfsGroup = "PFS14",
    ///                     SaDataSizeKilobytes = 429497,
    ///                     SaLifeTimeSeconds = 86472,
    ///                 },
    ///             },
    ///             VpnProtocols = 
    ///             {
    ///                 "IkeV2",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// 
    /// ```
    /// </summary>
    public partial class P2sVpnServerConfiguration : Pulumi.CustomResource
    {
        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// The name of the P2SVpnServerConfiguration that is unique within a VirtualWan in a resource group. This name can be used to access the resource along with Parent VirtualWan resource name.
        /// </summary>
        [Output("name")]
        public Output<string?> Name { get; private set; } = null!;

        [Output("p2SVpnGateways")]
        public Output<ImmutableArray<Outputs.SubResourceResponseResult>> P2SVpnGateways { get; private set; } = null!;

        /// <summary>
        /// Radius client root certificate of P2SVpnServerConfiguration.
        /// </summary>
        [Output("p2SVpnServerConfigRadiusClientRootCertificates")]
        public Output<ImmutableArray<Outputs.P2SVpnServerConfigRadiusClientRootCertificateResponseResult>> P2SVpnServerConfigRadiusClientRootCertificates { get; private set; } = null!;

        /// <summary>
        /// Radius Server root certificate of P2SVpnServerConfiguration.
        /// </summary>
        [Output("p2SVpnServerConfigRadiusServerRootCertificates")]
        public Output<ImmutableArray<Outputs.P2SVpnServerConfigRadiusServerRootCertificateResponseResult>> P2SVpnServerConfigRadiusServerRootCertificates { get; private set; } = null!;

        /// <summary>
        /// VPN client revoked certificate of P2SVpnServerConfiguration.
        /// </summary>
        [Output("p2SVpnServerConfigVpnClientRevokedCertificates")]
        public Output<ImmutableArray<Outputs.P2SVpnServerConfigVpnClientRevokedCertificateResponseResult>> P2SVpnServerConfigVpnClientRevokedCertificates { get; private set; } = null!;

        /// <summary>
        /// VPN client root certificate of P2SVpnServerConfiguration.
        /// </summary>
        [Output("p2SVpnServerConfigVpnClientRootCertificates")]
        public Output<ImmutableArray<Outputs.P2SVpnServerConfigVpnClientRootCertificateResponseResult>> P2SVpnServerConfigVpnClientRootCertificates { get; private set; } = null!;

        /// <summary>
        /// The provisioning state of the P2SVpnServerConfiguration resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// The radius server address property of the P2SVpnServerConfiguration resource for point to site client connection.
        /// </summary>
        [Output("radiusServerAddress")]
        public Output<string?> RadiusServerAddress { get; private set; } = null!;

        /// <summary>
        /// The radius secret property of the P2SVpnServerConfiguration resource for point to site client connection.
        /// </summary>
        [Output("radiusServerSecret")]
        public Output<string?> RadiusServerSecret { get; private set; } = null!;

        /// <summary>
        /// VpnClientIpsecPolicies for P2SVpnServerConfiguration.
        /// </summary>
        [Output("vpnClientIpsecPolicies")]
        public Output<ImmutableArray<Outputs.IpsecPolicyResponseResult>> VpnClientIpsecPolicies { get; private set; } = null!;

        /// <summary>
        /// vpnProtocols for the P2SVpnServerConfiguration.
        /// </summary>
        [Output("vpnProtocols")]
        public Output<ImmutableArray<string>> VpnProtocols { get; private set; } = null!;


        /// <summary>
        /// Create a P2sVpnServerConfiguration resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public P2sVpnServerConfiguration(string name, P2sVpnServerConfigurationArgs args, CustomResourceOptions? options = null)
            : base("azurerm:network/v20180801:P2sVpnServerConfiguration", name, args ?? new P2sVpnServerConfigurationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private P2sVpnServerConfiguration(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:network/v20180801:P2sVpnServerConfiguration", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:network/latest:P2sVpnServerConfiguration"},
                    new Pulumi.Alias { Type = "azurerm:network/v20181001:P2sVpnServerConfiguration"},
                    new Pulumi.Alias { Type = "azurerm:network/v20181101:P2sVpnServerConfiguration"},
                    new Pulumi.Alias { Type = "azurerm:network/v20181201:P2sVpnServerConfiguration"},
                    new Pulumi.Alias { Type = "azurerm:network/v20190201:P2sVpnServerConfiguration"},
                    new Pulumi.Alias { Type = "azurerm:network/v20190401:P2sVpnServerConfiguration"},
                    new Pulumi.Alias { Type = "azurerm:network/v20190601:P2sVpnServerConfiguration"},
                    new Pulumi.Alias { Type = "azurerm:network/v20190701:P2sVpnServerConfiguration"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing P2sVpnServerConfiguration resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static P2sVpnServerConfiguration Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new P2sVpnServerConfiguration(name, id, options);
        }
    }

    public sealed class P2sVpnServerConfigurationArgs : Pulumi.ResourceArgs
    {
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

        /// <summary>
        /// The name of the P2SVpnServerConfiguration that is unique within a VirtualWan in a resource group. This name can be used to access the resource along with Parent VirtualWan resource name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("p2SVpnServerConfigRadiusClientRootCertificates")]
        private InputList<Inputs.P2SVpnServerConfigRadiusClientRootCertificateArgs>? _p2SVpnServerConfigRadiusClientRootCertificates;

        /// <summary>
        /// Radius client root certificate of P2SVpnServerConfiguration.
        /// </summary>
        public InputList<Inputs.P2SVpnServerConfigRadiusClientRootCertificateArgs> P2SVpnServerConfigRadiusClientRootCertificates
        {
            get => _p2SVpnServerConfigRadiusClientRootCertificates ?? (_p2SVpnServerConfigRadiusClientRootCertificates = new InputList<Inputs.P2SVpnServerConfigRadiusClientRootCertificateArgs>());
            set => _p2SVpnServerConfigRadiusClientRootCertificates = value;
        }

        [Input("p2SVpnServerConfigRadiusServerRootCertificates")]
        private InputList<Inputs.P2SVpnServerConfigRadiusServerRootCertificateArgs>? _p2SVpnServerConfigRadiusServerRootCertificates;

        /// <summary>
        /// Radius Server root certificate of P2SVpnServerConfiguration.
        /// </summary>
        public InputList<Inputs.P2SVpnServerConfigRadiusServerRootCertificateArgs> P2SVpnServerConfigRadiusServerRootCertificates
        {
            get => _p2SVpnServerConfigRadiusServerRootCertificates ?? (_p2SVpnServerConfigRadiusServerRootCertificates = new InputList<Inputs.P2SVpnServerConfigRadiusServerRootCertificateArgs>());
            set => _p2SVpnServerConfigRadiusServerRootCertificates = value;
        }

        [Input("p2SVpnServerConfigVpnClientRevokedCertificates")]
        private InputList<Inputs.P2SVpnServerConfigVpnClientRevokedCertificateArgs>? _p2SVpnServerConfigVpnClientRevokedCertificates;

        /// <summary>
        /// VPN client revoked certificate of P2SVpnServerConfiguration.
        /// </summary>
        public InputList<Inputs.P2SVpnServerConfigVpnClientRevokedCertificateArgs> P2SVpnServerConfigVpnClientRevokedCertificates
        {
            get => _p2SVpnServerConfigVpnClientRevokedCertificates ?? (_p2SVpnServerConfigVpnClientRevokedCertificates = new InputList<Inputs.P2SVpnServerConfigVpnClientRevokedCertificateArgs>());
            set => _p2SVpnServerConfigVpnClientRevokedCertificates = value;
        }

        [Input("p2SVpnServerConfigVpnClientRootCertificates")]
        private InputList<Inputs.P2SVpnServerConfigVpnClientRootCertificateArgs>? _p2SVpnServerConfigVpnClientRootCertificates;

        /// <summary>
        /// VPN client root certificate of P2SVpnServerConfiguration.
        /// </summary>
        public InputList<Inputs.P2SVpnServerConfigVpnClientRootCertificateArgs> P2SVpnServerConfigVpnClientRootCertificates
        {
            get => _p2SVpnServerConfigVpnClientRootCertificates ?? (_p2SVpnServerConfigVpnClientRootCertificates = new InputList<Inputs.P2SVpnServerConfigVpnClientRootCertificateArgs>());
            set => _p2SVpnServerConfigVpnClientRootCertificates = value;
        }

        /// <summary>
        /// The name of the P2SVpnServerConfiguration.
        /// </summary>
        [Input("p2SVpnServerConfigurationName", required: true)]
        public Input<string> P2SVpnServerConfigurationName { get; set; } = null!;

        /// <summary>
        /// The radius server address property of the P2SVpnServerConfiguration resource for point to site client connection.
        /// </summary>
        [Input("radiusServerAddress")]
        public Input<string>? RadiusServerAddress { get; set; }

        /// <summary>
        /// The radius secret property of the P2SVpnServerConfiguration resource for point to site client connection.
        /// </summary>
        [Input("radiusServerSecret")]
        public Input<string>? RadiusServerSecret { get; set; }

        /// <summary>
        /// The resource group name of the VirtualWan.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the VirtualWan.
        /// </summary>
        [Input("virtualWanName", required: true)]
        public Input<string> VirtualWanName { get; set; } = null!;

        [Input("vpnClientIpsecPolicies")]
        private InputList<Inputs.IpsecPolicyArgs>? _vpnClientIpsecPolicies;

        /// <summary>
        /// VpnClientIpsecPolicies for P2SVpnServerConfiguration.
        /// </summary>
        public InputList<Inputs.IpsecPolicyArgs> VpnClientIpsecPolicies
        {
            get => _vpnClientIpsecPolicies ?? (_vpnClientIpsecPolicies = new InputList<Inputs.IpsecPolicyArgs>());
            set => _vpnClientIpsecPolicies = value;
        }

        [Input("vpnProtocols")]
        private InputList<string>? _vpnProtocols;

        /// <summary>
        /// vpnProtocols for the P2SVpnServerConfiguration.
        /// </summary>
        public InputList<string> VpnProtocols
        {
            get => _vpnProtocols ?? (_vpnProtocols = new InputList<string>());
            set => _vpnProtocols = value;
        }

        public P2sVpnServerConfigurationArgs()
        {
        }
    }
}
