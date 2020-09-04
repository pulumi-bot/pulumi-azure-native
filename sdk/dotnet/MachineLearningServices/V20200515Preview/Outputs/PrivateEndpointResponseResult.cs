// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.MachineLearningServices.V20200515Preview.Outputs
{

    [OutputType]
    public sealed class PrivateEndpointResponseResult
    {
        /// <summary>
        /// The ARM identifier for Private Endpoint
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private PrivateEndpointResponseResult(string id)
        {
            Id = id;
        }
    }
}
