// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DocumentDB.V20200601Preview.Outputs
{

    [OutputType]
    public sealed class ExcludedPathResponseResult
    {
        /// <summary>
        /// The path for which the indexing behavior applies to. Index paths typically start with root and end with wildcard (/path/*)
        /// </summary>
        public readonly string? Path;

        [OutputConstructor]
        private ExcludedPathResponseResult(string? path)
        {
            Path = path;
        }
    }
}
