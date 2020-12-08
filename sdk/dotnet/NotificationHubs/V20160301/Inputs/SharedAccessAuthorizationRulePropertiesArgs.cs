// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.NotificationHubs.V20160301.Inputs
{

    /// <summary>
    /// SharedAccessAuthorizationRule properties.
    /// </summary>
    public sealed class SharedAccessAuthorizationRulePropertiesArgs : Pulumi.ResourceArgs
    {
        [Input("rights")]
        private InputList<Pulumi.AzureNextGen.NotificationHubs.V20160301.AccessRights>? _rights;

        /// <summary>
        /// The rights associated with the rule.
        /// </summary>
        public InputList<Pulumi.AzureNextGen.NotificationHubs.V20160301.AccessRights> Rights
        {
            get => _rights ?? (_rights = new InputList<Pulumi.AzureNextGen.NotificationHubs.V20160301.AccessRights>());
            set => _rights = value;
        }

        public SharedAccessAuthorizationRulePropertiesArgs()
        {
        }
    }
}
