// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20180601
{
    /// <summary>
    /// VpnConnection Resource.
    /// </summary>
    public partial class VpnConnection : Pulumi.CustomResource
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
        /// Parameters for VpnConnection
        /// </summary>
        [Output("properties")]
        public Output<Outputs.VpnConnectionPropertiesResponseResult> Properties { get; private set; } = null!;


        /// <summary>
        /// Create a VpnConnection resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VpnConnection(string name, VpnConnectionArgs args, CustomResourceOptions? options = null)
            : base("azurerm:network/v20180601:VpnConnection", name, args ?? new VpnConnectionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VpnConnection(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:network/v20180601:VpnConnection", name, null, MakeResourceOptions(options, id))
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
        /// The connection status.
        /// </summary>
        [Input("connectionStatus")]
        public Input<string>? ConnectionStatus { get; set; }

        /// <summary>
        /// EnableBgp flag
        /// </summary>
        [Input("enableBgp")]
        public Input<bool>? EnableBgp { get; set; }

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
        /// The name of the connection.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The provisioning state of the resource.
        /// </summary>
        [Input("provisioningState")]
        public Input<string>? ProvisioningState { get; set; }

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

        public VpnConnectionArgs()
        {
        }
    }
}
