// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20170801.Inputs
{

    /// <summary>
    /// Properties of ResourceNavigationLink.
    /// </summary>
    public sealed class ResourceNavigationLinkFormatResponseArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// Link to the external resource
        /// </summary>
        [Input("link")]
        public string? Link { get; set; }

        /// <summary>
        /// Resource type of the linked resource.
        /// </summary>
        [Input("linkedResourceType")]
        public string? LinkedResourceType { get; set; }

        /// <summary>
        /// Provisioning state of the ResourceNavigationLink resource.
        /// </summary>
        [Input("provisioningState", required: true)]
        public string ProvisioningState { get; set; } = null!;

        public ResourceNavigationLinkFormatResponseArgs()
        {
        }
    }
}