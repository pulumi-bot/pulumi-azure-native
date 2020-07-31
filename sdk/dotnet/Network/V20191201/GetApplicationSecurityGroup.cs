// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20191201
{
    public static class GetApplicationSecurityGroup
    {
        public static Task<GetApplicationSecurityGroupResult> InvokeAsync(GetApplicationSecurityGroupArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetApplicationSecurityGroupResult>("azurerm:network/v20191201:getApplicationSecurityGroup", args ?? new GetApplicationSecurityGroupArgs(), options.WithVersion());
    }


    public sealed class GetApplicationSecurityGroupArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the application security group.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetApplicationSecurityGroupArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetApplicationSecurityGroupResult
    {
        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// Resource location.
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// Resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Properties of the application security group.
        /// </summary>
        public readonly Outputs.ApplicationSecurityGroupPropertiesFormatResponseResult Properties;
        /// <summary>
        /// Resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Resource type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetApplicationSecurityGroupResult(
            string etag,

            string? location,

            string name,

            Outputs.ApplicationSecurityGroupPropertiesFormatResponseResult properties,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            Etag = etag;
            Location = location;
            Name = name;
            Properties = properties;
            Tags = tags;
            Type = type;
        }
    }
}