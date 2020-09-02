// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Compute.V20180630Preview.Inputs
{

    /// <summary>
    /// Guest configuration assignment properties.
    /// </summary>
    public sealed class GuestConfigurationAssignmentPropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The source which initiated the guest configuration assignment. Ex: Azure Policy
        /// </summary>
        [Input("context")]
        public Input<string>? Context { get; set; }

        /// <summary>
        /// The guest configuration to assign.
        /// </summary>
        [Input("guestConfiguration")]
        public Input<Inputs.GuestConfigurationNavigationArgs>? GuestConfiguration { get; set; }

        public GuestConfigurationAssignmentPropertiesArgs()
        {
        }
    }
}
