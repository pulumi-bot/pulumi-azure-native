// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.ContainerInstance.V20201101.Outputs
{

    [OutputType]
    public sealed class ContainerExecResponse
    {
        /// <summary>
        /// The commands to execute within the container.
        /// </summary>
        public readonly ImmutableArray<string> Command;

        [OutputConstructor]
        private ContainerExecResponse(ImmutableArray<string> command)
        {
            Command = command;
        }
    }
}
