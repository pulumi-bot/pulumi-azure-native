// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.EventGrid.V20200401Preview.Inputs
{

    /// <summary>
    /// The identity information for the resource.
    /// </summary>
    public sealed class IdentityInfoArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The principal ID of resource identity.
        /// </summary>
        [Input("principalId")]
        public Input<string>? PrincipalId { get; set; }

        /// <summary>
        /// The tenant ID of resource.
        /// </summary>
        [Input("tenantId")]
        public Input<string>? TenantId { get; set; }

        /// <summary>
        /// The type of managed identity used. The type 'SystemAssigned, UserAssigned' includes both an implicitly created identity and a set of user-assigned identities. The type 'None' will remove any identity.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        [Input("userAssignedIdentities")]
        private InputMap<Inputs.UserIdentityPropertiesArgs>? _userAssignedIdentities;

        /// <summary>
        /// The list of user identities associated with the resource. The user identity dictionary key references will be ARM resource ids in the form:
        /// '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}'.
        /// This property is currently not used and reserved for future usage.
        /// </summary>
        public InputMap<Inputs.UserIdentityPropertiesArgs> UserAssignedIdentities
        {
            get => _userAssignedIdentities ?? (_userAssignedIdentities = new InputMap<Inputs.UserIdentityPropertiesArgs>());
            set => _userAssignedIdentities = value;
        }

        public IdentityInfoArgs()
        {
        }
    }
}
