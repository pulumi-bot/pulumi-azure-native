// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Media.V20180601Preview.Outputs
{

    [OutputType]
    public sealed class IPAccessControlResponseResult
    {
        /// <summary>
        /// The IP allow list.
        /// </summary>
        public readonly ImmutableArray<Outputs.IPRangeResponseResult> Allow;

        [OutputConstructor]
        private IPAccessControlResponseResult(ImmutableArray<Outputs.IPRangeResponseResult> allow)
        {
            Allow = allow;
        }
    }
}
