// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.SignalRService.V20200701Preview.Outputs
{

    [OutputType]
    public sealed class PrivateLinkServiceConnectionStateResponseResult
    {
        /// <summary>
        /// A message indicating if changes on the service provider require any updates on the consumer.
        /// </summary>
        public readonly string? ActionsRequired;
        /// <summary>
        /// The reason for approval/rejection of the connection.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Indicates whether the connection has been Approved/Rejected/Removed by the owner of the service.
        /// </summary>
        public readonly string? Status;

        [OutputConstructor]
        private PrivateLinkServiceConnectionStateResponseResult(
            string? actionsRequired,

            string? description,

            string? status)
        {
            ActionsRequired = actionsRequired;
            Description = description;
            Status = status;
        }
    }
}
