// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DevTestLab.V20150521Preview.Inputs
{

    /// <summary>
    /// Property overrides on a subnet of a virtual network.
    /// </summary>
    public sealed class SubnetOverrideArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name given to the subnet within the lab.
        /// </summary>
        [Input("labSubnetName")]
        public Input<string>? LabSubnetName { get; set; }

        /// <summary>
        /// The resource identifier of the subnet.
        /// </summary>
        [Input("resourceId")]
        public Input<string>? ResourceId { get; set; }

        /// <summary>
        /// Indicates whether this subnet can be used during virtual machine creation.
        /// </summary>
        [Input("useInVmCreationPermission")]
        public Input<string>? UseInVmCreationPermission { get; set; }

        /// <summary>
        /// Indicates whether public IP addresses can be assigned to virtual machines on this subnet.
        /// </summary>
        [Input("usePublicIpAddressPermission")]
        public Input<string>? UsePublicIpAddressPermission { get; set; }

        public SubnetOverrideArgs()
        {
        }
    }
}
