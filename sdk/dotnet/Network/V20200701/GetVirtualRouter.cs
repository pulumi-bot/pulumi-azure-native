// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Network.V20200701
{
    public static class GetVirtualRouter
    {
        public static Task<GetVirtualRouterResult> InvokeAsync(GetVirtualRouterArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetVirtualRouterResult>("azure-nextgen:network/v20200701:getVirtualRouter", args ?? new GetVirtualRouterArgs(), options.WithVersion());
    }


    public sealed class GetVirtualRouterArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Expands referenced resources.
        /// </summary>
        [Input("expand")]
        public string? Expand { get; set; }

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the Virtual Router.
        /// </summary>
        [Input("virtualRouterName", required: true)]
        public string VirtualRouterName { get; set; } = null!;

        public GetVirtualRouterArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetVirtualRouterResult
    {
        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// The Gateway on which VirtualRouter is hosted.
        /// </summary>
        public readonly Outputs.SubResourceResponse? HostedGateway;
        /// <summary>
        /// The Subnet on which VirtualRouter is hosted.
        /// </summary>
        public readonly Outputs.SubResourceResponse? HostedSubnet;
        /// <summary>
        /// Resource location.
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// Resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// List of references to VirtualRouterPeerings.
        /// </summary>
        public readonly ImmutableArray<Outputs.SubResourceResponse> Peerings;
        /// <summary>
        /// The provisioning state of the resource.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// VirtualRouter ASN.
        /// </summary>
        public readonly int? VirtualRouterAsn;
        /// <summary>
        /// VirtualRouter IPs.
        /// </summary>
        public readonly ImmutableArray<string> VirtualRouterIps;

        [OutputConstructor]
        private GetVirtualRouterResult(
            string etag,

            Outputs.SubResourceResponse? hostedGateway,

            Outputs.SubResourceResponse? hostedSubnet,

            string? location,

            string name,

            ImmutableArray<Outputs.SubResourceResponse> peerings,

            string provisioningState,

            ImmutableDictionary<string, string>? tags,

            string type,

            int? virtualRouterAsn,

            ImmutableArray<string> virtualRouterIps)
        {
            Etag = etag;
            HostedGateway = hostedGateway;
            HostedSubnet = hostedSubnet;
            Location = location;
            Name = name;
            Peerings = peerings;
            ProvisioningState = provisioningState;
            Tags = tags;
            Type = type;
            VirtualRouterAsn = virtualRouterAsn;
            VirtualRouterIps = virtualRouterIps;
        }
    }
}
