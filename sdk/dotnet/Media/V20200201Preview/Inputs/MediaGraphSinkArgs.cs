// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Media.V20200201Preview.Inputs
{

    /// <summary>
    /// Media Sink.
    /// </summary>
    public sealed class MediaGraphSinkArgs : Pulumi.ResourceArgs
    {
        [Input("inputs", required: true)]
        private InputList<string>? _inputs;

        /// <summary>
        /// Sink inputs.
        /// </summary>
        public InputList<string> Inputs
        {
            get => _inputs ?? (_inputs = new InputList<string>());
            set => _inputs = value;
        }

        /// <summary>
        /// Sink name.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The discriminator for derived types.
        /// </summary>
        [Input("odataType", required: true)]
        public Input<string> OdataType { get; set; } = null!;

        public MediaGraphSinkArgs()
        {
        }
    }
}
