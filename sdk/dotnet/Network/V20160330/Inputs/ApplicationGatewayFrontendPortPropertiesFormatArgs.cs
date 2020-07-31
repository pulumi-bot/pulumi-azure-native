// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20160330.Inputs
{

    /// <summary>
    /// Properties of Frontend Port of application gateway
    /// </summary>
    public sealed class ApplicationGatewayFrontendPortPropertiesFormatArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Gets or sets the frontend port
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// Gets or sets Provisioning state of the frontend port resource Updating/Deleting/Failed
        /// </summary>
        [Input("provisioningState")]
        public Input<string>? ProvisioningState { get; set; }

        public ApplicationGatewayFrontendPortPropertiesFormatArgs()
        {
        }
    }
}