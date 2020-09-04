// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.AzureData.V20190724Preview
{
    public static class GetDataController
    {
        public static Task<GetDataControllerResult> InvokeAsync(GetDataControllerArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDataControllerResult>("azurerm:azuredata/v20190724preview:getDataController", args ?? new GetDataControllerArgs(), options.WithVersion());
    }


    public sealed class GetDataControllerArgs : Pulumi.InvokeArgs
    {
        [Input("dataControllerName", required: true)]
        public string DataControllerName { get; set; } = null!;

        /// <summary>
        /// The name of the Azure resource group
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetDataControllerArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDataControllerResult
    {
        /// <summary>
        /// The geo-location where the resource lives
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Properties from the on premise data controller
        /// </summary>
        public readonly Outputs.OnPremisePropertyResponseResult OnPremiseProperty;
        /// <summary>
        /// Read only system data
        /// </summary>
        public readonly Outputs.SystemDataResponseResult SystemData;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetDataControllerResult(
            string location,

            string name,

            Outputs.OnPremisePropertyResponseResult onPremiseProperty,

            Outputs.SystemDataResponseResult systemData,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            Location = location;
            Name = name;
            OnPremiseProperty = onPremiseProperty;
            SystemData = systemData;
            Tags = tags;
            Type = type;
        }
    }
}
