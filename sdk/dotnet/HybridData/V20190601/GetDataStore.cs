// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.HybridData.V20190601
{
    public static class GetDataStore
    {
        public static Task<GetDataStoreResult> InvokeAsync(GetDataStoreArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDataStoreResult>("azurerm:hybriddata/v20190601:getDataStore", args ?? new GetDataStoreArgs(), options.WithVersion());
    }


    public sealed class GetDataStoreArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the DataManager Resource within the specified resource group. DataManager names must be between 3 and 24 characters in length and use any alphanumeric and underscore only
        /// </summary>
        [Input("dataManagerName", required: true)]
        public string DataManagerName { get; set; } = null!;

        /// <summary>
        /// The data store/repository name queried.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The Resource Group Name
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetDataStoreArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDataStoreResult
    {
        /// <summary>
        /// Name of the object.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// DataStore properties.
        /// </summary>
        public readonly Outputs.DataStorePropertiesResponseResult Properties;
        /// <summary>
        /// Type of the object.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetDataStoreResult(
            string name,

            Outputs.DataStorePropertiesResponseResult properties,

            string type)
        {
            Name = name;
            Properties = properties;
            Type = type;
        }
    }
}