// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ApiManagement.V20180601Preview.Inputs
{

    /// <summary>
    /// OAuth acquire token request body parameter (www-url-form-encoded).
    /// </summary>
    public sealed class TokenBodyParameterContractArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// body parameter name.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// body parameter value.
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public TokenBodyParameterContractArgs()
        {
        }
    }
}
