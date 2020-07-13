// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Logic.Outputs
{

    [OutputType]
    public sealed class ApiResourceMetadataResponseResult
    {
        /// <summary>
        /// The api type.
        /// </summary>
        public readonly string? ApiType;
        /// <summary>
        /// The brand color.
        /// </summary>
        public readonly string? BrandColor;
        /// <summary>
        /// The connection type.
        /// </summary>
        public readonly string? ConnectionType;
        /// <summary>
        /// The connector deployment parameters metadata.
        /// </summary>
        public readonly Outputs.ApiDeploymentParameterMetadataSetResponseResult? DeploymentParameters;
        /// <summary>
        /// The hide key.
        /// </summary>
        public readonly string? HideKey;
        /// <summary>
        /// The provisioning state.
        /// </summary>
        public readonly string? ProvisioningState;
        /// <summary>
        /// The source.
        /// </summary>
        public readonly string? Source;
        /// <summary>
        /// The tags.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// The WSDL import method.
        /// </summary>
        public readonly string? WsdlImportMethod;
        /// <summary>
        /// The WSDL service.
        /// </summary>
        public readonly Outputs.WsdlServiceResponseResult? WsdlService;

        [OutputConstructor]
        private ApiResourceMetadataResponseResult(
            string? ApiType,

            string? brandColor,

            string? connectionType,

            Outputs.ApiDeploymentParameterMetadataSetResponseResult? deploymentParameters,

            string? hideKey,

            string? provisioningState,

            string? source,

            ImmutableDictionary<string, string>? tags,

            string? wsdlImportMethod,

            Outputs.WsdlServiceResponseResult? wsdlService)
        {
            this.ApiType = ApiType;
            BrandColor = brandColor;
            ConnectionType = connectionType;
            DeploymentParameters = deploymentParameters;
            HideKey = hideKey;
            ProvisioningState = provisioningState;
            Source = source;
            Tags = tags;
            WsdlImportMethod = wsdlImportMethod;
            WsdlService = wsdlService;
        }
    }
}