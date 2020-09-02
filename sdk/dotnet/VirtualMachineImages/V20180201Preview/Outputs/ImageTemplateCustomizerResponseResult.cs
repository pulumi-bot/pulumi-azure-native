// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.VirtualMachineImages.V20180201Preview.Outputs
{

    [OutputType]
    public sealed class ImageTemplateCustomizerResponseResult
    {
        /// <summary>
        /// Friendly Name to provide context on what this customization step does
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The type of customization tool you want to use on the Image. For example, "shell" can be shellCustomizer
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private ImageTemplateCustomizerResponseResult(
            string? name,

            string type)
        {
            Name = name;
            Type = type;
        }
    }
}
