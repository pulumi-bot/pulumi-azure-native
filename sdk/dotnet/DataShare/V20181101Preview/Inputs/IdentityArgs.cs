// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DataShare.V20181101Preview.Inputs
{

    /// <summary>
    /// Identity of resource
    /// </summary>
    public sealed class IdentityArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Identity Type
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public IdentityArgs()
        {
        }
    }
}
