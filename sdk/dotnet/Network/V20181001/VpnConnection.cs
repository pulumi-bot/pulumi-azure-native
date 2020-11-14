// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Network.V20181001
{
    /// <summary>
    /// VpnConnection Resource.
    /// </summary>
    public partial class VpnConnection : Pulumi.CustomResource
    {
        /// <summary>
        /// Expected bandwidth in MBPS.
        /// </summary>
        [Output("connectionBandwidth")]
        public Output<int?> ConnectionBandwidth { get; private set; } = null!;

        /// <summary>
        /// The connection status.
        /// </summary>
        [Output("connectionStatus")]
        public Output<string> ConnectionStatus { get; private set; } = null!;

        /// <summary>
        /// Egress bytes transferred.
        /// </summary>
        [Output("egressBytesTransferred")]
        public Output<int> EgressBytesTransferred { get; private set; } = null!;

        /// <summary>
        /// EnableBgp flag
        /// </summary>
        [Output("enableBgp")]
        public Output<bool?> EnableBgp { get; private set; } = null!;

        /// <summary>
        /// Enable internet security
        /// </summary>
        [Output("enableInternetSecurity")]
        public Output<bool?> EnableInternetSecurity { get; private set; } = null!;

        /// <summary>
        /// EnableBgp flag
        /// </summary>
        [Output("enableRateLimiting")]
        public Output<bool?> EnableRateLimiting { get; private set; } = null!;

        /// <summary>
        /// Gets a unique read-only string that changes whenever the resource is updated.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// Ingress bytes transferred.
        /// </summary>
        [Output("ingressBytesTransferred")]
        public Output<int> IngressBytesTransferred { get; private set; } = null!;

        /// <summary>
        /// The IPSec Policies to be considered by this connection.
        /// </summary>
        [Output("ipsecPolicies")]
        public Output<ImmutableArray<Outputs.IpsecPolicyResponse>> IpsecPolicies { get; private set; } = null!;

        /// <summary>
        /// The name of the resource that is unique within a resource group. This name can be used to access the resource.
        /// </summary>
        [Output("name")]
        public Output<string?> Name { get; private set; } = null!;

        /// <summary>
        /// The provisioning state of the resource.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// Id of the connected vpn site.
        /// </summary>
        [Output("remoteVpnSite")]
        public Output<Outputs.SubResourceResponse?> RemoteVpnSite { get; private set; } = null!;

        /// <summary>
        /// routing weight for vpn connection.
        /// </summary>
        [Output("routingWeight")]
        public Output<int?> RoutingWeight { get; private set; } = null!;

        /// <summary>
        /// SharedKey for the vpn connection.
        /// </summary>
        [Output("sharedKey")]
        public Output<string?> SharedKey { get; private set; } = null!;

        /// <summary>
        /// Connection protocol used for this connection
        /// </summary>
        [Output("vpnConnectionProtocolType")]
        public Output<string?> VpnConnectionProtocolType { get; private set; } = null!;


        /// <summary>
        /// Create a VpnConnection resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VpnConnection(string name, VpnConnectionArgs args, CustomResourceOptions? options = null)
            : base("azure-nextgen:network/v20181001:VpnConnection", name, args ?? new VpnConnectionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VpnConnection(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-nextgen:network/v20181001:VpnConnection", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azure-nextgen:network/latest:VpnConnection"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20180401:VpnConnection"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20180601:VpnConnection"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20180701:VpnConnection"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20180801:VpnConnection"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20181101:VpnConnection"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20181201:VpnConnection"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20190201:VpnConnection"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20190401:VpnConnection"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20190601:VpnConnection"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20190701:VpnConnection"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20190801:VpnConnection"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20190901:VpnConnection"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20191101:VpnConnection"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20191201:VpnConnection"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20200301:VpnConnection"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20200401:VpnConnection"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20200501:VpnConnection"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20200601:VpnConnection"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20200701:VpnConnection"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing VpnConnection resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VpnConnection Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new VpnConnection(name, id, options);
        }
    }

    public sealed class VpnConnectionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Expected bandwidth in MBPS.
        /// </summary>
        [Input("connectionBandwidth")]
        public Input<int>? ConnectionBandwidth { get; set; }

        /// <summary>
        /// The name of the connection.
        /// </summary>
        [Input("connectionName", required: true)]
        public Input<string> ConnectionName { get; set; } = null!;

        /// <summary>
        /// EnableBgp flag
        /// </summary>
        [Input("enableBgp")]
        public Input<bool>? EnableBgp { get; set; }

        /// <summary>
        /// Enable internet security
        /// </summary>
        [Input("enableInternetSecurity")]
        public Input<bool>? EnableInternetSecurity { get; set; }

        /// <summary>
        /// EnableBgp flag
        /// </summary>
        [Input("enableRateLimiting")]
        public Input<bool>? EnableRateLimiting { get; set; }

        /// <summary>
        /// The name of the gateway.
        /// </summary>
        [Input("gatewayName", required: true)]
        public Input<string> GatewayName { get; set; } = null!;

        /// <summary>
        /// Resource ID.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("ipsecPolicies")]
        private InputList<Inputs.IpsecPolicyArgs>? _ipsecPolicies;

        /// <summary>
        /// The IPSec Policies to be considered by this connection.
        /// </summary>
        public InputList<Inputs.IpsecPolicyArgs> IpsecPolicies
        {
            get => _ipsecPolicies ?? (_ipsecPolicies = new InputList<Inputs.IpsecPolicyArgs>());
            set => _ipsecPolicies = value;
        }

        /// <summary>
        /// The name of the resource that is unique within a resource group. This name can be used to access the resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Id of the connected vpn site.
        /// </summary>
        [Input("remoteVpnSite")]
        public Input<Inputs.SubResourceArgs>? RemoteVpnSite { get; set; }

        /// <summary>
        /// The resource group name of the VpnGateway.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// routing weight for vpn connection.
        /// </summary>
        [Input("routingWeight")]
        public Input<int>? RoutingWeight { get; set; }

        /// <summary>
        /// SharedKey for the vpn connection.
        /// </summary>
        [Input("sharedKey")]
        public Input<string>? SharedKey { get; set; }

        /// <summary>
        /// Connection protocol used for this connection
        /// </summary>
        [Input("vpnConnectionProtocolType")]
        public Input<string>? VpnConnectionProtocolType { get; set; }

        public VpnConnectionArgs()
        {
        }
    }
}
