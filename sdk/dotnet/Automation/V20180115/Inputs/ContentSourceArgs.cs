// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Automation.V20180115.Inputs
{

    /// <summary>
    /// Definition of the content source.
    /// </summary>
    public sealed class ContentSourceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Gets or sets the hash.
        /// </summary>
        [Input("hash")]
        public Input<Inputs.ContentHashArgs>? Hash { get; set; }

        /// <summary>
        /// Gets or sets the content source type.
        /// </summary>
        [Input("type")]
        public InputUnion<string, Pulumi.AzureNextGen.Automation.V20180115.ContentSourceType>? Type { get; set; }

        /// <summary>
        /// Gets or sets the value of the content. This is based on the content source type.
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        /// <summary>
        /// Gets or sets the version of the content.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public ContentSourceArgs()
        {
        }
    }
}
