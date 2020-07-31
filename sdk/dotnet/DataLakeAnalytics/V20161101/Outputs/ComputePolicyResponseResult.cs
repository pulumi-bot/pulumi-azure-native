// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DataLakeAnalytics.V20161101.Outputs
{

    [OutputType]
    public sealed class ComputePolicyResponseResult
    {
        /// <summary>
        /// The resource identifier.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The compute policy properties.
        /// </summary>
        public readonly Outputs.ComputePolicyPropertiesResponseResult Properties;
        /// <summary>
        /// The resource type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private ComputePolicyResponseResult(
            string id,

            string name,

            Outputs.ComputePolicyPropertiesResponseResult properties,

            string type)
        {
            Id = id;
            Name = name;
            Properties = properties;
            Type = type;
        }
    }
}