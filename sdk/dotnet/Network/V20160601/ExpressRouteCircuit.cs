// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20160601
{
    /// <summary>
    /// ExpressRouteCircuit resource
    /// </summary>
    public partial class ExpressRouteCircuit : Pulumi.CustomResource
    {
        /// <summary>
        /// allow classic operations
        /// </summary>
        [Output("allowClassicOperations")]
        public Output<bool?> AllowClassicOperations { get; private set; } = null!;

        /// <summary>
        /// Gets or sets list of authorizations
        /// </summary>
        [Output("authorizations")]
        public Output<ImmutableArray<Outputs.ExpressRouteCircuitAuthorizationResponseResult>> Authorizations { get; private set; } = null!;

        /// <summary>
        /// Gets or sets CircuitProvisioningState state of the resource 
        /// </summary>
        [Output("circuitProvisioningState")]
        public Output<string?> CircuitProvisioningState { get; private set; } = null!;

        /// <summary>
        /// Gets a unique read-only string that changes whenever the resource is updated
        /// </summary>
        [Output("etag")]
        public Output<string?> Etag { get; private set; } = null!;

        /// <summary>
        /// Gets or sets the GatewayManager Etag
        /// </summary>
        [Output("gatewayManagerEtag")]
        public Output<string?> GatewayManagerEtag { get; private set; } = null!;

        /// <summary>
        /// Resource location
        /// </summary>
        [Output("location")]
        public Output<string?> Location { get; private set; } = null!;

        /// <summary>
        /// Resource name
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Gets or sets list of peerings
        /// </summary>
        [Output("peerings")]
        public Output<ImmutableArray<Outputs.ExpressRouteCircuitPeeringResponseResult>> Peerings { get; private set; } = null!;

        /// <summary>
        /// Gets provisioning state of the PublicIP resource Updating/Deleting/Failed
        /// </summary>
        [Output("provisioningState")]
        public Output<string?> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// Gets or sets ServiceKey
        /// </summary>
        [Output("serviceKey")]
        public Output<string?> ServiceKey { get; private set; } = null!;

        /// <summary>
        /// Gets or sets ServiceProviderNotes
        /// </summary>
        [Output("serviceProviderNotes")]
        public Output<string?> ServiceProviderNotes { get; private set; } = null!;

        /// <summary>
        /// Gets or sets ServiceProviderProperties
        /// </summary>
        [Output("serviceProviderProperties")]
        public Output<Outputs.ExpressRouteCircuitServiceProviderPropertiesResponseResult?> ServiceProviderProperties { get; private set; } = null!;

        /// <summary>
        /// Gets or sets ServiceProviderProvisioningState state of the resource 
        /// </summary>
        [Output("serviceProviderProvisioningState")]
        public Output<string?> ServiceProviderProvisioningState { get; private set; } = null!;

        /// <summary>
        /// Gets or sets sku
        /// </summary>
        [Output("sku")]
        public Output<Outputs.ExpressRouteCircuitSkuResponseResult?> Sku { get; private set; } = null!;

        /// <summary>
        /// Resource tags
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Resource type
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
            : base("azurerm:network/v20160601:ExpressRouteCircuit", name, args ?? new ExpressRouteCircuitArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ExpressRouteCircuit(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:network/v20160601:ExpressRouteCircuit", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:network/latest:ExpressRouteCircuit"},
                    new Pulumi.Alias { Type = "azurerm:network/v20150501preview:ExpressRouteCircuit"},
                    new Pulumi.Alias { Type = "azurerm:network/v20150615:ExpressRouteCircuit"},
                    new Pulumi.Alias { Type = "azurerm:network/v20160330:ExpressRouteCircuit"},
                    new Pulumi.Alias { Type = "azurerm:network/v20160901:ExpressRouteCircuit"},
                    new Pulumi.Alias { Type = "azurerm:network/v20161201:ExpressRouteCircuit"},
                    new Pulumi.Alias { Type = "azurerm:network/v20170301:ExpressRouteCircuit"},
                    new Pulumi.Alias { Type = "azurerm:network/v20170601:ExpressRouteCircuit"},
                    new Pulumi.Alias { Type = "azurerm:network/v20170801:ExpressRouteCircuit"},
                    new Pulumi.Alias { Type = "azurerm:network/v20170901:ExpressRouteCircuit"},
                    new Pulumi.Alias { Type = "azurerm:network/v20171001:ExpressRouteCircuit"},
                    new Pulumi.Alias { Type = "azurerm:network/v20171101:ExpressRouteCircuit"},
                    new Pulumi.Alias { Type = "azurerm:network/v20180101:ExpressRouteCircuit"},
                    new Pulumi.Alias { Type = "azurerm:network/v20180201:ExpressRouteCircuit"},
                    new Pulumi.Alias { Type = "azurerm:network/v20180401:ExpressRouteCircuit"},
                    new Pulumi.Alias { Type = "azurerm:network/v20180601:ExpressRouteCircuit"},
                    new Pulumi.Alias { Type = "azurerm:network/v20180701:ExpressRouteCircuit"},
                    new Pulumi.Alias { Type = "azurerm:network/v20180801:ExpressRouteCircuit"},
                    new Pulumi.Alias { Type = "azurerm:network/v20181001:ExpressRouteCircuit"},
                    new Pulumi.Alias { Type = "azurerm:network/v20181101:ExpressRouteCircuit"},
                    new Pulumi.Alias { Type = "azurerm:network/v20181201:ExpressRouteCircuit"},
                    new Pulumi.Alias { Type = "azurerm:network/v20190201:ExpressRouteCircuit"},
                    new Pulumi.Alias { Type = "azurerm:network/v20190401:ExpressRouteCircuit"},
                    new Pulumi.Alias { Type = "azurerm:network/v20190601:ExpressRouteCircuit"},
                    new Pulumi.Alias { Type = "azurerm:network/v20190701:ExpressRouteCircuit"},
                    new Pulumi.Alias { Type = "azurerm:network/v20190801:ExpressRouteCircuit"},
                    new Pulumi.Alias { Type = "azurerm:network/v20190901:ExpressRouteCircuit"},
                    new Pulumi.Alias { Type = "azurerm:network/v20191101:ExpressRouteCircuit"},
                    new Pulumi.Alias { Type = "azurerm:network/v20191201:ExpressRouteCircuit"},
                    new Pulumi.Alias { Type = "azurerm:network/v20200301:ExpressRouteCircuit"},
                    new Pulumi.Alias { Type = "azurerm:network/v20200401:ExpressRouteCircuit"},
                    new Pulumi.Alias { Type = "azurerm:network/v20200501:ExpressRouteCircuit"},
                    new Pulumi.Alias { Type = "azurerm:network/v20200601:ExpressRouteCircuit"},
                },
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
        /// allow classic operations
        /// </summary>
        [Input("allowClassicOperations")]
        public Input<bool>? AllowClassicOperations { get; set; }

        [Input("authorizations")]
        private InputList<Inputs.ExpressRouteCircuitAuthorizationArgs>? _authorizations;

        /// <summary>
        /// Gets or sets list of authorizations
        /// </summary>
        public InputList<Inputs.ExpressRouteCircuitAuthorizationArgs> Authorizations
        {
            get => _authorizations ?? (_authorizations = new InputList<Inputs.ExpressRouteCircuitAuthorizationArgs>());
            set => _authorizations = value;
        }

        /// <summary>
        /// The name of the circuit.
        /// </summary>
        [Input("circuitName", required: true)]
        public Input<string> CircuitName { get; set; } = null!;

        /// <summary>
        /// Gets or sets CircuitProvisioningState state of the resource 
        /// </summary>
        [Input("circuitProvisioningState")]
        public Input<string>? CircuitProvisioningState { get; set; }

        /// <summary>
        /// Gets a unique read-only string that changes whenever the resource is updated
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// Gets or sets the GatewayManager Etag
        /// </summary>
        [Input("gatewayManagerEtag")]
        public Input<string>? GatewayManagerEtag { get; set; }

        /// <summary>
        /// Resource Id
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Resource location
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        [Input("peerings")]
        private InputList<Inputs.ExpressRouteCircuitPeeringArgs>? _peerings;

        /// <summary>
        /// Gets or sets list of peerings
        /// </summary>
        public InputList<Inputs.ExpressRouteCircuitPeeringArgs> Peerings
        {
            get => _peerings ?? (_peerings = new InputList<Inputs.ExpressRouteCircuitPeeringArgs>());
            set => _peerings = value;
        }

        /// <summary>
        /// Gets provisioning state of the PublicIP resource Updating/Deleting/Failed
        /// </summary>
        [Input("provisioningState")]
        public Input<string>? ProvisioningState { get; set; }

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Gets or sets ServiceKey
        /// </summary>
        [Input("serviceKey")]
        public Input<string>? ServiceKey { get; set; }

        /// <summary>
        /// Gets or sets ServiceProviderNotes
        /// </summary>
        [Input("serviceProviderNotes")]
        public Input<string>? ServiceProviderNotes { get; set; }

        /// <summary>
        /// Gets or sets ServiceProviderProperties
        /// </summary>
        [Input("serviceProviderProperties")]
        public Input<Inputs.ExpressRouteCircuitServiceProviderPropertiesArgs>? ServiceProviderProperties { get; set; }

        /// <summary>
        /// Gets or sets ServiceProviderProvisioningState state of the resource 
        /// </summary>
        [Input("serviceProviderProvisioningState")]
        public Input<string>? ServiceProviderProvisioningState { get; set; }

        /// <summary>
        /// Gets or sets sku
        /// </summary>
        [Input("sku")]
        public Input<Inputs.ExpressRouteCircuitSkuArgs>? Sku { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Resource tags
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
