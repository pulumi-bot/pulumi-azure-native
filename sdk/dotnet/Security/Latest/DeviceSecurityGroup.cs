// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Security.Latest
{
    /// <summary>
    /// The device security group resource
    /// 
    /// ## Example Usage
    /// ### Create or update a device security group for the specified IoT hub resource
    /// ```csharp
    /// using Pulumi;
    /// using AzureRM = Pulumi.AzureRM;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var deviceSecurityGroup = new AzureRM.Security.Latest.DeviceSecurityGroup("deviceSecurityGroup", new AzureRM.Security.Latest.DeviceSecurityGroupArgs
    ///         {
    ///             DeviceSecurityGroupName = "samplesecuritygroup",
    ///             ResourceId = "subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/SampleRG/providers/Microsoft.Devices/iotHubs/sampleiothub",
    ///             TimeWindowRules = 
    ///             {
    ///                 new AzureRM.Security.Latest.Inputs.TimeWindowCustomAlertRuleArgs
    ///                 {
    ///                     IsEnabled = true,
    ///                     MaxThreshold = 30,
    ///                     MinThreshold = 0,
    ///                     RuleType = "ActiveConnectionsNotInAllowedRange",
    ///                     TimeWindowSize = "PT05M",
    ///                 },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// 
    /// ```
    /// </summary>
    public partial class DeviceSecurityGroup : Pulumi.CustomResource
    {
        /// <summary>
        /// The allow-list custom alert rules.
        /// </summary>
        [Output("allowlistRules")]
        public Output<ImmutableArray<Outputs.AllowlistCustomAlertRuleResponseResult>> AllowlistRules { get; private set; } = null!;

        /// <summary>
        /// The deny-list custom alert rules.
        /// </summary>
        [Output("denylistRules")]
        public Output<ImmutableArray<Outputs.DenylistCustomAlertRuleResponseResult>> DenylistRules { get; private set; } = null!;

        /// <summary>
        /// Resource name
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The list of custom alert threshold rules.
        /// </summary>
        [Output("thresholdRules")]
        public Output<ImmutableArray<Outputs.ThresholdCustomAlertRuleResponseResult>> ThresholdRules { get; private set; } = null!;

        /// <summary>
        /// The list of custom alert time-window rules.
        /// </summary>
        [Output("timeWindowRules")]
        public Output<ImmutableArray<Outputs.TimeWindowCustomAlertRuleResponseResult>> TimeWindowRules { get; private set; } = null!;

        /// <summary>
        /// Resource type
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a DeviceSecurityGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DeviceSecurityGroup(string name, DeviceSecurityGroupArgs args, CustomResourceOptions? options = null)
            : base("azurerm:security/latest:DeviceSecurityGroup", name, args ?? new DeviceSecurityGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DeviceSecurityGroup(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:security/latest:DeviceSecurityGroup", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:security/v20170801preview:DeviceSecurityGroup"},
                    new Pulumi.Alias { Type = "azurerm:security/v20190801:DeviceSecurityGroup"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing DeviceSecurityGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DeviceSecurityGroup Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new DeviceSecurityGroup(name, id, options);
        }
    }

    public sealed class DeviceSecurityGroupArgs : Pulumi.ResourceArgs
    {
        [Input("allowlistRules")]
        private InputList<Inputs.AllowlistCustomAlertRuleArgs>? _allowlistRules;

        /// <summary>
        /// The allow-list custom alert rules.
        /// </summary>
        public InputList<Inputs.AllowlistCustomAlertRuleArgs> AllowlistRules
        {
            get => _allowlistRules ?? (_allowlistRules = new InputList<Inputs.AllowlistCustomAlertRuleArgs>());
            set => _allowlistRules = value;
        }

        [Input("denylistRules")]
        private InputList<Inputs.DenylistCustomAlertRuleArgs>? _denylistRules;

        /// <summary>
        /// The deny-list custom alert rules.
        /// </summary>
        public InputList<Inputs.DenylistCustomAlertRuleArgs> DenylistRules
        {
            get => _denylistRules ?? (_denylistRules = new InputList<Inputs.DenylistCustomAlertRuleArgs>());
            set => _denylistRules = value;
        }

        /// <summary>
        /// The name of the device security group. Note that the name of the device security group is case insensitive.
        /// </summary>
        [Input("deviceSecurityGroupName", required: true)]
        public Input<string> DeviceSecurityGroupName { get; set; } = null!;

        /// <summary>
        /// The identifier of the resource.
        /// </summary>
        [Input("resourceId", required: true)]
        public Input<string> ResourceId { get; set; } = null!;

        [Input("thresholdRules")]
        private InputList<Inputs.ThresholdCustomAlertRuleArgs>? _thresholdRules;

        /// <summary>
        /// The list of custom alert threshold rules.
        /// </summary>
        public InputList<Inputs.ThresholdCustomAlertRuleArgs> ThresholdRules
        {
            get => _thresholdRules ?? (_thresholdRules = new InputList<Inputs.ThresholdCustomAlertRuleArgs>());
            set => _thresholdRules = value;
        }

        [Input("timeWindowRules")]
        private InputList<Inputs.TimeWindowCustomAlertRuleArgs>? _timeWindowRules;

        /// <summary>
        /// The list of custom alert time-window rules.
        /// </summary>
        public InputList<Inputs.TimeWindowCustomAlertRuleArgs> TimeWindowRules
        {
            get => _timeWindowRules ?? (_timeWindowRules = new InputList<Inputs.TimeWindowCustomAlertRuleArgs>());
            set => _timeWindowRules = value;
        }

        public DeviceSecurityGroupArgs()
        {
        }
    }
}
