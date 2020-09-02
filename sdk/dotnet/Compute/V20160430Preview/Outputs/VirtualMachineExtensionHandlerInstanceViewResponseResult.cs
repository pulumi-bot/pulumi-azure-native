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
    public sealed class VirtualMachineExtensionHandlerInstanceViewResponseResult
    {
        /// <summary>
        /// The extension handler status.
        /// </summary>
        public readonly Outputs.InstanceViewStatusResponseResult? Status;
        /// <summary>
        /// Specifies the type of the extension; an example is "CustomScriptExtension".
        /// </summary>
        public readonly string? Type;
        /// <summary>
        /// Specifies the version of the script handler.
        /// </summary>
        public readonly string? TypeHandlerVersion;

        [OutputConstructor]
        private VirtualMachineExtensionHandlerInstanceViewResponseResult(
            Outputs.InstanceViewStatusResponseResult? status,

            string? type,

            string? typeHandlerVersion)
        {
            Status = status;
            Type = type;
            TypeHandlerVersion = typeHandlerVersion;
        }
    }
}
