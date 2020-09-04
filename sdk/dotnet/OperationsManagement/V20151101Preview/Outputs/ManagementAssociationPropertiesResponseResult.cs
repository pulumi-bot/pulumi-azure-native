// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.OperationsManagement.V20151101Preview.Outputs
{

    [OutputType]
    public sealed class ManagementAssociationPropertiesResponseResult
    {
        /// <summary>
        /// The applicationId of the appliance for this association.
        /// </summary>
        public readonly string ApplicationId;

        [OutputConstructor]
        private ManagementAssociationPropertiesResponseResult(string applicationId)
        {
            ApplicationId = applicationId;
        }
    }
}
