// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.DataBoxEdge.Latest.Inputs
{

    /// <summary>
    /// Specifies the mapping between this particular user and the type of access he has on shares on this device.
    /// </summary>
    public sealed class ShareAccessRightArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Type of access to be allowed on the share for this user.
        /// </summary>
        [Input("accessType", required: true)]
        public InputUnion<string, Pulumi.AzureNextGen.DataBoxEdge.Latest.ShareAccessType> AccessType { get; set; } = null!;

        /// <summary>
        /// The share ID.
        /// </summary>
        [Input("shareId", required: true)]
        public Input<string> ShareId { get; set; } = null!;

        public ShareAccessRightArgs()
        {
        }
    }
}
