// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.RecoveryServices.V20180710.Inputs
{

    /// <summary>
    /// Recovery plan creation properties.
    /// </summary>
    public sealed class CreateRecoveryPlanInputPropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The failover deployment model.
        /// </summary>
        [Input("failoverDeploymentModel")]
        public InputUnion<string, Pulumi.AzureNextGen.RecoveryServices.V20180710.FailoverDeploymentModel>? FailoverDeploymentModel { get; set; }

        [Input("groups", required: true)]
        private InputList<Inputs.RecoveryPlanGroupArgs>? _groups;

        /// <summary>
        /// The recovery plan groups.
        /// </summary>
        public InputList<Inputs.RecoveryPlanGroupArgs> Groups
        {
            get => _groups ?? (_groups = new InputList<Inputs.RecoveryPlanGroupArgs>());
            set => _groups = value;
        }

        /// <summary>
        /// The primary fabric Id.
        /// </summary>
        [Input("primaryFabricId", required: true)]
        public Input<string> PrimaryFabricId { get; set; } = null!;

        [Input("providerSpecificInput")]
        private InputList<Inputs.RecoveryPlanA2AInputArgs>? _providerSpecificInput;

        /// <summary>
        /// The provider specific input.
        /// </summary>
        public InputList<Inputs.RecoveryPlanA2AInputArgs> ProviderSpecificInput
        {
            get => _providerSpecificInput ?? (_providerSpecificInput = new InputList<Inputs.RecoveryPlanA2AInputArgs>());
            set => _providerSpecificInput = value;
        }

        /// <summary>
        /// The recovery fabric Id.
        /// </summary>
        [Input("recoveryFabricId", required: true)]
        public Input<string> RecoveryFabricId { get; set; } = null!;

        public CreateRecoveryPlanInputPropertiesArgs()
        {
        }
    }
}
