// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Compute.Outputs
{

    [OutputType]
    public sealed class VirtualMachineScaleSetVMProtectionPolicyResponse
    {
        /// <summary>
        /// Indicates that the virtual machine scale set VM shouldn't be considered for deletion during a scale-in operation.
        /// </summary>
        public readonly bool? ProtectFromScaleIn;
        /// <summary>
        /// Indicates that model updates or actions (including scale-in) initiated on the virtual machine scale set should not be applied to the virtual machine scale set VM.
        /// </summary>
        public readonly bool? ProtectFromScaleSetActions;

        [OutputConstructor]
        private VirtualMachineScaleSetVMProtectionPolicyResponse(
            bool? protectFromScaleIn,

            bool? protectFromScaleSetActions)
        {
            ProtectFromScaleIn = protectFromScaleIn;
            ProtectFromScaleSetActions = protectFromScaleSetActions;
        }
    }
}