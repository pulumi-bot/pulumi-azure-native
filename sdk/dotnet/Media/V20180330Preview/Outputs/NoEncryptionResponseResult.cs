// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Media.V20180330Preview.Outputs
{

    [OutputType]
    public sealed class NoEncryptionResponseResult
    {
        /// <summary>
        /// Representing supported protocols
        /// </summary>
        public readonly Outputs.EnabledProtocolsResponseResult? EnabledProtocols;

        [OutputConstructor]
        private NoEncryptionResponseResult(Outputs.EnabledProtocolsResponseResult? enabledProtocols)
        {
            EnabledProtocols = enabledProtocols;
        }
    }
}
