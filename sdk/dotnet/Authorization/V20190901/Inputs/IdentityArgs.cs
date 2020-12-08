// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Authorization.V20190901.Inputs
{

    /// <summary>
    /// Identity for the resource.
    /// </summary>
    public sealed class IdentityArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The identity type. This is the only required field when adding a system assigned identity to a resource.
        /// </summary>
        [Input("type")]
        public Input<Pulumi.AzureNextGen.Authorization.V20190901.ResourceIdentityType>? Type { get; set; }

        public IdentityArgs()
        {
        }
    }
}
