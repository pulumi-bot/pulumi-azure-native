// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ContainerRegistry
{
    public static class ListRegistryTaskDetails
    {
        public static Task<ListRegistryTaskDetailsResult> InvokeAsync(ListRegistryTaskDetailsArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<ListRegistryTaskDetailsResult>("azurerm:containerregistry:listRegistryTaskDetails", args ?? new ListRegistryTaskDetailsArgs(), options.WithVersion());
    }


    public sealed class ListRegistryTaskDetailsArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the container registry task.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the container registry.
        /// </summary>
        [Input("registryName", required: true)]
        public string RegistryName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group to which the container registry belongs.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public ListRegistryTaskDetailsArgs()
        {
        }
    }


    [OutputType]
    public sealed class ListRegistryTaskDetailsResult
    {
        /// <summary>
        /// Identity for the resource.
        /// </summary>
        public readonly Outputs.IdentityPropertiesResponseResult? Identity;
        /// <summary>
        /// The location of the resource. This cannot be changed after the resource is created.
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// The name of the resource.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The properties of a task.
        /// </summary>
        public readonly Outputs.TaskPropertiesResponseResult Properties;
        /// <summary>
        /// The tags of the resource.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// The type of the resource.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private ListRegistryTaskDetailsResult(
            Outputs.IdentityPropertiesResponseResult? identity,

            string location,

            string name,

            Outputs.TaskPropertiesResponseResult properties,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            Identity = identity;
            Location = location;
            Name = name;
            Properties = properties;
            Tags = tags;
            Type = type;
        }
    }
}