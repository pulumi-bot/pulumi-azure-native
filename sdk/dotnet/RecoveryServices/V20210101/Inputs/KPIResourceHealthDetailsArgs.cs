// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.RecoveryServices.V20210101.Inputs
{

    /// <summary>
    /// KPI Resource Health Details
    /// </summary>
    public sealed class KPIResourceHealthDetailsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Resource Health Status
        /// </summary>
        [Input("resourceHealthStatus")]
        public InputUnion<string, Pulumi.AzureNextGen.RecoveryServices.V20210101.ResourceHealthStatus>? ResourceHealthStatus { get; set; }

        public KPIResourceHealthDetailsArgs()
        {
        }
    }
}
