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
    public sealed class WindowsOsInfoResponseResult
    {
        /// <summary>
        /// The state of the Windows OS.
        /// </summary>
        public readonly string? WindowsOsState;

        [OutputConstructor]
        private WindowsOsInfoResponseResult(string? windowsOsState)
        {
            WindowsOsState = windowsOsState;
        }
    }
}
