// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Insights.V20180601Preview.Outputs
{

    [OutputType]
    public sealed class EtwEventConfigurationResponseResult
    {
        public readonly string? Filter;
        public readonly int Id;
        public readonly string Name;

        [OutputConstructor]
        private EtwEventConfigurationResponseResult(
            string? filter,

            int id,

            string name)
        {
            Filter = filter;
            Id = id;
            Name = name;
        }
    }
}
