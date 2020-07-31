// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ApiManagement.V20170301.Inputs
{

    /// <summary>
    /// Parameters supplied to the Create TagDescription operation.
    /// </summary>
    public sealed class TagDescriptionBasePropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description of the Tag.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Description of the external resources describing the tag.
        /// </summary>
        [Input("externalDocsDescription")]
        public Input<string>? ExternalDocsDescription { get; set; }

        /// <summary>
        /// Absolute URL of external resources describing the tag.
        /// </summary>
        [Input("externalDocsUrl")]
        public Input<string>? ExternalDocsUrl { get; set; }

        public TagDescriptionBasePropertiesArgs()
        {
        }
    }
}