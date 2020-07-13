// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Automation.Outputs
{

    [OutputType]
    public sealed class ContentSourceResponseResult
    {
        /// <summary>
        /// Gets or sets the hash.
        /// </summary>
        public readonly Outputs.ContentHashResponseResult? Hash;
        /// <summary>
        /// Gets or sets the content source type.
        /// </summary>
        public readonly string? Type;
        /// <summary>
        /// Gets or sets the value of the content. This is based on the content source type.
        /// </summary>
        public readonly string? Value;
        /// <summary>
        /// Gets or sets the version of the content.
        /// </summary>
        public readonly string? Version;

        [OutputConstructor]
        private ContentSourceResponseResult(
            Outputs.ContentHashResponseResult? hash,

            string? type,

            string? value,

            string? version)
        {
            Hash = hash;
            Type = type;
            Value = value;
            Version = version;
        }
    }
}