// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Network.V20190701
{
    /// <summary>
    /// Subnet in a virtual network resource.
    /// </summary>
    public partial class Subnet : Pulumi.CustomResource
    {
        /// <summary>
        /// The address prefix for the subnet.
        /// </summary>
        [Output("addressPrefix")]
        public Output<string?> AddressPrefix { get; private set; } = null!;

        /// <summary>
        /// List of address prefixes for the subnet.
        /// </summary>
        [Output("addressPrefixes")]
        public Output<ImmutableArray<string>> AddressPrefixes { get; private set; } = null!;

        /// <summary>
        /// An array of references to the delegations on the subnet.
        /// </summary>
        [Output("delegations")]
        public Output<ImmutableArray<Outputs.DelegationResponse>> Delegations { get; private set; } = null!;

        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        [Output("etag")]
        public Output<string?> Etag { get; private set; } = null!;

        /// <summary>
        /// Array of IP configuration profiles which reference this subnet.
        /// </summary>
        [Output("ipConfigurationProfiles")]
        public Output<ImmutableArray<Outputs.IPConfigurationProfileResponse>> IpConfigurationProfiles { get; private set; } = null!;

        /// <summary>
        /// An array of references to the network interface IP configurations using subnet.
        /// </summary>
        [Output("ipConfigurations")]
        public Output<ImmutableArray<Outputs.IPConfigurationResponse>> IpConfigurations { get; private set; } = null!;

        /// <summary>
        /// The name of the resource that is unique within a resource group. This name can be used to access the resource.
        /// </summary>
        [Output("name")]
        public Output<string?> Name { get; private set; } = null!;

        /// <summary>
        /// Nat gateway associated with this subnet.
        /// </summary>
        [Output("natGateway")]
        public Output<Outputs.SubResourceResponse?> NatGateway { get; private set; } = null!;

        /// <summary>
        /// The reference of the NetworkSecurityGroup resource.
        /// </summary>
        [Output("networkSecurityGroup")]
        public Output<Outputs.NetworkSecurityGroupResponse?> NetworkSecurityGroup { get; private set; } = null!;

        /// <summary>
        /// Enable or Disable apply network policies on private end point in the subnet.
        /// </summary>
        [Output("privateEndpointNetworkPolicies")]
        public Output<string?> PrivateEndpointNetworkPolicies { get; private set; } = null!;

        /// <summary>
        /// An array of references to private endpoints.
        /// </summary>
        [Output("privateEndpoints")]
        public Output<ImmutableArray<Outputs.PrivateEndpointResponse>> PrivateEndpoints { get; private set; } = null!;

        /// <summary>
        /// Enable or Disable apply network policies on private link service in the subnet.
        /// </summary>
        [Output("privateLinkServiceNetworkPolicies")]
        public Output<string?> PrivateLinkServiceNetworkPolicies { get; private set; } = null!;

        /// <summary>
        /// The provisioning state of the subnet resource.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// A read-only string identifying the intention of use for this subnet based on delegations and other user-defined properties.
        /// </summary>
        [Output("purpose")]
        public Output<string> Purpose { get; private set; } = null!;

        /// <summary>
        /// An array of references to the external resources using subnet.
        /// </summary>
        [Output("resourceNavigationLinks")]
        public Output<ImmutableArray<Outputs.ResourceNavigationLinkResponse>> ResourceNavigationLinks { get; private set; } = null!;

        /// <summary>
        /// The reference of the RouteTable resource.
        /// </summary>
        [Output("routeTable")]
        public Output<Outputs.RouteTableResponse?> RouteTable { get; private set; } = null!;

        /// <summary>
        /// An array of references to services injecting into this subnet.
        /// </summary>
        [Output("serviceAssociationLinks")]
        public Output<ImmutableArray<Outputs.ServiceAssociationLinkResponse>> ServiceAssociationLinks { get; private set; } = null!;

        /// <summary>
        /// An array of service endpoint policies.
        /// </summary>
        [Output("serviceEndpointPolicies")]
        public Output<ImmutableArray<Outputs.ServiceEndpointPolicyResponse>> ServiceEndpointPolicies { get; private set; } = null!;

        /// <summary>
        /// An array of service endpoints.
        /// </summary>
        [Output("serviceEndpoints")]
        public Output<ImmutableArray<Outputs.ServiceEndpointPropertiesFormatResponse>> ServiceEndpoints { get; private set; } = null!;


        /// <summary>
        /// Create a Subnet resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Subnet(string name, SubnetArgs args, CustomResourceOptions? options = null)
            : base("azure-nextgen:network/v20190701:Subnet", name, args ?? new SubnetArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Subnet(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-nextgen:network/v20190701:Subnet", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azure-nextgen:network/latest:Subnet"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20150501preview:Subnet"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20150615:Subnet"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20160330:Subnet"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20160601:Subnet"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20160901:Subnet"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20161201:Subnet"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20170301:Subnet"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20170601:Subnet"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20170801:Subnet"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20170901:Subnet"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20171001:Subnet"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20171101:Subnet"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20180101:Subnet"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20180201:Subnet"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20180401:Subnet"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20180601:Subnet"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20180701:Subnet"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20180801:Subnet"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20181001:Subnet"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20181101:Subnet"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20181201:Subnet"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20190201:Subnet"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20190401:Subnet"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20190601:Subnet"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20190801:Subnet"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20190901:Subnet"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20191101:Subnet"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20191201:Subnet"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20200301:Subnet"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20200401:Subnet"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20200501:Subnet"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20200601:Subnet"},
                    new Pulumi.Alias { Type = "azure-nextgen:network/v20200701:Subnet"},
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
        /// The address prefix for the subnet.
        /// </summary>
        [Input("addressPrefix")]
        public Input<string>? AddressPrefix { get; set; }

        [Input("addressPrefixes")]
        private InputList<string>? _addressPrefixes;

        /// <summary>
        /// List of address prefixes for the subnet.
        /// </summary>
        public InputList<string> AddressPrefixes
        {
            get => _addressPrefixes ?? (_addressPrefixes = new InputList<string>());
            set => _addressPrefixes = value;
        }

        [Input("delegations")]
        private InputList<Inputs.DelegationArgs>? _delegations;

        /// <summary>
        /// An array of references to the delegations on the subnet.
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
        /// The name of the resource that is unique within a resource group. This name can be used to access the resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Nat gateway associated with this subnet.
        /// </summary>
        [Input("natGateway")]
        public Input<Inputs.SubResourceArgs>? NatGateway { get; set; }

        /// <summary>
        /// The reference of the NetworkSecurityGroup resource.
        /// </summary>
        [Input("networkSecurityGroup")]
        public Input<Inputs.NetworkSecurityGroupArgs>? NetworkSecurityGroup { get; set; }

        /// <summary>
        /// Enable or Disable apply network policies on private end point in the subnet.
        /// </summary>
        [Input("privateEndpointNetworkPolicies")]
        public Input<string>? PrivateEndpointNetworkPolicies { get; set; }

        /// <summary>
        /// Enable or Disable apply network policies on private link service in the subnet.
        /// </summary>
        [Input("privateLinkServiceNetworkPolicies")]
        public Input<string>? PrivateLinkServiceNetworkPolicies { get; set; }

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        [Input("resourceNavigationLinks")]
        private InputList<Inputs.ResourceNavigationLinkArgs>? _resourceNavigationLinks;

        /// <summary>
        /// An array of references to the external resources using subnet.
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
        /// An array of references to services injecting into this subnet.
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
