// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Automation.Latest.Outputs
{

    [OutputType]
    public sealed class TagSettingsPropertiesResponse
    {
        /// <summary>
        /// Filter VMs by Any or All specified tags.
        /// </summary>
        public readonly string? FilterOperator;
        /// <summary>
        /// Dictionary of tags with its list of values.
        /// </summary>
        public readonly ImmutableDictionary<string, ImmutableArray<string>>? Tags;

        [OutputConstructor]
        private TagSettingsPropertiesResponse(
            string? filterOperator,

            ImmutableDictionary<string, ImmutableArray<string>>? tags)
        {
            FilterOperator = filterOperator;
            Tags = tags;
        }
    }
}
