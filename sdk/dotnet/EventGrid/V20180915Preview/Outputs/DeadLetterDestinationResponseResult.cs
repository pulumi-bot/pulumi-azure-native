// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.EventGrid.V20180915Preview.Outputs
{

    [OutputType]
    public sealed class DeadLetterDestinationResponseResult
    {
        /// <summary>
        /// Type of the endpoint for the dead letter destination
        /// </summary>
        public readonly string EndpointType;

        [OutputConstructor]
        private DeadLetterDestinationResponseResult(string endpointType)
        {
            EndpointType = endpointType;
        }
    }
}
