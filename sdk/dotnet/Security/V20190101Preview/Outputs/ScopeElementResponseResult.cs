// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Security.V20190101Preview.Outputs
{

    [OutputType]
    public sealed class ScopeElementResponseResult
    {
        /// <summary>
        /// The alert entity type to suppress by.
        /// </summary>
        public readonly string? Field;

        [OutputConstructor]
        private ScopeElementResponseResult(string? field)
        {
            Field = field;
        }
    }
}
