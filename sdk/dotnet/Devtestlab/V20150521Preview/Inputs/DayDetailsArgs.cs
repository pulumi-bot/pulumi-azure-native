// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DevTestLab.V20150521Preview.Inputs
{

    /// <summary>
    /// Properties of a daily schedule.
    /// </summary>
    public sealed class DayDetailsArgs : Pulumi.ResourceArgs
    {
        [Input("time")]
        public Input<string>? Time { get; set; }

        public DayDetailsArgs()
        {
        }
    }
}
