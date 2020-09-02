// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DataMigration.V20171115Preview.Outputs
{

    [OutputType]
    public sealed class ProjectTaskPropertiesResponseResult
    {
        /// <summary>
        /// Array of errors. This is ignored if submitted.
        /// </summary>
        public readonly ImmutableArray<Outputs.ODataErrorResponseResult> Errors;
        /// <summary>
        /// The state of the task. This is ignored if submitted.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// Task type.
        /// </summary>
        public readonly string TaskType;

        [OutputConstructor]
        private ProjectTaskPropertiesResponseResult(
            ImmutableArray<Outputs.ODataErrorResponseResult> errors,

            string state,

            string taskType)
        {
            Errors = errors;
            State = state;
            TaskType = taskType;
        }
    }
}
