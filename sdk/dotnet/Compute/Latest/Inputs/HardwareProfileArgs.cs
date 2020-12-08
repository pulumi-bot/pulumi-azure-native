// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Compute.Latest.Inputs
{

    /// <summary>
    /// Specifies the hardware settings for the virtual machine.
    /// </summary>
    public sealed class HardwareProfileArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the size of the virtual machine. For more information about virtual machine sizes, see [Sizes for virtual machines](https://docs.microsoft.com/en-us/azure/virtual-machines/sizes). &lt;br&gt;&lt;br&gt; The available VM sizes depend on region and availability set. For a list of available sizes use these APIs:  &lt;br&gt;&lt;br&gt; [List all available virtual machine sizes in an availability set](https://docs.microsoft.com/rest/api/compute/availabilitysets/listavailablesizes) &lt;br&gt;&lt;br&gt; [List all available virtual machine sizes in a region]( https://docs.microsoft.com/en-us/rest/api/compute/resourceskus/list) &lt;br&gt;&lt;br&gt; [List all available virtual machine sizes for resizing](https://docs.microsoft.com/rest/api/compute/virtualmachines/listavailablesizes). &lt;br&gt;&lt;br&gt; This list of sizes is no longer updated and the **VirtualMachineSizeTypes** string constants will be removed from the subsequent REST API specification. Use [List all available virtual machine sizes in a region]( https://docs.microsoft.com/en-us/rest/api/compute/resourceskus/list) to get the latest sizes.
        /// </summary>
        [Input("vmSize")]
        public InputUnion<string, Pulumi.AzureNextGen.Compute.Latest.VirtualMachineSizeTypes>? VmSize { get; set; }

        public HardwareProfileArgs()
        {
        }
    }
}
