// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ServiceFabric.V20191101Preview.Inputs
{

    /// <summary>
    /// Describes how the service is partitioned.
    /// </summary>
    public sealed class PartitionSchemeDescriptionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies how the service is partitioned.
        /// </summary>
        [Input("partitionScheme", required: true)]
        public Input<string> PartitionScheme { get; set; } = null!;

        public PartitionSchemeDescriptionArgs()
        {
        }
    }
}
