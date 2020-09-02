// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ServiceFabric.V20191101Preview
{
    public static class GetApplicationType
    {
        public static Task<GetApplicationTypeResult> InvokeAsync(GetApplicationTypeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetApplicationTypeResult>("azurerm:servicefabric/v20191101preview:getApplicationType", args ?? new GetApplicationTypeArgs(), options.WithVersion());
    }


    public sealed class GetApplicationTypeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the application type name resource.
        /// </summary>
        [Input("applicationTypeName", required: true)]
        public string ApplicationTypeName { get; set; } = null!;

        /// <summary>
        /// The name of the cluster resource.
        /// </summary>
        [Input("clusterName", required: true)]
        public string ClusterName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetApplicationTypeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetApplicationTypeResult
    {
        /// <summary>
        /// Azure resource etag.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// It will be deprecated in New API, resource location depends on the parent resource.
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// Azure resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The current deployment or provisioning state, which only appears in the response.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Azure resource tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Azure resource type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetApplicationTypeResult(
            string etag,

            string? location,

            string name,

            string provisioningState,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            Etag = etag;
            Location = location;
            Name = name;
            ProvisioningState = provisioningState;
            Tags = tags;
            Type = type;
        }
    }
}
