// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ApiManagement.V20191201Preview
{
    public static class GetApiDiagnostic
    {
        public static Task<GetApiDiagnosticResult> InvokeAsync(GetApiDiagnosticArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetApiDiagnosticResult>("azurerm:apimanagement/v20191201preview:getApiDiagnostic", args ?? new GetApiDiagnosticArgs(), options.WithVersion());
    }


    public sealed class GetApiDiagnosticArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// API identifier. Must be unique in the current API Management service instance.
        /// </summary>
        [Input("apiId", required: true)]
        public string ApiId { get; set; } = null!;

        /// <summary>
        /// Diagnostic identifier. Must be unique in the current API Management service instance.
        /// </summary>
        [Input("diagnosticId", required: true)]
        public string DiagnosticId { get; set; } = null!;

        /// <summary>
        /// The name of the resource group.
        /// </summary>
        [Input("resourceGroupName", required: true)]
        public string ResourceGroupName { get; set; } = null!;

        /// <summary>
        /// The name of the API Management service.
        /// </summary>
        [Input("serviceName", required: true)]
        public string ServiceName { get; set; } = null!;

        public GetApiDiagnosticArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetApiDiagnosticResult
    {
        /// <summary>
        /// Specifies for what type of messages sampling settings should not apply.
        /// </summary>
        public readonly string? AlwaysLog;
        /// <summary>
        /// Diagnostic settings for incoming/outgoing HTTP messages to the Backend
        /// </summary>
        public readonly Outputs.PipelineDiagnosticSettingsResponseResult? Backend;
        /// <summary>
        /// Diagnostic settings for incoming/outgoing HTTP messages to the Gateway.
        /// </summary>
        public readonly Outputs.PipelineDiagnosticSettingsResponseResult? Frontend;
        /// <summary>
        /// Sets correlation protocol to use for Application Insights diagnostics.
        /// </summary>
        public readonly string? HttpCorrelationProtocol;
        /// <summary>
        /// Log the ClientIP. Default is false.
        /// </summary>
        public readonly bool? LogClientIp;
        /// <summary>
        /// Resource Id of a target logger.
        /// </summary>
        public readonly string LoggerId;
        /// <summary>
        /// Resource name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Sampling settings for Diagnostic.
        /// </summary>
        public readonly Outputs.SamplingSettingsResponseResult? Sampling;
        /// <summary>
        /// Resource type for API Management resource.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// The verbosity level applied to traces emitted by trace policies.
        /// </summary>
        public readonly string? Verbosity;

        [OutputConstructor]
        private GetApiDiagnosticResult(
            string? alwaysLog,

            Outputs.PipelineDiagnosticSettingsResponseResult? backend,

            Outputs.PipelineDiagnosticSettingsResponseResult? frontend,

            string? httpCorrelationProtocol,

            bool? logClientIp,

            string loggerId,

            string name,

            Outputs.SamplingSettingsResponseResult? sampling,

            string type,

            string? verbosity)
        {
            AlwaysLog = alwaysLog;
            Backend = backend;
            Frontend = frontend;
            HttpCorrelationProtocol = httpCorrelationProtocol;
            LogClientIp = logClientIp;
            LoggerId = loggerId;
            Name = name;
            Sampling = sampling;
            Type = type;
            Verbosity = verbosity;
        }
    }
}
