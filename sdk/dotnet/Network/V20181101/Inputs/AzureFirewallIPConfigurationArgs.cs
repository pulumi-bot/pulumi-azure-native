// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20181101.Inputs
{

    /// <summary>
    /// IP configuration of an Azure Firewall.
    /// </summary>
    public sealed class AzureFirewallIPConfigurationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Resource ID.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Name of the resource that is unique within a resource group. This name can be used to access the resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The provisioning state of the resource.
        /// </summary>
        [Input("provisioningState")]
        public Input<string>? ProvisioningState { get; set; }

        /// <summary>
        /// Reference of the PublicIP resource. This field is a mandatory input if subnet is not null.
        /// </summary>
        [Input("publicIPAddress")]
        public Input<Inputs.SubResourceArgs>? PublicIPAddress { get; set; }

        /// <summary>
        /// Reference of the subnet resource. This resource must be named 'AzureFirewallSubnet'.
        /// </summary>
        [Input("subnet")]
        public Input<Inputs.SubResourceArgs>? Subnet { get; set; }

        public AzureFirewallIPConfigurationArgs()
        {
        }
    }
}
