// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Synapse.V20190601Preview.Inputs
{

    /// <summary>
    /// The entity reference.
    /// </summary>
    public sealed class EntityReferenceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of this referenced entity.
        /// </summary>
        [Input("referenceName")]
        public Input<string>? ReferenceName { get; set; }

        /// <summary>
        /// The type of this referenced entity.
        /// </summary>
        [Input("type")]
        public InputUnion<string, Pulumi.AzureNextGen.Synapse.V20190601Preview.IntegrationRuntimeEntityReferenceType>? Type { get; set; }

        public EntityReferenceArgs()
        {
        }
    }
}
