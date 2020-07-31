// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Automation.V20151031.Inputs
{

    /// <summary>
    /// The parameters supplied to the create or update module properties.
    /// </summary>
    public sealed class ModuleCreateOrUpdatePropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Gets or sets the module content link.
        /// </summary>
        [Input("contentLink", required: true)]
        public Input<Inputs.ContentLinkArgs> ContentLink { get; set; } = null!;

        public ModuleCreateOrUpdatePropertiesArgs()
        {
        }
    }
}