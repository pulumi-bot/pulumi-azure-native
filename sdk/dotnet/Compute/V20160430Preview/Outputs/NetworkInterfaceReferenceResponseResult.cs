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
    public sealed class NetworkInterfaceReferenceResponseResult
    {
        /// <summary>
        /// Resource Id
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// Specifies the primary network interface in case the virtual machine has more than 1 network interface.
        /// </summary>
        public readonly bool? Primary;

        [OutputConstructor]
        private NetworkInterfaceReferenceResponseResult(
            string? id,

            bool? primary)
        {
            Id = id;
            Primary = primary;
        }
    }
}
