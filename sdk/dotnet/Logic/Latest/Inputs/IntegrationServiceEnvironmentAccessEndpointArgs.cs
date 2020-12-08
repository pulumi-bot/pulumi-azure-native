// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Logic.Latest.Inputs
{

    /// <summary>
    /// The integration service environment access endpoint.
    /// </summary>
    public sealed class IntegrationServiceEnvironmentAccessEndpointArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The access endpoint type.
        /// </summary>
        [Input("type")]
        public InputUnion<string, Pulumi.AzureNextGen.Logic.Latest.IntegrationServiceEnvironmentAccessEndpointType>? Type { get; set; }

        public IntegrationServiceEnvironmentAccessEndpointArgs()
        {
        }
    }
}
