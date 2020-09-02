// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Solutions.V20200821Preview.Inputs
{

    /// <summary>
    /// Managed application deployment policy.
    /// </summary>
    public sealed class ApplicationDeploymentPolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The managed application deployment mode.
        /// </summary>
        [Input("deploymentMode", required: true)]
        public Input<string> DeploymentMode { get; set; } = null!;

        public ApplicationDeploymentPolicyArgs()
        {
        }
    }
}
