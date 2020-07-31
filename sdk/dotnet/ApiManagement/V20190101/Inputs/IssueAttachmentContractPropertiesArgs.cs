// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ApiManagement.V20190101.Inputs
{

    /// <summary>
    /// Issue Attachment contract Properties.
    /// </summary>
    public sealed class IssueAttachmentContractPropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// An HTTP link or Base64-encoded binary data.
        /// </summary>
        [Input("content", required: true)]
        public Input<string> Content { get; set; } = null!;

        /// <summary>
        /// Either 'link' if content is provided via an HTTP link or the MIME type of the Base64-encoded binary data provided in the 'content' property.
        /// </summary>
        [Input("contentFormat", required: true)]
        public Input<string> ContentFormat { get; set; } = null!;

        /// <summary>
        /// Filename by which the binary data will be saved.
        /// </summary>
        [Input("title", required: true)]
        public Input<string> Title { get; set; } = null!;

        public IssueAttachmentContractPropertiesArgs()
        {
        }
    }
}