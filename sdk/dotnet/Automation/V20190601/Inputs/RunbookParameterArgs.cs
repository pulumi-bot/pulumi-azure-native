// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Automation.V20190601.Inputs
{

    /// <summary>
    /// Definition of the runbook parameter type.
    /// </summary>
    public sealed class RunbookParameterArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Gets or sets the default value of parameter.
        /// </summary>
        [Input("defaultValue")]
        public Input<string>? DefaultValue { get; set; }

        /// <summary>
        /// Gets or sets a Boolean value to indicate whether the parameter is mandatory or not.
        /// </summary>
        [Input("isMandatory")]
        public Input<bool>? IsMandatory { get; set; }

        /// <summary>
        /// Get or sets the position of the parameter.
        /// </summary>
        [Input("position")]
        public Input<int>? Position { get; set; }

        /// <summary>
        /// Gets or sets the type of the parameter.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public RunbookParameterArgs()
        {
        }
    }
}
