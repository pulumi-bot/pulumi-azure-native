// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.SignalRService.V20200501
{
    public static class GetSignalR
    {
        public static Task<GetSignalRResult> InvokeAsync(GetSignalRArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSignalRResult>("azure-nextgen:signalrservice/v20200501:getSignalR", args ?? new GetSignalRArgs(), options.WithVersion());
    }


    public sealed class GetSignalRArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the SignalR resource.
        /// </summary>
        [Input("resourceName", required: true)]
        public string ResourceName { get; set; } = null!;

        public GetSignalRArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetSignalRResult
    {
        /// <summary>
        /// Cross-Origin Resource Sharing (CORS) settings.
        /// </summary>
        public readonly Outputs.SignalRCorsSettingsResponse? Cors;
        /// <summary>
        /// The publicly accessible IP of the SignalR service.
        /// </summary>
        public readonly string ExternalIP;
        /// <summary>
        /// List of SignalR featureFlags. e.g. ServiceMode.
        /// 
        /// FeatureFlags that are not included in the parameters for the update operation will not be modified.
        /// And the response will only include featureFlags that are explicitly set. 
        /// When a featureFlag is not explicitly set, SignalR service will use its globally default value. 
        /// But keep in mind, the default value doesn't mean "false". It varies in terms of different FeatureFlags.
        /// </summary>
        public readonly ImmutableArray<Outputs.SignalRFeatureResponse> Features;
        /// <summary>
        /// FQDN of the SignalR service instance. Format: xxx.service.signalr.net
        /// </summary>
        public readonly string HostName;
        /// <summary>
        /// Prefix for the hostName of the SignalR service. Retained for future use.
        /// The hostname will be of format: &amp;lt;hostNamePrefix&amp;gt;.service.signalr.net.
        /// </summary>
        public readonly string HostNamePrefix;
        /// <summary>
        /// Fully qualified resource Id for the resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The kind of the service - e.g. "SignalR", or "RawWebSockets" for "Microsoft.SignalRService/SignalR"
        /// </summary>
        public readonly string? Kind;
        /// <summary>
        /// The GEO location of the SignalR service. e.g. West US | East US | North Central US | South Central US.
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// The name of the resource.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Network ACLs
        /// </summary>
        public readonly Outputs.SignalRNetworkACLsResponse? NetworkACLs;
        /// <summary>
        /// Private endpoint connections to the SignalR resource.
        /// </summary>
        public readonly ImmutableArray<Outputs.PrivateEndpointConnectionResponse> PrivateEndpointConnections;
        /// <summary>
        /// Provisioning state of the resource.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// The publicly accessible port of the SignalR service which is designed for browser/client side usage.
        /// </summary>
        public readonly int PublicPort;
        /// <summary>
        /// The publicly accessible port of the SignalR service which is designed for customer server side usage.
        /// </summary>
        public readonly int ServerPort;
        /// <summary>
        /// The billing information of the resource.(e.g. Free, Standard)
        /// </summary>
        public readonly Outputs.ResourceSkuResponse? Sku;
        /// <summary>
        /// Tags of the service which is a list of key value pairs that describe the resource.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// The type of the resource - e.g. "Microsoft.SignalRService/SignalR"
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Upstream settings when the Azure SignalR is in server-less mode.
        /// </summary>
        public readonly Outputs.ServerlessUpstreamSettingsResponse? Upstream;
        /// <summary>
        /// Version of the SignalR resource. Probably you need the same or higher version of client SDKs.
        /// </summary>
        public readonly string Version;

        [OutputConstructor]
        private GetSignalRResult(
            Outputs.SignalRCorsSettingsResponse? cors,

            string externalIP,

            ImmutableArray<Outputs.SignalRFeatureResponse> features,

            string hostName,

            string hostNamePrefix,

            string id,

            string? kind,

            string? location,

            string name,

            Outputs.SignalRNetworkACLsResponse? networkACLs,

            ImmutableArray<Outputs.PrivateEndpointConnectionResponse> privateEndpointConnections,

            string provisioningState,

            int publicPort,

            int serverPort,

            Outputs.ResourceSkuResponse? sku,

            ImmutableDictionary<string, string>? tags,

            string type,

            Outputs.ServerlessUpstreamSettingsResponse? upstream,

            string version)
        {
            Cors = cors;
            ExternalIP = externalIP;
            Features = features;
            HostName = hostName;
            HostNamePrefix = hostNamePrefix;
            Id = id;
            Kind = kind;
            Location = location;
            Name = name;
            NetworkACLs = networkACLs;
            PrivateEndpointConnections = privateEndpointConnections;
            ProvisioningState = provisioningState;
            PublicPort = publicPort;
            ServerPort = serverPort;
            Sku = sku;
            Tags = tags;
            Type = type;
            Upstream = upstream;
            Version = version;
        }
    }
}
