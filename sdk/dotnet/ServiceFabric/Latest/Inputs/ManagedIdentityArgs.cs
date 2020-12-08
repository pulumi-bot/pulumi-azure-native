// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.ServiceFabric.Latest.Inputs
{

    /// <summary>
    /// Describes the managed identities for an Azure resource.
    /// </summary>
    public sealed class ManagedIdentityArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The type of managed identity for the resource.
        /// </summary>
        [Input("type")]
        public Input<Pulumi.AzureNextGen.ServiceFabric.Latest.ManagedIdentityType>? Type { get; set; }

        [Input("userAssignedIdentities")]
        private InputMap<object>? _userAssignedIdentities;

        /// <summary>
        /// The list of user identities associated with the resource. The user identity dictionary key references will be ARM resource ids in the form:
        /// '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}'.
        /// </summary>
        public InputMap<object> UserAssignedIdentities
        {
            get => _userAssignedIdentities ?? (_userAssignedIdentities = new InputMap<object>());
            set => _userAssignedIdentities = value;
        }

        public ManagedIdentityArgs()
        {
        }
    }
}
