// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Migrate.V20180901Preview.Inputs
{

    /// <summary>
    /// Gets or sets the tags.
    /// </summary>
    public sealed class MigrateProjectTagsArgs : Pulumi.ResourceArgs
    {
        [Input("additionalProperties")]
        public Input<string>? AdditionalProperties { get; set; }

        public MigrateProjectTagsArgs()
        {
        }
    }
}
