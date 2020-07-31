// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ContainerService.V20170701.Outputs
{

    [OutputType]
    public sealed class ContainerServiceVMDiagnosticsResponseResult
    {
        /// <summary>
        /// Whether the VM diagnostic agent is provisioned on the VM.
        /// </summary>
        public readonly bool Enabled;
        /// <summary>
        /// The URI of the storage account where diagnostics are stored.
        /// </summary>
        public readonly string StorageUri;

        [OutputConstructor]
        private ContainerServiceVMDiagnosticsResponseResult(
            bool enabled,

            string storageUri)
        {
            Enabled = enabled;
            StorageUri = storageUri;
        }
    }
}