// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Sql.V20180601Preview.Inputs
{

    public sealed class PrivateEndpointPropertyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Resource id of the private endpoint.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        public PrivateEndpointPropertyArgs()
        {
        }
    }
}
