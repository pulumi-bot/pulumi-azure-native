// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Confluent.V20200301Preview.Outputs
{

    [OutputType]
    public sealed class OrganizationResourcePropertiesResponseOfferDetail
    {
        /// <summary>
        /// Offer Id
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// Offer Plan Id
        /// </summary>
        public readonly string? PlanId;
        /// <summary>
        /// Offer Plan Name
        /// </summary>
        public readonly string? PlanName;
        /// <summary>
        /// Publisher Id
        /// </summary>
        public readonly string? PublisherId;
        /// <summary>
        /// SaaS Offer Status
        /// </summary>
        public readonly string? Status;
        /// <summary>
        /// Offer Plan Term unit
        /// </summary>
        public readonly string? TermUnit;

        [OutputConstructor]
        private OrganizationResourcePropertiesResponseOfferDetail(
            string? id,

            string? planId,

            string? planName,

            string? publisherId,

            string? status,

            string? termUnit)
        {
            Id = id;
            PlanId = planId;
            PlanName = planName;
            PublisherId = publisherId;
            Status = status;
            TermUnit = termUnit;
        }
    }
}
