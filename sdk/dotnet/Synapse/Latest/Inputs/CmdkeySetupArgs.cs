// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Synapse.Latest.Inputs
{

    /// <summary>
    /// The custom setup of running cmdkey commands.
    /// </summary>
    public sealed class CmdkeySetupArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The password of data source access.
        /// </summary>
        [Input("password", required: true)]
        public Input<Inputs.SecureStringArgs> Password { get; set; } = null!;

        /// <summary>
        /// The server name of data source access.
        /// </summary>
        [Input("targetName", required: true)]
        public Input<object> TargetName { get; set; } = null!;

        /// <summary>
        /// The type of custom setup.
        /// Expected value is 'CmdkeySetup'.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// The user name of data source access.
        /// </summary>
        [Input("userName", required: true)]
        public Input<object> UserName { get; set; } = null!;

        public CmdkeySetupArgs()
        {
        }
    }
}
