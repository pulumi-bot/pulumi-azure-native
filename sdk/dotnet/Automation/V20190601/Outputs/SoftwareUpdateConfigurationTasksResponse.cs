// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Automation.V20190601.Outputs
{

    [OutputType]
    public sealed class SoftwareUpdateConfigurationTasksResponse
    {
        /// <summary>
        /// Post task properties.
        /// </summary>
        public readonly Outputs.TaskPropertiesResponse? PostTask;
        /// <summary>
        /// Pre task properties.
        /// </summary>
        public readonly Outputs.TaskPropertiesResponse? PreTask;

        [OutputConstructor]
        private SoftwareUpdateConfigurationTasksResponse(
            Outputs.TaskPropertiesResponse? postTask,

            Outputs.TaskPropertiesResponse? preTask)
        {
            PostTask = postTask;
            PreTask = preTask;
        }
    }
}
