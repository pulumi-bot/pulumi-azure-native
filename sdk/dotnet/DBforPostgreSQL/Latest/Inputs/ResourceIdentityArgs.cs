// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.DBforPostgreSQL.Latest.Inputs
{

    /// <summary>
    /// Azure Active Directory identity configuration for a resource.
    /// </summary>
    public sealed class ResourceIdentityArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The identity type. Set this to 'SystemAssigned' in order to automatically create and assign an Azure Active Directory principal for the resource.
        /// </summary>
        [Input("type")]
        public InputUnion<string, Pulumi.AzureNextGen.DBforPostgreSQL.Latest.IdentityType>? Type { get; set; }

        public ResourceIdentityArgs()
        {
        }
    }
}
