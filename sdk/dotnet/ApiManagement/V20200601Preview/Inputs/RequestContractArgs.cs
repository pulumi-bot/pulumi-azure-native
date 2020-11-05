// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.ApiManagement.V20200601Preview.Inputs
{

    /// <summary>
    /// Operation request details.
    /// </summary>
    public sealed class RequestContractArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Operation request description.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("headers")]
        private InputList<Inputs.ParameterContractArgs>? _headers;

        /// <summary>
        /// Collection of operation request headers.
        /// </summary>
        public InputList<Inputs.ParameterContractArgs> Headers
        {
            get => _headers ?? (_headers = new InputList<Inputs.ParameterContractArgs>());
            set => _headers = value;
        }

        [Input("queryParameters")]
        private InputList<Inputs.ParameterContractArgs>? _queryParameters;

        /// <summary>
        /// Collection of operation request query parameters.
        /// </summary>
        public InputList<Inputs.ParameterContractArgs> QueryParameters
        {
            get => _queryParameters ?? (_queryParameters = new InputList<Inputs.ParameterContractArgs>());
            set => _queryParameters = value;
        }

        [Input("representations")]
        private InputList<Inputs.RepresentationContractArgs>? _representations;

        /// <summary>
        /// Collection of operation request representations.
        /// </summary>
        public InputList<Inputs.RepresentationContractArgs> Representations
        {
            get => _representations ?? (_representations = new InputList<Inputs.RepresentationContractArgs>());
            set => _representations = value;
        }

        public RequestContractArgs()
        {
        }
    }
}
