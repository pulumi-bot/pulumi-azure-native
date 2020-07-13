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
    public sealed class VirtualMachineScaleSetExtensionProfileResponse
    {
        /// <summary>
        /// The virtual machine scale set child extension resources.
        /// </summary>
        public readonly ImmutableArray<Outputs.VirtualMachineScaleSetExtensionResponse> Extensions;
        /// <summary>
        /// Specifies the time alloted for all extensions to start. The time duration should be between 15 minutes and 120 minutes (inclusive) and should be specified in ISO 8601 format. The default value is 90 minutes (PT1H30M). &lt;br&gt;&lt;br&gt; Minimum api-version: 2020-06-01
        /// </summary>
        public readonly string? ExtensionsTimeBudget;

        [OutputConstructor]
        private VirtualMachineScaleSetExtensionProfileResponse(
            ImmutableArray<Outputs.VirtualMachineScaleSetExtensionResponse> extensions,

            string? extensionsTimeBudget)
        {
            Extensions = extensions;
            ExtensionsTimeBudget = extensionsTimeBudget;
        }
    }
}