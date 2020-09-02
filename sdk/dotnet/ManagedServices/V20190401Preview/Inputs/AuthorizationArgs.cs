// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ManagedServices.V20190401Preview.Inputs
{

    /// <summary>
    /// Authorization tuple containing principal Id (of user/service principal/security group) and role definition id.
    /// </summary>
    public sealed class AuthorizationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Principal Id of the security group/service principal/user that would be assigned permissions to the projected subscription
        /// </summary>
        [Input("principalId", required: true)]
        public Input<string> PrincipalId { get; set; } = null!;

        /// <summary>
        /// The role definition identifier. This role will define all the permissions that the security group/service principal/user must have on the projected subscription. This role cannot be an owner role.
        /// </summary>
        [Input("roleDefinitionId", required: true)]
        public Input<string> RoleDefinitionId { get; set; } = null!;

        public AuthorizationArgs()
        {
        }
    }
}
