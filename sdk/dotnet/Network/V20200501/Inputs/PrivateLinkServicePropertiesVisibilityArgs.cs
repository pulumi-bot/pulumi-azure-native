// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20200501.Inputs
{

    /// <summary>
    /// The visibility list of the private link service.
    /// </summary>
    public sealed class PrivateLinkServicePropertiesVisibilityArgs : Pulumi.ResourceArgs
    {
        [Input("subscriptions")]
        private InputList<string>? _subscriptions;

        /// <summary>
        /// The list of subscriptions.
        /// </summary>
        public InputList<string> Subscriptions
        {
            get => _subscriptions ?? (_subscriptions = new InputList<string>());
            set => _subscriptions = value;
        }

        public PrivateLinkServicePropertiesVisibilityArgs()
        {
        }
    }
}