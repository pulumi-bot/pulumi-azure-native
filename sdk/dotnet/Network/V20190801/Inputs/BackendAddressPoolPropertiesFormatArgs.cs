// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20190801.Inputs
{

    /// <summary>
    /// Properties of the backend address pool.
    /// </summary>
    public sealed class BackendAddressPoolPropertiesFormatArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The provisioning state of the backend address pool resource.
        /// </summary>
        [Input("provisioningState")]
        public Input<string>? ProvisioningState { get; set; }

        public BackendAddressPoolPropertiesFormatArgs()
        {
        }
    }
}