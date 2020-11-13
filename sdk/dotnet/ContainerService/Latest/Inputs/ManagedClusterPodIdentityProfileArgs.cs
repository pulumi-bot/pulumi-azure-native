// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.ContainerService.Latest.Inputs
{

    public sealed class ManagedClusterPodIdentityProfileArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether the pod identity addon is enabled.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("userAssignedIdentities")]
        private InputList<Inputs.ManagedClusterPodIdentityArgs>? _userAssignedIdentities;

        /// <summary>
        /// User assigned pod identity settings.
        /// </summary>
        public InputList<Inputs.ManagedClusterPodIdentityArgs> UserAssignedIdentities
        {
            get => _userAssignedIdentities ?? (_userAssignedIdentities = new InputList<Inputs.ManagedClusterPodIdentityArgs>());
            set => _userAssignedIdentities = value;
        }

        [Input("userAssignedIdentityExceptions")]
        private InputList<Inputs.ManagedClusterPodIdentityExceptionArgs>? _userAssignedIdentityExceptions;

        /// <summary>
        /// User assigned pod identity exception settings.
        /// </summary>
        public InputList<Inputs.ManagedClusterPodIdentityExceptionArgs> UserAssignedIdentityExceptions
        {
            get => _userAssignedIdentityExceptions ?? (_userAssignedIdentityExceptions = new InputList<Inputs.ManagedClusterPodIdentityExceptionArgs>());
            set => _userAssignedIdentityExceptions = value;
        }

        public ManagedClusterPodIdentityProfileArgs()
        {
        }
    }
}
