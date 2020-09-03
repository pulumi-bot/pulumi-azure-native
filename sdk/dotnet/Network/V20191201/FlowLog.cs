// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20191201
{
    /// <summary>
    /// A flow log resource.
    /// 
    /// ## Example Usage
    /// ### Create or update flow log
    /// ```csharp
    /// using Pulumi;
    /// using AzureRM = Pulumi.AzureRM;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var flowLog = new AzureRM.Network.V20191201.FlowLog("flowLog", new AzureRM.Network.V20191201.FlowLogArgs
    ///         {
    ///             Enabled = true,
    ///             FlowLogName = "fl",
    ///             Format = new AzureRM.Network.V20191201.Inputs.FlowLogFormatParametersArgs
    ///             {
    ///                 Type = "JSON",
    ///                 Version = 1,
    ///             },
    ///             Location = "centraluseuap",
    ///             NetworkWatcherName = "nw1",
    ///             ResourceGroupName = "rg1",
    ///             StorageId = "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Storage/storageAccounts/nwtest1mgvbfmqsigdxe",
    ///             TargetResourceId = "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/networkSecurityGroups/desmondcentral-nsg",
    ///         });
    ///     }
    /// 
    /// }
    /// 
    /// ```
    /// </summary>
    public partial class FlowLog : Pulumi.CustomResource
    {
        /// <summary>
        /// Flag to enable/disable flow logging.
        /// </summary>
        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;

        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// Parameters that define the configuration of traffic analytics.
        /// </summary>
        [Output("flowAnalyticsConfiguration")]
        public Output<Outputs.TrafficAnalyticsPropertiesResponseResult?> FlowAnalyticsConfiguration { get; private set; } = null!;

        /// <summary>
        /// Parameters that define the flow log format.
        /// </summary>
        [Output("format")]
        public Output<Outputs.FlowLogFormatParametersResponseResult?> Format { get; private set; } = null!;

        /// <summary>
        /// Resource location.
        /// </summary>
        [Output("location")]
        public Output<string?> Location { get; private set; } = null!;

        /// <summary>
        /// Resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The provisioning state of the flow log.
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// Parameters that define the retention policy for flow log.
        /// </summary>
        [Output("retentionPolicy")]
        public Output<Outputs.RetentionPolicyParametersResponseResult?> RetentionPolicy { get; private set; } = null!;

        /// <summary>
        /// ID of the storage account which is used to store the flow log.
        /// </summary>
        [Output("storageId")]
        public Output<string> StorageId { get; private set; } = null!;

        /// <summary>
        /// Resource tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Guid of network security group to which flow log will be applied.
        /// </summary>
        [Output("targetResourceGuid")]
        public Output<string> TargetResourceGuid { get; private set; } = null!;

        /// <summary>
        /// ID of network security group to which flow log will be applied.
        /// </summary>
        [Output("targetResourceId")]
        public Output<string> TargetResourceId { get; private set; } = null!;

        /// <summary>
        /// Resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a FlowLog resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FlowLog(string name, FlowLogArgs args, CustomResourceOptions? options = null)
            : base("azurerm:network/v20191201:FlowLog", name, args ?? new FlowLogArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FlowLog(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:network/v20191201:FlowLog", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:network/latest:FlowLog"},
                    new Pulumi.Alias { Type = "azurerm:network/v20191101:FlowLog"},
                    new Pulumi.Alias { Type = "azurerm:network/v20200301:FlowLog"},
                    new Pulumi.Alias { Type = "azurerm:network/v20200401:FlowLog"},
                    new Pulumi.Alias { Type = "azurerm:network/v20200501:FlowLog"},
                    new Pulumi.Alias { Type = "azurerm:network/v20200601:FlowLog"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing FlowLog resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FlowLog Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new FlowLog(name, id, options);
        }
    }

    public sealed class FlowLogArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Flag to enable/disable flow logging.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// Parameters that define the configuration of traffic analytics.
        /// </summary>
        [Input("flowAnalyticsConfiguration")]
        public Input<Inputs.TrafficAnalyticsPropertiesArgs>? FlowAnalyticsConfiguration { get; set; }

        /// <summary>
        /// The name of the flow log.
        /// </summary>
        [Input("flowLogName", required: true)]
        public Input<string> FlowLogName { get; set; } = null!;

        /// <summary>
        /// Parameters that define the flow log format.
        /// </summary>
        [Input("format")]
        public Input<Inputs.FlowLogFormatParametersArgs>? Format { get; set; }

        /// <summary>
        /// Resource ID.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Resource location.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The name of the network watcher.
        /// </summary>
        [Input("networkWatcherName", required: true)]
        public Input<string> NetworkWatcherName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Parameters that define the retention policy for flow log.
        /// </summary>
        [Input("retentionPolicy")]
        public Input<Inputs.RetentionPolicyParametersArgs>? RetentionPolicy { get; set; }

        /// <summary>
        /// ID of the storage account which is used to store the flow log.
        /// </summary>
        [Input("storageId", required: true)]
        public Input<string> StorageId { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Resource tags.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// ID of network security group to which flow log will be applied.
        /// </summary>
        [Input("targetResourceId", required: true)]
        public Input<string> TargetResourceId { get; set; } = null!;

        public FlowLogArgs()
        {
        }
    }
}
