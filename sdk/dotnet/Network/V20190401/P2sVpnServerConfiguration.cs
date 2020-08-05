// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20190401
{
    /// <summary>
    /// P2SVpnServerConfiguration Resource.
    /// </summary>
    public partial class P2sVpnServerConfiguration : Pulumi.CustomResource
    {
        /// <summary>
        /// Gets a unique read-only string that changes whenever the resource is updated.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// The name of the resource that is unique within a resource group. This name can be used to access the resource.
        /// </summary>
        [Output("name")]
        public Output<string?> Name { get; private set; } = null!;

        /// <summary>
        /// Properties of the P2SVpnServer configuration.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.P2SVpnServerConfigurationPropertiesResponseResult> Properties { get; private set; } = null!;


        /// <summary>
        /// Create a P2sVpnServerConfiguration resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public P2sVpnServerConfiguration(string name, P2sVpnServerConfigurationArgs args, CustomResourceOptions? options = null)
            : base("azurerm:network/v20190401:P2sVpnServerConfiguration", name, args ?? new P2sVpnServerConfigurationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private P2sVpnServerConfiguration(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:network/v20190401:P2sVpnServerConfiguration", name, null, MakeResourceOptions(options, id))
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
        /// The name of the P2SVpnServerConfiguration.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

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
        /// VPN protocols for the P2SVpnServerConfiguration.
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
