// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.EdgeOrder.V20201201Preview.Outputs
{

    [OutputType]
    public sealed class BillingModelResponseResult
    {
        /// <summary>
        /// String to represent the billing model
        /// </summary>
        public readonly string Model;

        [OutputConstructor]
        private BillingModelResponseResult(string model)
        {
            Model = model;
        }
    }
}
