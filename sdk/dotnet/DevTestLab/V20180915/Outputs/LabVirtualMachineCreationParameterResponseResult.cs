// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.DevTestLab.V20180915.Outputs
{

    [OutputType]
    public sealed class LabVirtualMachineCreationParameterResponseResult
    {
        /// <summary>
        /// The location of the new virtual machine or environment
        /// </summary>
        public readonly string? Location;
        /// <summary>
        /// The name of the virtual machine or environment
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The properties of the resource.
        /// </summary>
        public readonly Outputs.LabVirtualMachineCreationParameterPropertiesResponseResult? Properties;
        /// <summary>
        /// The tags of the resource.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;

        [OutputConstructor]
        private LabVirtualMachineCreationParameterResponseResult(
            string? location,

            string? name,

            Outputs.LabVirtualMachineCreationParameterPropertiesResponseResult? properties,

            ImmutableDictionary<string, string>? tags)
        {
            Location = location;
            Name = name;
            Properties = properties;
            Tags = tags;
        }
    }
}