// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.HybridCompute.V20200815Preview.Inputs
{

    public sealed class MachineIdentityArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The identity type.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public MachineIdentityArgs()
        {
        }
    }
}
