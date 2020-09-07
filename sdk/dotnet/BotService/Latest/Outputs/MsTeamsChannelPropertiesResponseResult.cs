// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.BotService.Latest.Outputs
{

    [OutputType]
    public sealed class MsTeamsChannelPropertiesResponseResult
    {
        /// <summary>
        /// Webhook for Microsoft Teams channel calls
        /// </summary>
        public readonly string? CallingWebHook;
        /// <summary>
        /// Enable calling for Microsoft Teams channel
        /// </summary>
        public readonly bool? EnableCalling;
        /// <summary>
        /// Whether this channel is enabled for the bot
        /// </summary>
        public readonly bool IsEnabled;

        [OutputConstructor]
        private MsTeamsChannelPropertiesResponseResult(
            string? callingWebHook,

            bool? enableCalling,

            bool isEnabled)
        {
            CallingWebHook = callingWebHook;
            EnableCalling = enableCalling;
            IsEnabled = isEnabled;
        }
    }
}
