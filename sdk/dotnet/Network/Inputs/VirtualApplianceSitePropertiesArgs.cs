// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.Inputs
{

    /// <summary>
    /// Properties of the rule group.
    /// </summary>
    public sealed class VirtualApplianceSitePropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Address Prefix.
        /// </summary>
        [Input("addressPrefix")]
        public Input<string>? AddressPrefix { get; set; }

        /// <summary>
        /// Office 365 Policy.
        /// </summary>
        [Input("o365Policy")]
        public Input<Inputs.Office365PolicyPropertiesArgs>? O365Policy { get; set; }

        public VirtualApplianceSitePropertiesArgs()
        {
        }
    }
}