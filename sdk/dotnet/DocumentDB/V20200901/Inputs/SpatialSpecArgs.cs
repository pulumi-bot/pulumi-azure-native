// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.DocumentDB.V20200901.Inputs
{

    public sealed class SpatialSpecArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The path for which the indexing behavior applies to. Index paths typically start with root and end with wildcard (/path/*)
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        [Input("types")]
        private InputList<string>? _types;

        /// <summary>
        /// List of path's spatial type
        /// </summary>
        public InputList<string> Types
        {
            get => _types ?? (_types = new InputList<string>());
            set => _types = value;
        }

        public SpatialSpecArgs()
        {
        }
    }
}
