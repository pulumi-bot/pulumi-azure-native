// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ServiceFabric.V20191101Preview.Inputs
{

    public sealed class ApplicationUserAssignedIdentityArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The friendly name of user assigned identity.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The principal id of user assigned identity.
        /// </summary>
        [Input("principalId", required: true)]
        public Input<string> PrincipalId { get; set; } = null!;

        public ApplicationUserAssignedIdentityArgs()
        {
        }
    }
}
