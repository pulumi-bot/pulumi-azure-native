// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20190201
{
    /// <summary>
    /// ExpressRouteCircuit resource
    /// </summary>
    public partial class ExpressRouteCircuit : Pulumi.CustomResource
    {
        /// <summary>
        /// Gets a unique read-only string that changes whenever the resource is updated.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

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
        /// Properties of the express route circuit.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.ExpressRouteCircuitPropertiesFormatResponseResult> Properties { get; private set; } = null!;

        /// <summary>
        /// The SKU.
        /// </summary>
        [Output("sku")]
        public Output<Outputs.ExpressRouteCircuitSkuResponseResult?> Sku { get; private set; } = null!;

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
        /// Create a ExpressRouteCircuit resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ExpressRouteCircuit(string name, ExpressRouteCircuitArgs args, CustomResourceOptions? options = null)
            : base("azurerm:network/v20190201:ExpressRouteCircuit", name, args ?? new ExpressRouteCircuitArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ExpressRouteCircuit(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:network/v20190201:ExpressRouteCircuit", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing ExpressRouteCircuit resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ExpressRouteCircuit Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ExpressRouteCircuit(name, id, options);
        }
    }

    public sealed class ExpressRouteCircuitArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Allow classic operations
        /// </summary>
        [Input("allowClassicOperations")]
        public Input<bool>? AllowClassicOperations { get; set; }

        [Input("authorizations")]
        private InputList<Inputs.ExpressRouteCircuitAuthorizationArgs>? _authorizations;

        /// <summary>
        /// The list of authorizations.
        /// </summary>
        public InputList<Inputs.ExpressRouteCircuitAuthorizationArgs> Authorizations
        {
            get => _authorizations ?? (_authorizations = new InputList<Inputs.ExpressRouteCircuitAuthorizationArgs>());
            set => _authorizations = value;
        }

        /// <summary>
        /// The bandwidth of the circuit when the circuit is provisioned on an ExpressRoutePort resource.
        /// </summary>
        [Input("bandwidthInGbps")]
        public Input<double>? BandwidthInGbps { get; set; }

        /// <summary>
        /// The CircuitProvisioningState state of the resource.
        /// </summary>
        [Input("circuitProvisioningState")]
        public Input<string>? CircuitProvisioningState { get; set; }

        /// <summary>
        /// The reference to the ExpressRoutePort resource when the circuit is provisioned on an ExpressRoutePort resource.
        /// </summary>
        [Input("expressRoutePort")]
        public Input<Inputs.SubResourceArgs>? ExpressRoutePort { get; set; }

        /// <summary>
        /// The GatewayManager Etag.
        /// </summary>
        [Input("gatewayManagerEtag")]
        public Input<string>? GatewayManagerEtag { get; set; }

        /// <summary>
        /// Flag denoting Global reach status.
        /// </summary>
        [Input("globalReachEnabled")]
        public Input<bool>? GlobalReachEnabled { get; set; }

        /// <summary>
        /// Resource ID.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Resource location.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the circuit.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("peerings")]
        private InputList<Inputs.ExpressRouteCircuitPeeringArgs>? _peerings;

        /// <summary>
        /// The list of peerings.
        /// </summary>
        public InputList<Inputs.ExpressRouteCircuitPeeringArgs> Peerings
        {
            get => _peerings ?? (_peerings = new InputList<Inputs.ExpressRouteCircuitPeeringArgs>());
            set => _peerings = value;
        }

        /// <summary>
        /// Gets the provisioning state of the public IP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
        /// </summary>
        [Input("provisioningState")]
        public Input<string>? ProvisioningState { get; set; }

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The ServiceKey.
        /// </summary>
        [Input("serviceKey")]
        public Input<string>? ServiceKey { get; set; }

        /// <summary>
        /// The ServiceProviderNotes.
        /// </summary>
        [Input("serviceProviderNotes")]
        public Input<string>? ServiceProviderNotes { get; set; }

        /// <summary>
        /// The ServiceProviderProperties.
        /// </summary>
        [Input("serviceProviderProperties")]
        public Input<Inputs.ExpressRouteCircuitServiceProviderPropertiesArgs>? ServiceProviderProperties { get; set; }

        /// <summary>
        /// The ServiceProviderProvisioningState state of the resource.
        /// </summary>
        [Input("serviceProviderProvisioningState")]
        public Input<string>? ServiceProviderProvisioningState { get; set; }

        /// <summary>
        /// The SKU.
        /// </summary>
        [Input("sku")]
        public Input<Inputs.ExpressRouteCircuitSkuArgs>? Sku { get; set; }

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

        public ExpressRouteCircuitArgs()
        {
        }
    }
}
