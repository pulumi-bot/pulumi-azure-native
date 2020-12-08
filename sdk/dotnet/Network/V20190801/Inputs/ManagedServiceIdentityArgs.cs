// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Network.V20190801.Inputs
{

    /// <summary>
    /// Identity for the resource.
    /// </summary>
    public sealed class ManagedServiceIdentityArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The type of identity used for the resource. The type 'SystemAssigned, UserAssigned' includes both an implicitly created identity and a set of user assigned identities. The type 'None' will remove any identities from the virtual machine.
        /// </summary>
        [Input("type")]
        public Input<Pulumi.AzureNextGen.Network.V20190801.ResourceIdentityType>? Type { get; set; }

        [Input("userAssignedIdentities")]
        private InputMap<object>? _userAssignedIdentities;

        /// <summary>
        /// The list of user identities associated with resource. The user identity dictionary key references will be ARM resource ids in the form: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}'.
        /// </summary>
        public InputMap<object> UserAssignedIdentities
        {
            get => _userAssignedIdentities ?? (_userAssignedIdentities = new InputMap<object>());
            set => _userAssignedIdentities = value;
        }

        public ManagedServiceIdentityArgs()
        {
        }
    }
}
