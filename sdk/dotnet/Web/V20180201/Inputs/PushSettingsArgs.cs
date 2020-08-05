// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Web.V20180201.Inputs
{

    /// <summary>
    /// Push settings for the App.
    /// </summary>
    public sealed class PushSettingsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Gets or sets a JSON string containing a list of dynamic tags that will be evaluated from user claims in the push registration endpoint.
        /// </summary>
        [Input("dynamicTagsJson")]
        public Input<string>? DynamicTagsJson { get; set; }

        /// <summary>
        /// Gets or sets a flag indicating whether the Push endpoint is enabled.
        /// </summary>
        [Input("isPushEnabled", required: true)]
        public Input<bool> IsPushEnabled { get; set; } = null!;

        /// <summary>
        /// Kind of resource.
        /// </summary>
        [Input("kind")]
        public Input<string>? Kind { get; set; }

        /// <summary>
        /// Gets or sets a JSON string containing a list of tags that are whitelisted for use by the push registration endpoint.
        /// </summary>
        [Input("tagWhitelistJson")]
        public Input<string>? TagWhitelistJson { get; set; }

        /// <summary>
        /// Gets or sets a JSON string containing a list of tags that require user authentication to be used in the push registration endpoint.
        /// Tags can consist of alphanumeric characters and the following:
        /// '_', '@', '#', '.', ':', '-'. 
        /// Validation should be performed at the PushRequestHandler.
        /// </summary>
        [Input("tagsRequiringAuth")]
        public Input<string>? TagsRequiringAuth { get; set; }

        public PushSettingsArgs()
        {
        }
    }
}
