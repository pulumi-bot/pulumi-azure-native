// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.HanaOnAzure.Inputs
{

    /// <summary>
    /// Specifies the hardware settings for the HANA instance.
    /// </summary>
    public sealed class HardwareProfileArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the HANA instance SKU.
        /// </summary>
        [Input("hanaInstanceSize")]
        public InputUnion<string, Pulumi.AzureNative.HanaOnAzure.HanaInstanceSizeNamesEnum>? HanaInstanceSize { get; set; }

        /// <summary>
        /// Name of the hardware type (vendor and/or their product name)
        /// </summary>
        [Input("hardwareType")]
        public InputUnion<string, Pulumi.AzureNative.HanaOnAzure.HanaHardwareTypeNamesEnum>? HardwareType { get; set; }

        public HardwareProfileArgs()
        {
        }
    }
}
