// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Logic.V20160601
{
    /// <summary>
    /// The integration account RosettaNet process configuration.
    /// 
    /// ## Example Usage
    /// ### Create or update an RosettaNetProcessConfiguration
    /// ```csharp
    /// using Pulumi;
    /// using AzureRM = Pulumi.AzureRM;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var rosettaNetProcessConfiguration = new AzureRM.Logic.V20160601.RosettaNetProcessConfiguration("rosettaNetProcessConfiguration", new AzureRM.Logic.V20160601.RosettaNetProcessConfigurationArgs
    ///         {
    ///             ActivitySettings = new AzureRM.Logic.V20160601.Inputs.RosettaNetPipActivitySettingsArgs
    ///             {
    ///                 AcknowledgmentOfReceiptSettings = new AzureRM.Logic.V20160601.Inputs.RosettaNetPipAcknowledgmentOfReceiptSettingsArgs
    ///                 {
    ///                     IsNonRepudiationRequired = false,
    ///                     TimeToAcknowledgeInSeconds = 600,
    ///                 },
    ///                 ActivityBehavior = new AzureRM.Logic.V20160601.Inputs.RosettaNetPipActivityBehaviorArgs
    ///                 {
    ///                     ActionType = "DoubleAction",
    ///                     IsAuthorizationRequired = false,
    ///                     IsSecuredTransportRequired = false,
    ///                     NonRepudiationOfOriginAndContent = false,
    ///                     PersistentConfidentialityScope = "None",
    ///                     ResponseType = "Async",
    ///                     RetryCount = 2,
    ///                     TimeToPerformInSeconds = 7200,
    ///                 },
    ///                 ActivityType = "RequestResponse",
    ///             },
    ///             Description = "Test description",
    ///             InitiatorRoleSettings = new AzureRM.Logic.V20160601.Inputs.RosettaNetPipRoleSettingsArgs
    ///             {
    ///                 Action = "Purchase Order Request",
    ///                 BusinessDocument = new AzureRM.Logic.V20160601.Inputs.RosettaNetPipBusinessDocumentArgs
    ///                 {
    ///                     Description = "A request to accept a purchase order for fulfillment..",
    ///                     Name = "Purchase Order Request",
    ///                     Version = "V02.02.00",
    ///                 },
    ///                 Description = "This partner role creates a demand for a product or service.",
    ///                 Role = "Buyer",
    ///                 RoleType = "Functional",
    ///                 Service = "Buyer Service",
    ///                 ServiceClassification = "Business Service",
    ///             },
    ///             IntegrationAccountName = "testia123",
    ///             ProcessCode = "3A4",
    ///             ProcessName = "Request Purchase Order",
    ///             ProcessVersion = "V02.02.00",
    ///             ResourceGroupName = "testrg123",
    ///             ResponderRoleSettings = new AzureRM.Logic.V20160601.Inputs.RosettaNetPipRoleSettingsArgs
    ///             {
    ///                 Action = "Purchase Order Confirmation Action",
    ///                 BusinessDocument = new AzureRM.Logic.V20160601.Inputs.RosettaNetPipBusinessDocumentArgs
    ///                 {
    ///                     Description = "Formally confirms the status of line item(s) in a Purchase Order. A Purchase Order line item may have one of the following states: accepted, rejected, or pending.",
    ///                     Name = "Purchase Order Confirmation",
    ///                     Version = "V02.02.00",
    ///                 },
    ///                 Description = "An organization that sells products to partners in the supply chain.",
    ///                 Role = "Seller",
    ///                 RoleType = "Organizational",
    ///                 Service = "Seller Service",
    ///                 ServiceClassification = "Business Service",
    ///             },
    ///             RosettaNetProcessConfigurationName = "3A4",
    ///         });
    ///     }
    /// 
    /// }
    /// 
    /// ```
    /// </summary>
    public partial class RosettaNetProcessConfiguration : Pulumi.CustomResource
    {
        /// <summary>
        /// The RosettaNet process configuration activity settings.
        /// </summary>
        [Output("activitySettings")]
        public Output<Outputs.RosettaNetPipActivitySettingsResponseResult> ActivitySettings { get; private set; } = null!;

        /// <summary>
        /// The changed time.
        /// </summary>
        [Output("changedTime")]
        public Output<string> ChangedTime { get; private set; } = null!;

        /// <summary>
        /// The created time.
        /// </summary>
        [Output("createdTime")]
        public Output<string> CreatedTime { get; private set; } = null!;

        /// <summary>
        /// The integration account RosettaNet ProcessConfiguration properties.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The RosettaNet initiator role settings.
        /// </summary>
        [Output("initiatorRoleSettings")]
        public Output<Outputs.RosettaNetPipRoleSettingsResponseResult> InitiatorRoleSettings { get; private set; } = null!;

        /// <summary>
        /// The resource location.
        /// </summary>
        [Output("location")]
        public Output<string?> Location { get; private set; } = null!;

        /// <summary>
        /// The metadata.
        /// </summary>
        [Output("metadata")]
        public Output<ImmutableDictionary<string, string>?> Metadata { get; private set; } = null!;

        /// <summary>
        /// Gets the resource name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The integration account RosettaNet process code.
        /// </summary>
        [Output("processCode")]
        public Output<string> ProcessCode { get; private set; } = null!;

        /// <summary>
        /// The integration account RosettaNet process name.
        /// </summary>
        [Output("processName")]
        public Output<string> ProcessName { get; private set; } = null!;

        /// <summary>
        /// The integration account RosettaNet process version.
        /// </summary>
        [Output("processVersion")]
        public Output<string> ProcessVersion { get; private set; } = null!;

        /// <summary>
        /// The RosettaNet responder role settings.
        /// </summary>
        [Output("responderRoleSettings")]
        public Output<Outputs.RosettaNetPipRoleSettingsResponseResult> ResponderRoleSettings { get; private set; } = null!;

        /// <summary>
        /// The resource tags.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// Gets the resource type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a RosettaNetProcessConfiguration resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RosettaNetProcessConfiguration(string name, RosettaNetProcessConfigurationArgs args, CustomResourceOptions? options = null)
            : base("azurerm:logic/v20160601:RosettaNetProcessConfiguration", name, args ?? new RosettaNetProcessConfigurationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RosettaNetProcessConfiguration(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:logic/v20160601:RosettaNetProcessConfiguration", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:logic/latest:RosettaNetProcessConfiguration"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing RosettaNetProcessConfiguration resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RosettaNetProcessConfiguration Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new RosettaNetProcessConfiguration(name, id, options);
        }
    }

    public sealed class RosettaNetProcessConfigurationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The RosettaNet process configuration activity settings.
        /// </summary>
        [Input("activitySettings", required: true)]
        public Input<Inputs.RosettaNetPipActivitySettingsArgs> ActivitySettings { get; set; } = null!;

        /// <summary>
        /// The integration account RosettaNet ProcessConfiguration properties.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The RosettaNet initiator role settings.
        /// </summary>
        [Input("initiatorRoleSettings", required: true)]
        public Input<Inputs.RosettaNetPipRoleSettingsArgs> InitiatorRoleSettings { get; set; } = null!;

        /// <summary>
        /// The integration account name.
        /// </summary>
        [Input("integrationAccountName", required: true)]
        public Input<string> IntegrationAccountName { get; set; } = null!;

        /// <summary>
        /// The resource location.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        [Input("metadata")]
        private InputMap<string>? _metadata;

        /// <summary>
        /// The metadata.
        /// </summary>
        public InputMap<string> Metadata
        {
            get => _metadata ?? (_metadata = new InputMap<string>());
            set => _metadata = value;
        }

        /// <summary>
        /// The integration account RosettaNet process code.
        /// </summary>
        [Input("processCode", required: true)]
        public Input<string> ProcessCode { get; set; } = null!;

        /// <summary>
        /// The integration account RosettaNet process name.
        /// </summary>
        [Input("processName", required: true)]
        public Input<string> ProcessName { get; set; } = null!;

        /// <summary>
        /// The integration account RosettaNet process version.
        /// </summary>
        [Input("processVersion", required: true)]
        public Input<string> ProcessVersion { get; set; } = null!;

        /// <summary>
        /// The resource group name.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The RosettaNet responder role settings.
        /// </summary>
        [Input("responderRoleSettings", required: true)]
        public Input<Inputs.RosettaNetPipRoleSettingsArgs> ResponderRoleSettings { get; set; } = null!;

        /// <summary>
        /// The integration account RosettaNet ProcessConfiguration name.
        /// </summary>
        [Input("rosettaNetProcessConfigurationName", required: true)]
        public Input<string> RosettaNetProcessConfigurationName { get; set; } = null!;

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

        public RosettaNetProcessConfigurationArgs()
        {
        }
    }
}
