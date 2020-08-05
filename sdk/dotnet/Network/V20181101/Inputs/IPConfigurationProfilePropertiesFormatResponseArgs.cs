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
    /// IP configuration profile properties.
    /// </summary>
    public sealed class IPConfigurationProfilePropertiesFormatResponseArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The provisioning state of the resource.
        /// </summary>
        [Input("provisioningState", required: true)]
        public string ProvisioningState { get; set; } = null!;

        /// <summary>
        /// The reference of the subnet resource to create a container network interface ip configuration.
        /// </summary>
        [Input("subnet")]
        public Inputs.SubnetResponseArgs? Subnet { get; set; }

        public IPConfigurationProfilePropertiesFormatResponseArgs()
        {
        }
    }
}
