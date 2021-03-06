// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Compute.Outputs
{

    [OutputType]
    public sealed class LoadBalancerFrontendIPConfigurationResponse
    {
        public readonly string? Name;
        /// <summary>
        /// Describes a cloud service IP Configuration
        /// </summary>
        public readonly Outputs.LoadBalancerFrontendIPConfigurationPropertiesResponse? Properties;

        [OutputConstructor]
        private LoadBalancerFrontendIPConfigurationResponse(
            string? name,

            Outputs.LoadBalancerFrontendIPConfigurationPropertiesResponse? properties)
        {
            Name = name;
            Properties = properties;
        }
    }
}
