// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Scheduler.V20140801Preview.Inputs
{

    public sealed class SkuArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Gets or set the SKU.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public SkuArgs()
        {
        }
    }
}
