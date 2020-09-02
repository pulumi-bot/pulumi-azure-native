// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Solutions.V20160901Preview.Outputs
{

    [OutputType]
    public sealed class ApplianceArtifactResponseResult
    {
        /// <summary>
        /// The appliance artifact name.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The appliance artifact type.
        /// </summary>
        public readonly string? Type;
        /// <summary>
        /// The appliance artifact blob uri.
        /// </summary>
        public readonly string? Uri;

        [OutputConstructor]
        private ApplianceArtifactResponseResult(
            string? name,

            string? type,

            string? uri)
        {
            Name = name;
            Type = type;
            Uri = uri;
        }
    }
}
