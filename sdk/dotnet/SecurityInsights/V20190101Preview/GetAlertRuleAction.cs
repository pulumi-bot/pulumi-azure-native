// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.SecurityInsights.V20190101Preview
{
    public static class GetAlertRuleAction
    {
        public static Task<GetAlertRuleActionResult> InvokeAsync(GetAlertRuleActionArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAlertRuleActionResult>("azurerm:securityinsights/v20190101preview:getAlertRuleAction", args ?? new GetAlertRuleActionArgs(), options.WithVersion());
    }


    public sealed class GetAlertRuleActionArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Action ID
        /// </summary>
        [Input("actionId", required: true)]
        public string ActionId { get; set; } = null!;

        /// <summary>
        /// The namespace of workspaces resource provider- Microsoft.OperationalInsights.
        /// </summary>
        [Input("operationalInsightsResourceProvider", required: true)]
        public string OperationalInsightsResourceProvider { get; set; } = null!;

        /// <summary>
        /// The name of the resource group within the user's subscription. The name is case insensitive.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// Alert rule ID
        /// </summary>
        [Input("ruleId", required: true)]
        public string RuleId { get; set; } = null!;

        /// <summary>
        /// The name of the workspace.
        /// </summary>
        [Input("workspaceName", required: true)]
        public string WorkspaceName { get; set; } = null!;

        public GetAlertRuleActionArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAlertRuleActionResult
    {
        /// <summary>
        /// Etag of the action.
        /// </summary>
        public readonly string? Etag;
        /// <summary>
        /// Logic App Resource Id, /subscriptions/{my-subscription}/resourceGroups/{my-resource-group}/providers/Microsoft.Logic/workflows/{my-workflow-id}.
        /// </summary>
        public readonly string? LogicAppResourceId;
        /// <summary>
        /// Azure resource name
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Azure resource type
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The name of the logic app's workflow.
        /// </summary>
        public readonly string? WorkflowId;

        [OutputConstructor]
        private GetAlertRuleActionResult(
            string? etag,

            string? logicAppResourceId,

            string name,

            string type,

            string? workflowId)
        {
            Etag = etag;
            LogicAppResourceId = logicAppResourceId;
            Name = name;
            Type = type;
            WorkflowId = workflowId;
        }
    }
}
