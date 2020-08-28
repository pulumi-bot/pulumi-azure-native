// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Search.V20200801.Outputs
{

    [OutputType]
    public sealed class PrivateEndpointConnectionPropertiesResponsePrivateEndpointResult
    {
        /// <summary>
        /// The resource id of the private endpoint resource from Microsoft.Network provider.
        /// </summary>
        public readonly string? Id;

        [OutputConstructor]
        private PrivateEndpointConnectionPropertiesResponsePrivateEndpointResult(string? id)
        {
            Id = id;
        }
    }
}
