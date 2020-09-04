// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ServiceFabric.V20190301Preview.Outputs
{

    [OutputType]
    public sealed class PartitionSchemeDescriptionResponseResult
    {
        /// <summary>
        /// Specifies how the service is partitioned.
        /// </summary>
        public readonly string PartitionScheme;

        [OutputConstructor]
        private PartitionSchemeDescriptionResponseResult(string partitionScheme)
        {
            PartitionScheme = partitionScheme;
        }
    }
}
