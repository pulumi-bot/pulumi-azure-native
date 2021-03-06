// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.Synapse.V20201201.Outputs
{

    [OutputType]
    public sealed class SsisEnvironmentResponseResult
    {
        /// <summary>
        /// Metadata description.
        /// </summary>
        public readonly string? Description;
        /// <summary>
        /// Folder id which contains environment.
        /// </summary>
        public readonly double? FolderId;
        /// <summary>
        /// Metadata id.
        /// </summary>
        public readonly double? Id;
        /// <summary>
        /// Metadata name.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Type of metadata.
        /// Expected value is 'Environment'.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Variable in environment
        /// </summary>
        public readonly ImmutableArray<Outputs.SsisVariableResponseResult> Variables;

        [OutputConstructor]
        private SsisEnvironmentResponseResult(
            string? description,

            double? folderId,

            double? id,

            string? name,

            string type,

            ImmutableArray<Outputs.SsisVariableResponseResult> variables)
        {
            Description = description;
            FolderId = folderId;
            Id = id;
            Name = name;
            Type = type;
            Variables = variables;
        }
    }
}
