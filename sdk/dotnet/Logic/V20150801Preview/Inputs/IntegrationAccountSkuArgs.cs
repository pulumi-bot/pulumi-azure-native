// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Logic.V20150801Preview.Inputs
{

    public sealed class IntegrationAccountSkuArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The sku name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public IntegrationAccountSkuArgs()
        {
        }
    }
}
