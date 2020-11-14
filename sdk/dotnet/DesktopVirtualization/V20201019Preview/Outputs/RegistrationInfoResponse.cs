// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.DesktopVirtualization.V20201019Preview.Outputs
{

    [OutputType]
    public sealed class RegistrationInfoResponse
    {
        /// <summary>
        /// Expiration time of registration token.
        /// </summary>
        public readonly string? ExpirationTime;
        /// <summary>
        /// The type of resetting the token.
        /// </summary>
        public readonly string? RegistrationTokenOperation;
        /// <summary>
        /// The registration token base64 encoded string.
        /// </summary>
        public readonly string? Token;

        [OutputConstructor]
        private RegistrationInfoResponse(
            string? expirationTime,

            string? registrationTokenOperation,

            string? token)
        {
            ExpirationTime = expirationTime;
            RegistrationTokenOperation = registrationTokenOperation;
            Token = token;
        }
    }
}
