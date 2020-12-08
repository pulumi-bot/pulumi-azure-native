// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Network.V20191101.Inputs
{

    /// <summary>
    /// ExpressRouteLink child resource definition.
    /// </summary>
    public sealed class ExpressRouteLinkArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Administrative state of the physical port.
        /// </summary>
        [Input("adminState")]
        public InputUnion<string, Pulumi.AzureNextGen.Network.V20191101.ExpressRouteLinkAdminState>? AdminState { get; set; }

        /// <summary>
        /// Resource ID.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// MacSec configuration.
        /// </summary>
        [Input("macSecConfig")]
        public Input<Inputs.ExpressRouteLinkMacSecConfigArgs>? MacSecConfig { get; set; }

        /// <summary>
        /// Name of child port resource that is unique among child port resources of the parent.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public ExpressRouteLinkArgs()
        {
        }
    }
}
