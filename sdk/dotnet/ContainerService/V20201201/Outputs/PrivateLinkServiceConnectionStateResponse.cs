// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.ContainerService.V20201201.Outputs
{

    [OutputType]
    public sealed class PrivateLinkServiceConnectionStateResponse
    {
        /// <summary>
        /// The private link service connection description.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// The private link service connection status.
        /// </summary>
        public readonly string? Status;

        [OutputConstructor]
        private PrivateLinkServiceConnectionStateResponse(
            string? description,

            string? status)
        {
            Description = description;
            Status = status;
        }
    }
}
