// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.SignalRService.V20200701Preview
{
    /// <summary>
    /// A class represent a SignalR service resource.
    /// </summary>
    public partial class SignalR : Pulumi.CustomResource
    {
        /// <summary>
        /// Cross-Origin Resource Sharing (CORS) settings.
        /// </summary>
        [Output("cors")]
        public Output<Outputs.SignalRCorsSettingsResponseResult?> Cors { get; private set; } = null!;

        /// <summary>
        /// The publicly accessible IP of the SignalR service.
        /// </summary>
        [Output("externalIP")]
        public Output<string> ExternalIP { get; private set; } = null!;

        /// <summary>
        /// List of SignalR featureFlags. e.g. ServiceMode.
        /// 
        /// FeatureFlags that are not included in the parameters for the update operation will not be modified.
        /// And the response will only include featureFlags that are explicitly set. 
        /// When a featureFlag is not explicitly set, SignalR service will use its globally default value. 
        /// But keep in mind, the default value doesn't mean "false". It varies in terms of different FeatureFlags.
        /// </summary>
        [Output("features")]
        public Output<ImmutableArray<Outputs.SignalRFeatureResponseResult>> Features { get; private set; } = null!;

        /// <summary>
        /// FQDN of the SignalR service instance. Format: xxx.service.signalr.net
        /// </summary>
        [Output("hostName")]
        public Output<string> HostName { get; private set; } = null!;

        /// <summary>
        /// The managed identity response
        /// </summary>
        [Output("identity")]
        public Output<Outputs.ManagedIdentityResponseResult?> Identity { get; private set; } = null!;

        /// <summary>
        /// The kind of the service - e.g. "SignalR", or "RawWebSockets" for "Microsoft.SignalRService/SignalR"
        /// </summary>
        [Output("kind")]
        public Output<string?> Kind { get; private set; } = null!;

        /// <summary>
        /// The GEO location of the SignalR service. e.g. West US | East US | North Central US | South Central US.
        /// </summary>
        [Output("location")]
        public Output<string?> Location { get; private set; } = null!;

        /// <summary>
        /// The name of the resource.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Network ACLs
        /// </summary>
        [Output("networkACLs")]
        public Output<Outputs.SignalRNetworkACLsResponseResult?> NetworkACLs { get; private set; } = null!;

        /// <summary>
        /// Private endpoint connections to the SignalR resource.
        /// </summary>
        [Output("privateEndpointConnections")]
        public Output<ImmutableArray<Outputs.PrivateEndpointConnectionResponseResult>> PrivateEndpointConnections { get; private set; } = null!;

        /// <summary>
        /// Provisioning state of the resource.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// The publicly accessible port of the SignalR service which is designed for browser/client side usage.
        /// </summary>
        [Output("publicPort")]
        public Output<int> PublicPort { get; private set; } = null!;

        /// <summary>
        /// The publicly accessible port of the SignalR service which is designed for customer server side usage.
        /// </summary>
        [Output("serverPort")]
        public Output<int> ServerPort { get; private set; } = null!;

        /// <summary>
        /// The billing information of the resource.(e.g. Free, Standard)
        /// </summary>
        [Output("sku")]
        public Output<Outputs.ResourceSkuResponseResult?> Sku { get; private set; } = null!;

        /// <summary>
        /// Tags of the service which is a list of key value pairs that describe the resource.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// TLS settings.
        /// </summary>
        [Output("tls")]
        public Output<Outputs.SignalRTlsSettingsResponseResult?> Tls { get; private set; } = null!;

        /// <summary>
        /// The type of the resource - e.g. "Microsoft.SignalRService/SignalR"
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// Upstream settings when the Azure SignalR is in server-less mode.
        /// </summary>
        [Output("upstream")]
        public Output<Outputs.ServerlessUpstreamSettingsResponseResult?> Upstream { get; private set; } = null!;

        /// <summary>
        /// Version of the SignalR resource. Probably you need the same or higher version of client SDKs.
        /// </summary>
        [Output("version")]
        public Output<string> Version { get; private set; } = null!;


        /// <summary>
        /// Create a SignalR resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SignalR(string name, SignalRArgs args, CustomResourceOptions? options = null)
            : base("azurerm:signalrservice/v20200701preview:SignalR", name, args ?? new SignalRArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SignalR(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:signalrservice/v20200701preview:SignalR", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:signalrservice/latest:SignalR"},
                    new Pulumi.Alias { Type = "azurerm:signalrservice/v20180301preview:SignalR"},
                    new Pulumi.Alias { Type = "azurerm:signalrservice/v20181001:SignalR"},
                    new Pulumi.Alias { Type = "azurerm:signalrservice/v20200501:SignalR"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing SignalR resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SignalR Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new SignalR(name, id, options);
        }
    }

    public sealed class SignalRArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Cross-Origin Resource Sharing (CORS) settings.
        /// </summary>
        [Input("cors")]
        public Input<Inputs.SignalRCorsSettingsArgs>? Cors { get; set; }

        [Input("features")]
        private InputList<Inputs.SignalRFeatureArgs>? _features;

        /// <summary>
        /// List of SignalR featureFlags. e.g. ServiceMode.
        /// 
        /// FeatureFlags that are not included in the parameters for the update operation will not be modified.
        /// And the response will only include featureFlags that are explicitly set. 
        /// When a featureFlag is not explicitly set, SignalR service will use its globally default value. 
        /// But keep in mind, the default value doesn't mean "false". It varies in terms of different FeatureFlags.
        /// </summary>
        public InputList<Inputs.SignalRFeatureArgs> Features
        {
            get => _features ?? (_features = new InputList<Inputs.SignalRFeatureArgs>());
            set => _features = value;
        }

        /// <summary>
        /// The managed identity response
        /// </summary>
        [Input("identity")]
        public Input<Inputs.ManagedIdentityArgs>? Identity { get; set; }

        /// <summary>
        /// The kind of the service - e.g. "SignalR", or "RawWebSockets" for "Microsoft.SignalRService/SignalR"
        /// </summary>
        [Input("kind")]
        public Input<string>? Kind { get; set; }

        /// <summary>
        /// The GEO location of the SignalR service. e.g. West US | East US | North Central US | South Central US.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Network ACLs
        /// </summary>
        [Input("networkACLs")]
        public Input<Inputs.SignalRNetworkACLsArgs>? NetworkACLs { get; set; }

        /// <summary>
        /// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the SignalR resource.
        /// </summary>
        [Input("resourceName", required: true)]
        public Input<string> ResourceName { get; set; } = null!;

        /// <summary>
        /// The billing information of the resource.(e.g. Free, Standard)
        /// </summary>
        [Input("sku")]
        public Input<Inputs.ResourceSkuArgs>? Sku { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Tags of the service which is a list of key value pairs that describe the resource.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// TLS settings.
        /// </summary>
        [Input("tls")]
        public Input<Inputs.SignalRTlsSettingsArgs>? Tls { get; set; }

        /// <summary>
        /// Upstream settings when the Azure SignalR is in server-less mode.
        /// </summary>
        [Input("upstream")]
        public Input<Inputs.ServerlessUpstreamSettingsArgs>? Upstream { get; set; }

        public SignalRArgs()
        {
        }
    }
}
