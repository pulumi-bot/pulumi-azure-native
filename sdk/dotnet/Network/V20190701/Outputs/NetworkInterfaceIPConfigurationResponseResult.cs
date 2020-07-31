// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20190701.Outputs
{

    [OutputType]
    public sealed class NetworkInterfaceIPConfigurationResponseResult
    {
        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        public readonly string? Etag;
        /// <summary>
        /// Resource ID.
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// The name of the resource that is unique within a resource group. This name can be used to access the resource.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Network interface IP configuration properties.
        /// </summary>
        public readonly Outputs.NetworkInterfaceIPConfigurationPropertiesFormatResponseResult? Properties;

        [OutputConstructor]
        private NetworkInterfaceIPConfigurationResponseResult(
            string? etag,

            string? id,

            string? name,

            Outputs.NetworkInterfaceIPConfigurationPropertiesFormatResponseResult? properties)
        {
            Etag = etag;
            Id = id;
            Name = name;
            Properties = properties;
        }
    }
}