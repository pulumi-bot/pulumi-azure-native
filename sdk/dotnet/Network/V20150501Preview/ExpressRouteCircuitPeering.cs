// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Network.V20150501Preview
{
    /// <summary>
    /// Peering in a ExpressRouteCircuit resource
    /// </summary>
    public partial class ExpressRouteCircuitPeering : Pulumi.CustomResource
    {
        /// <summary>
        /// Gets or sets the azure ASN
        /// </summary>
        [Output("azureASN")]
        public Output<int?> AzureASN { get; private set; } = null!;

        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated
        /// </summary>
        [Output("etag")]
        public Output<string?> Etag { get; private set; } = null!;

        /// <summary>
        /// Gets or sets the Microsoft peering config
        /// </summary>
        [Output("microsoftPeeringConfig")]
        public Output<Outputs.ExpressRouteCircuitPeeringConfigResponse?> MicrosoftPeeringConfig { get; private set; } = null!;

        /// <summary>
        /// Gets name of the resource that is unique within a resource group. This name can be used to access the resource
        /// </summary>
        [Output("name")]
        public Output<string?> Name { get; private set; } = null!;

        /// <summary>
        /// Gets or sets the peer ASN
        /// </summary>
        [Output("peerASN")]
        public Output<int?> PeerASN { get; private set; } = null!;

        /// <summary>
        /// Gets or sets PeeringType
        /// </summary>
        [Output("peeringType")]
        public Output<string?> PeeringType { get; private set; } = null!;

        /// <summary>
        /// Gets or sets the primary port
        /// </summary>
        [Output("primaryAzurePort")]
        public Output<string?> PrimaryAzurePort { get; private set; } = null!;

        /// <summary>
        /// Gets or sets the primary address prefix
        /// </summary>
        [Output("primaryPeerAddressPrefix")]
        public Output<string?> PrimaryPeerAddressPrefix { get; private set; } = null!;

        /// <summary>
        /// Gets or sets Provisioning state of the PublicIP resource Updating/Deleting/Failed
        /// </summary>
        [Output("provisioningState")]
        public Output<string?> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// Gets or sets the secondary port
        /// </summary>
        [Output("secondaryAzurePort")]
        public Output<string?> SecondaryAzurePort { get; private set; } = null!;

        /// <summary>
        /// Gets or sets the secondary address prefix
        /// </summary>
        [Output("secondaryPeerAddressPrefix")]
        public Output<string?> SecondaryPeerAddressPrefix { get; private set; } = null!;

        /// <summary>
        /// Gets or sets the shared key
        /// </summary>
        [Output("sharedKey")]
        public Output<string?> SharedKey { get; private set; } = null!;

        /// <summary>
        /// Gets or sets state of Peering
        /// </summary>
        [Output("state")]
        public Output<string?> State { get; private set; } = null!;

        /// <summary>
        /// Gets or peering stats
        /// </summary>
        [Output("stats")]
        public Output<Outputs.ExpressRouteCircuitStatsResponse?> Stats { get; private set; } = null!;

        /// <summary>
        /// Gets or sets the vlan id
        /// </summary>
        [Output("vlanId")]
        public Output<int?> VlanId { get; private set; } = null!;


        /// <summary>
        /// Create a ExpressRouteCircuitPeering resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ExpressRouteCircuitPeering(string name, ExpressRouteCircuitPeeringArgs args, CustomResourceOptions? options = null)
            : base("azure-nextgen:network/v20150501preview:ExpressRouteCircuitPeering", name, args ?? new ExpressRouteCircuitPeeringArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ExpressRouteCircuitPeering(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-nextgen:network/v20150501preview:ExpressRouteCircuitPeering", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azure-nextgen:network/latest:ExpressRouteCircuitPeering"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20150615:ExpressRouteCircuitPeering"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20160330:ExpressRouteCircuitPeering"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20160601:ExpressRouteCircuitPeering"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20160901:ExpressRouteCircuitPeering"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20161201:ExpressRouteCircuitPeering"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20170301:ExpressRouteCircuitPeering"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20170601:ExpressRouteCircuitPeering"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20170801:ExpressRouteCircuitPeering"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20170901:ExpressRouteCircuitPeering"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20171001:ExpressRouteCircuitPeering"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20171101:ExpressRouteCircuitPeering"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20180101:ExpressRouteCircuitPeering"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20180201:ExpressRouteCircuitPeering"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20180401:ExpressRouteCircuitPeering"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20180601:ExpressRouteCircuitPeering"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20180701:ExpressRouteCircuitPeering"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20180801:ExpressRouteCircuitPeering"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20181001:ExpressRouteCircuitPeering"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20181101:ExpressRouteCircuitPeering"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20181201:ExpressRouteCircuitPeering"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20190201:ExpressRouteCircuitPeering"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20190401:ExpressRouteCircuitPeering"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20190601:ExpressRouteCircuitPeering"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20190701:ExpressRouteCircuitPeering"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20190801:ExpressRouteCircuitPeering"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20190901:ExpressRouteCircuitPeering"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20191101:ExpressRouteCircuitPeering"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20191201:ExpressRouteCircuitPeering"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20200301:ExpressRouteCircuitPeering"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20200401:ExpressRouteCircuitPeering"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20200501:ExpressRouteCircuitPeering"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20200601:ExpressRouteCircuitPeering"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20200701:ExpressRouteCircuitPeering"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ExpressRouteCircuitPeering resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ExpressRouteCircuitPeering Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ExpressRouteCircuitPeering(name, id, options);
        }
    }

    public sealed class ExpressRouteCircuitPeeringArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Gets or sets the azure ASN
        /// </summary>
        [Input("azureASN")]
        public Input<int>? AzureASN { get; set; }

        /// <summary>
        /// The name of the express route circuit.
        /// </summary>
        [Input("circuitName", required: true)]
        public Input<string> CircuitName { get; set; } = null!;

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

        /// <summary>
        /// Gets or sets the Microsoft peering config
        /// </summary>
        [Input("microsoftPeeringConfig")]
        public Input<Inputs.ExpressRouteCircuitPeeringConfigArgs>? MicrosoftPeeringConfig { get; set; }

        /// <summary>
        /// Gets name of the resource that is unique within a resource group. This name can be used to access the resource
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Gets or sets the peer ASN
        /// </summary>
        [Input("peerASN")]
        public Input<int>? PeerASN { get; set; }

        /// <summary>
        /// The name of the peering.
        /// </summary>
        [Input("peeringName", required: true)]
        public Input<string> PeeringName { get; set; } = null!;

        /// <summary>
        /// Gets or sets PeeringType
        /// </summary>
        [Input("peeringType")]
        public Input<string>? PeeringType { get; set; }

        /// <summary>
        /// Gets or sets the primary port
        /// </summary>
        [Input("primaryAzurePort")]
        public Input<string>? PrimaryAzurePort { get; set; }

        /// <summary>
        /// Gets or sets the primary address prefix
        /// </summary>
        [Input("primaryPeerAddressPrefix")]
        public Input<string>? PrimaryPeerAddressPrefix { get; set; }

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
        /// Gets or sets the secondary port
        /// </summary>
        [Input("secondaryAzurePort")]
        public Input<string>? SecondaryAzurePort { get; set; }

        /// <summary>
        /// Gets or sets the secondary address prefix
        /// </summary>
        [Input("secondaryPeerAddressPrefix")]
        public Input<string>? SecondaryPeerAddressPrefix { get; set; }

        /// <summary>
        /// Gets or sets the shared key
        /// </summary>
        [Input("sharedKey")]
        public Input<string>? SharedKey { get; set; }

        /// <summary>
        /// Gets or sets state of Peering
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        /// <summary>
        /// Gets or peering stats
        /// </summary>
        [Input("stats")]
        public Input<Inputs.ExpressRouteCircuitStatsArgs>? Stats { get; set; }

        /// <summary>
        /// Gets or sets the vlan id
        /// </summary>
        [Input("vlanId")]
        public Input<int>? VlanId { get; set; }

        public ExpressRouteCircuitPeeringArgs()
        {
        }
    }
}
