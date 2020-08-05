// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20181101.Inputs
{

    /// <summary>
    /// Details the service to which the subnet is delegated.
    /// </summary>
    public sealed class DelegationResponseArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        [Input("etag")]
        public string? Etag { get; set; }

        /// <summary>
        /// Resource ID.
        /// </summary>
        [Input("id")]
        public string? Id { get; set; }

        /// <summary>
        /// The name of the resource that is unique within a subnet. This name can be used to access the resource.
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// Properties of the subnet.
        /// </summary>
        [Input("properties")]
        public Inputs.ServiceDelegationPropertiesFormatResponseArgs? Properties { get; set; }

        public DelegationResponseArgs()
        {
        }
    }
}
