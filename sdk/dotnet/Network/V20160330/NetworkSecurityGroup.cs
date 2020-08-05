// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20160330
{
    /// <summary>
    /// NetworkSecurityGroup resource
    /// </summary>
    public partial class NetworkSecurityGroup : Pulumi.CustomResource
    {
        /// <summary>
        /// Gets a unique read-only string that changes whenever the resource is updated
        /// </summary>
        [Output("etag")]
        public Output<string?> Etag { get; private set; } = null!;

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
        /// Network Security Group resource
        /// </summary>
        [Output("properties")]
        public Output<Outputs.NetworkSecurityGroupPropertiesFormatResponseResult> Properties { get; private set; } = null!;

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
        /// Create a NetworkSecurityGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NetworkSecurityGroup(string name, NetworkSecurityGroupArgs args, CustomResourceOptions? options = null)
            : base("azurerm:network/v20160330:NetworkSecurityGroup", name, args ?? new NetworkSecurityGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NetworkSecurityGroup(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:network/v20160330:NetworkSecurityGroup", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing NetworkSecurityGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NetworkSecurityGroup Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new NetworkSecurityGroup(name, id, options);
        }
    }

    public sealed class NetworkSecurityGroupArgs : Pulumi.ResourceArgs
    {
        [Input("defaultSecurityRules")]
        private InputList<Inputs.SecurityRuleArgs>? _defaultSecurityRules;

        /// <summary>
        /// Gets or sets Default security rules of network security group
        /// </summary>
        public InputList<Inputs.SecurityRuleArgs> DefaultSecurityRules
        {
            get => _defaultSecurityRules ?? (_defaultSecurityRules = new InputList<Inputs.SecurityRuleArgs>());
            set => _defaultSecurityRules = value;
        }

        /// <summary>
        /// Gets a unique read-only string that changes whenever the resource is updated
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

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

        /// <summary>
        /// The name of the network security group.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("networkInterfaces")]
        private InputList<Inputs.NetworkInterfaceArgs>? _networkInterfaces;

        /// <summary>
        /// Gets collection of references to Network Interfaces
        /// </summary>
        public InputList<Inputs.NetworkInterfaceArgs> NetworkInterfaces
        {
            get => _networkInterfaces ?? (_networkInterfaces = new InputList<Inputs.NetworkInterfaceArgs>());
            set => _networkInterfaces = value;
        }

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
        /// Gets or sets resource GUID property of the network security group resource
        /// </summary>
        [Input("resourceGuid")]
        public Input<string>? ResourceGuid { get; set; }

        [Input("securityRules")]
        private InputList<Inputs.SecurityRuleArgs>? _securityRules;

        /// <summary>
        /// Gets or sets Security rules of network security group
        /// </summary>
        public InputList<Inputs.SecurityRuleArgs> SecurityRules
        {
            get => _securityRules ?? (_securityRules = new InputList<Inputs.SecurityRuleArgs>());
            set => _securityRules = value;
        }

        [Input("subnets")]
        private InputList<Inputs.SubnetArgs>? _subnets;

        /// <summary>
        /// Gets collection of references to subnets
        /// </summary>
        public InputList<Inputs.SubnetArgs> Subnets
        {
            get => _subnets ?? (_subnets = new InputList<Inputs.SubnetArgs>());
            set => _subnets = value;
        }

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

        public NetworkSecurityGroupArgs()
        {
        }
    }
}
