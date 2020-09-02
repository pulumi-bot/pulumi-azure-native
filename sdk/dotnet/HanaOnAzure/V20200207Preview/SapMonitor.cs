// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.HanaOnAzure.V20200207Preview
{
    /// <summary>
    /// SAP monitor info on Azure (ARM properties and SAP monitor properties)
    /// </summary>
    public partial class SapMonitor : Pulumi.CustomResource
    {
        /// <summary>
        /// The value indicating whether to send analytics to Microsoft
        /// </summary>
        [Output("enableCustomerAnalytics")]
        public Output<bool?> EnableCustomerAnalytics { get; private set; } = null!;

        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The ARM ID of the Log Analytics Workspace that is used for monitoring
        /// </summary>
        [Output("logAnalyticsWorkspaceArmId")]
        public Output<string?> LogAnalyticsWorkspaceArmId { get; private set; } = null!;

        /// <summary>
        /// The workspace ID of the log analytics workspace to be used for monitoring
        /// </summary>
        [Output("logAnalyticsWorkspaceId")]
        public Output<string?> LogAnalyticsWorkspaceId { get; private set; } = null!;

        /// <summary>
        /// The shared key of the log analytics workspace that is used for monitoring
        /// </summary>
        [Output("logAnalyticsWorkspaceSharedKey")]
        public Output<string?> LogAnalyticsWorkspaceSharedKey { get; private set; } = null!;

        /// <summary>
        /// The name of the resource group the SAP Monitor resources get deployed into.
        /// </summary>
        [Output("managedResourceGroupName")]
        public Output<string> ManagedResourceGroupName { get; private set; } = null!;

        /// <summary>
        /// The subnet which the SAP monitor will be deployed in
        /// </summary>
        [Output("monitorSubnet")]
        public Output<string?> MonitorSubnet { get; private set; } = null!;

        /// <summary>
        /// The name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// State of provisioning of the HanaInstance
        /// </summary>
        [Output("provisioningState")]
        public Output<string> ProvisioningState { get; private set; } = null!;

        /// <summary>
        /// The version of the payload running in the Collector VM
        /// </summary>
        [Output("sapMonitorCollectorVersion")]
        public Output<string> SapMonitorCollectorVersion { get; private set; } = null!;

        /// <summary>
        /// Resource tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a SapMonitor resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SapMonitor(string name, SapMonitorArgs args, CustomResourceOptions? options = null)
            : base("azurerm:hanaonazure/v20200207preview:SapMonitor", name, args ?? new SapMonitorArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SapMonitor(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:hanaonazure/v20200207preview:SapMonitor", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:hanaonazure/v20171103preview:SapMonitor"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing SapMonitor resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SapMonitor Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new SapMonitor(name, id, options);
        }
    }

    public sealed class SapMonitorArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The value indicating whether to send analytics to Microsoft
        /// </summary>
        [Input("enableCustomerAnalytics")]
        public Input<bool>? EnableCustomerAnalytics { get; set; }

        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// The ARM ID of the Log Analytics Workspace that is used for monitoring
        /// </summary>
        [Input("logAnalyticsWorkspaceArmId")]
        public Input<string>? LogAnalyticsWorkspaceArmId { get; set; }

        /// <summary>
        /// The workspace ID of the log analytics workspace to be used for monitoring
        /// </summary>
        [Input("logAnalyticsWorkspaceId")]
        public Input<string>? LogAnalyticsWorkspaceId { get; set; }

        /// <summary>
        /// The shared key of the log analytics workspace that is used for monitoring
        /// </summary>
        [Input("logAnalyticsWorkspaceSharedKey")]
        public Input<string>? LogAnalyticsWorkspaceSharedKey { get; set; }

        /// <summary>
        /// The subnet which the SAP monitor will be deployed in
        /// </summary>
        [Input("monitorSubnet")]
        public Input<string>? MonitorSubnet { get; set; }

        /// <summary>
        /// Name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Name of the SAP monitor resource.
        /// </summary>
        [Input("sapMonitorName", required: true)]
        public Input<string> SapMonitorName { get; set; } = null!;

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

        public SapMonitorArgs()
        {
        }
    }
}
