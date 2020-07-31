// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Authorization.V20150101
{
    public static class GetManagementLockAtResourceGroupLevel
    {
        public static Task<GetManagementLockAtResourceGroupLevelResult> InvokeAsync(GetManagementLockAtResourceGroupLevelArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetManagementLockAtResourceGroupLevelResult>("azurerm:authorization/v20150101:getManagementLockAtResourceGroupLevel", args ?? new GetManagementLockAtResourceGroupLevelArgs(), options.WithVersion());
    }


    public sealed class GetManagementLockAtResourceGroupLevelArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The lock name.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The resource group name.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetManagementLockAtResourceGroupLevelArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetManagementLockAtResourceGroupLevelResult
    {
        /// <summary>
        /// The name of the lock.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The properties of the lock.
        /// </summary>
        public readonly Outputs.ManagementLockPropertiesResponseResult Properties;
        /// <summary>
        /// The type of the lock.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetManagementLockAtResourceGroupLevelResult(
            string? name,

            Outputs.ManagementLockPropertiesResponseResult properties,

            string type)
        {
            Name = name;
            Properties = properties;
            Type = type;
        }
    }
}