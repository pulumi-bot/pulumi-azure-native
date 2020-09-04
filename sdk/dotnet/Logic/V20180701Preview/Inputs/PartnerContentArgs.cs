// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Logic.V20180701Preview.Inputs
{

    /// <summary>
    /// The integration account partner content.
    /// </summary>
    public sealed class PartnerContentArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The B2B partner content.
        /// </summary>
        [Input("b2b")]
        public Input<Inputs.B2BPartnerContentArgs>? B2b { get; set; }

        public PartnerContentArgs()
        {
        }
    }
}
