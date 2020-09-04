// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.VirtualMachineImages.V20180201Preview.Inputs
{

    public sealed class ImageTemplateSourceArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the type of source image you want to start with.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public ImageTemplateSourceArgs()
        {
        }
    }
}
