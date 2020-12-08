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
    /// The parameters to provide for the Enterprise Channel.
    /// </summary>
    public sealed class EnterpriseChannelPropertiesArgs : Pulumi.ResourceArgs
    {
        [Input("nodes", required: true)]
        private InputList<Inputs.EnterpriseChannelNodeArgs>? _nodes;

        /// <summary>
        /// The nodes associated with the Enterprise Channel.
        /// </summary>
        public InputList<Inputs.EnterpriseChannelNodeArgs> Nodes
        {
            get => _nodes ?? (_nodes = new InputList<Inputs.EnterpriseChannelNodeArgs>());
            set => _nodes = value;
        }

        /// <summary>
        /// The current state of the Enterprise Channel.
        /// </summary>
        [Input("state")]
        public InputUnion<string, Pulumi.AzureNextGen.BotService.Latest.EnterpriseChannelState>? State { get; set; }

        public EnterpriseChannelPropertiesArgs()
        {
        }
    }
}
