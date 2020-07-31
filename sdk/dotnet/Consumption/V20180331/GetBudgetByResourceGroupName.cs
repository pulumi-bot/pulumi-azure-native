// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Consumption.V20180331
{
    public static class GetBudgetByResourceGroupName
    {
        public static Task<GetBudgetByResourceGroupNameResult> InvokeAsync(GetBudgetByResourceGroupNameArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetBudgetByResourceGroupNameResult>("azurerm:consumption/v20180331:getBudgetByResourceGroupName", args ?? new GetBudgetByResourceGroupNameArgs(), options.WithVersion());
    }


    public sealed class GetBudgetByResourceGroupNameArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Budget Name.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Azure Resource Group Name.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetBudgetByResourceGroupNameArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetBudgetByResourceGroupNameResult
    {
        /// <summary>
        /// eTag of the resource. To handle concurrent update scenario, this field will be used to determine whether the user is updating the latest version or not.
        /// </summary>
        public readonly string? ETag;
        /// <summary>
        /// Resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The properties of the budget.
        /// </summary>
        public readonly Outputs.BudgetPropertiesResponseResult Properties;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetBudgetByResourceGroupNameResult(
            string? eTag,

            string name,

            Outputs.BudgetPropertiesResponseResult properties,

            string type)
        {
            ETag = eTag;
            Name = name;
            Properties = properties;
            Type = type;
        }
    }
}