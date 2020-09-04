// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.SqlVirtualMachine.V20170301Preview.Inputs
{

    /// <summary>
    /// A load balancer configuration for an availability group listener.
    /// </summary>
    public sealed class LoadBalancerConfigurationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Resource id of the load balancer.
        /// </summary>
        [Input("loadBalancerResourceId")]
        public Input<string>? LoadBalancerResourceId { get; set; }

        /// <summary>
        /// Private IP address.
        /// </summary>
        [Input("privateIpAddress")]
        public Input<Inputs.PrivateIPAddressArgs>? PrivateIpAddress { get; set; }

        /// <summary>
        /// Probe port.
        /// </summary>
        [Input("probePort")]
        public Input<int>? ProbePort { get; set; }

        /// <summary>
        /// Resource id of the public IP.
        /// </summary>
        [Input("publicIpAddressResourceId")]
        public Input<string>? PublicIpAddressResourceId { get; set; }

        [Input("sqlVirtualMachineInstances")]
        private InputList<string>? _sqlVirtualMachineInstances;

        /// <summary>
        /// List of the SQL virtual machine instance resource id's that are enrolled into the availability group listener.
        /// </summary>
        public InputList<string> SqlVirtualMachineInstances
        {
            get => _sqlVirtualMachineInstances ?? (_sqlVirtualMachineInstances = new InputList<string>());
            set => _sqlVirtualMachineInstances = value;
        }

        public LoadBalancerConfigurationArgs()
        {
        }
    }
}
