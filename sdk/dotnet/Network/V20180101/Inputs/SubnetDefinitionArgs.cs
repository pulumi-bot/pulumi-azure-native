// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20180101.Inputs
{

    /// <summary>
    /// Subnet in a virtual network resource.
    /// </summary>
    public sealed class SubnetDefinitionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The address prefix for the subnet.
        /// </summary>
        [Input("addressPrefix")]
        public Input<string>? AddressPrefix { get; set; }

        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// Resource ID.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The name of the resource that is unique within a resource group. This name can be used to access the resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The reference of the NetworkSecurityGroup resource.
        /// </summary>
        [Input("networkSecurityGroup")]
        public Input<Inputs.NetworkSecurityGroupArgs>? NetworkSecurityGroup { get; set; }

        /// <summary>
        /// The provisioning state of the resource.
        /// </summary>
        [Input("provisioningState")]
        public Input<string>? ProvisioningState { get; set; }

        [Input("resourceNavigationLinks")]
        private InputList<Inputs.ResourceNavigationLinkArgs>? _resourceNavigationLinks;

        /// <summary>
        /// Gets an array of references to the external resources using subnet.
        /// </summary>
        public InputList<Inputs.ResourceNavigationLinkArgs> ResourceNavigationLinks
        {
            get => _resourceNavigationLinks ?? (_resourceNavigationLinks = new InputList<Inputs.ResourceNavigationLinkArgs>());
            set => _resourceNavigationLinks = value;
        }

        /// <summary>
        /// The reference of the RouteTable resource.
        /// </summary>
        [Input("routeTable")]
        public Input<Inputs.RouteTableArgs>? RouteTable { get; set; }

        [Input("serviceEndpoints")]
        private InputList<Inputs.ServiceEndpointPropertiesFormatArgs>? _serviceEndpoints;

        /// <summary>
        /// An array of service endpoints.
        /// </summary>
        public InputList<Inputs.ServiceEndpointPropertiesFormatArgs> ServiceEndpoints
        {
            get => _serviceEndpoints ?? (_serviceEndpoints = new InputList<Inputs.ServiceEndpointPropertiesFormatArgs>());
            set => _serviceEndpoints = value;
        }

        public SubnetDefinitionArgs()
        {
        }
    }
}
