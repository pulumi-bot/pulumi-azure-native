// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureRM.ServiceFabric.V20170701Preview.Outputs
{

    [OutputType]
    public sealed class SettingsSectionDescriptionResponseResult
    {
        /// <summary>
        /// The section name of the fabric settings.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The collection of parameters in the section.
        /// </summary>
        public readonly ImmutableArray<Outputs.SettingsParameterDescriptionResponseResult> Parameters;

        [OutputConstructor]
        private SettingsSectionDescriptionResponseResult(
            string name,

            ImmutableArray<Outputs.SettingsParameterDescriptionResponseResult> parameters)
        {
            Name = name;
            Parameters = parameters;
        }
    }
}
