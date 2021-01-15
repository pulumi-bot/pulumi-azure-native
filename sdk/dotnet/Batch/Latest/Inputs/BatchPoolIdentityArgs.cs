// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Batch.Latest.Inputs
{

    /// <summary>
    /// The identity of the Batch pool, if configured. If the pool identity is updated during update an existing pool, only the new vms which are created after the pool shrinks to 0 will have the updated identities
    /// </summary>
    public sealed class BatchPoolIdentityArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The type of identity used for the Batch Pool.
        /// </summary>
        [Input("type", required: true)]
        public Input<Pulumi.AzureNextGen.Batch.Latest.PoolIdentityType> Type { get; set; } = null!;

        [Input("userAssignedIdentities")]
        private InputMap<object>? _userAssignedIdentities;

        /// <summary>
        /// The list of user identities associated with the Batch pool. The user identity dictionary key references will be ARM resource ids in the form: '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}'.
        /// </summary>
        public InputMap<object> UserAssignedIdentities
        {
            get => _userAssignedIdentities ?? (_userAssignedIdentities = new InputMap<object>());
            set => _userAssignedIdentities = value;
        }

        public BatchPoolIdentityArgs()
        {
        }
    }
}
