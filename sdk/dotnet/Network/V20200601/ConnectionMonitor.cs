// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20200601
{
    /// <summary>
    /// Information about the connection monitor.
    /// 
    /// ## Example Usage
    /// ### Create connection monitor V1
    /// ```csharp
    /// using Pulumi;
    /// using AzureRM = Pulumi.AzureRM;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var connectionMonitor = new AzureRM.Network.V20200601.ConnectionMonitor("connectionMonitor", new AzureRM.Network.V20200601.ConnectionMonitorArgs
    ///         {
    ///             ConnectionMonitorName = "cm1",
    ///             Destination = new AzureRM.Network.V20200601.Inputs.ConnectionMonitorDestinationArgs
    ///             {
    ///                 Address = "bing.com",
    ///                 Port = 80,
    ///             },
    ///             MonitoringIntervalInSeconds = 60,
    ///             NetworkWatcherName = "nw1",
    ///             ResourceGroupName = "rg1",
    ///             Source = new AzureRM.Network.V20200601.Inputs.ConnectionMonitorSourceArgs
    ///             {
    ///                 ResourceId = "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Compute/virtualMachines/vm1",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// 
    /// ```
    /// ### Create connection monitor V2
    /// ```csharp
    /// using Pulumi;
    /// using AzureRM = Pulumi.AzureRM;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var connectionMonitor = new AzureRM.Network.V20200601.ConnectionMonitor("connectionMonitor", new AzureRM.Network.V20200601.ConnectionMonitorArgs
    ///         {
    ///             ConnectionMonitorName = "cm1",
    ///             Endpoints = 
    ///             {
    ///                 new AzureRM.Network.V20200601.Inputs.ConnectionMonitorEndpointArgs
    ///                 {
    ///                     Name = "vm1",
    ///                     ResourceId = "/subscriptions/96e68903-0a56-4819-9987-8d08ad6a1f99/resourceGroups/NwRgIrinaCentralUSEUAP/providers/Microsoft.Compute/virtualMachines/vm1",
    ///                 },
    ///                 new AzureRM.Network.V20200601.Inputs.ConnectionMonitorEndpointArgs
    ///                 {
    ///                     Filter = new AzureRM.Network.V20200601.Inputs.ConnectionMonitorEndpointFilterArgs
    ///                     {
    ///                         Items = 
    ///                         {
    ///                             new AzureRM.Network.V20200601.Inputs.ConnectionMonitorEndpointFilterItemArgs
    ///                             {
    ///                                 Address = "npmuser",
    ///                                 Type = "AgentAddress",
    ///                             },
    ///                         },
    ///                         Type = "Include",
    ///                     },
    ///                     Name = "CanaryWorkspaceVamshi",
    ///                     ResourceId = "/subscriptions/96e68903-0a56-4819-9987-8d08ad6a1f99/resourceGroups/vasamudrRG/providers/Microsoft.OperationalInsights/workspaces/vasamudrWorkspace",
    ///                 },
    ///                 new AzureRM.Network.V20200601.Inputs.ConnectionMonitorEndpointArgs
    ///                 {
    ///                     Address = "bing.com",
    ///                     Name = "bing",
    ///                 },
    ///                 new AzureRM.Network.V20200601.Inputs.ConnectionMonitorEndpointArgs
    ///                 {
    ///                     Address = "google.com",
    ///                     Name = "google",
    ///                 },
    ///             },
    ///             NetworkWatcherName = "nw1",
    ///             Outputs = {},
    ///             ResourceGroupName = "rg1",
    ///             TestConfigurations = 
    ///             {
    ///                 new AzureRM.Network.V20200601.Inputs.ConnectionMonitorTestConfigurationArgs
    ///                 {
    ///                     Name = "testConfig1",
    ///                     Protocol = "Tcp",
    ///                     TcpConfiguration = new AzureRM.Network.V20200601.Inputs.ConnectionMonitorTcpConfigurationArgs
    ///                     {
    ///                         DisableTraceRoute = false,
    ///                         Port = 80,
    ///                     },
    ///                     TestFrequencySec = 60,
    ///                 },
    ///             },
    ///             TestGroups = 
    ///             {
    ///                 new AzureRM.Network.V20200601.Inputs.ConnectionMonitorTestGroupArgs
    ///                 {
    ///                     Destinations = 
    ///                     {
    ///                         "bing",
    ///                         "google",
    ///                     },
    ///                     Disable = false,
    ///                     Name = "test1",
    ///                     Sources = 
    ///                     {
    ///                         "vm1",
    ///                         "CanaryWorkspaceVamshi",
    ///                     },
    ///                     TestConfigurations = 
    ///                     {
    ///                         "testConfig1",
    ///                     },
    ///                 },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// 
    /// ```
    /// </summary>
    public partial class ConnectionMonitor : Pulumi.CustomResource
    {
        /// <summary>
        /// Determines if the connection monitor will start automatically once created.
        /// </summary>
        [Output("autoStart")]
        public Output<bool?> AutoStart { get; private set; } = null!;

        /// <summary>
        /// Type of connection monitor.
        /// </summary>
        [Output("connectionMonitorType")]
        public Output<string> ConnectionMonitorType { get; private set; } = null!;

        /// <summary>
        /// Describes the destination of connection monitor.
        /// </summary>
        [Output("destination")]
        public Output<Outputs.ConnectionMonitorDestinationResponseResult?> Destination { get; private set; } = null!;

        /// <summary>
        /// List of connection monitor endpoints.
        /// </summary>
        [Output("endpoints")]
        public Output<ImmutableArray<Outputs.ConnectionMonitorEndpointResponseResult>> Endpoints { get; private set; } = null!;

        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// Connection monitor location.
        /// </summary>
        [Output("location")]
        public Output<string?> Location { get; private set; } = null!;

        /// <summary>
        /// Monitoring interval in seconds.
        /// </summary>
        [Output("monitoringIntervalInSeconds")]
        public Output<int?> MonitoringIntervalInSeconds { get; private set; } = null!;

        /// <summary>
        /// The monitoring status of the connection monitor.
        /// </summary>
        [Output("monitoringStatus")]
        public Output<string> MonitoringStatus { get; private set; } = null!;

        /// <summary>
        /// Name of the connection monitor.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Optional notes to be associated with the connection monitor.
        /// </summary>
        [Output("notes")]
        public Output<string?> Notes { get; private set; } = null!;

        /// <summary>
        /// List of connection monitor outputs.
        /// </summary>
        [Output("outputs")]
        public Output<ImmutableArray<Outputs.ConnectionMonitorOutputResponseResult>> Outputs { get; private set; } = null!;

        /// <summary>
        /// The provisioning state of the connection monitor.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// Describes the source of connection monitor.
        /// </summary>
        [Output("source")]
        public Output<Outputs.ConnectionMonitorSourceResponseResult?> Source { get; private set; } = null!;

        /// <summary>
        /// The date and time when the connection monitor was started.
        /// </summary>
        [Output("startTime")]
        public Output<string> StartTime { get; private set; } = null!;

        /// <summary>
        /// Connection monitor tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// List of connection monitor test configurations.
        /// </summary>
        [Output("testConfigurations")]
        public Output<ImmutableArray<Outputs.ConnectionMonitorTestConfigurationResponseResult>> TestConfigurations { get; private set; } = null!;

        /// <summary>
        /// List of connection monitor test groups.
        /// </summary>
        [Output("testGroups")]
        public Output<ImmutableArray<Outputs.ConnectionMonitorTestGroupResponseResult>> TestGroups { get; private set; } = null!;

        /// <summary>
        /// Connection monitor type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a ConnectionMonitor resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ConnectionMonitor(string name, ConnectionMonitorArgs args, CustomResourceOptions? options = null)
            : base("azurerm:network/v20200601:ConnectionMonitor", name, args ?? new ConnectionMonitorArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ConnectionMonitor(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:network/v20200601:ConnectionMonitor", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:network/latest:ConnectionMonitor"},
                    new Pulumi.Alias { Type = "azurerm:network/v20171001:ConnectionMonitor"},
                    new Pulumi.Alias { Type = "azurerm:network/v20171101:ConnectionMonitor"},
                    new Pulumi.Alias { Type = "azurerm:network/v20180101:ConnectionMonitor"},
                    new Pulumi.Alias { Type = "azurerm:network/v20180201:ConnectionMonitor"},
                    new Pulumi.Alias { Type = "azurerm:network/v20180401:ConnectionMonitor"},
                    new Pulumi.Alias { Type = "azurerm:network/v20180601:ConnectionMonitor"},
                    new Pulumi.Alias { Type = "azurerm:network/v20180701:ConnectionMonitor"},
                    new Pulumi.Alias { Type = "azurerm:network/v20180801:ConnectionMonitor"},
                    new Pulumi.Alias { Type = "azurerm:network/v20181001:ConnectionMonitor"},
                    new Pulumi.Alias { Type = "azurerm:network/v20181101:ConnectionMonitor"},
                    new Pulumi.Alias { Type = "azurerm:network/v20181201:ConnectionMonitor"},
                    new Pulumi.Alias { Type = "azurerm:network/v20190201:ConnectionMonitor"},
                    new Pulumi.Alias { Type = "azurerm:network/v20190401:ConnectionMonitor"},
                    new Pulumi.Alias { Type = "azurerm:network/v20190601:ConnectionMonitor"},
                    new Pulumi.Alias { Type = "azurerm:network/v20190701:ConnectionMonitor"},
                    new Pulumi.Alias { Type = "azurerm:network/v20190801:ConnectionMonitor"},
                    new Pulumi.Alias { Type = "azurerm:network/v20190901:ConnectionMonitor"},
                    new Pulumi.Alias { Type = "azurerm:network/v20191101:ConnectionMonitor"},
                    new Pulumi.Alias { Type = "azurerm:network/v20191201:ConnectionMonitor"},
                    new Pulumi.Alias { Type = "azurerm:network/v20200301:ConnectionMonitor"},
                    new Pulumi.Alias { Type = "azurerm:network/v20200401:ConnectionMonitor"},
                    new Pulumi.Alias { Type = "azurerm:network/v20200501:ConnectionMonitor"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ConnectionMonitor resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ConnectionMonitor Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ConnectionMonitor(name, id, options);
        }
    }

    public sealed class ConnectionMonitorArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Determines if the connection monitor will start automatically once created.
        /// </summary>
        [Input("autoStart")]
        public Input<bool>? AutoStart { get; set; }

        /// <summary>
        /// The name of the connection monitor.
        /// </summary>
        [Input("connectionMonitorName", required: true)]
        public Input<string> ConnectionMonitorName { get; set; } = null!;

        /// <summary>
        /// Describes the destination of connection monitor.
        /// </summary>
        [Input("destination")]
        public Input<Inputs.ConnectionMonitorDestinationArgs>? Destination { get; set; }

        [Input("endpoints")]
        private InputList<Inputs.ConnectionMonitorEndpointArgs>? _endpoints;

        /// <summary>
        /// List of connection monitor endpoints.
        /// </summary>
        public InputList<Inputs.ConnectionMonitorEndpointArgs> Endpoints
        {
            get => _endpoints ?? (_endpoints = new InputList<Inputs.ConnectionMonitorEndpointArgs>());
            set => _endpoints = value;
        }

        /// <summary>
        /// Connection monitor location.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Monitoring interval in seconds.
        /// </summary>
        [Input("monitoringIntervalInSeconds")]
        public Input<int>? MonitoringIntervalInSeconds { get; set; }

        /// <summary>
        /// The name of the Network Watcher resource.
        /// </summary>
        [Input("networkWatcherName", required: true)]
        public Input<string> NetworkWatcherName { get; set; } = null!;

        /// <summary>
        /// Optional notes to be associated with the connection monitor.
        /// </summary>
        [Input("notes")]
        public Input<string>? Notes { get; set; }

        [Input("outputs")]
        private InputList<Inputs.ConnectionMonitorOutputArgs>? _outputs;

        /// <summary>
        /// List of connection monitor outputs.
        /// </summary>
        public InputList<Inputs.ConnectionMonitorOutputArgs> Outputs
        {
            get => _outputs ?? (_outputs = new InputList<Inputs.ConnectionMonitorOutputArgs>());
            set => _outputs = value;
        }

        /// <summary>
        /// The name of the resource group containing Network Watcher.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Describes the source of connection monitor.
        /// </summary>
        [Input("source")]
        public Input<Inputs.ConnectionMonitorSourceArgs>? Source { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Connection monitor tags.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("testConfigurations")]
        private InputList<Inputs.ConnectionMonitorTestConfigurationArgs>? _testConfigurations;

        /// <summary>
        /// List of connection monitor test configurations.
        /// </summary>
        public InputList<Inputs.ConnectionMonitorTestConfigurationArgs> TestConfigurations
        {
            get => _testConfigurations ?? (_testConfigurations = new InputList<Inputs.ConnectionMonitorTestConfigurationArgs>());
            set => _testConfigurations = value;
        }

        [Input("testGroups")]
        private InputList<Inputs.ConnectionMonitorTestGroupArgs>? _testGroups;

        /// <summary>
        /// List of connection monitor test groups.
        /// </summary>
        public InputList<Inputs.ConnectionMonitorTestGroupArgs> TestGroups
        {
            get => _testGroups ?? (_testGroups = new InputList<Inputs.ConnectionMonitorTestGroupArgs>());
            set => _testGroups = value;
        }

        public ConnectionMonitorArgs()
        {
        }
    }
}
