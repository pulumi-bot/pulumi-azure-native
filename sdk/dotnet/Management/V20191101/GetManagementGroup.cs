// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Management.V20191101
{
    public static class GetManagementGroup
    {
        public static Task<GetManagementGroupResult> InvokeAsync(GetManagementGroupArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetManagementGroupResult>("azurerm:management/v20191101:getManagementGroup", args ?? new GetManagementGroupArgs(), options.WithVersion());
    }


    public sealed class GetManagementGroupArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Management Group ID.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetManagementGroupArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetManagementGroupResult
    {
        /// <summary>
        /// The name of the management group. For example, 00000000-0000-0000-0000-000000000000
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The generic properties of a management group.
        /// </summary>
        public readonly Outputs.ManagementGroupPropertiesResponseResult Properties;
        /// <summary>
        /// The type of the resource.  For example, Microsoft.Management/managementGroups
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetManagementGroupResult(
            string name,

            Outputs.ManagementGroupPropertiesResponseResult properties,

            string type)
        {
            Name = name;
            Properties = properties;
            Type = type;
        }
    }
}