// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DeploymentManager.V20180901Preview.Inputs
{

    /// <summary>
    /// Defines the authentication method and properties to access the artifacts.
    /// </summary>
    public sealed class AuthenticationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The authentication type
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public AuthenticationArgs()
        {
        }
    }
}
