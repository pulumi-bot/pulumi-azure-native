// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ApiManagement.Outputs
{

    [OutputType]
    public sealed class GatewayContractPropertiesResponse
    {
        /// <summary>
        /// Gateway description
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Gateway location.
        /// </summary>
        public readonly Outputs.ResourceLocationDataContractResponse? LocationData;

        [OutputConstructor]
        private GatewayContractPropertiesResponse(
            string? description,

            Outputs.ResourceLocationDataContractResponse? locationData)
        {
            Description = description;
            LocationData = locationData;
        }
    }
}