// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Insights.Outputs
{

    [OutputType]
    public sealed class PrivateLinkScopedResourceResponse
    {
        /// <summary>
        /// The full resource Id of the private link scope resource.
        /// </summary>
        public readonly string? ResourceId;
        /// <summary>
        /// The private link scope unique Identifier.
        /// </summary>
        public readonly string? ScopeId;

        [OutputConstructor]
        private PrivateLinkScopedResourceResponse(
            string? ResourceId,

            string? ScopeId)
        {
            this.ResourceId = ResourceId;
            this.ScopeId = ScopeId;
        }
    }
}