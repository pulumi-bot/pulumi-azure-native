// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.Media.V20180601Preview.Inputs
{

    /// <summary>
    /// Describes the properties of a rectangular window applied to the input media before processing it.
    /// </summary>
    public sealed class RectangleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The height of the rectangular region in pixels. This can be absolute pixel value (e.g 100), or relative to the size of the video (For example, 50%).
        /// </summary>
        [Input("height")]
        public Input<string>? Height { get; set; }

        /// <summary>
        /// The number of pixels from the left-margin. This can be absolute pixel value (e.g 100), or relative to the size of the video (For example, 50%).
        /// </summary>
        [Input("left")]
        public Input<string>? Left { get; set; }

        /// <summary>
        /// The number of pixels from the top-margin. This can be absolute pixel value (e.g 100), or relative to the size of the video (For example, 50%).
        /// </summary>
        [Input("top")]
        public Input<string>? Top { get; set; }

        /// <summary>
        /// The width of the rectangular region in pixels. This can be absolute pixel value (e.g 100), or relative to the size of the video (For example, 50%).
        /// </summary>
        [Input("width")]
        public Input<string>? Width { get; set; }

        public RectangleArgs()
        {
        }
    }
}
