// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.MachineLearning.Latest.Inputs
{

    /// <summary>
    /// Asset output port
    /// </summary>
    public sealed class OutputPortArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Port data type.
        /// </summary>
        [Input("type")]
        public InputUnion<string, Pulumi.AzureNextGen.MachineLearning.Latest.OutputPortType>? Type { get; set; }

        public OutputPortArgs()
        {
        }
    }
}
