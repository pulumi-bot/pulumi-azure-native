// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Compute.V20160430Preview.Outputs
{

    [OutputType]
    public sealed class VirtualMachineExtensionInstanceViewResponseResult
    {
        /// <summary>
        /// The virtual machine extension name.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The resource status information.
        /// </summary>
        public readonly ImmutableArray<Outputs.InstanceViewStatusResponseResult> Statuses;
        /// <summary>
        /// The resource status information.
        /// </summary>
        public readonly ImmutableArray<Outputs.InstanceViewStatusResponseResult> Substatuses;
        /// <summary>
        /// Specifies the type of the extension; an example is "CustomScriptExtension".
        /// </summary>
        public readonly string? Type;
        /// <summary>
        /// Specifies the version of the script handler.
        /// </summary>
        public readonly string? TypeHandlerVersion;

        [OutputConstructor]
        private VirtualMachineExtensionInstanceViewResponseResult(
            string? name,

            ImmutableArray<Outputs.InstanceViewStatusResponseResult> statuses,

            ImmutableArray<Outputs.InstanceViewStatusResponseResult> substatuses,

            string? type,

            string? typeHandlerVersion)
        {
            Name = name;
            Statuses = statuses;
            Substatuses = substatuses;
            Type = type;
            TypeHandlerVersion = typeHandlerVersion;
        }
    }
}
