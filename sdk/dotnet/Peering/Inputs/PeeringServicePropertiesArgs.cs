// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Peering.Inputs
{

    /// <summary>
    /// The properties that define connectivity to the Peering Service.
    /// </summary>
    public sealed class PeeringServicePropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The PeeringServiceLocation of the Customer.
        /// </summary>
        [Input("peeringServiceLocation")]
        public Input<string>? PeeringServiceLocation { get; set; }

        /// <summary>
        /// The MAPS Provider Name.
        /// </summary>
        [Input("peeringServiceProvider")]
        public Input<string>? PeeringServiceProvider { get; set; }

        public PeeringServicePropertiesArgs()
        {
        }
    }
}