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
    /// Describes the role profile for the cloud service.
    /// </summary>
    public sealed class CloudServiceRoleProfileArgs : Pulumi.ResourceArgs
    {
        [Input("roles")]
        private InputList<Inputs.CloudServiceRoleProfilePropertiesArgs>? _roles;

        /// <summary>
        /// List of roles for the cloud service.
        /// </summary>
        public InputList<Inputs.CloudServiceRoleProfilePropertiesArgs> Roles
        {
            get => _roles ?? (_roles = new InputList<Inputs.CloudServiceRoleProfilePropertiesArgs>());
            set => _roles = value;
        }

        public CloudServiceRoleProfileArgs()
        {
        }
    }
}
