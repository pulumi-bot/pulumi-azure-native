// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.SoftwarePlan.V20190601Preview.Outputs
{

    [OutputType]
    public sealed class SkuResponseResult
    {
        /// <summary>
        /// Name of the SKU to be applied
        /// </summary>
        public readonly string? Name;

        [OutputConstructor]
        private SkuResponseResult(string? name)
        {
            Name = name;
        }
    }
}
