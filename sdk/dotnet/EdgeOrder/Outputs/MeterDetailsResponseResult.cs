// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.EdgeOrder.Outputs
{

    [OutputType]
    public sealed class MeterDetailsResponseResult
    {
        /// <summary>
        /// Billing model to represent billing cycle, i.e. Monthly, biweekly, daily, hourly etc.
        /// </summary>
        public readonly Outputs.BillingModelResponseResult BillingModel;
        /// <summary>
        /// MeterId/ Billing Guid against which the product system will be charged
        /// </summary>
        public readonly string MeterId;
        /// <summary>
        /// Category of the billing meter.
        /// </summary>
        public readonly string MeterType;

        [OutputConstructor]
        private MeterDetailsResponseResult(
            Outputs.BillingModelResponseResult billingModel,

            string meterId,

            string meterType)
        {
            BillingModel = billingModel;
            MeterId = meterId;
            MeterType = meterType;
        }
    }
}
