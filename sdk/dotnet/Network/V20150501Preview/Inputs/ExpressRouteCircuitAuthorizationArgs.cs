// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20150501Preview.Inputs
{

    /// <summary>
    /// Authorization in a ExpressRouteCircuit resource
    /// </summary>
    public sealed class ExpressRouteCircuitAuthorizationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Gets or sets the authorization key
        /// </summary>
        [Input("authorizationKey")]
        public Input<string>? AuthorizationKey { get; set; }

        /// <summary>
        /// Gets or sets AuthorizationUseStatus
        /// </summary>
        [Input("authorizationUseStatus")]
        public Input<string>? AuthorizationUseStatus { get; set; }

        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// Resource Id
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Gets name of the resource that is unique within a resource group. This name can be used to access the resource
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Gets or sets Provisioning state of the PublicIP resource Updating/Deleting/Failed
        /// </summary>
        [Input("provisioningState")]
        public Input<string>? ProvisioningState { get; set; }

        public ExpressRouteCircuitAuthorizationArgs()
        {
        }
    }
}
