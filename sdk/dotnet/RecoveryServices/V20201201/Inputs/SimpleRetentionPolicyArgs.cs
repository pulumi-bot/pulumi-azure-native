// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.RecoveryServices.V20201201.Inputs
{

    /// <summary>
    /// Simple policy retention.
    /// </summary>
    public sealed class SimpleRetentionPolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Retention duration of the protection policy.
        /// </summary>
        [Input("retentionDuration")]
        public Input<Inputs.RetentionDurationArgs>? RetentionDuration { get; set; }

        /// <summary>
        /// This property will be used as the discriminator for deciding the specific types in the polymorphic chain of types.
        /// Expected value is 'SimpleRetentionPolicy'.
        /// </summary>
        [Input("retentionPolicyType", required: true)]
        public Input<string> RetentionPolicyType { get; set; } = null!;

        public SimpleRetentionPolicyArgs()
        {
        }
    }
}
