// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DataFactory.V20170901Preview.Outputs
{

    [OutputType]
    public sealed class ActivityDependencyResponseResult
    {
        /// <summary>
        /// Activity name.
        /// </summary>
        public readonly string Activity;
        /// <summary>
        /// Match-Condition for the dependency.
        /// </summary>
        public readonly ImmutableArray<string> DependencyConditions;

        [OutputConstructor]
        private ActivityDependencyResponseResult(
            string activity,

            ImmutableArray<string> dependencyConditions)
        {
            Activity = activity;
            DependencyConditions = dependencyConditions;
        }
    }
}
