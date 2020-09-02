// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.AppPlatform.V20190501Preview.Outputs
{

    [OutputType]
    public sealed class UserSourceInfoResponseResult
    {
        /// <summary>
        /// Selector for the artifact to be used for the deployment for multi-module projects. This should be
        /// the relative path to the target module/project.
        /// </summary>
        public readonly string? ArtifactSelector;
        /// <summary>
        /// Relative path of the storage which stores the source
        /// </summary>
        public readonly string? RelativePath;
        /// <summary>
        /// Type of the source uploaded
        /// </summary>
        public readonly string? Type;
        /// <summary>
        /// Version of the source
        /// </summary>
        public readonly string? Version;

        [OutputConstructor]
        private UserSourceInfoResponseResult(
            string? artifactSelector,

            string? relativePath,

            string? type,

            string? version)
        {
            ArtifactSelector = artifactSelector;
            RelativePath = relativePath;
            Type = type;
            Version = version;
        }
    }
}
