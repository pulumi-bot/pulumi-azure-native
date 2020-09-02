// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Migrate.V20191001Preview
{
    public static class GetMoveResource
    {
        public static Task<GetMoveResourceResult> InvokeAsync(GetMoveResourceArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetMoveResourceResult>("azurerm:migrate/v20191001preview:getMoveResource", args ?? new GetMoveResourceArgs(), options.WithVersion());
    }


    public sealed class GetMoveResourceArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The Move Collection Name.
        /// </summary>
        [Input("moveCollectionName", required: true)]
        public string MoveCollectionName { get; set; } = null!;

        /// <summary>
        /// The Move Resource Name.
        /// </summary>
        [Input("moveResourceName", required: true)]
        public string MoveResourceName { get; set; } = null!;

        /// <summary>
        /// The Resource Group Name.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetMoveResourceArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetMoveResourceResult
    {
        /// <summary>
        /// The name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Defines the move resource properties.
        /// </summary>
        public readonly Outputs.MoveResourcePropertiesResponseResult Properties;
        /// <summary>
        /// The type of the resource.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetMoveResourceResult(
            string name,

            Outputs.MoveResourcePropertiesResponseResult properties,

            string type)
        {
            Name = name;
            Properties = properties;
            Type = type;
        }
    }
}
