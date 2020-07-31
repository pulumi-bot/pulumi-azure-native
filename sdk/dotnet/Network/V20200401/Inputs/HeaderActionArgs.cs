// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20200401.Inputs
{

    /// <summary>
    /// An action that can manipulate an http header.
    /// </summary>
    public sealed class HeaderActionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Which type of manipulation to apply to the header.
        /// </summary>
        [Input("headerActionType", required: true)]
        public Input<string> HeaderActionType { get; set; } = null!;

        /// <summary>
        /// The name of the header this action will apply to.
        /// </summary>
        [Input("headerName", required: true)]
        public Input<string> HeaderName { get; set; } = null!;

        /// <summary>
        /// The value to update the given header name with. This value is not used if the actionType is Delete.
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public HeaderActionArgs()
        {
        }
    }
}