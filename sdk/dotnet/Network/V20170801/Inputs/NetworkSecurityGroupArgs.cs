// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20170801.Inputs
{

    /// <summary>
    /// NetworkSecurityGroup resource.
    /// </summary>
    public sealed class NetworkSecurityGroupArgs : Pulumi.ResourceArgs
    {
        [Input("defaultSecurityRules")]
        private InputList<Inputs.SecurityRuleArgs>? _defaultSecurityRules;

        /// <summary>
        /// The default security rules of network security group.
        /// </summary>
        public InputList<Inputs.SecurityRuleArgs> DefaultSecurityRules
        {
            get => _defaultSecurityRules ?? (_defaultSecurityRules = new InputList<Inputs.SecurityRuleArgs>());
            set => _defaultSecurityRules = value;
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
        /// Resource location.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The provisioning state of the public IP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
        /// </summary>
        [Input("provisioningState")]
        public Input<string>? ProvisioningState { get; set; }

        /// <summary>
        /// The resource GUID property of the network security group resource.
        /// </summary>
        [Input("resourceGuid")]
        public Input<string>? ResourceGuid { get; set; }

        [Input("securityRules")]
        private InputList<Inputs.SecurityRuleArgs>? _securityRules;

        /// <summary>
        /// A collection of security rules of the network security group.
        /// </summary>
        public InputList<Inputs.SecurityRuleArgs> SecurityRules
        {
            get => _securityRules ?? (_securityRules = new InputList<Inputs.SecurityRuleArgs>());
            set => _securityRules = value;
        }

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

        public NetworkSecurityGroupArgs()
        {
        }
    }
}
