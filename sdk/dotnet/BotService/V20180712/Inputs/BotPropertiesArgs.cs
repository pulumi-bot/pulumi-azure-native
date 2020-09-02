// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.BotService.V20180712.Inputs
{

    /// <summary>
    /// The parameters to provide for the Bot.
    /// </summary>
    public sealed class BotPropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the bot
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The Application Insights key
        /// </summary>
        [Input("developerAppInsightKey")]
        public Input<string>? DeveloperAppInsightKey { get; set; }

        /// <summary>
        /// The Application Insights Api Key
        /// </summary>
        [Input("developerAppInsightsApiKey")]
        public Input<string>? DeveloperAppInsightsApiKey { get; set; }

        /// <summary>
        /// The Application Insights App Id
        /// </summary>
        [Input("developerAppInsightsApplicationId")]
        public Input<string>? DeveloperAppInsightsApplicationId { get; set; }

        /// <summary>
        /// The Name of the bot
        /// </summary>
        [Input("displayName", required: true)]
        public Input<string> DisplayName { get; set; } = null!;

        /// <summary>
        /// The bot's endpoint
        /// </summary>
        [Input("endpoint", required: true)]
        public Input<string> Endpoint { get; set; } = null!;

        /// <summary>
        /// The Icon Url of the bot
        /// </summary>
        [Input("iconUrl")]
        public Input<string>? IconUrl { get; set; }

        [Input("luisAppIds")]
        private InputList<string>? _luisAppIds;

        /// <summary>
        /// Collection of LUIS App Ids
        /// </summary>
        public InputList<string> LuisAppIds
        {
            get => _luisAppIds ?? (_luisAppIds = new InputList<string>());
            set => _luisAppIds = value;
        }

        /// <summary>
        /// The LUIS Key
        /// </summary>
        [Input("luisKey")]
        public Input<string>? LuisKey { get; set; }

        /// <summary>
        /// Microsoft App Id for the bot
        /// </summary>
        [Input("msaAppId", required: true)]
        public Input<string> MsaAppId { get; set; } = null!;

        public BotPropertiesArgs()
        {
        }
    }
}
