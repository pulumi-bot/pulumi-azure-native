// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20181001.Inputs
{

    /// <summary>
    /// Network Security Group resource.
    /// </summary>
    public sealed class NetworkSecurityGroupPropertiesFormatResponseArgs : Pulumi.InvokeArgs
    {
        [Input("defaultSecurityRules")]
        private List<Inputs.SecurityRuleResponseArgs>? _defaultSecurityRules;

        /// <summary>
        /// The default security rules of network security group.
        /// </summary>
        public List<Inputs.SecurityRuleResponseArgs> DefaultSecurityRules
        {
            get => _defaultSecurityRules ?? (_defaultSecurityRules = new List<Inputs.SecurityRuleResponseArgs>());
            set => _defaultSecurityRules = value;
        }

        [Input("networkInterfaces", required: true)]
        private List<Inputs.NetworkInterfaceResponseArgs>? _networkInterfaces;

        /// <summary>
        /// A collection of references to network interfaces.
        /// </summary>
        public List<Inputs.NetworkInterfaceResponseArgs> NetworkInterfaces
        {
            get => _networkInterfaces ?? (_networkInterfaces = new List<Inputs.NetworkInterfaceResponseArgs>());
            set => _networkInterfaces = value;
        }

        /// <summary>
        /// The provisioning state of the public IP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
        /// </summary>
        [Input("provisioningState")]
        public string? ProvisioningState { get; set; }

        /// <summary>
        /// The resource GUID property of the network security group resource.
        /// </summary>
        [Input("resourceGuid")]
        public string? ResourceGuid { get; set; }

        [Input("securityRules")]
        private List<Inputs.SecurityRuleResponseArgs>? _securityRules;

        /// <summary>
        /// A collection of security rules of the network security group.
        /// </summary>
        public List<Inputs.SecurityRuleResponseArgs> SecurityRules
        {
            get => _securityRules ?? (_securityRules = new List<Inputs.SecurityRuleResponseArgs>());
            set => _securityRules = value;
        }

        [Input("subnets", required: true)]
        private List<Inputs.SubnetResponseArgs>? _subnets;

        /// <summary>
        /// A collection of references to subnets.
        /// </summary>
        public List<Inputs.SubnetResponseArgs> Subnets
        {
            get => _subnets ?? (_subnets = new List<Inputs.SubnetResponseArgs>());
            set => _subnets = value;
        }

        public NetworkSecurityGroupPropertiesFormatResponseArgs()
        {
        }
    }
}
