// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20190601.Inputs
{

    /// <summary>
    /// IP configuration profile child resource.
    /// </summary>
    public sealed class IPConfigurationProfileArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// Resource ID.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// The name of the resource. This name can be used to access the resource.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Properties of the IP configuration profile.
        /// </summary>
        [Input("properties")]
        public Input<Inputs.IPConfigurationProfilePropertiesFormatArgs>? Properties { get; set; }

        public IPConfigurationProfileArgs()
        {
        }
    }
}