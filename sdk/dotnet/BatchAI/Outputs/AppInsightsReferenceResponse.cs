// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.BatchAI.Outputs
{

    [OutputType]
    public sealed class AppInsightsReferenceResponse
    {
        /// <summary>
        /// Azure Application Insights component resource ID.
        /// </summary>
        public readonly Outputs.ResourceIdResponse Component;
        /// <summary>
        /// Value of the Azure Application Insights instrumentation key.
        /// </summary>
        public readonly string? InstrumentationKey;
        /// <summary>
        /// KeyVault Store and Secret which contains Azure Application Insights instrumentation key. One of instrumentationKey or instrumentationKeySecretReference must be specified.
        /// </summary>
        public readonly Outputs.KeyVaultSecretReferenceResponse? InstrumentationKeySecretReference;

        [OutputConstructor]
        private AppInsightsReferenceResponse(
            Outputs.ResourceIdResponse component,

            string? instrumentationKey,

            Outputs.KeyVaultSecretReferenceResponse? instrumentationKeySecretReference)
        {
            Component = component;
            InstrumentationKey = instrumentationKey;
            InstrumentationKeySecretReference = instrumentationKeySecretReference;
        }
    }
}
