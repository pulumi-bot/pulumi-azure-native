// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.HealthcareApis.V20180820Preview.Inputs
{

    /// <summary>
    /// An access policy entry.
    /// </summary>
    public sealed class ServiceAccessPolicyEntryArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// An object ID that is allowed access to the FHIR service.
        /// </summary>
        [Input("objectId", required: true)]
        public Input<string> ObjectId { get; set; } = null!;

        public ServiceAccessPolicyEntryArgs()
        {
        }
    }
}
