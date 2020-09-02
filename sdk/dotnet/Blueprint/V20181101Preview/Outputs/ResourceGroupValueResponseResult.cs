// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Blueprint.V20181101Preview.Outputs
{

    [OutputType]
    public sealed class ResourceGroupValueResponseResult
    {
        /// <summary>
        /// Location of the resource group.
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// Name of the resource group.
        /// </summary>
        public readonly string? Name;

        [OutputConstructor]
        private ResourceGroupValueResponseResult(
            string? location,

            string? name)
        {
            Location = location;
            Name = name;
        }
    }
}
