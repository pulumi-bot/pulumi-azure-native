// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.BotService.Latest.Inputs
{

    /// <summary>
    /// The properties specific to an Enterprise Channel Node.
    /// </summary>
    public sealed class EnterpriseChannelNodeArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The location of the Enterprise Channel Node.
        /// </summary>
        [Input("azureLocation", required: true)]
        public Input<string> AzureLocation { get; set; } = null!;

        /// <summary>
        /// The sku of the Enterprise Channel Node.
        /// </summary>
        [Input("azureSku", required: true)]
        public Input<string> AzureSku { get; set; } = null!;

        /// <summary>
        /// The name of the Enterprise Channel Node.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The current state of the Enterprise Channel Node.
        /// </summary>
        [Input("state")]
        public InputUnion<string, Pulumi.AzureNextGen.BotService.Latest.EnterpriseChannelNodeState>? State { get; set; }

        public EnterpriseChannelNodeArgs()
        {
        }
    }
}
