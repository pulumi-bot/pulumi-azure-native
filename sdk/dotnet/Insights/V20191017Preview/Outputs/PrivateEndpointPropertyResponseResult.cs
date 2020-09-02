// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Insights.V20191017Preview.Outputs
{

    [OutputType]
    public sealed class PrivateEndpointPropertyResponseResult
    {
        /// <summary>
        /// Resource id of the private endpoint.
        /// </summary>
        public readonly string? Id;

        [OutputConstructor]
        private PrivateEndpointPropertyResponseResult(string? id)
        {
            Id = id;
        }
    }
}
