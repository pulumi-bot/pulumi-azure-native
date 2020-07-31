// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Network.V20160601.Inputs
{

    /// <summary>
    /// Network security rule
    /// </summary>
    public sealed class SecurityRuleResponseArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// A unique read-only string that changes whenever the resource is updated
        /// </summary>
        [Input("etag")]
        public string? Etag { get; set; }

        /// <summary>
        /// Resource Id
        /// </summary>
        [Input("id")]
        public string? Id { get; set; }

        /// <summary>
        /// Gets name of the resource that is unique within a resource group. This name can be used to access the resource
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        [Input("properties")]
        public Inputs.SecurityRulePropertiesFormatResponseArgs? Properties { get; set; }

        public SecurityRuleResponseArgs()
        {
        }
    }
}