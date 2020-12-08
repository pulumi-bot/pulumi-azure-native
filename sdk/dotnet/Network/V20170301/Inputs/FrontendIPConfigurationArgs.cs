// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Network.V20170301.Inputs
{

    /// <summary>
    /// Frontend IP address of the load balancer.
    /// </summary>
    public sealed class FrontendIPConfigurationArgs : Pulumi.ResourceArgs
    {
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
        /// The private IP address of the IP configuration.
        /// </summary>
        [Input("privateIPAddress")]
        public Input<string>? PrivateIPAddress { get; set; }

        /// <summary>
        /// The Private IP allocation method. Possible values are: 'Static' and 'Dynamic'.
        /// </summary>
        [Input("privateIPAllocationMethod")]
        public InputUnion<string, Pulumi.AzureNextGen.Network.V20170301.IPAllocationMethod>? PrivateIPAllocationMethod { get; set; }

        /// <summary>
        /// Gets the provisioning state of the public IP resource. Possible values are: 'Updating', 'Deleting', and 'Failed'.
        /// </summary>
        [Input("provisioningState")]
        public Input<string>? ProvisioningState { get; set; }

        /// <summary>
        /// The reference of the Public IP resource.
        /// </summary>
        [Input("publicIPAddress")]
        public Input<Inputs.PublicIPAddressArgs>? PublicIPAddress { get; set; }

        /// <summary>
        /// The reference of the subnet resource.
        /// </summary>
        [Input("subnet")]
        public Input<Inputs.SubnetArgs>? Subnet { get; set; }

        public FrontendIPConfigurationArgs()
        {
        }
    }
}
