// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Network.V20200801.Inputs
{

    /// <summary>
    /// Reference to another subresource.
    /// </summary>
    public sealed class SubResourceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Resource ID.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        public SubResourceArgs()
        {
        }
    }
}
