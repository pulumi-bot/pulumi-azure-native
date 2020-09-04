// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DataBoxEdge.V20200501Preview.Inputs
{

    /// <summary>
    /// The mapping between a particular user and the access type on the SMB share.
    /// </summary>
    public sealed class UserAccessRightArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Type of access to be allowed for the user.
        /// </summary>
        [Input("accessType", required: true)]
        public Input<string> AccessType { get; set; } = null!;

        /// <summary>
        /// User ID (already existing in the device).
        /// </summary>
        [Input("userId", required: true)]
        public Input<string> UserId { get; set; } = null!;

        public UserAccessRightArgs()
        {
        }
    }
}
