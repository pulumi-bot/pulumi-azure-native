// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.RecoveryServices.Latest.Inputs
{

    /// <summary>
    /// Policy creation properties.
    /// </summary>
    public sealed class CreatePolicyInputPropertiesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ReplicationProviderSettings.
        /// </summary>
        [Input("providerSpecificInput")]
        public InputUnion<Inputs.A2APolicyCreationInputArgs, InputUnion<Inputs.HyperVReplicaAzurePolicyInputArgs, InputUnion<Inputs.HyperVReplicaBluePolicyInputArgs, InputUnion<Inputs.HyperVReplicaPolicyInputArgs, InputUnion<Inputs.InMageAzureV2PolicyInputArgs, InputUnion<Inputs.InMagePolicyInputArgs, InputUnion<Inputs.InMageRcmPolicyCreationInputArgs, Inputs.VMwareCbtPolicyCreationInputArgs>>>>>>>? ProviderSpecificInput { get; set; }

        public CreatePolicyInputPropertiesArgs()
        {
        }
    }
}
