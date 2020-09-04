// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20150504Preview.Inputs
{

    /// <summary>
    /// An A record.
    /// </summary>
    public sealed class ARecordArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Gets or sets the IPv4 address of this A record in string notation.
        /// </summary>
        [Input("ipv4Address")]
        public Input<string>? Ipv4Address { get; set; }

        public ARecordArgs()
        {
        }
    }
}
