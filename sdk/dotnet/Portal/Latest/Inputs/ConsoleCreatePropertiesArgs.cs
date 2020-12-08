// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Portal.Latest.Inputs
{

    /// <summary>
    /// Cloud shell properties for creating a console.
    /// </summary>
    public sealed class ConsoleCreatePropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The operating system type of the cloud shell.
        /// </summary>
        [Input("osType", required: true)]
        public InputUnion<string, Pulumi.AzureNextGen.Portal.Latest.OsType> OsType { get; set; } = null!;

        /// <summary>
        /// Provisioning state of the console.
        /// </summary>
        [Input("provisioningState")]
        public InputUnion<string, Pulumi.AzureNextGen.Portal.Latest.ProvisioningState>? ProvisioningState { get; set; }

        /// <summary>
        /// Uri of the console.
        /// </summary>
        [Input("uri")]
        public Input<string>? Uri { get; set; }

        public ConsoleCreatePropertiesArgs()
        {
        }
    }
}
