// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Compute.Outputs
{

    [OutputType]
    public sealed class BootDiagnosticsInstanceViewResponseResult
    {
        /// <summary>
        /// The console screenshot blob URI. &lt;br&gt;&lt;br&gt;NOTE: This will **not** be set if boot diagnostics is currently enabled with managed storage.
        /// </summary>
        public readonly string ConsoleScreenshotBlobUri;
        /// <summary>
        /// The serial console log blob Uri. &lt;br&gt;&lt;br&gt;NOTE: This will **not** be set if boot diagnostics is currently enabled with managed storage.
        /// </summary>
        public readonly string SerialConsoleLogBlobUri;
        /// <summary>
        /// The boot diagnostics status information for the VM. &lt;br&gt;&lt;br&gt; NOTE: It will be set only if there are errors encountered in enabling boot diagnostics.
        /// </summary>
        public readonly Outputs.InstanceViewStatusResponseResult Status;

        [OutputConstructor]
        private BootDiagnosticsInstanceViewResponseResult(
            string consoleScreenshotBlobUri,

            string serialConsoleLogBlobUri,

            Outputs.InstanceViewStatusResponseResult status)
        {
            ConsoleScreenshotBlobUri = consoleScreenshotBlobUri;
            SerialConsoleLogBlobUri = serialConsoleLogBlobUri;
            Status = status;
        }
    }
}