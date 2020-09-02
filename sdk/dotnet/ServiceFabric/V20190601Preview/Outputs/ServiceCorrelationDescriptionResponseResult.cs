// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ServiceFabric.V20190601Preview.Outputs
{

    [OutputType]
    public sealed class ServiceCorrelationDescriptionResponseResult
    {
        /// <summary>
        /// The ServiceCorrelationScheme which describes the relationship between this service and the service specified via ServiceName.
        /// </summary>
        public readonly string Scheme;
        /// <summary>
        /// The name of the service that the correlation relationship is established with.
        /// </summary>
        public readonly string ServiceName;

        [OutputConstructor]
        private ServiceCorrelationDescriptionResponseResult(
            string scheme,

            string serviceName)
        {
            Scheme = scheme;
            ServiceName = serviceName;
        }
    }
}
