// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Web.V20160601.Inputs
{

    /// <summary>
    /// The WSDL definition
    /// </summary>
    public sealed class WsdlDefinitionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The WSDL content
        /// </summary>
        [Input("content")]
        public Input<string>? Content { get; set; }

        /// <summary>
        /// The WSDL import method
        /// </summary>
        [Input("importMethod")]
        public InputUnion<string, Pulumi.AzureNextGen.Web.V20160601.WsdlImportMethod>? ImportMethod { get; set; }

        /// <summary>
        /// The service with name and endpoint names
        /// </summary>
        [Input("service")]
        public Input<Inputs.WsdlServiceArgs>? Service { get; set; }

        /// <summary>
        /// The WSDL URL
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        public WsdlDefinitionArgs()
        {
        }
    }
}
