// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Network.V20190701.Inputs
{

    /// <summary>
    /// Contains the DDoS protection settings of the public IP.
    /// </summary>
    public sealed class DdosSettingsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The DDoS custom policy associated with the public IP.
        /// </summary>
        [Input("ddosCustomPolicy")]
        public Input<Inputs.SubResourceArgs>? DdosCustomPolicy { get; set; }

        /// <summary>
        /// The DDoS protection policy customizability of the public IP. Only standard coverage will have the ability to be customized.
        /// </summary>
        [Input("protectionCoverage")]
        public InputUnion<string, Pulumi.AzureNextGen.Network.V20190701.DdosSettingsProtectionCoverage>? ProtectionCoverage { get; set; }

        public DdosSettingsArgs()
        {
        }
    }
}
