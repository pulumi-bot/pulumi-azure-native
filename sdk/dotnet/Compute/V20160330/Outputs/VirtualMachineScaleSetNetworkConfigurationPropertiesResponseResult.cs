// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Compute.V20160330.Outputs
{

    [OutputType]
    public sealed class VirtualMachineScaleSetNetworkConfigurationPropertiesResponseResult
    {
        /// <summary>
        /// The virtual machine scale set IP Configuration.
        /// </summary>
        public readonly ImmutableArray<Outputs.VirtualMachineScaleSetIPConfigurationResponseResult> IpConfigurations;
        /// <summary>
        /// Whether this is a primary NIC on a virtual machine.
        /// </summary>
        public readonly bool? Primary;

        [OutputConstructor]
        private VirtualMachineScaleSetNetworkConfigurationPropertiesResponseResult(
            ImmutableArray<Outputs.VirtualMachineScaleSetIPConfigurationResponseResult> ipConfigurations,

            bool? primary)
        {
            IpConfigurations = ipConfigurations;
            Primary = primary;
        }
    }
}