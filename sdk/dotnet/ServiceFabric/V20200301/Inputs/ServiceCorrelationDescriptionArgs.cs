// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.ServiceFabric.V20200301.Inputs
{

    /// <summary>
    /// Creates a particular correlation between services.
    /// </summary>
    public sealed class ServiceCorrelationDescriptionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ServiceCorrelationScheme which describes the relationship between this service and the service specified via ServiceName.
        /// </summary>
        [Input("scheme", required: true)]
        public InputUnion<string, Pulumi.AzureNextGen.ServiceFabric.V20200301.ServiceCorrelationScheme> Scheme { get; set; } = null!;

        /// <summary>
        /// The name of the service that the correlation relationship is established with.
        /// </summary>
        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        public ServiceCorrelationDescriptionArgs()
        {
        }
    }
}
