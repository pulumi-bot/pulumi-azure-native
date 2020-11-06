// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Storage.V20200801Preview.Outputs
{

    [OutputType]
    public sealed class ChangeFeedResponse
    {
        /// <summary>
        /// Indicates whether change feed event logging is enabled for the Blob service.
        /// </summary>
        public readonly bool? Enabled;

        [OutputConstructor]
        private ChangeFeedResponse(bool? enabled)
        {
            Enabled = enabled;
        }
    }
}
