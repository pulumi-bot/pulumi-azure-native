// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.StorageCache.V20190801Preview.Outputs
{

    [OutputType]
    public sealed class UnknownTargetResponseResult
    {
        /// <summary>
        /// Dictionary of string-&gt;string pairs containing information about the StorageTarget.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? UnknownMap;

        [OutputConstructor]
        private UnknownTargetResponseResult(ImmutableDictionary<string, string>? unknownMap)
        {
            UnknownMap = unknownMap;
        }
    }
}
