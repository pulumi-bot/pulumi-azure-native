// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.BotService.V20180712.Outputs
{

    [OutputType]
    public sealed class ChannelResponseResult
    {
        /// <summary>
        /// The channel name
        /// </summary>
        public readonly string ChannelName;

        [OutputConstructor]
        private ChannelResponseResult(string channelName)
        {
            ChannelName = channelName;
        }
    }
}
