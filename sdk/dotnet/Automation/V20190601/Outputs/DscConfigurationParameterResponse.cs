// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AzureNextGen.Automation.V20190601.Outputs
{

    [OutputType]
    public sealed class DscConfigurationParameterResponse
    {
        /// <summary>
        /// Gets or sets the default value of parameter.
        /// </summary>
        public readonly string? DefaultValue;
        /// <summary>
        /// Gets or sets a Boolean value to indicate whether the parameter is mandatory or not.
        /// </summary>
        public readonly bool? IsMandatory;
        /// <summary>
        /// Get or sets the position of the parameter.
        /// </summary>
        public readonly int? Position;
        /// <summary>
        /// Gets or sets the type of the parameter.
        /// </summary>
        public readonly string? Type;

        [OutputConstructor]
        private DscConfigurationParameterResponse(
            string? defaultValue,

            bool? isMandatory,

            int? position,

            string? type)
        {
            DefaultValue = defaultValue;
            IsMandatory = isMandatory;
            Position = position;
            Type = type;
        }
    }
}
