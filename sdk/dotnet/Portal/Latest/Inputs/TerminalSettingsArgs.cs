// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Portal.Latest.Inputs
{

    /// <summary>
    /// Settings for terminal appearance.
    /// </summary>
    public sealed class TerminalSettingsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Size of terminal font.
        /// </summary>
        [Input("fontSize")]
        public Input<string>? FontSize { get; set; }

        /// <summary>
        /// Style of terminal font.
        /// </summary>
        [Input("fontStyle")]
        public Input<string>? FontStyle { get; set; }

        public TerminalSettingsArgs()
        {
        }
    }
}