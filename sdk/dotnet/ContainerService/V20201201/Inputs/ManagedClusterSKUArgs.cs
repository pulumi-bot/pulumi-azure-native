// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.ContainerService.V20201201.Inputs
{

    public sealed class ManagedClusterSKUArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of a managed cluster SKU.
        /// </summary>
        [Input("name")]
        public InputUnion<string, Pulumi.AzureNextGen.ContainerService.V20201201.ManagedClusterSKUName>? Name { get; set; }

        /// <summary>
        /// Tier of a managed cluster SKU.
        /// </summary>
        [Input("tier")]
        public InputUnion<string, Pulumi.AzureNextGen.ContainerService.V20201201.ManagedClusterSKUTier>? Tier { get; set; }

        public ManagedClusterSKUArgs()
        {
        }
    }
}
