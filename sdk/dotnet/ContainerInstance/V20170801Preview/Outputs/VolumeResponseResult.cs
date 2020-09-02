// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ContainerInstance.V20170801Preview.Outputs
{

    [OutputType]
    public sealed class VolumeResponseResult
    {
        /// <summary>
        /// The name of the Azure File volume.
        /// </summary>
        public readonly Outputs.AzureFileVolumeResponseResult AzureFile;
        /// <summary>
        /// The name of the volume.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private VolumeResponseResult(
            Outputs.AzureFileVolumeResponseResult azureFile,

            string name)
        {
            AzureFile = azureFile;
            Name = name;
        }
    }
}
