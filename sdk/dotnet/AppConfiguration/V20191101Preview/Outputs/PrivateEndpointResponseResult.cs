// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.AppConfiguration.V20191101Preview.Outputs
{

    [OutputType]
    public sealed class PrivateEndpointResponseResult
    {
        /// <summary>
        /// The resource Id for private endpoint
        /// </summary>
        public readonly string? Id;

        [OutputConstructor]
        private PrivateEndpointResponseResult(string? id)
        {
            Id = id;
        }
    }
}
