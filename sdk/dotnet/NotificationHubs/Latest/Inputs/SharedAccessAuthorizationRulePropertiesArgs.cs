// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNative.NotificationHubs.Latest.Inputs
{

    /// <summary>
    /// SharedAccessAuthorizationRule properties.
    /// </summary>
    public sealed class SharedAccessAuthorizationRulePropertiesArgs : Pulumi.ResourceArgs
    {
        [Input("rights")]
        private InputList<Pulumi.AzureNative.NotificationHubs.Latest.AccessRights>? _rights;

        /// <summary>
        /// The rights associated with the rule.
        /// </summary>
        public InputList<Pulumi.AzureNative.NotificationHubs.Latest.AccessRights> Rights
        {
            get => _rights ?? (_rights = new InputList<Pulumi.AzureNative.NotificationHubs.Latest.AccessRights>());
            set => _rights = value;
        }

        public SharedAccessAuthorizationRulePropertiesArgs()
        {
        }
    }
}
