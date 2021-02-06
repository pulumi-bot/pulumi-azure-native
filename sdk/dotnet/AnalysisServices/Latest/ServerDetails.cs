// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.AnalysisServices.Latest
{
    /// <summary>
    /// Represents an instance of an Analysis Services resource.
    /// Latest API Version: 2017-08-01.
    /// </summary>
    [AzureNextGenResourceType("azure-nextgen:analysisservices/latest:ServerDetails")]
    public partial class ServerDetails : Pulumi.CustomResource
    {
        /// <summary>
        /// A collection of AS server administrators
        /// </summary>
        [Output("asAdministrators")]
        public Output<Outputs.ServerAdministratorsResponse?> AsAdministrators { get; private set; } = null!;

        /// <summary>
        /// The SAS container URI to the backup container.
        /// </summary>
        [Output("backupBlobContainerUri")]
        public Output<string?> BackupBlobContainerUri { get; private set; } = null!;

        /// <summary>
        /// The gateway details configured for the AS server.
        /// </summary>
        [Output("gatewayDetails")]
        public Output<Outputs.GatewayDetailsResponse?> GatewayDetails { get; private set; } = null!;

        /// <summary>
        /// The firewall settings for the AS server.
        /// </summary>
        [Output("ipV4FirewallSettings")]
        public Output<Outputs.IPv4FirewallSettingsResponse?> IpV4FirewallSettings { get; private set; } = null!;

        /// <summary>
        /// Location of the Analysis Services resource.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The managed mode of the server (0 = not managed, 1 = managed).
        /// </summary>
        [Output("managedMode")]
        public Output<int?> ManagedMode { get; private set; } = null!;

        /// <summary>
        /// The name of the Analysis Services resource.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The current deployment state of Analysis Services resource. The provisioningState is to indicate states for resource provisioning.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// How the read-write server's participation in the query pool is controlled.&lt;br/&gt;It can have the following values: &lt;ul&gt;&lt;li&gt;readOnly - indicates that the read-write server is intended not to participate in query operations&lt;/li&gt;&lt;li&gt;all - indicates that the read-write server can participate in query operations&lt;/li&gt;&lt;/ul&gt;Specifying readOnly when capacity is 1 results in error.
        /// </summary>
        [Output("querypoolConnectionMode")]
        public Output<string?> QuerypoolConnectionMode { get; private set; } = null!;

        /// <summary>
        /// The full name of the Analysis Services resource.
        /// </summary>
        [Output("serverFullName")]
        public Output<string> ServerFullName { get; private set; } = null!;

        /// <summary>
        /// The server monitor mode for AS server
        /// </summary>
        [Output("serverMonitorMode")]
        public Output<int?> ServerMonitorMode { get; private set; } = null!;

        /// <summary>
        /// The SKU of the Analysis Services resource.
        /// </summary>
        [Output("sku")]
        public Output<Outputs.ResourceSkuResponse> Sku { get; private set; } = null!;

        /// <summary>
        /// The current state of Analysis Services resource. The state is to indicate more states outside of resource provisioning.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// Key-value pairs of additional resource provisioning properties.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The type of the Analysis Services resource.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a ServerDetails resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServerDetails(string name, ServerDetailsArgs args, CustomResourceOptions? options = null)
            : base("azure-nextgen:analysisservices/latest:ServerDetails", name, args ?? new ServerDetailsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServerDetails(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azure-nextgen:analysisservices/latest:ServerDetails", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azure-nextgen:analysisservices/v20160516:ServerDetails"},
                    new Pulumi.Alias { Type = "azure-nextgen:analysisservices/v20170714:ServerDetails"},
                    new Pulumi.Alias { Type = "azure-nextgen:analysisservices/v20170801:ServerDetails"},
                    new Pulumi.Alias { Type = "azure-nextgen:analysisservices/v20170801beta:ServerDetails"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ServerDetails resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServerDetails Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ServerDetails(name, id, options);
        }
    }

    public sealed class ServerDetailsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A collection of AS server administrators
        /// </summary>
        [Input("asAdministrators")]
        public Input<Inputs.ServerAdministratorsArgs>? AsAdministrators { get; set; }

        /// <summary>
        /// The SAS container URI to the backup container.
        /// </summary>
        [Input("backupBlobContainerUri")]
        public Input<string>? BackupBlobContainerUri { get; set; }

        /// <summary>
        /// The gateway details configured for the AS server.
        /// </summary>
        [Input("gatewayDetails")]
        public Input<Inputs.GatewayDetailsArgs>? GatewayDetails { get; set; }

        /// <summary>
        /// The firewall settings for the AS server.
        /// </summary>
        [Input("ipV4FirewallSettings")]
        public Input<Inputs.IPv4FirewallSettingsArgs>? IpV4FirewallSettings { get; set; }

        /// <summary>
        /// Location of the Analysis Services resource.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The managed mode of the server (0 = not managed, 1 = managed).
        /// </summary>
        [Input("managedMode")]
        public Input<int>? ManagedMode { get; set; }

        /// <summary>
        /// How the read-write server's participation in the query pool is controlled.&lt;br/&gt;It can have the following values: &lt;ul&gt;&lt;li&gt;readOnly - indicates that the read-write server is intended not to participate in query operations&lt;/li&gt;&lt;li&gt;all - indicates that the read-write server can participate in query operations&lt;/li&gt;&lt;/ul&gt;Specifying readOnly when capacity is 1 results in error.
        /// </summary>
        [Input("querypoolConnectionMode")]
        public Input<Pulumi.AzureNextGen.AnalysisServices.Latest.ConnectionMode>? QuerypoolConnectionMode { get; set; }

        /// <summary>
        /// The name of the Azure Resource group of which a given Analysis Services server is part. This name must be at least 1 character in length, and no more than 90.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The server monitor mode for AS server
        /// </summary>
        [Input("serverMonitorMode")]
        public Input<int>? ServerMonitorMode { get; set; }

        /// <summary>
        /// The name of the Analysis Services server. It must be a minimum of 3 characters, and a maximum of 63.
        /// </summary>
        [Input("serverName", required: true)]
        public Input<string> ServerName { get; set; } = null!;

        /// <summary>
        /// The SKU of the Analysis Services resource.
        /// </summary>
        [Input("sku", required: true)]
        public Input<Inputs.ResourceSkuArgs> Sku { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Key-value pairs of additional resource provisioning properties.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public ServerDetailsArgs()
        {
            ManagedMode = 1;
            QuerypoolConnectionMode = Pulumi.AzureNextGen.AnalysisServices.Latest.ConnectionMode.All;
            ServerMonitorMode = 1;
        }
    }
}
