// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.HealthBot.V20201208Preview.Outputs
{

    [OutputType]
    public sealed class HealthBotPropertiesResponse
    {
        /// <summary>
        /// The link.
        /// </summary>
        public readonly string BotManagementPortalLink;
        /// <summary>
        /// The provisioning state of the Healthcare bot resource.
        /// </summary>
        public readonly string ProvisioningState;

        [OutputConstructor]
        private HealthBotPropertiesResponse(
            string botManagementPortalLink,

            string provisioningState)
        {
            BotManagementPortalLink = botManagementPortalLink;
            ProvisioningState = provisioningState;
        }
    }
}
