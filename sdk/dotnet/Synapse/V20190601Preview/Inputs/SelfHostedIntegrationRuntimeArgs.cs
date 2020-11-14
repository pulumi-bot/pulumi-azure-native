// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Synapse.V20190601Preview.Inputs
{

    /// <summary>
    /// Self-hosted integration runtime.
    /// </summary>
    public sealed class SelfHostedIntegrationRuntimeArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Integration runtime description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Linked integration runtime type from data factory
        /// </summary>
        [Input("linkedInfo")]
        public InputUnion<Inputs.LinkedIntegrationRuntimeKeyAuthorizationArgs, Inputs.LinkedIntegrationRuntimeRbacAuthorizationArgs>? LinkedInfo { get; set; }

        /// <summary>
        /// Type of integration runtime.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public SelfHostedIntegrationRuntimeArgs()
        {
        }
    }
}
