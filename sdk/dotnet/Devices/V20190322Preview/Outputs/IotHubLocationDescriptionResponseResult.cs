// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Devices.V20190322Preview.Outputs
{

    [OutputType]
    public sealed class IotHubLocationDescriptionResponseResult
    {
        /// <summary>
        /// Azure Geo Regions
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// Specific Role assigned to this location
        /// </summary>
        public readonly string? Role;

        [OutputConstructor]
        private IotHubLocationDescriptionResponseResult(
            string? location,

            string? role)
        {
            Location = location;
            Role = role;
        }
    }
}
