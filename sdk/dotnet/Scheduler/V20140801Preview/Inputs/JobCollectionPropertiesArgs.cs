// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Scheduler.V20140801Preview.Inputs
{

    public sealed class JobCollectionPropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Gets or sets the job collection quota.
        /// </summary>
        [Input("quota")]
        public Input<Inputs.JobCollectionQuotaArgs>? Quota { get; set; }

        /// <summary>
        /// Gets or sets the SKU.
        /// </summary>
        [Input("sku")]
        public Input<Inputs.SkuArgs>? Sku { get; set; }

        /// <summary>
        /// Gets or sets the state.
        /// </summary>
        [Input("state")]
        public Input<string>? State { get; set; }

        public JobCollectionPropertiesArgs()
        {
        }
    }
}
