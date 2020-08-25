// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Aad.V20200101.Outputs
{

    [OutputType]
    public sealed class HealthMonitorResponseResult
    {
        /// <summary>
        /// Health Monitor Details
        /// </summary>
        public readonly string Details;
        /// <summary>
        /// Health Monitor Id
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Health Monitor Name
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private HealthMonitorResponseResult(
            string details,

            string id,

            string name)
        {
            Details = details;
            Id = id;
            Name = name;
        }
    }
}
