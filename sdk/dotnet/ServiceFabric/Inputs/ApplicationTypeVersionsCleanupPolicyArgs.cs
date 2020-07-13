// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ServiceFabric.Inputs
{

    public sealed class ApplicationTypeVersionsCleanupPolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Number of unused versions per application type to keep.
        /// </summary>
        [Input("maxUnusedVersionsToKeep", required: true)]
        public Input<int> MaxUnusedVersionsToKeep { get; set; } = null!;

        public ApplicationTypeVersionsCleanupPolicyArgs()
        {
        }
    }
}