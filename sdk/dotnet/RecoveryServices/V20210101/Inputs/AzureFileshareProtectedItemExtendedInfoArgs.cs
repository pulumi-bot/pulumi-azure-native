// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.RecoveryServices.V20210101.Inputs
{

    /// <summary>
    /// Additional information about Azure File Share backup item.
    /// </summary>
    public sealed class AzureFileshareProtectedItemExtendedInfoArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The oldest backup copy available for this item in the service.
        /// </summary>
        [Input("oldestRecoveryPoint")]
        public Input<string>? OldestRecoveryPoint { get; set; }

        /// <summary>
        /// Indicates consistency of policy object and policy applied to this backup item.
        /// </summary>
        [Input("policyState")]
        public Input<string>? PolicyState { get; set; }

        /// <summary>
        /// Number of available backup copies associated with this backup item.
        /// </summary>
        [Input("recoveryPointCount")]
        public Input<int>? RecoveryPointCount { get; set; }

        public AzureFileshareProtectedItemExtendedInfoArgs()
        {
        }
    }
}
