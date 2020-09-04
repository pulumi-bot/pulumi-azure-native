// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ContainerService.V20180801Preview.Inputs
{

    /// <summary>
    /// SSH configuration for Linux-based VMs running on Azure.
    /// </summary>
    public sealed class ContainerServiceSshConfigurationArgs : Pulumi.ResourceArgs
    {
        [Input("publicKeys", required: true)]
        private InputList<Inputs.ContainerServiceSshPublicKeyArgs>? _publicKeys;

        /// <summary>
        /// The list of SSH public keys used to authenticate with Linux-based VMs. Only expect one key specified.
        /// </summary>
        public InputList<Inputs.ContainerServiceSshPublicKeyArgs> PublicKeys
        {
            get => _publicKeys ?? (_publicKeys = new InputList<Inputs.ContainerServiceSshPublicKeyArgs>());
            set => _publicKeys = value;
        }

        public ContainerServiceSshConfigurationArgs()
        {
        }
    }
}
