// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.DesktopVirtualization.V20201019Preview.Inputs
{

    /// <summary>
    /// Represents a RegistrationInfo definition.
    /// </summary>
    public sealed class RegistrationInfoArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Expiration time of registration token.
        /// </summary>
        [Input("expirationTime")]
        public Input<string>? ExpirationTime { get; set; }

        /// <summary>
        /// The type of resetting the token.
        /// </summary>
        [Input("registrationTokenOperation")]
        public InputUnion<string, Pulumi.AzureNextGen.DesktopVirtualization.V20201019Preview.RegistrationTokenOperation>? RegistrationTokenOperation { get; set; }

        /// <summary>
        /// The registration token base64 encoded string.
        /// </summary>
        [Input("token")]
        public Input<string>? Token { get; set; }

        public RegistrationInfoArgs()
        {
        }
    }
}
