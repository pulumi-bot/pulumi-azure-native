// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.ServiceBus.V20180101Preview.Inputs
{

    /// <summary>
    /// Properties supplied for Subnet
    /// </summary>
    public sealed class SubnetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Resource ID of Virtual Network Subnet
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        public SubnetArgs()
        {
        }
    }
}
