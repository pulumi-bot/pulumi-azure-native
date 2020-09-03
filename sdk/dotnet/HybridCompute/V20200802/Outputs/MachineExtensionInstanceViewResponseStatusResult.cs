// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.HybridCompute.V20200802.Outputs
{

    [OutputType]
    public sealed class MachineExtensionInstanceViewResponseStatusResult
    {
        /// <summary>
        /// The status code.
        /// </summary>
        public readonly string Code;
        /// <summary>
        /// The short localizable label for the status.
        /// </summary>
        public readonly string DisplayStatus;
        /// <summary>
        /// The level code.
        /// </summary>
        public readonly string Level;
        /// <summary>
        /// The detailed status message, including for alerts and error messages.
        /// </summary>
        public readonly string Message;
        /// <summary>
        /// The time of the status.
        /// </summary>
        public readonly string Time;

        [OutputConstructor]
        private MachineExtensionInstanceViewResponseStatusResult(
            string code,

            string displayStatus,

            string level,

            string message,

            string time)
        {
            Code = code;
            DisplayStatus = displayStatus;
            Level = level;
            Message = message;
            Time = time;
        }
    }
}
