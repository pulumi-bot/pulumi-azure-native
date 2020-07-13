// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.PolicyInsights.Outputs
{

    [OutputType]
    public sealed class RemediationFiltersResponse
    {
        /// <summary>
        /// The resource locations that will be remediated.
        /// </summary>
        public readonly ImmutableArray<string> Locations;

        [OutputConstructor]
        private RemediationFiltersResponse(ImmutableArray<string> locations)
        {
            Locations = locations;
        }
    }
}