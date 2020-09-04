// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ServiceFabricMesh.V20180701Preview.Outputs
{

    [OutputType]
    public sealed class SettingResponseResult
    {
        /// <summary>
        /// The name of the setting.
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// The value of the setting.
        /// </summary>
        public readonly string? Value;

        [OutputConstructor]
        private SettingResponseResult(
            string? name,

            string? value)
        {
            Name = name;
            Value = value;
        }
    }
}
