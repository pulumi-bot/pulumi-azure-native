// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ApiManagement.V20191201Preview.Inputs
{

    /// <summary>
    /// Identity properties of the Api Management service resource.
    /// </summary>
    public sealed class ApiManagementServiceIdentityArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The type of identity used for the resource. The type 'SystemAssigned, UserAssigned' includes both an implicitly created identity and a set of user assigned identities. The type 'None' will remove any identities from the service.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        [Input("userAssignedIdentities")]
        private InputMap<Inputs.UserIdentityPropertiesArgs>? _userAssignedIdentities;

        /// <summary>
        /// The list of user identities associated with the resource. The user identity 
        /// dictionary key references will be ARM resource ids in the form: 
        /// '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/
        ///     providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}'.
        /// </summary>
        public InputMap<Inputs.UserIdentityPropertiesArgs> UserAssignedIdentities
        {
            get => _userAssignedIdentities ?? (_userAssignedIdentities = new InputMap<Inputs.UserIdentityPropertiesArgs>());
            set => _userAssignedIdentities = value;
        }

        public ApiManagementServiceIdentityArgs()
        {
        }
    }
}
