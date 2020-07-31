// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20191201.Inputs
{

    /// <summary>
    /// Properties specific to ExpressRouteLink resources.
    /// </summary>
    public sealed class ExpressRouteLinkPropertiesFormatArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Administrative state of the physical port.
        /// </summary>
        [Input("adminState")]
        public Input<string>? AdminState { get; set; }

        /// <summary>
        /// MacSec configuration.
        /// </summary>
        [Input("macSecConfig")]
        public Input<Inputs.ExpressRouteLinkMacSecConfigArgs>? MacSecConfig { get; set; }

        public ExpressRouteLinkPropertiesFormatArgs()
        {
        }
    }
}