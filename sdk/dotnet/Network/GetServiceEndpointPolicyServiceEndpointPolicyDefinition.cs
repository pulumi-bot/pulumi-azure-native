// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network
{
    public static class GetServiceEndpointPolicyServiceEndpointPolicyDefinition
    {
        public static Task<GetServiceEndpointPolicyServiceEndpointPolicyDefinitionResult> InvokeAsync(GetServiceEndpointPolicyServiceEndpointPolicyDefinitionArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetServiceEndpointPolicyServiceEndpointPolicyDefinitionResult>("azurerm:network:getServiceEndpointPolicyServiceEndpointPolicyDefinition", args ?? new GetServiceEndpointPolicyServiceEndpointPolicyDefinitionArgs(), options.WithVersion());
    }


    public sealed class GetServiceEndpointPolicyServiceEndpointPolicyDefinitionArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the service endpoint policy definition name.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the service endpoint policy name.
        /// </summary>
        [Input("serviceEndpointPolicyName", required: true)]
        public string ServiceEndpointPolicyName { get; set; } = null!;

        public GetServiceEndpointPolicyServiceEndpointPolicyDefinitionArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetServiceEndpointPolicyServiceEndpointPolicyDefinitionResult
    {
        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// The name of the resource that is unique within a resource group. This name can be used to access the resource.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Properties of the service endpoint policy definition.
        /// </summary>
        public readonly Outputs.ServiceEndpointPolicyDefinitionPropertiesFormatResponseResult Properties;

        [OutputConstructor]
        private GetServiceEndpointPolicyServiceEndpointPolicyDefinitionResult(
            string etag,

            string? name,

            Outputs.ServiceEndpointPolicyDefinitionPropertiesFormatResponseResult properties)
        {
            Etag = etag;
            Name = name;
            Properties = properties;
        }
    }
}