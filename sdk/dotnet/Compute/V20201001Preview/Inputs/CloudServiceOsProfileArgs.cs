// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Compute.V20201001Preview.Inputs
{

    /// <summary>
    /// Describes the OS profile for the cloud service.
    /// </summary>
    public sealed class CloudServiceOsProfileArgs : Pulumi.ResourceArgs
    {
        [Input("secrets")]
        private InputList<Inputs.CloudServiceVaultSecretGroupArgs>? _secrets;

        /// <summary>
        /// Specifies set of certificates that should be installed onto the role instances.
        /// </summary>
        public InputList<Inputs.CloudServiceVaultSecretGroupArgs> Secrets
        {
            get => _secrets ?? (_secrets = new InputList<Inputs.CloudServiceVaultSecretGroupArgs>());
            set => _secrets = value;
        }

        public CloudServiceOsProfileArgs()
        {
        }
    }
}
