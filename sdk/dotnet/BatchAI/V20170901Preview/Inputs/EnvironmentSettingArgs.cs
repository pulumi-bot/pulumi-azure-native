// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.BatchAI.V20170901Preview.Inputs
{

    /// <summary>
    /// A collection of environment variables to set.
    /// </summary>
    public sealed class EnvironmentSettingArgs : Pulumi.ResourceArgs
    {
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("value")]
        public Input<string>? Value { get; set; }

        public EnvironmentSettingArgs()
        {
        }
    }
}
