// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Automation
{
    public static class GetAutomationAccountSchedule
    {
        public static Task<GetAutomationAccountScheduleResult> InvokeAsync(GetAutomationAccountScheduleArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAutomationAccountScheduleResult>("azurerm:automation:getAutomationAccountSchedule", args ?? new GetAutomationAccountScheduleArgs(), options.WithVersion());
    }


    public sealed class GetAutomationAccountScheduleArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the automation account.
        /// </summary>
        [Input("automationAccountName", required: true)]
        public string AutomationAccountName { get; set; } = null!;

        /// <summary>
        /// The schedule name.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Name of an Azure Resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetAutomationAccountScheduleArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAutomationAccountScheduleResult
    {
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Gets or sets the properties of the schedule.
        /// </summary>
        public readonly Outputs.SchedulePropertiesResponseResult Properties;
        /// <summary>
        /// The type of the resource.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetAutomationAccountScheduleResult(
            string name,

            Outputs.SchedulePropertiesResponseResult properties,

            string type)
        {
            Name = name;
            Properties = properties;
            Type = type;
        }
    }
}