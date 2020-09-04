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
    public sealed class FallbackRoutePropertiesResponseResult
    {
        /// <summary>
        /// The condition which is evaluated in order to apply the fallback route. If the condition is not provided it will evaluate to true by default. For grammar, See: https://docs.microsoft.com/azure/iot-hub/iot-hub-devguide-query-language
        /// </summary>
        public readonly string? Condition;
        /// <summary>
        /// The list of endpoints to which the messages that satisfy the condition are routed to. Currently only 1 endpoint is allowed.
        /// </summary>
        public readonly ImmutableArray<string> EndpointNames;
        /// <summary>
        /// Used to specify whether the fallback route is enabled.
        /// </summary>
        public readonly bool IsEnabled;
        /// <summary>
        /// The name of the route. The name can only include alphanumeric characters, periods, underscores, hyphens, has a maximum length of 64 characters, and must be unique.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The source to which the routing rule is to be applied to. For example, DeviceMessages
        /// </summary>
        public readonly string Source;

        [OutputConstructor]
        private FallbackRoutePropertiesResponseResult(
            string? condition,

            ImmutableArray<string> endpointNames,

            bool isEnabled,

            string? name,

            string source)
        {
            Condition = condition;
            EndpointNames = endpointNames;
            IsEnabled = isEnabled;
            Name = name;
            Source = source;
        }
    }
}
