// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20150504Preview.Outputs
{

    [OutputType]
    public sealed class NsRecordResponseResult
    {
        /// <summary>
        /// Gets or sets the name server name for this record, without a terminating dot.
        /// </summary>
        public readonly string? Nsdname;

        [OutputConstructor]
        private NsRecordResponseResult(string? nsdname)
        {
            Nsdname = nsdname;
        }
    }
}
