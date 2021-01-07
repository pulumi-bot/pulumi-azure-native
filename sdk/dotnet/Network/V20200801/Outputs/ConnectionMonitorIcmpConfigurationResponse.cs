// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Network.V20200801.Outputs
{

    [OutputType]
    public sealed class ConnectionMonitorIcmpConfigurationResponse
    {
        /// <summary>
        /// Value indicating whether path evaluation with trace route should be disabled.
        /// </summary>
        public readonly bool? DisableTraceRoute;

        [OutputConstructor]
        private ConnectionMonitorIcmpConfigurationResponse(bool? disableTraceRoute)
        {
            DisableTraceRoute = disableTraceRoute;
        }
    }
}
