// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Insights.Latest
{
    /// <summary>
    /// The metric alert resource.
    /// 
    /// ## Example Usage
    /// ### Create or update a dynamic alert rule for Multiple Resources
    /// ```csharp
    /// using Pulumi;
    /// using AzureRM = Pulumi.AzureRM;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var metricAlert = new AzureRM.Insights.Latest.MetricAlert("metricAlert", new AzureRM.Insights.Latest.MetricAlertArgs
    ///         {
    ///             Actions = 
    ///             {
    ///                 new AzureRM.Insights.Latest.Inputs.MetricAlertActionArgs
    ///                 {
    ///                     ActionGroupId = "/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/gigtest/providers/microsoft.insights/notificationgroups/group2",
    ///                     WebHookProperties = 
    ///                     {
    ///                         { "key11", "value11" },
    ///                         { "key12", "value12" },
    ///                     },
    ///                 },
    ///             },
    ///             AutoMitigate = false,
    ///             Criteria = new AzureRM.Insights.Latest.Inputs.MetricAlertCriteriaArgs
    ///             {
    ///                 OdataType = "Microsoft.Azure.Monitor.MultipleResourceMultipleMetricCriteria",
    ///             },
    ///             Description = "This is the description of the rule1",
    ///             Enabled = true,
    ///             EvaluationFrequency = "PT1M",
    ///             Location = "global",
    ///             ResourceGroupName = "gigtest",
    ///             RuleName = "MetricAlertOnMultipleResources",
    ///             Scopes = 
    ///             {
    ///                 "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/gigtest/providers/Microsoft.Compute/virtualMachines/gigwadme1",
    ///                 "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/gigtest/providers/Microsoft.Compute/virtualMachines/gigwadme2",
    ///             },
    ///             Severity = 3,
    ///             Tags = ,
    ///             TargetResourceRegion = "southcentralus",
    ///             TargetResourceType = "Microsoft.Compute/virtualMachines",
    ///             WindowSize = "PT15M",
    ///         });
    ///     }
    /// 
    /// }
    /// 
    /// ```
    /// ### Create or update a dynamic alert rule for Single Resource
    /// ```csharp
    /// using Pulumi;
    /// using AzureRM = Pulumi.AzureRM;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var metricAlert = new AzureRM.Insights.Latest.MetricAlert("metricAlert", new AzureRM.Insights.Latest.MetricAlertArgs
    ///         {
    ///             Actions = 
    ///             {
    ///                 new AzureRM.Insights.Latest.Inputs.MetricAlertActionArgs
    ///                 {
    ///                     ActionGroupId = "/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/gigtest/providers/microsoft.insights/notificationgroups/group2",
    ///                     WebHookProperties = 
    ///                     {
    ///                         { "key11", "value11" },
    ///                         { "key12", "value12" },
    ///                     },
    ///                 },
    ///             },
    ///             AutoMitigate = false,
    ///             Criteria = new AzureRM.Insights.Latest.Inputs.MetricAlertCriteriaArgs
    ///             {
    ///                 OdataType = "Microsoft.Azure.Monitor.MultipleResourceMultipleMetricCriteria",
    ///             },
    ///             Description = "This is the description of the rule1",
    ///             Enabled = true,
    ///             EvaluationFrequency = "PT1M",
    ///             Location = "global",
    ///             ResourceGroupName = "gigtest",
    ///             RuleName = "chiricutin",
    ///             Scopes = 
    ///             {
    ///                 "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/gigtest/providers/Microsoft.Compute/virtualMachines/gigwadme",
    ///             },
    ///             Severity = 3,
    ///             Tags = ,
    ///             TargetResourceRegion = "southcentralus",
    ///             TargetResourceType = "Microsoft.Compute/virtualMachines",
    ///             WindowSize = "PT15M",
    ///         });
    ///     }
    /// 
    /// }
    /// 
    /// ```
    /// ### Create or update a web test alert rule
    /// ```csharp
    /// using Pulumi;
    /// using AzureRM = Pulumi.AzureRM;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var metricAlert = new AzureRM.Insights.Latest.MetricAlert("metricAlert", new AzureRM.Insights.Latest.MetricAlertArgs
    ///         {
    ///             Actions = {},
    ///             Criteria = new AzureRM.Insights.Latest.Inputs.MetricAlertCriteriaArgs
    ///             {
    ///                 OdataType = "Microsoft.Azure.Monitor.WebtestLocationAvailabilityCriteria",
    ///             },
    ///             Description = "Automatically created alert rule for availability test \"component-example\" a",
    ///             Enabled = true,
    ///             EvaluationFrequency = "PT1M",
    ///             Location = "global",
    ///             ResourceGroupName = "rg-example",
    ///             RuleName = "webtest-name-example",
    ///             Scopes = 
    ///             {
    ///                 "/subscriptions/12345678-1234-1234-1234-123456789101/resourcegroups/rg-example/providers/microsoft.insights/webtests/component-example",
    ///                 "/subscriptions/12345678-1234-1234-1234-123456789101/resourcegroups/rg-example/providers/microsoft.insights/components/webtest-name-example",
    ///             },
    ///             Severity = 4,
    ///             Tags = 
    ///             {
    ///                 { "hidden-link:/subscriptions/12345678-1234-1234-1234-123456789101/resourcegroups/rg-example/providers/microsoft.insights/components/webtest-name-example", "Resource" },
    ///                 { "hidden-link:/subscriptions/12345678-1234-1234-1234-123456789101/resourcegroups/rg-example/providers/microsoft.insights/webtests/component-example", "Resource" },
    ///             },
    ///             WindowSize = "PT15M",
    ///         });
    ///     }
    /// 
    /// }
    /// 
    /// ```
    /// ### Create or update an alert rule for Multiple Resource
    /// ```csharp
    /// using Pulumi;
    /// using AzureRM = Pulumi.AzureRM;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var metricAlert = new AzureRM.Insights.Latest.MetricAlert("metricAlert", new AzureRM.Insights.Latest.MetricAlertArgs
    ///         {
    ///             Actions = 
    ///             {
    ///                 new AzureRM.Insights.Latest.Inputs.MetricAlertActionArgs
    ///                 {
    ///                     ActionGroupId = "/subscriptions/14ddf0c5-77c5-4b53-84f6-e1fa43ad68f7/resourcegroups/gigtest/providers/microsoft.insights/notificationgroups/group2",
    ///                     WebHookProperties = 
    ///                     {
    ///                         { "key11", "value11" },
    ///                         { "key12", "value12" },
    ///                     },
    ///                 },
    ///             },
    ///             AutoMitigate = false,
    ///             Criteria = new AzureRM.Insights.Latest.Inputs.MetricAlertCriteriaArgs
    ///             {
    ///                 OdataType = "Microsoft.Azure.Monitor.MultipleResourceMultipleMetricCriteria",
    ///             },
    ///             Description = "This is the description of the rule1",
    ///             Enabled = true,
    ///             EvaluationFrequency = "PT1M",
    ///             Location = "global",
    ///             ResourceGroupName = "gigtest",
    ///             RuleName = "MetricAlertOnMultipleResources",
    ///             Scopes = 
    ///             {
    ///                 "/subscriptions/14ddf0c5-77c5-4b53-84f6-e1fa43ad68f7/resourceGroups/gigtest/providers/Microsoft.Compute/virtualMachines/gigwadme1",
    ///                 "/subscriptions/14ddf0c5-77c5-4b53-84f6-e1fa43ad68f7/resourceGroups/gigtest/providers/Microsoft.Compute/virtualMachines/gigwadme2",
    ///             },
    ///             Severity = 3,
    ///             Tags = ,
    ///             TargetResourceRegion = "southcentralus",
    ///             TargetResourceType = "Microsoft.Compute/virtualMachines",
    ///             WindowSize = "PT15M",
    ///         });
    ///     }
    /// 
    /// }
    /// 
    /// ```
    /// ### Create or update an alert rule for Single Resource
    /// ```csharp
    /// using Pulumi;
    /// using AzureRM = Pulumi.AzureRM;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var metricAlert = new AzureRM.Insights.Latest.MetricAlert("metricAlert", new AzureRM.Insights.Latest.MetricAlertArgs
    ///         {
    ///             Actions = 
    ///             {
    ///                 new AzureRM.Insights.Latest.Inputs.MetricAlertActionArgs
    ///                 {
    ///                     ActionGroupId = "/subscriptions/14ddf0c5-77c5-4b53-84f6-e1fa43ad68f7/resourcegroups/gigtest/providers/microsoft.insights/notificationgroups/group2",
    ///                     WebHookProperties = 
    ///                     {
    ///                         { "key11", "value11" },
    ///                         { "key12", "value12" },
    ///                     },
    ///                 },
    ///             },
    ///             AutoMitigate = false,
    ///             Criteria = new AzureRM.Insights.Latest.Inputs.MetricAlertCriteriaArgs
    ///             {
    ///                 OdataType = "Microsoft.Azure.Monitor.SingleResourceMultipleMetricCriteria",
    ///             },
    ///             Description = "This is the description of the rule1",
    ///             Enabled = true,
    ///             EvaluationFrequency = "Pt1m",
    ///             Location = "global",
    ///             ResourceGroupName = "gigtest",
    ///             RuleName = "chiricutin",
    ///             Scopes = 
    ///             {
    ///                 "/subscriptions/14ddf0c5-77c5-4b53-84f6-e1fa43ad68f7/resourceGroups/gigtest/providers/Microsoft.Compute/virtualMachines/gigwadme",
    ///             },
    ///             Severity = 3,
    ///             Tags = ,
    ///             WindowSize = "Pt15m",
    ///         });
    ///     }
    /// 
    /// }
    /// 
    /// ```
    /// ### Create or update an alert rule on Resource group(s)
    /// ```csharp
    /// using Pulumi;
    /// using AzureRM = Pulumi.AzureRM;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var metricAlert = new AzureRM.Insights.Latest.MetricAlert("metricAlert", new AzureRM.Insights.Latest.MetricAlertArgs
    ///         {
    ///             Actions = 
    ///             {
    ///                 new AzureRM.Insights.Latest.Inputs.MetricAlertActionArgs
    ///                 {
    ///                     ActionGroupId = "/subscriptions/14ddf0c5-77c5-4b53-84f6-e1fa43ad68f7/resourcegroups/gigtest/providers/microsoft.insights/notificationgroups/group2",
    ///                     WebHookProperties = 
    ///                     {
    ///                         { "key11", "value11" },
    ///                         { "key12", "value12" },
    ///                     },
    ///                 },
    ///             },
    ///             AutoMitigate = false,
    ///             Criteria = new AzureRM.Insights.Latest.Inputs.MetricAlertCriteriaArgs
    ///             {
    ///                 OdataType = "Microsoft.Azure.Monitor.MultipleResourceMultipleMetricCriteria",
    ///             },
    ///             Description = "This is the description of the rule1",
    ///             Enabled = true,
    ///             EvaluationFrequency = "PT1M",
    ///             Location = "global",
    ///             ResourceGroupName = "gigtest1",
    ///             RuleName = "MetricAlertAtResourceGroupLevel",
    ///             Scopes = 
    ///             {
    ///                 "/subscriptions/14ddf0c5-77c5-4b53-84f6-e1fa43ad68f7/resourceGroups/gigtest1",
    ///                 "/subscriptions/14ddf0c5-77c5-4b53-84f6-e1fa43ad68f7/resourceGroups/gigtest2",
    ///             },
    ///             Severity = 3,
    ///             Tags = ,
    ///             TargetResourceRegion = "southcentralus",
    ///             TargetResourceType = "Microsoft.Compute/virtualMachines",
    ///             WindowSize = "PT15M",
    ///         });
    ///     }
    /// 
    /// }
    /// 
    /// ```
    /// ### Create or update an alert rule on Subscription
    /// ```csharp
    /// using Pulumi;
    /// using AzureRM = Pulumi.AzureRM;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var metricAlert = new AzureRM.Insights.Latest.MetricAlert("metricAlert", new AzureRM.Insights.Latest.MetricAlertArgs
    ///         {
    ///             Actions = 
    ///             {
    ///                 new AzureRM.Insights.Latest.Inputs.MetricAlertActionArgs
    ///                 {
    ///                     ActionGroupId = "/subscriptions/14ddf0c5-77c5-4b53-84f6-e1fa43ad68f7/resourcegroups/gigtest/providers/microsoft.insights/notificationgroups/group2",
    ///                     WebHookProperties = 
    ///                     {
    ///                         { "key11", "value11" },
    ///                         { "key12", "value12" },
    ///                     },
    ///                 },
    ///             },
    ///             AutoMitigate = false,
    ///             Criteria = new AzureRM.Insights.Latest.Inputs.MetricAlertCriteriaArgs
    ///             {
    ///                 OdataType = "Microsoft.Azure.Monitor.MultipleResourceMultipleMetricCriteria",
    ///             },
    ///             Description = "This is the description of the rule1",
    ///             Enabled = true,
    ///             EvaluationFrequency = "PT1M",
    ///             Location = "global",
    ///             ResourceGroupName = "gigtest",
    ///             RuleName = "MetricAlertAtSubscriptionLevel",
    ///             Scopes = 
    ///             {
    ///                 "/subscriptions/14ddf0c5-77c5-4b53-84f6-e1fa43ad68f7",
    ///             },
    ///             Severity = 3,
    ///             Tags = ,
    ///             TargetResourceRegion = "southcentralus",
    ///             TargetResourceType = "Microsoft.Compute/virtualMachines",
    ///             WindowSize = "PT15M",
    ///         });
    ///     }
    /// 
    /// }
    /// 
    /// ```
    /// </summary>
    public partial class MetricAlert : Pulumi.CustomResource
    {
        /// <summary>
        /// the array of actions that are performed when the alert rule becomes active, and when an alert condition is resolved.
        /// </summary>
        [Output("actions")]
        public Output<ImmutableArray<Outputs.MetricAlertActionResponseResult>> Actions { get; private set; } = null!;

        /// <summary>
        /// the flag that indicates whether the alert should be auto resolved or not. The default is true.
        /// </summary>
        [Output("autoMitigate")]
        public Output<bool?> AutoMitigate { get; private set; } = null!;

        /// <summary>
        /// defines the specific alert criteria information.
        /// </summary>
        [Output("criteria")]
        public Output<Outputs.MetricAlertCriteriaResponseResult> Criteria { get; private set; } = null!;

        /// <summary>
        /// the description of the metric alert that will be included in the alert email.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// the flag that indicates whether the metric alert is enabled.
        /// </summary>
        [Output("enabled")]
        public Output<bool> Enabled { get; private set; } = null!;

        /// <summary>
        /// how often the metric alert is evaluated represented in ISO 8601 duration format.
        /// </summary>
        [Output("evaluationFrequency")]
        public Output<string> EvaluationFrequency { get; private set; } = null!;

        /// <summary>
        /// Last time the rule was updated in ISO8601 format.
        /// </summary>
        [Output("lastUpdatedTime")]
        public Output<string> LastUpdatedTime { get; private set; } = null!;

        /// <summary>
        /// Resource location
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// Azure resource name
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// the list of resource id's that this metric alert is scoped to.
        /// </summary>
        [Output("scopes")]
        public Output<ImmutableArray<string>> Scopes { get; private set; } = null!;

        /// <summary>
        /// Alert severity {0, 1, 2, 3, 4}
        /// </summary>
        [Output("severity")]
        public Output<int> Severity { get; private set; } = null!;

        /// <summary>
        /// Resource tags
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        /// <summary>
        /// the region of the target resource(s) on which the alert is created/updated. Mandatory for MultipleResourceMultipleMetricCriteria.
        /// </summary>
        [Output("targetResourceRegion")]
        public Output<string?> TargetResourceRegion { get; private set; } = null!;

        /// <summary>
        /// the resource type of the target resource(s) on which the alert is created/updated. Mandatory for MultipleResourceMultipleMetricCriteria.
        /// </summary>
        [Output("targetResourceType")]
        public Output<string?> TargetResourceType { get; private set; } = null!;

        /// <summary>
        /// Azure resource type
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// the period of time (in ISO 8601 duration format) that is used to monitor alert activity based on the threshold.
        /// </summary>
        [Output("windowSize")]
        public Output<string> WindowSize { get; private set; } = null!;


        /// <summary>
        /// Create a MetricAlert resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MetricAlert(string name, MetricAlertArgs args, CustomResourceOptions? options = null)
            : base("azurerm:insights/latest:MetricAlert", name, args ?? new MetricAlertArgs(), MakeResourceOptions(options, ""))
        {
        }

        private MetricAlert(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("azurerm:insights/latest:MetricAlert", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new Pulumi.Alias { Type = "azurerm:insights/v20180301:MetricAlert"},
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing MetricAlert resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static MetricAlert Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new MetricAlert(name, id, options);
        }
    }

    public sealed class MetricAlertArgs : Pulumi.ResourceArgs
    {
        [Input("actions")]
        private InputList<Inputs.MetricAlertActionArgs>? _actions;

        /// <summary>
        /// the array of actions that are performed when the alert rule becomes active, and when an alert condition is resolved.
        /// </summary>
        public InputList<Inputs.MetricAlertActionArgs> Actions
        {
            get => _actions ?? (_actions = new InputList<Inputs.MetricAlertActionArgs>());
            set => _actions = value;
        }

        /// <summary>
        /// the flag that indicates whether the alert should be auto resolved or not. The default is true.
        /// </summary>
        [Input("autoMitigate")]
        public Input<bool>? AutoMitigate { get; set; }

        /// <summary>
        /// defines the specific alert criteria information.
        /// </summary>
        [Input("criteria", required: true)]
        public Input<Inputs.MetricAlertCriteriaArgs> Criteria { get; set; } = null!;

        /// <summary>
        /// the description of the metric alert that will be included in the alert email.
        /// </summary>
        [Input("description", required: true)]
        public Input<string> Description { get; set; } = null!;

        /// <summary>
        /// the flag that indicates whether the metric alert is enabled.
        /// </summary>
        [Input("enabled", required: true)]
        public Input<bool> Enabled { get; set; } = null!;

        /// <summary>
        /// how often the metric alert is evaluated represented in ISO 8601 duration format.
        /// </summary>
        [Input("evaluationFrequency", required: true)]
        public Input<string> EvaluationFrequency { get; set; } = null!;

        /// <summary>
        /// Resource location
        /// </summary>
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public Input<string> ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the rule.
        /// </summary>
        [Input("ruleName", required: true)]
        public Input<string> RuleName { get; set; } = null!;

        [Input("scopes")]
        private InputList<string>? _scopes;

        /// <summary>
        /// the list of resource id's that this metric alert is scoped to.
        /// </summary>
        public InputList<string> Scopes
        {
            get => _scopes ?? (_scopes = new InputList<string>());
            set => _scopes = value;
        }

        /// <summary>
        /// Alert severity {0, 1, 2, 3, 4}
        /// </summary>
        [Input("severity", required: true)]
        public Input<int> Severity { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// Resource tags
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        /// <summary>
        /// the region of the target resource(s) on which the alert is created/updated. Mandatory for MultipleResourceMultipleMetricCriteria.
        /// </summary>
        [Input("targetResourceRegion")]
        public Input<string>? TargetResourceRegion { get; set; }

        /// <summary>
        /// the resource type of the target resource(s) on which the alert is created/updated. Mandatory for MultipleResourceMultipleMetricCriteria.
        /// </summary>
        [Input("targetResourceType")]
        public Input<string>? TargetResourceType { get; set; }

        /// <summary>
        /// the period of time (in ISO 8601 duration format) that is used to monitor alert activity based on the threshold.
        /// </summary>
        [Input("windowSize", required: true)]
        public Input<string> WindowSize { get; set; } = null!;

        public MetricAlertArgs()
        {
        }
    }
}
