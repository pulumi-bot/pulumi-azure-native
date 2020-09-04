// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.EventGrid.V20200101Preview
{
    public static class GetDomain
    {
        public static Task<GetDomainResult> InvokeAsync(GetDomainArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDomainResult>("azurerm:eventgrid/v20200101preview:getDomain", args ?? new GetDomainArgs(), options.WithVersion());
    }


    public sealed class GetDomainArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Name of the domain
        /// </summary>
        [Input("domainName", required: true)]
        public string DomainName { get; set; } = null!;

        /// <summary>
        /// The name of the resource group within the user's subscription.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        public GetDomainArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetDomainResult
    {
        /// <summary>
        /// Endpoint for the domain.
        /// </summary>
        public readonly string Endpoint;
        /// <summary>
        /// This determines the format that Event Grid should expect for incoming events published to the domain.
        /// </summary>
        public readonly string? InputSchema;
        /// <summary>
        /// Information about the InputSchemaMapping which specified the info about mapping event payload.
        /// </summary>
        public readonly Outputs.InputSchemaMappingResponseResult? InputSchemaMapping;
        /// <summary>
        /// Location of the resource
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// Metric resource id for the domain.
        /// </summary>
        public readonly string MetricResourceId;
        /// <summary>
        /// Name of the resource
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Provisioning state of the domain.
        /// </summary>
        public readonly string ProvisioningState;
        /// <summary>
        /// Tags of the resource
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// Type of the resource
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetDomainResult(
            string endpoint,

            string? inputSchema,

            Outputs.InputSchemaMappingResponseResult? inputSchemaMapping,

            string location,

            string metricResourceId,

            string name,

            string provisioningState,

            ImmutableDictionary<string, string>? tags,

            string type)
        {
            Endpoint = endpoint;
            InputSchema = inputSchema;
            InputSchemaMapping = inputSchemaMapping;
            Location = location;
            MetricResourceId = metricResourceId;
            Name = name;
            ProvisioningState = provisioningState;
            Tags = tags;
            Type = type;
        }
    }
}
