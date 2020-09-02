// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DataMigration.V20171115Preview.Inputs
{

    /// <summary>
    /// An Azure SKU instance
    /// </summary>
    public sealed class ServiceSkuArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The capacity of the SKU, if it supports scaling
        /// </summary>
        [Input("capacity")]
        public Input<int>? Capacity { get; set; }

        /// <summary>
        /// The SKU family, used when the service has multiple performance classes within a tier, such as 'A', 'D', etc. for virtual machines
        /// </summary>
        [Input("family")]
        public Input<string>? Family { get; set; }

        /// <summary>
        /// The unique name of the SKU, such as 'P3'
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The size of the SKU, used when the name alone does not denote a service size or when a SKU has multiple performance classes within a family, e.g. 'A1' for virtual machines
        /// </summary>
        [Input("size")]
        public Input<string>? Size { get; set; }

        /// <summary>
        /// The tier of the SKU, such as 'Free', 'Basic', 'Standard', or 'Premium'
        /// </summary>
        [Input("tier")]
        public Input<string>? Tier { get; set; }

        public ServiceSkuArgs()
        {
        }
    }
}
