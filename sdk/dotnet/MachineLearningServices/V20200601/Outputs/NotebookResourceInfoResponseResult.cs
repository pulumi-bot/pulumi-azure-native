// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.MachineLearningServices.V20200601.Outputs
{

    [OutputType]
    public sealed class NotebookResourceInfoResponseResult
    {
        public readonly string? Fqdn;
        /// <summary>
        /// The error that occurs when preparing notebook.
        /// </summary>
        public readonly Outputs.NotebookPreparationErrorResponseResult? NotebookPreparationError;
        /// <summary>
        /// the data plane resourceId that used to initialize notebook component
        /// </summary>
        public readonly string? ResourceId;

        [OutputConstructor]
        private NotebookResourceInfoResponseResult(
            string? fqdn,

            Outputs.NotebookPreparationErrorResponseResult? notebookPreparationError,

            string? resourceId)
        {
            Fqdn = fqdn;
            NotebookPreparationError = notebookPreparationError;
            ResourceId = resourceId;
        }
    }
}