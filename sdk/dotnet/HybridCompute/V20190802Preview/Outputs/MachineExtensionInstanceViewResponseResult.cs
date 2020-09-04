// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.HybridCompute.V20190802Preview.Outputs
{

    [OutputType]
    public sealed class MachineExtensionInstanceViewResponseResult
    {
        /// <summary>
        /// The machine extension name.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Instance view status.
        /// </summary>
        public readonly Outputs.MachineExtensionInstanceViewResponseStatusResult? Status;
        /// <summary>
        /// Specifies the type of the extension; an example is "CustomScriptExtension".
        /// </summary>
        public readonly string? Type;
        /// <summary>
        /// Specifies the version of the script handler.
        /// </summary>
        public readonly string? TypeHandlerVersion;

        [OutputConstructor]
        private MachineExtensionInstanceViewResponseResult(
            string? name,

            Outputs.MachineExtensionInstanceViewResponseStatusResult? status,

            string? type,

            string? typeHandlerVersion)
        {
            Name = name;
            Status = status;
            Type = type;
            TypeHandlerVersion = typeHandlerVersion;
        }
    }
}
