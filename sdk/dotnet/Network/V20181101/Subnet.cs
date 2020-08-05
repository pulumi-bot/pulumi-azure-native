// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20181101
{
    /// <summary>
    /// Subnet in a virtual network resource.
    /// </summary>
    public partial class Subnet : Pulumi.CustomResource
    {
        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        [Output("etag")]
        public Output<string?> Etag { get; private set; } = null!;

        /// <summary>
        /// The name of the resource that is unique within a resource group. This name can be used to access the resource.
        /// </summary>
        [Output("name")]
        public Output<string?> Name { get; private set; } = null!;

        /// <summary>
        /// Properties of the subnet.
        /// </summary>
        [Output("properties")]
        public Output<Outputs.SubnetPropertiesFormatResponseResult> Properties { get; private set; } = null!;


        /// <summary>
        /// Create a Subnet resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Subnet(string name, SubnetArgs args, CustomResourceOptions? options = null)
            : base("azurerm:network/v20181101:Subnet", name, args ?? new SubnetArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Subnet(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:network/v20181101:Subnet", name, null, MakeResourceOptions(options, id))
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
        /// The address prefix for the subnet.
        /// </summary>
        [Input("addressPrefix")]
        public Input<string>? AddressPrefix { get; set; }

        [Input("addressPrefixes")]
        private InputList<string>? _addressPrefixes;

        /// <summary>
        /// List of  address prefixes for the subnet.
        /// </summary>
        public InputList<string> AddressPrefixes
        {
            get => _addressPrefixes ?? (_addressPrefixes = new InputList<string>());
            set => _addressPrefixes = value;
        }

        [Input("delegations")]
        private InputList<Inputs.DelegationArgs>? _delegations;

        /// <summary>
        /// Gets an array of references to the delegations on the subnet.
        /// </summary>
        public InputList<Inputs.DelegationArgs> Delegations
        {
            get => _delegations ?? (_delegations = new InputList<Inputs.DelegationArgs>());
            set => _delegations = value;
        }

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
        /// The name of the subnet.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The reference of the NetworkSecurityGroup resource.
        /// </summary>
        [Input("networkSecurityGroup")]
        public Input<Inputs.NetworkSecurityGroupArgs>? NetworkSecurityGroup { get; set; }

        /// <summary>
        /// The provisioning state of the resource.
        /// </summary>
        [Input("provisioningState")]
        public Input<string>? ProvisioningState { get; set; }

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("resourceNavigationLinks")]
        private InputList<Inputs.ResourceNavigationLinkArgs>? _resourceNavigationLinks;

        /// <summary>
        /// Gets an array of references to the external resources using subnet.
        /// </summary>
        public InputList<Inputs.ResourceNavigationLinkArgs> ResourceNavigationLinks
        {
            get => _resourceNavigationLinks ?? (_resourceNavigationLinks = new InputList<Inputs.ResourceNavigationLinkArgs>());
            set => _resourceNavigationLinks = value;
        }

        /// <summary>
        /// The reference of the RouteTable resource.
        /// </summary>
        [Input("routeTable")]
        public Input<Inputs.RouteTableArgs>? RouteTable { get; set; }

        [Input("serviceAssociationLinks")]
        private InputList<Inputs.ServiceAssociationLinkArgs>? _serviceAssociationLinks;

        /// <summary>
        /// Gets an array of references to services injecting into this subnet.
        /// </summary>
        public InputList<Inputs.ServiceAssociationLinkArgs> ServiceAssociationLinks
        {
            get => _serviceAssociationLinks ?? (_serviceAssociationLinks = new InputList<Inputs.ServiceAssociationLinkArgs>());
            set => _serviceAssociationLinks = value;
        }

        [Input("serviceEndpointPolicies")]
        private InputList<Inputs.ServiceEndpointPolicyArgs>? _serviceEndpointPolicies;

        /// <summary>
        /// An array of service endpoint policies.
        /// </summary>
        public InputList<Inputs.ServiceEndpointPolicyArgs> ServiceEndpointPolicies
        {
            get => _serviceEndpointPolicies ?? (_serviceEndpointPolicies = new InputList<Inputs.ServiceEndpointPolicyArgs>());
            set => _serviceEndpointPolicies = value;
        }

        [Input("serviceEndpoints")]
        private InputList<Inputs.ServiceEndpointPropertiesFormatArgs>? _serviceEndpoints;

        /// <summary>
        /// An array of service endpoints.
        /// </summary>
        public InputList<Inputs.ServiceEndpointPropertiesFormatArgs> ServiceEndpoints
        {
            get => _serviceEndpoints ?? (_serviceEndpoints = new InputList<Inputs.ServiceEndpointPropertiesFormatArgs>());
            set => _serviceEndpoints = value;
        }

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
