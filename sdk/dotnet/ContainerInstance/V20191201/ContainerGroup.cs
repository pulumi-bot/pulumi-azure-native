// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ContainerInstance.V20191201
{
    /// <summary>
    /// A container group.
    /// 
    /// ## Example Usage
    /// ### ContainerGroupsCreateOrUpdate
    /// ```csharp
    /// using Pulumi;
    /// using AzureRM = Pulumi.AzureRM;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var containerGroup = new AzureRM.ContainerInstance.V20191201.ContainerGroup("containerGroup", new AzureRM.ContainerInstance.V20191201.ContainerGroupArgs
    ///         {
    ///             ContainerGroupName = "demo1",
    ///             Containers = 
    ///             {
    ///                 new AzureRM.ContainerInstance.V20191201.Inputs.ContainerArgs
    ///                 {
    ///                     Name = "demo1",
    ///                 },
    ///             },
    ///             Diagnostics = new AzureRM.ContainerInstance.V20191201.Inputs.ContainerGroupDiagnosticsArgs
    ///             {
    ///                 LogAnalytics = new AzureRM.ContainerInstance.V20191201.Inputs.LogAnalyticsArgs
    ///                 {
    ///                     LogType = "ContainerInsights",
    ///                     Metadata = 
    ///                     {
    ///                         { "test-key", "test-metadata-value" },
    ///                     },
    ///                     WorkspaceId = "workspaceid",
    ///                     WorkspaceKey = "workspaceKey",
    ///                 },
    ///             },
    ///             DnsConfig = new AzureRM.ContainerInstance.V20191201.Inputs.DnsConfigurationArgs
    ///             {
    ///                 NameServers = 
    ///                 {
    ///                     "1.1.1.1",
    ///                 },
    ///                 Options = "ndots:2",
    ///                 SearchDomains = "cluster.local svc.cluster.local",
    ///             },
    ///             Identity = new AzureRM.ContainerInstance.V20191201.Inputs.ContainerGroupIdentityArgs
    ///             {
    ///                 Type = "SystemAssigned, UserAssigned",
    ///             },
    ///             ImageRegistryCredentials = {},
    ///             IpAddress = new AzureRM.ContainerInstance.V20191201.Inputs.IpAddressArgs
    ///             {
    ///                 DnsNameLabel = "dnsnamelabel1",
    ///                 Ports = 
    ///                 {
    ///                     new AzureRM.ContainerInstance.V20191201.Inputs.PortArgs
    ///                     {
    ///                         Port = 80,
    ///                         Protocol = "TCP",
    ///                     },
    ///                 },
    ///                 Type = "Public",
    ///             },
    ///             Location = "west us",
    ///             NetworkProfile = new AzureRM.ContainerInstance.V20191201.Inputs.ContainerGroupNetworkProfileArgs
    ///             {
    ///                 Id = "test-network-profile-id",
    ///             },
    ///             OsType = "Linux",
    ///             ResourceGroupName = "demo",
    ///             Volumes = 
    ///             {
    ///                 new AzureRM.ContainerInstance.V20191201.Inputs.VolumeArgs
    ///                 {
    ///                     AzureFile = new AzureRM.ContainerInstance.V20191201.Inputs.AzureFileVolumeArgs
    ///                     {
    ///                         ShareName = "shareName",
    ///                         StorageAccountKey = "accountKey",
    ///                         StorageAccountName = "accountName",
    ///                     },
    ///                     Name = "volume1",
    ///                 },
    ///                 new AzureRM.ContainerInstance.V20191201.Inputs.VolumeArgs
    ///                 {
    ///                     EmptyDir = ,
    ///                     Name = "volume2",
    ///                 },
    ///                 new AzureRM.ContainerInstance.V20191201.Inputs.VolumeArgs
    ///                 {
    ///                     Name = "volume3",
    ///                     Secret = 
    ///                     {
    ///                         { "secretKey1", "SecretValue1InBase64" },
    ///                         { "secretKey2", "SecretValue2InBase64" },
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
    public partial class ContainerGroup : Pulumi.CustomResource
    {
        /// <summary>
        /// The containers within the container group.
        /// </summary>
        [Output("containers")]
        public Output<ImmutableArray<Outputs.ContainerResponseResult>> Containers { get; private set; } = null!;

        /// <summary>
        /// The diagnostic information for a container group.
        /// </summary>
        [Output("diagnostics")]
        public Output<Outputs.ContainerGroupDiagnosticsResponseResult?> Diagnostics { get; private set; } = null!;

        /// <summary>
        /// The DNS config information for a container group.
        /// </summary>
        [Output("dnsConfig")]
        public Output<Outputs.DnsConfigurationResponseResult?> DnsConfig { get; private set; } = null!;

        /// <summary>
        /// The encryption properties for a container group.
        /// </summary>
        [Output("encryptionProperties")]
        public Output<Outputs.EncryptionPropertiesResponseResult?> EncryptionProperties { get; private set; } = null!;

        /// <summary>
        /// The identity of the container group, if configured.
        /// </summary>
        [Output("identity")]
        public Output<Outputs.ContainerGroupIdentityResponseResult?> Identity { get; private set; } = null!;

        /// <summary>
        /// The image registry credentials by which the container group is created from.
        /// </summary>
        [Output("imageRegistryCredentials")]
        public Output<ImmutableArray<Outputs.ImageRegistryCredentialResponseResult>> ImageRegistryCredentials { get; private set; } = null!;

        /// <summary>
        /// The init containers for a container group.
        /// </summary>
        [Output("initContainers")]
        public Output<ImmutableArray<Outputs.InitContainerDefinitionResponseResult>> InitContainers { get; private set; } = null!;

        /// <summary>
        /// The instance view of the container group. Only valid in response.
        /// </summary>
        [Output("instanceView")]
        public Output<Outputs.ContainerGroupResponseInstanceViewResult> InstanceView { get; private set; } = null!;

        /// <summary>
        /// The IP address type of the container group.
        /// </summary>
        [Output("ipAddress")]
        public Output<Outputs.IpAddressResponseResult?> IpAddress { get; private set; } = null!;

        /// <summary>
        /// The resource location.
        /// </summary>
        [Output("location")]
        public Output<string?> Location { get; private set; } = null!;

        /// <summary>
        /// The resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The network profile information for a container group.
        /// </summary>
        [Output("networkProfile")]
        public Output<Outputs.ContainerGroupNetworkProfileResponseResult?> NetworkProfile { get; private set; } = null!;

        /// <summary>
        /// The operating system type required by the containers in the container group.
        /// </summary>
        [Output("osType")]
        public Output<string> OsType { get; private set; } = null!;

        /// <summary>
        /// The provisioning state of the container group. This only appears in the response.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// Restart policy for all containers within the container group. 
        /// - `Always` Always restart
        /// - `OnFailure` Restart on failure
        /// - `Never` Never restart
        /// </summary>
        [Output("restartPolicy")]
        public Output<string?> RestartPolicy { get; private set; } = null!;

        /// <summary>
        /// The SKU for a container group.
        /// </summary>
        [Output("sku")]
        public Output<string?> Sku { get; private set; } = null!;

        /// <summary>
        /// The resource tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// The list of volumes that can be mounted by containers in this container group.
        /// </summary>
        [Output("volumes")]
        public Output<ImmutableArray<Outputs.VolumeResponseResult>> Volumes { get; private set; } = null!;


        /// <summary>
        /// Create a ContainerGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ContainerGroup(string name, ContainerGroupArgs args, CustomResourceOptions? options = null)
            : base("azurerm:containerinstance/v20191201:ContainerGroup", name, args ?? new ContainerGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ContainerGroup(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:containerinstance/v20191201:ContainerGroup", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:containerinstance/latest:ContainerGroup"},
                    new Pulumi.Alias { Type = "azurerm:containerinstance/v20170801preview:ContainerGroup"},
                    new Pulumi.Alias { Type = "azurerm:containerinstance/v20171001preview:ContainerGroup"},
                    new Pulumi.Alias { Type = "azurerm:containerinstance/v20171201preview:ContainerGroup"},
                    new Pulumi.Alias { Type = "azurerm:containerinstance/v20180201preview:ContainerGroup"},
                    new Pulumi.Alias { Type = "azurerm:containerinstance/v20180401:ContainerGroup"},
                    new Pulumi.Alias { Type = "azurerm:containerinstance/v20180601:ContainerGroup"},
                    new Pulumi.Alias { Type = "azurerm:containerinstance/v20180901:ContainerGroup"},
                    new Pulumi.Alias { Type = "azurerm:containerinstance/v20181001:ContainerGroup"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ContainerGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ContainerGroup Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ContainerGroup(name, id, options);
        }
    }

    public sealed class ContainerGroupArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the container group.
        /// </summary>
        [Input("containerGroupName", required: true)]
        public Input<string> ContainerGroupName { get; set; } = null!;

        [Input("containers", required: true)]
        private InputList<Inputs.ContainerArgs>? _containers;

        /// <summary>
        /// The containers within the container group.
        /// </summary>
        public InputList<Inputs.ContainerArgs> Containers
        {
            get => _containers ?? (_containers = new InputList<Inputs.ContainerArgs>());
            set => _containers = value;
        }

        /// <summary>
        /// The diagnostic information for a container group.
        /// </summary>
        [Input("diagnostics")]
        public Input<Inputs.ContainerGroupDiagnosticsArgs>? Diagnostics { get; set; }

        /// <summary>
        /// The DNS config information for a container group.
        /// </summary>
        [Input("dnsConfig")]
        public Input<Inputs.DnsConfigurationArgs>? DnsConfig { get; set; }

        /// <summary>
        /// The encryption properties for a container group.
        /// </summary>
        [Input("encryptionProperties")]
        public Input<Inputs.EncryptionPropertiesArgs>? EncryptionProperties { get; set; }

        /// <summary>
        /// The identity of the container group, if configured.
        /// </summary>
        [Input("identity")]
        public Input<Inputs.ContainerGroupIdentityArgs>? Identity { get; set; }

        [Input("imageRegistryCredentials")]
        private InputList<Inputs.ImageRegistryCredentialArgs>? _imageRegistryCredentials;

        /// <summary>
        /// The image registry credentials by which the container group is created from.
        /// </summary>
        public InputList<Inputs.ImageRegistryCredentialArgs> ImageRegistryCredentials
        {
            get => _imageRegistryCredentials ?? (_imageRegistryCredentials = new InputList<Inputs.ImageRegistryCredentialArgs>());
            set => _imageRegistryCredentials = value;
        }

        [Input("initContainers")]
        private InputList<Inputs.InitContainerDefinitionArgs>? _initContainers;

        /// <summary>
        /// The init containers for a container group.
        /// </summary>
        public InputList<Inputs.InitContainerDefinitionArgs> InitContainers
        {
            get => _initContainers ?? (_initContainers = new InputList<Inputs.InitContainerDefinitionArgs>());
            set => _initContainers = value;
        }

        /// <summary>
        /// The IP address type of the container group.
        /// </summary>
        [Input("ipAddress")]
        public Input<Inputs.IpAddressArgs>? IpAddress { get; set; }

        /// <summary>
        /// The resource location.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The network profile information for a container group.
        /// </summary>
        [Input("networkProfile")]
        public Input<Inputs.ContainerGroupNetworkProfileArgs>? NetworkProfile { get; set; }

        /// <summary>
        /// The operating system type required by the containers in the container group.
        /// </summary>
        [Input("osType", required: true)]
        public Input<string> OsType { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Restart policy for all containers within the container group. 
        /// - `Always` Always restart
        /// - `OnFailure` Restart on failure
        /// - `Never` Never restart
        /// </summary>
        [Input("restartPolicy")]
        public Input<string>? RestartPolicy { get; set; }

        /// <summary>
        /// The SKU for a container group.
        /// </summary>
        [Input("sku")]
        public Input<string>? Sku { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// The resource tags.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("volumes")]
        private InputList<Inputs.VolumeArgs>? _volumes;

        /// <summary>
        /// The list of volumes that can be mounted by containers in this container group.
        /// </summary>
        public InputList<Inputs.VolumeArgs> Volumes
        {
            get => _volumes ?? (_volumes = new InputList<Inputs.VolumeArgs>());
            set => _volumes = value;
        }

        public ContainerGroupArgs()
        {
        }
    }
}
