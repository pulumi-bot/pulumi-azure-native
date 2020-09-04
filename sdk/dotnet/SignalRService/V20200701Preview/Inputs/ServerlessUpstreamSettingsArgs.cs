// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.SignalRService.V20200701Preview.Inputs
{

    /// <summary>
    /// The settings for the Upstream when the Azure SignalR is in server-less mode.
    /// </summary>
    public sealed class ServerlessUpstreamSettingsArgs : Pulumi.ResourceArgs
    {
        [Input("templates")]
        private InputList<Inputs.UpstreamTemplateArgs>? _templates;

        /// <summary>
        /// Gets or sets the list of Upstream URL templates. Order matters, and the first matching template takes effects.
        /// </summary>
        public InputList<Inputs.UpstreamTemplateArgs> Templates
        {
            get => _templates ?? (_templates = new InputList<Inputs.UpstreamTemplateArgs>());
            set => _templates = value;
        }

        public ServerlessUpstreamSettingsArgs()
        {
        }
    }
}
