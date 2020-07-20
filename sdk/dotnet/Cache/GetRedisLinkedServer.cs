// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Cache
{
    public static class GetRedisLinkedServer
    {
        public static Task<GetRedisLinkedServerResult> InvokeAsync(GetRedisLinkedServerArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRedisLinkedServerResult>("azurerm:cache:getRedisLinkedServer", args ?? new GetRedisLinkedServerArgs(), options.WithVersion());
    }


    public sealed class GetRedisLinkedServerArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the linked server.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetRedisLinkedServerArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetRedisLinkedServerResult
    {
        /// <summary>
        /// Resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Properties of the linked server.
        /// </summary>
        public readonly Outputs.RedisLinkedServerPropertiesResponseResult Properties;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetRedisLinkedServerResult(
            string name,

            Outputs.RedisLinkedServerPropertiesResponseResult properties,

            string type)
        {
            Name = name;
            Properties = properties;
            Type = type;
        }
    }
}
