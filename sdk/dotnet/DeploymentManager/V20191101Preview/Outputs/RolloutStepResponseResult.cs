// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.DeploymentManager.V20191101Preview.Outputs
{

    [OutputType]
    public sealed class RolloutStepResponseResult
    {
        /// <summary>
        /// Supplementary informative messages during rollout.
        /// </summary>
        public readonly ImmutableArray<Outputs.MessageResponseResult> Messages;
        /// <summary>
        /// Name of the step.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Detailed information of specific action execution.
        /// </summary>
        public readonly Outputs.StepOperationInfoResponseResult OperationInfo;
        /// <summary>
        /// Set of resource operations that were performed, if any, on an Azure resource.
        /// </summary>
        public readonly ImmutableArray<Outputs.ResourceOperationResponseResult> ResourceOperations;
        /// <summary>
        /// Current state of the step.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// The step group the current step is part of.
        /// </summary>
        public readonly string? StepGroup;

        [OutputConstructor]
        private RolloutStepResponseResult(
            ImmutableArray<Outputs.MessageResponseResult> messages,

            string name,

            Outputs.StepOperationInfoResponseResult operationInfo,

            ImmutableArray<Outputs.ResourceOperationResponseResult> resourceOperations,

            string status,

            string? stepGroup)
        {
            Messages = messages;
            Name = name;
            OperationInfo = operationInfo;
            ResourceOperations = resourceOperations;
            Status = status;
            StepGroup = stepGroup;
        }
    }
}
