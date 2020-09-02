// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Compute.V20180630Preview.Outputs
{

    [OutputType]
    public sealed class GuestConfigurationAssignmentPropertiesResponseResult
    {
        /// <summary>
        /// Combined hash of the configuration package and parameters.
        /// </summary>
        public readonly string AssignmentHash;
        /// <summary>
        /// A value indicating compliance status of the virtual machine for the assigned guest configuration.
        /// </summary>
        public readonly string ComplianceStatus;
        /// <summary>
        /// The source which initiated the guest configuration assignment. Ex: Azure Policy
        /// </summary>
        public readonly string? Context;
        /// <summary>
        /// The guest configuration to assign.
        /// </summary>
        public readonly Outputs.GuestConfigurationNavigationResponseResult? GuestConfiguration;
        /// <summary>
        /// Date and time when last compliance status was checked.
        /// </summary>
        public readonly string LastComplianceStatusChecked;
        /// <summary>
        /// Id of the latest report for the guest configuration assignment. 
        /// </summary>
        public readonly string LatestReportId;
        /// <summary>
        /// The provisioning state, which only appears in the response.
        /// </summary>
        public readonly string ProvisioningState;

        [OutputConstructor]
        private GuestConfigurationAssignmentPropertiesResponseResult(
            string assignmentHash,

            string complianceStatus,

            string? context,

            Outputs.GuestConfigurationNavigationResponseResult? guestConfiguration,

            string lastComplianceStatusChecked,

            string latestReportId,

            string provisioningState)
        {
            AssignmentHash = assignmentHash;
            ComplianceStatus = complianceStatus;
            Context = context;
            GuestConfiguration = guestConfiguration;
            LastComplianceStatusChecked = lastComplianceStatusChecked;
            LatestReportId = latestReportId;
            ProvisioningState = provisioningState;
        }
    }
}
