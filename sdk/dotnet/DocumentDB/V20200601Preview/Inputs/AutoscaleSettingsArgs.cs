// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DocumentDB.V20200601Preview.Inputs
{

    public sealed class AutoscaleSettingsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Represents maximum throughput, the resource can scale up to.
        /// </summary>
        [Input("maxThroughput")]
        public Input<int>? MaxThroughput { get; set; }

        public AutoscaleSettingsArgs()
        {
        }
    }
}
