// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ServiceFabric.V20191101Preview.Outputs
{

    [OutputType]
    public sealed class ApplicationTypeVersionsCleanupPolicyResponseResult
    {
        /// <summary>
        /// Number of unused versions per application type to keep.
        /// </summary>
        public readonly int MaxUnusedVersionsToKeep;

        [OutputConstructor]
        private ApplicationTypeVersionsCleanupPolicyResponseResult(int maxUnusedVersionsToKeep)
        {
            MaxUnusedVersionsToKeep = maxUnusedVersionsToKeep;
        }
    }
}
