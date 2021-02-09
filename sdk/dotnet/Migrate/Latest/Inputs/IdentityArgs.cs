// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Migrate.Latest.Inputs
{

    /// <summary>
    /// Defines the MSI properties of the Move Collection.
    /// </summary>
    public sealed class IdentityArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Gets or sets the principal id.
        /// </summary>
        [Input("principalId")]
        public Input<string>? PrincipalId { get; set; }

        /// <summary>
        /// Gets or sets the tenant id.
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        /// <summary>
        /// The type of identity used for the resource mover service.
        /// </summary>
        [Input("type")]
        public InputUnion<string, Pulumi.AzureNextGen.Migrate.Latest.ResourceIdentityType>? Type { get; set; }

        public IdentityArgs()
        {
        }
    }
}
