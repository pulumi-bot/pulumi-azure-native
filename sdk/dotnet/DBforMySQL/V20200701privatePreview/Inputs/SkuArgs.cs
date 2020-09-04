// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DBForMySql.V20200701privatePreview.Inputs
{

    /// <summary>
    /// Billing information related properties of a server.
    /// </summary>
    public sealed class SkuArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the sku, e.g. Standard_D32s_v3.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The tier of the particular SKU, e.g. GeneralPurpose.
        /// </summary>
        [Input("tier")]
        public Input<string>? Tier { get; set; }

        public SkuArgs()
        {
        }
    }
}
