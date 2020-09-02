// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DevTestLab.V20150521Preview.Inputs
{

    /// <summary>
    /// Properties for creating a custom image from a VHD.
    /// </summary>
    public sealed class CustomImagePropertiesCustomArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The image name.
        /// </summary>
        [Input("imageName")]
        public Input<string>? ImageName { get; set; }

        /// <summary>
        /// Indicates whether sysprep has been run on the VHD.
        /// </summary>
        [Input("sysPrep")]
        public Input<bool>? SysPrep { get; set; }

        public CustomImagePropertiesCustomArgs()
        {
        }
    }
}
