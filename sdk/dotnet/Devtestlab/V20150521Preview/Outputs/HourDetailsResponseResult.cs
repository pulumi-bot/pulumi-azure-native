// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DevTestLab.V20150521Preview.Outputs
{

    [OutputType]
    public sealed class HourDetailsResponseResult
    {
        /// <summary>
        /// Minutes of the hour the schedule will run.
        /// </summary>
        public readonly int? Minute;

        [OutputConstructor]
        private HourDetailsResponseResult(int? minute)
        {
            Minute = minute;
        }
    }
}
