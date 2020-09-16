// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.RecoveryServices.V20180710.Inputs
{

    /// <summary>
    /// InMageRcm fabric provider specific settings.
    /// </summary>
    public sealed class InMageRcmFabricCreationInputArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The certificate to be used for AAD authentication.
        /// </summary>
        [Input("authCertificate")]
        public Input<string>? AuthCertificate { get; set; }

        /// <summary>
        /// Gets the class type.
        /// </summary>
        [Input("instanceType")]
        public Input<string>? InstanceType { get; set; }

        /// <summary>
        /// The ARM Id of the physical site.
        /// </summary>
        [Input("physicalSiteId")]
        public Input<string>? PhysicalSiteId { get; set; }

        /// <summary>
        /// The identity provider input for source agent authentication.
        /// </summary>
        [Input("sourceAgentIdentity")]
        public Input<Inputs.IdentityProviderInputArgs>? SourceAgentIdentity { get; set; }

        /// <summary>
        /// The ARM Id of the VMware site.
        /// </summary>
        [Input("vmwareSiteId")]
        public Input<string>? VmwareSiteId { get; set; }

        public InMageRcmFabricCreationInputArgs()
        {
        }
    }
}
