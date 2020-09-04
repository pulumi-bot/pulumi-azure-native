// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DocumentDB.V20200601Preview.Inputs
{

    public sealed class CompositePathArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Sort order for composite paths.
        /// </summary>
        [Input("order")]
        public Input<string>? Order { get; set; }

        /// <summary>
        /// The path for which the indexing behavior applies to. Index paths typically start with root and end with wildcard (/path/*)
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        public CompositePathArgs()
        {
        }
    }
}
