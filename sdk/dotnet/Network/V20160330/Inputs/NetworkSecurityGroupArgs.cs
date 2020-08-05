// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20160330.Inputs
{

    /// <summary>
    /// NetworkSecurityGroup resource
    /// </summary>
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
        private InputList<Inputs.SubnetDefinitionArgs>? _subnets;

        /// <summary>
        /// Gets collection of references to subnets
        /// </summary>
        public InputList<Inputs.SubnetDefinitionArgs> Subnets
        {
            get => _subnets ?? (_subnets = new InputList<Inputs.SubnetDefinitionArgs>());
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
