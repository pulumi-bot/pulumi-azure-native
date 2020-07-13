// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.NotificationHubs.Outputs
{

    [OutputType]
    public sealed class ApnsCredentialResponse
    {
        /// <summary>
        /// Properties of NotificationHub ApnsCredential.
        /// </summary>
        public readonly Outputs.ApnsCredentialPropertiesResponse? Properties;

        [OutputConstructor]
        private ApnsCredentialResponse(Outputs.ApnsCredentialPropertiesResponse? properties)
        {
            Properties = properties;
        }
    }
}