// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.VirtualMachineImages.V20190201Preview.Outputs
{

    [OutputType]
    public sealed class ProvisioningErrorResponseResult
    {
        /// <summary>
        /// Verbose error message about the provisioning failure
        /// </summary>
        public readonly string? Message;
        /// <summary>
        /// Error code of the provisioning failure
        /// </summary>
        public readonly string? ProvisioningErrorCode;

        [OutputConstructor]
        private ProvisioningErrorResponseResult(
            string? message,

            string? provisioningErrorCode)
        {
            Message = message;
            ProvisioningErrorCode = provisioningErrorCode;
        }
    }
}
